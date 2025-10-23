// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [eccVersion] class.
var (
	EccVersionClass     _eccVersionClass
	EccVersionClassOnce sync.Once
)

func geteccVersionClass() _eccVersionClass {
	EccVersionClassOnce.Do(func() {
		EccVersionClass = _eccVersionClass{objc.GetClass("eccVersion")}
	})
	return EccVersionClass
}

type _eccVersionClass struct {
	class objc.Class
}

// An interface definition for the [eccVersion] class.
type IeccVersion interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-c.ivar
type eccVersion struct {
	objectivec.Object
}

// eccVersionFrom constructs a [eccVersion] from an unsafe.Pointer.
func eccVersionFrom(ptr unsafe.Pointer) eccVersion {
	return eccVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _eccVersionClass) Alloc() eccVersion {
	rv := objc.Send[eccVersion](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _eccVersionClass) New() eccVersion {
	rv := objc.Send[eccVersion](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ eccVersion) Init() eccVersion {
	rv := objc.Send[eccVersion](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ eccVersion) Autorelease() eccVersion {
	rv := objc.Send[eccVersion](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NeweccVersion creates a new eccVersion instance.
func NeweccVersion() eccVersion {
	return geteccVersionClass().New()
}




