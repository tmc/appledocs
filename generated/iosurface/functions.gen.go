// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOSurface Functions (51 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOSurfaceDecrementUseCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceIncrementUseCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceAlignProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceAllowsPixelSizeCasting func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceCopyAllValues func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceCopyValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceCreate func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceCreateMachPort func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceCreateXPCObject func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetAllocSize func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBaseAddress func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBaseAddressOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBitDepthOfComponentOfPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBitOffsetOfComponentOfPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBytesPerElement func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBytesPerElementOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBytesPerRow func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetBytesPerRowOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetElementHeight func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetElementHeightOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetElementWidth func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetElementWidthOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetHeight func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetHeightOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetID func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetNameOfComponentOfPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetNumberOfComponentsOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetPixelFormat func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetPlaneCount func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetPropertyAlignment func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetPropertyMaximum func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetRangeOfComponentOfPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetSeed func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetSubsampling func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetTypeID func() unsafe.Pointer
	_IOSurfaceGetTypeOfComponentOfPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetUseCount func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetWidth func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceGetWidthOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceIsInUse func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceLock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceLookup func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceLookupFromMachPort func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceLookupFromXPCObject func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_IOSurfaceRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceSetOwnershipIdentity func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceSetPurgeable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceSetValues func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOSurfaceUnlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOSurfaceDecrementUseCount, lib, "IOSurfaceDecrementUseCount")
	tryRegister(&_IOSurfaceIncrementUseCount, lib, "IOSurfaceIncrementUseCount")
	tryRegister(&_IOSurfaceAlignProperty, lib, "IOSurfaceAlignProperty")
	tryRegister(&_IOSurfaceAllowsPixelSizeCasting, lib, "IOSurfaceAllowsPixelSizeCasting")
	tryRegister(&_IOSurfaceCopyAllValues, lib, "IOSurfaceCopyAllValues")
	tryRegister(&_IOSurfaceCopyValue, lib, "IOSurfaceCopyValue")
	tryRegister(&_IOSurfaceCreate, lib, "IOSurfaceCreate")
	tryRegister(&_IOSurfaceCreateMachPort, lib, "IOSurfaceCreateMachPort")
	tryRegister(&_IOSurfaceCreateXPCObject, lib, "IOSurfaceCreateXPCObject")
	tryRegister(&_IOSurfaceGetAllocSize, lib, "IOSurfaceGetAllocSize")
	tryRegister(&_IOSurfaceGetBaseAddress, lib, "IOSurfaceGetBaseAddress")
	tryRegister(&_IOSurfaceGetBaseAddressOfPlane, lib, "IOSurfaceGetBaseAddressOfPlane")
	tryRegister(&_IOSurfaceGetBitDepthOfComponentOfPlane, lib, "IOSurfaceGetBitDepthOfComponentOfPlane")
	tryRegister(&_IOSurfaceGetBitOffsetOfComponentOfPlane, lib, "IOSurfaceGetBitOffsetOfComponentOfPlane")
	tryRegister(&_IOSurfaceGetBytesPerElement, lib, "IOSurfaceGetBytesPerElement")
	tryRegister(&_IOSurfaceGetBytesPerElementOfPlane, lib, "IOSurfaceGetBytesPerElementOfPlane")
	tryRegister(&_IOSurfaceGetBytesPerRow, lib, "IOSurfaceGetBytesPerRow")
	tryRegister(&_IOSurfaceGetBytesPerRowOfPlane, lib, "IOSurfaceGetBytesPerRowOfPlane")
	tryRegister(&_IOSurfaceGetElementHeight, lib, "IOSurfaceGetElementHeight")
	tryRegister(&_IOSurfaceGetElementHeightOfPlane, lib, "IOSurfaceGetElementHeightOfPlane")
	tryRegister(&_IOSurfaceGetElementWidth, lib, "IOSurfaceGetElementWidth")
	tryRegister(&_IOSurfaceGetElementWidthOfPlane, lib, "IOSurfaceGetElementWidthOfPlane")
	tryRegister(&_IOSurfaceGetHeight, lib, "IOSurfaceGetHeight")
	tryRegister(&_IOSurfaceGetHeightOfPlane, lib, "IOSurfaceGetHeightOfPlane")
	tryRegister(&_IOSurfaceGetID, lib, "IOSurfaceGetID")
	tryRegister(&_IOSurfaceGetNameOfComponentOfPlane, lib, "IOSurfaceGetNameOfComponentOfPlane")
	tryRegister(&_IOSurfaceGetNumberOfComponentsOfPlane, lib, "IOSurfaceGetNumberOfComponentsOfPlane")
	tryRegister(&_IOSurfaceGetPixelFormat, lib, "IOSurfaceGetPixelFormat")
	tryRegister(&_IOSurfaceGetPlaneCount, lib, "IOSurfaceGetPlaneCount")
	tryRegister(&_IOSurfaceGetPropertyAlignment, lib, "IOSurfaceGetPropertyAlignment")
	tryRegister(&_IOSurfaceGetPropertyMaximum, lib, "IOSurfaceGetPropertyMaximum")
	tryRegister(&_IOSurfaceGetRangeOfComponentOfPlane, lib, "IOSurfaceGetRangeOfComponentOfPlane")
	tryRegister(&_IOSurfaceGetSeed, lib, "IOSurfaceGetSeed")
	tryRegister(&_IOSurfaceGetSubsampling, lib, "IOSurfaceGetSubsampling")
	tryRegister(&_IOSurfaceGetTypeID, lib, "IOSurfaceGetTypeID")
	tryRegister(&_IOSurfaceGetTypeOfComponentOfPlane, lib, "IOSurfaceGetTypeOfComponentOfPlane")
	tryRegister(&_IOSurfaceGetUseCount, lib, "IOSurfaceGetUseCount")
	tryRegister(&_IOSurfaceGetWidth, lib, "IOSurfaceGetWidth")
	tryRegister(&_IOSurfaceGetWidthOfPlane, lib, "IOSurfaceGetWidthOfPlane")
	tryRegister(&_IOSurfaceIsInUse, lib, "IOSurfaceIsInUse")
	tryRegister(&_IOSurfaceLock, lib, "IOSurfaceLock")
	tryRegister(&_IOSurfaceLookup, lib, "IOSurfaceLookup")
	tryRegister(&_IOSurfaceLookupFromMachPort, lib, "IOSurfaceLookupFromMachPort")
	tryRegister(&_IOSurfaceLookupFromXPCObject, lib, "IOSurfaceLookupFromXPCObject")
	tryRegister(&_IOSurfaceRemoveAllValues, lib, "IOSurfaceRemoveAllValues")
	tryRegister(&_IOSurfaceRemoveValue, lib, "IOSurfaceRemoveValue")
	tryRegister(&_IOSurfaceSetOwnershipIdentity, lib, "IOSurfaceSetOwnershipIdentity")
	tryRegister(&_IOSurfaceSetPurgeable, lib, "IOSurfaceSetPurgeable")
	tryRegister(&_IOSurfaceSetValue, lib, "IOSurfaceSetValue")
	tryRegister(&_IOSurfaceSetValues, lib, "IOSurfaceSetValues")
	tryRegister(&_IOSurfaceUnlock, lib, "IOSurfaceUnlock")
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



// Decrements the per-process usage count for an . [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/1419377-iosurfacedecrementusecount
func IOSurfaceDecrementUseCount(buffer unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceDecrementUseCount(buffer, p1)
	}


// Increments the per-process usage count for an . [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iosurface/1419455-iosurfaceincrementusecount
func IOSurfaceIncrementUseCount(buffer unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceIncrementUseCount(buffer, p1)
	}


// Returns the smallest aligned value greater than or equal to the specified value. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceAlignProperty(_:_:)
func IOSurfaceAlignProperty(property unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceAlignProperty(property, value)
	}


// IOSurfaceAllowsPixelSizeCasting is a IOSurface function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceAllowsPixelSizeCasting(_:)
func IOSurfaceAllowsPixelSizeCasting(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceAllowsPixelSizeCasting(buffer)
	}


// IOSurfaceCopyAllValues is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyAllValues(_:)
func IOSurfaceCopyAllValues(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceCopyAllValues(buffer)
	}


// Retrieves a value from the dictionary associated with the buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyValue(_:_:)
func IOSurfaceCopyValue(buffer unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceCopyValue(buffer, key)
	}


// Creates a brand new IOSurface object [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreate(_:)
func IOSurfaceCreate(properties unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceCreate(properties)
	}


