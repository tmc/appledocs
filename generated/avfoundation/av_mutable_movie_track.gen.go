// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableMovieTrack] class.
var (
	MutableMovieTrackClass     _MutableMovieTrackClass
	MutableMovieTrackClassOnce sync.Once
)

func getMutableMovieTrackClass() _MutableMovieTrackClass {
	MutableMovieTrackClassOnce.Do(func() {
		MutableMovieTrackClass = _MutableMovieTrackClass{objc.GetClass("AVMutableMovieTrack")}
	})
	return MutableMovieTrackClass
}

type _MutableMovieTrackClass struct {
	class objc.Class
}





// An interface definition for the [MutableMovieTrack] class.
type IMutableMovieTrack interface {
	IMovieTrack
	

	// properties:
	AlternateGroupID() int
	SetAlternateGroupID(value int)
	AvailableMetadataFormats() MetadataFormat get /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */)
	AvailableTrackAssociationTypes() objc.IObject
	SetAvailableTrackAssociationTypes(value objc.IObject)
	CanProvideSampleCursors() objectivec.IObject
	SetCanProvideSampleCursors(value objectivec.IObject)
	CleanApertureDimensions() corefoundation.CGSize
	SetCleanApertureDimensions(value corefoundation.CGSize)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	EncodedPixelsDimensions() corefoundation.CGSize
	SetEncodedPixelsDimensions(value corefoundation.CGSize)
	EstimatedDataRate() objectivec.IObject
	SetEstimatedDataRate(value objectivec.IObject)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	FormatDescriptions() objectivec.IObject
	SetFormatDescriptions(value objectivec.IObject)
	HasAudioSampleDependencies() objectivec.IObject
	SetHasAudioSampleDependencies(value objectivec.IObject)
	HasProtectedContent() bool
	IsDecodable() objectivec.IObject
	SetIsDecodable(value objectivec.IObject)
	Enabled() bool
	SetEnabled(value bool)
	Modified() bool
	SetModified(value bool)
	IsPlayable() objectivec.IObject
	SetIsPlayable(value objectivec.IObject)
	IsSelfContained() objectivec.IObject
	SetIsSelfContained(value objectivec.IObject)
	LanguageCode() objc.IObject /* cross-framework: NSString */
	SetLanguageCode(value objc.IObject /* cross-framework: NSString */)
	Layer() int
	SetLayer(value int)
	MediaDataStorage() IAVMediaDataStorage
	SetMediaDataStorage(value IAVMediaDataStorage)
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MinFrameDuration() Time get /* not a class type */
	SetMinFrameDuration(value Time get /* not a class type */)
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	NaturalTimeScale() TimeScale get /* not a class type */
	SetNaturalTimeScale(value TimeScale get /* not a class type */)
	NominalFrameRate() objectivec.IObject
	SetNominalFrameRate(value objectivec.IObject)
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	PreferredMediaChunkDuration() objc.IObject /* cross-framework: Time */
	SetPreferredMediaChunkDuration(value objc.IObject /* cross-framework: Time */)
	PreferredMediaChunkSize() int
	SetPreferredMediaChunkSize(value int)
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	ProductionApertureDimensions() corefoundation.CGSize
	SetProductionApertureDimensions(value corefoundation.CGSize)
	RequiresFrameReordering() objectivec.IObject
	SetRequiresFrameReordering(value objectivec.IObject)
	SampleReferenceBaseURL() objc.IObject /* cross-framework: NSURL */
	SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: NSURL */)
	Segments() IAVAssetTrackSegment
	SetSegments(value IAVAssetTrackSegment)
	TimeRange() TimeRange get /* not a class type */
	SetTimeRange(value TimeRange get /* not a class type */)
	Timescale() TimeScale /* not a class type */
	SetTimescale(value TimeScale /* not a class type */)
	TotalSampleDataLength() objectivec.IObject
	SetTotalSampleDataLength(value objectivec.IObject)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsModified() bool
	SetIsModified(value bool)


	

	// methods:
	AddTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType TrackAssociationType /* typedef */)
	AssociatedTracksOfType(trackAssociationType TrackAssociationType /* typedef */) []AssetTrack
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool
	InsertEmptyTimeRange(timeRange TimeRange /* not a class type */)
	InsertMediaTimeRangeIntoTimeRange(mediaTimeRange TimeRange /* not a class type */, trackTimeRange TimeRange /* not a class type */) bool
	InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange TimeRange /* not a class type */, track IAVAssetTrack, startTime objc.IObject /* cross-framework: Time */, copySampleData bool, outError objectivec.IObject) bool
	MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem
	RemoveTimeRange(timeRange TimeRange /* not a class type */)
	RemoveTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType TrackAssociationType /* typedef */)
	ReplaceFormatDescriptionWithFormatDescription(formatDescription FormatDescriptionRef /* not a class type */, newFormatDescription FormatDescriptionRef /* not a class type */)
	SamplePresentationTimeForTrackTime(trackTime objc.IObject /* cross-framework: Time */) objc.IObject /* cross-framework: Time */
	ScaleTimeRangeToDuration(timeRange TimeRange /* not a class type */, duration objc.IObject /* cross-framework: Time */)
	SegmentForTrackTime(trackTime objc.IObject /* cross-framework: Time */) IAssetTrackSegment


}





