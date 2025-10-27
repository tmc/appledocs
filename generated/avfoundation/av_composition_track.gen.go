// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CompositionTrack] class.
var (
	CompositionTrackClass     _CompositionTrackClass
	CompositionTrackClassOnce sync.Once
)

func getCompositionTrackClass() _CompositionTrackClass {
	CompositionTrackClassOnce.Do(func() {
		CompositionTrackClass = _CompositionTrackClass{objc.GetClass("AVCompositionTrack")}
	})
	return CompositionTrackClass
}

type _CompositionTrackClass struct {
	class objc.Class
}





// An interface definition for the [CompositionTrack] class.
type ICompositionTrack interface {
	IAssetTrack
	

	// properties:
	AvailableMetadataFormats() MetadataFormat get /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */)
	AvailableTrackAssociationTypes() objectivec.IObject
	SetAvailableTrackAssociationTypes(value objectivec.IObject)
	CanProvideSampleCursors() objectivec.IObject
	SetCanProvideSampleCursors(value objectivec.IObject)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	EstimatedDataRate() objectivec.IObject
	SetEstimatedDataRate(value objectivec.IObject)
	ExtendedLanguageTag() objectivec.IObject
	SetExtendedLanguageTag(value objectivec.IObject)
	FormatDescriptionReplacements() []CompositionTrackFormatDescriptionReplacement
	FormatDescriptions() objectivec.IObject
	SetFormatDescriptions(value objectivec.IObject)
	HasAudioSampleDependencies() objectivec.IObject
	SetHasAudioSampleDependencies(value objectivec.IObject)
	IsDecodable() objectivec.IObject
	SetIsDecodable(value objectivec.IObject)
	IsEnabled() objectivec.IObject
	SetIsEnabled(value objectivec.IObject)
	IsPlayable() objectivec.IObject
	SetIsPlayable(value objectivec.IObject)
	IsSelfContained() objectivec.IObject
	SetIsSelfContained(value objectivec.IObject)
	LanguageCode() objectivec.IObject
	SetLanguageCode(value objectivec.IObject)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinFrameDuration() Time get /* not a class type */
	SetMinFrameDuration(value Time get /* not a class type */)
	NaturalSize() Size get /* not a class type */
	SetNaturalSize(value Size get /* not a class type */)
	NaturalTimeScale() TimeScale get /* not a class type */
	SetNaturalTimeScale(value TimeScale get /* not a class type */)
	NominalFrameRate() objectivec.IObject
	SetNominalFrameRate(value objectivec.IObject)
	PreferredTransform() AffineTransform get /* not a class type */
	SetPreferredTransform(value AffineTransform get /* not a class type */)
	PreferredVolume() objectivec.IObject
	SetPreferredVolume(value objectivec.IObject)
	RequiresFrameReordering() objectivec.IObject
	SetRequiresFrameReordering(value objectivec.IObject)
	Segments() []CompositionTrackSegment
	TimeRange() TimeRange get /* not a class type */
	SetTimeRange(value TimeRange get /* not a class type */)
	TotalSampleDataLength() objectivec.IObject
	SetTotalSampleDataLength(value objectivec.IObject)


	

	// methods:
	AssociatedTracksOfType(trackAssociationType TrackAssociationType) []AssetTrack
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool
	MetadataForFormat(format MetadataFormat) []MetadataItem
	SamplePresentationTimeForTrackTime(trackTime objectivec.IObject) objectivec.IObject
	SegmentForTrackTime(trackTime objectivec.IObject) ICompositionTrackSegment


}





// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackClass) Alloc() CompositionTrack {
	rv := objc.Send[CompositionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CompositionTrackClass) New() CompositionTrack {
	rv := objc.Send[CompositionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositionTrack) Init() CompositionTrack {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositionTrack) Autorelease() CompositionTrack {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositionTrack creates a new CompositionTrack instance.
func NewCompositionTrack() CompositionTrack {
	return getCompositionTrackClass().New()
}





// A track in a composition that presents media of a uniform type.
//
// This object provides an immutable composition track. The framework also provides a mutable subclass, .


// A track in a composition that presents media of a uniform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack
type CompositionTrack struct {
	AssetTrack
}

// CompositionTrackFrom constructs a [CompositionTrack] from an unsafe.Pointer.
//
// A track in a composition that presents media of a uniform type.
func CompositionTrackFrom(ptr unsafe.Pointer) CompositionTrack {
	return CompositionTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}




















// Returns an array of associated tracks that have the specified association type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/associatedTracks(ofType:)
func (c_ CompositionTrack) AssociatedTracksOfType(trackAssociationType TrackAssociationType) []AssetTrack {
	rv := objc.Send[[]AssetTrack](c_.ID, objc.Sel("associatedTracksOfType:"), trackAssociationType)
	return rv
}


// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasMediaCharacteristic(_:)
func (c_ CompositionTrack) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c_ CompositionTrack) MetadataForFormat(format MetadataFormat) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}


// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/samplePresentationTime(forTrackTime:)
func (c_ CompositionTrack) SamplePresentationTimeForTrackTime(trackTime objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return rv
}


// Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segment(forTrackTime:)
func (c_ CompositionTrack) SegmentForTrackTime(trackTime objectivec.IObject) ICompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](c_.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return rv
}







// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c_ CompositionTrack) AvailableMetadataFormats() MetadataFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c_ CompositionTrack) SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableTrackAssociationTypes
func (c_ CompositionTrack) AvailableTrackAssociationTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableTrackAssociationTypes
func (c_ CompositionTrack) SetAvailableTrackAssociationTypes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c_ CompositionTrack) CanProvideSampleCursors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c_ CompositionTrack) SetCanProvideSampleCursors(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c_ CompositionTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c_ CompositionTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c_ CompositionTrack) EstimatedDataRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c_ CompositionTrack) SetEstimatedDataRate(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedDataRate:"), value)
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
func (c_ CompositionTrack) ExtendedLanguageTag() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
func (c_ CompositionTrack) SetExtendedLanguageTag(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptionReplacements
func (c_ CompositionTrack) FormatDescriptionReplacements() []CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[[]CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
func (c_ CompositionTrack) FormatDescriptions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
func (c_ CompositionTrack) SetFormatDescriptions(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptions:"), value)
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c_ CompositionTrack) HasAudioSampleDependencies() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c_ CompositionTrack) SetHasAudioSampleDependencies(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c_ CompositionTrack) IsDecodable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isDecodable"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c_ CompositionTrack) SetIsDecodable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDecodable:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c_ CompositionTrack) IsEnabled() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c_ CompositionTrack) SetIsEnabled(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isPlayable
func (c_ CompositionTrack) IsPlayable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isPlayable
func (c_ CompositionTrack) SetIsPlayable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c_ CompositionTrack) IsSelfContained() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isSelfContained"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c_ CompositionTrack) SetIsSelfContained(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelfContained:"), value)
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/languageCode
func (c_ CompositionTrack) LanguageCode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/languageCode
func (c_ CompositionTrack) SetLanguageCode(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageCode:"), value)
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c_ CompositionTrack) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c_ CompositionTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
func (c_ CompositionTrack) MinFrameDuration() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
func (c_ CompositionTrack) SetMinFrameDuration(value Time get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
func (c_ CompositionTrack) NaturalSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
func (c_ CompositionTrack) SetNaturalSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c_ CompositionTrack) NaturalTimeScale() TimeScale get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c_ CompositionTrack) SetNaturalTimeScale(value TimeScale get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalTimeScale:"), value)
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/nominalFrameRate
func (c_ CompositionTrack) NominalFrameRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/nominalFrameRate
func (c_ CompositionTrack) SetNominalFrameRate(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNominalFrameRate:"), value)
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
func (c_ CompositionTrack) PreferredTransform() AffineTransform get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
func (c_ CompositionTrack) SetPreferredTransform(value AffineTransform get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c_ CompositionTrack) PreferredVolume() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c_ CompositionTrack) SetPreferredVolume(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c_ CompositionTrack) RequiresFrameReordering() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c_ CompositionTrack) SetRequiresFrameReordering(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segments
func (c_ CompositionTrack) Segments() []CompositionTrackSegment {
	rv := objc.Send[[]CompositionTrackSegment](c_.ID, objc.Sel("segments"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
func (c_ CompositionTrack) TimeRange() TimeRange get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
func (c_ CompositionTrack) SetTimeRange(value TimeRange get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeRange:"), value)
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c_ CompositionTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c_ CompositionTrack) SetTotalSampleDataLength(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}








