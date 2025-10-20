package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/virtualization"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("Virtualization Framework Example")
	fmt.Println("================================")

	// Example 0: Test framework availability by creating a VM instance
	fmt.Println("\n0. Testing Framework Availability:")
	vm := virtualization.NewVZVirtualMachine()
	fmt.Printf("   Created VZVirtualMachine: %v\n", vm.ID != 0)
	fmt.Printf("   VM Object ID: 0x%x (non-zero indicates valid object)\n", vm.ID)
	fmt.Println("   Note: Calling methods like CanStart() requires proper VM configuration")

	// Example 1: Framework Overview
	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   Virtualization framework enables:")
	fmt.Println("   - Creating and running VMs on macOS")
	fmt.Println("   - Running macOS and Linux guest OSes")
	fmt.Println("   - Hardware acceleration via Apple Silicon")
	fmt.Println("   - Network, storage, and graphics emulation")

	// Example 2: Key Classes
	fmt.Println("\n2. Key Classes:")

	classes := []string{
		"VZVirtualMachine - Main VM instance",
		"VZVirtualMachineConfiguration - VM settings",
		"VZBootLoader - Boot configuration",
		"VZMacOSBootLoader - macOS specific boot",
		"VZLinuxBootLoader - Linux kernel boot",
		"VZDiskImageStorageDeviceAttachment - Virtual disks",
		"VZVirtioBlockDeviceConfiguration - Block devices",
		"VZVirtioNetworkDeviceConfiguration - Network devices",
		"VZVirtioGraphicsDeviceConfiguration - Graphics",
		"VZMacPlatformConfiguration - Mac platform settings",
		"VZGenericPlatformConfiguration - Generic platform",
	}

	for i, class := range classes {
		fmt.Printf("   %d. %s\n", i+1, class)
	}

	// Example 3: VM Creation Workflow
	fmt.Println("\n3. VM Creation Workflow:")

	workflow := []string{
		"1. Check VZVirtualMachine.IsSupported() - verify support",
		"2. Create VZVirtualMachineConfiguration",
		"3. Set platform configuration (Mac or Generic)",
		"4. Configure boot loader (macOS or Linux)",
		"5. Add storage devices (disk images)",
		"6. Configure network devices",
		"7. Set up graphics and display",
		"8. Add serial ports for console I/O",
		"9. Configure memory and CPU count",
		"10. Validate configuration",
		"11. Create VZVirtualMachine with config",
		"12. Start VM with completion handler",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 4: Platform Configurations
	fmt.Println("\n4. Platform Configurations:")

	platforms := map[string]string{
		"VZMacPlatformConfiguration": "For macOS guests (Apple Silicon only)",
		"VZGenericPlatformConfiguration": "For Linux and other guests",
	}

	for platform, description := range platforms {
		fmt.Printf("   %-35s: %s\n", platform, description)
	}

	// Example 5: Boot Loaders
	fmt.Println("\n5. Boot Loaders:")

	bootLoaders := []string{
		"VZMacOSBootLoader - Boot macOS from IPSW restore images",
		"VZLinuxBootLoader - Boot Linux with kernel, initrd, cmdline",
		"VZEFIBootLoader - Boot using EFI firmware",
	}

	for i, loader := range bootLoaders {
		fmt.Printf("   %d. %s\n", i+1, loader)
	}

	// Example 6: Storage Configuration
	fmt.Println("\n6. Storage Configuration:")

	storage := []string{
		"VZDiskImageStorageDeviceAttachment - Attach disk images",
		"VZVirtioBlockDeviceConfiguration - Virtio block devices",
		"VZVirtioTraditionalMemoryBalloonDeviceConfiguration - Memory balloon",
		"Support for raw, qcow2, and other formats",
		"Read-only and read-write modes",
	}

	for i, item := range storage {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 7: Network Configuration
	fmt.Println("\n7. Network Configuration:")

	network := []string{
		"VZVirtioNetworkDeviceConfiguration - Virtio network adapter",
		"VZNATNetworkDeviceAttachment - NAT networking",
		"VZBridgedNetworkDeviceAttachment - Bridged networking",
		"VZFileHandleNetworkDeviceAttachment - Custom networking",
		"MAC address configuration",
	}

	for i, item := range network {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 8: Graphics and Display
	fmt.Println("\n8. Graphics and Display:")

	graphics := []string{
		"VZVirtioGraphicsDeviceConfiguration - Virtio GPU",
		"VZVirtioGraphicsScanoutConfiguration - Display config",
		"VZMacGraphicsDeviceConfiguration - Mac graphics (Apple Silicon)",
		"VZMacGraphicsDisplayConfiguration - Mac displays",
		"Hardware acceleration support",
	}

	for i, item := range graphics {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 9: Console and Serial Ports
	fmt.Println("\n9. Console and Serial Ports:")

	console := []string{
		"VZVirtioConsoleDeviceConfiguration - Virtio console",
		"VZVirtioConsolePortConfiguration - Console ports",
		"VZFileHandleSerialPortAttachment - File-based I/O",
		"VZFileSerialPortAttachment - File attachment",
		"Useful for debugging and system console access",
	}

	for i, item := range console {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 10: Audio Configuration
	fmt.Println("\n10. Audio Configuration:")

	audio := []string{
		"VZVirtioSoundDeviceConfiguration - Virtio sound device",
		"VZVirtioSoundDeviceInputStreamConfiguration - Audio input",
		"VZVirtioSoundDeviceOutputStreamConfiguration - Audio output",
		"VZHostAudioInputStreamSource - Host microphone",
		"VZHostAudioOutputStreamSink - Host speakers",
	}

	for i, item := range audio {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 11: VM States
	fmt.Println("\n11. VM States:")

	states := []string{
		"stopped - VM is not running",
		"running - VM is actively running",
		"paused - VM is paused",
		"error - VM encountered an error",
		"starting - VM is starting up",
		"pausing - VM is pausing",
		"resuming - VM is resuming",
		"stopping - VM is shutting down",
	}

	for i, state := range states {
		fmt.Printf("   %d. %s\n", i+1, state)
	}

	// Example 12: Memory and CPU
	fmt.Println("\n12. Memory and CPU Configuration:")

	resources := []string{
		"memorySize - RAM allocation (in bytes)",
		"cpuCount - Number of virtual CPUs",
		"Must be within host machine capabilities",
		"Consider guest OS requirements",
		"Balance performance vs resource usage",
	}

	for i, item := range resources {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 13: File Sharing
	fmt.Println("\n13. File Sharing (Directory Sharing):")

	sharing := []string{
		"VZDirectorySharingDeviceConfiguration - Share folders",
		"VZSharedDirectory - Directory to share",
		"VZVirtioFileSystemDeviceConfiguration - Virtio FS",
		"Read-only and read-write modes",
		"Mount points in guest OS",
	}

	for i, item := range sharing {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 14: Requirements and Limitations
	fmt.Println("\n14. Requirements and Limitations:")

	requirements := []string{
		"macOS 11.0+ (Big Sur or later)",
		"Apple Silicon for macOS guests",
		"Intel or Apple Silicon for Linux guests",
		"Requires Virtualization entitlement",
		"Sufficient RAM and disk space",
		"Guest OS must support virtualized hardware",
	}

	for i, req := range requirements {
		fmt.Printf("   %d. %s\n", i+1, req)
	}

	// Example 15: Use Cases
	fmt.Println("\n15. Use Cases:")

	useCases := map[string]string{
		"Development":     "Test apps on different OS versions",
		"CI/CD":           "Automated testing in VMs",
		"Sandboxing":      "Run untrusted code safely",
		"Server Apps":     "Host server applications",
		"Cross-Platform":  "Develop for Linux on macOS",
		"Education":       "Teaching OS internals",
		"Security":        "Malware analysis in isolation",
	}

	for useCase, description := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, description)
	}

	// Example 16: macOS Guest Specifics
	fmt.Println("\n16. macOS Guest Requirements (Apple Silicon):")

	macOSReqs := []string{
		"Download macOS restore image (IPSW)",
		"Use VZMacOSRestoreImage to load IPSW",
		"Create auxiliary storage for macOS",
		"Use VZMacPlatformConfiguration",
		"Configure machine identifier",
		"Set up hardware model",
	}

	for i, req := range macOSReqs {
		fmt.Printf("   %d. %s\n", i+1, req)
	}

	// Example 17: Linux Guest Specifics
	fmt.Println("\n17. Linux Guest Requirements:")

	linuxReqs := []string{
		"Linux kernel image (vmlinuz)",
		"Initial ramdisk (initrd)",
		"Kernel command line parameters",
		"Root disk image with Linux installed",
		"Use VZLinuxBootLoader",
		"Configure VZGenericPlatformConfiguration",
	}

	for i, req := range linuxReqs {
		fmt.Printf("   %d. %s\n", i+1, req)
	}

	// Example 18: Error Handling
	fmt.Println("\n18. Error Handling:")

	errors := []string{
		"Check configuration validation errors",
		"Handle VM startup failures",
		"Monitor VM state changes",
		"Implement VZVirtualMachineDelegate",
		"Handle guestDidStop notifications",
		"Log errors for debugging",
	}

	for i, item := range errors {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 19: Performance Tips
	fmt.Println("\n19. Performance Tips:")

	tips := []string{
		"Use Virtio devices for best performance",
		"Allocate sufficient RAM for guest OS",
		"Use SSD-backed disk images",
		"Enable hardware acceleration",
		"Configure appropriate CPU count",
		"Use NAT networking for simplicity",
		"Consider memory ballooning for efficiency",
	}

	for i, tip := range tips {
		fmt.Printf("   %d. %s\n", i+1, tip)
	}

	// Example 20: Code Structure Example
	fmt.Println("\n20. Example Code Structure:")
	fmt.Println("   // Check support")
	fmt.Println("   if !virtualization.VZVirtualMachine.IsSupported() {")
	fmt.Println("       panic(\"Virtualization not supported\")")
	fmt.Println("   }")
	fmt.Println("   ")
	fmt.Println("   // Create configuration")
	fmt.Println("   config := virtualization.NewVZVirtualMachineConfiguration()")
	fmt.Println("   ")
	fmt.Println("   // Set platform (example)")
	fmt.Println("   // platform := virtualization.NewVZGenericPlatformConfiguration()")
	fmt.Println("   // config.SetPlatform(platform)")
	fmt.Println("   ")
	fmt.Println("   // Configure boot loader, storage, network, etc.")
	fmt.Println("   // ...")
	fmt.Println("   ")
	fmt.Println("   // Validate configuration")
	fmt.Println("   // valid, err := config.ValidateWithError()")
	fmt.Println("   ")
	fmt.Println("   // Create and start VM")
	fmt.Println("   // vm := virtualization.NewVZVirtualMachineWithConfiguration(config)")
	fmt.Println("   // vm.StartWithCompletionHandler(completionHandler)")

	fmt.Println("\n✓ Virtualization framework overview completed!")
	fmt.Println("\nNote: Full VM creation requires:")
	fmt.Println("  - Disk images (raw, qcow2, etc.)")
	fmt.Println("  - Boot loader configuration")
	fmt.Println("  - Guest OS installation media")
	fmt.Println("  - Proper entitlements in app")
	fmt.Println("  - macOS 11.0+ on Apple Silicon for macOS guests")
	fmt.Println("\nFor complete examples, see:")
	fmt.Println("  - https://developer.apple.com/documentation/virtualization")
	fmt.Println("  - Sample code: Running macOS in a Virtual Machine")
	fmt.Println("  - Sample code: Running Linux in a Virtual Machine")
}
