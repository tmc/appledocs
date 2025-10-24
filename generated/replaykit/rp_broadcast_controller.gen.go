// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class RPBroadcastController */


/* debug [class_header]: Header for RPBroadcastController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastController */
// An interface definition for the [RPBroadcastController] class.
type IRPBroadcastController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RPBroadcastController */
	// properties:
	BroadcastURL() objc.IObject /* cross-framework: NSURL */
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Broadcasting() bool
	Paused() bool
	ServiceInfo() foundation.IDictionary
	IsBroadcasting() bool
	SetIsBroadcasting(value bool)
	IsPaused() bool
	SetIsPaused(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastController */
	// methods:
	FinishBroadcastWithHandler(handler unsafe.Pointer)
	PauseBroadcast()
	ResumeBroadcast()
	StartBroadcastWithHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastController */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastControllerClass) Alloc() RPBroadcastController {
	rv := objc.Send[RPBroadcastController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastController */

// Stops the current broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/finishBroadcast(handler:)
func (r_ RPBroadcastController) FinishBroadcastWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("finishBroadcastWithHandler:"), handler)
}/* debug [instance_methods/method]: FinishBroadcastWithHandler */


// Pauses the current broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/pauseBroadcast()
func (r_ RPBroadcastController) PauseBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("pauseBroadcast"))
}/* debug [instance_methods/method]: PauseBroadcast */


// Resumes a paused broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/resumeBroadcast()
func (r_ RPBroadcastController) ResumeBroadcast() {
	objc.Send[objc.ID](r_.ID, objc.Sel("resumeBroadcast"))
}/* debug [instance_methods/method]: ResumeBroadcast */


// Starts a broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/startBroadcast(handler:)
func (r_ RPBroadcastController) StartBroadcastWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startBroadcastWithHandler:"), handler)
}/* debug [instance_methods/method]: StartBroadcastWithHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastController */

// A URL that redirects users to an ongoing or completed broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/broadcastURL
func (r_ RPBroadcastController) BroadcastURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("broadcastURL"))
	return rv
}/* debug [instance_properties/getter]: broadcastURL */


// The delegate for the broadcast controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the broadcast controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/delegate
func (r_ RPBroadcastController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isBroadcasting
func (r_ RPBroadcastController) Broadcasting() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("broadcasting"))
	return rv
}/* debug [instance_properties/getter]: broadcasting */


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/isPaused
func (r_ RPBroadcastController) Paused() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// Information updated by the service during a broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/serviceInfo
func (r_ RPBroadcastController) ServiceInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("serviceInfo"))
	return rv
}/* debug [instance_properties/getter]: serviceInfo */


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/isbroadcasting
func (r_ RPBroadcastController) IsBroadcasting() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isBroadcasting"))
	return rv
}/* debug [instance_properties/getter]: isBroadcasting */


// A Boolean value indicating whether the controller is broadcasting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/isbroadcasting
func (r_ RPBroadcastController) SetIsBroadcasting(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsBroadcasting:"), value)
}/* debug [instance_properties/setter]: isBroadcasting */


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/ispaused
func (r_ RPBroadcastController) IsPaused() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value indicating whether the broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpbroadcastcontroller/ispaused
func (r_ RPBroadcastController) SetIsPaused(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastController */


