package functions

import (
	"ScrCat/utils"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

// ScreenShotForFull 全屏截图
func ScreenShotForFull() {
	//获取活跃显示器设备的数量
	monitor := utils.GetMouseMonitor()
	bounds := screenshot.GetDisplayBounds(monitor)
	img, _ := screenshot.CaptureRect(bounds)
	file, _ := os.Create("screenshot.png")
	defer file.Close()
	png.Encode(file, img)
}
