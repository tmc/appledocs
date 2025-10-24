// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemLegibleOutput] class.
var (
	PlayerItemLegibleOutputClass     _PlayerItemLegibleOutputClass
	PlayerItemLegibleOutputClassOnce sync.Once
)

func getPlayerItemLegibleOutputClass() _PlayerItemLegibleOutputClass {
	PlayerItemLegibleOutputClassOnce.Do(func() {
		PlayerItemLegibleOutputClass = _PlayerItemLegibleOutputClass{objc.GetClass("AVPlayerItemLegibleOutput")}
	})
	return PlayerItemLegibleOutputClass
}

type _PlayerItemLegibleOutputClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemLegibleOutput] class.
type IPlayerItemLegibleOutput interface {
	IPlayerItemOutput
	

	// properties:
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
	TextStylingResolution() PlayerItemLegibleOutputTextStylingResolution /* typedef */
	SetTextStylingResolution(value PlayerItemLegibleOutputTextStylingResolution /* typedef */)


	

	// methods:
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemLegibleOutputClass) Alloc() PlayerItemLegibleOutput {
	rv := objc.Send[PlayerItemLegibleOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemLegibleOutputClass) New() PlayerItemLegibleOutput {
	rv := objc.Send[PlayerItemLegibleOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemLegibleOutput) Init() PlayerItemLegibleOutput {
	rv := objc.Send[PlayerItemLegibleOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemLegibleOutput) Autorelease() PlayerItemLegibleOutput {
	rv := objc.Send[PlayerItemLegibleOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemLegibleOutput creates a new PlayerItemLegibleOutput instance.
func NewPlayerItemLegibleOutput() PlayerItemLegibleOutput {
	return getPlayerItemLegibleOutputClass().New()
}





// An object that vends attributed strings for media with a legible characteristic.


// An object that vends attributed strings for media with a legible characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput
type PlayerItemLegibleOutput struct {
	PlayerItemOutput
}

// PlayerItemLegibleOutputFrom constructs a [PlayerItemLegibleOutput] from an unsafe.Pointer.
//
// An object that vends attributed strings for media with a legible characteristic.
func PlayerItemLegibleOutputFrom(ptr unsafe.Pointer) PlayerItemLegibleOutput {
	return PlayerItemLegibleOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}






// Creates an initialized legible-output object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/init(mediaSubtypesForNativeRepresentation:)
func NewPlayerItemLegibleOutputWithMediaSubtypesForNativeRepresentation(subtypes []foundation.Number) PlayerItemLegibleOutput {
	instance := getPlayerItemLegibleOutputClass().Alloc()
	rv := objc.Send[PlayerItemLegibleOutput](instance.ID, objc.Sel("initWithMediaSubtypesForNativeRepresentation:"), subtypes)
	rv.Autorelease()
	return rv
}

















// Sets the receiver’s delegate and a dispatch queue on which the delegate is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/setDelegate(_:queue:)
func (p_ PlayerItemLegibleOutput) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}







// The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemLegibleOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}


// The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}


// The delegate of the output class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/delegate
func (p_ PlayerItemLegibleOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// The dispatch queue on which the delegate is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/delegateQueue
func (p_ PlayerItemLegibleOutput) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("delegateQueue"))
	return rv
}


// A string identifier indicating the degree of text styling to be applied to attributed strings vended by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/textStylingResolution-swift.property
func (p_ PlayerItemLegibleOutput) TextStylingResolution() PlayerItemLegibleOutputTextStylingResolution /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("textStylingResolution"))
	return rv
}


// A string identifier indicating the degree of text styling to be applied to attributed strings vended by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/textStylingResolution-swift.property
func (p_ PlayerItemLegibleOutput) SetTextStylingResolution(value PlayerItemLegibleOutputTextStylingResolution /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextStylingResolution:"), value)
}







