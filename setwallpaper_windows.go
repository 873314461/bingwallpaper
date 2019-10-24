package main

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"
)

func setWallpaper(wallpaperPath string) error {
	wallpaperPath, err := filepath.Abs(wallpaperPath)
	if err != nil {
		return fmt.Errorf("get abs path [%s] error: %v", wallpaperPath, err)
	}

	pathPtr, err := syscall.UTF16PtrFromString(wallpaperPath)
	if err != nil {
		return fmt.Errorf("get string ptr error: %v", err)
	}

	uiAction, uiParam, fWinIni := uint32(SPI_SETDESKWALLPAPER), uint32(0), uint32(SPIF_UPDATEINIFILE)

	user32, err := syscall.LoadLibrary("User32.dll")
	if err != nil {
		return fmt.Errorf("load library error: %v", err)
	}
	defer syscall.FreeLibrary(user32)

	systemParametersInfo, err := syscall.GetProcAddress(user32, "SystemParametersInfoW")
	if err != nil {
		return fmt.Errorf("get proc address error: %v", err)
	}

	r, _, sysErr := syscall.Syscall6(
		uintptr(systemParametersInfo), 4,
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(fWinIni),
		0, 0)
	if r == 0 {
		return fmt.Errorf("system call error: %v", sysErr)
	}
	return nil
}
