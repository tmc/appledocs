// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemMetadataOutput] class.
var (
	PlayerItemMetadataOutputClass     _PlayerItemMetadataOutputClass
	PlayerItemMetadataOutputClassOnce sync.Once
)

func getPlayerItemMetadataOutputClass() _PlayerItemMetadataOutputClass {
	PlayerItemMetadataOutputClassOnce.Do(func() {
		PlayerItemMetadataOutputClass = _PlayerItemMetadataOutputClass{objc.GetClass("AVPlayerItemMetadataOutput")}
	})
	return PlayerItemMetadataOutputClass
}

type _PlayerItemMetadataOutputClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemMetadataOutput] class.
type IPlayerItemMetadataOutput interface {
	IPlayerItemOutput
	

	// properties:
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
	SuppressesPlayerRendering() bool
	SetSuppressesPlayerRendering(value bool)


	

	// methods:
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemMetadataOutputClass) Alloc() PlayerItemMetadataOutput {
	rv := objc.Send[PlayerItemMetadataOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemMetadataOutputClass) New() PlayerItemMetadataOutput {
	rv := objc.Send[PlayerItemMetadataOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemMetadataOutput) Init() PlayerItemMetadataOutput {
	rv := objc.Send[PlayerItemMetadataOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemMetadataOutput) Autorelease() PlayerItemMetadataOutput {
	rv := objc.Send[PlayerItemMetadataOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemMetadataOutput creates a new PlayerItemMetadataOutput instance.
func NewPlayerItemMetadataOutput() PlayerItemMetadataOutput {
	return getPlayerItemMetadataOutputClass().New()
}





// An object that vends collections of metadata items that a player item’s tracks carry.


// An object that vends collections of metadata items that a player item’s tracks carry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput
type PlayerItemMetadataOutput struct {
	PlayerItemOutput
}

// PlayerItemMetadataOutputFrom constructs a [PlayerItemMetadataOutput] from an unsafe.Pointer.
//
// An object that vends collections of metadata items that a player item’s tracks carry.
func PlayerItemMetadataOutputFrom(ptr unsafe.Pointer) PlayerItemMetadataOutput {
	return PlayerItemMetadataOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}






// Creates an instance of AVPlayerItemMetadataOutput.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/init(identifiers:)
func NewPlayerItemMetadataOutputWithIdentifiers(identifiers []string) PlayerItemMetadataOutput {
	instance := getPlayerItemMetadataOutputClass().Alloc()
	rv := objc.Send[PlayerItemMetadataOutput](instance.ID, objc.Sel("initWithIdentifiers:"), identifiers)
	rv.Autorelease()
	return rv
}

















// Sets the delegate and a dispatch queue on which the delegate is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/setDelegate(_:queue:)
func (p_ PlayerItemMetadataOutput) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}







// The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemMetadataOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}


// The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/advanceIntervalForDelegateInvocation
func (p_ PlayerItemMetadataOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}


// The delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/delegate
func (p_ PlayerItemMetadataOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// The dispatch queue on which messages are sent to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/delegateQueue
func (p_ PlayerItemMetadataOutput) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("delegateQueue"))
	return rv
}


// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/suppressesplayerrendering
func (p_ PlayerItemMetadataOutput) SuppressesPlayerRendering() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("suppressesPlayerRendering"))
	return rv
}


// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/suppressesplayerrendering
func (p_ PlayerItemMetadataOutput) SetSuppressesPlayerRendering(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSuppressesPlayerRendering:"), value)
}







