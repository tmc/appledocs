// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	AllMediaSelections() []MediaSelection
	AvailableChapterLocales() []foundation.Locale
	AvailableMediaCharacteristicsWithMediaSelectionOptions() []string
	AvailableMetadataFormats() []string
	CanContainFragments() bool
	CommonMetadata() []MetadataItem
	ContainsFragments() bool
	CreationDate() IAVMetadataItem
	Duration() objectivec.IObject
	HasProtectedContent() bool
	CompatibleWithAirPlayVideo() bool
	Composable() bool
	Exportable() bool
	Playable() bool
	Readable() bool
	Lyrics() foundation.foundation.INSString
	Metadata() []MetadataItem
	MinimumTimeOffsetFromLive() objectivec.IObject
	OverallDurationHint() objectivec.IObject
	PreferredMediaSelection() IAVMediaSelection
	PreferredRate() float32
	PreferredTransform() corefoundation.CGAffineTransform
	PreferredVolume() float32
	ProvidesPreciseDurationAndTiming() bool
	ReferenceRestrictions() AssetReferenceRestrictions
	TrackGroups() []AssetTrackGroup
	Tracks() []AssetTrack
	IsCompatibleWithAirPlayVideo() bool
	SetIsCompatibleWithAirPlayVideo(value bool)
	IsCompatibleWithSavedPhotosAlbum() bool
	SetIsCompatibleWithSavedPhotosAlbum(value bool)
	IsComposable() bool
	SetIsComposable(value bool)
	IsExportable() bool
	SetIsExportable(value bool)
	IsPlayable() bool
	SetIsPlayable(value bool)
	IsReadable() bool
	SetIsReadable(value bool)


	

	// methods:
	CancelLoading()
	FindUnusedTrackIDWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.Locale, commonKeys []string, completionHandler unsafe.Pointer)
	LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler unsafe.Pointer)
	LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer)
	LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler unsafe.Pointer)
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetClass) Alloc() Asset {
	rv := objc.Send[Asset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that models timed audiovisual media.
//
// An asset models file-based media like a QuickTime movie or an MP3 audio file, and also media streamed using HTTP Live Streaming (HLS). An asset is a container object for one or more instances of that model the uniformly typed tracks of media. The most commonly used track types are audio and video, but assets may also contain supplementary tracks, like closed captions, subtitles, and timed metadata. You load the tracks for an asset by asynchronously loading its property. In some cases, you may want to perform operations on a subset of an asset’s tracks rather than on its complete collection. For those situations, an asset provides methods to retrieve subsets of tracks according to particular criteria, such as identifier, media type, or characteristic.


// An object that models timed audiovisual media.
//
// [Full Topic]
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






// Creates an asset that models the media at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewAssetWithURL(URL foundation.foundation.INSURL) Asset {
	rv := objc.Send[Asset](objc.ID(getAssetClass().class), objc.Sel("assetWithURL:"), URL)
	return rv
}







// Creates an asset that models the media at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func (ac _AssetClass) AssetWithURL(URL foundation.foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWithURL:"), URL)
	return rv
}












// Cancels all pending requests to asynchronously load property values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/cancelLoading()
func (a_ Asset) CancelLoading() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelLoading"))
}


// Loads an identifier that no other track in the asset uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/findUnusedTrackID(completionHandler:)
func (a_ Asset) FindUnusedTrackIDWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("findUnusedTrackIDWithCompletionHandler:"), completionHandler)
}


// Loads chapter metadata that contains the specified title locale and common keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadChapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:completionHandler:
func (a_ Asset) LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.Locale, commonKeys []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadChapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:completionHandler:"), locale, commonKeys, completionHandler)
}


// Loads chapter metadata with a locale that best matches the list of preferred languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadChapterMetadataGroups(bestMatchingPreferredLanguages:completionHandler:)
func (a_ Asset) LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadChapterMetadataGroupsBestMatchingPreferredLanguages:completionHandler:"), preferredLanguages, completionHandler)
}


// Loads a media selection group that contains one or more options with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadMediaSelectionGroup(for:completionHandler:)
func (a_ Asset) LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMediaSelectionGroupForMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// Loads an array of metadata items that the asset contains for the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadMetadata(for:completionHandler:)
func (a_ Asset) LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}


// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTrack(withTrackID:completionHandler:)
func (a_ Asset) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a_ Asset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaType:completionHandler:)
func (a_ Asset) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}







// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/allMediaSelections
func (a_ Asset) AllMediaSelections() []MediaSelection {
	rv := objc.Send[[]MediaSelection](a_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableChapterLocales
func (a_ Asset) AvailableChapterLocales() []foundation.Locale {
	rv := objc.Send[[]foundation.Locale](a_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableMediaCharacteristicsWithMediaSelectionOptions
func (a_ Asset) AvailableMediaCharacteristicsWithMediaSelectionOptions() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableMetadataFormats
func (a_ Asset) AvailableMetadataFormats() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/canContainFragments
func (a_ Asset) CanContainFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canContainFragments"))
	return rv
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/commonMetadata
func (a_ Asset) CommonMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/containsFragments
func (a_ Asset) ContainsFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsFragments"))
	return rv
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/creationDate
func (a_ Asset) CreationDate() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("creationDate"))
	return rv
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/duration
func (a_ Asset) Duration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("duration"))
	return rv
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/hasProtectedContent
func (a_ Asset) HasProtectedContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isCompatibleWithAirPlayVideo
func (a_ Asset) CompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compatibleWithAirPlayVideo"))
	return rv
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isComposable
func (a_ Asset) Composable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("composable"))
	return rv
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isExportable
func (a_ Asset) Exportable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("exportable"))
	return rv
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isPlayable
func (a_ Asset) Playable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playable"))
	return rv
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isReadable
func (a_ Asset) Readable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("readable"))
	return rv
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/lyrics
func (a_ Asset) Lyrics() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("lyrics"))
	return rv
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/metadata
func (a_ Asset) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/minimumTimeOffsetFromLive
func (a_ Asset) MinimumTimeOffsetFromLive() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/overallDurationHint
func (a_ Asset) OverallDurationHint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredMediaSelection
func (a_ Asset) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](a_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredRate
func (a_ Asset) PreferredRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredRate"))
	return rv
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredTransform
func (a_ Asset) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredVolume
func (a_ Asset) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/providesPreciseDurationAndTiming
func (a_ Asset) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/referenceRestrictions
func (a_ Asset) ReferenceRestrictions() AssetReferenceRestrictions {
	rv := objc.Send[AssetReferenceRestrictions](a_.ID, objc.Sel("referenceRestrictions"))
	return rv
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/trackGroups
func (a_ Asset) TrackGroups() []AssetTrackGroup {
	rv := objc.Send[[]AssetTrackGroup](a_.ID, objc.Sel("trackGroups"))
	return rv
}


// The tracks an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/tracks
func (a_ Asset) Tracks() []AssetTrack {
	rv := objc.Send[[]AssetTrack](a_.ID, objc.Sel("tracks"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}


// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) IsCompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}


// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithSavedPhotosAlbum:"), value)
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) IsComposable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isComposable"))
	return rv
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) SetIsComposable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsComposable:"), value)
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) IsExportable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isExportable"))
	return rv
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) SetIsExportable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsExportable:"), value)
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) IsPlayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) SetIsPlayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) IsReadable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadable"))
	return rv
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) SetIsReadable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadable:"), value)
}







