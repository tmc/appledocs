// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterContentAppObserver */


/* debug [class_header]: Header for MTRClusterContentAppObserver */
// The class instance for the [MTRClusterContentAppObserver] class.
var (
	MTRClusterContentAppObserverClass     _MTRClusterContentAppObserverClass
	MTRClusterContentAppObserverClassOnce sync.Once
)

func getMTRClusterContentAppObserverClass() _MTRClusterContentAppObserverClass {
	MTRClusterContentAppObserverClassOnce.Do(func() {
		MTRClusterContentAppObserverClass = _MTRClusterContentAppObserverClass{objc.GetClass("MTRClusterContentAppObserver")}
	})
	return MTRClusterContentAppObserverClass
}

type _MTRClusterContentAppObserverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterContentAppObserver */
// An interface definition for the [MTRClusterContentAppObserver] class.
type IMTRClusterContentAppObserver interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterContentAppObserver */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterContentAppObserver */
	// methods:
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterContentAppObserver */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterContentAppObserverClass) Alloc() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterContentAppObserverClass) New() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterContentAppObserver) Init() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterContentAppObserver) Autorelease() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterContentAppObserver creates a new MTRClusterContentAppObserver instance.
func NewMTRClusterContentAppObserver() MTRClusterContentAppObserver {
	return getMTRClusterContentAppObserverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterContentAppObserver */
// Cluster Content App Observer This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.


// Cluster Content App Observer This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver
type MTRClusterContentAppObserver struct {
	MTRGenericCluster
}

// MTRClusterContentAppObserverFrom constructs a [MTRClusterContentAppObserver] from an unsafe.Pointer.
//
// Cluster Content App Observer This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.
func MTRClusterContentAppObserverFrom(ptr unsafe.Pointer) MTRClusterContentAppObserver {
	return MTRClusterContentAppObserver{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterContentAppObserver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterContentAppObserver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterContentAppObserver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterContentAppObserver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterContentAppObserver */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterContentAppObserver */



