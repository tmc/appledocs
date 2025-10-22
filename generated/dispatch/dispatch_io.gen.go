// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DispatchIO] class.
var (
	DispatchIOClass     _DispatchIOClass
	DispatchIOClassOnce sync.Once
)

func getDispatchIOClass() _DispatchIOClass {
	DispatchIOClassOnce.Do(func() {
		DispatchIOClass = _DispatchIOClass{objc.GetClass("DispatchIO")}
	})
	return DispatchIOClass
}

type _DispatchIOClass struct {
	class objc.Class
}

// An interface definition for the [DispatchIO] class.
type IDispatchIO interface {
	objectivec.IObject
}

// An object that manages operations on a file descriptor using either stream-based or random-access semantics.


// An object that manages operations on a file descriptor using either stream-based or random-access semantics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO

type DispatchIO struct {
	objectivec.Object
}

// DispatchIOFrom constructs a [DispatchIO] from an unsafe.Pointer.
//
// An object that manages operations on a file descriptor using either stream-based or random-access semantics.
func DispatchIOFrom(ptr unsafe.Pointer) DispatchIO {
	return DispatchIO{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchIOClass) Alloc() DispatchIO {
	rv := objc.Send[DispatchIO](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchIOClass) New() DispatchIO {
	rv := objc.Send[DispatchIO](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchIO) Init() DispatchIO {
	rv := objc.Send[DispatchIO](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchIO) Autorelease() DispatchIO {
	rv := objc.Send[DispatchIO](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchIO creates a new DispatchIO instance.
func NewDispatchIO() DispatchIO {
	return getDispatchIOClass().New()
}




