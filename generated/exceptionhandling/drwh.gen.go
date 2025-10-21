// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [drwh] class.
var (
	DrwhClass     _drwhClass
	DrwhClassOnce sync.Once
)

func getdrwhClass() _drwhClass {
	DrwhClassOnce.Do(func() {
		DrwhClass = _drwhClass{objc.GetClass("drwh")}
	})
	return DrwhClass
}

type _drwhClass struct {
	class objc.Class
}

// An interface definition for the [drwh] class.
type Idrwh interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/drwh
type drwh struct {
	objectivec.Object
}

// drwhFrom constructs a [drwh] from an unsafe.Pointer.
func drwhFrom(ptr unsafe.Pointer) drwh {
	return drwh{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _drwhClass) Alloc() drwh {
	rv := objc.Send[drwh](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _drwhClass) New() drwh {
	rv := objc.Send[drwh](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ drwh) Init() drwh {
	rv := objc.Send[drwh](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ drwh) Autorelease() drwh {
	rv := objc.Send[drwh](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdrwh creates a new drwh instance.
func Newdrwh() drwh {
	return getdrwhClass().New()
}




