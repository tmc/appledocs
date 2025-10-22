// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
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
	ActiveVideoCallSourceView() appkit.View
	ActiveVideoCallContentViewController() AVPictureInPictureVideoCallViewController
	SetActiveVideoCallContentViewController(value IAVPictureInPictureVideoCallViewController)
	PlayerLayer() avfoundation.PlayerLayer
	SetPlayerLayer(value avfoundation.IPlayerLayer)
	SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer
	SetSampleBufferDisplayLayer(value avfoundation.ISampleBufferDisplayLayer)
	SampleBufferPlaybackDelegate() unsafe.Pointer
	SetSampleBufferPlaybackDelegate(value unsafe.Pointer)
	ContentSource() AVPictureInPictureControllerContentSource
	SetContentSource(value IAVPictureInPictureControllerContentSource)
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

// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureControllerContentSourceClass) Alloc() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The view that contains the video content of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/activeVideoCallSourceView

func (p_ PictureInPictureControllerContentSource) ActiveVideoCallSourceView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("activeVideoCallSourceView"))
	return rv
}


// The view controller that presents the video call content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallcontentviewcontroller

func (p_ PictureInPictureControllerContentSource) ActiveVideoCallContentViewController() AVPictureInPictureVideoCallViewController {
	rv := objc.Send[AVPictureInPictureVideoCallViewController](p_.ID, objc.Sel("activeVideoCallContentViewController"))
	return rv
}


// The view controller that presents the video call content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallcontentviewcontroller

func (p_ PictureInPictureControllerContentSource) SetActiveVideoCallContentViewController(value IAVPictureInPictureVideoCallViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActiveVideoCallContentViewController:"), value)
}


// The presenting player layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/playerlayer

func (p_ PictureInPictureControllerContentSource) PlayerLayer() avfoundation.PlayerLayer {
	rv := objc.Send[avfoundation.PlayerLayer](p_.ID, objc.Sel("playerLayer"))
	return rv
}


// The presenting player layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/playerlayer

func (p_ PictureInPictureControllerContentSource) SetPlayerLayer(value avfoundation.IPlayerLayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayerLayer:"), value)
}


// The presenting sample buffer display layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/samplebufferdisplaylayer

func (p_ PictureInPictureControllerContentSource) SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer {
	rv := objc.Send[avfoundation.SampleBufferDisplayLayer](p_.ID, objc.Sel("sampleBufferDisplayLayer"))
	return rv
}


// The presenting sample buffer display layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/samplebufferdisplaylayer

func (p_ PictureInPictureControllerContentSource) SetSampleBufferDisplayLayer(value avfoundation.ISampleBufferDisplayLayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSampleBufferDisplayLayer:"), value)
}


// A delegate object that responds to sample buffer playback events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/samplebufferplaybackdelegate

func (p_ PictureInPictureControllerContentSource) SampleBufferPlaybackDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sampleBufferPlaybackDelegate"))
	return rv
}


// A delegate object that responds to sample buffer playback events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/samplebufferplaybackdelegate

func (p_ PictureInPictureControllerContentSource) SetSampleBufferPlaybackDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSampleBufferPlaybackDelegate:"), value)
}


// The source of the controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.property

func (p_ PictureInPictureControllerContentSource) ContentSource() AVPictureInPictureControllerContentSource {
	rv := objc.Send[AVPictureInPictureControllerContentSource](p_.ID, objc.Sel("contentSource"))
	return rv
}


// The source of the controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.property

func (p_ PictureInPictureControllerContentSource) SetContentSource(value IAVPictureInPictureControllerContentSource) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentSource:"), value)
}



