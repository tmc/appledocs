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
	// properties:
	EstimatedSampleLocation() SampleCursorStorageRange /* not a class type */
	RefinementDataLocation() SampleCursorStorageRange /* not a class type */
	ByteSource() IMEByteSource
	SetByteSource(value IMEByteSource)
	// methods:
}

// An object that provides information about the estimated sample location with the media.


// An object that provides information about the estimated sample location with the media.
//
// [Full Topic]
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



// Creates an estimated sample location object with the byte source, sample location, and data location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/init(byteSource:estimatedSampleLocation:refinementDataLocation:)
func NewMEEstimatedSampleLocationWithByteSourceEstimatedSampleLocationRefinementDataLocation(byteSource IMEByteSource, estimatedSampleLocation SampleCursorStorageRange /* not a class type */, refinementDataLocation SampleCursorStorageRange /* not a class type */) MEEstimatedSampleLocation {
	instance := getMEEstimatedSampleLocationClass().Alloc()
	rv := objc.Send[MEEstimatedSampleLocation](instance.ID, objc.Sel("initWithByteSource:estimatedSampleLocation:refinementDataLocation:"), byteSource, estimatedSampleLocation, refinementDataLocation)
	rv.Autorelease()
	return rv
}



// The estimated starting file offset and size in bytes of the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/estimatedSampleLocation
func (m_ MEEstimatedSampleLocation) EstimatedSampleLocation() SampleCursorStorageRange /* not a class type */ {
	rv := objc.Send[SampleCursorStorageRange](m_.ID, objc.Sel("estimatedSampleLocation"))
	return rv
}


// The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/refinementDataLocation
func (m_ MEEstimatedSampleLocation) RefinementDataLocation() SampleCursorStorageRange /* not a class type */ {
	rv := objc.Send[SampleCursorStorageRange](m_.ID, objc.Sel("refinementDataLocation"))
	return rv
}


// The byte source to use to read the data for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/meestimatedsamplelocation/bytesource
func (m_ MEEstimatedSampleLocation) ByteSource() IMEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}


// The byte source to use to read the data for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/meestimatedsamplelocation/bytesource
func (m_ MEEstimatedSampleLocation) SetByteSource(value IMEByteSource) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setByteSource:"), value)
}


