// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ptr] class.
var (
	PtrClass     _ptrClass
	PtrClassOnce sync.Once
)

func getptrClass() _ptrClass {
	PtrClassOnce.Do(func() {
		PtrClass = _ptrClass{objc.GetClass("ptr")}
	})
	return PtrClass
}

type _ptrClass struct {
	class objc.Class
}





// An interface definition for the [ptr] class.
type Iptr interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _ptrClass) Alloc() ptr {
	rv := objc.Send[ptr](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ptrClass) New() ptr {
	rv := objc.Send[ptr](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ptr) Init() ptr {
	rv := objc.Send[ptr](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ptr) Autorelease() ptr {
	rv := objc.Send[ptr](p_.ID, objc.Sel("autorelease"))
	return rv
}

// Newptr creates a new ptr instance.
func Newptr() ptr {
	return getptrClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/ptr
type ptr struct {
	objectivec.Object
}

// ptrFrom constructs a [ptr] from an unsafe.Pointer.
func ptrFrom(ptr unsafe.Pointer) ptr {
	return ptr{objectivec.Object{objc.ID(ptr)}}
}































