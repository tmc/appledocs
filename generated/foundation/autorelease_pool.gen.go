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

// An interface definition for the [AutoreleasePool] class.
type IAutoreleasePool interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (ac _AutoreleasePoolClass) Alloc() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AutoreleasePoolClass) New() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutoreleasePool) Init() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutoreleasePool) Autorelease() AutoreleasePool {
	rv := objc.Send[AutoreleasePool](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutoreleasePool creates a new AutoreleasePool instance.
func NewAutoreleasePool() AutoreleasePool {
	return autoreleasePoolClass.New()
}




