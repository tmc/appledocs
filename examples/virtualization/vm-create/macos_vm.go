// Package main - macOS VM creator using generated Virtualization framework bindings
//
// This demonstrates creating a macOS virtual machine using the bindings generated
// from Apple documentation at github.com/tmc/appledocs/generated/virtualization
//
// Based on Code-Hex/vz example/macOS but using our generated bindings instead of cgo.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/virtualization"
)

// VMPaths holds paths for VM bundle contents
type VMPaths struct {
	BundlePath           string
	DiskImagePath        string
	AuxiliaryStoragePath string
	HardwareModelPath    string
	MachineIdentifierPath string
	RestoreImagePath     string
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

// CreateVMConfig creates a macOS VM configuration using generated bindings
func CreateVMConfig() virtualization.VZVirtualMachineConfiguration {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	paths := GetVMPaths()

	// Ensure VM bundle exists
	os.MkdirAll(paths.BundlePath, 0755)

	// Create platform configuration
	platform := createMacPlatformConfig(paths)

	// Create bootloader
	bootLoader := virtualization.NewVZMacOSBootLoader()

	// Create main configuration
	config := virtualization.NewVZVirtualMachineConfiguration()
	config.SetBootLoader(unsafe.Pointer(bootLoader.ID))
	config.SetPlatform(unsafe.Pointer(platform.ID))

	// Set CPU count
	cpuCount := uint(runtime.NumCPU() - 1)
	if cpuCount < 1 {
		cpuCount = 1
	}
	maxCPU := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedCPUCount()
	minCPU := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedCPUCount()
	if cpuCount > maxCPU {
		cpuCount = maxCPU
	}
	if cpuCount < minCPU {
		cpuCount = minCPU
	}
	config.SetCpuCount(cpuCount)

	// Set memory size (4GB)
	memorySize := uint64(4 * 1024 * 1024 * 1024)
	maxMem := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedMemorySize()
	minMem := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedMemorySize()
	if memorySize > maxMem {
		memorySize = maxMem
	}
	if memorySize < minMem {
		memorySize = minMem
	}
	config.SetMemorySize(memorySize)

	// Add block device for disk
	diskAttachment := createDiskAttachment(paths.DiskImagePath, 64*1024*1024*1024)
	blockDevice := virtualization.NewVZVirtioBlockDeviceConfigurationWithAttachment(
		unsafe.Pointer(diskAttachment.ID),
	)

	// Set storage devices
	storageDevices := []unsafe.Pointer{unsafe.Pointer(blockDevice.ID)}
	nsArray := createNSArray(storageDevices)
	config.SetStorageDevices(unsafe.Pointer(nsArray))

	// Add graphics device
	graphicsDevice := virtualization.NewVZMacGraphicsDeviceConfiguration()
	display := virtualization.NewVZMacGraphicsDisplayConfigurationWithWidthHeightPixelsPerInch(
		1920, 1200, 80,
	)

	// Set displays on graphics device
	displays := []unsafe.Pointer{unsafe.Pointer(display.ID)}
	displayArray := createNSArray(displays)
	graphicsDevice.SetDisplays(unsafe.Pointer(displayArray))

	// Set graphics devices
	graphicsDevices := []unsafe.Pointer{unsafe.Pointer(graphicsDevice.ID)}
	graphicsArray := createNSArray(graphicsDevices)
	config.SetGraphicsDevices(unsafe.Pointer(graphicsArray))

	// Add network device
	natAttachment := virtualization.NewVZNATNetworkDeviceAttachment()
	networkDevice := virtualization.NewVZVirtioNetworkDeviceConfigurationWithAttachment(
		unsafe.Pointer(natAttachment.ID),
	)

	networkDevices := []unsafe.Pointer{unsafe.Pointer(networkDevice.ID)}
	networkArray := createNSArray(networkDevices)
	config.SetNetworkDevices(unsafe.Pointer(networkArray))

	// Add pointing device
	pointingDevice := virtualization.NewVZUSBScreenCoordinatePointingDeviceConfiguration()
	pointingDevices := []unsafe.Pointer{unsafe.Pointer(pointingDevice.ID)}
	pointingArray := createNSArray(pointingDevices)
	config.SetPointingDevices(unsafe.Pointer(pointingArray))

	// Add keyboard
	keyboard := virtualization.NewVZMacKeyboardConfiguration()
	keyboards := []unsafe.Pointer{unsafe.Pointer(keyboard.ID)}
	keyboardArray := createNSArray(keyboards)
	config.SetKeyboards(unsafe.Pointer(keyboardArray))

	// Add trackpad
	trackpad := virtualization.NewVZMacTrackpadConfiguration()
	pointingDevices = append(pointingDevices, unsafe.Pointer(trackpad.ID))
	pointingArray = createNSArray(pointingDevices)
	config.SetPointingDevices(unsafe.Pointer(pointingArray))

	fmt.Printf("Created VM configuration:\n")
	fmt.Printf("  CPUs: %d\n", cpuCount)
	fmt.Printf("  Memory: %d GB\n", memorySize/(1024*1024*1024))
	fmt.Printf("  Disk: %s\n", paths.DiskImagePath)

	return config
}

// createMacPlatformConfig creates a Mac platform configuration
func createMacPlatformConfig(paths VMPaths) virtualization.VZMacPlatformConfiguration {
	// Load or create auxiliary storage
	auxStorage := virtualization.NewVZMacAuxiliaryStorageWithURL(
		unsafe.Pointer(pathToNSURL(paths.AuxiliaryStoragePath).ID),
	)

	// Load or create hardware model
	hardwareModel := loadOrCreateHardwareModel(paths.HardwareModelPath)

	// Load or create machine identifier
	machineID := loadOrCreateMachineIdentifier(paths.MachineIdentifierPath)

	// Create platform configuration
	platform := virtualization.NewVZMacPlatformConfiguration()
	platform.SetAuxiliaryStorage(unsafe.Pointer(auxStorage.ID))
	platform.SetHardwareModel(unsafe.Pointer(hardwareModel.ID))
	platform.SetMachineIdentifier(unsafe.Pointer(machineID.ID))

	return platform
}

// loadOrCreateHardwareModel loads existing or creates new hardware model
func loadOrCreateHardwareModel(path string) virtualization.VZMacHardwareModel {
	if _, err := os.Stat(path); err == nil {
		// Load existing
		data, err := os.ReadFile(path)
		if err == nil {
			nsData := bytesToNSData(data)
			model := virtualization.NewVZMacHardwareModelWithDataRepresentation(
				unsafe.Pointer(nsData),
			)
			if model.ID != 0 {
				return model
			}
		}
	}

	// Create new - use supported model
	model := virtualization.VZMacHardwareModelClass.Supported()
	if model.ID != 0 {
		// Save for next time
		data := nsDataToBytes(unsafe.Pointer(model.DataRepresentation()))
		os.WriteFile(path, data, 0644)
	}

	return model
}

// loadOrCreateMachineIdentifier loads existing or creates new machine identifier
func loadOrCreateMachineIdentifier(path string) virtualization.VZMacMachineIdentifier {
	if _, err := os.Stat(path); err == nil {
		// Load existing
		data, err := os.ReadFile(path)
		if err == nil {
			nsData := bytesToNSData(data)
			identifier := virtualization.NewVZMacMachineIdentifierWithDataRepresentation(
				unsafe.Pointer(nsData),
			)
			if identifier.ID != 0 {
				return identifier
			}
		}
	}

	// Create new
	identifier := virtualization.NewVZMacMachineIdentifier()
	if identifier.ID != 0 {
		// Save for next time
		data := nsDataToBytes(unsafe.Pointer(identifier.DataRepresentation()))
		os.WriteFile(path, data, 0644)
	}

	return identifier
}

// createDiskAttachment creates a disk image attachment
func createDiskAttachment(path string, sizeBytes uint64) virtualization.VZDiskImageStorageDeviceAttachment {
	// Create disk if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create empty disk file
		f, err := os.Create(path)
		if err != nil {
			panic(fmt.Sprintf("failed to create disk: %v", err))
		}
		f.Truncate(int64(sizeBytes))
		f.Close()
		fmt.Printf("Created disk image: %s (%d GB)\n", path, sizeBytes/(1024*1024*1024))
	}

	url := pathToNSURL(path)
	attachment := virtualization.NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError(
		unsafe.Pointer(url.ID),
		false, // not read-only
		nil,   // error pointer
	)

	return attachment
}

