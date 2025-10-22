// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNAssetInfo] class.
var (
	CNAssetInfoClass     _CNAssetInfoClass
	CNAssetInfoClassOnce sync.Once
)

func getCNAssetInfoClass() _CNAssetInfoClass {
	CNAssetInfoClassOnce.Do(func() {
		CNAssetInfoClass = _CNAssetInfoClass{objc.GetClass("CNAssetInfo")}
	})
	return CNAssetInfoClass
}

type _CNAssetInfoClass struct {
	class objc.Class
}

// An interface definition for the [CNAssetInfo] class.
type ICNAssetInfo interface {
	objectivec.IObject
	FrameTimingTrack() avfoundation.AssetTrack
	VideoCompositionTracks() []avfoundation.AssetTrack
}

// An object that provides Cinematic-specific information about an asset, including its tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t
type CNAssetInfo struct {
	objectivec.Object
}

// CNAssetInfoFrom constructs a [CNAssetInfo] from an unsafe.Pointer.
//
// An object that provides Cinematic-specific information about an asset, including its tracks.
func CNAssetInfoFrom(ptr unsafe.Pointer) CNAssetInfo {
	return CNAssetInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNAssetInfoClass) Alloc() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNAssetInfoClass) New() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNAssetInfo) Init() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNAssetInfo) Autorelease() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNAssetInfo creates a new CNAssetInfo instance.
func NewCNAssetInfo() CNAssetInfo {
	return getCNAssetInfoClass().New()
}


// Determines if the asset is Cinematic.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/checkIfCinematic:completionHandler:
func (cc _CNAssetInfoClass) CheckIfCinematicCompletionHandler(asset avfoundation.IAsset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("checkIfCinematic:completionHandler:"), asset, completionHandler)
}

// The track used for Cinematic frame timing.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/frameTimingTrack
func (c_ CNAssetInfo) FrameTimingTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("frameTimingTrack"))
	return rv
}

// Tracks required to construct the video composition output.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/videoCompositionTracks
func (c_ CNAssetInfo) VideoCompositionTracks() []avfoundation.AssetTrack {
	rv := objc.Send[[]avfoundation.AssetTrack](c_.ID, objc.Sel("videoCompositionTracks"))
	return rv
}



