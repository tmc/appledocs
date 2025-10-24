// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVSynchronizedLayer */


/* debug [class_header]: Header for AVSynchronizedLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SynchronizedLayer */
// An interface definition for the [SynchronizedLayer] class.
type ISynchronizedLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for SynchronizedLayer */
	// properties:
	PlayerItem() IAVPlayerItem
	SetPlayerItem(value IAVPlayerItem)
	AVCoreAnimationBeginTimeAtZero() float64
	BeginTime() float64
	SetBeginTime(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SynchronizedLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SynchronizedLayer */
// Alloc allocates a new instance without initialization.
func (sc _SynchronizedLayerClass) Alloc() SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SynchronizedLayer */
// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
//
// You can create an arbitrary number of synchronized layers from the same object. A synchronized layer is similar to a object in that it doesn’t display anything itself, it just confers state upon its layer subtree. confers its timing state, synchronizing the timing of layers in its subtree with that of a player item. Any layer with animation property set that is added as a sublayer of should set the animation’s property to a non-zero positive value so animations will be interpreted on the player item’s timeline. replaces the default of 0.0 with . To start the animation from time 0, use a small positive value like . You might use a layer as shown in the following example:


// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer
type SynchronizedLayer struct {
	Layer
}

// SynchronizedLayerFrom constructs a [SynchronizedLayer] from an unsafe.Pointer.
//
// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
func SynchronizedLayerFrom(ptr unsafe.Pointer) SynchronizedLayer {
	return SynchronizedLayer{
		Layer: LayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SynchronizedLayer */

// Creates a new synchronized layer with timing synchronized with a given player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/init(playerItem:)
func NewSynchronizedLayerWithPlayerItem(playerItem IAVPlayerItem) SynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](objc.ID(getSynchronizedLayerClass().class), objc.Sel("synchronizedLayerWithPlayerItem:"), playerItem)
	return rv
}/* debug [class_init_methods/constructor]: NewSynchronizedLayerWithPlayerItem */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SynchronizedLayer */

// Creates a new synchronized layer with timing synchronized with a given player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/init(playerItem:)
func (sc _SynchronizedLayerClass) SynchronizedLayerWithPlayerItem(playerItem IAVPlayerItem) ISynchronizedLayer {
	rv := objc.Send[SynchronizedLayer](objc.ID(sc.class), objc.Sel("synchronizedLayerWithPlayerItem:"), playerItem)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SynchronizedLayerWithPlayerItem) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SynchronizedLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SynchronizedLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SynchronizedLayer */

// The player item to which the timing of the layer is synchronized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/playerItem
func (s_ SynchronizedLayer) PlayerItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](s_.ID, objc.Sel("playerItem"))
	return rv
}/* debug [instance_properties/getter]: playerItem */


// The player item to which the timing of the layer is synchronized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/playerItem
func (s_ SynchronizedLayer) SetPlayerItem(value IAVPlayerItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPlayerItem:"), value)
}/* debug [instance_properties/setter]: playerItem */


// A value that sets an animation begin time to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoreanimationbegintimeatzero
func (s_ SynchronizedLayer) AVCoreAnimationBeginTimeAtZero() float64 {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("AVCoreAnimationBeginTimeAtZero"))
	return rv
}/* debug [instance_properties/getter]: AVCoreAnimationBeginTimeAtZero */


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (s_ SynchronizedLayer) BeginTime() float64 {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("beginTime"))
	return rv
}/* debug [instance_properties/getter]: beginTime */


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (s_ SynchronizedLayer) SetBeginTime(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBeginTime:"), value)
}/* debug [instance_properties/setter]: beginTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSynchronizedLayer */


