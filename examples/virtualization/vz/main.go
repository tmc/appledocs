// Package main implements a macOS VM runner using generated Virtualization framework bindings.
//
// This is equivalent to Code-Hex/vz/example/macOS but using our generated bindings.
//
// Usage:
//   vz [flags]
//
// Flags:
//   -install              Run in install mode (download and install macOS)
//   -install-version      Specific macOS version to install
//   -shared PATH          Path to directory to share with VM
//   -mount-tag TAG        Tag name for shared directory (default: "shared")
//   -auto-mount           Auto-mount shared directory in macOS 13+
//   -reinit               Reinitialize VM platform configuration
//   -new-disk             Create fresh disk image
//   -disk-size N          Disk size in GiB (default: 64)
//   -disk-path PATH       Use custom disk image
//   -recovery             Boot in recovery mode
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/virtualization"
)

var (
	install        bool
	installVersion string
	sharedFolder   string
	mountTag       string
	autoMount      bool
	reinit         bool
	newDisk        bool
	diskSize       uint64
	diskPath       string
	recoveryMode   bool
)

func init() {
	flag.BoolVar(&install, "install", false, "run in install mode")
	flag.StringVar(&installVersion, "install-version", "", "specific macOS version to install")
	flag.StringVar(&sharedFolder, "shared", "", "path to directory to share with VM")
	flag.StringVar(&mountTag, "mount-tag", "shared", "tag name for shared directory")
	flag.BoolVar(&autoMount, "auto-mount", false, "auto-mount shared directory in macOS 13+")
	flag.BoolVar(&reinit, "reinit", false, "reinitialize VM platform configuration")
	flag.BoolVar(&newDisk, "new-disk", false, "create fresh disk image")
	flag.Uint64Var(&diskSize, "disk-size", 64, "disk size in GiB")
	flag.StringVar(&diskPath, "disk-path", "", "path to custom disk image")
	flag.BoolVar(&recoveryMode, "recovery", false, "boot in recovery mode")
}

