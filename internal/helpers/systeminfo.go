package helpers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return "unknown"
}

func checkInternet() (bool, time.Duration) {
	url := "https://www.google.com/generate_204"
	start := time.Now()

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false, 0
	}
	resp.Body.Close()

	return true, time.Since(start)
}

func ShowSystemInfo(is_print bool) (
	osName string,
	arch string,
	cpuCount string,
	localIP string,
	hasInternet bool,
	ping time.Duration,
) {
	osName = runtime.GOOS
	arch = runtime.GOARCH
	cpuCount = strconv.Itoa(runtime.NumCPU())
	localIP = getLocalIP()
	hasInternet, ping = checkInternet()
	if is_print {
		fmt.Println("-----------------------------------------")
		log.Println("[*] Operating System:", osName)
		log.Println("[*] Architecture:", arch)
		log.Println("[*] Number of CPUs:", cpuCount)
		log.Println("[*] Local IP:", localIP)
		if hasInternet {
			log.Printf("[*] Internet: yes (ping ~ %v)\n", ping)
		} else {
			log.Println("[*] Internet: no")
		}
		fmt.Println("-----------------------------------------")
	}
	return
}
