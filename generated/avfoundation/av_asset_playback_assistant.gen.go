// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetPlaybackAssistant] class.
var (
	AssetPlaybackAssistantClass     _AssetPlaybackAssistantClass
	AssetPlaybackAssistantClassOnce sync.Once
)

func getAssetPlaybackAssistantClass() _AssetPlaybackAssistantClass {
	AssetPlaybackAssistantClassOnce.Do(func() {
		AssetPlaybackAssistantClass = _AssetPlaybackAssistantClass{objc.GetClass("AVAssetPlaybackAssistant")}
	})
	return AssetPlaybackAssistantClass
}

type _AssetPlaybackAssistantClass struct {
	class objc.Class
}





// An interface definition for the [AssetPlaybackAssistant] class.
type IAssetPlaybackAssistant interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	LoadPlaybackConfigurationOptionsWithCompletionHandler(completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetPlaybackAssistantClass) Alloc() AssetPlaybackAssistant {
	rv := objc.Send[AssetPlaybackAssistant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetPlaybackAssistantClass) New() AssetPlaybackAssistant {
	rv := objc.Send[AssetPlaybackAssistant](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetPlaybackAssistant) Init() AssetPlaybackAssistant {
	rv := objc.Send[AssetPlaybackAssistant](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetPlaybackAssistant) Autorelease() AssetPlaybackAssistant {
	rv := objc.Send[AssetPlaybackAssistant](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetPlaybackAssistant creates a new AssetPlaybackAssistant instance.
func NewAssetPlaybackAssistant() AssetPlaybackAssistant {
	return getAssetPlaybackAssistantClass().New()
}





// An object that provides playback information for an asset.


// An object that provides playback information for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant
type AssetPlaybackAssistant struct {
	objectivec.Object
}

// AssetPlaybackAssistantFrom constructs a [AssetPlaybackAssistant] from an unsafe.Pointer.
//
// An object that provides playback information for an asset.
func AssetPlaybackAssistantFrom(ptr unsafe.Pointer) AssetPlaybackAssistant {
	return AssetPlaybackAssistant{objectivec.Object{objc.ID(ptr)}}
}






// Creates a playback assistant to inspect the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant/init(asset:)
func NewAssetPlaybackAssistantWithAsset(asset IAVAsset) AssetPlaybackAssistant {
	rv := objc.Send[AssetPlaybackAssistant](objc.ID(getAssetPlaybackAssistantClass().class), objc.Sel("assetPlaybackAssistantWithAsset:"), asset)
	return rv
}







// Creates a playback assistant to inspect the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant/init(asset:)
func (ac _AssetPlaybackAssistantClass) AssetPlaybackAssistantWithAsset(asset IAVAsset) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetPlaybackAssistantWithAsset:"), asset)
	return rv
}












// Loads playback configuration options for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant/loadPlaybackConfigurationOptions(completionHandler:)
func (a_ AssetPlaybackAssistant) LoadPlaybackConfigurationOptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadPlaybackConfigurationOptionsWithCompletionHandler:"), completionHandler)
}












