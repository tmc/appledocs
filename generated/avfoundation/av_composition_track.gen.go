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
	MetadataForFormat(format MetadataFormat) []MetadataItem
}

// A track in a composition that presents media of a uniform type.
//
// This object provides an immutable composition track. The framework also provides a mutable subclass, .
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c_ CompositionTrack) MetadataForFormat(format MetadataFormat) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}

// An array of metadata formats available for the track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availablemetadataformats
func (c_ CompositionTrack) AvailableMetadataFormats() MetadataFormat {
	rv := objc.Send[MetadataFormat](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// SetAvailableMetadataFormats sets the value of the availableMetadataFormats property.
// An array of metadata formats available for the track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availablemetadataformats
func (c_ CompositionTrack) SetAvailableMetadataFormats(value MetadataFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}

// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availabletrackassociationtypes
func (c_ CompositionTrack) AvailableTrackAssociationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// SetAvailableTrackAssociationTypes sets the value of the availableTrackAssociationTypes property.
// An array of association types that the track uses to associate with other tracks.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/availabletrackassociationtypes
func (c_ CompositionTrack) SetAvailableTrackAssociationTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}

// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/canprovidesamplecursors
func (c_ CompositionTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// SetCanProvideSampleCursors sets the value of the canProvideSampleCursors property.
// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/canprovidesamplecursors
func (c_ CompositionTrack) SetCanProvideSampleCursors(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}

// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/commonmetadata
func (c_ CompositionTrack) CommonMetadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}


// SetCommonMetadata sets the value of the commonMetadata property.
// An array of metadata items for all common metadata keys that have a value.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/commonmetadata
func (c_ CompositionTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}

// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/estimateddatarate
func (c_ CompositionTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// SetEstimatedDataRate sets the value of the estimatedDataRate property.
// The estimated data rate, in bits per second, of the media that the track references.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/estimateddatarate
func (c_ CompositionTrack) SetEstimatedDataRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedDataRate:"), value)
}

// The language tag of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/extendedlanguagetag
func (c_ CompositionTrack) ExtendedLanguageTag() string {
	rv := objc.Send[string](c_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// SetExtendedLanguageTag sets the value of the extendedLanguageTag property.
// The language tag of the track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/extendedlanguagetag
func (c_ CompositionTrack) SetExtendedLanguageTag(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}

// The replacement format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrack) FormatDescriptionReplacements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}


// SetFormatDescriptionReplacements sets the value of the formatDescriptionReplacements property.
// The replacement format descriptions.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrack) SetFormatDescriptionReplacements(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptionReplacements:"), value)
}

// The format descriptions of the media samples that a track references.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptions
func (c_ CompositionTrack) FormatDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// SetFormatDescriptions sets the value of the formatDescriptions property.
// The format descriptions of the media samples that a track references.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptions
func (c_ CompositionTrack) SetFormatDescriptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptions:"), value)
}

// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/hasaudiosampledependencies
func (c_ CompositionTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// SetHasAudioSampleDependencies sets the value of the hasAudioSampleDependencies property.
// A Boolean value that indicates whether the track has sample dependencies.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/hasaudiosampledependencies
func (c_ CompositionTrack) SetHasAudioSampleDependencies(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}

// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isdecodable
func (c_ CompositionTrack) IsDecodable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDecodable"))
	return rv
}


// SetIsDecodable sets the value of the isDecodable property.
// A Boolean value that indicates whether the track is decodable in the current environment.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isdecodable
func (c_ CompositionTrack) SetIsDecodable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDecodable:"), value)
}

// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isenabled
func (c_ CompositionTrack) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the track’s container enables it.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isenabled
func (c_ CompositionTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isplayable
func (c_ CompositionTrack) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// SetIsPlayable sets the value of the isPlayable property.
// A Boolean value that indicates whether the track is playable in the current environment.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isplayable
func (c_ CompositionTrack) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}

// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isselfcontained
func (c_ CompositionTrack) IsSelfContained() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelfContained"))
	return rv
}


// SetIsSelfContained sets the value of the isSelfContained property.
// A Boolean value that indicates whether this track references sample data only within its container file.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/isselfcontained
func (c_ CompositionTrack) SetIsSelfContained(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelfContained:"), value)
}

// The language code of the track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/languagecode
func (c_ CompositionTrack) LanguageCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("languageCode"))
	return rv
}


// SetLanguageCode sets the value of the languageCode property.
// The language code of the track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/languagecode
func (c_ CompositionTrack) SetLanguageCode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}

// An array of metadata items for all metadata identifiers that have a value.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/metadata
func (c_ CompositionTrack) Metadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// An array of metadata items for all metadata identifiers that have a value.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/metadata
func (c_ CompositionTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}

// The minimum duration of the track’s frames.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/minframeduration
func (c_ CompositionTrack) MinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// SetMinFrameDuration sets the value of the minFrameDuration property.
// The minimum duration of the track’s frames.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/minframeduration
func (c_ CompositionTrack) SetMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}

// The natural dimensions of the media data that the track references.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturalsize
func (c_ CompositionTrack) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("naturalSize"))
	return rv
}


// SetNaturalSize sets the value of the naturalSize property.
// The natural dimensions of the media data that the track references.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturalsize
func (c_ CompositionTrack) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}

// The natural time scale of the media that a track references.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturaltimescale
func (c_ CompositionTrack) NaturalTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// SetNaturalTimeScale sets the value of the naturalTimeScale property.
// The natural time scale of the media that a track references.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/naturaltimescale
func (c_ CompositionTrack) SetNaturalTimeScale(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalTimeScale:"), value)
}

// The frame rate of the track, in frames per second.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/nominalframerate
func (c_ CompositionTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// SetNominalFrameRate sets the value of the nominalFrameRate property.
// The frame rate of the track, in frames per second.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/nominalframerate
func (c_ CompositionTrack) SetNominalFrameRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNominalFrameRate:"), value)
}

// The track’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredtransform
func (c_ CompositionTrack) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}


// SetPreferredTransform sets the value of the preferredTransform property.
// The track’s transform preference to apply to its visual content during presentation or processing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredtransform
func (c_ CompositionTrack) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}

// The track’s volume preference for playing its audible media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredvolume
func (c_ CompositionTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("preferredVolume"))
	return rv
}


// SetPreferredVolume sets the value of the preferredVolume property.
// The track’s volume preference for playing its audible media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/preferredvolume
func (c_ CompositionTrack) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}

// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/requiresframereordering
func (c_ CompositionTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// SetRequiresFrameReordering sets the value of the requiresFrameReordering property.
// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/requiresframereordering
func (c_ CompositionTrack) SetRequiresFrameReordering(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}

// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/segments
func (c_ CompositionTrack) Segments() AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c_.ID, objc.Sel("segments"))
	return rv
}


// SetSegments sets the value of the segments property.
// The time mappings from the track’s media samples to its timeline.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/segments
func (c_ CompositionTrack) SetSegments(value IAVCompositionTrackSegment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSegments:"), value)
}

// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/timerange
func (c_ CompositionTrack) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeRange"))
	return rv
}


// SetTimeRange sets the value of the timeRange property.
// The time range of the track within the overall timeline of the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/timerange
func (c_ CompositionTrack) SetTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeRange:"), value)
}

// The total number of bytes of sample data the track requires.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/totalsampledatalength
func (c_ CompositionTrack) TotalSampleDataLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// SetTotalSampleDataLength sets the value of the totalSampleDataLength property.
// The total number of bytes of sample data the track requires.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/totalsampledatalength
func (c_ CompositionTrack) SetTotalSampleDataLength(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}



