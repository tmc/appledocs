// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [METrackInfo] class.
var (
	METrackInfoClass     _METrackInfoClass
	METrackInfoClassOnce sync.Once
)

func getMETrackInfoClass() _METrackInfoClass {
	METrackInfoClassOnce.Do(func() {
		METrackInfoClass = _METrackInfoClass{objc.GetClass("METrackInfo")}
	})
	return METrackInfoClass
}

type _METrackInfoClass struct {
	class objc.Class
}

// An interface definition for the [METrackInfo] class.
type IMETrackInfo interface {
	objectivec.IObject
}

// An object that includes track properties parsed from the media asset.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo
type METrackInfo struct {
	objectivec.Object
}

// METrackInfoFrom constructs a [METrackInfo] from an unsafe.Pointer.
//
// An object that includes track properties parsed from the media asset.
func METrackInfoFrom(ptr unsafe.Pointer) METrackInfo {
	return METrackInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _METrackInfoClass) Alloc() METrackInfo {
	rv := objc.Send[METrackInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _METrackInfoClass) New() METrackInfo {
	rv := objc.Send[METrackInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ METrackInfo) Init() METrackInfo {
	rv := objc.Send[METrackInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ METrackInfo) Autorelease() METrackInfo {
	rv := objc.Send[METrackInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMETrackInfo creates a new METrackInfo instance.
func NewMETrackInfo() METrackInfo {
	return getMETrackInfoClass().New()
}




// Creates a new track info object with the media type, track ID, and format descriptions that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/initWithMediaType:trackID:formatDescriptions:
func NewMETrackInfoWithMediaTypeTrackIDFormatDescriptions(mediaType unsafe.Pointer, trackID unsafe.Pointer, formatDescriptions objc.ID) METrackInfo {
	instance := getMETrackInfoClass().Alloc()
	rv := objc.Send[METrackInfo](instance.ID, objc.Sel("initWithMediaType:trackID:formatDescriptions:"), mediaType, trackID, formatDescriptions)
	rv.Autorelease()
	return rv
}


// The natural timescale of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) NaturalTimescale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("naturalTimescale"))
	return rv
}


// SetNaturalTimescale sets the value of the naturalTimescale property.
// The natural timescale of the track.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) SetNaturalTimescale(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalTimescale:"), value)
}

// The frame rate of the track in frames per second, as a 32-bit floating point number.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) NominalFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// SetNominalFrameRate sets the value of the nominalFrameRate property.
// The frame rate of the track in frames per second, as a 32-bit floating point number.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) SetNominalFrameRate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalFrameRate:"), value)
}

// An array of edit segments for the given track.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) TrackEdits() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](m_.ID, objc.Sel("trackEdits"))
	return rv
}


// SetTrackEdits sets the value of the trackEdits property.
// An array of edit segments for the given track.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) SetTrackEdits(value []unsafe.Pointer) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackEdits:"), nsArray)
}