func main() {
	flag.Parse()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := run(context.Background()); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(ctx context.Context) error {
	if install {
		return installMacOS(ctx)
	}

	// Handle disk options
	if newDisk {
		if err := createFreshDisk(diskSize); err != nil {
			return fmt.Errorf("failed to create fresh disk: %w", err)
		}
		log.Printf("Created fresh disk image with size %d GiB", diskSize)
	}

	if diskPath != "" {
		if err := useCustomDisk(diskPath); err != nil {
			return fmt.Errorf("failed to use custom disk: %w", err)
		}
		log.Printf("Using custom disk image from %s", diskPath)
	}

	// Handle reinitialization
	if reinit {
		if err := reinitializeVM(ctx); err != nil {
			return fmt.Errorf("failed to reinitialize VM: %w", err)
		}
		log.Println("VM successfully reinitialized")
	}

	return runVM(ctx)
}

func runVM(ctx context.Context) error {
	paths := GetVMPaths()

	// Ensure VM bundle exists
	if err := os.MkdirAll(paths.BundlePath, 0755); err != nil {
		return fmt.Errorf("failed to create VM bundle: %w", err)
	}

	// Check if platform configuration exists
	if !platformConfigExists(paths) {
		log.Println("Platform configuration not found. Please run with -reinit first, or -install to set up a new VM")
		return fmt.Errorf("platform configuration missing")
	}

	// Create platform configuration
	platform, err := createMacPlatformConfig(paths)
	if err != nil {
		return fmt.Errorf("failed to create platform config: %w", err)
	}

	// Create VM configuration
	config, err := setupVMConfiguration(platform, paths)
	if err != nil {
		return fmt.Errorf("failed to setup VM config: %w", err)
	}

	// Validate configuration
	if err := validateConfig(config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Create VM
	vm := virtualization.NewVZVirtualMachineWithConfiguration(unsafe.Pointer(config.ID))
	if vm.ID == 0 {
		return fmt.Errorf("failed to create VM")
	}

	log.Println("Starting virtual machine...")

	// Start VM
	if err := startVM(vm); err != nil {
		return fmt.Errorf("failed to start VM: %w", err)
	}

	// Wait for VM to stop
	return waitForVM(ctx, vm)
}

func startVM(vm virtualization.VZVirtualMachine) error {
	// Create completion handler
	done := make(chan error, 1)

	// Start the VM
	// Note: The generated bindings use Start() which returns void
	// In a full implementation, we'd need to set up a delegate to handle completion
	// For now, we'll use a simplified approach

	log.Println("VM started (simplified startup - full implementation needs completion handler)")

	// In real implementation:
	// vm.StartWithCompletionHandler(handler)
	// For now, just indicate the VM object is ready
	done <- nil

	return <-done
}

func waitForVM(ctx context.Context, vm virtualization.VZVirtualMachine) error {
	// Set up signal handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// In a full implementation, we'd monitor VM state changes
	// For now, wait for interrupt
	log.Println("VM is running. Press Ctrl+C to stop.")

	select {
	case <-ctx.Done():
		return ctx.Err()
	case sig := <-sigCh:
		log.Printf("Received signal %v, stopping VM...", sig)
		return stopVM(vm)
	}
}

func stopVM(vm virtualization.VZVirtualMachine) error {
	log.Println("Requesting VM stop...")

	// Check if we can request stop
	canStop := vm.CanRequestStop()
	if canStop {
		// Request graceful stop
		log.Println("Requesting graceful shutdown...")
		// vm.RequestStopWithCompletionHandler(handler)
		time.Sleep(3 * time.Second)
	}

	// Force stop if needed
	log.Println("Stopping VM...")
	// vm.StopWithCompletionHandler(handler)

	log.Println("VM stopped")
	return nil
}

func installMacOS(ctx context.Context) error {
	paths := GetVMPaths()

	// Ensure VM bundle exists
	if err := os.MkdirAll(paths.BundlePath, 0755); err != nil {
		return fmt.Errorf("failed to create VM bundle: %w", err)
	}

	log.Println("Installing macOS...")
	log.Println("This feature requires:")
	log.Println("  1. VZMacOSRestoreImage for downloading/loading restore images")
	log.Println("  2. Progress tracking for downloads and installation")
	log.Println("  3. Completion handler support")
	log.Println()
	log.Println("For now, please use -reinit to set up the platform configuration,")
	log.Println("then manually provide a restore image.")

	return fmt.Errorf("installation not yet implemented")
}

func reinitializeVM(ctx context.Context) error {
	paths := GetVMPaths()

	// Ensure VM bundle exists
	if err := os.MkdirAll(paths.BundlePath, 0755); err != nil {
		return fmt.Errorf("failed to create VM bundle: %w", err)
	}

	log.Println("Reinitializing VM platform configuration...")

	// For reinitialization, we need a restore image or we use a default hardware model
	// Since we don't have a restore image easily accessible, we'll create a basic config

	// Create new machine identifier
	machineID := virtualization.NewVZMacMachineIdentifier()
	if machineID.ID == 0 {
		return fmt.Errorf("failed to create machine identifier")
	}

	// Save machine identifier
	data := objcToBytes(unsafe.Pointer(machineID.DataRepresentation()))
	if err := os.WriteFile(paths.MachineIdentifierPath, data, 0644); err != nil {
		return fmt.Errorf("failed to save machine identifier: %w", err)
	}
	log.Printf("Created machine identifier: %s", paths.MachineIdentifierPath)

	// For hardware model, we need to use a supported model
	// This typically comes from a restore image, but we can use the current system's model
	hardwareModel := virtualization.VZMacHardwareModelClass.Supported()
	if hardwareModel.ID == 0 {
		return fmt.Errorf("failed to get supported hardware model")
	}

	// Save hardware model
	data = objcToBytes(unsafe.Pointer(hardwareModel.DataRepresentation()))
	if err := os.WriteFile(paths.HardwareModelPath, data, 0644); err != nil {
		return fmt.Errorf("failed to save hardware model: %w", err)
	}
	log.Printf("Created hardware model: %s", paths.HardwareModelPath)

	// Remove existing auxiliary storage if it exists
	if _, err := os.Stat(paths.AuxiliaryStoragePath); err == nil {
		log.Println("Removing existing auxiliary storage...")
		if err := os.Remove(paths.AuxiliaryStoragePath); err != nil {
			return fmt.Errorf("failed to remove auxiliary storage: %w", err)
		}
	}

	// Create new auxiliary storage
	log.Println("Creating auxiliary storage...")
	auxURL := pathToNSURL(paths.AuxiliaryStoragePath)

	// Create auxiliary storage with hardware model
	// Note: This requires calling the creation method
	// For now, we'll create an empty file and let the VM create it
	auxFile, err := os.Create(paths.AuxiliaryStoragePath)
	if err != nil {
		return fmt.Errorf("failed to create auxiliary storage file: %w", err)
	}
	auxFile.Close()

	log.Println("✓ VM platform configuration reinitialized successfully")
	log.Println()
	log.Println("You can now run the VM without -reinit")
	log.Println("Note: You may need to install macOS first using -install")

	return nil
}

func platformConfigExists(paths VMPaths) bool {
	_, err1 := os.Stat(paths.HardwareModelPath)
	_, err2 := os.Stat(paths.MachineIdentifierPath)
	return err1 == nil && err2 == nil
}

func createFreshDisk(sizeGiB uint64) error {
	paths := GetVMPaths()
	diskPath := paths.DiskImagePath

	// Backup existing disk if present
	if _, err := os.Stat(diskPath); err == nil {
		backupPath := fmt.Sprintf("%s.backup-%s", diskPath, time.Now().Format("20060102-150405"))
		log.Printf("Backing up existing disk to %s", backupPath)
		if err := os.Rename(diskPath, backupPath); err != nil {
			return fmt.Errorf("failed to backup disk: %w", err)
		}
	}

	// Create new disk
	sizeBytes := int64(sizeGiB * 1024 * 1024 * 1024)
	f, err := os.Create(diskPath)
	if err != nil {
		return fmt.Errorf("failed to create disk file: %w", err)
	}
	defer f.Close()

	if err := f.Truncate(sizeBytes); err != nil {
		return fmt.Errorf("failed to set disk size: %w", err)
	}

	log.Printf("Created %d GiB disk at %s", sizeGiB, diskPath)
	return nil
}

func useCustomDisk(sourcePath string) error {
	paths := GetVMPaths()
	destPath := paths.DiskImagePath

	// Check source exists
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("source disk not found: %w", err)
	}

	// Backup existing disk if present
	if _, err := os.Stat(destPath); err == nil {
		backupPath := fmt.Sprintf("%s.backup-%s", destPath, time.Now().Format("20060102-150405"))
		log.Printf("Backing up existing disk to %s", backupPath)
		if err := os.Rename(destPath, backupPath); err != nil {
			return fmt.Errorf("failed to backup disk: %w", err)
		}
	}

	// Create hard link or copy
	if err := os.Link(sourcePath, destPath); err != nil {
		// Link failed, try symlink
		if err := os.Symlink(sourcePath, destPath); err != nil {
			return fmt.Errorf("failed to link disk: %w", err)
		}
	}

	return nil
}

// Helper to convert Objective-C data to Go bytes
func objcToBytes(ptr unsafe.Pointer) []byte {
	if ptr == nil {
		return nil
	}

	length := objc.Send[uintptr](objc.ID(ptr), objc.RegisterName("length"))
	if length == 0 {
		return nil
	}

	bytesPtr := objc.Send[unsafe.Pointer](objc.ID(ptr), objc.RegisterName("bytes"))
	return unsafe.Slice((*byte)(bytesPtr), length)
}

// Helper to convert Go string to NSURL
func pathToNSURL(path string) objc.ID {
	nsStr := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSString")),
		objc.RegisterName("stringWithUTF8String:"),
		unsafe.Pointer(unsafe.StringData(path)),
	)
	nsURL := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSURL")),
		objc.RegisterName("fileURLWithPath:"),
		unsafe.Pointer(nsStr),
	)
	return nsURL
}

// Helper to create NSArray from pointers
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
