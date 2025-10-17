// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AutoreleasePool] class.
var autoreleasePoolClass = _AutoreleasePoolClass{objc.GetClass("NSAutoreleasePool")}

type _AutoreleasePoolClass struct {
	class objc.Class
}

// An object that supports Cocoa’s reference-counted memory management system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool

type AutoreleasePool struct {
	objectivec.Object
}

// AutoreleasePoolFrom constructs a [AutoreleasePool] from an unsafe.Pointer.
//
// An object that supports Cocoa’s reference-counted memory management system.
func AutoreleasePoolFrom(ptr unsafe.Pointer) AutoreleasePool {
	return AutoreleasePool{objectivec.Object{objc.ID(ptr)}}
}