// Alloc allocates a new instance without initialization.
func (mc _MutableMovieTrackClass) Alloc() MutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableMovieTrackClass) New() MutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableMovieTrack) Init() MutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableMovieTrack) Autorelease() MutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableMovieTrack creates a new MutableMovieTrack instance.
func NewMutableMovieTrack() MutableMovieTrack {
	return getMutableMovieTrackClass().New()
}





// A mutable track that conforms to the QuickTime or ISO base media file format.


// A mutable track that conforms to the QuickTime or ISO base media file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack
type MutableMovieTrack struct {
	MovieTrack
}

// MutableMovieTrackFrom constructs a [MutableMovieTrack] from an unsafe.Pointer.
//
// A mutable track that conforms to the QuickTime or ISO base media file format.
func MutableMovieTrackFrom(ptr unsafe.Pointer) MutableMovieTrack {
	return MutableMovieTrack{
		MovieTrack: MovieTrackFrom(ptr),
	}
}




















// Creates a specific type of track association between two tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/addTrackAssociation(to:type:)
func (m_ MutableMovieTrack) AddTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType TrackAssociationType /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addTrackAssociationToTrack:type:"), movieTrack, trackAssociationType)
}


// Returns an array of associated tracks that have the specified association type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/associatedTracks(ofType:)
func (m_ MutableMovieTrack) AssociatedTracksOfType(trackAssociationType TrackAssociationType /* typedef */) []AssetTrack {
	rv := objc.Send[[]AssetTrack](m_.ID, objc.Sel("associatedTracksOfType:"), trackAssociationType)
	return rv
}


// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasMediaCharacteristic(_:)
func (m_ MutableMovieTrack) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Adds an empty time range to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertEmptyTimeRange(_:)
func (m_ MutableMovieTrack) InsertEmptyTimeRange(timeRange TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}


// Inserts a reference to a media time range into a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertMediaTimeRange(_:into:)
func (m_ MutableMovieTrack) InsertMediaTimeRangeIntoTimeRange(mediaTimeRange TimeRange /* not a class type */, trackTimeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertMediaTimeRange:intoTimeRange:"), mediaTimeRange, trackTimeRange)
	return rv
}


// Inserts a portion of an asset track into the target movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertTimeRange(_:of:at:copySampleData:)
func (m_ MutableMovieTrack) InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange TimeRange /* not a class type */, track IAVAssetTrack, startTime objc.IObject /* cross-framework: Time */, copySampleData bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertTimeRange:ofTrack:atTime:copySampleData:error:"), timeRange, track, startTime, copySampleData, outError)
	return rv
}


// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/metadata(forFormat:)
func (m_ MutableMovieTrack) MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}


// Removes the specified time range from a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/removeTimeRange(_:)
func (m_ MutableMovieTrack) RemoveTimeRange(timeRange TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTimeRange:"), timeRange)
}


// Removes a specific type of track association between two tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/removeTrackAssociation(to:type:)
func (m_ MutableMovieTrack) RemoveTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType TrackAssociationType /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTrackAssociationToTrack:type:"), movieTrack, trackAssociationType)
}


// Replaces the track’s format description with a new format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/replaceFormatDescription(_:with:)
func (m_ MutableMovieTrack) ReplaceFormatDescriptionWithFormatDescription(formatDescription FormatDescriptionRef /* not a class type */, newFormatDescription FormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceFormatDescription:withFormatDescription:"), formatDescription, newFormatDescription)
}


// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/samplePresentationTime(forTrackTime:)
func (m_ MutableMovieTrack) SamplePresentationTimeForTrackTime(trackTime objc.IObject /* cross-framework: Time */) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return rv
}


// Changes the duration of a time range in a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/scaleTimeRange(_:toDuration:)
func (m_ MutableMovieTrack) ScaleTimeRangeToDuration(timeRange TimeRange /* not a class type */, duration objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}


// Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/segment(forTrackTime:)
func (m_ MutableMovieTrack) SegmentForTrackTime(trackTime objc.IObject /* cross-framework: Time */) IAssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](m_.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return rv
}







// A number that identifies the track as a member of a particular alternate group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/alternateGroupID
func (m_ MutableMovieTrack) AlternateGroupID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("alternateGroupID"))
	return rv
}


