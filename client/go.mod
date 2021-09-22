module github.com/detecc/detecc-cpu/client

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/detecc/deteccted v1.0.0
	github.com/shirou/gopsutil v3.21.8+incompatible
	github.com/tklauser/go-sysconf v0.3.9 // indirect
)

replace github.com/detecc/deteccted => ../../deteccted
