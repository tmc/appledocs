// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Composition] class.
var (
	CompositionClass     _CompositionClass
	CompositionClassOnce sync.Once
)

func getCompositionClass() _CompositionClass {
	CompositionClassOnce.Do(func() {
		CompositionClass = _CompositionClass{objc.GetClass("AVComposition")}
	})
	return CompositionClass
}

type _CompositionClass struct {
	class objc.Class
}

// An interface definition for the [Composition] class.
type IComposition interface {
	IAsset
	// properties:
	AllMediaSelections() IAVMediaSelection
	SetAllMediaSelections(value IAVMediaSelection)
	AvailableChapterLocales() objc.IObject /* cross-framework: Locale */
	SetAvailableChapterLocales(value objc.IObject /* cross-framework: Locale */)
	AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic /* not a class type */
	SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic /* not a class type */)
	AvailableMetadataFormats() MetadataFormat /* not a class type */
	SetAvailableMetadataFormats(value MetadataFormat /* not a class type */)
	CanContainFragments() bool
	SetCanContainFragments(value bool)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	ContainsFragments() bool
	SetContainsFragments(value bool)
	CreationDate() IAVMetadataItem
	SetCreationDate(value IAVMetadataItem)
	Duration() objc.IObject /* cross-framework: Time */
	SetDuration(value objc.IObject /* cross-framework: Time */)
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
	Lyrics() objc.IObject /* cross-framework: NSString */
	SetLyrics(value objc.IObject /* cross-framework: NSString */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinimumTimeOffsetFromLive() objc.IObject /* cross-framework: Time */
	SetMinimumTimeOffsetFromLive(value objc.IObject /* cross-framework: Time */)
	NaturalSize() objc.IObject /* cross-framework: Size */
	SetNaturalSize(value objc.IObject /* cross-framework: Size */)
	OverallDurationHint() objc.IObject /* cross-framework: Time */
	SetOverallDurationHint(value objc.IObject /* cross-framework: Time */)
	PreferredDisplayCriteria() objc.IObject /* cross-framework: DisplayCriteria */
	SetPreferredDisplayCriteria(value objc.IObject /* cross-framework: DisplayCriteria */)
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)
	PreferredRate() float32
	SetPreferredRate(value float32)
	PreferredTransform() objc.IObject /* cross-framework: AffineTransform */
	SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	ProvidesPreciseDurationAndTiming() bool
	SetProvidesPreciseDurationAndTiming(value bool)
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
	Tracks() objc.IObject /* cross-framework: CompositionTrack */
	SetTracks(value objc.IObject /* cross-framework: CompositionTrack */)
	UrlAssetInitializationOptions() objc.IObject /* cross-framework: NSString */
	SetUrlAssetInitializationOptions(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
//
// A composition is a container for one or more tracks of media. Its tracks are instances of that present media of a uniform type like audio or video. A track itself is a container for one or more segments of media, which are instances of , a type that represents a region of media in the source track.


// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition
type Composition struct {
	Asset
}

// CompositionFrom constructs a [Composition] from an unsafe.Pointer.
//
// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
func CompositionFrom(ptr unsafe.Pointer) Composition {
	return Composition{
		Asset: AssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositionClass) New() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Composition) Init() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Composition) Autorelease() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComposition creates a new Composition instance.
func NewComposition() Composition {
	return getCompositionClass().New()
}



// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/allmediaselections
func (c_ Composition) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[MediaSelection](c_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/allmediaselections
func (c_ Composition) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllMediaSelections:"), value)
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablechapterlocales
func (c_ Composition) AvailableChapterLocales() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](c_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablechapterlocales
func (c_ Composition) SetAvailableChapterLocales(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemediacharacteristicswithmediaselectionoptions
func (c_ Composition) AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic /* not a class type */ {
	rv := objc.Send[MediaCharacteristic](c_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemediacharacteristicswithmediaselectionoptions
func (c_ Composition) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemetadataformats
func (c_ Composition) AvailableMetadataFormats() MetadataFormat /* not a class type */ {
	rv := objc.Send[MetadataFormat](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemetadataformats
func (c_ Composition) SetAvailableMetadataFormats(value MetadataFormat /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/cancontainfragments
func (c_ Composition) CanContainFragments() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canContainFragments"))
	return rv
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/cancontainfragments
func (c_ Composition) SetCanContainFragments(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanContainFragments:"), value)
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/commonmetadata
func (c_ Composition) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/commonmetadata
func (c_ Composition) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/containsfragments
func (c_ Composition) ContainsFragments() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsFragments"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/containsfragments
func (c_ Composition) SetContainsFragments(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainsFragments:"), value)
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/creationdate
func (c_ Composition) CreationDate() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("creationDate"))
	return rv
}


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/creationdate
func (c_ Composition) SetCreationDate(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreationDate:"), value)
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/duration
func (c_ Composition) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](c_.ID, objc.Sel("duration"))
	return rv
}


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/duration
func (c_ Composition) SetDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/hasprotectedcontent
func (c_ Composition) HasProtectedContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/hasprotectedcontent
func (c_ Composition) SetHasProtectedContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasProtectedContent:"), value)
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithairplayvideo
func (c_ Composition) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithairplayvideo
func (c_ Composition) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}


// A Boolean value that indicates whether you can write the composition to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithsavedphotosalbum
func (c_ Composition) IsCompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}


// A Boolean value that indicates whether you can write the composition to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithsavedphotosalbum
func (c_ Composition) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCompatibleWithSavedPhotosAlbum:"), value)
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscomposable
func (c_ Composition) IsComposable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComposable"))
	return rv
}


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscomposable
func (c_ Composition) SetIsComposable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComposable:"), value)
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isexportable
func (c_ Composition) IsExportable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExportable"))
	return rv
}


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isexportable
func (c_ Composition) SetIsExportable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExportable:"), value)
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isplayable
func (c_ Composition) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isplayable
func (c_ Composition) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isreadable
func (c_ Composition) IsReadable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isReadable"))
	return rv
}


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isreadable
func (c_ Composition) SetIsReadable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsReadable:"), value)
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/lyrics
func (c_ Composition) Lyrics() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("lyrics"))
	return rv
}


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/lyrics
func (c_ Composition) SetLyrics(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyrics:"), value)
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/metadata
func (c_ Composition) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/metadata
func (c_ Composition) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/minimumtimeoffsetfromlive
func (c_ Composition) MinimumTimeOffsetFromLive() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](c_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/minimumtimeoffsetfromlive
func (c_ Composition) SetMinimumTimeOffsetFromLive(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}


// The authored size of the visual portion of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/naturalsize
func (c_ Composition) NaturalSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](c_.ID, objc.Sel("naturalSize"))
	return rv
}


// The authored size of the visual portion of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/naturalsize
func (c_ Composition) SetNaturalSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/overalldurationhint
func (c_ Composition) OverallDurationHint() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](c_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/overalldurationhint
func (c_ Composition) SetOverallDurationHint(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferreddisplaycriteria
func (c_ Composition) PreferredDisplayCriteria() objc.IObject /* cross-framework: DisplayCriteria */ {
	rv := objc.Send[DisplayCriteria](c_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}


// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferreddisplaycriteria
func (c_ Composition) SetPreferredDisplayCriteria(value objc.IObject /* cross-framework: DisplayCriteria */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredDisplayCriteria:"), value)
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredmediaselection
func (c_ Composition) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](c_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredmediaselection
func (c_ Composition) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredMediaSelection:"), value)
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredrate
func (c_ Composition) PreferredRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("preferredRate"))
	return rv
}


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredrate
func (c_ Composition) SetPreferredRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredRate:"), value)
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredtransform
func (c_ Composition) PreferredTransform() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[corefoundation.AffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredtransform
func (c_ Composition) SetPreferredTransform(value objc.IObject /* cross-framework: AffineTransform */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredvolume
func (c_ Composition) PreferredVolume() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredvolume
func (c_ Composition) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/providesprecisedurationandtiming
func (c_ Composition) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/providesprecisedurationandtiming
func (c_ Composition) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/trackgroups
func (c_ Composition) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](c_.ID, objc.Sel("trackGroups"))
	return rv
}


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/trackgroups
func (c_ Composition) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrackGroups:"), value)
}


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/tracks
func (c_ Composition) Tracks() objc.IObject /* cross-framework: CompositionTrack */ {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("tracks"))
	return rv
}


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/tracks
func (c_ Composition) SetTracks(value objc.IObject /* cross-framework: CompositionTrack */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTracks:"), value)
}


// The options you used to create a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/urlassetinitializationoptions
func (c_ Composition) UrlAssetInitializationOptions() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("urlAssetInitializationOptions"))
	return rv
}


// The options you used to create a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/urlassetinitializationoptions
func (c_ Composition) SetUrlAssetInitializationOptions(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrlAssetInitializationOptions:"), value)
}



