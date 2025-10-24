// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AVComposition */


/* debug [class_header]: Header for AVComposition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Composition */
// An interface definition for the [Composition] class.
type IComposition interface {
	IAsset
	
/* debug [class_interface_properties]: Properties for Composition */
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
	Duration() Time get /* not a class type */
	SetDuration(value Time get /* not a class type */)
	HasProtectedContent() objectivec.IObject
	SetHasProtectedContent(value objectivec.IObject)
	IsCompatibleWithAirPlayVideo() objectivec.IObject
	SetIsCompatibleWithAirPlayVideo(value objectivec.IObject)
	IsComposable() objectivec.IObject
	SetIsComposable(value objectivec.IObject)
	IsExportable() objectivec.IObject
	SetIsExportable(value objectivec.IObject)
	IsPlayable() objectivec.IObject
	SetIsPlayable(value objectivec.IObject)
	IsReadable() objectivec.IObject
	SetIsReadable(value objectivec.IObject)
	Lyrics() objectivec.IObject
	SetLyrics(value objectivec.IObject)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MinimumTimeOffsetFromLive() Time get /* not a class type */
	SetMinimumTimeOffsetFromLive(value Time get /* not a class type */)
	NaturalSize() corefoundation.CGSize
	OverallDurationHint() Time get /* not a class type */
	SetOverallDurationHint(value Time get /* not a class type */)
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)
	PreferredRate() objectivec.IObject
	SetPreferredRate(value objectivec.IObject)
	PreferredTransform() AffineTransform get /* not a class type */
	SetPreferredTransform(value AffineTransform get /* not a class type */)
	PreferredVolume() objectivec.IObject
	SetPreferredVolume(value objectivec.IObject)
	ProvidesPreciseDurationAndTiming() objectivec.IObject
	SetProvidesPreciseDurationAndTiming(value objectivec.IObject)
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
	Tracks() []CompositionTrack
	URLAssetInitializationOptions() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Composition */
	// methods:
	ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []TimedMetadataGroup
	ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.Locale, commonKeys []string) []TimedMetadataGroup
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer)
	MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IMediaSelectionGroup
	MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem
	TrackWithTrackID(trackID PersistentTrackID /* not a class type */) ICompositionTrack
	TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) []CompositionTrack
	TracksWithMediaType(mediaType MediaType /* typedef */) []CompositionTrack
	UnusedTrackID() PersistentTrackID /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Composition */
// Alloc allocates a new instance without initialization.
func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Composition */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Composition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Composition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Composition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Composition */

// Returns an array of chapters with a locale that best matches the list of preferred languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(bestMatchingPreferredLanguages:)
func (c_ Composition) ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](c_.ID, objc.Sel("chapterMetadataGroupsBestMatchingPreferredLanguages:"), preferredLanguages)
	return rv
}/* debug [instance_methods/method]: ChapterMetadataGroupsBestMatchingPreferredLanguages */


// Returns an array of chapters that contain the specified title locale and common keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (c_ Composition) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.Locale, commonKeys []string) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](c_.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, commonKeys)
	return rv
}/* debug [instance_methods/method]: ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys */


// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/loadTrack(withTrackID:completionHandler:)
func (c_ Composition) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}/* debug [instance_methods/method]: LoadTrackWithTrackIDCompletionHandler */


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/loadTracks(withMediaCharacteristic:completionHandler:)
func (c_ Composition) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaCharacteristicCompletionHandler */


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/loadTracks(withMediaType:completionHandler:)
func (c_ Composition) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaTypeCompletionHandler */


// Returns a media selection group that contains one or more options with the specified media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/mediaSelectionGroup(forMediaCharacteristic:)
func (c_ Composition) MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) IMediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](c_.ID, objc.Sel("mediaSelectionGroupForMediaCharacteristic:"), mediaCharacteristic)
	return rv
}/* debug [instance_methods/method]: MediaSelectionGroupForMediaCharacteristic */


// Returns an array of metadata items from the container with the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata(forFormat:)
func (c_ Composition) MetadataForFormat(format MetadataFormat /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}/* debug [instance_methods/method]: MetadataForFormat */


// Returns a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/track(withTrackID:)
func (c_ Composition) TrackWithTrackID(trackID PersistentTrackID /* not a class type */) ICompositionTrack {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("trackWithTrackID:"), trackID)
	return rv
}/* debug [instance_methods/method]: TrackWithTrackID */


// Returns tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks(withMediaCharacteristic:)
func (c_ Composition) TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) []CompositionTrack {
	rv := objc.Send[[]CompositionTrack](c_.ID, objc.Sel("tracksWithMediaCharacteristic:"), mediaCharacteristic)
	return rv
}/* debug [instance_methods/method]: TracksWithMediaCharacteristic */


