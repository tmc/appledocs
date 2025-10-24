// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableMovie] class.
var (
	MutableMovieClass     _MutableMovieClass
	MutableMovieClassOnce sync.Once
)

func getMutableMovieClass() _MutableMovieClass {
	MutableMovieClassOnce.Do(func() {
		MutableMovieClass = _MutableMovieClass{objc.GetClass("AVMutableMovie")}
	})
	return MutableMovieClass
}

type _MutableMovieClass struct {
	class objc.Class
}





// An interface definition for the [MutableMovie] class.
type IMutableMovie interface {
	IMovie
	

	// properties:
	AllMediaSelections() IAVMediaSelection
	SetAllMediaSelections(value IAVMediaSelection)
	AvailableChapterLocales() objectivec.IObject
	SetAvailableChapterLocales(value objectivec.IObject)
	AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic get /* not a class type */
	SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic get /* not a class type */)
	AvailableMetadataFormats() MetadataFormat get /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */)
	CanContainFragments() objectivec.IObject
	SetCanContainFragments(value objectivec.IObject)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	ContainsFragments() objectivec.IObject
	SetContainsFragments(value objectivec.IObject)
	CreationDate() IAVMetadataItem
	SetCreationDate(value IAVMetadataItem)
	DefaultMediaDataStorage() IAVMediaDataStorage
	SetDefaultMediaDataStorage(value IAVMediaDataStorage)
	Duration() Time get /* not a class type */
	SetDuration(value Time get /* not a class type */)
	HasProtectedContent() objectivec.IObject
	SetHasProtectedContent(value objectivec.IObject)
	InterleavingPeriod() objc.IObject /* cross-framework: Time */
	SetInterleavingPeriod(value objc.IObject /* cross-framework: Time */)
	IsCompatibleWithAirPlayVideo() objectivec.IObject
	SetIsCompatibleWithAirPlayVideo(value objectivec.IObject)
	IsComposable() objectivec.IObject
	SetIsComposable(value objectivec.IObject)
	IsExportable() objectivec.IObject
	SetIsExportable(value objectivec.IObject)
	Modified() bool
	SetModified(value bool)
	IsPlayable() objectivec.IObject
	SetIsPlayable(value objectivec.IObject)
	IsReadable() objectivec.IObject
	SetIsReadable(value objectivec.IObject)
	Lyrics() objectivec.IObject
	SetLyrics(value objectivec.IObject)
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MinimumTimeOffsetFromLive() Time get /* not a class type */
	SetMinimumTimeOffsetFromLive(value Time get /* not a class type */)
	OverallDurationHint() Time get /* not a class type */
	SetOverallDurationHint(value Time get /* not a class type */)
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)
	PreferredRate() float32
	SetPreferredRate(value float32)
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	ProvidesPreciseDurationAndTiming() objectivec.IObject
	SetProvidesPreciseDurationAndTiming(value objectivec.IObject)
	Timescale() TimeScale /* not a class type */
	SetTimescale(value TimeScale /* not a class type */)
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
	Tracks() []MutableMovieTrack
	AVAssetExportPresetPassthrough() objc.IObject /* cross-framework: NSString */
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	IsModified() bool
	SetIsModified(value bool)


	

	// methods:
	AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType MediaType /* typedef */, track IAVAssetTrack, options foundation.IDictionary) IMutableMovieTrack
	AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []AssetTrack, options foundation.IDictionary) []MutableMovieTrack
	ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []TimedMetadataGroup
	ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.Locale, commonKeys []string) []TimedMetadataGroup
	InsertEmptyTimeRange(timeRange TimeRange /* not a class type */)
	InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange TimeRange /* not a class type */, asset IAVAsset, startTime objc.IObject /* cross-framework: Time */, copySampleData bool, outError objectivec.IObject) bool
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer)
	MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IMediaSelectionGroup
	MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem
	MutableTrackCompatibleWithTrack(track IAVAssetTrack) IMutableMovieTrack
	RemoveTimeRange(timeRange TimeRange /* not a class type */)
	RemoveTrack(track IAVMovieTrack)
	ScaleTimeRangeToDuration(timeRange TimeRange /* not a class type */, duration objc.IObject /* cross-framework: Time */)
	TrackWithTrackID(trackID PersistentTrackID /* not a class type */) IMutableMovieTrack
	TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) []MutableMovieTrack
	TracksWithMediaType(mediaType MediaType /* typedef */) []MutableMovieTrack
	UnusedTrackID() PersistentTrackID /* not a class type */


}