// Helper functions for Objective-C interop

func pathToNSURL(path string) foundation.URL {
	absPath, _ := filepath.Abs(path)
	nsStr := stringToNSString(absPath)
	nsURL := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSURL")),
		objc.RegisterName("fileURLWithPath:"),
		unsafe.Pointer(nsStr.ID),
	)
	return foundation.URLFrom(unsafe.Pointer(nsURL))
}

func stringToNSString(s string) foundation.String {
	cstr := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSString")),
		objc.RegisterName("stringWithUTF8String:"),
		unsafe.Pointer(unsafe.StringData(s)),
	)
	return foundation.StringFrom(unsafe.Pointer(cstr))
}

func bytesToNSData(data []byte) objc.ID {
	if len(data) == 0 {
		return objc.Send[objc.ID](
			objc.ID(objc.GetClass("NSData")),
			objc.RegisterName("data"),
		)
	}
	return objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSData")),
		objc.RegisterName("dataWithBytes:length:"),
		unsafe.Pointer(&data[0]),
		uintptr(len(data)),
	)
}

func nsDataToBytes(nsData unsafe.Pointer) []byte {
	if nsData == nil {
		return nil
	}

	length := objc.Send[uintptr](objc.ID(nsData), objc.RegisterName("length"))
	if length == 0 {
		return nil
	}

	bytesPtr := objc.Send[unsafe.Pointer](objc.ID(nsData), objc.RegisterName("bytes"))
	return unsafe.Slice((*byte)(bytesPtr), length)
}

func createNSArray(objects []unsafe.Pointer) objc.ID {
	if len(objects) == 0 {
		return objc.Send[objc.ID](
			objc.ID(objc.GetClass("NSArray")),
			objc.RegisterName("array"),
		)
	}

	return objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSArray")),
		objc.RegisterName("arrayWithObjects:count:"),
		unsafe.Pointer(&objects[0]),
		uintptr(len(objects)),
	)
}

// RunMacOSVM demonstrates running a macOS VM
func RunMacOSVM() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	config := CreateVMConfig()

	// Create VM
	vm := virtualization.NewVZVirtualMachineWithConfiguration(unsafe.Pointer(config.ID))
	if vm.ID == 0 {
		return fmt.Errorf("failed to create virtual machine")
	}

	fmt.Println("✓ Virtual machine created successfully")
	fmt.Println("\nNote: To run the VM, you need to:")
	fmt.Println("  1. Start the VM with vm.Start()")
	fmt.Println("  2. Handle state change notifications")
	fmt.Println("  3. Set up graphics output (VZVirtualMachineView or framebuffer)")
	fmt.Println("  4. Handle lifecycle events")
	fmt.Println("\nSee the full vz example for complete implementation")

	return nil
}
