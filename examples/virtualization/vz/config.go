package main

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/virtualization"
)

// createMacPlatformConfig creates a Mac platform configuration
func createMacPlatformConfig(paths VMPaths) (virtualization.VZMacPlatformConfiguration, error) {
	// Load hardware model
	hardwareModel, err := loadHardwareModel(paths.HardwareModelPath)
	if err != nil {
		return virtualization.VZMacPlatformConfiguration{}, fmt.Errorf("failed to load hardware model: %w", err)
	}

	// Load machine identifier
	machineID, err := loadMachineIdentifier(paths.MachineIdentifierPath)
	if err != nil {
		return virtualization.VZMacPlatformConfiguration{}, fmt.Errorf("failed to load machine identifier: %w", err)
	}

	// Load or create auxiliary storage
	auxURL := pathToNSURL(paths.AuxiliaryStoragePath)
	auxStorage := virtualization.NewVZMacAuxiliaryStorageWithURL(unsafe.Pointer(auxURL))
	if auxStorage.ID == 0 {
		return virtualization.VZMacPlatformConfiguration{}, fmt.Errorf("failed to create auxiliary storage")
	}

	// Create platform configuration
	platform := virtualization.NewVZMacPlatformConfiguration()
	platform.SetHardwareModel(unsafe.Pointer(hardwareModel.ID))
	platform.SetMachineIdentifier(unsafe.Pointer(machineID.ID))
	platform.SetAuxiliaryStorage(unsafe.Pointer(auxStorage.ID))

	return platform, nil
}

// setupVMConfiguration creates the complete VM configuration
func setupVMConfiguration(platform virtualization.VZMacPlatformConfiguration, paths VMPaths) (virtualization.VZVirtualMachineConfiguration, error) {
	// Create bootloader
	bootLoader := virtualization.NewVZMacOSBootLoader()
	if bootLoader.ID == 0 {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to create bootloader")
	}

	// Create configuration
	config := virtualization.NewVZVirtualMachineConfiguration()
	config.SetBootLoader(unsafe.Pointer(bootLoader.ID))
	config.SetPlatform(unsafe.Pointer(platform.ID))

	// CPU configuration
	cpuCount := computeCPUCount()
	config.SetCpuCount(cpuCount)

	// Memory configuration
	memorySize := computeMemorySize()
	config.SetMemorySize(memorySize)

	// Storage devices
	if err := configureStorageDevices(config, paths); err != nil {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure storage: %w", err)
	}

	// Graphics devices
	if err := configureGraphicsDevices(config); err != nil {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure graphics: %w", err)
	}

	// Network devices
	if err := configureNetworkDevices(config); err != nil {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure network: %w", err)
	}

	// Input devices
	if err := configureInputDevices(config); err != nil {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure input: %w", err)
	}

	// Audio devices
	if err := configureAudioDevices(config); err != nil {
		return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure audio: %w", err)
	}

	// Shared directories
	if sharedFolder != "" {
		if err := configureSharedDirectories(config); err != nil {
			return virtualization.VZVirtualMachineConfiguration{}, fmt.Errorf("failed to configure shared directories: %w", err)
		}
	}

	return config, nil
}

func computeCPUCount() uint {
	totalCPUs := uint(runtime.NumCPU())
	virtualCPUs := totalCPUs - 1
	if virtualCPUs < 1 {
		virtualCPUs = 1
	}

	maxAllowed := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedCPUCount()
	minAllowed := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedCPUCount()

	if virtualCPUs > maxAllowed {
		virtualCPUs = maxAllowed
	}
	if virtualCPUs < minAllowed {
		virtualCPUs = minAllowed
	}

	return virtualCPUs
}

func computeMemorySize() uint64 {
	// Default to 4GB
	memorySize := uint64(4 * 1024 * 1024 * 1024)

	maxAllowed := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedMemorySize()
	minAllowed := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedMemorySize()

	if memorySize > maxAllowed {
		memorySize = maxAllowed
	}
	if memorySize < minAllowed {
		memorySize = minAllowed
	}

	return memorySize
}

