package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	// cNum, err := cpu.Counts(true)
	// result, err := cpu.PercentWithContext(context.Background(), time.Second*10, true)
	result, err := cpu.Info()
	if err != nil {
		fmt.Println("failed to get cpu counts", err)
		return
	}
	fmt.Println(result)
}