// Returns a mach_port_t that holds a reference to the IOSurface. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateMachPort(_:)
func IOSurfaceCreateMachPort(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceCreateMachPort(buffer)
	}


// Returns an xpc_object_t that holds a reference to the IOSurface. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateXPCObject(_:)
func IOSurfaceCreateXPCObject(aSurface unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceCreateXPCObject(aSurface)
	}


// Returns the total allocation size of the buffer including all planes. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetAllocSize(_:)
func IOSurfaceGetAllocSize(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetAllocSize(buffer)
	}


// Returns the address of the first byte of data in a particular buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddress(_:)
func IOSurfaceGetBaseAddress(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBaseAddress(buffer)
	}


// Returns the address of the first byte of data in the specified plane. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddressOfPlane(_:_:)
func IOSurfaceGetBaseAddressOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBaseAddressOfPlane(buffer, planeIndex)
	}


// IOSurfaceGetBitDepthOfComponentOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitDepthOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitDepthOfComponentOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer, componentIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBitDepthOfComponentOfPlane(buffer, planeIndex, componentIndex)
	}


// IOSurfaceGetBitOffsetOfComponentOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitOffsetOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitOffsetOfComponentOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer, componentIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBitOffsetOfComponentOfPlane(buffer, planeIndex, componentIndex)
	}


// Returns the length (in bytes) of each element in a particular buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElement(_:)
func IOSurfaceGetBytesPerElement(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBytesPerElement(buffer)
	}


// Returns the size of each element (in bytes) in the specified plane. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElementOfPlane(_:_:)
func IOSurfaceGetBytesPerElementOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBytesPerElementOfPlane(buffer, planeIndex)
	}


// Returns the length (in bytes) of each row in a particular buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRow(_:)
func IOSurfaceGetBytesPerRow(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBytesPerRow(buffer)
	}


// Returns the size of each row (in bytes) in the specified plane. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRowOfPlane(_:_:)
func IOSurfaceGetBytesPerRowOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetBytesPerRowOfPlane(buffer, planeIndex)
	}


// Returns the height (in pixels) of each element in a particular buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeight(_:)
func IOSurfaceGetElementHeight(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetElementHeight(buffer)
	}


// Returns the height (in pixels) of each element in the specified plane. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeightOfPlane(_:_:)
func IOSurfaceGetElementHeightOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetElementHeightOfPlane(buffer, planeIndex)
	}


// Returns the width (in pixels) of each element in a particular buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidth(_:)
func IOSurfaceGetElementWidth(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetElementWidth(buffer)
	}


// Returns the width (in pixels) of each element in the specified plane. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidthOfPlane(_:_:)
func IOSurfaceGetElementWidthOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetElementWidthOfPlane(buffer, planeIndex)
	}


// Returns the height of the IOSurface buffer in pixels. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeight(_:)
func IOSurfaceGetHeight(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetHeight(buffer)
	}


// Returns the height of the specified plane (in pixels). [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeightOfPlane(_:_:)
func IOSurfaceGetHeightOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetHeightOfPlane(buffer, planeIndex)
	}


// Retrieves the unique value for an . [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetID(_:)
func IOSurfaceGetID(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetID(buffer)
	}


// IOSurfaceGetNameOfComponentOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNameOfComponentOfPlane(_:_:_:)
func IOSurfaceGetNameOfComponentOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer, componentIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetNameOfComponentOfPlane(buffer, planeIndex, componentIndex)
	}


// IOSurfaceGetNumberOfComponentsOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNumberOfComponentsOfPlane(_:_:)
func IOSurfaceGetNumberOfComponentsOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetNumberOfComponentsOfPlane(buffer, planeIndex)
	}


// Returns an unsigned integer that contains the traditional macOS buffer format. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPixelFormat(_:)
func IOSurfaceGetPixelFormat(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetPixelFormat(buffer)
	}


// IOSurfaceGetPlaneCount is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPlaneCount(_:)
func IOSurfaceGetPlaneCount(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetPlaneCount(buffer)
	}


// Returns the alignment requirements for a property (if any). [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyAlignment(_:)
func IOSurfaceGetPropertyAlignment(property unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetPropertyAlignment(property)
	}


// Returns the maximum value for a given property that is guaranteed to be compatible with all of the current devices (GPUs, etc.) in the system. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyMaximum(_:)
func IOSurfaceGetPropertyMaximum(property unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetPropertyMaximum(property)
	}


// IOSurfaceGetRangeOfComponentOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetRangeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetRangeOfComponentOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer, componentIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetRangeOfComponentOfPlane(buffer, planeIndex, componentIndex)
	}


// IOSurfaceGetSeed is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSeed(_:)
func IOSurfaceGetSeed(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetSeed(buffer)
	}


// IOSurfaceGetSubsampling is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSubsampling(_:)
func IOSurfaceGetSubsampling(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetSubsampling(buffer)
	}


// IOSurfaceGetTypeID is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeID()
func IOSurfaceGetTypeID() unsafe.Pointer {
	return _IOSurfaceGetTypeID()
	}


// IOSurfaceGetTypeOfComponentOfPlane is a IOSurface function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetTypeOfComponentOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer, componentIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetTypeOfComponentOfPlane(buffer, planeIndex, componentIndex)
	}


// Returns the per-process usage count for an . [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetUseCount(_:)
func IOSurfaceGetUseCount(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetUseCount(buffer)
	}


// Returns the width of the IOSurface buffer in pixels. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidth(_:)
func IOSurfaceGetWidth(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetWidth(buffer)
	}


// Returns the width of the specified plane (in pixels). [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidthOfPlane(_:_:)
func IOSurfaceGetWidthOfPlane(buffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceGetWidthOfPlane(buffer, planeIndex)
	}


// Returns true of an IOSurface is in use by any process in the system, otherwise false. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceIsInUse(_:)
func IOSurfaceIsInUse(buffer unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceIsInUse(buffer)
	}


// “Lock” an IOSurface for reading or writing. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLock(_:_:_:)
func IOSurfaceLock(buffer unsafe.Pointer, options unsafe.Pointer, seed unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceLock(buffer, options, seed)
	}


// Performs an atomic lookup and retain of an IOSurface by its IOSurfaceID. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookup(_:)
func IOSurfaceLookup(csid unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceLookup(csid)
	}


// Recreates an IOSurfaceRef from a mach port. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromMachPort(_:)
func IOSurfaceLookupFromMachPort(port unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceLookupFromMachPort(port)
	}


// IOSurfaceLookupFromXPCObject is a IOSurface function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromXPCObject(_:)
func IOSurfaceLookupFromXPCObject(xobj unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceLookupFromXPCObject(xobj)
	}


// IOSurfaceRemoveAllValues is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveAllValues(_:)
func IOSurfaceRemoveAllValues(buffer unsafe.Pointer) {
	_IOSurfaceRemoveAllValues(buffer)
	}


// Deletes a value in the dictionary associated with the buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveValue(_:_:)
func IOSurfaceRemoveValue(buffer unsafe.Pointer, key unsafe.Pointer) {
	_IOSurfaceRemoveValue(buffer, key)
	}


// IOSurfaceSetOwnershipIdentity is a IOSurface function. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetOwnershipIdentity(_:_:_:_:)
func IOSurfaceSetOwnershipIdentity(buffer unsafe.Pointer, task_id_token unsafe.Pointer, newLedgerTag int, newLedgerOptions unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceSetOwnershipIdentity(buffer, task_id_token, newLedgerTag, newLedgerOptions)
	}


// IOSurfaceSetPurgeable is a IOSurface function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetPurgeable(_:_:_:)
func IOSurfaceSetPurgeable(buffer unsafe.Pointer, newState unsafe.Pointer, oldState unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceSetPurgeable(buffer, newState, oldState)
	}


// Sets a value in the dictionary associated with the buffer. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValue(_:_:_:)
func IOSurfaceSetValue(buffer unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_IOSurfaceSetValue(buffer, key, value)
	}


// IOSurfaceSetValues is a IOSurface function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValues(_:_:)
func IOSurfaceSetValues(buffer unsafe.Pointer, keysAndValues unsafe.Pointer) {
	_IOSurfaceSetValues(buffer, keysAndValues)
	}


// “Unlock” an for reading or writing. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/IOSurface/IOSurfaceUnlock(_:_:_:)
func IOSurfaceUnlock(buffer unsafe.Pointer, options unsafe.Pointer, seed unsafe.Pointer) unsafe.Pointer {
	return _IOSurfaceUnlock(buffer, options, seed)
	}




