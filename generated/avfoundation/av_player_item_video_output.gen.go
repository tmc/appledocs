// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemVideoOutput */


/* debug [class_header]: Header for AVPlayerItemVideoOutput */
// The class instance for the [PlayerItemVideoOutput] class.
var (
	PlayerItemVideoOutputClass     _PlayerItemVideoOutputClass
	PlayerItemVideoOutputClassOnce sync.Once
)

func getPlayerItemVideoOutputClass() _PlayerItemVideoOutputClass {
	PlayerItemVideoOutputClassOnce.Do(func() {
		PlayerItemVideoOutputClass = _PlayerItemVideoOutputClass{objc.GetClass("AVPlayerItemVideoOutput")}
	})
	return PlayerItemVideoOutputClass
}

type _PlayerItemVideoOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemVideoOutput */
// An interface definition for the [PlayerItemVideoOutput] class.
type IPlayerItemVideoOutput interface {
	IPlayerItemOutput
	
/* debug [class_interface_properties]: Properties for PlayerItemVideoOutput */
	// properties:
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemVideoOutput */
	// methods:
	HasNewPixelBufferForItemTime(itemTime objc.IObject /* cross-framework: Time */) bool
	RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval float64)
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemVideoOutput */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemVideoOutputClass) Alloc() PlayerItemVideoOutput {
	rv := objc.Send[PlayerItemVideoOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemVideoOutputClass) New() PlayerItemVideoOutput {
	rv := objc.Send[PlayerItemVideoOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemVideoOutput) Init() PlayerItemVideoOutput {
	rv := objc.Send[PlayerItemVideoOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemVideoOutput) Autorelease() PlayerItemVideoOutput {
	rv := objc.Send[PlayerItemVideoOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemVideoOutput creates a new PlayerItemVideoOutput instance.
func NewPlayerItemVideoOutput() PlayerItemVideoOutput {
	return getPlayerItemVideoOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemVideoOutput */
// An object that outputs video frames from a player item.


// An object that outputs video frames from a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput
type PlayerItemVideoOutput struct {
	PlayerItemOutput
}

// PlayerItemVideoOutputFrom constructs a [PlayerItemVideoOutput] from an unsafe.Pointer.
//
// An object that outputs video frames from a player item.
func PlayerItemVideoOutputFrom(ptr unsafe.Pointer) PlayerItemVideoOutput {
	return PlayerItemVideoOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemVideoOutput */

// Creates a video output object initialized with the specified output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/init(outputSettings:)
func NewPlayerItemVideoOutputWithOutputSettings(outputSettings foundation.IDictionary) PlayerItemVideoOutput {
	instance := getPlayerItemVideoOutputClass().Alloc()
	rv := objc.Send[PlayerItemVideoOutput](instance.ID, objc.Sel("initWithOutputSettings:"), outputSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerItemVideoOutputWithOutputSettings */


// Creates a video output object using the specified pixel buffer attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/init(pixelBufferAttributes:)-7n7v8
func NewPlayerItemVideoOutputWithPixelBufferAttributes(pixelBufferAttributes foundation.IDictionary) PlayerItemVideoOutput {
	instance := getPlayerItemVideoOutputClass().Alloc()
	rv := objc.Send[PlayerItemVideoOutput](instance.ID, objc.Sel("initWithPixelBufferAttributes:"), pixelBufferAttributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerItemVideoOutputWithPixelBufferAttributes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemVideoOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemVideoOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemVideoOutput */

// Returns a Boolean value that indicates whether video output is available for the specified item time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/hasNewPixelBuffer(forItemTime:)
func (p_ PlayerItemVideoOutput) HasNewPixelBufferForItemTime(itemTime objc.IObject /* cross-framework: Time */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasNewPixelBufferForItemTime:"), itemTime)
	return rv
}/* debug [instance_methods/method]: HasNewPixelBufferForItemTime */


// Tells the receiver that the video out put client is entering a quiescent state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/requestNotificationOfMediaDataChange(withAdvanceInterval:)
func (p_ PlayerItemVideoOutput) RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("requestNotificationOfMediaDataChangeWithAdvanceInterval:"), interval)
}/* debug [instance_methods/method]: RequestNotificationOfMediaDataChangeWithAdvanceInterval */


// Sets the delegate and dispatch queue for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/setDelegate(_:queue:)
func (p_ PlayerItemVideoOutput) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemVideoOutput */

// The delegate for the video output object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/delegate
func (p_ PlayerItemVideoOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue on which to call delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/delegateQueue
func (p_ PlayerItemVideoOutput) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemVideoOutput */


