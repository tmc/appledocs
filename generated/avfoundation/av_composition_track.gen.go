// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	MetadataForFormat(format IMetadataFormat) []MetadataItem
	AvailableMetadataFormats() MetadataFormat
	SetAvailableMetadataFormats(value IMetadataFormat)
	AvailableTrackAssociationTypes() unsafe.Pointer
	SetAvailableTrackAssociationTypes(value unsafe.Pointer)
	CanProvideSampleCursors() bool
	SetCanProvideSampleCursors(value bool)
	CommonMetadata() AVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	EstimatedDataRate() float32
	SetEstimatedDataRate(value float32)
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	FormatDescriptionReplacements() unsafe.Pointer
	SetFormatDescriptionReplacements(value unsafe.Pointer)
	FormatDescriptions() unsafe.Pointer
	SetFormatDescriptions(value unsafe.Pointer)
	HasAudioSampleDependencies() bool
	SetHasAudioSampleDependencies(value bool)
	IsDecodable() bool
	SetIsDecodable(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsPlayable() bool
	SetIsPlayable(value bool)
	IsSelfContained() bool
	SetIsSelfContained(value bool)
	LanguageCode() string
	SetLanguageCode(value string)
	Metadata() AVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinFrameDuration() unsafe.Pointer
	SetMinFrameDuration(value unsafe.Pointer)
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	NaturalTimeScale() unsafe.Pointer
	SetNaturalTimeScale(value unsafe.Pointer)
	NominalFrameRate() float32
	SetNominalFrameRate(value float32)
	PreferredTransform() coregraphics.CGAffineTransform
	SetPreferredTransform(value coregraphics.CGAffineTransform)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)
	Segments() AVCompositionTrackSegment
	SetSegments(value IAVCompositionTrackSegment)
	TimeRange() unsafe.Pointer
	SetTimeRange(value unsafe.Pointer)
	TotalSampleDataLength() unsafe.Pointer
	SetTotalSampleDataLength(value unsafe.Pointer)
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

// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackClass) Alloc() CompositionTrack {
	rv := objc.Send[CompositionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c_ CompositionTrack) MetadataForFormat(format IMetadataFormat) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availablemetadataformats
func (c_ CompositionTrack) AvailableMetadataFormats() MetadataFormat {
	rv := objc.Send[MetadataFormat](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availablemetadataformats
func (c_ CompositionTrack) SetAvailableMetadataFormats(value IMetadataFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availabletrackassociationtypes
func (c_ CompositionTrack) AvailableTrackAssociationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availabletrackassociationtypes
func (c_ CompositionTrack) SetAvailableTrackAssociationTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/canprovidesamplecursors
func (c_ CompositionTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/canprovidesamplecursors
func (c_ CompositionTrack) SetCanProvideSampleCursors(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/commonmetadata
func (c_ CompositionTrack) CommonMetadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/commonmetadata
func (c_ CompositionTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/estimateddatarate
func (c_ CompositionTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/estimateddatarate
func (c_ CompositionTrack) SetEstimatedDataRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedDataRate:"), value)
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/extendedlanguagetag
func (c_ CompositionTrack) ExtendedLanguageTag() string {
	rv := objc.Send[string](c_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/extendedlanguagetag
func (c_ CompositionTrack) SetExtendedLanguageTag(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrack) FormatDescriptionReplacements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrack) SetFormatDescriptionReplacements(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptionReplacements:"), value)
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptions
func (c_ CompositionTrack) FormatDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptions
func (c_ CompositionTrack) SetFormatDescriptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptions:"), value)
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/hasaudiosampledependencies
func (c_ CompositionTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/hasaudiosampledependencies
func (c_ CompositionTrack) SetHasAudioSampleDependencies(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isdecodable
func (c_ CompositionTrack) IsDecodable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDecodable"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isdecodable
func (c_ CompositionTrack) SetIsDecodable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDecodable:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isenabled
func (c_ CompositionTrack) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isenabled
func (c_ CompositionTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isplayable
func (c_ CompositionTrack) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isplayable
func (c_ CompositionTrack) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isselfcontained
func (c_ CompositionTrack) IsSelfContained() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelfContained"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isselfcontained
func (c_ CompositionTrack) SetIsSelfContained(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelfContained:"), value)
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/languagecode
func (c_ CompositionTrack) LanguageCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/languagecode
func (c_ CompositionTrack) SetLanguageCode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/metadata
func (c_ CompositionTrack) Metadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/metadata
func (c_ CompositionTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/minframeduration
func (c_ CompositionTrack) MinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/minframeduration
func (c_ CompositionTrack) SetMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturalsize
func (c_ CompositionTrack) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural dimensions of the media data that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturalsize
func (c_ CompositionTrack) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturaltimescale
func (c_ CompositionTrack) NaturalTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturaltimescale
func (c_ CompositionTrack) SetNaturalTimeScale(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalTimeScale:"), value)
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/nominalframerate
func (c_ CompositionTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/nominalframerate
func (c_ CompositionTrack) SetNominalFrameRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNominalFrameRate:"), value)
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredtransform
func (c_ CompositionTrack) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredtransform
func (c_ CompositionTrack) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredvolume
func (c_ CompositionTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The track’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredvolume
func (c_ CompositionTrack) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/requiresframereordering
func (c_ CompositionTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/requiresframereordering
func (c_ CompositionTrack) SetRequiresFrameReordering(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/segments
func (c_ CompositionTrack) Segments() AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c_.ID, objc.Sel("segments"))
	return rv
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/segments
func (c_ CompositionTrack) SetSegments(value IAVCompositionTrackSegment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSegments:"), value)
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/timerange
func (c_ CompositionTrack) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/timerange
func (c_ CompositionTrack) SetTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeRange:"), value)
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/totalsampledatalength
func (c_ CompositionTrack) TotalSampleDataLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/totalsampledatalength
func (c_ CompositionTrack) SetTotalSampleDataLength(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}