// A number that identifies the track as a member of a particular alternate group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/alternateGroupID
func (m_ MutableMovieTrack) SetAlternateGroupID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlternateGroupID:"), value)
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableMetadataFormats
func (m_ MutableMovieTrack) AvailableMetadataFormats() MetadataFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of metadata formats available for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableMetadataFormats
func (m_ MutableMovieTrack) SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableTrackAssociationTypes
func (m_ MutableMovieTrack) AvailableTrackAssociationTypes() objc.IObject {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}


// An array of association types that the track uses to associate with other tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableTrackAssociationTypes
func (m_ MutableMovieTrack) SetAvailableTrackAssociationTypes(value objc.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/canProvideSampleCursors
func (m_ MutableMovieTrack) CanProvideSampleCursors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}


// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/canProvideSampleCursors
func (m_ MutableMovieTrack) SetCanProvideSampleCursors(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}


// The clean aperture dimension of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/cleanApertureDimensions
func (m_ MutableMovieTrack) CleanApertureDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("cleanApertureDimensions"))
	return rv
}


// The clean aperture dimension of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/cleanApertureDimensions
func (m_ MutableMovieTrack) SetCleanApertureDimensions(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCleanApertureDimensions:"), value)
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/commonMetadata
func (m_ MutableMovieTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("commonMetadata"))
	return rv
}


// An array of metadata items for all common metadata keys that have a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/commonMetadata
func (m_ MutableMovieTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommonMetadata:"), value)
}


// The encoded pixels dimensions of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/encodedPixelsDimensions
func (m_ MutableMovieTrack) EncodedPixelsDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("encodedPixelsDimensions"))
	return rv
}


// The encoded pixels dimensions of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/encodedPixelsDimensions
func (m_ MutableMovieTrack) SetEncodedPixelsDimensions(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodedPixelsDimensions:"), value)
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/estimatedDataRate
func (m_ MutableMovieTrack) EstimatedDataRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("estimatedDataRate"))
	return rv
}


// The estimated data rate, in bits per second, of the media that the track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/estimatedDataRate
func (m_ MutableMovieTrack) SetEstimatedDataRate(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEstimatedDataRate:"), value)
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/extendedLanguageTag
func (m_ MutableMovieTrack) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The language tag of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/extendedLanguageTag
func (m_ MutableMovieTrack) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/formatDescriptions
func (m_ MutableMovieTrack) FormatDescriptions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("formatDescriptions"))
	return rv
}


// The format descriptions of the media samples that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/formatDescriptions
func (m_ MutableMovieTrack) SetFormatDescriptions(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFormatDescriptions:"), value)
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasAudioSampleDependencies
func (m_ MutableMovieTrack) HasAudioSampleDependencies() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}


// A Boolean value that indicates whether the track has sample dependencies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasAudioSampleDependencies
func (m_ MutableMovieTrack) SetHasAudioSampleDependencies(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
}


// A Boolean value that indicates whether a track contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasProtectedContent
func (m_ MutableMovieTrack) HasProtectedContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isDecodable
func (m_ MutableMovieTrack) IsDecodable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isDecodable"))
	return rv
}


// A Boolean value that indicates whether the track is decodable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isDecodable
func (m_ MutableMovieTrack) SetIsDecodable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDecodable:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isEnabled
func (m_ MutableMovieTrack) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isEnabled
func (m_ MutableMovieTrack) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean value that indicates whether a track is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isModified
func (m_ MutableMovieTrack) Modified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("modified"))
	return rv
}


// A Boolean value that indicates whether a track is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isModified
func (m_ MutableMovieTrack) SetModified(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModified:"), value)
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isPlayable
func (m_ MutableMovieTrack) IsPlayable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the track is playable in the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isPlayable
func (m_ MutableMovieTrack) SetIsPlayable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isSelfContained
func (m_ MutableMovieTrack) IsSelfContained() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isSelfContained"))
	return rv
}


// A Boolean value that indicates whether this track references sample data only within its container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isSelfContained
func (m_ MutableMovieTrack) SetIsSelfContained(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSelfContained:"), value)
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/languageCode
func (m_ MutableMovieTrack) LanguageCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/languageCode
func (m_ MutableMovieTrack) SetLanguageCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLanguageCode:"), value)
}


// The layer level for the visual media of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/layer
func (m_ MutableMovieTrack) Layer() int {
	rv := objc.Send[int](m_.ID, objc.Sel("layer"))
	return rv
}


// The layer level for the visual media of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/layer
func (m_ MutableMovieTrack) SetLayer(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLayer:"), value)
}


// A storage container for the media data to be added to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/mediaDataStorage
func (m_ MutableMovieTrack) MediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("mediaDataStorage"))
	return rv
}


// A storage container for the media data to be added to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/mediaDataStorage
func (m_ MutableMovieTrack) SetMediaDataStorage(value IAVMediaDataStorage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaDataStorage:"), value)
}


