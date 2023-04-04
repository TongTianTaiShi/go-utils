package gpucheck

import "testing"

func TestGetGpuInfo(t *testing.T) {
	infos, err := GetGPUInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(infos)
}
