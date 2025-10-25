// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetTrack */


/* debug [class_header]: Header for AVAssetTrack */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetTrack */
// An interface definition for the [AssetTrack] class.
type IAssetTrack interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetTrack */
	// properties:
	Asset() IAVAsset
	AvailableMetadataFormats() []string
	AvailableTrackAssociationTypes() []string
	CanProvideSampleCursors() bool
	CommonMetadata() []MetadataItem
	EstimatedDataRate() float32
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	FormatDescriptions() objc.IObject /* cross-framework: NSArray */
	HasAudioSampleDependencies() bool
	Decodable() bool
	Enabled() bool
	Playable() bool
	SelfContained() bool
	LanguageCode() objc.IObject /* cross-framework: NSString */
	MediaType() MediaType /* typedef */
	Metadata() []MetadataItem
	MinFrameDuration() objc.IObject /* cross-framework: Time */
	NaturalSize() corefoundation.CGSize
	NaturalTimeScale() TimeScale /* not a class type */
	NominalFrameRate() float32
	PreferredTransform() corefoundation.CGAffineTransform
	PreferredVolume() float32
	RequiresFrameReordering() bool
	Segments() []AssetTrackSegment
	TimeRange() TimeRange /* not a class type */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetTrack */
	// methods:
	LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType /* typedef */, completionHandler unsafe.Pointer)
	LoadMetadataForFormatCompletionHandler(format MetadataFormat /* typedef */, completionHandler unsafe.Pointer)
	LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	LoadSegmentForTrackTimeCompletionHandler(trackTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp objc.IObject /* cross-framework: Time */) ISampleCursor
	MakeSampleCursorAtFirstSampleInDecodeOrder() ISampleCursor
	MakeSampleCursorAtLastSampleInDecodeOrder() ISampleCursor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetTrack */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetTrack */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetTrack */

// Loads associated tracks that have the specified association type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadAssociatedTracks(ofType:completionHandler:)
func (a_ AssetTrack) LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadAssociatedTracksOfType:completionHandler:"), trackAssociationType, completionHandler)
}/* debug [instance_methods/method]: LoadAssociatedTracksOfTypeCompletionHandler */


// Loads metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadMetadata(for:completionHandler:)
func (a_ AssetTrack) LoadMetadataForFormatCompletionHandler(format MetadataFormat /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}/* debug [instance_methods/method]: LoadMetadataForFormatCompletionHandler */


// Loads a sample presentation time that maps to the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSamplePresentationTime(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSamplePresentationTimeForTrackTime:completionHandler:"), trackTime, completionHandler)
}/* debug [instance_methods/method]: LoadSamplePresentationTimeForTrackTimeCompletionHandler */


// Loads a segment with a target time range that contains, or is closest to, the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSegment(forTrackTime:completionHandler:)
func (a_ AssetTrack) LoadSegmentForTrackTimeCompletionHandler(trackTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadSegmentForTrackTime:completionHandler:"), trackTime, completionHandler)
}/* debug [instance_methods/method]: LoadSegmentForTrackTimeCompletionHandler */


// Creates a sample cursor and positions it at or near the specified presentation timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursor(presentationTimeStamp:)
func (a_ AssetTrack) MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp objc.IObject /* cross-framework: Time */) ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorWithPresentationTimeStamp:"), presentationTimeStamp)
	return rv
}/* debug [instance_methods/method]: MakeSampleCursorWithPresentationTimeStamp */


// Creates a sample cursor and positions it at the track’s first media sample in decode order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtFirstSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtFirstSampleInDecodeOrder() ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorAtFirstSampleInDecodeOrder"))
	return rv
}/* debug [instance_methods/method]: MakeSampleCursorAtFirstSampleInDecodeOrder */


// Creates a sample cursor and positions it at the track’s last media sample in decode order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtLastSampleInDecodeOrder()
func (a_ AssetTrack) MakeSampleCursorAtLastSampleInDecodeOrder() ISampleCursor {
	rv := objc.Send[SampleCursor](a_.ID, objc.Sel("makeSampleCursorAtLastSampleInDecodeOrder"))
	return rv
}/* debug [instance_methods/method]: MakeSampleCursorAtLastSampleInDecodeOrder */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetTrack */

// The asset object that contains this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/asset
func (a_ AssetTrack) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableMetadataFormats
func (a_ AssetTrack) AvailableMetadataFormats() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}/* debug [instance_properties/getter]: availableMetadataFormats */


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
func (a_ AssetTrack) AvailableTrackAssociationTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}/* debug [instance_properties/getter]: availableTrackAssociationTypes */


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/canProvideSampleCursors
func (a_ AssetTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}/* debug [instance_properties/getter]: canProvideSampleCursors */


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/commonMetadata
func (a_ AssetTrack) CommonMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}/* debug [instance_properties/getter]: commonMetadata */


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/estimatedDataRate
func (a_ AssetTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("estimatedDataRate"))
	return rv
}/* debug [instance_properties/getter]: estimatedDataRate */


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/extendedLanguageTag
func (a_ AssetTrack) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}/* debug [instance_properties/getter]: extendedLanguageTag */


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/formatDescriptions
func (a_ AssetTrack) FormatDescriptions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("formatDescriptions"))
	return rv
}/* debug [instance_properties/getter]: formatDescriptions */


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/hasAudioSampleDependencies
func (a_ AssetTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}/* debug [instance_properties/getter]: hasAudioSampleDependencies */


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isDecodable
func (a_ AssetTrack) Decodable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("decodable"))
	return rv
}/* debug [instance_properties/getter]: decodable */


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isEnabled
func (a_ AssetTrack) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isPlayable
func (a_ AssetTrack) Playable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playable"))
	return rv
}/* debug [instance_properties/getter]: playable */


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/isSelfContained
func (a_ AssetTrack) SelfContained() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selfContained"))
	return rv
}/* debug [instance_properties/getter]: selfContained */


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/languageCode
func (a_ AssetTrack) LanguageCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("languageCode"))
	return rv
}/* debug [instance_properties/getter]: languageCode */


// The type of media that a track presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/mediaType
func (a_ AssetTrack) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/metadata
func (a_ AssetTrack) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/minFrameDuration
func (a_ AssetTrack) MinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("minFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: minFrameDuration */


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalSize
func (a_ AssetTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}/* debug [instance_properties/getter]: naturalSize */


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/naturalTimeScale
func (a_ AssetTrack) NaturalTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("naturalTimeScale"))
	return rv
}/* debug [instance_properties/getter]: naturalTimeScale */


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/nominalFrameRate
func (a_ AssetTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}/* debug [instance_properties/getter]: nominalFrameRate */


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredTransform
func (a_ AssetTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredVolume
func (a_ AssetTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}/* debug [instance_properties/getter]: preferredVolume */


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/requiresFrameReordering
func (a_ AssetTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}/* debug [instance_properties/getter]: requiresFrameReordering */


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/segments
func (a_ AssetTrack) Segments() []AssetTrackSegment {
	rv := objc.Send[[]AssetTrackSegment](a_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/timeRange
func (a_ AssetTrack) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/totalSampleDataLength
func (a_ AssetTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}/* debug [instance_properties/getter]: totalSampleDataLength */


// The persistent unique identifier for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/trackID
func (a_ AssetTrack) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) IsDecodable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDecodable"))
	return rv
}/* debug [instance_properties/getter]: isDecodable */


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) SetIsDecodable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDecodable:"), value)
}/* debug [instance_properties/setter]: isDecodable */


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) IsPlayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}/* debug [instance_properties/getter]: isPlayable */


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) SetIsPlayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}/* debug [instance_properties/setter]: isPlayable */


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) IsSelfContained() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSelfContained"))
	return rv
}/* debug [instance_properties/getter]: isSelfContained */


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) SetIsSelfContained(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSelfContained:"), value)
}/* debug [instance_properties/setter]: isSelfContained */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetTrack */



