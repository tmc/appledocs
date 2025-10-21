// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	FinishBroadcastWithHandler(handler func(error objc.ID))
	PauseBroadcast()
	ResumeBroadcast()
	StartBroadcastWithHandler(handler func(error objc.ID))
}

// An object containing methods for starting and controlling a broadcast.
//
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
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/finishBroadcast(handler:)
func (r_ RPBroadcastController) FinishBroadcastWithHandler(handler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("finishBroadcastWithHandler:"), handler)
}

// Pauses the current broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/pauseBroadcast()
func (r_ RPBroadcastController) PauseBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("pauseBroadcast"))
}

// Resumes a paused broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/resumeBroadcast()
func (r_ RPBroadcastController) ResumeBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("resumeBroadcast"))
}

// Starts a broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/startBroadcast(handler:)
func (r_ RPBroadcastController) StartBroadcastWithHandler(handler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startBroadcastWithHandler:"), handler)
}

// The bundle ID for the selected broadcast service.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/broadcastExtensionBundleID
func (r_ RPBroadcastController) BroadcastExtensionBundleID() string {
	rv := objc.Send[string](r_.ID, objc.Sel("broadcastExtensionBundleID"))
	return rv
}

// A URL that redirects users to an ongoing or completed broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/broadcastURL
func (r_ RPBroadcastController) BroadcastURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("broadcastURL"))
	return rv
}

// The delegate for the broadcast controller.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the broadcast controller.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isBroadcasting
func (r_ RPBroadcastController) Broadcasting() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("broadcasting"))
	return rv
}

// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isPaused
func (r_ RPBroadcastController) Paused() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("paused"))
	return rv
}

// Information updated by the service during a broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/serviceInfo
func (r_ RPBroadcastController) ServiceInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("serviceInfo"))
	return rv
}



