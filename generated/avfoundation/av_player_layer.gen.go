// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [AVPlayerLayer] class.
var aVPlayerLayerClass = _AVPlayerLayerClass{objc.GetClass("AVPlayerLayer")}

type _AVPlayerLayerClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerLayer] class.
type IAVPlayerLayer interface {
	quartzcore.ILayer
}

// An object that presents the visual contents of a player object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer

type AVPlayerLayer struct {
	quartzcore.Layer
}

// AVPlayerLayerFrom constructs a [AVPlayerLayer] from an unsafe.Pointer.
//
// An object that presents the visual contents of a player object.
func AVPlayerLayerFrom(ptr unsafe.Pointer) AVPlayerLayer {
	return AVPlayerLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _AVPlayerLayerClass) Alloc() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVPlayerLayerClass) New() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerLayer) Init() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerLayer) Autorelease() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerLayer creates a new AVPlayerLayer instance.
func NewAVPlayerLayer() AVPlayerLayer {
	return aVPlayerLayerClass.New()
}




