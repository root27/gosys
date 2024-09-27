package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

func onExit() {

	log.Println("Exit called")

}

func updateMetrics(mCpu, mMemory, diskUsage *systray.MenuItem) {

	cpuInfo, err := cpu.Percent(time.Second, false)

	if err != nil {
		log.Fatal(err)
	}

	memInfo, err := mem.VirtualMemory()

	if err != nil {
		log.Fatal(err)
	}

	dUsage, err := disk.Usage("/")

	if err != nil {
		log.Fatal(err)
	}

	mCpu.SetTitle(fmt.Sprintf("CPU Usage: %.1f%%", cpuInfo[0]))

	mMemory.SetTitle(fmt.Sprintf("Memory Usage: %.1f%%", memInfo.UsedPercent))

	diskUsage.SetTitle(fmt.Sprintf("Disk Usage: %.1f%%", dUsage.UsedPercent))

}

func onReady() {

	systray.SetTitle("System Monitor")

	systray.SetTooltip("System Monitor")

	mCpu := systray.AddMenuItem("CPU Usage", "Show CPU Usage")

	mMemory := systray.AddMenuItem("Memory Usage", "Show Memory Usage")

	diskUsage := systray.AddMenuItem("Disk Usage", "Show Disk Usage")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			updateMetrics(mCpu, mMemory, diskUsage)

			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		<-mQuit.ClickedCh

		systray.Quit()
	}()
}

func main() {

	systray.Run(onReady, onExit)

}