// Alloc allocates a new instance without initialization.
func (mc _MutableMovieClass) Alloc() MutableMovie {
	rv := objc.Send[MutableMovie](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableMovieClass) New() MutableMovie {
	rv := objc.Send[MutableMovie](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableMovie) Init() MutableMovie {
	rv := objc.Send[MutableMovie](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableMovie) Autorelease() MutableMovie {
	rv := objc.Send[MutableMovie](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableMovie creates a new MutableMovie instance.
func NewMutableMovie() MutableMovie {
	return getMutableMovieClass().New()
}





// A mutable object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// This class is a mutable subclass of that provides methods that support movie editing. For example, you can use a mutable movie to copy media data from one track and paste it into another. You can also use this object to create track references from one track to another (for example, to set one track as a chapter track of another track). To perform editing operations on individual tracks, use the associated classes and . You use movie objects only when operating on format-specific features of a QuickTime or ISO base media file. You typically don’t use these classes to open and play QuickTime movie files or ISO base media files. Instead, you use and . When performing media insertions, a movie interleaves media data from tracks in the source asset to optimize the movie file for playback. However, performing a series of media insertions may result in a movie file that’s not optimally interleaved. You can optimize a movie file for playback by exporting it with an object using the export preset , and setting the property value to .


// A mutable object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie
type MutableMovie struct {
	Movie
}

// MutableMovieFrom constructs a [MutableMovie] from an unsafe.Pointer.
//
// A mutable object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
func MutableMovieFrom(ptr unsafe.Pointer) MutableMovie {
	return MutableMovie{
		Movie: MovieFrom(ptr),
	}
}






// Creates a mutable movie object from a movie stored in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(data:options:error:)
func NewMutableMovieWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, outError objectivec.IObject) MutableMovie {
	instance := getMutableMovieClass().Alloc()
	rv := objc.Send[MutableMovie](instance.ID, objc.Sel("initWithData:options:error:"), data, options, outError)
	rv.Autorelease()
	return rv
}


// Creates a mutable movie object without tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(settingsFrom:options:)
func NewMutableMovieWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.IDictionary, outError objectivec.IObject) MutableMovie {
	instance := getMutableMovieClass().Alloc()
	rv := objc.Send[MutableMovie](instance.ID, objc.Sel("initWithSettingsFromMovie:options:error:"), movie, options, outError)
	rv.Autorelease()
	return rv
}


// Creates a mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(url:options:error:)
func NewMutableMovieWithURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, outError objectivec.IObject) MutableMovie {
	instance := getMutableMovieClass().Alloc()
	rv := objc.Send[MutableMovie](instance.ID, objc.Sel("initWithURL:options:error:"), URL, options, outError)
	rv.Autorelease()
	return rv
}







// Returns a new mutable movie object from a movie stored in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithData:options:error:
func (mc _MutableMovieClass) MovieWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("movieWithData:options:error:"), data, options, outError)
	return rv
}


// Returns a new mutable movie object without tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithSettingsFromMovie:options:error:
func (mc _MutableMovieClass) MovieWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.IDictionary, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("movieWithSettingsFromMovie:options:error:"), movie, options, outError)
	return rv
}


// Returns a new mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithURL:options:error:
func (mc _MutableMovieClass) MovieWithURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("movieWithURL:options:error:"), URL, options, outError)
	return rv
}












// Adds an empty track to the target movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/addMutableTrack(withMediaType:copySettingsFrom:options:)
func (m_ MutableMovie) AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType MediaType /* typedef */, track IAVAssetTrack, options foundation.IDictionary) IMutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](m_.ID, objc.Sel("addMutableTrackWithMediaType:copySettingsFromTrack:options:"), mediaType, track, options)
	return rv
}


// Adds one or more empty tracks to the target movie and copies the track settings from the source tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/addMutableTracksCopyingSettings(from:options:)
func (m_ MutableMovie) AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []AssetTrack, options foundation.IDictionary) []MutableMovieTrack {
	rv := objc.Send[[]MutableMovieTrack](m_.ID, objc.Sel("addMutableTracksCopyingSettingsFromTracks:options:"), existingTracks, options)
	return rv
}


// Returns an array of chapters with a locale that best matches the list of preferred languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/chapterMetadataGroups(bestMatchingPreferredLanguages:)
func (m_ MutableMovie) ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](m_.ID, objc.Sel("chapterMetadataGroupsBestMatchingPreferredLanguages:"), preferredLanguages)
	return rv
}


// Returns an array of chapters that contain the specified title locale and common keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (m_ MutableMovie) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.Locale, commonKeys []string) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](m_.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, commonKeys)
	return rv
}


// Adds an empty time range to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/insertEmptyTimeRange(_:)
func (m_ MutableMovie) InsertEmptyTimeRange(timeRange TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}


// Inserts all of the tracks in a specified time range of an asset into a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/insertTimeRange(_:of:at:copySampleData:)
func (m_ MutableMovie) InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange TimeRange /* not a class type */, asset IAVAsset, startTime objc.IObject /* cross-framework: Time */, copySampleData bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertTimeRange:ofAsset:atTime:copySampleData:error:"), timeRange, asset, startTime, copySampleData, outError)
	return rv
}


// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/loadTrack(withTrackID:completionHandler:)
func (m_ MutableMovie) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/loadTracks(withMediaCharacteristic:completionHandler:)
func (m_ MutableMovie) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/loadTracks(withMediaType:completionHandler:)
func (m_ MutableMovie) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}


// Returns a media selection group that contains one or more options with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/mediaSelectionGroup(forMediaCharacteristic:)
func (m_ MutableMovie) MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IMediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](m_.ID, objc.Sel("mediaSelectionGroupForMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Returns an array of metadata items from the container with the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/metadata(forFormat:)
func (m_ MutableMovie) MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}


// Provides a reference to a track from a mutable movie into which you can insert any time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/mutableTrack(compatibleWith:)
func (m_ MutableMovie) MutableTrackCompatibleWithTrack(track IAVAssetTrack) IMutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](m_.ID, objc.Sel("mutableTrackCompatibleWithTrack:"), track)
	return rv
}


// Removes the specified time range from a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/removeTimeRange(_:)
func (m_ MutableMovie) RemoveTimeRange(timeRange TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTimeRange:"), timeRange)
}


// Removes the specified track from the target movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/removeTrack(_:)
func (m_ MutableMovie) RemoveTrack(track IAVMovieTrack) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTrack:"), track)
}


// Changes the duration of a time range in a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/scale(_:toDuration:)
func (m_ MutableMovie) ScaleTimeRangeToDuration(timeRange TimeRange /* not a class type */, duration objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}


// Retrieves a track in the movie that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/track(withTrackID:)
func (m_ MutableMovie) TrackWithTrackID(trackID PersistentTrackID /* not a class type */) IMutableMovieTrack {
	rv := objc.Send[MutableMovieTrack](m_.ID, objc.Sel("trackWithTrackID:"), trackID)
	return rv
}


// Retrieve tracks in the movie that present media of the specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/tracks(withMediaCharacteristic:)
func (m_ MutableMovie) TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) []MutableMovieTrack {
	rv := objc.Send[[]MutableMovieTrack](m_.ID, objc.Sel("tracksWithMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Retrieves tracks in the movie that present media of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/tracks(withMediaType:)
func (m_ MutableMovie) TracksWithMediaType(mediaType MediaType /* typedef */) []MutableMovieTrack {
	rv := objc.Send[[]MutableMovieTrack](m_.ID, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}


// Returns an identifier that no other tracks in the asset use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/unusedTrackID()
func (m_ MutableMovie) UnusedTrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("unusedTrackID"))
	return rv
}







// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/allMediaSelections
func (m_ MutableMovie) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/allMediaSelections
func (m_ MutableMovie) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllMediaSelections:"), value)
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableChapterLocales
func (m_ MutableMovie) AvailableChapterLocales() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableChapterLocales
func (m_ MutableMovie) SetAvailableChapterLocales(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMediaCharacteristicsWithMediaSelectionOptions
func (m_ MutableMovie) AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMediaCharacteristicsWithMediaSelectionOptions
func (m_ MutableMovie) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMetadataFormats
func (m_ MutableMovie) AvailableMetadataFormats() MetadataFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMetadataFormats
func (m_ MutableMovie) SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/canContainFragments
func (m_ MutableMovie) CanContainFragments() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("canContainFragments"))
	return rv
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/canContainFragments
func (m_ MutableMovie) SetCanContainFragments(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanContainFragments:"), value)
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/commonMetadata
func (m_ MutableMovie) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("commonMetadata"))
	return rv
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/commonMetadata
func (m_ MutableMovie) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommonMetadata:"), value)
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/containsFragments
func (m_ MutableMovie) ContainsFragments() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("containsFragments"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/containsFragments
func (m_ MutableMovie) SetContainsFragments(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContainsFragments:"), value)
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/creationDate
func (m_ MutableMovie) CreationDate() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("creationDate"))
	return rv
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/creationDate
func (m_ MutableMovie) SetCreationDate(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreationDate:"), value)
}


// The default storage container for media data that you add to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/defaultMediaDataStorage
func (m_ MutableMovie) DefaultMediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("defaultMediaDataStorage"))
	return rv
}


// The default storage container for media data that you add to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/defaultMediaDataStorage
func (m_ MutableMovie) SetDefaultMediaDataStorage(value IAVMediaDataStorage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultMediaDataStorage:"), value)
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/duration
func (m_ MutableMovie) Duration() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("duration"))
	return rv
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/duration
func (m_ MutableMovie) SetDuration(value Time get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/hasProtectedContent
func (m_ MutableMovie) HasProtectedContent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/hasProtectedContent
func (m_ MutableMovie) SetHasProtectedContent(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasProtectedContent:"), value)
}


// A time period indicating the duration for interleaving runs of samples for each track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/interleavingPeriod
func (m_ MutableMovie) InterleavingPeriod() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("interleavingPeriod"))
	return rv
}


