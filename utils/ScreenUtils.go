package utils

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

// GetMouseMonitor 返回鼠标所在的显示器编号（从 1 开始），如果不在任何显示器范围内返回 -1
func GetMouseMonitor() int {
	// 初始化 GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("无法初始化 GLFW:", err)
	}
	defer glfw.Terminate()
	// 设置窗口为隐藏
	glfw.WindowHint(glfw.Visible, glfw.False)
	// 创建一个窗口作为上下文
	window, err := glfw.CreateWindow(1, 1, "Hidden Window", nil, nil)
	if err != nil {
		log.Fatalln("无法创建隐藏窗口:", err)
	}
	defer window.Destroy()

	// 设置当前窗口上下文
	window.MakeContextCurrent()

	// 获取鼠标位置（相对于窗口）
	x, y := window.GetCursorPos()

	// 获取所有显示器
	monitors := glfw.GetMonitors()
	if monitors == nil {
		log.Fatalln("无法获取显示器列表")
	}

	// 遍历每个显示器，检查鼠标是否在其范围内
	for i, monitor := range monitors {
		// 获取显示器的工作区
		vidMode := monitor.GetVideoMode()
		xPos, yPos := monitor.GetPos()
		width, height := vidMode.Width, vidMode.Height

		// 检查鼠标位置是否在当前显示器范围内
		if int(x) >= xPos && int(x) < xPos+width &&
			int(y) >= yPos && int(y) < yPos+height {
			return i // 返回显示器编号（从 0 开始）
		}
	}

	return -1 // 如果鼠标不在任何显示器范围内，返回 -1
}