// Returns tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks(withMediaType:)
func (c_ Composition) TracksWithMediaType(mediaType MediaType /* typedef */) []CompositionTrack {
	rv := objc.Send[[]CompositionTrack](c_.ID, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}/* debug [instance_methods/method]: TracksWithMediaType */


// Returns an identifier that no other tracks in the asset use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/unusedTrackID()
func (c_ Composition) UnusedTrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](c_.ID, objc.Sel("unusedTrackID"))
	return rv
}/* debug [instance_methods/method]: UnusedTrackID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Composition */

// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/allMediaSelections
func (c_ Composition) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[MediaSelection](c_.ID, objc.Sel("allMediaSelections"))
	return rv
}/* debug [instance_properties/getter]: allMediaSelections */


// The array of available media selections for this asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/allMediaSelections
func (c_ Composition) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllMediaSelections:"), value)
}/* debug [instance_properties/setter]: allMediaSelections */


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableChapterLocales
func (c_ Composition) AvailableChapterLocales() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("availableChapterLocales"))
	return rv
}/* debug [instance_properties/getter]: availableChapterLocales */


// The locales of the asset’s chapter metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableChapterLocales
func (c_ Composition) SetAvailableChapterLocales(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableChapterLocales:"), value)
}/* debug [instance_properties/setter]: availableChapterLocales */


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMediaCharacteristicsWithMediaSelectionOptions
func (c_ Composition) AvailableMediaCharacteristicsWithMediaSelectionOptions() MediaCharacteristic get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}/* debug [instance_properties/getter]: availableMediaCharacteristicsWithMediaSelectionOptions */


// An array of media characteristics for which a media selection option is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMediaCharacteristicsWithMediaSelectionOptions
func (c_ Composition) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value MediaCharacteristic get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), value)
}/* debug [instance_properties/setter]: availableMediaCharacteristicsWithMediaSelectionOptions */


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMetadataFormats
func (c_ Composition) AvailableMetadataFormats() MetadataFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}/* debug [instance_properties/getter]: availableMetadataFormats */


// The metadata formats this asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMetadataFormats
func (c_ Composition) SetAvailableMetadataFormats(value MetadataFormat get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}/* debug [instance_properties/setter]: availableMetadataFormats */


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/canContainFragments
func (c_ Composition) CanContainFragments() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("canContainFragments"))
	return rv
}/* debug [instance_properties/getter]: canContainFragments */


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/canContainFragments
func (c_ Composition) SetCanContainFragments(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanContainFragments:"), value)
}/* debug [instance_properties/setter]: canContainFragments */


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/commonMetadata
func (c_ Composition) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("commonMetadata"))
	return rv
}/* debug [instance_properties/getter]: commonMetadata */


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/commonMetadata
func (c_ Composition) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommonMetadata:"), value)
}/* debug [instance_properties/setter]: commonMetadata */


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/containsFragments
func (c_ Composition) ContainsFragments() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("containsFragments"))
	return rv
}/* debug [instance_properties/getter]: containsFragments */


// A Boolean value that indicates whether at least one movie fragment extends the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/containsFragments
func (c_ Composition) SetContainsFragments(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainsFragments:"), value)
}/* debug [instance_properties/setter]: containsFragments */


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/creationDate
func (c_ Composition) CreationDate() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// A metadata item that indicates the asset’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/creationDate
func (c_ Composition) SetCreationDate(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreationDate:"), value)
}/* debug [instance_properties/setter]: creationDate */


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/duration
func (c_ Composition) Duration() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A time value that indicates the asset’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/duration
func (c_ Composition) SetDuration(value Time get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/hasProtectedContent
func (c_ Composition) HasProtectedContent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("hasProtectedContent"))
	return rv
}/* debug [instance_properties/getter]: hasProtectedContent */


// A Boolean value that indicates whether the asset contains protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/hasProtectedContent
func (c_ Composition) SetHasProtectedContent(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasProtectedContent:"), value)
}/* debug [instance_properties/setter]: hasProtectedContent */


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isCompatibleWithAirPlayVideo
func (c_ Composition) IsCompatibleWithAirPlayVideo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}/* debug [instance_properties/getter]: isCompatibleWithAirPlayVideo */


// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isCompatibleWithAirPlayVideo
func (c_ Composition) SetIsCompatibleWithAirPlayVideo(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCompatibleWithAirPlayVideo:"), value)
}/* debug [instance_properties/setter]: isCompatibleWithAirPlayVideo */


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isComposable
func (c_ Composition) IsComposable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isComposable"))
	return rv
}/* debug [instance_properties/getter]: isComposable */


// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isComposable
func (c_ Composition) SetIsComposable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComposable:"), value)
}/* debug [instance_properties/setter]: isComposable */


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isExportable
func (c_ Composition) IsExportable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isExportable"))
	return rv
}/* debug [instance_properties/getter]: isExportable */


