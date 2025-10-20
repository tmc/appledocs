package main

import (
	"flag"
	"fmt"
	"runtime"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("LocalAuthenticationEmbeddedUI Framework Example")
	fmt.Println("================================================")

	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   LocalAuthenticationEmbeddedUI provides:")
	fmt.Println("   - Embedded biometric authentication UI")
	fmt.Println("   - Custom authentication views")
	fmt.Println("   - Integration with LocalAuthentication")

	fmt.Println("\n2. Key Features:")
	fmt.Println("   - LAAuthenticationView - Embedded auth UI")
	fmt.Println("   - Custom UI placement")
	fmt.Println("   - TouchID/FaceID integration")

	fmt.Println("\n3. Requirements:")
	fmt.Println("   - macOS 12.0+")
	fmt.Println("   - Biometric hardware")
	fmt.Println("   - UI integration")

	fmt.Println("\n✓ LocalAuthenticationEmbeddedUI overview completed!")
}
