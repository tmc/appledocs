package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/metal"
)

func main() {
	fmt.Println("Metal Device Information")
	fmt.Println("========================")

	// Get the system's default Metal device
	devicePtr := metal.MTLCreateSystemDefaultDevice()
	if devicePtr == nil {
		fmt.Println("Error: No Metal-capable device found")
		return
	}

	device := objc.ID(devicePtr)
	defer device.Send(objc.RegisterName("release"))

	// Get device name
	nameObj := device.Send(objc.RegisterName("name"))
	if nameObj != 0 {
		nameStr := objc.ID(nameObj).Send(objc.RegisterName("UTF8String"))
		if nameStr != 0 {
			name := (*byte)(unsafe.Pointer(nameStr))
			nameBytes := make([]byte, 0, 256)
			for i := 0; ; i++ {
				b := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(name)) + uintptr(i)))
				if b == 0 {
					break
				}
				nameBytes = append(nameBytes, b)
			}
			fmt.Printf("\nDevice Name: %s\n", string(nameBytes))
		}
	}

	// Check if device is headless (compute-only, no displays)
	headlessSel := objc.RegisterName("isHeadless")
	isHeadless := device.Send(headlessSel)
	fmt.Printf("Headless: %v\n", isHeadless != 0)

	// Check if device is low power (integrated GPU vs discrete)
	lowPowerSel := objc.RegisterName("isLowPower")
	isLowPower := device.Send(lowPowerSel)
	fmt.Printf("Low Power: %v\n", isLowPower != 0)

	// Check if device is removable (external GPU)
	removableSel := objc.RegisterName("isRemovable")
	isRemovable := device.Send(removableSel)
	fmt.Printf("Removable: %v\n", isRemovable != 0)

	// Get registry ID
	registryIDSel := objc.RegisterName("registryID")
	registryID := device.Send(registryIDSel)
	fmt.Printf("Registry ID: %d\n", registryID)

	fmt.Println("\nMetal device information retrieved successfully!")
	fmt.Println("\nNote: This example demonstrates basic device queries.")
	fmt.Println("More advanced queries (like maxThreadsPerThreadgroup, feature sets, etc.)")
	fmt.Println("require more complex struct handling and are best done through MetalKit.")
}
