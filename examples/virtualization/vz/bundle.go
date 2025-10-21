package main

import (
	"os"
	"path/filepath"
)

// VMPaths holds paths for VM bundle contents
type VMPaths struct {
	BundlePath            string
	DiskImagePath         string
	AuxiliaryStoragePath  string
	HardwareModelPath     string
	MachineIdentifierPath string
	RestoreImagePath      string
}

// GetVMPaths returns standard paths for macOS VM bundle
func GetVMPaths() VMPaths {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	bundlePath := filepath.Join(home, "VM.bundle")

	return VMPaths{
		BundlePath:            bundlePath,
		DiskImagePath:         filepath.Join(bundlePath, "Disk.img"),
		AuxiliaryStoragePath:  filepath.Join(bundlePath, "AuxiliaryStorage"),
		HardwareModelPath:     filepath.Join(bundlePath, "HardwareModel"),
		MachineIdentifierPath: filepath.Join(bundlePath, "MachineIdentifier"),
		RestoreImagePath:      filepath.Join(bundlePath, "RestoreImage.ipsw"),
	}
}
