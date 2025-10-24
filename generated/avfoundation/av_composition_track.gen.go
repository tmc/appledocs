// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
)

/* debug [class.gen.go]: Generating class AVCompositionTrack */


/* debug [class_header]: Header for AVCompositionTrack */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompositionTrack */
// An interface definition for the [CompositionTrack] class.
type ICompositionTrack interface {
	IAssetTrack
	
/* debug [class_interface_properties]: Properties for CompositionTrack */
	// properties:
	AvailableMetadataFormats() MetadataFormat get /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */)
	AvailableTrackAssociationTypes() objc.IObject
	SetAvailableTrackAssociationTypes(value objc.IObject)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompositionTrack */
	// methods:
	AssociatedTracksOfType(trackAssociationType TrackAssociationType /* typedef */) []AssetTrack
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool
	MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem
	SamplePresentationTimeForTrackTime(trackTime objc.IObject /* cross-framework: Time */) objc.IObject /* cross-framework: Time */
	SegmentForTrackTime(trackTime objc.IObject /* cross-framework: Time */) ICompositionTrackSegment
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompositionTrack */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompositionTrack */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompositionTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompositionTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompositionTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompositionTrack */

// Returns an array of associated tracks that have the specified association type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/associatedTracks(ofType:)
func (c_ CompositionTrack) AssociatedTracksOfType(trackAssociationType TrackAssociationType /* typedef */) []AssetTrack {
	rv := objc.Send[[]AssetTrack](c_.ID, objc.Sel("associatedTracksOfType:"), trackAssociationType)
	return rv
}/* debug [instance_methods/method]: AssociatedTracksOfType */


// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasMediaCharacteristic(_:)
func (c_ CompositionTrack) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}/* debug [instance_methods/method]: HasMediaCharacteristic */


// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c_ CompositionTrack) MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}/* debug [instance_methods/method]: MetadataForFormat */


// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/samplePresentationTime(forTrackTime:)
func (c_ CompositionTrack) SamplePresentationTimeForTrackTime(trackTime objc.IObject /* cross-framework: Time */) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return rv
}/* debug [instance_methods/method]: SamplePresentationTimeForTrackTime */


// Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segment(forTrackTime:)
func (c_ CompositionTrack) SegmentForTrackTime(trackTime objc.IObject /* cross-framework: Time */) ICompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](c_.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return rv
}/* debug [instance_methods/method]: SegmentForTrackTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompositionTrack */

// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c_ CompositionTrack) AvailableMetadataFormats() MetadataFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}/* debug [instance_properties/getter]: availableMetadataFormats */


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c_ CompositionTrack) SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}/* debug [instance_properties/setter]: availableMetadataFormats */


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableTrackAssociationTypes
func (c_ CompositionTrack) AvailableTrackAssociationTypes() objc.IObject {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}/* debug [instance_properties/getter]: availableTrackAssociationTypes */


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableTrackAssociationTypes
func (c_ CompositionTrack) SetAvailableTrackAssociationTypes(value objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}/* debug [instance_properties/setter]: availableTrackAssociationTypes */


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c_ CompositionTrack) CanProvideSampleCursors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}/* debug [instance_properties/getter]: canProvideSampleCursors */


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c_ CompositionTrack) SetCanProvideSampleCursors(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}/* debug [instance_properties/setter]: canProvideSampleCursors */


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c_ CompositionTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}/* debug [instance_properties/getter]: commonMetadata */


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c_ CompositionTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}/* debug [instance_properties/setter]: commonMetadata */


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c_ CompositionTrack) EstimatedDataRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("estimatedDataRate"))
	return rv
}/* debug [instance_properties/getter]: estimatedDataRate */


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c_ CompositionTrack) SetEstimatedDataRate(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedDataRate:"), value)
}/* debug [instance_properties/setter]: estimatedDataRate */


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
func (c_ CompositionTrack) ExtendedLanguageTag() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}/* debug [instance_properties/getter]: extendedLanguageTag */


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
func (c_ CompositionTrack) SetExtendedLanguageTag(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}/* debug [instance_properties/setter]: extendedLanguageTag */


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptionReplacements
func (c_ CompositionTrack) FormatDescriptionReplacements() []CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[[]CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}/* debug [instance_properties/getter]: formatDescriptionReplacements */


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
func (c_ CompositionTrack) FormatDescriptions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("formatDescriptions"))
	return rv
}/* debug [instance_properties/getter]: formatDescriptions */


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
func (c_ CompositionTrack) SetFormatDescriptions(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptions:"), value)
}/* debug [instance_properties/setter]: formatDescriptions */


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c_ CompositionTrack) HasAudioSampleDependencies() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}/* debug [instance_properties/getter]: hasAudioSampleDependencies */


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c_ CompositionTrack) SetHasAudioSampleDependencies(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}/* debug [instance_properties/setter]: hasAudioSampleDependencies */


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c_ CompositionTrack) IsDecodable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isDecodable"))
	return rv
}/* debug [instance_properties/getter]: isDecodable */


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c_ CompositionTrack) SetIsDecodable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDecodable:"), value)
}/* debug [instance_properties/setter]: isDecodable */


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c_ CompositionTrack) IsEnabled() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c_ CompositionTrack) SetIsEnabled(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isPlayable
func (c_ CompositionTrack) IsPlayable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isPlayable"))
	return rv
}/* debug [instance_properties/getter]: isPlayable */


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isPlayable
func (c_ CompositionTrack) SetIsPlayable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}/* debug [instance_properties/setter]: isPlayable */


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c_ CompositionTrack) IsSelfContained() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isSelfContained"))
	return rv
}/* debug [instance_properties/getter]: isSelfContained */


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c_ CompositionTrack) SetIsSelfContained(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelfContained:"), value)
}/* debug [instance_properties/setter]: isSelfContained */


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/languageCode
func (c_ CompositionTrack) LanguageCode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("languageCode"))
	return rv
}/* debug [instance_properties/getter]: languageCode */


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/languageCode
func (c_ CompositionTrack) SetLanguageCode(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageCode:"), value)
}/* debug [instance_properties/setter]: languageCode */


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c_ CompositionTrack) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c_ CompositionTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}/* debug [instance_properties/setter]: metadata */


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
func (c_ CompositionTrack) MinFrameDuration() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: minFrameDuration */


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
func (c_ CompositionTrack) SetMinFrameDuration(value Time get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}/* debug [instance_properties/setter]: minFrameDuration */


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
func (c_ CompositionTrack) NaturalSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("naturalSize"))
	return rv
}/* debug [instance_properties/getter]: naturalSize */


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
func (c_ CompositionTrack) SetNaturalSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}/* debug [instance_properties/setter]: naturalSize */


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c_ CompositionTrack) NaturalTimeScale() TimeScale get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("naturalTimeScale"))
	return rv
}/* debug [instance_properties/getter]: naturalTimeScale */


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c_ CompositionTrack) SetNaturalTimeScale(value TimeScale get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalTimeScale:"), value)
}/* debug [instance_properties/setter]: naturalTimeScale */


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/nominalFrameRate
func (c_ CompositionTrack) NominalFrameRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("nominalFrameRate"))
	return rv
}/* debug [instance_properties/getter]: nominalFrameRate */


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/nominalFrameRate
func (c_ CompositionTrack) SetNominalFrameRate(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNominalFrameRate:"), value)
}/* debug [instance_properties/setter]: nominalFrameRate */


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
func (c_ CompositionTrack) PreferredTransform() AffineTransform get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
func (c_ CompositionTrack) SetPreferredTransform(value AffineTransform get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}/* debug [instance_properties/setter]: preferredTransform */


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c_ CompositionTrack) PreferredVolume() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("preferredVolume"))
	return rv
}/* debug [instance_properties/getter]: preferredVolume */


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c_ CompositionTrack) SetPreferredVolume(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}/* debug [instance_properties/setter]: preferredVolume */


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c_ CompositionTrack) RequiresFrameReordering() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}/* debug [instance_properties/getter]: requiresFrameReordering */


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c_ CompositionTrack) SetRequiresFrameReordering(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}/* debug [instance_properties/setter]: requiresFrameReordering */


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segments
func (c_ CompositionTrack) Segments() []CompositionTrackSegment {
	rv := objc.Send[[]CompositionTrackSegment](c_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
func (c_ CompositionTrack) TimeRange() TimeRange get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
func (c_ CompositionTrack) SetTimeRange(value TimeRange get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeRange:"), value)
}/* debug [instance_properties/setter]: timeRange */


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c_ CompositionTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}/* debug [instance_properties/getter]: totalSampleDataLength */


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c_ CompositionTrack) SetTotalSampleDataLength(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}/* debug [instance_properties/setter]: totalSampleDataLength */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCompositionTrack */



