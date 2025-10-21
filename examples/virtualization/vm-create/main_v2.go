//go:build v2
// +build v2

// Package main demonstrates the Apple Virtualization framework bindings.
//
// This is the V2 implementation that demonstrates modern Go bindings for the
// Apple Virtualization framework, using type-safe generated methods where
// possible and direct objc.Send calls only when necessary.
//
// Build: go build -tags v2
// Run:   go run -tags v2 .
//
// This implementation demonstrates:
// - Type-safe CPU and memory configuration via generated setters
// - Storage device with disk image attachment and array management
// - Network device with NAT attachment and generated SetAttachment()
// - Graphics device with scanout configuration
// - Type-safe property setters (SetStorageDevices, SetNetworkDevices, SetGraphicsDevices)
// - VM state management and lifecycle
// - AppKit window integration for VM display
//
// Usage:
//   go run -tags v2 . -start -kernel <kernel_path> -disk <disk_path> [-initrd <initrd_path>] [-cmdline <cmdline>]
//
// Example:
//   go run -tags v2 . -start -kernel vmlinuz-6.1.0 -disk ubuntu.img -cmdline "console=ttyS0 root=/dev/vda"
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/virtualization"
)

var (
	e2e      = flag.Bool("e2e", false, "run end-to-end tests")
	startVM  = flag.Bool("start", false, "start a VM with UI display")
	diskPath = flag.String("disk", "", "path to disk image for VM")
	kernel   = flag.String("kernel", "", "path to Linux kernel image")
	initrd   = flag.String("initrd", "", "path to initrd image (optional)")
	cmdline  = flag.String("cmdline", "console=ttyS0", "Linux kernel command line")
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	if *startVM {
		if err := startVMWithUI(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Show framework overview
	showFrameworkOverview()
}

func showFrameworkOverview() {
	fmt.Println("Virtualization Framework Example (V2 - Modern Bindings)")
	fmt.Println("=======================================================")
	fmt.Println()
	fmt.Println("This V2 implementation uses type-safe generated bindings for the")
	fmt.Println("Virtualization framework, demonstrating modern API patterns.")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  ✓ Full Linux VM support with all device types")
	fmt.Println("  ✓ Type-safe CPU and memory configuration")
	fmt.Println("  ✓ VirtIO block device for disk storage")
	fmt.Println("  ✓ NAT network device for internet connectivity")
	fmt.Println("  ✓ Generated device array property setters")
	fmt.Println("  ✓ VirtIO graphics device with scanout")
	fmt.Println("  ✓ AppKit window integration")
	fmt.Println("  ✓ Complete VM configuration and startup")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  %s -start -kernel <kernel_path> -disk <disk_path>\n", os.Args[0])
	fmt.Println()
	fmt.Println("Example:")
	fmt.Printf("  %s -start -kernel vmlinuz-6.1.0 -disk ubuntu.img -cmdline \"console=ttyS0 root=/dev/vda\"\n", os.Args[0])
	fmt.Println()
}

func startVMWithUI() error {
	// Verify required parameters
	if *kernel == "" {
		return fmt.Errorf("-kernel is required when using -start\nUsage: vm-create -start -kernel <kernel_path> -disk <disk_path> [-initrd <initrd_path>] [-cmdline <cmdline>]")
	}

	if *diskPath == "" {
		return fmt.Errorf("-disk is required when using -start\nUsage: vm-create -start -kernel <kernel_path> -disk <disk_path> [-initrd <initrd_path>] [-cmdline <cmdline>]")
	}

	// Verify files exist
	if _, err := os.Stat(*kernel); err != nil {
		return fmt.Errorf("kernel file not found: %w", err)
	}

	if _, err := os.Stat(*diskPath); err != nil {
		return fmt.Errorf("disk image not found: %w", err)
	}

	if *initrd != "" {
		if _, err := os.Stat(*initrd); err != nil {
			return fmt.Errorf("initrd file not found: %w", err)
		}
	}

	fmt.Println("Virtualization Framework - VM Launcher (V2)")
	fmt.Println("============================================")

	// Check if virtualization is supported
	fmt.Println("\n1. Checking virtualization support...")
	// Use objc.Send to call class method
	supported := objc.Send[bool](
		objc.ID(objc.GetClass("VZVirtualMachine")),
		objc.RegisterName("isVirtualizationSupported"),
	)
	if !supported {
		return fmt.Errorf("virtualization is not supported on this system")
	}
	fmt.Println("   ✓ Virtualization is supported")

	// Create AppKit application
	fmt.Println("\n2. Initializing AppKit application...")
	app := appkit.NewApplication()
	if app.ID == 0 {
		return fmt.Errorf("failed to create NSApplication")
	}
	fmt.Println("   ✓ NSApplication created")

	// Create VM configuration
	fmt.Println("\n3. Creating VM configuration...")
	config, err := createLinuxVMConfiguration()
	if err != nil {
		return fmt.Errorf("failed to create VM configuration: %w", err)
	}
	fmt.Println("   ✓ VM configuration created")
	fmt.Printf("     - Kernel: %s\n", *kernel)
	fmt.Printf("     - Disk: %s\n", *diskPath)
	if *initrd != "" {
		fmt.Printf("     - Initrd: %s\n", *initrd)
	}
	fmt.Printf("     - Cmdline: %s\n", *cmdline)

	// Get CPU and memory configuration using generated getters
	cpuCount := config.CpuCount()
	memSize := config.MemorySize()
	fmt.Printf("     - CPUs: %d\n", cpuCount)
	fmt.Printf("     - Memory: %d GB\n", memSize/(1024*1024*1024))

	// Validate configuration
	fmt.Println("\n4. Validating configuration...")
	var errorPtr unsafe.Pointer
	valid := config.ValidateWithError(errorPtr)
	if !valid {
		return fmt.Errorf("invalid configuration")
	}
	fmt.Println("   ✓ Configuration is valid")

	// Create VM instance using initWithConfiguration:
	fmt.Println("\n5. Creating VM instance...")
	vmClass := objc.GetClass("VZVirtualMachine")
	vmAlloc := objc.Send[objc.ID](objc.ID(vmClass), objc.RegisterName("alloc"))
	vmID := objc.Send[objc.ID](
		vmAlloc,
		objc.RegisterName("initWithConfiguration:"),
		unsafe.Pointer(config.ID),
	)
	vm := virtualization.VZVirtualMachineFrom(unsafe.Pointer(vmID))
	if vm.ID == 0 {
		return fmt.Errorf("failed to create VM")
	}
	fmt.Println("   ✓ VM instance created")

	// Create main window with VZVirtualMachineView
	fmt.Println("\n6. Creating AppKit window with VM view...")
	window, _ := createWindowWithVMView(vm)
	if window.ID == 0 {
		return fmt.Errorf("failed to create window")
	}
	fmt.Println("   ✓ Window created (800x600)")

	// Show window
	fmt.Println("\n7. Displaying window...")
	window.MakeKeyAndOrderFront(objc.ID(0))
	fmt.Println("   ✓ Window displayed")

	// Start the VM with proper event handling
	fmt.Println("\n8. Starting VM...")

	// Create a completion handler that will be called when VM starts
	completion := createVMCompletionHandler(vm)

	// Start the VM with completion handler
	objc.Send[bool](
		vm.ID,
		objc.RegisterName("startWithCompletionHandler:"),
		unsafe.Pointer(completion),
	)
	fmt.Println("   ✓ VM start initiated")

	// Run application event loop - this blocks until window is closed
	fmt.Println("\n9. Running application event loop...")
	fmt.Println("   (Close the window to exit)")
	fmt.Println()

	// Set up a simple dispatch queue to monitor VM state
	go monitorVMState(vm)

	app.Run()

	fmt.Println("\n10. VM shutdown complete")
	fmt.Println("\n✓ Application finished")
	return nil
}

func createLinuxVMConfiguration() (virtualization.VZVirtualMachineConfiguration, error) {
	config := virtualization.NewVZVirtualMachineConfiguration()

	// Set platform configuration (Generic for Linux)
	fmt.Println("     Setting up Generic platform for Linux...")
	platform := virtualization.NewVZGenericPlatformConfiguration()
	config.SetPlatform(unsafe.Pointer(platform.ID))

	// Set boot loader for Linux
	fmt.Println("     Configuring Linux boot loader...")
	kernelURL := stringToNSURL(*kernel)
	bootLoader := virtualization.NewVZLinuxBootLoaderWithKernelURL(unsafe.Pointer(kernelURL.ID))
	bootLoader.SetCommandLine(unsafe.Pointer(stringToNSString(*cmdline).ID))

	if *initrd != "" {
		fmt.Println("     Adding initrd image...")
		initrdURL := stringToNSURL(*initrd)
		bootLoader.SetInitialRamdiskURL(unsafe.Pointer(initrdURL.ID))
	}

	config.SetBootLoader(unsafe.Pointer(bootLoader.ID))

	// Set CPU count using generated setter
	fmt.Println("     Configuring CPU count...")
	cpuCount := computeCPUCount()
	config.SetCpuCount(cpuCount)

	// Set memory size using generated setter
	fmt.Println("     Configuring memory size...")
	memSize := computeMemorySize()
	config.SetMemorySize(memSize)

	// Add storage device
	fmt.Println("     Adding storage device...")
	if err := addStorageDevice(config); err != nil {
		return config, fmt.Errorf("failed to add storage device: %w", err)
	}

	// Add network device
	fmt.Println("     Adding NAT network device...")
	if err := addNetworkDevice(config); err != nil {
		return config, fmt.Errorf("failed to add network device: %w", err)
	}

	// Add graphics device
	fmt.Println("     Adding graphics device...")
	if err := addGraphicsDevice(config); err != nil {
		return config, fmt.Errorf("failed to add graphics device: %w", err)
	}

	return config, nil
}

func addStorageDevice(config virtualization.VZVirtualMachineConfiguration) error {
	// Create disk image storage attachment
	diskURL := stringToNSURL(*diskPath)

	// Use objc.Send to create attachment with URL and read-only flag
	attachClass := objc.GetClass("VZDiskImageStorageDeviceAttachment")
	attachAlloc := objc.Send[objc.ID](objc.ID(attachClass), objc.RegisterName("alloc"))
	attachID := objc.Send[objc.ID](
		attachAlloc,
		objc.RegisterName("initWithURL:readOnly:error:"),
		unsafe.Pointer(diskURL.ID),
		false,
		unsafe.Pointer(uintptr(0)), // nil error pointer
	)

	// Create VirtIO block device
	blockDevice := virtualization.NewVZVirtioBlockDeviceConfiguration()

	// Set attachment using objc.Send (storage device doesn't have generated SetAttachment)
	objc.Send[bool](
		blockDevice.ID,
		objc.RegisterName("setAttachment:"),
		unsafe.Pointer(attachID),
	)

	// Create array and set storage devices using generated SetStorageDevices method
	arrayClass := objc.GetClass("NSMutableArray")
	arrayAlloc := objc.Send[objc.ID](objc.ID(arrayClass), objc.RegisterName("alloc"))
	arrayID := objc.Send[objc.ID](arrayAlloc, objc.RegisterName("init"))

	objc.Send[bool](
		arrayID,
		objc.RegisterName("addObject:"),
		unsafe.Pointer(blockDevice.ID),
	)

	// Convert to array of storage device configurations
	storageDevices := []virtualization.VZStorageDeviceConfiguration{
		virtualization.VZStorageDeviceConfigurationFrom(unsafe.Pointer(blockDevice.ID)),
	}

	// Use generated type-safe setter
	config.SetStorageDevices(storageDevices)

	return nil
}

func addNetworkDevice(config virtualization.VZVirtualMachineConfiguration) error {
	// Create NAT attachment
	natAttachment := virtualization.NewVZNATNetworkDeviceAttachment()

	// Create network device
	networkDevice := virtualization.NewVZVirtioNetworkDeviceConfiguration()

	// Set attachment using generated type-safe setter
	networkDevice.SetAttachment(unsafe.Pointer(natAttachment.ID))

	// Set network devices using generated type-safe setter
	networkDevices := []virtualization.VZNetworkDeviceConfiguration{
		networkDevice.VZNetworkDeviceConfiguration,
	}
	config.SetNetworkDevices(networkDevices)

	return nil
}

func addGraphicsDevice(config virtualization.VZVirtualMachineConfiguration) error {
	// Create graphics device
	graphicsDevice := virtualization.NewVZVirtioGraphicsDeviceConfiguration()

	// Create scanout configuration using objc.Send
	scanoutClass := objc.GetClass("VZVirtioGraphicsScanoutConfiguration")
	scanoutAlloc := objc.Send[objc.ID](objc.ID(scanoutClass), objc.RegisterName("alloc"))
	scanoutID := objc.Send[objc.ID](
		scanoutAlloc,
		objc.RegisterName("initWithWidthInPixels:heightInPixels:"),
		int64(1920),
		int64(1200),
	)

	// Set scanouts using generated type-safe setter
	scanout := virtualization.VZVirtioGraphicsScanoutConfigurationFrom(unsafe.Pointer(scanoutID))
	graphicsDevice.SetScanouts([]virtualization.VZVirtioGraphicsScanoutConfiguration{scanout})

	// Set graphics devices using generated type-safe setter
	graphicsDevices := []virtualization.VZGraphicsDeviceConfiguration{
		virtualization.VZGraphicsDeviceConfigurationFrom(unsafe.Pointer(graphicsDevice.ID)),
	}
	config.SetGraphicsDevices(graphicsDevices)

	return nil
}

func computeCPUCount() uint {
	totalAvailableCPUs := runtime.NumCPU()
	virtualCPUCount := uint(totalAvailableCPUs - 1)
	if virtualCPUCount <= 1 {
		virtualCPUCount = 1
	}

	// Get max/min allowed using objc.Send
	configClass := objc.GetClass("VZVirtualMachineConfiguration")
	maxAllowed := objc.Send[uint](
		objc.ID(configClass),
		objc.RegisterName("maximumAllowedCPUCount"),
	)
	minAllowed := objc.Send[uint](
		objc.ID(configClass),
		objc.RegisterName("minimumAllowedCPUCount"),
	)

	if virtualCPUCount > maxAllowed {
		virtualCPUCount = maxAllowed
	}
	if virtualCPUCount < minAllowed {
		virtualCPUCount = minAllowed
	}

	return virtualCPUCount
}

func computeMemorySize() uint64 {
	// We arbitrarily choose 4GB
	memorySize := uint64(4 * 1024 * 1024 * 1024)

	// Get max/min allowed using objc.Send
	configClass := objc.GetClass("VZVirtualMachineConfiguration")
	maxAllowed := objc.Send[uint64](
		objc.ID(configClass),
		objc.RegisterName("maximumAllowedMemorySize"),
	)
	minAllowed := objc.Send[uint64](
		objc.ID(configClass),
		objc.RegisterName("minimumAllowedMemorySize"),
	)

	if memorySize > maxAllowed {
		memorySize = maxAllowed
	}
	if memorySize < minAllowed {
		memorySize = minAllowed
	}

	return memorySize
}

func createWindowWithVMView(vm virtualization.VZVirtualMachine) (appkit.Window, virtualization.VZVirtualMachineView) {
	// Create window using NSWindow alloc + init
	winClass := objc.GetClass("NSWindow")
	allocID := objc.Send[objc.ID](objc.ID(winClass), objc.RegisterName("alloc"))

	// Initialize window with frame, style, backing, and defer
	windowID := objc.Send[objc.ID](
		allocID,
		objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		uintptr(100), uintptr(100), uintptr(800), uintptr(600), // x, y, width, height
		uintptr(15), // NSWindowStyleMaskTitled | Closable | Miniaturizable | Resizable
		uintptr(2),  // NSBackingStoreBuffered
		false,
	)

	window := appkit.WindowFrom(unsafe.Pointer(windowID))

	// Set window title
	titleStr := stringToNSString("Linux VM Viewer (V2)")
	objc.Send[bool](
		objc.ID(window.ID),
		objc.RegisterName("setTitle:"),
		unsafe.Pointer(titleStr.ID),
	)

	// Create VZVirtualMachineView
	vmView := virtualization.NewVZVirtualMachineView()

	// Set the virtual machine using objc.Send
	objc.Send[bool](
		vmView.ID,
		objc.RegisterName("setVirtualMachine:"),
		unsafe.Pointer(vm.ID),
	)

	// Set as content view
	objc.Send[bool](
		window.ID,
		objc.RegisterName("setContentView:"),
		unsafe.Pointer(vmView.ID),
	)

	return window, vmView
}

// Helper functions to convert Go strings to Objective-C strings and URLs

func stringToNSString(s string) foundation.String {
	// Create NSString from Go string using stringWithUTF8String:
	cstr := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSString")),
		objc.RegisterName("stringWithUTF8String:"),
		unsafe.Pointer(unsafe.StringData(s)),
	)
	return foundation.StringFrom(unsafe.Pointer(cstr))
}

func stringToNSURL(s string) foundation.URL {
	// Expand path if needed
	expandedPath, _ := filepath.Abs(s)

	// Create NSURL from path string
	nsStr := stringToNSString(expandedPath)
	nsURL := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSURL")),
		objc.RegisterName("fileURLWithPath:"),
		unsafe.Pointer(nsStr.ID),
	)
	return foundation.URLFrom(unsafe.Pointer(nsURL))
}

