// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	NaturalTimescale() TimeScale /* not a class type */
	SetNaturalTimescale(value TimeScale /* not a class type */)
	NominalFrameRate() unsafe.Pointer
	SetNominalFrameRate(value unsafe.Pointer)
	TrackEdits() []objc.IObject /* cross-framework: Value */
	SetTrackEdits(value []objc.IObject /* cross-framework: Value */)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	IsEnabled() bool
	SetIsEnabled(value bool)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	NaturalSize() objc.IObject /* cross-framework: Size */
	SetNaturalSize(value objc.IObject /* cross-framework: Size */)
	PreferredTransform() objc.IObject /* cross-framework: AffineTransform */
	SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */)
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)
	// methods:
}

// An object that includes track properties parsed from the media asset.


// An object that includes track properties parsed from the media asset.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/initWithMediaType:trackID:formatDescriptions:
func NewMETrackInfoWithMediaTypeTrackIDFormatDescriptions(mediaType MediaType /* not a class type */, trackID PersistentTrackID /* not a class type */, formatDescriptions objc.IObject /* cross-framework: NSArray */) METrackInfo {
	instance := getMETrackInfoClass().Alloc()
	rv := objc.Send[METrackInfo](instance.ID, objc.Sel("initWithMediaType:trackID:formatDescriptions:"), mediaType, trackID, formatDescriptions)
	rv.Autorelease()
	return rv
}



// The natural timescale of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) NaturalTimescale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](m_.ID, objc.Sel("naturalTimescale"))
	return rv
}


// The natural timescale of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) SetNaturalTimescale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalTimescale:"), value)
}


// The frame rate of the track in frames per second, as a 32-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) NominalFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The frame rate of the track in frames per second, as a 32-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) SetNominalFrameRate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalFrameRate:"), value)
}


// An array of edit segments for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) TrackEdits() []objc.IObject /* cross-framework: Value */ {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("trackEdits"))
	return rv
}


// An array of edit segments for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) SetTrackEdits(value []objc.IObject /* cross-framework: Value */) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/extendedlanguagetag
func (m_ METrackInfo) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/extendedlanguagetag
func (m_ METrackInfo) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}


// The media type of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/mediatype
func (m_ METrackInfo) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}


// The media type of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/mediatype
func (m_ METrackInfo) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaType:"), value)
}


// Indicates the natural dimensions of the media data referenced by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/naturalsize
func (m_ METrackInfo) NaturalSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// Indicates the natural dimensions of the media data referenced by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/naturalsize
func (m_ METrackInfo) SetNaturalSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}


// Indicates the preferred affine display transform of the track media for visual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/preferredtransform
func (m_ METrackInfo) PreferredTransform() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[corefoundation.AffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}


// Indicates the preferred affine display transform of the track media for visual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/preferredtransform
func (m_ METrackInfo) SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}


// A Boolean value that indicates whether frame reordering occurs in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/requiresframereordering
func (m_ METrackInfo) RequiresFrameReordering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// A Boolean value that indicates whether frame reordering occurs in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/requiresframereordering
func (m_ METrackInfo) SetRequiresFrameReordering(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}


// An integer that identifies the track within the media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/trackid
func (m_ METrackInfo) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}


// An integer that identifies the track within the media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/trackid
func (m_ METrackInfo) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}


