// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RPBroadcastController] class.
var (
	RPBroadcastControllerClass     _RPBroadcastControllerClass
	RPBroadcastControllerClassOnce sync.Once
)

func getRPBroadcastControllerClass() _RPBroadcastControllerClass {
	RPBroadcastControllerClassOnce.Do(func() {
		RPBroadcastControllerClass = _RPBroadcastControllerClass{objc.GetClass("RPBroadcastController")}
	})
	return RPBroadcastControllerClass
}

type _RPBroadcastControllerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastController] class.
type IRPBroadcastController interface {
	objectivec.IObject
	// properties:
	BroadcastURL() objc.IObject /* cross-framework: NSURL */
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Broadcasting() bool
	Paused() bool
	ServiceInfo() foundation.IDictionary
	IsBroadcasting() bool
	SetIsBroadcasting(value bool)
	IsPaused() bool
	SetIsPaused(value bool)
	// methods:
	FinishBroadcastWithHandler(handler func(unsafe.Pointer))
	PauseBroadcast()
	ResumeBroadcast()
	StartBroadcastWithHandler(handler func(unsafe.Pointer))
}

// An object containing methods for starting and controlling a broadcast.


// An object containing methods for starting and controlling a broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController
type RPBroadcastController struct {
	objectivec.Object
}

// RPBroadcastControllerFrom constructs a [RPBroadcastController] from an unsafe.Pointer.
//
// An object containing methods for starting and controlling a broadcast.
func RPBroadcastControllerFrom(ptr unsafe.Pointer) RPBroadcastController {
	return RPBroadcastController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastControllerClass) Alloc() RPBroadcastController {
	rv := objc.Send[RPBroadcastController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastControllerClass) New() RPBroadcastController {
	rv := objc.Send[RPBroadcastController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastController) Init() RPBroadcastController {
	rv := objc.Send[RPBroadcastController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastController) Autorelease() RPBroadcastController {
	rv := objc.Send[RPBroadcastController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastController creates a new RPBroadcastController instance.
func NewRPBroadcastController() RPBroadcastController {
	return getRPBroadcastControllerClass().New()
}



// Stops the current broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/finishBroadcast(handler:)
func (r_ RPBroadcastController) FinishBroadcastWithHandler(handler func(unsafe.Pointer)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("finishBroadcastWithHandler:"), handler)
}


// Pauses the current broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/pauseBroadcast()
func (r_ RPBroadcastController) PauseBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("pauseBroadcast"))
}


// Resumes a paused broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/resumeBroadcast()
func (r_ RPBroadcastController) ResumeBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("resumeBroadcast"))
}


// Starts a broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/startBroadcast(handler:)
func (r_ RPBroadcastController) StartBroadcastWithHandler(handler func(unsafe.Pointer)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startBroadcastWithHandler:"), handler)
}


// A URL that redirects users to an ongoing or completed broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/broadcastURL
func (r_ RPBroadcastController) BroadcastURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("broadcastURL"))
	return rv
}


// The delegate for the broadcast controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the broadcast controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isBroadcasting
func (r_ RPBroadcastController) Broadcasting() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("broadcasting"))
	return rv
}


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isPaused
func (r_ RPBroadcastController) Paused() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("paused"))
	return rv
}


// Information updated by the service during a broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/serviceInfo
func (r_ RPBroadcastController) ServiceInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("serviceInfo"))
	return rv
}


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/isbroadcasting
func (r_ RPBroadcastController) IsBroadcasting() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isBroadcasting"))
	return rv
}


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/isbroadcasting
func (r_ RPBroadcastController) SetIsBroadcasting(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsBroadcasting:"), value)
}


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/ispaused
func (r_ RPBroadcastController) IsPaused() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/ispaused
func (r_ RPBroadcastController) SetIsPaused(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsPaused:"), value)
}


