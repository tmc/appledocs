// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetReaderVideoCompositionOutput] class.
var (
	AssetReaderVideoCompositionOutputClass     _AssetReaderVideoCompositionOutputClass
	AssetReaderVideoCompositionOutputClassOnce sync.Once
)

func getAssetReaderVideoCompositionOutputClass() _AssetReaderVideoCompositionOutputClass {
	AssetReaderVideoCompositionOutputClassOnce.Do(func() {
		AssetReaderVideoCompositionOutputClass = _AssetReaderVideoCompositionOutputClass{objc.GetClass("AVAssetReaderVideoCompositionOutput")}
	})
	return AssetReaderVideoCompositionOutputClass
}

type _AssetReaderVideoCompositionOutputClass struct {
	class objc.Class
}





// An interface definition for the [AssetReaderVideoCompositionOutput] class.
type IAssetReaderVideoCompositionOutput interface {
	IAssetReaderOutput
	

	// properties:
	CustomVideoCompositor() unsafe.Pointer
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
	VideoSettings() foundation.IDictionary
	VideoTracks() []AssetTrack


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetReaderVideoCompositionOutputClass) Alloc() AssetReaderVideoCompositionOutput {
	rv := objc.Send[AssetReaderVideoCompositionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderVideoCompositionOutputClass) New() AssetReaderVideoCompositionOutput {
	rv := objc.Send[AssetReaderVideoCompositionOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderVideoCompositionOutput) Init() AssetReaderVideoCompositionOutput {
	rv := objc.Send[AssetReaderVideoCompositionOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderVideoCompositionOutput) Autorelease() AssetReaderVideoCompositionOutput {
	rv := objc.Send[AssetReaderVideoCompositionOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderVideoCompositionOutput creates a new AssetReaderVideoCompositionOutput instance.
func NewAssetReaderVideoCompositionOutput() AssetReaderVideoCompositionOutput {
	return getAssetReaderVideoCompositionOutputClass().New()
}





// An object that reads composited video frames from one or more tracks of an asset.


// An object that reads composited video frames from one or more tracks of an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput
type AssetReaderVideoCompositionOutput struct {
	AssetReaderOutput
}

// AssetReaderVideoCompositionOutputFrom constructs a [AssetReaderVideoCompositionOutput] from an unsafe.Pointer.
//
// An object that reads composited video frames from one or more tracks of an asset.
func AssetReaderVideoCompositionOutputFrom(ptr unsafe.Pointer) AssetReaderVideoCompositionOutput {
	return AssetReaderVideoCompositionOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}






// Creates an object that reads composited video frames from the specified video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/init(videoTracks:videoSettings:)
func NewAssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []AssetTrack, videoSettings foundation.IDictionary) AssetReaderVideoCompositionOutput {
	instance := getAssetReaderVideoCompositionOutputClass().Alloc()
	rv := objc.Send[AssetReaderVideoCompositionOutput](instance.ID, objc.Sel("initWithVideoTracks:videoSettings:"), videoTracks, videoSettings)
	rv.Autorelease()
	return rv
}







// Returns a new object that reads composited video from the specified video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/assetReaderVideoCompositionOutputWithVideoTracks:videoSettings:
func (ac _AssetReaderVideoCompositionOutputClass) AssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []AssetTrack, videoSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderVideoCompositionOutputWithVideoTracks:videoSettings:"), videoTracks, videoSettings)
	return rv
}

















// A custom video compositor for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/customVideoCompositor
func (a_ AssetReaderVideoCompositionOutput) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// The video composition to use for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoComposition
func (a_ AssetReaderVideoCompositionOutput) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}


// The video composition to use for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoComposition
func (a_ AssetReaderVideoCompositionOutput) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}


// The video settings that the output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoSettings
func (a_ AssetReaderVideoCompositionOutput) VideoSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("videoSettings"))
	return rv
}


// The tracks from which the output reads the composited video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoTracks
func (a_ AssetReaderVideoCompositionOutput) VideoTracks() []AssetTrack {
	rv := objc.Send[[]AssetTrack](a_.ID, objc.Sel("videoTracks"))
	return rv
}







