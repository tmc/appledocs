// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceType] class.
var (
	MTRDeviceTypeClass     _MTRDeviceTypeClass
	MTRDeviceTypeClassOnce sync.Once
)

func getMTRDeviceTypeClass() _MTRDeviceTypeClass {
	MTRDeviceTypeClassOnce.Do(func() {
		MTRDeviceTypeClass = _MTRDeviceTypeClass{objc.GetClass("MTRDeviceType")}
	})
	return MTRDeviceTypeClass
}

type _MTRDeviceTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceType] class.
type IMTRDeviceType interface {
	objectivec.IObject
	Id() foundation.Number
	IsUtility() bool
	Name() string
}

// Meta-data about a device type defined in the Matter specification.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType
type MTRDeviceType struct {
	objectivec.Object
}

// MTRDeviceTypeFrom constructs a [MTRDeviceType] from an unsafe.Pointer.
//
// Meta-data about a device type defined in the Matter specification.
func MTRDeviceTypeFrom(ptr unsafe.Pointer) MTRDeviceType {
	return MTRDeviceType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceTypeClass) Alloc() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceTypeClass) New() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceType) Init() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceType) Autorelease() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceType creates a new MTRDeviceType instance.
func NewMTRDeviceType() MTRDeviceType {
	return getMTRDeviceTypeClass().New()
}




// Returns an MTRDeviceType for the given ID, if the ID is known. Returns nil for unknown IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/init(forID:)
func NewMTRDeviceTypeForID(deviceTypeID foundation.INumber) MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(getMTRDeviceTypeClass().class), objc.Sel("deviceTypeForID:"), deviceTypeID)
	return rv
}


// Returns an MTRDeviceType for the given ID, if the ID is known. Returns nil for unknown IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/init(forID:)
func (mc _MTRDeviceTypeClass) DeviceTypeForID(deviceTypeID foundation.INumber) MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("deviceTypeForID:"), deviceTypeID)
	return rv
}

// The identifier of the device type (32-bit unsigned integer).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/id
func (m_ MTRDeviceType) Id() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("id"))
	return rv
}

// Returns whether this is a utility device type.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/isUtility
func (m_ MTRDeviceType) IsUtility() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUtility"))
	return rv
}

// Returns the name of the device type.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/name
func (m_ MTRDeviceType) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


