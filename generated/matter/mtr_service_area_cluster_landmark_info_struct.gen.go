// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterLandmarkInfoStruct] class.
var (
	MTRServiceAreaClusterLandmarkInfoStructClass     _MTRServiceAreaClusterLandmarkInfoStructClass
	MTRServiceAreaClusterLandmarkInfoStructClassOnce sync.Once
)

func getMTRServiceAreaClusterLandmarkInfoStructClass() _MTRServiceAreaClusterLandmarkInfoStructClass {
	MTRServiceAreaClusterLandmarkInfoStructClassOnce.Do(func() {
		MTRServiceAreaClusterLandmarkInfoStructClass = _MTRServiceAreaClusterLandmarkInfoStructClass{objc.GetClass("MTRServiceAreaClusterLandmarkInfoStruct")}
	})
	return MTRServiceAreaClusterLandmarkInfoStructClass
}

type _MTRServiceAreaClusterLandmarkInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterLandmarkInfoStruct] class.
type IMTRServiceAreaClusterLandmarkInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct
type MTRServiceAreaClusterLandmarkInfoStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterLandmarkInfoStructFrom constructs a [MTRServiceAreaClusterLandmarkInfoStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterLandmarkInfoStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterLandmarkInfoStruct {
	return MTRServiceAreaClusterLandmarkInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterLandmarkInfoStructClass) Alloc() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterLandmarkInfoStructClass) New() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) Init() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) Autorelease() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterLandmarkInfoStruct creates a new MTRServiceAreaClusterLandmarkInfoStruct instance.
func NewMTRServiceAreaClusterLandmarkInfoStruct() MTRServiceAreaClusterLandmarkInfoStruct {
	return getMTRServiceAreaClusterLandmarkInfoStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/landmarkTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) LandmarkTag() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("landmarkTag"))
	return rv
}


// SetLandmarkTag sets the value of the landmarkTag property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/landmarkTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) SetLandmarkTag(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLandmarkTag:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/relativePositionTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) RelativePositionTag() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("relativePositionTag"))
	return rv
}


// SetRelativePositionTag sets the value of the relativePositionTag property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/relativePositionTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) SetRelativePositionTag(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRelativePositionTag:"), value)
}



