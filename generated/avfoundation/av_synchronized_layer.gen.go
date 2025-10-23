// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [SynchronizedLayer] class.
var (
	SynchronizedLayerClass     _SynchronizedLayerClass
	SynchronizedLayerClassOnce sync.Once
)

func getSynchronizedLayerClass() _SynchronizedLayerClass {
	SynchronizedLayerClassOnce.Do(func() {
		SynchronizedLayerClass = _SynchronizedLayerClass{objc.GetClass("AVSynchronizedLayer")}
	})
	return SynchronizedLayerClass
}

type _SynchronizedLayerClass struct {
	class objc.Class
}

// An interface definition for the [SynchronizedLayer] class.
type ISynchronizedLayer interface {
	quartzcore.ILayer
	AVCoreAnimationBeginTimeAtZero() foundation.TimeInterval
	PlayerItem() IAVPlayerItem
	SetPlayerItem(value IAVPlayerItem)
	BeginTime() foundation.TimeInterval
	SetBeginTime(value foundation.TimeInterval)
}

// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
//
// You can create an arbitrary number of synchronized layers from the same object. A synchronized layer is similar to a object in that it doesn’t display anything itself, it just confers state upon its layer subtree. confers its timing state, synchronizing the timing of layers in its subtree with that of a player item. Any layer with animation property set that is added as a sublayer of should set the animation’s property to a non-zero positive value so animations will be interpreted on the player item’s timeline. replaces the default of 0.0 with . To start the animation from time 0, use a small positive value like . You might use a layer as shown in the following example:


// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer
type SynchronizedLayer struct {
	quartzcore.Layer
}

// SynchronizedLayerFrom constructs a [SynchronizedLayer] from an unsafe.Pointer.
//
// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
func SynchronizedLayerFrom(ptr unsafe.Pointer) SynchronizedLayer {
	return SynchronizedLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SynchronizedLayerClass) Alloc() SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SynchronizedLayerClass) New() SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SynchronizedLayer) Init() SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SynchronizedLayer) Autorelease() SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSynchronizedLayer creates a new SynchronizedLayer instance.
func NewSynchronizedLayer() SynchronizedLayer {
	return getSynchronizedLayerClass().New()
}



// A value that sets an animation begin time to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoreanimationbegintimeatzero
func (s_ SynchronizedLayer) AVCoreAnimationBeginTimeAtZero() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](s_.ID, objc.Sel("AVCoreAnimationBeginTimeAtZero"))
	return rv
}


// The player item to which the timing of the layer is synchronized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/playeritem
func (s_ SynchronizedLayer) PlayerItem() IAVPlayerItem {
	rv := objc.Send[AVPlayerItem](s_.ID, objc.Sel("playerItem"))
	return rv
}


// The player item to which the timing of the layer is synchronized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsynchronizedlayer/playeritem
func (s_ SynchronizedLayer) SetPlayerItem(value IAVPlayerItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPlayerItem:"), value)
}


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (s_ SynchronizedLayer) BeginTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](s_.ID, objc.Sel("beginTime"))
	return rv
}


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (s_ SynchronizedLayer) SetBeginTime(value foundation.TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBeginTime:"), value)
}



