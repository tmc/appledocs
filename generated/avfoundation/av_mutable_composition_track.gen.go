// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableCompositionTrack] class.
var (
	MutableCompositionTrackClass     _MutableCompositionTrackClass
	MutableCompositionTrackClassOnce sync.Once
)

func getMutableCompositionTrackClass() _MutableCompositionTrackClass {
	MutableCompositionTrackClassOnce.Do(func() {
		MutableCompositionTrackClass = _MutableCompositionTrackClass{objc.GetClass("AVMutableCompositionTrack")}
	})
	return MutableCompositionTrackClass
}

type _MutableCompositionTrackClass struct {
	class objc.Class
}





// An interface definition for the [MutableCompositionTrack] class.
type IMutableCompositionTrack interface {
	ICompositionTrack
	

	// properties:
	ExtendedLanguageTag() foundation.foundation.INSString
	SetExtendedLanguageTag(value foundation.foundation.INSString)
	Enabled() bool
	SetEnabled(value bool)
	LanguageCode() foundation.foundation.INSString
	SetLanguageCode(value foundation.foundation.INSString)
	NaturalTimeScale() TimeScale /* not a class type */
	SetNaturalTimeScale(value TimeScale /* not a class type */)
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	Segments() []CompositionTrackSegment
	SetSegments(value []CompositionTrackSegment)
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:
	AddTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType TrackAssociationType)
	InsertEmptyTimeRange(timeRange objectivec.IObject)
	InsertTimeRangeOfTrackAtTimeError(timeRange objectivec.IObject, track IAVAssetTrack, startTime objectivec.IObject, outError foundation.foundation.INSError) bool
	InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.Value, tracks []AssetTrack, startTime objectivec.IObject, outError foundation.foundation.INSError) bool
	RemoveTimeRange(timeRange objectivec.IObject)
	RemoveTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType TrackAssociationType)
	ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription FormatDescriptionRef /* not a class type */, replacementFormatDescription FormatDescriptionRef /* not a class type */)
	ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject)
	ValidateTrackSegmentsError(trackSegments []CompositionTrackSegment, outError foundation.foundation.INSError) bool


}





// Alloc allocates a new instance without initialization.
func (mc _MutableCompositionTrackClass) Alloc() MutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableCompositionTrackClass) New() MutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableCompositionTrack) Init() MutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableCompositionTrack) Autorelease() MutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableCompositionTrack creates a new MutableCompositionTrack instance.
func NewMutableCompositionTrack() MutableCompositionTrack {
	return getMutableCompositionTrackClass().New()
}





// A mutable track in a composition that you use to insert, remove, and scale track segments without affecting their low-level representation.
//
// Use this object to define constraints for the temporal arrangement of the track segments. If you set the composition’s track segments, you can test whether they meet the constraints by calling the method.


// A mutable track in a composition that you use to insert, remove, and scale track segments without affecting their low-level representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack
type MutableCompositionTrack struct {
	CompositionTrack
}

// MutableCompositionTrackFrom constructs a [MutableCompositionTrack] from an unsafe.Pointer.
//
// A mutable track in a composition that you use to insert, remove, and scale track segments without affecting their low-level representation.
func MutableCompositionTrackFrom(ptr unsafe.Pointer) MutableCompositionTrack {
	return MutableCompositionTrack{
		CompositionTrack: CompositionTrackFrom(ptr),
	}
}




















// Establishes a track association of a specific type between two tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/addTrackAssociation(to:type:)
func (m_ MutableCompositionTrack) AddTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType TrackAssociationType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addTrackAssociationToTrack:type:"), compositionTrack, trackAssociationType)
}


// Adds or extends an empty time range within the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertEmptyTimeRange(_:)
func (m_ MutableCompositionTrack) InsertEmptyTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}


