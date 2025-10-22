// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
	NaturalTimescale() unsafe.Pointer
	SetNaturalTimescale(value unsafe.Pointer)
	NominalFrameRate() unsafe.Pointer
	SetNominalFrameRate(value unsafe.Pointer)
	TrackEdits() []foundation.Value
	SetTrackEdits(value []foundation.IValue)
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	IsEnabled() bool
	SetIsEnabled(value bool)
	MediaType() unsafe.Pointer
	SetMediaType(value unsafe.Pointer)
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	PreferredTransform() coregraphics.CGAffineTransform
	SetPreferredTransform(value coregraphics.CGAffineTransform)
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)
	TrackID() unsafe.Pointer
	SetTrackID(value unsafe.Pointer)
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
func NewMETrackInfoWithMediaTypeTrackIDFormatDescriptions(mediaType unsafe.Pointer, trackID unsafe.Pointer, formatDescriptions objectivec.IObject) METrackInfo {
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
func (m_ METrackInfo) TrackEdits() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("trackEdits"))
	return rv
}


// SetTrackEdits sets the value of the trackEdits property.
// An array of edit segments for the given track.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) SetTrackEdits(value []foundation.IValue) {
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

// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/extendedlanguagetag
func (m_ METrackInfo) ExtendedLanguageTag() string {
	rv := objc.Send[string](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// SetExtendedLanguageTag sets the value of the extendedLanguageTag property.
// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/extendedlanguagetag
func (m_ METrackInfo) SetExtendedLanguageTag(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}

// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the track is enabled by default.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}

// The media type of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/mediatype
func (m_ METrackInfo) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The media type of the track.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/mediatype
func (m_ METrackInfo) SetMediaType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaType:"), value)
}

// Indicates the natural dimensions of the media data referenced by the track.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/naturalsize
func (m_ METrackInfo) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// SetNaturalSize sets the value of the naturalSize property.
// Indicates the natural dimensions of the media data referenced by the track.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/naturalsize
func (m_ METrackInfo) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}

// Indicates the preferred affine display transform of the track media for visual display.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/preferredtransform
func (m_ METrackInfo) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}


// SetPreferredTransform sets the value of the preferredTransform property.
// Indicates the preferred affine display transform of the track media for visual display.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/preferredtransform
func (m_ METrackInfo) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}

// A Boolean value that indicates whether frame reordering occurs in the track.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/requiresframereordering
func (m_ METrackInfo) RequiresFrameReordering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// SetRequiresFrameReordering sets the value of the requiresFrameReordering property.
// A Boolean value that indicates whether frame reordering occurs in the track.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/requiresframereordering
func (m_ METrackInfo) SetRequiresFrameReordering(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}

// An integer that identifies the track within the media asset.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/trackid
func (m_ METrackInfo) TrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("trackID"))
	return rv
}


// SetTrackID sets the value of the trackID property.
// An integer that identifies the track within the media asset.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/trackid
func (m_ METrackInfo) SetTrackID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}


