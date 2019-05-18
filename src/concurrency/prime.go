package concurrency

// 参考解释
//https://www.jianshu.com/p/21e7e8c23090


// 生成数
func Generate() chan int {

	ch := make(chan int)

	go func() {

		for i := 2; ; i++ {

			ch <- i

		}

	}()

	return ch
}

// 过滤数据
// 每一轮数据流中最小的那个就是素数，
// 原理是：一个素数不能整除的那个比它自身大的最小的那个数就是素数<比如说素数2可以得出素数3，
// 素数3可以得出素数5， 5求出7...>
func Filter(in chan int, primer int) chan int {

	out := make(chan int)

	go func() {

		for {

			i:= <- in

			if i%primer != 0{

				out <- i

			}
		}

	}()

	return out
}


// 筛选素数
// @todo 在循环的最后一步，把out赋值给ch，这一步很重要，是为了下一步循环做铺垫，所谓的赋值，即是把下一轮的数据流ch指定为本轮的输出流out：本轮输出流流向下轮输入流，如此往复...


func Sieve() chan int {

	out := make(chan int) // 输出管道

	go func() {

		ch := Generate() // 生成数列

		for {

			prime := <- ch // 取出数列

			ch = Filter(ch, prime) //过滤数列

			out <- prime //返回数列

		}

	}()

	return out
}