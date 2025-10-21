// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AssociatedTracksOfType(trackAssociationType unsafe.Pointer) []AssetTrack
	HasMediaCharacteristic(mediaCharacteristic unsafe.Pointer) bool
	LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadMetadataForFormatCompletionHandler(format unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadSegmentForTrackTimeCompletionHandler(trackTime unsafe.Pointer, completionHandler unsafe.Pointer)
	MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp unsafe.Pointer) unsafe.Pointer
	MakeSampleCursorAtFirstSampleInDecodeOrder() unsafe.Pointer
	MakeSampleCursorAtLastSampleInDecodeOrder() unsafe.Pointer
	MetadataForFormat(format unsafe.Pointer) []MetadataItem
	SamplePresentationTimeForTrackTime(trackTime unsafe.Pointer) unsafe.Pointer
	SegmentForTrackTime(trackTime unsafe.Pointer) unsafe.Pointer
}

// An object that models a track of media that an asset contains.
//
// An asset contains one or more tracks of media that the framework models using the class. A track object holds the uniformly typed media that an asset provides such as audio, video, or closed captions. A track, like its containing , doesn’t load all of its media upon creation. Instead, it defers loading its data until you perform an operation that requires it. Because loading the data can take time, an asset track adopts the protocol so you can load its property values asynchronously by calling the method.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AssetTrackClass) Alloc() AssetTrack {
	rv := objc.Send[AssetTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns an array of associated tracks that have the specified association type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/associatedTracks(ofType:)
func (a_ AssetTrack) AssociatedTracksOfType(trackAssociationType unsafe.Pointer) []AssetTrack {
	rv := objc.Send[[]AssetTrack](a_.ID, objc.Sel("associatedTracksOfType:"), trackAssociationType)
	return rv
}

// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/hasMediaCharacteristic(_:)
func (a_ AssetTrack) HasMediaCharacteristic(mediaCharacteristic unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Loads associated tracks that have the specified association type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadAssociatedTracks(ofType:completionHandler:)
func (a_ AssetTrack) LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadAssociatedTracksOfType:completionHandler:"), trackAssociationType, completionHandler)
}

// Loads metadata items that a track contains for the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadMetadata(for:completionHandler:)
func (a_ AssetTrack) LoadMetadataForFormatCompletionHandler(format unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}

// Loads a sample presentation time that maps to the specified track time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSamplePresentationTime(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSamplePresentationTimeForTrackTime:completionHandler:"), trackTime, completionHandler)
}

// Loads a segment with a target time range that contains, or is closest to, the specified track time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSegment(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSegmentForTrackTimeCompletionHandler(trackTime unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSegmentForTrackTime:completionHandler:"), trackTime, completionHandler)
}

// Creates a sample cursor and positions it at or near the specified presentation timestamp.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursor(presentationTimeStamp:)
func (a_ AssetTrack) MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("makeSampleCursorWithPresentationTimeStamp:"), presentationTimeStamp)
	return rv
}

// Creates a sample cursor and positions it at the track’s first media sample in decode order.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtFirstSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtFirstSampleInDecodeOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("makeSampleCursorAtFirstSampleInDecodeOrder"))
	return rv
}

// Creates a sample cursor and positions it at the track’s last media sample in decode order.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtLastSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtLastSampleInDecodeOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("makeSampleCursorAtLastSampleInDecodeOrder"))
	return rv
}

// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/metadata(forFormat:)
func (a_ AssetTrack) MetadataForFormat(format unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}

// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/samplePresentationTime(forTrackTime:)
func (a_ AssetTrack) SamplePresentationTimeForTrackTime(trackTime unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return rv
}

// Retrieves a segment with a target time range that contains, or is closest to, the specified track time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/segment(forTrackTime:)
func (a_ AssetTrack) SegmentForTrackTime(trackTime unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return rv
}

// The asset object that contains this track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/asset
func (a_ AssetTrack) Asset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("asset"))
	return rv
}

// An array of metadata formats available for the track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableMetadataFormats
func (a_ AssetTrack) AvailableMetadataFormats() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}

// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
func (a_ AssetTrack) AvailableTrackAssociationTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}

// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/canProvideSampleCursors
func (a_ AssetTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}

// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/commonMetadata
func (a_ AssetTrack) CommonMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}

// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/estimatedDataRate
func (a_ AssetTrack) EstimatedDataRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("estimatedDataRate"))
	return rv
}

// The language tag of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/extendedLanguageTag
func (a_ AssetTrack) ExtendedLanguageTag() string {
	rv := objc.Send[string](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}

// The format descriptions of the media samples that a track references.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/formatDescriptions
func (a_ AssetTrack) FormatDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("formatDescriptions"))
	return rv
}

// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/hasAudioSampleDependencies
func (a_ AssetTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}

// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isDecodable
func (a_ AssetTrack) Decodable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("decodable"))
	return rv
}

// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isEnabled
func (a_ AssetTrack) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}

// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isPlayable
func (a_ AssetTrack) Playable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playable"))
	return rv
}

// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isSelfContained
func (a_ AssetTrack) SelfContained() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selfContained"))
	return rv
}

// The language code of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/languageCode
func (a_ AssetTrack) LanguageCode() string {
	rv := objc.Send[string](a_.ID, objc.Sel("languageCode"))
	return rv
}

// The type of media that a track presents.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/mediaType
func (a_ AssetTrack) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaType"))
	return rv
}

// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/metadata
func (a_ AssetTrack) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}

// The minimum duration of the track’s frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/minFrameDuration
func (a_ AssetTrack) MinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("minFrameDuration"))
	return rv
}

// The natural dimensions of the media data that the track references.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalSize
func (a_ AssetTrack) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}

// The natural time scale of the media that a track references.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalTimeScale
func (a_ AssetTrack) NaturalTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("naturalTimeScale"))
	return rv
}

// The frame rate of the track, in frames per second.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/nominalFrameRate
func (a_ AssetTrack) NominalFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}

// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredTransform
func (a_ AssetTrack) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}

// The track’s volume preference for playing its audible media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredVolume
func (a_ AssetTrack) PreferredVolume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredVolume"))
	return rv
}

// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/requiresFrameReordering
func (a_ AssetTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}

// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/segments
func (a_ AssetTrack) Segments() []AssetTrackSegment {
	rv := objc.Send[[]AssetTrackSegment](a_.ID, objc.Sel("segments"))
	return rv
}

// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/timeRange
func (a_ AssetTrack) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("timeRange"))
	return rv
}

// The total number of bytes of sample data the track requires.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/totalSampleDataLength
func (a_ AssetTrack) TotalSampleDataLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}

// The persistent unique identifier for this track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/trackID
func (a_ AssetTrack) TrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackID"))
	return rv
}



