// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [exceptionHangingMask] class.
var (
	ExceptionHangingMaskClass     _exceptionHangingMaskClass
	ExceptionHangingMaskClassOnce sync.Once
)

func getexceptionHangingMaskClass() _exceptionHangingMaskClass {
	ExceptionHangingMaskClassOnce.Do(func() {
		ExceptionHangingMaskClass = _exceptionHangingMaskClass{objc.GetClass("exceptionHangingMask")}
	})
	return ExceptionHangingMaskClass
}

type _exceptionHangingMaskClass struct {
	class objc.Class
}

// An interface definition for the [exceptionHangingMask] class.
type IexceptionHangingMask interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/exceptionHangingMask
type exceptionHangingMask struct {
	objectivec.Object
}

// exceptionHangingMaskFrom constructs a [exceptionHangingMask] from an unsafe.Pointer.
func exceptionHangingMaskFrom(ptr unsafe.Pointer) exceptionHangingMask {
	return exceptionHangingMask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _exceptionHangingMaskClass) Alloc() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _exceptionHangingMaskClass) New() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ exceptionHangingMask) Init() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ exceptionHangingMask) Autorelease() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewexceptionHangingMask creates a new exceptionHangingMask instance.
func NewexceptionHangingMask() exceptionHangingMask {
	return getexceptionHangingMaskClass().New()
}




