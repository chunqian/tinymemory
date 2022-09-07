package memory

import unsafe "unsafe"

func PX_ASSERT() {
}
func PX_ERROR(fmt *int8) {
	PX_ASSERT()
}
func PX_GETLOG() *int8 {
	return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
}
func PX_LOG(fmt *int8) {
}
