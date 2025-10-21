// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CBAttribute] class.
var (
	CBAttributeClass     _CBAttributeClass
	CBAttributeClassOnce sync.Once
)

func getCBAttributeClass() _CBAttributeClass {
	CBAttributeClassOnce.Do(func() {
		CBAttributeClass = _CBAttributeClass{objc.GetClass("CBAttribute")}
	})
	return CBAttributeClass
}

type _CBAttributeClass struct {
	class objc.Class
}

// An interface definition for the [CBAttribute] class.
type ICBAttribute interface {
	objectivec.IObject
}

// A representation of common aspects of services offered by a peripheral.
//
// Concrete subclasses of (and their mutable counterparts) represent the services a peripheral offers, the characteristics of those services, and the descriptors attached to those characteristics. The concrete subclasses are:
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute
type CBAttribute struct {
	objectivec.Object
}

// CBAttributeFrom constructs a [CBAttribute] from an unsafe.Pointer.
//
// A representation of common aspects of services offered by a peripheral.
func CBAttributeFrom(ptr unsafe.Pointer) CBAttribute {
	return CBAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBAttributeClass) Alloc() CBAttribute {
	rv := objc.Send[CBAttribute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBAttributeClass) New() CBAttribute {
	rv := objc.Send[CBAttribute](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBAttribute) Init() CBAttribute {
	rv := objc.Send[CBAttribute](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBAttribute) Autorelease() CBAttribute {
	rv := objc.Send[CBAttribute](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBAttribute creates a new CBAttribute instance.
func NewCBAttribute() CBAttribute {
	return getCBAttributeClass().New()
}


// The Bluetooth-specific UUID of the attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute/uuid
func (c_ CBAttribute) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("UUID"))
	return rv
}



