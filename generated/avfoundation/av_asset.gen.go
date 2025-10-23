// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
	AllMediaSelections() IAVMediaSelection
	SetAllMediaSelections(value IAVMediaSelection)
	AvailableChapterLocales() foundation.Locale
	SetAvailableChapterLocales(value foundation.Locale)
	AvailableMediaCharacteristicsWithMediaSelectionOptions() unsafe.Pointer
	SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value unsafe.Pointer)
	AvailableMetadataFormats() unsafe.Pointer
	SetAvailableMetadataFormats(value unsafe.Pointer)
	CanContainFragments() bool
	SetCanContainFragments(value bool)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	ContainsFragments() bool
	SetContainsFragments(value bool)
	CreationDate() IAVMetadataItem
	SetCreationDate(value IAVMetadataItem)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	HasProtectedContent() bool
	SetHasProtectedContent(value bool)
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
	Lyrics() string
	SetLyrics(value string)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinimumTimeOffsetFromLive() unsafe.Pointer
	SetMinimumTimeOffsetFromLive(value unsafe.Pointer)
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	OverallDurationHint() unsafe.Pointer
	SetOverallDurationHint(value unsafe.Pointer)
	PreferredDisplayCriteria() DisplayCriteria
	SetPreferredDisplayCriteria(value DisplayCriteria)
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)
	PreferredRate() float32
	SetPreferredRate(value float32)
	PreferredTransform() coregraphics.CGAffineTransform
	SetPreferredTransform(value coregraphics.CGAffineTransform)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	ProvidesPreciseDurationAndTiming() bool
	SetProvidesPreciseDurationAndTiming(value bool)
	ReferenceRestrictions() unsafe.Pointer
	SetReferenceRestrictions(value unsafe.Pointer)
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
	Tracks() IAVAssetTrack
	SetTracks(value IAVAssetTrack)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer)
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



// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a_ Asset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/allmediaselections
func (a_ Asset) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[MediaSelection](a_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/allmediaselections
func (a_ Asset) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllMediaSelections:"), value)
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablechapterlocales
func (a_ Asset) AvailableChapterLocales() foundation.Locale {
	rv := objc.Send[foundation.Locale](a_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablechapterlocales
func (a_ Asset) SetAvailableChapterLocales(value foundation.Locale) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) AvailableMediaCharacteristicsWithMediaSelectionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) AvailableMetadataFormats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) SetAvailableMetadataFormats(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (a_ Asset) CanContainFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canContainFragments"))
	return rv
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (a_ Asset) SetCanContainFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanContainFragments:"), value)
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (a_ Asset) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("commonMetadata"))
	return rv
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (a_ Asset) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCommonMetadata:"), value)
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/containsfragments
func (a_ Asset) ContainsFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsFragments"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/containsfragments
func (a_ Asset) SetContainsFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContainsFragments:"), value)
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/creationdate
func (a_ Asset) CreationDate() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("creationDate"))
	return rv
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/creationdate
func (a_ Asset) SetCreationDate(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCreationDate:"), value)
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/duration
func (a_ Asset) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("duration"))
	return rv
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/duration
func (a_ Asset) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/hasprotectedcontent
func (a_ Asset) HasProtectedContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/hasprotectedcontent
func (a_ Asset) SetHasProtectedContent(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasProtectedContent:"), value)
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


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) Lyrics() string {
	rv := objc.Send[string](a_.ID, objc.Sel("lyrics"))
	return rv
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) SetLyrics(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLyrics:"), objc.String(value))
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/metadata
func (a_ Asset) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/metadata
func (a_ Asset) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/minimumtimeoffsetfromlive
func (a_ Asset) MinimumTimeOffsetFromLive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/minimumtimeoffsetfromlive
func (a_ Asset) SetMinimumTimeOffsetFromLive(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/naturalsize
func (a_ Asset) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/naturalsize
func (a_ Asset) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) OverallDurationHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) SetOverallDurationHint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferreddisplaycriteria
func (a_ Asset) PreferredDisplayCriteria() DisplayCriteria {
	rv := objc.Send[DisplayCriteria](a_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferreddisplaycriteria
func (a_ Asset) SetPreferredDisplayCriteria(value DisplayCriteria) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredDisplayCriteria:"), value)
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredmediaselection
func (a_ Asset) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](a_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredmediaselection
func (a_ Asset) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaSelection:"), value)
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredrate
func (a_ Asset) PreferredRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredRate"))
	return rv
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredrate
func (a_ Asset) SetPreferredRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredRate:"), value)
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredvolume
func (a_ Asset) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredvolume
func (a_ Asset) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}


// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/referencerestrictions
func (a_ Asset) ReferenceRestrictions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("referenceRestrictions"))
	return rv
}


// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/referencerestrictions
func (a_ Asset) SetReferenceRestrictions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReferenceRestrictions:"), value)
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/trackgroups
func (a_ Asset) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](a_.ID, objc.Sel("trackGroups"))
	return rv
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/trackgroups
func (a_ Asset) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackGroups:"), value)
}


// The tracks an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/tracks
func (a_ Asset) Tracks() IAVAssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("tracks"))
	return rv
}


// The tracks an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/tracks
func (a_ Asset) SetTracks(value IAVAssetTrack) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTracks:"), value)
}



