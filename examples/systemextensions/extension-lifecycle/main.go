package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/systemextensions"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("SystemExtensions Framework Example")
	fmt.Println("===================================")

	// Prevent unused import error
	_ = systemextensions.NewOSSystemExtensionRequest

	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   SystemExtensions enables:")
	fmt.Println("   - Installing and managing system extensions")
	fmt.Println("   - Network extensions")
	fmt.Println("   - Endpoint security extensions")
	fmt.Println("   - Driver extensions")

	fmt.Println("\n2. Key Classes:")
	fmt.Println("   OSSystemExtensionRequest - Request to install/uninstall")
	fmt.Println("   OSSystemExtensionManager - Manages extension lifecycle")

	fmt.Println("\n3. Extension Types:")
	fmt.Println("   - Network Extensions (content filtering, DNS proxy)")
	fmt.Println("   - Endpoint Security (system monitoring)")
	fmt.Println("   - DriverKit Extensions (USB, PCI, HID drivers)")

	fmt.Println("\n4. Requirements:")
	fmt.Println("   - macOS 10.15+")
	fmt.Println("   - System Extension entitlement")
	fmt.Println("   - User approval required")

	fmt.Println("\n✓ SystemExtensions overview completed!")
}
