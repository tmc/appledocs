// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MESampleLocation] class.
var (
	MESampleLocationClass     _MESampleLocationClass
	MESampleLocationClassOnce sync.Once
)

func getMESampleLocationClass() _MESampleLocationClass {
	MESampleLocationClassOnce.Do(func() {
		MESampleLocationClass = _MESampleLocationClass{objc.GetClass("MESampleLocation")}
	})
	return MESampleLocationClass
}

type _MESampleLocationClass struct {
	class objc.Class
}

// An interface definition for the [MESampleLocation] class.
type IMESampleLocation interface {
	objectivec.IObject
	ByteSource() MEByteSource
	SampleLocation() unsafe.Pointer
}

// An object that provides information about the sample location with the media.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation
type MESampleLocation struct {
	objectivec.Object
}

// MESampleLocationFrom constructs a [MESampleLocation] from an unsafe.Pointer.
//
// An object that provides information about the sample location with the media.
func MESampleLocationFrom(ptr unsafe.Pointer) MESampleLocation {
	return MESampleLocation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MESampleLocationClass) Alloc() MESampleLocation {
	rv := objc.Send[MESampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MESampleLocationClass) New() MESampleLocation {
	rv := objc.Send[MESampleLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MESampleLocation) Init() MESampleLocation {
	rv := objc.Send[MESampleLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MESampleLocation) Autorelease() MESampleLocation {
	rv := objc.Send[MESampleLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMESampleLocation creates a new MESampleLocation instance.
func NewMESampleLocation() MESampleLocation {
	return getMESampleLocationClass().New()
}




// Creates a sample location object with the byte source and sample location that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/init(byteSource:sampleLocation:)
func NewMESampleLocationWithByteSourceSampleLocation(byteSource IMEByteSource, sampleLocation unsafe.Pointer) MESampleLocation {
	instance := getMESampleLocationClass().Alloc()
	rv := objc.Send[MESampleLocation](instance.ID, objc.Sel("initWithByteSource:sampleLocation:"), byteSource, sampleLocation)
	rv.Autorelease()
	return rv
}


// The byte source to use to read the data for the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/byteSource
func (m_ MESampleLocation) ByteSource() MEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}

// The starting file offset and size in bytes of the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/sampleLocation
func (m_ MESampleLocation) SampleLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sampleLocation"))
	return rv
}


