// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class RPBroadcastHandler */


/* debug [class_header]: Header for RPBroadcastHandler */
// The class instance for the [RPBroadcastHandler] class.
var (
	RPBroadcastHandlerClass     _RPBroadcastHandlerClass
	RPBroadcastHandlerClassOnce sync.Once
)

func getRPBroadcastHandlerClass() _RPBroadcastHandlerClass {
	RPBroadcastHandlerClassOnce.Do(func() {
		RPBroadcastHandlerClass = _RPBroadcastHandlerClass{objc.GetClass("RPBroadcastHandler")}
	})
	return RPBroadcastHandlerClass
}

type _RPBroadcastHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastHandler */
// An interface definition for the [RPBroadcastHandler] class.
type IRPBroadcastHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RPBroadcastHandler */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastHandler */
	// methods:
	UpdateBroadcastURL(broadcastURL objc.IObject /* cross-framework: NSURL */)
	UpdateServiceInfo(serviceInfo foundation.IDictionary)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastHandler */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastHandlerClass) Alloc() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPBroadcastHandlerClass) New() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastHandler) Init() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastHandler) Autorelease() RPBroadcastHandler {
	rv := objc.Send[RPBroadcastHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastHandler creates a new RPBroadcastHandler instance.
func NewRPBroadcastHandler() RPBroadcastHandler {
	return getRPBroadcastHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastHandler */
// An object that sends messages to the broadcasting app.


// An object that sends messages to the broadcasting app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler
type RPBroadcastHandler struct {
	objectivec.Object
}

// RPBroadcastHandlerFrom constructs a [RPBroadcastHandler] from an unsafe.Pointer.
//
// An object that sends messages to the broadcasting app.
func RPBroadcastHandlerFrom(ptr unsafe.Pointer) RPBroadcastHandler {
	return RPBroadcastHandler{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastHandler */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastHandler */

// Sends the current broadcast URL to the broadcast controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler/updateBroadcast(_:)
func (r_ RPBroadcastHandler) UpdateBroadcastURL(broadcastURL objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateBroadcastURL:"), broadcastURL)
}/* debug [instance_methods/method]: UpdateBroadcastURL */


// Sends information about the current broadcast to the broadcasting app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastHandler/updateServiceInfo(_:)
func (r_ RPBroadcastHandler) UpdateServiceInfo(serviceInfo foundation.IDictionary) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateServiceInfo:"), serviceInfo)
}/* debug [instance_methods/method]: UpdateServiceInfo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastHandler */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastHandler */



