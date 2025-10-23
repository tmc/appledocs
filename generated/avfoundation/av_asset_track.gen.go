// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AvailableTrackAssociationTypes() []string /* primitive/slice/pointer */
	Segments() []AssetTrackSegment /* primitive/slice/pointer */
	TimeRange() TimeRange /* not a class type */
	Asset() IAVAsset
	SetAsset(value IAVAsset)
	AvailableMetadataFormats() MetadataFormat /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat /* not a class type */)
	CanProvideSampleCursors() bool /* primitive/slice/pointer */
	SetCanProvideSampleCursors(value bool /* primitive/slice/pointer */)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	EstimatedDataRate() float32 /* primitive/slice/pointer */
	SetEstimatedDataRate(value float32 /* primitive/slice/pointer */)
	ExtendedLanguageTag() string /* primitive/slice/pointer */
	SetExtendedLanguageTag(value string /* primitive/slice/pointer */)
	FormatDescriptions() unsafe.Pointer
	SetFormatDescriptions(value unsafe.Pointer)
	HasAudioSampleDependencies() bool /* primitive/slice/pointer */
	SetHasAudioSampleDependencies(value bool /* primitive/slice/pointer */)
	IsDecodable() bool /* primitive/slice/pointer */
	SetIsDecodable(value bool /* primitive/slice/pointer */)
	IsEnabled() bool /* primitive/slice/pointer */
	SetIsEnabled(value bool /* primitive/slice/pointer */)
	IsPlayable() bool /* primitive/slice/pointer */
	SetIsPlayable(value bool /* primitive/slice/pointer */)
	IsSelfContained() bool /* primitive/slice/pointer */
	SetIsSelfContained(value bool /* primitive/slice/pointer */)
	LanguageCode() string /* primitive/slice/pointer */
	SetLanguageCode(value string /* primitive/slice/pointer */)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinFrameDuration() Time /* not a class type */
	SetMinFrameDuration(value Time /* not a class type */)
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	NaturalTimeScale() TimeScale /* not a class type */
	SetNaturalTimeScale(value TimeScale /* not a class type */)
	NominalFrameRate() float32 /* primitive/slice/pointer */
	SetNominalFrameRate(value float32 /* primitive/slice/pointer */)
	PreferredTransform() coregraphics.CGAffineTransform
	SetPreferredTransform(value coregraphics.CGAffineTransform)
	PreferredVolume() float32 /* primitive/slice/pointer */
	SetPreferredVolume(value float32 /* primitive/slice/pointer */)
	RequiresFrameReordering() bool /* primitive/slice/pointer */
	SetRequiresFrameReordering(value bool /* primitive/slice/pointer */)
	TotalSampleDataLength() unsafe.Pointer
	SetTotalSampleDataLength(value unsafe.Pointer)
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)
	// methods:
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



// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
func (a_ AssetTrack) AvailableTrackAssociationTypes() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/segments
func (a_ AssetTrack) Segments() []AssetTrackSegment /* primitive/slice/pointer */ {
	rv := objc.Send[[]AssetTrackSegment](a_.ID, objc.Sel("segments"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/timeRange
func (a_ AssetTrack) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The asset object that contains this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/asset
func (a_ AssetTrack) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// The asset object that contains this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/asset
func (a_ AssetTrack) SetAsset(value IAVAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAsset:"), value)
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/availablemetadataformats
func (a_ AssetTrack) AvailableMetadataFormats() MetadataFormat /* not a class type */ {
	rv := objc.Send[MetadataFormat](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/availablemetadataformats
func (a_ AssetTrack) SetAvailableMetadataFormats(value MetadataFormat /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/canprovidesamplecursors
func (a_ AssetTrack) CanProvideSampleCursors() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/canprovidesamplecursors
func (a_ AssetTrack) SetCanProvideSampleCursors(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/commonmetadata
func (a_ AssetTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/commonmetadata
func (a_ AssetTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCommonMetadata:"), value)
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/estimateddatarate
func (a_ AssetTrack) EstimatedDataRate() float32 /* primitive/slice/pointer */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/estimateddatarate
func (a_ AssetTrack) SetEstimatedDataRate(value float32 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEstimatedDataRate:"), value)
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/extendedlanguagetag
func (a_ AssetTrack) ExtendedLanguageTag() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/extendedlanguagetag
func (a_ AssetTrack) SetExtendedLanguageTag(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/formatdescriptions
func (a_ AssetTrack) FormatDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/formatdescriptions
func (a_ AssetTrack) SetFormatDescriptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormatDescriptions:"), value)
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/hasaudiosampledependencies
func (a_ AssetTrack) HasAudioSampleDependencies() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/hasaudiosampledependencies
func (a_ AssetTrack) SetHasAudioSampleDependencies(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) IsDecodable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDecodable"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isdecodable
func (a_ AssetTrack) SetIsDecodable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDecodable:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) IsEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isenabled
func (a_ AssetTrack) SetIsEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) IsPlayable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isplayable
func (a_ AssetTrack) SetIsPlayable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) IsSelfContained() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSelfContained"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/isselfcontained
func (a_ AssetTrack) SetIsSelfContained(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSelfContained:"), value)
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/languagecode
func (a_ AssetTrack) LanguageCode() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/languagecode
func (a_ AssetTrack) SetLanguageCode(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}


// The type of media that a track presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/mediatype
func (a_ AssetTrack) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](a_.ID, objc.Sel("mediaType"))
	return rv
}


// The type of media that a track presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/mediatype
func (a_ AssetTrack) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaType:"), value)
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/metadata
func (a_ AssetTrack) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/metadata
func (a_ AssetTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/minframeduration
func (a_ AssetTrack) MinFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/minframeduration
func (a_ AssetTrack) SetMinFrameDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/naturalsize
func (a_ AssetTrack) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/naturalsize
func (a_ AssetTrack) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/naturaltimescale
func (a_ AssetTrack) NaturalTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/naturaltimescale
func (a_ AssetTrack) SetNaturalTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalTimeScale:"), value)
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/nominalframerate
func (a_ AssetTrack) NominalFrameRate() float32 /* primitive/slice/pointer */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/nominalframerate
func (a_ AssetTrack) SetNominalFrameRate(value float32 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNominalFrameRate:"), value)
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/preferredtransform
func (a_ AssetTrack) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/preferredtransform
func (a_ AssetTrack) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/preferredvolume
func (a_ AssetTrack) PreferredVolume() float32 /* primitive/slice/pointer */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/preferredvolume
func (a_ AssetTrack) SetPreferredVolume(value float32 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/requiresframereordering
func (a_ AssetTrack) RequiresFrameReordering() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/requiresframereordering
func (a_ AssetTrack) SetRequiresFrameReordering(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/totalsampledatalength
func (a_ AssetTrack) TotalSampleDataLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/totalsampledatalength
func (a_ AssetTrack) SetTotalSampleDataLength(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}


// The persistent unique identifier for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/trackid
func (a_ AssetTrack) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}


// The persistent unique identifier for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/trackid
func (a_ AssetTrack) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackID:"), value)
}



