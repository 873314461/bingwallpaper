package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func setWallpaper(wallpaperPath string) error {
	wallpaperPath, err := filepath.Abs(wallpaperPath)
	if err != nil {
		return fmt.Errorf("get abs path [%s] error: %v", wallpaperPath, err)
	}
	cmd := fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", wallpaperPath)
	err = exec.Command("/usr/bin/osascript", "-e", cmd).Run()
	if err != nil {
		return fmt.Errorf("run command [/usr/bin/osascript -e %s] error: %v", cmd, err)
	}
	return nil
}
