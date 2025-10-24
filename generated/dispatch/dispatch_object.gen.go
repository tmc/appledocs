// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DispatchObject] class.
var (
	DispatchObjectClass     _DispatchObjectClass
	DispatchObjectClassOnce sync.Once
)

func getDispatchObjectClass() _DispatchObjectClass {
	DispatchObjectClassOnce.Do(func() {
		DispatchObjectClass = _DispatchObjectClass{objc.GetClass("DispatchObject")}
	})
	return DispatchObjectClass
}

type _DispatchObjectClass struct {
	class objc.Class
}

// An interface definition for the [DispatchObject] class.
type IDispatchObject interface {
	IOS_object
	// properties:
	// methods:
}

// The base class for most dispatch types.
//
// There are many types of dispatch objects, including , , and . The base dispatch object interfaces allow you to manage memory, pause and resume execution, define object context, log task data, and more.


// The base class for most dispatch types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject
type DispatchObject struct {
	OS_object
}

// DispatchObjectFrom constructs a [DispatchObject] from an unsafe.Pointer.
//
// The base class for most dispatch types.
func DispatchObjectFrom(ptr unsafe.Pointer) DispatchObject {
	return DispatchObject{
		OS_object: OS_objectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchObjectClass) Alloc() DispatchObject {
	rv := objc.Send[DispatchObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchObjectClass) New() DispatchObject {
	rv := objc.Send[DispatchObject](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchObject) Init() DispatchObject {
	rv := objc.Send[DispatchObject](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchObject) Autorelease() DispatchObject {
	rv := objc.Send[DispatchObject](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchObject creates a new DispatchObject instance.
func NewDispatchObject() DispatchObject {
	return getDispatchObjectClass().New()
}




