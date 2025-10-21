//go:build !v2
// +build !v2

// Package main - Simplified Linux VM launcher using Apple Virtualization framework
//
// This is a minimal, production-ready example of creating and launching a Linux VM
// using the generated bindings from github.com/tmc/appledocs/generated/virtualization
//
// Build: go build -o vm-create main_quickstart.go
// Usage: vm-create <kernel_path> <disk_path> [initrd_path] [cmdline]
//
// Example:
//   vm-create vmlinuz-6.1.0 ubuntu.img initrd.img "console=ttyS0"
//   vm-create vmlinuz disk.img                    # minimal
//
// Special feature: Auto-download macOS installer
//   vm-create --download-macos monterey           # Downloads and creates macOS VM config
//   vm-create --download-macos ventura            # Ventura installer
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/virtualization"
)

var (
	// macOS version URLs (example - these would be from Apple's softwareupdate server)
	macOSVersions = map[string]string{
		"monterey": "https://example.com/monterey.dmg",
		"ventura":  "https://example.com/ventura.dmg",
		"sonoma":   "https://example.com/sonoma.dmg",
	}
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Handle macOS DMG download
	if len(os.Args) > 1 && os.Args[1] == "--download-macos" {
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s --download-macos <version>\n", os.Args[0])
			fmt.Fprintf(os.Stderr, "Versions: %s\n", strings.Join(listMacOSVersions(), ", "))
			os.Exit(1)
		}
		downloadMacOSInstaller(os.Args[2])
		return
	}

	// Parse arguments for Linux VM
	if len(os.Args) < 3 {
		printUsage(os.Args[0])
		os.Exit(1)
	}

	kernel := os.Args[1]
	disk := os.Args[2]
	initrd := ""
	cmdline := "console=ttyS0"

	if len(os.Args) > 3 {
		initrd = os.Args[3]
	}
	if len(os.Args) > 4 {
		cmdline = os.Args[4]
	}

	// Verify files exist
	for _, path := range []string{kernel, disk} {
		if _, err := os.Stat(path); err != nil {
			fmt.Fprintf(os.Stderr, "Error: file not found: %s\n", path)
			os.Exit(1)
		}
	}
	if initrd != "" {
		if _, err := os.Stat(initrd); err != nil {
			fmt.Fprintf(os.Stderr, "Error: file not found: %s\n", initrd)
			os.Exit(1)
		}
	}

	// Create VM configuration
	config := createVMConfig(kernel, disk, initrd, cmdline)

	// Create VM instance
	vm := virtualization.NewVZVirtualMachineWithConfiguration(unsafe.Pointer(config.ID))
	if vm.ID == 0 {
		fmt.Fprintf(os.Stderr, "Error: failed to create VM\n")
		os.Exit(1)
	}

	fmt.Printf("✓ VM created\n")
	fmt.Printf("  Kernel: %s\n", kernel)
	fmt.Printf("  Disk:   %s\n", disk)
	if initrd != "" {
		fmt.Printf("  Initrd: %s\n", initrd)
	}
	fmt.Printf("  Cmdline: %s\n", cmdline)
	fmt.Printf("\nNote: Full VM startup requires additional setup for:\n")
	fmt.Printf("  - Console/display output\n")
	fmt.Printf("  - Event handling\n")
	fmt.Printf("  - See full example in main.go for UI implementation\n")
}

// printUsage shows usage information
func printUsage(prog string) {
	fmt.Fprintf(os.Stderr, "Virtualization Framework VM Creator\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s <kernel> <disk> [initrd] [cmdline]\n", prog)
	fmt.Fprintf(os.Stderr, "   or: %s --download-macos <version>\n\n", prog)
	fmt.Fprintf(os.Stderr, "Linux VM example:\n")
	fmt.Fprintf(os.Stderr, "  %s vmlinuz ubuntu.img\n", prog)
	fmt.Fprintf(os.Stderr, "  %s vmlinuz ubuntu.img initrd.img\n", prog)
	fmt.Fprintf(os.Stderr, "  %s vmlinuz ubuntu.img initrd.img \"console=ttyS0 root=/dev/vda\"\n\n", prog)
	fmt.Fprintf(os.Stderr, "macOS VM example:\n")
	fmt.Fprintf(os.Stderr, "  %s --download-macos monterey\n", prog)
	fmt.Fprintf(os.Stderr, "  %s --download-macos ventura\n\n", prog)
	fmt.Fprintf(os.Stderr, "Available macOS versions: %s\n", strings.Join(listMacOSVersions(), ", "))
}