// A time period indicating the duration for interleaving runs of samples for each track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/interleavingPeriod
func (m_ MutableMovie) SetInterleavingPeriod(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInterleavingPeriod:"), value)
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isCompatibleWithAirPlayVideo
func (m_ MutableMovie) IsCompatibleWithAirPlayVideo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isCompatibleWithAirPlayVideo
func (m_ MutableMovie) SetIsCompatibleWithAirPlayVideo(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isComposable
func (m_ MutableMovie) IsComposable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isComposable"))
	return rv
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isComposable
func (m_ MutableMovie) SetIsComposable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsComposable:"), value)
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isExportable
func (m_ MutableMovie) IsExportable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isExportable"))
	return rv
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isExportable
func (m_ MutableMovie) SetIsExportable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsExportable:"), value)
}


// A Boolean value that indicates whether the movie is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isModified
func (m_ MutableMovie) Modified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("modified"))
	return rv
}


// A Boolean value that indicates whether the movie is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isModified
func (m_ MutableMovie) SetModified(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModified:"), value)
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isPlayable
func (m_ MutableMovie) IsPlayable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isPlayable
func (m_ MutableMovie) SetIsPlayable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isReadable
func (m_ MutableMovie) IsReadable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isReadable"))
	return rv
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isReadable
func (m_ MutableMovie) SetIsReadable(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsReadable:"), value)
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/lyrics
func (m_ MutableMovie) Lyrics() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("lyrics"))
	return rv
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/lyrics
func (m_ MutableMovie) SetLyrics(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLyrics:"), value)
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/metadata
func (m_ MutableMovie) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/metadata
func (m_ MutableMovie) SetMetadata(value []MetadataItem) {
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


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/minimumTimeOffsetFromLive
func (m_ MutableMovie) MinimumTimeOffsetFromLive() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/minimumTimeOffsetFromLive
func (m_ MutableMovie) SetMinimumTimeOffsetFromLive(value Time get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/overallDurationHint
func (m_ MutableMovie) OverallDurationHint() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/overallDurationHint
func (m_ MutableMovie) SetOverallDurationHint(value Time get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredMediaSelection
func (m_ MutableMovie) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredMediaSelection
func (m_ MutableMovie) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredMediaSelection:"), value)
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredRate
func (m_ MutableMovie) PreferredRate() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("preferredRate"))
	return rv
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredRate
func (m_ MutableMovie) SetPreferredRate(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredRate:"), value)
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredTransform
func (m_ MutableMovie) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredTransform
func (m_ MutableMovie) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredVolume
func (m_ MutableMovie) PreferredVolume() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredVolume
func (m_ MutableMovie) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/providesPreciseDurationAndTiming
func (m_ MutableMovie) ProvidesPreciseDurationAndTiming() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/providesPreciseDurationAndTiming
func (m_ MutableMovie) SetProvidesPreciseDurationAndTiming(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/timescale
func (m_ MutableMovie) Timescale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](m_.ID, objc.Sel("timescale"))
	return rv
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/timescale
func (m_ MutableMovie) SetTimescale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimescale:"), value)
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/trackGroups
func (m_ MutableMovie) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](m_.ID, objc.Sel("trackGroups"))
	return rv
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/trackGroups
func (m_ MutableMovie) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackGroups:"), value)
}


// The tracks that a movie contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/tracks
func (m_ MutableMovie) Tracks() []MutableMovieTrack {
	rv := objc.Send[[]MutableMovieTrack](m_.ID, objc.Sel("tracks"))
	return rv
}


// A preset to export the asset in its current format, unless otherwise prohibited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportpresetpassthrough
func (m_ MutableMovie) AVAssetExportPresetPassthrough() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("AVAssetExportPresetPassthrough"))
	return rv
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/shouldoptimizefornetworkuse
func (m_ MutableMovie) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/shouldoptimizefornetworkuse
func (m_ MutableMovie) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}


// A Boolean value that indicates whether the movie is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/ismodified
func (m_ MutableMovie) IsModified() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isModified"))
	return rv
}


// A Boolean value that indicates whether the movie is in a modified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovie/ismodified
func (m_ MutableMovie) SetIsModified(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsModified:"), value)
}







