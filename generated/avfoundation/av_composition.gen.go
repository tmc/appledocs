// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale unsafe.Pointer, commonKeys unsafe.Pointer) []TimedMetadataGroup
	LoadTracksWithMediaTypeCompletionHandler(mediaType unsafe.Pointer, completionHandler unsafe.Pointer)
	MetadataForFormat(format unsafe.Pointer) []MetadataItem
}

// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
//
// A composition is a container for one or more tracks of media. Its tracks are instances of that present media of a uniform type like audio or video. A track itself is a container for one or more segments of media, which are instances of , a type that represents a region of media in the source track.
//
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


// Returns an array of chapters that contain the specified title locale and common keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (c_ Composition) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale unsafe.Pointer, commonKeys unsafe.Pointer) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](c_.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, commonKeys)
	return rv
}

// Loads tracks that contain media of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/loadTracks(withMediaType:completionHandler:)
func (c_ Composition) LoadTracksWithMediaTypeCompletionHandler(mediaType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}

// Returns an array of metadata items from the container with the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata(forFormat:)
func (c_ Composition) MetadataForFormat(format unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}

// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/commonMetadata
func (c_ Composition) CommonMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("commonMetadata"))
	return rv
}


// SetCommonMetadata sets the value of the commonMetadata property.
// The metadata items an asset contains for common metadata identifiers that provide a value.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/commonMetadata
func (c_ Composition) SetCommonMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}

// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata
func (c_ Composition) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// An array of metadata items for all metadata identifiers for which a value is available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata
func (c_ Composition) SetMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}

// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/overallDurationHint
func (c_ Composition) OverallDurationHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// SetOverallDurationHint sets the value of the overallDurationHint property.
// The total duration of fragments that currently exist, or may exist in the future.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/overallDurationHint
func (c_ Composition) SetOverallDurationHint(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOverallDurationHint:"), value)
}

// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredDisplayCriteria
func (c_ Composition) PreferredDisplayCriteria() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}


// SetPreferredDisplayCriteria sets the value of the preferredDisplayCriteria property.
// The asset’s display mode preference for optimal playback of its content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredDisplayCriteria
func (c_ Composition) SetPreferredDisplayCriteria(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredDisplayCriteria:"), value)
}

// The track groups an asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/trackGroups
func (c_ Composition) TrackGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("trackGroups"))
	return rv
}


// SetTrackGroups sets the value of the trackGroups property.
// The track groups an asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/trackGroups
func (c_ Composition) SetTrackGroups(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrackGroups:"), value)
}

// The tracks that a composition contains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks
func (c_ Composition) Tracks() []CompositionTrack {
	rv := objc.Send[[]CompositionTrack](c_.ID, objc.Sel("tracks"))
	return rv
}

// The array of available media selections for this asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/allmediaselections
func (c_ Composition) AllMediaSelections() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allMediaSelections"))
	return rv
}


// SetAllMediaSelections sets the value of the allMediaSelections property.
// The array of available media selections for this asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/allmediaselections
func (c_ Composition) SetAllMediaSelections(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllMediaSelections:"), value)
}

// The locales of the asset’s chapter metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablechapterlocales
func (c_ Composition) AvailableChapterLocales() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableChapterLocales"))
	return rv
}


// SetAvailableChapterLocales sets the value of the availableChapterLocales property.
// The locales of the asset’s chapter metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablechapterlocales
func (c_ Composition) SetAvailableChapterLocales(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}

// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemediacharacteristicswithmediaselectionoptions
func (c_ Composition) AvailableMediaCharacteristicsWithMediaSelectionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}


// SetAvailableMediaCharacteristicsWithMediaSelectionOptions sets the value of the availableMediaCharacteristicsWithMediaSelectionOptions property.
// An array of media characteristics for which a media selection option is available.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemediacharacteristicswithmediaselectionoptions
func (c_ Composition) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}

// The metadata formats this asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemetadataformats
func (c_ Composition) AvailableMetadataFormats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// SetAvailableMetadataFormats sets the value of the availableMetadataFormats property.
// The metadata formats this asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/availablemetadataformats
func (c_ Composition) SetAvailableMetadataFormats(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}

// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/cancontainfragments
func (c_ Composition) CanContainFragments() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canContainFragments"))
	return rv
}


// SetCanContainFragments sets the value of the canContainFragments property.
// A Boolean value that indicates whether you can extend the asset by fragments.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/cancontainfragments
func (c_ Composition) SetCanContainFragments(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanContainFragments:"), value)
}

// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/containsfragments
func (c_ Composition) ContainsFragments() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsFragments"))
	return rv
}


// SetContainsFragments sets the value of the containsFragments property.
// A Boolean value that indicates whether at least one movie fragment extends the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/containsfragments
func (c_ Composition) SetContainsFragments(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainsFragments:"), value)
}

// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/creationdate
func (c_ Composition) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// A metadata item that indicates the asset’s creation date.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/creationdate
func (c_ Composition) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreationDate:"), value)
}

// A time value that indicates the asset’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/duration
func (c_ Composition) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// A time value that indicates the asset’s duration.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/duration
func (c_ Composition) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}

// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/hasprotectedcontent
func (c_ Composition) HasProtectedContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasProtectedContent"))
	return rv
}


