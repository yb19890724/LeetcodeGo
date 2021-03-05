package _struct

// 发现差异了吗？我们只是添加了一个检查，以查看指针（r）是否为nil， 如果是，则返回return nil。
// 当我将上面的代码中，谁是担心呼叫的人GetDevice()在nil被抓他的头。

//
type Robot struct {
	devices []*device
}

//
type device struct {
	Name string
}

//
func newRobot() *Robot {
	return nil
}

// 原始代码：
/*func (r *Robot) GetDevice(name string) *device {
	for _, device := range r.devices {
		if device.Name == name {
			return device
		}
	}
	return nil
}*/

// 优化版本
func (r *Robot) device(name string) *device {
	if r == nil {
		return nil
	}
	for _, device := range r.devices {
		if device.Name == name {
			return device
		}
	}
	return nil
}

// nil指针具有类型的事实有两个重要含义，
// 第一个含义是，您可以很好地链接方法，而无需在调用方站点检查返回的值是否为nil。
// 第二个原因是您的方法应该预期会在nil指针上被调用，并且应该正确处理此类情况。
func TestCheckReceiver() {
	newRobot().device("")
}
