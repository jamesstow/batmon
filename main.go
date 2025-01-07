package main

import (
	"fmt"
	"jamesstow.co.uk/batmon/pkg"
)

func main() {
	batts, err := pkg.GetBatts()
	if err != nil {
		panic(err)
	}

	for _, batt := range batts {
		line := fmt.Sprintf("%s: %fW, %f%%", batt.Name, batt.Power, batt.Percent)
		fmt.Println(line)
	}
}
