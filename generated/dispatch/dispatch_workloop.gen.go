// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DispatchWorkloop] class.
var (
	DispatchWorkloopClass     _DispatchWorkloopClass
	DispatchWorkloopClassOnce sync.Once
)

func getDispatchWorkloopClass() _DispatchWorkloopClass {
	DispatchWorkloopClassOnce.Do(func() {
		DispatchWorkloopClass = _DispatchWorkloopClass{objc.GetClass("DispatchWorkloop")}
	})
	return DispatchWorkloopClass
}

type _DispatchWorkloopClass struct {
	class objc.Class
}

// An interface definition for the [DispatchWorkloop] class.
type IDispatchWorkloop interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchWorkloop
type DispatchWorkloop struct {
	objectivec.Object
}

// DispatchWorkloopFrom constructs a [DispatchWorkloop] from an unsafe.Pointer.
func DispatchWorkloopFrom(ptr unsafe.Pointer) DispatchWorkloop {
	return DispatchWorkloop{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchWorkloopClass) Alloc() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchWorkloopClass) New() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchWorkloop) Init() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchWorkloop) Autorelease() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchWorkloop creates a new DispatchWorkloop instance.
func NewDispatchWorkloop() DispatchWorkloop {
	return getDispatchWorkloopClass().New()
}




