// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PlayerItemRenderedLegibleOutput] class.
type IPlayerItemRenderedLegibleOutput interface {
	IPlayerItemOutput
	// properties:
	AdvanceIntervalForDelegateInvocation() unsafe.Pointer
	SetAdvanceIntervalForDelegateInvocation(value unsafe.Pointer)
	Delegate() PlayerItemRenderedLegibleOutputPushDelegate /* not a class type */
	SetDelegate(value PlayerItemRenderedLegibleOutputPushDelegate /* not a class type */)
	DelegateQueue() unsafe.Pointer
	SetDelegateQueue(value unsafe.Pointer)
	VideoDisplaySize() objc.IObject /* cross-framework: Size */
	SetVideoDisplaySize(value objc.IObject /* cross-framework: Size */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemRenderedLegibleOutputClass) Alloc() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Permits advance invocation of the associated delegate, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/advanceintervalfordelegateinvocation
func (p_ PlayerItemRenderedLegibleOutput) AdvanceIntervalForDelegateInvocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}


// Permits advance invocation of the associated delegate, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/advanceintervalfordelegateinvocation
func (p_ PlayerItemRenderedLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}


// A delegate object for this output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/delegate
func (p_ PlayerItemRenderedLegibleOutput) Delegate() PlayerItemRenderedLegibleOutputPushDelegate /* not a class type */ {
	rv := objc.Send[PlayerItemRenderedLegibleOutputPushDelegate](p_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object for this output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/delegate
func (p_ PlayerItemRenderedLegibleOutput) SetDelegate(value PlayerItemRenderedLegibleOutputPushDelegate /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// The dispatch queue on which the output calls the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/delegatequeue
func (p_ PlayerItemRenderedLegibleOutput) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegateQueue"))
	return rv
}


// The dispatch queue on which the output calls the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/delegatequeue
func (p_ PlayerItemRenderedLegibleOutput) SetDelegateQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegateQueue:"), value)
}


// Set the video display size to use for rendering of pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/videodisplaysize
func (p_ PlayerItemRenderedLegibleOutput) VideoDisplaySize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](p_.ID, objc.Sel("videoDisplaySize"))
	return rv
}


// Set the video display size to use for rendering of pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemrenderedlegibleoutput/videodisplaysize
func (p_ PlayerItemRenderedLegibleOutput) SetVideoDisplaySize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoDisplaySize:"), value)
}



