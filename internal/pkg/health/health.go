package pkg

import (
	"net/http"
	"time"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func CheckApplicationHealth(c *gin.Context) {

	logger := config.GetLogger("Health")

	cpuUsage, err := GetCpuUsage()
	if err != nil {
		logger.Errorf("Failed to get CPU usage: %v", err)
		return
	}

	memoryUsage, err := GetMemoryUsage()
	if err != nil {
		logger.Errorf("Failed to get memory usage: %v", err)
		return
	}

	checkApi, err := checkExternalAPIHealth()
	if err != nil {
		logger.Errorf("Failed to connect api external: %v", err)
		return
	}

	checkDataBase, err := checkDataBaseHealth()
	if err != nil {
		logger.Errorf("Failed initialize database: %v", err)
		return
	}

	if cpuUsage < 80 && memoryUsage < 80 && checkApi && checkDataBase {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
	}

}

func checkDataBaseHealth() (bool, error) {
	err := config.Init()
	if err != nil {
		return false, err
	}

	return true, nil
}
func checkExternalAPIHealth() (bool, error) {
	resp, err := http.Get("https://viacep.com.br/ws/01001000/json/")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	return true, nil
}

func GetCpuUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

func GetMemoryUsage() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return memInfo.UsedPercent, nil
}
