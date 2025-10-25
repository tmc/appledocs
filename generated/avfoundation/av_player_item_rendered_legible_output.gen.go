// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemRenderedLegibleOutput */


/* debug [class_header]: Header for AVPlayerItemRenderedLegibleOutput */
// The class instance for the [PlayerItemRenderedLegibleOutput] class.
var (
	PlayerItemRenderedLegibleOutputClass     _PlayerItemRenderedLegibleOutputClass
	PlayerItemRenderedLegibleOutputClassOnce sync.Once
)

func getPlayerItemRenderedLegibleOutputClass() _PlayerItemRenderedLegibleOutputClass {
	PlayerItemRenderedLegibleOutputClassOnce.Do(func() {
		PlayerItemRenderedLegibleOutputClass = _PlayerItemRenderedLegibleOutputClass{objc.GetClass("AVPlayerItemRenderedLegibleOutput")}
	})
	return PlayerItemRenderedLegibleOutputClass
}

type _PlayerItemRenderedLegibleOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemRenderedLegibleOutput */
// An interface definition for the [PlayerItemRenderedLegibleOutput] class.
type IPlayerItemRenderedLegibleOutput interface {
	IPlayerItemOutput
	
/* debug [class_interface_properties]: Properties for PlayerItemRenderedLegibleOutput */
	// properties:
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
	VideoDisplaySize() corefoundation.CGSize
	SetVideoDisplaySize(value corefoundation.CGSize)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemRenderedLegibleOutput */
	// methods:
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemRenderedLegibleOutput */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemRenderedLegibleOutputClass) Alloc() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemRenderedLegibleOutputClass) New() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemRenderedLegibleOutput) Init() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemRenderedLegibleOutput) Autorelease() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemRenderedLegibleOutput creates a new PlayerItemRenderedLegibleOutput instance.
func NewPlayerItemRenderedLegibleOutput() PlayerItemRenderedLegibleOutput {
	return getPlayerItemRenderedLegibleOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemRenderedLegibleOutput */
// A player item output that vends media with a legible characteristic as rendered pixel buffers.


// A player item output that vends media with a legible characteristic as rendered pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput
type PlayerItemRenderedLegibleOutput struct {
	PlayerItemOutput
}

// PlayerItemRenderedLegibleOutputFrom constructs a [PlayerItemRenderedLegibleOutput] from an unsafe.Pointer.
//
// A player item output that vends media with a legible characteristic as rendered pixel buffers.
func PlayerItemRenderedLegibleOutputFrom(ptr unsafe.Pointer) PlayerItemRenderedLegibleOutput {
	return PlayerItemRenderedLegibleOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemRenderedLegibleOutput */

// Creates a rendered legible output object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/init(videoDisplay:)
func NewPlayerItemRenderedLegibleOutputWithVideoDisplaySize(videoDisplaySize corefoundation.CGSize) PlayerItemRenderedLegibleOutput {
	instance := getPlayerItemRenderedLegibleOutputClass().Alloc()
	rv := objc.Send[PlayerItemRenderedLegibleOutput](instance.ID, objc.Sel("initWithVideoDisplaySize:"), videoDisplaySize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerItemRenderedLegibleOutputWithVideoDisplaySize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemRenderedLegibleOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemRenderedLegibleOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemRenderedLegibleOutput */

// Sets the delegate object and the queue on which it’s invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/setDelegate(_:queue:)
func (p_ PlayerItemRenderedLegibleOutput) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemRenderedLegibleOutput */

// Permits advance invocation of the associated delegate, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemRenderedLegibleOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}/* debug [instance_properties/getter]: advanceIntervalForDelegateInvocation */


// Permits advance invocation of the associated delegate, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemRenderedLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}/* debug [instance_properties/setter]: advanceIntervalForDelegateInvocation */


// A delegate object for this output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/delegate
func (p_ PlayerItemRenderedLegibleOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue on which the output calls the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/delegateQueue
func (p_ PlayerItemRenderedLegibleOutput) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */


// Set the video display size to use for rendering of pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/videoDisplaySize
func (p_ PlayerItemRenderedLegibleOutput) VideoDisplaySize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p_.ID, objc.Sel("videoDisplaySize"))
	return rv
}/* debug [instance_properties/getter]: videoDisplaySize */


// Set the video display size to use for rendering of pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/videoDisplaySize
func (p_ PlayerItemRenderedLegibleOutput) SetVideoDisplaySize(value corefoundation.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoDisplaySize:"), value)
}/* debug [instance_properties/setter]: videoDisplaySize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemRenderedLegibleOutput */


