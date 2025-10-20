// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [customAttributes] class.
var (
	CustomAttributesClass     _customAttributesClass
	CustomAttributesClassOnce sync.Once
)

func getcustomAttributesClass() _customAttributesClass {
	CustomAttributesClassOnce.Do(func() {
		CustomAttributesClass = _customAttributesClass{objc.GetClass("customAttributes")}
	})
	return CustomAttributesClass
}

type _customAttributesClass struct {
	class objc.Class
}

// An interface definition for the [customAttributes] class.
type IcustomAttributes interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-c.ivar
type customAttributes struct {
	objectivec.Object
}

// customAttributesFrom constructs a [customAttributes] from an unsafe.Pointer.
func customAttributesFrom(ptr unsafe.Pointer) customAttributes {
	return customAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _customAttributesClass) Alloc() customAttributes {
	rv := objc.Send[customAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _customAttributesClass) New() customAttributes {
	rv := objc.Send[customAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customAttributes) Init() customAttributes {
	rv := objc.Send[customAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customAttributes) Autorelease() customAttributes {
	rv := objc.Send[customAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomAttributes creates a new customAttributes instance.
func NewcustomAttributes() customAttributes {
	return getcustomAttributesClass().New()
}




