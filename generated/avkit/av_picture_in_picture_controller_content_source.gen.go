// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PictureInPictureControllerContentSource] class.
var (
	PictureInPictureControllerContentSourceClass     _PictureInPictureControllerContentSourceClass
	PictureInPictureControllerContentSourceClassOnce sync.Once
)

func getPictureInPictureControllerContentSourceClass() _PictureInPictureControllerContentSourceClass {
	PictureInPictureControllerContentSourceClassOnce.Do(func() {
		PictureInPictureControllerContentSourceClass = _PictureInPictureControllerContentSourceClass{objc.GetClass("AVPictureInPictureControllerContentSource")}
	})
	return PictureInPictureControllerContentSourceClass
}

type _PictureInPictureControllerContentSourceClass struct {
	class objc.Class
}





// An interface definition for the [PictureInPictureControllerContentSource] class.
type IPictureInPictureControllerContentSource interface {
	objectivec.IObject
	

	// properties:
	PlayerLayer() avfoundation.PlayerLayer
	SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer
	SampleBufferPlaybackDelegate() unsafe.Pointer
	ContentSource() IAVPictureInPictureControllerContentSource
	SetContentSource(value IAVPictureInPictureControllerContentSource)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureControllerContentSourceClass) Alloc() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PictureInPictureControllerContentSourceClass) New() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PictureInPictureControllerContentSource) Init() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PictureInPictureControllerContentSource) Autorelease() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPictureInPictureControllerContentSource creates a new PictureInPictureControllerContentSource instance.
func NewPictureInPictureControllerContentSource() PictureInPictureControllerContentSource {
	return getPictureInPictureControllerContentSourceClass().New()
}





// An object that represents the source of the content to present in Picture in Picture.
//
// The system supports displaying content from an or in a Picture in Picture window. Use an instance of this class to describe the source of your app’s content.


// An object that represents the source of the content to present in Picture in Picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class
type PictureInPictureControllerContentSource struct {
	objectivec.Object
}

// PictureInPictureControllerContentSourceFrom constructs a [PictureInPictureControllerContentSource] from an unsafe.Pointer.
//
// An object that represents the source of the content to present in Picture in Picture.
func PictureInPictureControllerContentSourceFrom(ptr unsafe.Pointer) PictureInPictureControllerContentSource {
	return PictureInPictureControllerContentSource{objectivec.Object{objc.ID(ptr)}}
}






// Creates a content source with an active video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/init(activeVideoCallSourceView:contentViewController:)
func NewPictureInPictureControllerContentSourceWithActiveVideoCallSourceViewContentViewController(sourceView appkit.View, contentViewController IAVPictureInPictureVideoCallViewController) PictureInPictureControllerContentSource {
	instance := getPictureInPictureControllerContentSourceClass().Alloc()
	rv := objc.Send[PictureInPictureControllerContentSource](instance.ID, objc.Sel("initWithActiveVideoCallSourceView:contentViewController:"), sourceView, contentViewController)
	rv.Autorelease()
	return rv
}


// Creates a content source with a player layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/init(playerLayer:)
func NewPictureInPictureControllerContentSourceWithPlayerLayer(playerLayer avfoundation.PlayerLayer) PictureInPictureControllerContentSource {
	instance := getPictureInPictureControllerContentSourceClass().Alloc()
	rv := objc.Send[PictureInPictureControllerContentSource](instance.ID, objc.Sel("initWithPlayerLayer:"), playerLayer)
	rv.Autorelease()
	return rv
}


// Creates a content source with a sample buffer display layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/init(sampleBufferDisplayLayer:playbackDelegate:)
func NewPictureInPictureControllerContentSourceWithSampleBufferDisplayLayerPlaybackDelegate(sampleBufferDisplayLayer avfoundation.SampleBufferDisplayLayer, playbackDelegate unsafe.Pointer) PictureInPictureControllerContentSource {
	instance := getPictureInPictureControllerContentSourceClass().Alloc()
	rv := objc.Send[PictureInPictureControllerContentSource](instance.ID, objc.Sel("initWithSampleBufferDisplayLayer:playbackDelegate:"), sampleBufferDisplayLayer, playbackDelegate)
	rv.Autorelease()
	return rv
}






















// The presenting player layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/playerLayer
func (p_ PictureInPictureControllerContentSource) PlayerLayer() avfoundation.PlayerLayer {
	rv := objc.Send[avfoundation.PlayerLayer](p_.ID, objc.Sel("playerLayer"))
	return rv
}


// The presenting sample buffer display layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/sampleBufferDisplayLayer
func (p_ PictureInPictureControllerContentSource) SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer {
	rv := objc.Send[avfoundation.SampleBufferDisplayLayer](p_.ID, objc.Sel("sampleBufferDisplayLayer"))
	return rv
}


// A delegate object that responds to sample buffer playback events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/sampleBufferPlaybackDelegate
func (p_ PictureInPictureControllerContentSource) SampleBufferPlaybackDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sampleBufferPlaybackDelegate"))
	return rv
}


// The source of the controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.property
func (p_ PictureInPictureControllerContentSource) ContentSource() IAVPictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](p_.ID, objc.Sel("contentSource"))
	return rv
}


// The source of the controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.property
func (p_ PictureInPictureControllerContentSource) SetContentSource(value IAVPictureInPictureControllerContentSource) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentSource:"), value)
}







