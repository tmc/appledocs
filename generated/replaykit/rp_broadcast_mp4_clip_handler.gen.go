// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class RPBroadcastMP4ClipHandler */


/* debug [class_header]: Header for RPBroadcastMP4ClipHandler */
// The class instance for the [RPBroadcastMP4ClipHandler] class.
var (
	RPBroadcastMP4ClipHandlerClass     _RPBroadcastMP4ClipHandlerClass
	RPBroadcastMP4ClipHandlerClassOnce sync.Once
)

func getRPBroadcastMP4ClipHandlerClass() _RPBroadcastMP4ClipHandlerClass {
	RPBroadcastMP4ClipHandlerClassOnce.Do(func() {
		RPBroadcastMP4ClipHandlerClass = _RPBroadcastMP4ClipHandlerClass{objc.GetClass("RPBroadcastMP4ClipHandler")}
	})
	return RPBroadcastMP4ClipHandlerClass
}

type _RPBroadcastMP4ClipHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastMP4ClipHandler */
// An interface definition for the [RPBroadcastMP4ClipHandler] class.
type IRPBroadcastMP4ClipHandler interface {
	IRPBroadcastHandler
	
/* debug [class_interface_properties]: Properties for RPBroadcastMP4ClipHandler */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastMP4ClipHandler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastMP4ClipHandler */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastMP4ClipHandlerClass) Alloc() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPBroadcastMP4ClipHandlerClass) New() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastMP4ClipHandler) Init() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastMP4ClipHandler) Autorelease() RPBroadcastMP4ClipHandler {
	rv := objc.Send[RPBroadcastMP4ClipHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastMP4ClipHandler creates a new RPBroadcastMP4ClipHandler instance.
func NewRPBroadcastMP4ClipHandler() RPBroadcastMP4ClipHandler {
	return getRPBroadcastMP4ClipHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastMP4ClipHandler */
// An object that processes MP4 movie clips from ReplayKit.
//
// Subclass this class to handle movie clips as ReplayKit records them during the broadcast. The system calls when a movie clip is available for processing.


// An object that processes MP4 movie clips from ReplayKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastMP4ClipHandler
type RPBroadcastMP4ClipHandler struct {
	RPBroadcastHandler
}

// RPBroadcastMP4ClipHandlerFrom constructs a [RPBroadcastMP4ClipHandler] from an unsafe.Pointer.
//
// An object that processes MP4 movie clips from ReplayKit.
func RPBroadcastMP4ClipHandlerFrom(ptr unsafe.Pointer) RPBroadcastMP4ClipHandler {
	return RPBroadcastMP4ClipHandler{
		RPBroadcastHandler: RPBroadcastHandlerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastMP4ClipHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastMP4ClipHandler */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastMP4ClipHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastMP4ClipHandler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastMP4ClipHandler */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastMP4ClipHandler */


