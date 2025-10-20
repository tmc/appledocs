// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [customQueryFunction] class.
var (
	CustomQueryFunctionClass     _customQueryFunctionClass
	CustomQueryFunctionClassOnce sync.Once
)

func getcustomQueryFunctionClass() _customQueryFunctionClass {
	CustomQueryFunctionClassOnce.Do(func() {
		CustomQueryFunctionClass = _customQueryFunctionClass{objc.GetClass("customQueryFunction")}
	})
	return CustomQueryFunctionClass
}

type _customQueryFunctionClass struct {
	class objc.Class
}

// An interface definition for the [customQueryFunction] class.
type IcustomQueryFunction interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-c.ivar
type customQueryFunction struct {
	objectivec.Object
}

// customQueryFunctionFrom constructs a [customQueryFunction] from an unsafe.Pointer.
func customQueryFunctionFrom(ptr unsafe.Pointer) customQueryFunction {
	return customQueryFunction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _customQueryFunctionClass) Alloc() customQueryFunction {
	rv := objc.Send[customQueryFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _customQueryFunctionClass) New() customQueryFunction {
	rv := objc.Send[customQueryFunction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customQueryFunction) Init() customQueryFunction {
	rv := objc.Send[customQueryFunction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customQueryFunction) Autorelease() customQueryFunction {
	rv := objc.Send[customQueryFunction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomQueryFunction creates a new customQueryFunction instance.
func NewcustomQueryFunction() customQueryFunction {
	return getcustomQueryFunctionClass().New()
}




