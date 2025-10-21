// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEEstimatedSampleLocation] class.
var (
	MEEstimatedSampleLocationClass     _MEEstimatedSampleLocationClass
	MEEstimatedSampleLocationClassOnce sync.Once
)

func getMEEstimatedSampleLocationClass() _MEEstimatedSampleLocationClass {
	MEEstimatedSampleLocationClassOnce.Do(func() {
		MEEstimatedSampleLocationClass = _MEEstimatedSampleLocationClass{objc.GetClass("MEEstimatedSampleLocation")}
	})
	return MEEstimatedSampleLocationClass
}

type _MEEstimatedSampleLocationClass struct {
	class objc.Class
}

// An interface definition for the [MEEstimatedSampleLocation] class.
type IMEEstimatedSampleLocation interface {
	objectivec.IObject
}

// An object that provides information about the estimated sample location with the media.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation
type MEEstimatedSampleLocation struct {
	objectivec.Object
}

// MEEstimatedSampleLocationFrom constructs a [MEEstimatedSampleLocation] from an unsafe.Pointer.
//
// An object that provides information about the estimated sample location with the media.
func MEEstimatedSampleLocationFrom(ptr unsafe.Pointer) MEEstimatedSampleLocation {
	return MEEstimatedSampleLocation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEEstimatedSampleLocationClass) Alloc() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEEstimatedSampleLocationClass) New() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEEstimatedSampleLocation) Init() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEEstimatedSampleLocation) Autorelease() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEstimatedSampleLocation creates a new MEEstimatedSampleLocation instance.
func NewMEEstimatedSampleLocation() MEEstimatedSampleLocation {
	return getMEEstimatedSampleLocationClass().New()
}


// The byte source to use to read the data for the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/byteSource
func (m_ MEEstimatedSampleLocation) ByteSource() MEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}

// The estimated starting file offset and size in bytes of the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/estimatedSampleLocation
func (m_ MEEstimatedSampleLocation) EstimatedSampleLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("estimatedSampleLocation"))
	return rv
}

// The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/meestimatedsamplelocation/refinementdatalocation
func (m_ MEEstimatedSampleLocation) RefinementDataLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("refinementDataLocation"))
	return rv
}


// SetRefinementDataLocation sets the value of the refinementDataLocation property.
// The starting file offset and size in bytes of the data necessary to provide an accurate sample location.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/meestimatedsamplelocation/refinementdatalocation
func (m_ MEEstimatedSampleLocation) SetRefinementDataLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRefinementDataLocation:"), value)
}



