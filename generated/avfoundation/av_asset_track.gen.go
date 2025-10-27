// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetTrack] class.
var (
	AssetTrackClass     _AssetTrackClass
	AssetTrackClassOnce sync.Once
)

func getAssetTrackClass() _AssetTrackClass {
	AssetTrackClassOnce.Do(func() {
		AssetTrackClass = _AssetTrackClass{objc.GetClass("AVAssetTrack")}
	})
	return AssetTrackClass
}

type _AssetTrackClass struct {
	class objc.Class
}





// An interface definition for the [AssetTrack] class.
type IAssetTrack interface {
	objectivec.IObject
	

	// properties:
	Asset() IAVAsset
	AvailableMetadataFormats() []string
	AvailableTrackAssociationTypes() []string
	CanProvideSampleCursors() bool
	CommonMetadata() []MetadataItem
	EstimatedDataRate() float32
	ExtendedLanguageTag() foundation.foundation.INSString
	FormatDescriptions() foundation.foundation.INSArray
	HasAudioSampleDependencies() bool
	Decodable() bool
	Enabled() bool
	Playable() bool
	SelfContained() bool
	LanguageCode() foundation.foundation.INSString
	MediaType() MediaType
	Metadata() []MetadataItem
	MinFrameDuration() objectivec.IObject
	NaturalSize() corefoundation.CGSize
	NaturalTimeScale() TimeScale /* not a class type */
	NominalFrameRate() float32
	PreferredTransform() corefoundation.CGAffineTransform
	PreferredVolume() float32
	RequiresFrameReordering() bool
	Segments() []AssetTrackSegment
	TimeRange() objectivec.IObject
	TotalSampleDataLength() objectivec.IObject
	TrackID() PersistentTrackID /* not a class type */
	IsDecodable() bool
	SetIsDecodable(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsPlayable() bool
	SetIsPlayable(value bool)
	IsSelfContained() bool
	SetIsSelfContained(value bool)


	

	// methods:
	LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType, completionHandler unsafe.Pointer)
	LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler unsafe.Pointer)
	LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime objectivec.IObject, completionHandler unsafe.Pointer)
	LoadSegmentForTrackTimeCompletionHandler(trackTime objectivec.IObject, completionHandler unsafe.Pointer)
	MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp objectivec.IObject) ISampleCursor
	MakeSampleCursorAtFirstSampleInDecodeOrder() ISampleCursor
	MakeSampleCursorAtLastSampleInDecodeOrder() ISampleCursor


}





// Alloc allocates a new instance without initialization.
func (ac _AssetTrackClass) Alloc() AssetTrack {
	rv := objc.Send[AssetTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetTrackClass) New() AssetTrack {
	rv := objc.Send[AssetTrack](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetTrack) Init() AssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetTrack) Autorelease() AssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetTrack creates a new AssetTrack instance.
func NewAssetTrack() AssetTrack {
	return getAssetTrackClass().New()
}





// An object that models a track of media that an asset contains.
//
// An asset contains one or more tracks of media that the framework models using the class. A track object holds the uniformly typed media that an asset provides such as audio, video, or closed captions. A track, like its containing , doesn’t load all of its media upon creation. Instead, it defers loading its data until you perform an operation that requires it. Because loading the data can take time, an asset track adopts the protocol so you can load its property values asynchronously by calling the method.


// An object that models a track of media that an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack
type AssetTrack struct {
	objectivec.Object
}

// AssetTrackFrom constructs a [AssetTrack] from an unsafe.Pointer.
//
// An object that models a track of media that an asset contains.
func AssetTrackFrom(ptr unsafe.Pointer) AssetTrack {
	return AssetTrack{objectivec.Object{objc.ID(ptr)}}
}




















// Loads associated tracks that have the specified association type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadAssociatedTracks(ofType:completionHandler:)
func (a_ AssetTrack) LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadAssociatedTracksOfType:completionHandler:"), trackAssociationType, completionHandler)
}


// Loads metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadMetadata(for:completionHandler:)
func (a_ AssetTrack) LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}


// Loads a sample presentation time that maps to the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSamplePresentationTime(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSamplePresentationTimeForTrackTime:completionHandler:"), trackTime, completionHandler)
}


// Loads a segment with a target time range that contains, or is closest to, the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSegment(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSegmentForTrackTimeCompletionHandler(trackTime objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSegmentForTrackTime:completionHandler:"), trackTime, completionHandler)
}


// Creates a sample cursor and positions it at or near the specified presentation timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursor(presentationTimeStamp:)
func (a_ AssetTrack) MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp objectivec.IObject) ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorWithPresentationTimeStamp:"), presentationTimeStamp)
	return rv
}


// Creates a sample cursor and positions it at the track’s first media sample in decode order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtFirstSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtFirstSampleInDecodeOrder() ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorAtFirstSampleInDecodeOrder"))
	return rv
}


// Creates a sample cursor and positions it at the track’s last media sample in decode order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtLastSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtLastSampleInDecodeOrder() ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorAtLastSampleInDecodeOrder"))
	return rv
}







// The asset object that contains this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/asset
func (a_ AssetTrack) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableMetadataFormats
func (a_ AssetTrack) AvailableMetadataFormats() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
func (a_ AssetTrack) AvailableTrackAssociationTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/canProvideSampleCursors
func (a_ AssetTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/commonMetadata
func (a_ AssetTrack) CommonMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/estimatedDataRate
func (a_ AssetTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/extendedLanguageTag
func (a_ AssetTrack) ExtendedLanguageTag() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/formatDescriptions
func (a_ AssetTrack) FormatDescriptions() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/hasAudioSampleDependencies
func (a_ AssetTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isDecodable
func (a_ AssetTrack) Decodable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("decodable"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isEnabled
func (a_ AssetTrack) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isPlayable
func (a_ AssetTrack) Playable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playable"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isSelfContained
func (a_ AssetTrack) SelfContained() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selfContained"))
	return rv
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/languageCode
func (a_ AssetTrack) LanguageCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("languageCode"))
	return rv
}


// The type of media that a track presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/mediaType
func (a_ AssetTrack) MediaType() MediaType {
	rv := objc.Send[MediaType](a_.ID, objc.Sel("mediaType"))
	return rv
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/metadata
func (a_ AssetTrack) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/minFrameDuration
func (a_ AssetTrack) MinFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalSize
func (a_ AssetTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalTimeScale
func (a_ AssetTrack) NaturalTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/nominalFrameRate
func (a_ AssetTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredTransform
func (a_ AssetTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredVolume
func (a_ AssetTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/requiresFrameReordering
func (a_ AssetTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/segments
func (a_ AssetTrack) Segments() []AssetTrackSegment {
	rv := objc.Send[[]AssetTrackSegment](a_.ID, objc.Sel("segments"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/timeRange
func (a_ AssetTrack) TimeRange() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/totalSampleDataLength
func (a_ AssetTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// The persistent unique identifier for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/trackID
func (a_ AssetTrack) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) IsDecodable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDecodable"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) SetIsDecodable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDecodable:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) IsPlayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) SetIsPlayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) IsSelfContained() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSelfContained"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) SetIsSelfContained(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSelfContained:"), value)
}