// Inserts a time range of media from a source track into a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertTimeRange(_:of:at:)
func (m_ MutableCompositionTrack) InsertTimeRangeOfTrackAtTimeError(timeRange objectivec.IObject, track IAVAssetTrack, startTime objectivec.IObject, outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertTimeRange:ofTrack:atTime:error:"), timeRange, track, startTime, outError)
	return rv
}


// Inserts the time ranges of multiple source tracks into a track of a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertTimeRanges(_:of:at:)
func (m_ MutableCompositionTrack) InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.Value, tracks []AssetTrack, startTime objectivec.IObject, outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertTimeRanges:ofTracks:atTime:error:"), timeRanges, tracks, startTime, outError)
	return rv
}


// Removes a time range of media from a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/removeTimeRange(_:)
func (m_ MutableCompositionTrack) RemoveTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTimeRange:"), timeRange)
}


// Removes an association from a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/removeTrackAssociation(to:type:)
func (m_ MutableCompositionTrack) RemoveTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType TrackAssociationType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTrackAssociationToTrack:type:"), compositionTrack, trackAssociationType)
}


// Replaces a format description with another or cancels a previous replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/replaceFormatDescription(_:with:)
func (m_ MutableCompositionTrack) ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription FormatDescriptionRef /* not a class type */, replacementFormatDescription FormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceFormatDescription:withFormatDescription:"), originalFormatDescription, replacementFormatDescription)
}


// Changes the duration of a time range of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/scaleTimeRange(_:toDuration:)
func (m_ MutableCompositionTrack) ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}


// Returns a Boolean value that indicates whether a given array of track segments conform to the timing rules for a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/validateSegments(_:)
func (m_ MutableCompositionTrack) ValidateTrackSegmentsError(trackSegments []CompositionTrackSegment, outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("validateTrackSegments:error:"), trackSegments, outError)
	return rv
}







// The language tag associated with the track, as an RFC 4646 language tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/extendedLanguageTag
func (m_ MutableCompositionTrack) ExtendedLanguageTag() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The language tag associated with the track, as an RFC 4646 language tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/extendedLanguageTag
func (m_ MutableCompositionTrack) SetExtendedLanguageTag(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A Boolean value that indicates whether the tracks is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/isEnabled
func (m_ MutableCompositionTrack) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the tracks is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/isEnabled
func (m_ MutableCompositionTrack) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}


// The language associated with the track, as an ISO 639-2/T language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/languageCode
func (m_ MutableCompositionTrack) LanguageCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("languageCode"))
	return rv
}


// The language associated with the track, as an ISO 639-2/T language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/languageCode
func (m_ MutableCompositionTrack) SetLanguageCode(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLanguageCode:"), value)
}


// The time scale in which you can perform time-based operations without extra numerical conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/naturalTimeScale
func (m_ MutableCompositionTrack) NaturalTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](m_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The time scale in which you can perform time-based operations without extra numerical conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/naturalTimeScale
func (m_ MutableCompositionTrack) SetNaturalTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalTimeScale:"), value)
}


// The preferred transformation of the visual media data for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/preferredTransform
func (m_ MutableCompositionTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The preferred transformation of the visual media data for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/preferredTransform
func (m_ MutableCompositionTrack) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The volume the track prefers for its audible media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/preferredVolume
func (m_ MutableCompositionTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The volume the track prefers for its audible media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/preferredVolume
func (m_ MutableCompositionTrack) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredVolume:"), value)
}


// The track segments that a composition track contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/segments
func (m_ MutableCompositionTrack) Segments() []CompositionTrackSegment {
	rv := objc.Send[[]CompositionTrackSegment](m_.ID, objc.Sel("segments"))
	return rv
}


// The track segments that a composition track contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/segments
func (m_ MutableCompositionTrack) SetSegments(value []CompositionTrackSegment) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegments:"), nsArray)
}


// A Boolean value that indicates whether the tracks is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/isenabled
func (m_ MutableCompositionTrack) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the tracks is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/isenabled
func (m_ MutableCompositionTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}








