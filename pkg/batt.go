package pkg

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Battery struct {
	Name    string
	Voltage float64
	Current float64
	Power   float64
	Percent float64
}

func readFloat(path string) (float64, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return -1, err
	}

	return strconv.ParseFloat(strings.TrimSpace(string(file)), 64)
}

func GetBatts() ([]Battery, error) {
	psuPath := "/sys/class/power_supply"
	supplies, err := os.ReadDir(psuPath)
	if err != nil {
		return nil, err
	}

	batteries := make([]Battery, 0)
	for _, sup := range supplies {
		if !strings.HasPrefix(sup.Name(), "BAT") {
			continue
		}

		batt := Battery{
			Name:    sup.Name(),
			Percent: -1,
			Power:   -1,
			Voltage: -1,
			Current: -1,
		}

		battPath := filepath.Join(psuPath, sup.Name())
		points, err := os.ReadDir(battPath)
		if err != nil {
			return nil, err
		}

		for _, point := range points {
			name := point.Name()
			path := filepath.Join(battPath, name)

			switch name {
			case "voltage_now":
				batt.Voltage, err = readFloat(path)
				batt.Voltage /= 1000000
			case "power_now":
				batt.Power, err = readFloat(path)
				batt.Power /= 1000000
			case "current_now":
				batt.Current, err = readFloat(path)
				batt.Current /= 1000000
			case "capacity":
				batt.Percent, err = readFloat(path)
			}
		}

		if batt.Power == -1 {
			if batt.Voltage != -1 && batt.Current != -1 {
				batt.Power = batt.Voltage * batt.Current
			}
		}

		batteries = append(batteries, batt)
	}

	return batteries, nil
}
