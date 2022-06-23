package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	chromePath := "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	// currentDir, _ := os.Getwd()
	// dir := filepath.Join(currentDir, ".cache")
	// ui, _ := lorca.New("http://baidu.com", dir, 800, 600)
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-chSignal:

		// case <-ui.Done():
	}
	ui.Close()
}