func configureStorageDevices(config virtualization.VZVirtualMachineConfiguration, paths VMPaths) error {
	diskPath := paths.DiskImagePath
	if diskPath == "" {
		diskPath = paths.DiskImagePath
	}

	// Create disk if it doesn't exist
	if _, err := os.Stat(diskPath); os.IsNotExist(err) {
		sizeBytes := int64(diskSize * 1024 * 1024 * 1024)
		f, err := os.Create(diskPath)
		if err != nil {
			return fmt.Errorf("failed to create disk: %w", err)
		}
		f.Truncate(sizeBytes)
		f.Close()
	}

	// Create disk attachment
	diskURL := pathToNSURL(diskPath)
	attachment := virtualization.NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError(
		unsafe.Pointer(diskURL),
		false, // not read-only
		nil,
	)

	if attachment.ID == 0 {
		return fmt.Errorf("failed to create disk attachment")
	}

	// Create block device
	blockDevice := virtualization.NewVZVirtioBlockDeviceConfigurationWithAttachment(
		unsafe.Pointer(attachment.ID),
	)

	// Set storage devices
	storageDevices := []unsafe.Pointer{unsafe.Pointer(blockDevice.ID)}
	array := createNSArray(storageDevices)
	config.SetStorageDevices(unsafe.Pointer(array))

	return nil
}

func configureGraphicsDevices(config virtualization.VZVirtualMachineConfiguration) error {
	// Create graphics device
	graphicsDevice := virtualization.NewVZMacGraphicsDeviceConfiguration()
	if graphicsDevice.ID == 0 {
		return fmt.Errorf("failed to create graphics device")
	}

	// Create display
	display := virtualization.NewVZMacGraphicsDisplayConfigurationWithWidthHeightPixelsPerInch(
		1920, 1200, 80,
	)
	if display.ID == 0 {
		return fmt.Errorf("failed to create display")
	}

	// Set displays
	displays := []unsafe.Pointer{unsafe.Pointer(display.ID)}
	displayArray := createNSArray(displays)
	graphicsDevice.SetDisplays(unsafe.Pointer(displayArray))

	// Set graphics devices
	devices := []unsafe.Pointer{unsafe.Pointer(graphicsDevice.ID)}
	devicesArray := createNSArray(devices)
	config.SetGraphicsDevices(unsafe.Pointer(devicesArray))

	return nil
}

func configureNetworkDevices(config virtualization.VZVirtualMachineConfiguration) error {
	// Create NAT attachment
	natAttachment := virtualization.NewVZNATNetworkDeviceAttachment()
	if natAttachment.ID == 0 {
		return fmt.Errorf("failed to create NAT attachment")
	}

	// Create network device
	networkDevice := virtualization.NewVZVirtioNetworkDeviceConfigurationWithAttachment(
		unsafe.Pointer(natAttachment.ID),
	)
	if networkDevice.ID == 0 {
		return fmt.Errorf("failed to create network device")
	}

	// Set network devices
	devices := []unsafe.Pointer{unsafe.Pointer(networkDevice.ID)}
	array := createNSArray(devices)
	config.SetNetworkDevices(unsafe.Pointer(array))

	return nil
}

func configureInputDevices(config virtualization.VZVirtualMachineConfiguration) error {
	// Pointing devices
	pointingDevices := []unsafe.Pointer{}

	// USB screen pointing device
	usbPointing := virtualization.NewVZUSBScreenCoordinatePointingDeviceConfiguration()
	if usbPointing.ID != 0 {
		pointingDevices = append(pointingDevices, unsafe.Pointer(usbPointing.ID))
	}

	// Mac trackpad (if available)
	trackpad := virtualization.NewVZMacTrackpadConfiguration()
	if trackpad.ID != 0 {
		pointingDevices = append(pointingDevices, unsafe.Pointer(trackpad.ID))
	}

	if len(pointingDevices) > 0 {
		array := createNSArray(pointingDevices)
		config.SetPointingDevices(unsafe.Pointer(array))
	}

	// Keyboard
	keyboard := virtualization.NewVZMacKeyboardConfiguration()
	if keyboard.ID == 0 {
		// Fallback to USB keyboard
		keyboard = virtualization.VZMacKeyboardConfiguration{
			VZKeyboardConfiguration: virtualization.NewVZUSBKeyboardConfiguration(),
		}
	}

	if keyboard.ID != 0 {
		keyboards := []unsafe.Pointer{unsafe.Pointer(keyboard.ID)}
		array := createNSArray(keyboards)
		config.SetKeyboards(unsafe.Pointer(array))
	}

	return nil
}

