// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterAreaStruct] class.
var (
	MTRServiceAreaClusterAreaStructClass     _MTRServiceAreaClusterAreaStructClass
	MTRServiceAreaClusterAreaStructClassOnce sync.Once
)

func getMTRServiceAreaClusterAreaStructClass() _MTRServiceAreaClusterAreaStructClass {
	MTRServiceAreaClusterAreaStructClassOnce.Do(func() {
		MTRServiceAreaClusterAreaStructClass = _MTRServiceAreaClusterAreaStructClass{objc.GetClass("MTRServiceAreaClusterAreaStruct")}
	})
	return MTRServiceAreaClusterAreaStructClass
}

type _MTRServiceAreaClusterAreaStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterAreaStruct] class.
type IMTRServiceAreaClusterAreaStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct
type MTRServiceAreaClusterAreaStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterAreaStructFrom constructs a [MTRServiceAreaClusterAreaStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterAreaStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterAreaStruct {
	return MTRServiceAreaClusterAreaStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterAreaStructClass) Alloc() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterAreaStructClass) New() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterAreaStruct) Init() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterAreaStruct) Autorelease() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterAreaStruct creates a new MTRServiceAreaClusterAreaStruct instance.
func NewMTRServiceAreaClusterAreaStruct() MTRServiceAreaClusterAreaStruct {
	return getMTRServiceAreaClusterAreaStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaID
func (m_ MTRServiceAreaClusterAreaStruct) AreaID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("areaID"))
	return rv
}


// SetAreaID sets the value of the areaID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaID
func (m_ MTRServiceAreaClusterAreaStruct) SetAreaID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaInfo
func (m_ MTRServiceAreaClusterAreaStruct) AreaInfo() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("areaInfo"))
	return rv
}


// SetAreaInfo sets the value of the areaInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaInfo
func (m_ MTRServiceAreaClusterAreaStruct) SetAreaInfo(value IMTRServiceAreaClusterAreaInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaInfo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/mapID
func (m_ MTRServiceAreaClusterAreaStruct) MapID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mapID"))
	return rv
}


// SetMapID sets the value of the mapID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/mapID
func (m_ MTRServiceAreaClusterAreaStruct) SetMapID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapID:"), value)
}



