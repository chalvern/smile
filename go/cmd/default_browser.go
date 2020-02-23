package cmd

import (
	"os/exec"
	"syscall"
)

// MacOSDefaultBrowser 在 mac 中使用默认的浏览器打开网页
func MacOSDefaultBrowser() {
	exec.Command(`open`, `https://jingwei.link`).Start()
}

// LinuxDefaultBrowser 在 linux 中
func LinuxDefaultBrowser() {
	exec.Command(`xdg-open`, `https://jingwei.link`).Start()
}

// WindowsOSDefaultBrowser 在 windows 中使用默认的浏览器打开网页
func WindowsOSDefaultBrowser() {
	// 有GUI调用
	exec.Command(`cmd`, `/c`, `start`, `https://jingwei.link`).Start()

	// 无GUI调用
	cmd := exec.Command(`cmd`, `/c`, `start`, `https://jingwei.link`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
