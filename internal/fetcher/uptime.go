package fetcher

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Time struct {
	Days    int
	Seconds int
	Minutes int
	Hours   int
}

func (t *Time) String() string {
	if t.Days == 0 && t.Hours == 0 && t.Minutes == 0 {
		return fmt.Sprintf("%ds", t.Seconds)
	} else if t.Days > 0 {
		return fmt.Sprintf("%dd %dh %dm", t.Days, t.Hours, t.Minutes)
	} else if t.Hours > 0 {
		return fmt.Sprintf("%dh %dm", t.Hours, t.Minutes)
	}
	return fmt.Sprintf("%dm", t.Minutes)
}

func GetUptime(filepath string) (*Time, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, fmt.Errorf("file is empty")
	}
	fields := strings.Fields(scanner.Text())

	uptime, err := convertToInt(fields[0])
	if err != nil {
		return nil, err
	}

	days := uptime / 86400
	hours := (uptime % 86400) / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60
	return &Time{
		Days:    days,
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
	}, err

}

func convertToInt(str string) (int, error) {
	integer, err := strconv.Atoi(strings.Split(str, ".")[0])
	if err != nil {
		return 0, err
	}
	return integer, nil
}