// SetHasProtectedContent sets the value of the hasProtectedContent property.
// A Boolean value that indicates whether the asset contains protected content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/hasprotectedcontent
func (c_ Composition) SetHasProtectedContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasProtectedContent:"), value)
}

// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithairplayvideo
func (c_ Composition) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}


// SetIsCompatibleWithAirPlayVideo sets the value of the isCompatibleWithAirPlayVideo property.
// A Boolean value that indicates whether the asset is compatible with AirPlay Video.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithairplayvideo
func (c_ Composition) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}

// A Boolean value that indicates whether you can write the composition to the Saved Photos album.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithsavedphotosalbum
func (c_ Composition) IsCompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}


// SetIsCompatibleWithSavedPhotosAlbum sets the value of the isCompatibleWithSavedPhotosAlbum property.
// A Boolean value that indicates whether you can write the composition to the Saved Photos album.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscompatiblewithsavedphotosalbum
func (c_ Composition) SetIsCompatibleWithSavedPhotosAlbum(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCompatibleWithSavedPhotosAlbum:"), value)
}

// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscomposable
func (c_ Composition) IsComposable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComposable"))
	return rv
}


// SetIsComposable sets the value of the isComposable property.
// A Boolean value that indicates whether you can use the asset as a segment of a composition track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/iscomposable
func (c_ Composition) SetIsComposable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComposable:"), value)
}

// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isexportable
func (c_ Composition) IsExportable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExportable"))
	return rv
}


// SetIsExportable sets the value of the isExportable property.
// A Boolean value that indicates whether you can export this asset using an export session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isexportable
func (c_ Composition) SetIsExportable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExportable:"), value)
}

// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isplayable
func (c_ Composition) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// SetIsPlayable sets the value of the isPlayable property.
// A Boolean value that indicates whether the asset has playable content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isplayable
func (c_ Composition) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}

// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isreadable
func (c_ Composition) IsReadable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isReadable"))
	return rv
}


// SetIsReadable sets the value of the isReadable property.
// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/isreadable
func (c_ Composition) SetIsReadable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsReadable:"), value)
}

// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/lyrics
func (c_ Composition) Lyrics() string {
	rv := objc.Send[string](c_.ID, objc.Sel("lyrics"))
	return rv
}


// SetLyrics sets the value of the lyrics property.
// The lyrics of the asset in a language suitable for the current locale.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/lyrics
func (c_ Composition) SetLyrics(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyrics:"), objc.String(value))
}

// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/minimumtimeoffsetfromlive
func (c_ Composition) MinimumTimeOffsetFromLive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}


// SetMinimumTimeOffsetFromLive sets the value of the minimumTimeOffsetFromLive property.
// A time value that indicates how closely playback follows the latest live stream content.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/minimumtimeoffsetfromlive
func (c_ Composition) SetMinimumTimeOffsetFromLive(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}

// The authored size of the visual portion of the composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/naturalsize
func (c_ Composition) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("naturalSize"))
	return rv
}


// SetNaturalSize sets the value of the naturalSize property.
// The authored size of the visual portion of the composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/naturalsize
func (c_ Composition) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNaturalSize:"), value)
}

// The default media selections for this asset’s media selection groups.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredmediaselection
func (c_ Composition) PreferredMediaSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}


// SetPreferredMediaSelection sets the value of the preferredMediaSelection property.
// The default media selections for this asset’s media selection groups.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredmediaselection
func (c_ Composition) SetPreferredMediaSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredMediaSelection:"), value)
}

// The asset’s rate preference for playing its media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredrate
func (c_ Composition) PreferredRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredRate"))
	return rv
}


// SetPreferredRate sets the value of the preferredRate property.
// The asset’s rate preference for playing its media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredrate
func (c_ Composition) SetPreferredRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredRate:"), value)
}

// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredtransform
func (c_ Composition) PreferredTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}


// SetPreferredTransform sets the value of the preferredTransform property.
// The asset’s transform preference to apply to its visual content during presentation or processing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredtransform
func (c_ Composition) SetPreferredTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}

// The asset’s volume preference for playing its audible media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredvolume
func (c_ Composition) PreferredVolume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredVolume"))
	return rv
}


// SetPreferredVolume sets the value of the preferredVolume property.
// The asset’s volume preference for playing its audible media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/preferredvolume
func (c_ Composition) SetPreferredVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}

// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/providesprecisedurationandtiming
func (c_ Composition) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}


// SetProvidesPreciseDurationAndTiming sets the value of the providesPreciseDurationAndTiming property.
// A Boolean value that indicates whether the asset provides precise duration and timing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/providesprecisedurationandtiming
func (c_ Composition) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}

// The options you used to create a composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/urlassetinitializationoptions
func (c_ Composition) UrlAssetInitializationOptions() string {
	rv := objc.Send[string](c_.ID, objc.Sel("urlAssetInitializationOptions"))
	return rv
}


// SetUrlAssetInitializationOptions sets the value of the urlAssetInitializationOptions property.
// The options you used to create a composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/urlassetinitializationoptions
func (c_ Composition) SetUrlAssetInitializationOptions(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrlAssetInitializationOptions:"), objc.String(value))
}



