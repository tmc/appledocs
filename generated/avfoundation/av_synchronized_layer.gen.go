// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [AVSynchronizedLayer] class.
var (
	aVSynchronizedLayerClass     _AVSynchronizedLayerClass
	aVSynchronizedLayerClassOnce sync.Once
)

func getAVSynchronizedLayerClass() _AVSynchronizedLayerClass {
	aVSynchronizedLayerClassOnce.Do(func() {
		aVSynchronizedLayerClass = _AVSynchronizedLayerClass{objc.GetClass("AVSynchronizedLayer")}
	})
	return aVSynchronizedLayerClass
}

type _AVSynchronizedLayerClass struct {
	class objc.Class
}

// An interface definition for the [AVSynchronizedLayer] class.
type IAVSynchronizedLayer interface {
	quartzcore.ILayer
}

// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer
type AVSynchronizedLayer struct {
	quartzcore.Layer
}

// AVSynchronizedLayerFrom constructs a [AVSynchronizedLayer] from an unsafe.Pointer.
//
// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
func AVSynchronizedLayerFrom(ptr unsafe.Pointer) AVSynchronizedLayer {
	return AVSynchronizedLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVSynchronizedLayerClass) Alloc() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVSynchronizedLayerClass) New() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVSynchronizedLayer) Init() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVSynchronizedLayer) Autorelease() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSynchronizedLayer creates a new AVSynchronizedLayer instance.
func NewAVSynchronizedLayer() AVSynchronizedLayer {
	return getAVSynchronizedLayerClass().New()
}




