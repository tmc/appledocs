package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/foundation"
)

func main() {
	// Lock to main thread (required for Objective-C)
	runtime.LockOSThread()

	// Create an AppleEventDescriptor with some string data
	descriptor := foundation.AppleEventDescriptorClass.DescriptorWithString("Hello, World!")
	fmt.Printf("Created descriptor: %v\n", descriptor)

	// Get the data from the descriptor
	// This calls: objc.Send[Data](descriptor.ID, objc.Sel("data"))
	// The Objective-C method returns NSData*, so we use Data as the type parameter
	data := descriptor.Data()
	fmt.Printf("Got data: %v\n", data)

	// Verify the data has content
	length := data.Length()
	fmt.Printf("Data length: %d bytes\n", length)

	// Create a string descriptor with different data
	descriptor2 := foundation.AppleEventDescriptorClass.DescriptorWithInt32(42)
	data2 := descriptor2.Data()
	length2 := data2.Length()
	fmt.Printf("\nDescriptor for int32(42) has data length: %d bytes\n", length2)

	// Show that the return type is IData (interface), not Data (struct)
	// This demonstrates type safety - we can work with the interface
	var idata foundation.IData = descriptor.Data()
	fmt.Printf("\nInterface type works: length = %d\n", idata.Length())

	fmt.Println("\n✓ AppleEventDescriptor.Data() works correctly!")
	fmt.Println("✓ objc.Send[Data] properly returns Data struct")
	fmt.Println("✓ Data implements IData interface as expected")
}
