package fetcher_test

import (
	"os"
	"testing"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func TestGetMemory(t *testing.T) {
	meminfo_data := `
MemTotal:       16384256 kB
MemFree:         2345678 kB
MemAvailable:   11234567 kB
Buffers:          345678 kB
Cached:          4567890 kB
SwapCached:            0 kB
Active:          6789012 kB
Inactive:        3456789 kB
SwapTotal:       2097148 kB
SwapFree:        2097148 kB
	`
	filename := "test-meminfo"
	file, err := os.CreateTemp("", filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte(meminfo_data)); err != nil {
		t.Fatal(err)
	}

	file.Close()

	expectedTotalMem := 15.63
	expectedFreeMem := 10.71
	expectedUsedMem := 4.91

	got, err := fetcher.GetMemory(file.Name())

	if err != nil {
		t.Fatal(err)
	}

	if expectedTotalMem != got.TotalMem {
		t.Errorf("Expected total memory to be %f but got %f", expectedTotalMem, got.TotalMem)
	}

	if expectedFreeMem != got.FreeMem {
		t.Errorf("Expected total memory to be %f but got %f", expectedFreeMem, got.FreeMem)
	}

	if expectedUsedMem != got.UsedMem {
		t.Errorf("Expected total memory to be %f but got %f", expectedUsedMem, got.UsedMem)
	}

}