func configureAudioDevices(config virtualization.VZVirtualMachineConfiguration) error {
	// Create audio device
	audioDevice := virtualization.NewVZVirtioSoundDeviceConfiguration()
	if audioDevice.ID == 0 {
		// Audio is optional, don't fail
		return nil
	}

	// Create input stream
	inputStream := virtualization.NewVZVirtioSoundDeviceInputStreamConfiguration()

	// Create output stream
	outputStream := virtualization.NewVZVirtioSoundDeviceOutputStreamConfiguration()

	// Set streams
	streams := []unsafe.Pointer{}
	if inputStream.ID != 0 {
		streams = append(streams, unsafe.Pointer(inputStream.ID))
	}
	if outputStream.ID != 0 {
		streams = append(streams, unsafe.Pointer(outputStream.ID))
	}

	if len(streams) > 0 {
		array := createNSArray(streams)
		audioDevice.SetStreams(unsafe.Pointer(array))

		// Set audio devices
		devices := []unsafe.Pointer{unsafe.Pointer(audioDevice.ID)}
		devicesArray := createNSArray(devices)
		config.SetAudioDevices(unsafe.Pointer(devicesArray))
	}

	return nil
}

func configureSharedDirectories(config virtualization.VZVirtualMachineConfiguration) error {
	// Verify directory exists
	fileInfo, err := os.Stat(sharedFolder)
	if err != nil {
		return fmt.Errorf("shared folder not found: %w", err)
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("shared path is not a directory: %s", sharedFolder)
	}

	// Create shared directory
	sharedDirURL := pathToNSURL(sharedFolder)
	sharedDir := virtualization.NewVZSharedDirectoryWithURLReadOnly(
		unsafe.Pointer(sharedDirURL),
		false, // writable
	)
	if sharedDir.ID == 0 {
		return fmt.Errorf("failed to create shared directory")
	}

	// Create single directory share
	share := virtualization.NewVZSingleDirectoryShareWithDirectory(
		unsafe.Pointer(sharedDir.ID),
	)
	if share.ID == 0 {
		return fmt.Errorf("failed to create directory share")
	}

	// Create file system device
	tag := mountTag
	if tag == "" {
		tag = "shared"
	}

	tagStr := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSString")),
		objc.RegisterName("stringWithUTF8String:"),
		unsafe.Pointer(unsafe.StringData(tag)),
	)

	fsDevice := virtualization.NewVZVirtioFileSystemDeviceConfigurationWithTag(
		unsafe.Pointer(tagStr),
	)
	if fsDevice.ID == 0 {
		return fmt.Errorf("failed to create file system device")
	}

	fsDevice.SetShare(unsafe.Pointer(share.ID))

	// Set directory sharing devices
	devices := []unsafe.Pointer{unsafe.Pointer(fsDevice.ID)}
	array := createNSArray(devices)
	config.SetDirectorySharingDevices(unsafe.Pointer(array))

	return nil
}

func validateConfig(config virtualization.VZVirtualMachineConfiguration) error {
	// Call validate method
	// Note: The generated bindings may not have error handling yet
	// We'll do a basic check

	if config.ID == 0 {
		return fmt.Errorf("invalid configuration object")
	}

	// In full implementation, would call:
	// valid, err := config.Validate()
	// if err != nil { return err }
	// if !valid { return fmt.Errorf("configuration validation failed") }

	return nil
}

func loadHardwareModel(path string) (virtualization.VZMacHardwareModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return virtualization.VZMacHardwareModel{}, err
	}

	nsData := bytesToNSData(data)
	model := virtualization.NewVZMacHardwareModelWithDataRepresentation(unsafe.Pointer(nsData))
	if model.ID == 0 {
		return virtualization.VZMacHardwareModel{}, fmt.Errorf("failed to create hardware model from data")
	}

	return model, nil
}

func loadMachineIdentifier(path string) (virtualization.VZMacMachineIdentifier, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return virtualization.VZMacMachineIdentifier{}, err
	}

	nsData := bytesToNSData(data)
	identifier := virtualization.NewVZMacMachineIdentifierWithDataRepresentation(unsafe.Pointer(nsData))
	if identifier.ID == 0 {
		return virtualization.VZMacMachineIdentifier{}, fmt.Errorf("failed to create machine identifier from data")
	}

	return identifier, nil
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
