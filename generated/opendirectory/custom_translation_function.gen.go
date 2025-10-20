// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [customTranslationFunction] class.
var (
	CustomTranslationFunctionClass     _customTranslationFunctionClass
	CustomTranslationFunctionClassOnce sync.Once
)

func getcustomTranslationFunctionClass() _customTranslationFunctionClass {
	CustomTranslationFunctionClassOnce.Do(func() {
		CustomTranslationFunctionClass = _customTranslationFunctionClass{objc.GetClass("customTranslationFunction")}
	})
	return CustomTranslationFunctionClass
}

type _customTranslationFunctionClass struct {
	class objc.Class
}

// An interface definition for the [customTranslationFunction] class.
type IcustomTranslationFunction interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-c.ivar
type customTranslationFunction struct {
	objectivec.Object
}

// customTranslationFunctionFrom constructs a [customTranslationFunction] from an unsafe.Pointer.
func customTranslationFunctionFrom(ptr unsafe.Pointer) customTranslationFunction {
	return customTranslationFunction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _customTranslationFunctionClass) Alloc() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _customTranslationFunctionClass) New() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customTranslationFunction) Init() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customTranslationFunction) Autorelease() customTranslationFunction {
	rv := objc.Send[customTranslationFunction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomTranslationFunction creates a new customTranslationFunction instance.
func NewcustomTranslationFunction() customTranslationFunction {
	return getcustomTranslationFunctionClass().New()
}




