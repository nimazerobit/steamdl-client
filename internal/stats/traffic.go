package stats

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"steam-lancache/internal/config"
)

var (
	bytesReceived int64
	mutex         sync.Mutex
)

func Load() {
	data, err := os.ReadFile(config.StatsFile)
	if err == nil {
		val, _ := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
		bytesReceived = val
	}
}

func AddBytes(n int64) {
	mutex.Lock()
	defer mutex.Unlock()
	bytesReceived += n
}

func StartSaver() {
	ticker := time.NewTicker(config.StatsUpdateFreq)
	for range ticker.C {
		mutex.Lock()
		currentBytes := bytesReceived
		mutex.Unlock()

		err := os.WriteFile(config.StatsFile, []byte(fmt.Sprintf("%d", currentBytes)), 0644)
		if err != nil {
			log.Printf("[Stats] Error saving: %v", err)
		}
	}
}