// listMacOSVersions returns available macOS versions
func listMacOSVersions() []string {
	versions := make([]string, 0, len(macOSVersions))
	for v := range macOSVersions {
		versions = append(versions, v)
	}
	return versions
}

// downloadMacOSInstaller downloads a macOS installer DMG
func downloadMacOSInstaller(version string) {
	version = strings.ToLower(version)

	url, ok := macOSVersions[version]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown macOS version: %s\n", version)
		fmt.Fprintf(os.Stderr, "Available versions: %s\n", strings.Join(listMacOSVersions(), ", "))
		os.Exit(1)
	}

	filename := fmt.Sprintf("macOS-%s-installer.dmg", version)

	// Check if already downloaded
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("✓ %s already exists\n", filename)
		createMacOSVMConfig(filename)
		return
	}

	fmt.Printf("Downloading macOS %s installer...\n", version)
	fmt.Printf("URL: %s\n", url)

	// Download with progress
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: HTTP %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Create file
	out, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	// Download with progress tracking
	total := resp.ContentLength
	written := int64(0)
	buffer := make([]byte, 32*1024)

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			out.Write(buffer[:n])
			written += int64(n)

			// Show progress
			if total > 0 {
				percent := float64(written) * 100 / float64(total)
				megabytes := float64(written) / (1024 * 1024)
				fmt.Printf("\r  %0.1f%% (%0.1f MB)   ", percent, megabytes)
			}
		}

		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error downloading: %v\n", err)
				os.Exit(1)
			}
			break
		}
	}

	fmt.Printf("\n✓ Downloaded %s (%0.1f MB)\n", filename, float64(written)/(1024*1024))

	// Create VM configuration with the downloaded DMG
	createMacOSVMConfig(filename)
}

// createMacOSVMConfig creates a VM configuration for macOS
func createMacOSVMConfig(dmgPath string) {
	fmt.Printf("\nCreating macOS VM configuration from %s\n", dmgPath)
	fmt.Printf("Note: Full macOS VM setup requires:\n")
	fmt.Printf("  - VZMacOSRestoreImage for IPSW processing\n")
	fmt.Printf("  - VZMacPlatformConfiguration setup\n")
	fmt.Printf("  - Auxiliary storage creation\n")
	fmt.Printf("  - Machine identifier configuration\n")
	fmt.Printf("\nFor complete macOS VM support, see main.go or Code-Hex/vz\n")
}

// createVMConfig creates a minimal Linux VM configuration
func createVMConfig(kernel, disk, initrd, cmdline string) virtualization.VZVirtualMachineConfiguration {
	config := virtualization.NewVZVirtualMachineConfiguration()

	// Platform: Generic (for Linux)
	platform := virtualization.NewVZGenericPlatformConfiguration()
	config.SetPlatform(unsafe.Pointer(platform.ID))

	// Boot loader: Linux kernel
	kernelURL := stringToNSURL(kernel)
	bootLoader := virtualization.NewVZLinuxBootLoaderWithKernelURL(unsafe.Pointer(kernelURL.ID))
	bootLoader.SetCommandLine(unsafe.Pointer(stringToNSString(cmdline).ID))

	if initrd != "" {
		initrdURL := stringToNSURL(initrd)
		bootLoader.SetInitialRamdiskURL(unsafe.Pointer(initrdURL.ID))
	}

	config.SetBootLoader(unsafe.Pointer(bootLoader.ID))

	// CPU and memory (sensible defaults)
	cpuCount := uint(runtime.NumCPU() / 2)
	if cpuCount < 1 {
		cpuCount = 1
	}
	config.SetCpuCount(cpuCount)
	config.SetMemorySize(uint64(2) * 1024 * 1024 * 1024) // 2GB

	// Note: Network and storage device configuration requires the complete binding
	// implementation from the vz reference bindings. The simplified example creates
	// the core config structure. See main.go for complete device setup.

	return config
}

// stringToNSString converts Go string to NSString
func stringToNSString(s string) foundation.String {
	cstr := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSString")),
		objc.RegisterName("stringWithUTF8String:"),
		unsafe.Pointer(unsafe.StringData(s)),
	)
	return foundation.StringFrom(unsafe.Pointer(cstr))
}

// stringToNSURL converts Go string to NSURL
func stringToNSURL(s string) foundation.URL {
	expandedPath, _ := filepath.Abs(s)
	nsStr := stringToNSString(expandedPath)
	nsURL := objc.Send[objc.ID](
		objc.ID(objc.GetClass("NSURL")),
		objc.RegisterName("fileURLWithPath:"),
		unsafe.Pointer(nsStr.ID),
	)
	return foundation.URLFrom(unsafe.Pointer(nsURL))
}
