// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDataTypeLocationDescriptorStruct] class.
var (
	MTRDataTypeLocationDescriptorStructClass     _MTRDataTypeLocationDescriptorStructClass
	MTRDataTypeLocationDescriptorStructClassOnce sync.Once
)

func getMTRDataTypeLocationDescriptorStructClass() _MTRDataTypeLocationDescriptorStructClass {
	MTRDataTypeLocationDescriptorStructClassOnce.Do(func() {
		MTRDataTypeLocationDescriptorStructClass = _MTRDataTypeLocationDescriptorStructClass{objc.GetClass("MTRDataTypeLocationDescriptorStruct")}
	})
	return MTRDataTypeLocationDescriptorStructClass
}

type _MTRDataTypeLocationDescriptorStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDataTypeLocationDescriptorStruct] class.
type IMTRDataTypeLocationDescriptorStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct
type MTRDataTypeLocationDescriptorStruct struct {
	objectivec.Object
}

// MTRDataTypeLocationDescriptorStructFrom constructs a [MTRDataTypeLocationDescriptorStruct] from an unsafe.Pointer.
func MTRDataTypeLocationDescriptorStructFrom(ptr unsafe.Pointer) MTRDataTypeLocationDescriptorStruct {
	return MTRDataTypeLocationDescriptorStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDataTypeLocationDescriptorStructClass) Alloc() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDataTypeLocationDescriptorStructClass) New() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDataTypeLocationDescriptorStruct) Init() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDataTypeLocationDescriptorStruct) Autorelease() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDataTypeLocationDescriptorStruct creates a new MTRDataTypeLocationDescriptorStruct instance.
func NewMTRDataTypeLocationDescriptorStruct() MTRDataTypeLocationDescriptorStruct {
	return getMTRDataTypeLocationDescriptorStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/areaType
func (m_ MTRDataTypeLocationDescriptorStruct) AreaType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("areaType"))
	return rv
}


// SetAreaType sets the value of the areaType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/areaType
func (m_ MTRDataTypeLocationDescriptorStruct) SetAreaType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/floorNumber
func (m_ MTRDataTypeLocationDescriptorStruct) FloorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("floorNumber"))
	return rv
}


// SetFloorNumber sets the value of the floorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/floorNumber
func (m_ MTRDataTypeLocationDescriptorStruct) SetFloorNumber(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFloorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/locationName
func (m_ MTRDataTypeLocationDescriptorStruct) LocationName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("locationName"))
	return rv
}


// SetLocationName sets the value of the locationName property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/locationName
func (m_ MTRDataTypeLocationDescriptorStruct) SetLocationName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocationName:"), objc.String(value))
}



