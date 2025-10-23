// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DispatchSource] class.
var (
	DispatchSourceClass     _DispatchSourceClass
	DispatchSourceClassOnce sync.Once
)

func getDispatchSourceClass() _DispatchSourceClass {
	DispatchSourceClassOnce.Do(func() {
		DispatchSourceClass = _DispatchSourceClass{objc.GetClass("DispatchSource")}
	})
	return DispatchSourceClass
}

type _DispatchSourceClass struct {
	class objc.Class
}

// An interface definition for the [DispatchSource] class.
type IDispatchSource interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that coordinates the processing of specific low-level system events, such as file-system events, timers, and UNIX signals.
//
// Use the methods of this class to construct new dispatch sources of the appropriate types.


// An object that coordinates the processing of specific low-level system events, such as file-system events, timers, and UNIX signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchSource
type DispatchSource struct {
	objectivec.Object
}

// DispatchSourceFrom constructs a [DispatchSource] from an unsafe.Pointer.
//
// An object that coordinates the processing of specific low-level system events, such as file-system events, timers, and UNIX signals.
func DispatchSourceFrom(ptr unsafe.Pointer) DispatchSource {
	return DispatchSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchSourceClass) Alloc() DispatchSource {
	rv := objc.Send[DispatchSource](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchSourceClass) New() DispatchSource {
	rv := objc.Send[DispatchSource](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchSource) Init() DispatchSource {
	rv := objc.Send[DispatchSource](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchSource) Autorelease() DispatchSource {
	rv := objc.Send[DispatchSource](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchSource creates a new DispatchSource instance.
func NewDispatchSource() DispatchSource {
	return getDispatchSourceClass().New()
}




