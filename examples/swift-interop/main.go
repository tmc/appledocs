// main.go - Call Swift functions from Go using purego

package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/ebitengine/purego"
)

func main() {
	// Open the Swift dynamic library
	lib, err := purego.Dlopen("./libhello.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		log.Fatalf("Failed to load libhello.dylib: %v", err)
	}

	// Register swift_hello (void function)
	var swiftHello func()
	purego.RegisterLibFunc(&swiftHello, lib, "swift_hello")

	fmt.Println("Calling swift_hello()...")
	swiftHello()

	// Register swift_add (function with parameters and return value)
	var swiftAdd func(int32, int32) int32
	purego.RegisterLibFunc(&swiftAdd, lib, "swift_add")

	result := swiftAdd(42, 13)
	fmt.Printf("swift_add(42, 13) = %d\n", result)

	// Register swift_greet (function with C string parameter)
	var swiftGreet func(*byte)
	purego.RegisterLibFunc(&swiftGreet, lib, "swift_greet")

	name := append([]byte("Go Developer"), 0) // null-terminated string
	fmt.Println("Calling swift_greet()...")
	swiftGreet(&name[0])

	// Register swift_get_message (function returning C string)
	var swiftGetMessage func() *byte
	purego.RegisterLibFunc(&swiftGetMessage, lib, "swift_get_message")

	messagePtr := swiftGetMessage()
	if messagePtr != nil {
		// Convert C string to Go string
		message := goCString(messagePtr)
		fmt.Printf("swift_get_message() = %q\n", message)
		// Note: In production, we should free the C string
		// but purego doesn't provide direct C free() access
	}
}

// goCString converts a null-terminated C string to a Go string
func goCString(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	var length int
	for {
		if *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(length))) == 0 {
			break
		}
		length++
	}
	return string(unsafe.Slice(ptr, length))
}
