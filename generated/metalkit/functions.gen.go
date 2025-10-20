// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MetalKit Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTKMetalVertexDescriptorFromModelIO func(unsafe.Pointer) unsafe.Pointer
	_MTKMetalVertexDescriptorFromModelIOWithError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTKMetalVertexFormatFromModelIO func(unsafe.Pointer) unsafe.Pointer
	_MTKModelIOVertexDescriptorFromMetal func(unsafe.Pointer) unsafe.Pointer
	_MTKModelIOVertexDescriptorFromMetalWithError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTKModelIOVertexFormatFromMetal func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTKMetalVertexDescriptorFromModelIO, lib, "MTKMetalVertexDescriptorFromModelIO")
	tryRegister(&_MTKMetalVertexDescriptorFromModelIOWithError, lib, "MTKMetalVertexDescriptorFromModelIOWithError")
	tryRegister(&_MTKMetalVertexFormatFromModelIO, lib, "MTKMetalVertexFormatFromModelIO")
	tryRegister(&_MTKModelIOVertexDescriptorFromMetal, lib, "MTKModelIOVertexDescriptorFromMetal")
	tryRegister(&_MTKModelIOVertexDescriptorFromMetalWithError, lib, "MTKModelIOVertexDescriptorFromMetalWithError")
	tryRegister(&_MTKModelIOVertexFormatFromMetal, lib, "MTKModelIOVertexFormatFromMetal")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns a partially converted Metal vertex descriptor. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIO(_:)
func MTKMetalVertexDescriptorFromModelIO(modelIODescriptor unsafe.Pointer) unsafe.Pointer {
	return _MTKMetalVertexDescriptorFromModelIO(modelIODescriptor)
	}


// Returns a partially converted Metal vertex descriptor, reporting any error that occurs. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIOWithError
func MTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _MTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor, error_)
	}


// Returns a converted Metal vertex format. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexFormatFromModelIO(_:)
func MTKMetalVertexFormatFromModelIO(vertexFormat unsafe.Pointer) unsafe.Pointer {
	return _MTKMetalVertexFormatFromModelIO(vertexFormat)
	}


// Returns a partially converted Model I/O vertex descriptor. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetal(_:)
func MTKModelIOVertexDescriptorFromMetal(metalDescriptor unsafe.Pointer) unsafe.Pointer {
	return _MTKModelIOVertexDescriptorFromMetal(metalDescriptor)
	}


// Returns a partially converted Model I/O vertex descriptor, reporting any error that occurs. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetalWithError
func MTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _MTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor, error_)
	}


// Returns a converted Model I/O vertex format. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexFormatFromMetal(_:)
func MTKModelIOVertexFormatFromMetal(vertexFormat unsafe.Pointer) unsafe.Pointer {
	return _MTKModelIOVertexFormatFromMetal(vertexFormat)
	}




