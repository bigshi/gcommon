package util

import "runtime"

//
// SetHalfCpuProcess
//  @Description: 使用一半cpu来处理协程
//
func SetHalfCpuProcess() {
	maxcpu := runtime.NumCPU()
	halfcpu := int(0.5 * float32(maxcpu))
	if halfcpu < 1 {
		halfcpu = 1
	}
	runtime.GOMAXPROCS(halfcpu)
}

//
// SetAllCpuProcess
//  @Description: 使用全部cpu来处理协程
//
func SetAllCpuProcess() {
	maxcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(maxcpu)
}
