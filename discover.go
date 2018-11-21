package logtic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func discoverDescriptorForProcess() (*instance, error) {
	if _, err := os.Stat(instancePath()); err != nil {
		return nil, nil
	}

	data, err := ioutil.ReadFile(instancePath())
	if err != nil {
		return nil, err
	}

	var settings instance
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}

type instance struct {
	FilePointer uint64 `json:"file_pointer_address"`
	Path        string `json:"path"`
	Level       int    `json:"level"`
}

func (i *instance) save() {
	data, err := json.Marshal(i)
	if err != nil {
		return
	}

	ioutil.WriteFile(instancePath(), data, 0644)
}

func deleteInstance() {
	os.Remove(instancePath())
}

func instancePath() string {
	pid := os.Getpid()
	tmpDir := os.TempDir()
	return path.Join(tmpDir, fmt.Sprintf("logtic_process_%d.fd", pid))
}
