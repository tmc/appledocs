// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetCache */


/* debug [class_header]: Header for AVAssetCache */
// The class instance for the [AssetCache] class.
var (
	AssetCacheClass     _AssetCacheClass
	AssetCacheClassOnce sync.Once
)

func getAssetCacheClass() _AssetCacheClass {
	AssetCacheClassOnce.Do(func() {
		AssetCacheClass = _AssetCacheClass{objc.GetClass("AVAssetCache")}
	})
	return AssetCacheClass
}

type _AssetCacheClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetCache */
// An interface definition for the [AssetCache] class.
type IAssetCache interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetCache */
	// properties:
	PlayableOffline() bool
	IsPlayableOffline() bool
	SetIsPlayableOffline(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetCache */
	// methods:
	MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string
	MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary
	MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []MediaSelectionOption
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetCache */
// Alloc allocates a new instance without initialization.
func (ac _AssetCacheClass) Alloc() AssetCache {
	rv := objc.Send[AssetCache](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetCacheClass) New() AssetCache {
	rv := objc.Send[AssetCache](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetCache) Init() AssetCache {
	rv := objc.Send[AssetCache](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetCache) Autorelease() AssetCache {
	rv := objc.Send[AssetCache](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetCache creates a new AssetCache instance.
func NewAssetCache() AssetCache {
	return getAssetCacheClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetCache */
// An object that you use to inspect locally cached media data.
//
// You can download HTTP Live Streaming assets to an iOS device using the and classes.


// An object that you use to inspect locally cached media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache
type AssetCache struct {
	objectivec.Object
}

// AssetCacheFrom constructs a [AssetCache] from an unsafe.Pointer.
//
// An object that you use to inspect locally cached media data.
func AssetCacheFrom(ptr unsafe.Pointer) AssetCache {
	return AssetCache{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetCache *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetCache */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetCache */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetCache */

// Returns an array of extended language tags for languages that can be selected for offline operations via use of the AVMediaSelectionGroup’s AVCustomMediaSelectionScheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationLanguages(for:)
func (a_ AssetCache) MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("mediaPresentationLanguagesForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}/* debug [instance_methods/method]: MediaPresentationLanguagesForMediaSelectionGroup */


// For each AVMediaPresentationSelector defined by the AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the AVMediaPresentationSettings that can be satisfied for offline operations, e.g. playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationSettings(for:)
func (a_ AssetCache) MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("mediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}/* debug [instance_methods/method]: MediaPresentationSettingsForMediaSelectionGroup */


// Returns an array of locally cached media selection options that are available for offline use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaSelectionOptions(in:)
func (a_ AssetCache) MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](a_.ID, objc.Sel("mediaSelectionOptionsInMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}/* debug [instance_methods/method]: MediaSelectionOptionsInMediaSelectionGroup */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetCache */

// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/isPlayableOffline
func (a_ AssetCache) PlayableOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playableOffline"))
	return rv
}/* debug [instance_properties/getter]: playableOffline */


// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/isplayableoffline
func (a_ AssetCache) IsPlayableOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayableOffline"))
	return rv
}/* debug [instance_properties/getter]: isPlayableOffline */


// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/isplayableoffline
func (a_ AssetCache) SetIsPlayableOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayableOffline:"), value)
}/* debug [instance_properties/setter]: isPlayableOffline */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetCache */



