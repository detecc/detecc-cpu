package main

import (
	"fmt"
	"github.com/detecc/deteccted/plugin"
	cpp "github.com/shirou/gopsutil/cpu"
	mem2 "github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

func init() {
	example := &Example{}
	plugin.Register(example.GetCmdName(), example)

	go example.Schedule(LogCpuMemUsage, 5 * time.Second)
}

type Example struct {
	plugin.Handler
}

func (e Example) Inject() {

}

func (e Example) GetCmdName() string {
	return "/example"
}

func (e Example) Execute(args interface{}) (interface{}, error) {
	log.Println(args)
	return args, nil
}

func (e Example) Schedule(f func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		f()
	}
	ticker.Stop()
}

func LogCpuMemUsage() {
	percent, _ := cpp.Percent(time.Second, false)
	mem, _ := mem2.VirtualMemory()
	fmt.Printf("       CPU Usage: %.2f\n", percent)
	fmt.Printf("Memory Available: %.2d\n", mem.Available/1024/1024)
	fmt.Printf("    Memory Total: %.2d\n", mem.Total/1024/1024)
	fmt.Printf("     Memory Used: %.2d\n", mem.Used/1024/1024)
}
