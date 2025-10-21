// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRServiceAreaClusterAreaInfoStruct] class.
var (
	MTRServiceAreaClusterAreaInfoStructClass     _MTRServiceAreaClusterAreaInfoStructClass
	MTRServiceAreaClusterAreaInfoStructClassOnce sync.Once
)

func getMTRServiceAreaClusterAreaInfoStructClass() _MTRServiceAreaClusterAreaInfoStructClass {
	MTRServiceAreaClusterAreaInfoStructClassOnce.Do(func() {
		MTRServiceAreaClusterAreaInfoStructClass = _MTRServiceAreaClusterAreaInfoStructClass{objc.GetClass("MTRServiceAreaClusterAreaInfoStruct")}
	})
	return MTRServiceAreaClusterAreaInfoStructClass
}

type _MTRServiceAreaClusterAreaInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterAreaInfoStruct] class.
type IMTRServiceAreaClusterAreaInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct
type MTRServiceAreaClusterAreaInfoStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterAreaInfoStructFrom constructs a [MTRServiceAreaClusterAreaInfoStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterAreaInfoStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterAreaInfoStruct {
	return MTRServiceAreaClusterAreaInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterAreaInfoStructClass) Alloc() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterAreaInfoStructClass) New() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterAreaInfoStruct) Init() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterAreaInfoStruct) Autorelease() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterAreaInfoStruct creates a new MTRServiceAreaClusterAreaInfoStruct instance.
func NewMTRServiceAreaClusterAreaInfoStruct() MTRServiceAreaClusterAreaInfoStruct {
	return getMTRServiceAreaClusterAreaInfoStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/landmarkInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) LandmarkInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("landmarkInfo"))
	return rv
}


// SetLandmarkInfo sets the value of the landmarkInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/landmarkInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) SetLandmarkInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLandmarkInfo:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/locationInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) LocationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("locationInfo"))
	return rv
}


// SetLocationInfo sets the value of the locationInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/locationInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) SetLocationInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocationInfo:"), value)
}


