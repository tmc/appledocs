// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Asset] class.
var (
	AssetClass     _AssetClass
	AssetClassOnce sync.Once
)

func getAssetClass() _AssetClass {
	AssetClassOnce.Do(func() {
		AssetClass = _AssetClass{objc.GetClass("AVAsset")}
	})
	return AssetClass
}

type _AssetClass struct {
	class objc.Class
}

// An interface definition for the [Asset] class.
type IAsset interface {
	objectivec.IObject
	CancelLoading()
	FindUnusedTrackIDWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer)
	MetadataForFormat(format unsafe.Pointer) []MetadataItem
	TrackWithTrackID(trackID unsafe.Pointer) unsafe.Pointer
	TracksWithMediaType(mediaType unsafe.Pointer) []AssetTrack
	UnusedTrackID() unsafe.Pointer
}

// An object that models timed audiovisual media.
//
// An asset models file-based media like a QuickTime movie or an MP3 audio file, and also media streamed using HTTP Live Streaming (HLS). An asset is a container object for one or more instances of that model the uniformly typed tracks of media. The most commonly used track types are audio and video, but assets may also contain supplementary tracks, like closed captions, subtitles, and timed metadata. You load the tracks for an asset by asynchronously loading its property. In some cases, you may want to perform operations on a subset of an asset’s tracks rather than on its complete collection. For those situations, an asset provides methods to retrieve subsets of tracks according to particular criteria, such as identifier, media type, or characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset
type Asset struct {
	objectivec.Object
}

// AssetFrom constructs a [Asset] from an unsafe.Pointer.
//
// An object that models timed audiovisual media.
func AssetFrom(ptr unsafe.Pointer) Asset {
	return Asset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetClass) Alloc() Asset {
	rv := objc.Send[Asset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetClass) New() Asset {
	rv := objc.Send[Asset](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Asset) Init() Asset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Asset) Autorelease() Asset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAsset creates a new Asset instance.
func NewAsset() Asset {
	return getAssetClass().New()
}


// Cancels all pending requests to asynchronously load property values.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/cancelLoading()
func (a_ Asset) CancelLoading() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelLoading"))
}

// Loads an identifier that no other track in the asset uses.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/findUnusedTrackID(completionHandler:)
func (a_ Asset) FindUnusedTrackIDWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("findUnusedTrackIDWithCompletionHandler:"), completionHandler)
}

// Loads chapter metadata with a locale that best matches the list of preferred languages.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadChapterMetadataGroups(bestMatchingPreferredLanguages:completionHandler:)
func (a_ Asset) LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadChapterMetadataGroupsBestMatchingPreferredLanguages:completionHandler:"), preferredLanguages, completionHandler)
}

// Loads a media selection group that contains one or more options with the specified media characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadMediaSelectionGroup(for:completionHandler:)
func (a_ Asset) LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMediaSelectionGroupForMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}

// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a_ Asset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}

// Returns an array of metadata items from the container with the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/metadata(forFormat:)
func (a_ Asset) MetadataForFormat(format unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}

// Returns a track that contains the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/track(withTrackID:)
func (a_ Asset) TrackWithTrackID(trackID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackWithTrackID:"), trackID)
	return rv
}

// Returns tracks that contain media of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/tracks(withMediaType:)
func (a_ Asset) TracksWithMediaType(mediaType unsafe.Pointer) []AssetTrack {
	rv := objc.Send[[]AssetTrack](a_.ID, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}

// Returns an identifier that no other tracks in the asset use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/unusedTrackID()
func (a_ Asset) UnusedTrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("unusedTrackID"))
	return rv
}

// The locales of the asset’s chapter metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableChapterLocales
func (a_ Asset) AvailableChapterLocales() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](a_.ID, objc.Sel("availableChapterLocales"))
	return rv
}

// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/hasProtectedContent
func (a_ Asset) HasProtectedContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasProtectedContent"))
	return rv
}

// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isCompatibleWithAirPlayVideo
func (a_ Asset) CompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compatibleWithAirPlayVideo"))
	return rv
}

// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isCompatibleWithSavedPhotosAlbum
func (a_ Asset) CompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compatibleWithSavedPhotosAlbum"))
	return rv
}

// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isExportable
func (a_ Asset) Exportable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("exportable"))
	return rv
}

// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/metadata
func (a_ Asset) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}

// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/naturalSize
func (a_ Asset) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}

// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredDisplayCriteria
func (a_ Asset) PreferredDisplayCriteria() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}

// The default media selections for this asset’s media selection groups.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredMediaSelection
func (a_ Asset) PreferredMediaSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}

// The asset’s volume preference for playing its audible media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredVolume
func (a_ Asset) PreferredVolume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredVolume"))
	return rv
}

// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/referenceRestrictions
func (a_ Asset) ReferenceRestrictions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("referenceRestrictions"))
	return rv
}

// The array of available media selections for this asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/allmediaselections
func (a_ Asset) AllMediaSelections() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// SetAllMediaSelections sets the value of the allMediaSelections property.
// The array of available media selections for this asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/allmediaselections
func (a_ Asset) SetAllMediaSelections(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllMediaSelections:"), value)
}

// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) AvailableMediaCharacteristicsWithMediaSelectionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// SetAvailableMediaCharacteristicsWithMediaSelectionOptions sets the value of the availableMediaCharacteristicsWithMediaSelectionOptions property.
// An array of media characteristics for which a media selection option is available.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}

// The metadata formats this asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) AvailableMetadataFormats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// SetAvailableMetadataFormats sets the value of the availableMetadataFormats property.
// The metadata formats this asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) SetAvailableMetadataFormats(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}

// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (a_ Asset) CanContainFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canContainFragments"))
	return rv
}


// SetCanContainFragments sets the value of the canContainFragments property.
// A Boolean value that indicates whether you can extend the asset by fragments.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (a_ Asset) SetCanContainFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanContainFragments:"), value)
}

// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (a_ Asset) CommonMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("commonMetadata"))
	return rv
}


// SetCommonMetadata sets the value of the commonMetadata property.
// The metadata items an asset contains for common metadata identifiers that provide a value.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (a_ Asset) SetCommonMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCommonMetadata:"), value)
}

// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/containsfragments
func (a_ Asset) ContainsFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsFragments"))
	return rv
}


// SetContainsFragments sets the value of the containsFragments property.
// A Boolean value that indicates whether at least one movie fragment extends the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/containsfragments
func (a_ Asset) SetContainsFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContainsFragments:"), value)
}

// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/creationdate
func (a_ Asset) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// A metadata item that indicates the asset’s creation date.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/creationdate
func (a_ Asset) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCreationDate:"), value)
}

// A time value that indicates the asset’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/duration
func (a_ Asset) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// A time value that indicates the asset’s duration.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/duration
func (a_ Asset) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}

// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// SetIsCompatibleWithAirPlayVideo sets the value of the isCompatibleWithAirPlayVideo property.
// A Boolean value that indicates whether the asset is compatible with AirPlay Video.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}

// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) IsCompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}


// SetIsCompatibleWithSavedPhotosAlbum sets the value of the isCompatibleWithSavedPhotosAlbum property.
// A Boolean value that indicates whether you can write the asset to the Saved Photos album.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithSavedPhotosAlbum:"), value)
}

// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) IsComposable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isComposable"))
	return rv
}


// SetIsComposable sets the value of the isComposable property.
// A Boolean value that indicates whether you can use the asset as a segment of a composition track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) SetIsComposable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsComposable:"), value)
}

// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) IsExportable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isExportable"))
	return rv
}


// SetIsExportable sets the value of the isExportable property.
// A Boolean value that indicates whether you can export this asset using an export session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) SetIsExportable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsExportable:"), value)
}

// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) IsPlayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}


// SetIsPlayable sets the value of the isPlayable property.
// A Boolean value that indicates whether the asset has playable content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) SetIsPlayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}

// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) IsReadable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadable"))
	return rv
}


// SetIsReadable sets the value of the isReadable property.
// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) SetIsReadable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadable:"), value)
}

// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) Lyrics() string {
	rv := objc.Send[string](a_.ID, objc.Sel("lyrics"))
	return rv
}


// SetLyrics sets the value of the lyrics property.
// The lyrics of the asset in a language suitable for the current locale.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) SetLyrics(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLyrics:"), objc.String(value))
}

// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/minimumtimeoffsetfromlive
func (a_ Asset) MinimumTimeOffsetFromLive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// SetMinimumTimeOffsetFromLive sets the value of the minimumTimeOffsetFromLive property.
// A time value that indicates how closely playback follows the latest live stream content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/minimumtimeoffsetfromlive
func (a_ Asset) SetMinimumTimeOffsetFromLive(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}

// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) OverallDurationHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// SetOverallDurationHint sets the value of the overallDurationHint property.
// The total duration of fragments that currently exist, or may exist in the future.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) SetOverallDurationHint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}

// The asset’s rate preference for playing its media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredrate
func (a_ Asset) PreferredRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredRate"))
	return rv
}


// SetPreferredRate sets the value of the preferredRate property.
// The asset’s rate preference for playing its media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredrate
func (a_ Asset) SetPreferredRate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredRate:"), value)
}

// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// SetPreferredTransform sets the value of the preferredTransform property.
// The asset’s transform preference to apply to its visual content during presentation or processing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredTransform:"), value)
}

// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// SetProvidesPreciseDurationAndTiming sets the value of the providesPreciseDurationAndTiming property.
// A Boolean value that indicates whether the asset provides precise duration and timing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}

// The track groups an asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/trackgroups
func (a_ Asset) TrackGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackGroups"))
	return rv
}


// SetTrackGroups sets the value of the trackGroups property.
// The track groups an asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/trackgroups
func (a_ Asset) SetTrackGroups(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackGroups:"), value)
}

// The tracks an asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/tracks
func (a_ Asset) Tracks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("tracks"))
	return rv
}


// SetTracks sets the value of the tracks property.
// The tracks an asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/tracks
func (a_ Asset) SetTracks(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTracks:"), value)
}



