// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceTypeRevision] class.
var (
	MTRDeviceTypeRevisionClass     _MTRDeviceTypeRevisionClass
	MTRDeviceTypeRevisionClassOnce sync.Once
)

func getMTRDeviceTypeRevisionClass() _MTRDeviceTypeRevisionClass {
	MTRDeviceTypeRevisionClassOnce.Do(func() {
		MTRDeviceTypeRevisionClass = _MTRDeviceTypeRevisionClass{objc.GetClass("MTRDeviceTypeRevision")}
	})
	return MTRDeviceTypeRevisionClass
}

type _MTRDeviceTypeRevisionClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceTypeRevision] class.
type IMTRDeviceTypeRevision interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceTypeRevision
type MTRDeviceTypeRevision struct {
	objectivec.Object
}

// MTRDeviceTypeRevisionFrom constructs a [MTRDeviceTypeRevision] from an unsafe.Pointer.
func MTRDeviceTypeRevisionFrom(ptr unsafe.Pointer) MTRDeviceTypeRevision {
	return MTRDeviceTypeRevision{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceTypeRevisionClass) Alloc() MTRDeviceTypeRevision {
	rv := objc.Send[MTRDeviceTypeRevision](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceTypeRevisionClass) New() MTRDeviceTypeRevision {
	rv := objc.Send[MTRDeviceTypeRevision](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceTypeRevision) Init() MTRDeviceTypeRevision {
	rv := objc.Send[MTRDeviceTypeRevision](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceTypeRevision) Autorelease() MTRDeviceTypeRevision {
	rv := objc.Send[MTRDeviceTypeRevision](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceTypeRevision creates a new MTRDeviceTypeRevision instance.
func NewMTRDeviceTypeRevision() MTRDeviceTypeRevision {
	return getMTRDeviceTypeRevisionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/devicetypeid
func (m_ MTRDeviceTypeRevision) DeviceTypeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("deviceTypeID"))
	return rv
}


// SetDeviceTypeID sets the value of the deviceTypeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/devicetypeid
func (m_ MTRDeviceTypeRevision) SetDeviceTypeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceTypeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/devicetyperevision
func (m_ MTRDeviceTypeRevision) DeviceTypeRevision() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("deviceTypeRevision"))
	return rv
}


// SetDeviceTypeRevision sets the value of the deviceTypeRevision property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/devicetyperevision
func (m_ MTRDeviceTypeRevision) SetDeviceTypeRevision(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceTypeRevision:"), value)
}

// Returns the MTRDeviceType corresponding to deviceTypeID,
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/typeinformation
func (m_ MTRDeviceTypeRevision) TypeInformation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("typeInformation"))
	return rv
}


// SetTypeInformation sets the value of the typeInformation property.
// Returns the MTRDeviceType corresponding to deviceTypeID,

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetyperevision/typeinformation
func (m_ MTRDeviceTypeRevision) SetTypeInformation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTypeInformation:"), value)
}



