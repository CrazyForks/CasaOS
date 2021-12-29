package v1

import (
	"net/http"
	"strings"
	"time"
	"unsafe"

	"github.com/IceWhaleTech/CasaOS/model"
	oasis_err2 "github.com/IceWhaleTech/CasaOS/pkg/utils/oasis_err"
	"github.com/IceWhaleTech/CasaOS/service"
	"github.com/gin-gonic/gin"
)

// @Summary 获取cpu信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getcpuinfo [get]
func CupInfo(c *gin.Context) {
	//检查参数是否正确
	cpu := service.MyService.ZiMa().GetCpuPercent()
	num := service.MyService.ZiMa().GetCpuCoreNum()
	data := make(map[string]interface{})
	data["percent"] = cpu
	data["num"] = num
	c.JSON(http.StatusOK, model.Result{Success: oasis_err2.SUCCESS, Message: oasis_err2.GetMsg(oasis_err2.SUCCESS), Data: data})

}

// @Summary 获取内存信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getmeminfo [get]
func MemInfo(c *gin.Context) {

	//检查参数是否正确
	mem := service.MyService.ZiMa().GetMemInfo()
	c.JSON(http.StatusOK, model.Result{Success: oasis_err2.SUCCESS, Message: oasis_err2.GetMsg(oasis_err2.SUCCESS), Data: mem})

}

// @Summary 获取硬盘信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getdiskinfo [get]
func DiskInfo(c *gin.Context) {
	disk := service.MyService.ZiMa().GetDiskInfo()
	c.JSON(http.StatusOK, model.Result{Success: oasis_err2.SUCCESS, Message: oasis_err2.GetMsg(oasis_err2.SUCCESS), Data: disk})
}

// @Summary 获取网络信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getnetinfo [get]
func NetInfo(c *gin.Context) {
	netList := service.MyService.ZiMa().GetNetInfo()

	newNet := []model.IOCountersStat{}
	for _, n := range netList {
		for _, netCardName := range service.MyService.ZiMa().GetNet(true) {
			if n.Name == netCardName {
				item := *(*model.IOCountersStat)(unsafe.Pointer(&n))
				item.State = strings.TrimSpace(service.MyService.ZiMa().GetNetState(n.Name))
				item.DateTime = time.Now()
				newNet = append(newNet, item)
				break
			}
		}
	}

	c.JSON(http.StatusOK, model.Result{Success: oasis_err2.SUCCESS, Message: oasis_err2.GetMsg(oasis_err2.SUCCESS), Data: newNet})
}

// @Summary 获取信息系统信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/sysinfo [get]
func SysInfo(c *gin.Context) {
	info := service.MyService.ZiMa().GetSysInfo()
	c.JSON(http.StatusOK, model.Result{Success: oasis_err2.SUCCESS, Message: oasis_err2.GetMsg(oasis_err2.SUCCESS), Data: info})
}