// A Boolean value that indicates whether you can export this asset using an export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isExportable
func (c_ Composition) SetIsExportable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExportable:"), value)
}/* debug [instance_properties/setter]: isExportable */


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isPlayable
func (c_ Composition) IsPlayable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isPlayable"))
	return rv
}/* debug [instance_properties/getter]: isPlayable */


// A Boolean value that indicates whether the asset has playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isPlayable
func (c_ Composition) SetIsPlayable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}/* debug [instance_properties/setter]: isPlayable */


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isReadable
func (c_ Composition) IsReadable() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isReadable"))
	return rv
}/* debug [instance_properties/getter]: isReadable */


// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isReadable
func (c_ Composition) SetIsReadable(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsReadable:"), value)
}/* debug [instance_properties/setter]: isReadable */


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/lyrics
func (c_ Composition) Lyrics() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("lyrics"))
	return rv
}/* debug [instance_properties/getter]: lyrics */


// The lyrics of the asset in a language suitable for the current locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/lyrics
func (c_ Composition) SetLyrics(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyrics:"), value)
}/* debug [instance_properties/setter]: lyrics */


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata
func (c_ Composition) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata
func (c_ Composition) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}/* debug [instance_properties/setter]: metadata */


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/minimumTimeOffsetFromLive
func (c_ Composition) MinimumTimeOffsetFromLive() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}/* debug [instance_properties/getter]: minimumTimeOffsetFromLive */


// A time value that indicates how closely playback follows the latest live stream content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/minimumTimeOffsetFromLive
func (c_ Composition) SetMinimumTimeOffsetFromLive(value Time get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}/* debug [instance_properties/setter]: minimumTimeOffsetFromLive */


// The authored size of the visual portion of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/naturalSize
func (c_ Composition) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("naturalSize"))
	return rv
}/* debug [instance_properties/getter]: naturalSize */


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/overallDurationHint
func (c_ Composition) OverallDurationHint() Time get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("overallDurationHint"))
	return rv
}/* debug [instance_properties/getter]: overallDurationHint */


// The total duration of fragments that currently exist, or may exist in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/overallDurationHint
func (c_ Composition) SetOverallDurationHint(value Time get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOverallDurationHint:"), value)
}/* debug [instance_properties/setter]: overallDurationHint */


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredMediaSelection
func (c_ Composition) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[MediaSelection](c_.ID, objc.Sel("preferredMediaSelection"))
	return rv
}/* debug [instance_properties/getter]: preferredMediaSelection */


// The default media selections for this asset’s media selection groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredMediaSelection
func (c_ Composition) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredMediaSelection:"), value)
}/* debug [instance_properties/setter]: preferredMediaSelection */


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredRate
func (c_ Composition) PreferredRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("preferredRate"))
	return rv
}/* debug [instance_properties/getter]: preferredRate */


// The asset’s rate preference for playing its media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredRate
func (c_ Composition) SetPreferredRate(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredRate:"), value)
}/* debug [instance_properties/setter]: preferredRate */


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredTransform
func (c_ Composition) PreferredTransform() AffineTransform get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The asset’s transform preference to apply to its visual content during presentation or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredTransform
func (c_ Composition) SetPreferredTransform(value AffineTransform get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredTransform:"), value)
}/* debug [instance_properties/setter]: preferredTransform */


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredVolume
func (c_ Composition) PreferredVolume() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("preferredVolume"))
	return rv
}/* debug [instance_properties/getter]: preferredVolume */


// The asset’s volume preference for playing its audible media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredVolume
func (c_ Composition) SetPreferredVolume(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVolume:"), value)
}/* debug [instance_properties/setter]: preferredVolume */


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/providesPreciseDurationAndTiming
func (c_ Composition) ProvidesPreciseDurationAndTiming() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}/* debug [instance_properties/getter]: providesPreciseDurationAndTiming */


// A Boolean value that indicates whether the asset provides precise duration and timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/providesPreciseDurationAndTiming
func (c_ Composition) SetProvidesPreciseDurationAndTiming(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}/* debug [instance_properties/setter]: providesPreciseDurationAndTiming */


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/trackGroups
func (c_ Composition) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](c_.ID, objc.Sel("trackGroups"))
	return rv
}/* debug [instance_properties/getter]: trackGroups */


// The track groups an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/trackGroups
func (c_ Composition) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrackGroups:"), value)
}/* debug [instance_properties/setter]: trackGroups */


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks
func (c_ Composition) Tracks() []CompositionTrack {
	rv := objc.Send[[]CompositionTrack](c_.ID, objc.Sel("tracks"))
	return rv
}/* debug [instance_properties/getter]: tracks */


// The options you used to create a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/urlAssetInitializationOptions
func (c_ Composition) URLAssetInitializationOptions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("URLAssetInitializationOptions"))
	return rv
}/* debug [instance_properties/getter]: URLAssetInitializationOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVComposition */


