package gpucheck

import (
	"fmt"
	"os/exec"
	"strings"
)

type Info struct {
	Index string `json:"index,omitempty"`
	UUID  string `json:"uuid,omitempty"`
	Name  string `json:"name,omitempty"`
	Power string `json:"power,omitempty"`
	Fan   string `json:"fan,omitempty"`
}

func GetGPUInfo() ([]Info, error) {
	var err error
	name := "nvidia-smi"
	args := []string{"--query-gpu=index,uuid,gpu_name,power.draw,fan.speed",
		"--format=csv,noheader"}

	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	gpuInfos, err := parseCmdOutput(string(output))
	if err != nil {
		return nil, err
	}

	return gpuInfos, nil
}

func parseCmdOutput(str string) ([]Info, error) {
	str = strings.Trim(str, "\n")
	lines := strings.Split(str, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no graphics card")
	}

	gpuInfos := make([]Info, len(lines))
	for i, line := range lines {
		ans := strings.SplitAfterN(line, ",", 5)
		if len(ans) != 5 {
			return nil, fmt.Errorf("invalid line format")
		}

		gpuInfos[i] = Info{
			Index: strings.TrimSpace(ans[0]),
			UUID:  strings.TrimSpace(ans[1]),
			Name:  strings.TrimSpace(ans[2]),
			Power: strings.TrimSpace(ans[3]),
			Fan:   strings.TrimSpace(ans[4]),
		}
	}

	return gpuInfos, nil
}