// An array of metadata stored by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/metadata
func (m_ MutableMovieTrack) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata stored by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/metadata
func (m_ MutableMovieTrack) SetMetadata(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadata:"), nsArray)
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/minFrameDuration
func (m_ MutableMovieTrack) MinFrameDuration() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The minimum duration of the track’s frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/minFrameDuration
func (m_ MutableMovieTrack) SetMinFrameDuration(value Time get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// The dimensions used to display the visual media data for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalSize
func (m_ MutableMovieTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// The dimensions used to display the visual media data for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalSize
func (m_ MutableMovieTrack) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalTimeScale
func (m_ MutableMovieTrack) NaturalTimeScale() TimeScale get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("naturalTimeScale"))
	return rv
}


// The natural time scale of the media that a track references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalTimeScale
func (m_ MutableMovieTrack) SetNaturalTimeScale(value TimeScale get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalTimeScale:"), value)
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/nominalFrameRate
func (m_ MutableMovieTrack) NominalFrameRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The frame rate of the track, in frames per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/nominalFrameRate
func (m_ MutableMovieTrack) SetNominalFrameRate(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalFrameRate:"), value)
}


// The boundary for media chunk alignment for file types that support media chunk alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkAlignment
func (m_ MutableMovieTrack) PreferredMediaChunkAlignment() int {
	rv := objc.Send[int](m_.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}


// The boundary for media chunk alignment for file types that support media chunk alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkAlignment
func (m_ MutableMovieTrack) SetPreferredMediaChunkAlignment(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}


// The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkDuration
func (m_ MutableMovieTrack) PreferredMediaChunkDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("preferredMediaChunkDuration"))
	return rv
}


// The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkDuration
func (m_ MutableMovieTrack) SetPreferredMediaChunkDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}


// The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkSize
func (m_ MutableMovieTrack) PreferredMediaChunkSize() int {
	rv := objc.Send[int](m_.ID, objc.Sel("preferredMediaChunkSize"))
	return rv
}


// The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkSize
func (m_ MutableMovieTrack) SetPreferredMediaChunkSize(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMediaChunkSize:"), value)
}


// The transform performed on the visual media data of the track for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredTransform
func (m_ MutableMovieTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The transform performed on the visual media data of the track for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredTransform
func (m_ MutableMovieTrack) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The preferred volume for the audible medata data of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredVolume
func (m_ MutableMovieTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The preferred volume for the audible medata data of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredVolume
func (m_ MutableMovieTrack) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredVolume:"), value)
}


// The production aperture dimensions of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/productionApertureDimensions
func (m_ MutableMovieTrack) ProductionApertureDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("productionApertureDimensions"))
	return rv
}


// The production aperture dimensions of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/productionApertureDimensions
func (m_ MutableMovieTrack) SetProductionApertureDimensions(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductionApertureDimensions:"), value)
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/requiresFrameReordering
func (m_ MutableMovieTrack) RequiresFrameReordering() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}


// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/requiresFrameReordering
func (m_ MutableMovieTrack) SetRequiresFrameReordering(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}


// The base URL for sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/sampleReferenceBaseURL
func (m_ MutableMovieTrack) SampleReferenceBaseURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("sampleReferenceBaseURL"))
	return rv
}


// The base URL for sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/sampleReferenceBaseURL
func (m_ MutableMovieTrack) SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/segments
func (m_ MutableMovieTrack) Segments() IAVAssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](m_.ID, objc.Sel("segments"))
	return rv
}


// The time mappings from the track’s media samples to its timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/segments
func (m_ MutableMovieTrack) SetSegments(value IAVAssetTrackSegment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegments:"), value)
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timeRange
func (m_ MutableMovieTrack) TimeRange() TimeRange get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the track within the overall timeline of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timeRange
func (m_ MutableMovieTrack) SetTimeRange(value TimeRange get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}


// The time scale for tracks that contain the atom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timescale
func (m_ MutableMovieTrack) Timescale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](m_.ID, objc.Sel("timescale"))
	return rv
}


// The time scale for tracks that contain the atom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timescale
func (m_ MutableMovieTrack) SetTimescale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimescale:"), value)
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/totalSampleDataLength
func (m_ MutableMovieTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("totalSampleDataLength"))
	return rv
}


// The total number of bytes of sample data the track requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/totalSampleDataLength
func (m_ MutableMovieTrack) SetTotalSampleDataLength(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalSampleDataLength:"), value)
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/isenabled
func (m_ MutableMovieTrack) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the track’s container enables it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/isenabled
func (m_ MutableMovieTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether a track is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/ismodified
func (m_ MutableMovieTrack) IsModified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isModified"))
	return rv
}


// A Boolean value that indicates whether a track is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/ismodified
func (m_ MutableMovieTrack) SetIsModified(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsModified:"), value)
}








