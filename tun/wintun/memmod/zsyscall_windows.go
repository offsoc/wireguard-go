// Code generated by 'go generate'; DO NOT EDIT.

package memmod

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGetProcAddress = modkernel32.NewProc("GetProcAddress")
	procIsBadReadPtr   = modkernel32.NewProc("IsBadReadPtr")
	procLoadLibraryA   = modkernel32.NewProc("LoadLibraryA")
)

func getProcAddress(module windows.Handle, procName *byte) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcAddress.Addr(), 2, uintptr(module), uintptr(unsafe.Pointer(procName)), 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

func isBadReadPtr(addr uintptr, ucb uintptr) (ret bool) {
	r0, _, _ := syscall.Syscall(procIsBadReadPtr.Addr(), 2, uintptr(addr), uintptr(ucb), 0)
	ret = r0 != 0
	return
}

func loadLibraryA(libFileName *byte) (module windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadLibraryA.Addr(), 1, uintptr(unsafe.Pointer(libFileName)), 0, 0)
	module = windows.Handle(r0)
	if module == 0 {
		err = errnoErr(e1)
	}
	return
}