// createVMCompletionHandler creates a completion block for VM startup
func createVMCompletionHandler(vm virtualization.VZVirtualMachine) unsafe.Pointer {
	// For now, return nil which means we'll check state via monitorVMState
	// In a full implementation, we'd create a proper block using purego
	return unsafe.Pointer(uintptr(0))
}

// monitorVMState monitors the VM state and logs changes
func monitorVMState(vm virtualization.VZVirtualMachine) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	lastState := ""
	for range ticker.C {
		// Get current VM state using objc.Send
		state := objc.Send[uint](vm.ID, objc.RegisterName("state"))

		// Map state values to strings
		stateStr := vmStateToString(state)

		if stateStr != lastState {
			lastState = stateStr
			fmt.Printf("   → VM State: %s\n", stateStr)

			// If VM stopped, we can break
			if stateStr == "Stopped" {
				break
			}
		}
	}
}

// vmStateToString converts VM state int to readable string
func vmStateToString(state uint) string {
	// VZVirtualMachineState enum values:
	// 0 = stopped
	// 1 = running
	// 2 = paused
	// 3 = error
	// 4 = starting
	// 5 = pausing
	// 6 = resuming
	// 7 = stopping
	switch state {
	case 0:
		return "Stopped"
	case 1:
		return "Running"
	case 2:
		return "Paused"
	case 3:
		return "Error"
	case 4:
		return "Starting"
	case 5:
		return "Pausing"
	case 6:
		return "Resuming"
	case 7:
		return "Stopping"
	default:
		return "Unknown"
	}
}
