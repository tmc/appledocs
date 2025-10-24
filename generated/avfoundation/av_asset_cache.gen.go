// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [AssetCache] class.
type IAssetCache interface {
	objectivec.IObject
	

	// properties:
	PlayableOffline() bool
	IsPlayableOffline() bool
	SetIsPlayableOffline(value bool)


	

	// methods:
	MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string
	MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary
	MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []MediaSelectionOption


}





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




















// Returns an array of extended language tags for languages that can be selected for offline operations via use of the AVMediaSelectionGroup’s AVCustomMediaSelectionScheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationLanguages(for:)
func (a_ AssetCache) MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("mediaPresentationLanguagesForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// For each AVMediaPresentationSelector defined by the AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the AVMediaPresentationSettings that can be satisfied for offline operations, e.g. playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationSettings(for:)
func (a_ AssetCache) MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("mediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Returns an array of locally cached media selection options that are available for offline use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaSelectionOptions(in:)
func (a_ AssetCache) MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](a_.ID, objc.Sel("mediaSelectionOptionsInMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}







// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/isPlayableOffline
func (a_ AssetCache) PlayableOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playableOffline"))
	return rv
}


// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/isplayableoffline
func (a_ AssetCache) IsPlayableOffline() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlayableOffline"))
	return rv
}


// A Boolean value that indicates whether the asset is playable without an internet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/isplayableoffline
func (a_ AssetCache) SetIsPlayableOffline(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlayableOffline:"), value)
}








