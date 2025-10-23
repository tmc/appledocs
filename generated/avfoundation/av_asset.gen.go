// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	CanContainFragments() bool /* primitive/slice/pointer. */
	ContainsFragments() bool /* primitive/slice/pointer. */
	Duration() Time /* not a class type */
	AllMediaSelections() IAVMediaSelection
	SetAllMediaSelections(value IAVMediaSelection)
	AvailableChapterLocales() objc.IObject /* cross-framework: Locale */
	SetAvailableChapterLocales(value objc.IObject /* cross-framework: Locale */)
	AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic /* not a class type */
	SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic /* not a class type */)
	AvailableMetadataFormats() MetadataFormat /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat /* not a class type */)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	CreationDate() IAVMetadataItem
	SetCreationDate(value IAVMetadataItem)
	HasProtectedContent() bool /* primitive/slice/pointer. */
	SetHasProtectedContent(value bool /* primitive/slice/pointer. */)
	IsCompatibleWithAirPlayVideo() bool /* primitive/slice/pointer. */
	SetIsCompatibleWithAirPlayVideo(value bool /* primitive/slice/pointer. */)
	IsCompatibleWithSavedPhotosAlbum() bool /* primitive/slice/pointer. */
	SetIsCompatibleWithSavedPhotosAlbum(value bool /* primitive/slice/pointer. */)
	IsComposable() bool /* primitive/slice/pointer. */
	SetIsComposable(value bool /* primitive/slice/pointer. */)
	IsExportable() bool /* primitive/slice/pointer. */
	SetIsExportable(value bool /* primitive/slice/pointer. */)
	IsPlayable() bool /* primitive/slice/pointer. */
	SetIsPlayable(value bool /* primitive/slice/pointer. */)
	IsReadable() bool /* primitive/slice/pointer. */
	SetIsReadable(value bool /* primitive/slice/pointer. */)
	Lyrics() objc.IObject /* cross-framework: NSString */
	SetLyrics(value objc.IObject /* cross-framework: NSString */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinimumTimeOffsetFromLive() Time /* not a class type */
	SetMinimumTimeOffsetFromLive(value Time /* not a class type */)
	NaturalSize() objc.IObject /* cross-framework: Size */
	SetNaturalSize(value objc.IObject /* cross-framework: Size */)
	OverallDurationHint() Time /* not a class type */
	SetOverallDurationHint(value Time /* not a class type */)
	PreferredDisplayCriteria() objc.IObject /* cross-framework: DisplayCriteria */
	SetPreferredDisplayCriteria(value objc.IObject /* cross-framework: DisplayCriteria */)
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)
	PreferredRate() float32 /* primitive/slice/pointer. */
	SetPreferredRate(value float32 /* primitive/slice/pointer. */)
	PreferredTransform() objc.IObject /* cross-framework: AffineTransform */
	SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */)
	PreferredVolume() float32 /* primitive/slice/pointer. */
	SetPreferredVolume(value float32 /* primitive/slice/pointer. */)
	ProvidesPreciseDurationAndTiming() bool /* primitive/slice/pointer. */
	SetProvidesPreciseDurationAndTiming(value bool /* primitive/slice/pointer. */)
	ReferenceRestrictions() AssetReferenceRestrictions /* not a class type */
	SetReferenceRestrictions(value AssetReferenceRestrictions /* not a class type */)
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
	Tracks() IAVAssetTrack
	SetTracks(value IAVAssetTrack)
	// methods:
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* not a class type */, completionHandler unsafe.Pointer)
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
func (a_ Asset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/canContainFragments
func (a_ Asset) CanContainFragments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canContainFragments"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/containsFragments
func (a_ Asset) ContainsFragments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsFragments"))
	return rv
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/duration
func (a_ Asset) Duration() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("duration"))
	return rv
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
func (a_ Asset) AvailableChapterLocales() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](a_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablechapterlocales
func (a_ Asset) SetAvailableChapterLocales(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic /* not a class type */ {
	rv := objc.Send[MediaCharacteristic](a_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemediacharacteristicswithmediaselectionoptions
func (a_ Asset) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) AvailableMetadataFormats() MetadataFormat /* not a class type */ {
	rv := objc.Send[MetadataFormat](a_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/availablemetadataformats
func (a_ Asset) SetAvailableMetadataFormats(value MetadataFormat /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
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


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/hasprotectedcontent
func (a_ Asset) HasProtectedContent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/hasprotectedcontent
func (a_ Asset) SetHasProtectedContent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasProtectedContent:"), value)
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) IsCompatibleWithAirPlayVideo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithairplayvideo
func (a_ Asset) SetIsCompatibleWithAirPlayVideo(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}


// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) IsCompatibleWithSavedPhotosAlbum() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}


// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscompatiblewithsavedphotosalbum
func (a_ Asset) SetIsCompatibleWithSavedPhotosAlbum(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompatibleWithSavedPhotosAlbum:"), value)
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) IsComposable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isComposable"))
	return rv
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/iscomposable
func (a_ Asset) SetIsComposable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsComposable:"), value)
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) IsExportable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isExportable"))
	return rv
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isexportable
func (a_ Asset) SetIsExportable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsExportable:"), value)
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) IsPlayable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isplayable
func (a_ Asset) SetIsPlayable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) IsReadable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadable"))
	return rv
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/isreadable
func (a_ Asset) SetIsReadable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadable:"), value)
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) Lyrics() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("lyrics"))
	return rv
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/lyrics
func (a_ Asset) SetLyrics(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLyrics:"), value)
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
func (a_ Asset) MinimumTimeOffsetFromLive() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/minimumtimeoffsetfromlive
func (a_ Asset) SetMinimumTimeOffsetFromLive(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/naturalsize
func (a_ Asset) NaturalSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/naturalsize
func (a_ Asset) SetNaturalSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) OverallDurationHint() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/overalldurationhint
func (a_ Asset) SetOverallDurationHint(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferreddisplaycriteria
func (a_ Asset) PreferredDisplayCriteria() objc.IObject /* cross-framework: DisplayCriteria */ {
	rv := objc.Send[DisplayCriteria](a_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferreddisplaycriteria
func (a_ Asset) SetPreferredDisplayCriteria(value objc.IObject /* cross-framework: DisplayCriteria */) {
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
func (a_ Asset) PreferredRate() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredRate"))
	return rv
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredrate
func (a_ Asset) SetPreferredRate(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredRate:"), value)
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) PreferredTransform() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[AffineTransform](a_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredtransform
func (a_ Asset) SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredvolume
func (a_ Asset) PreferredVolume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/preferredvolume
func (a_ Asset) SetPreferredVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) ProvidesPreciseDurationAndTiming() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/providesprecisedurationandtiming
func (a_ Asset) SetProvidesPreciseDurationAndTiming(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}


// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/referencerestrictions
func (a_ Asset) ReferenceRestrictions() AssetReferenceRestrictions /* not a class type */ {
	rv := objc.Send[AssetReferenceRestrictions](a_.ID, objc.Sel("referenceRestrictions"))
	return rv
}


// The restrictions that an asset places on how it resolves references to external media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/referencerestrictions
func (a_ Asset) SetReferenceRestrictions(value AssetReferenceRestrictions /* not a class type */) {
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



