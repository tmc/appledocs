// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterMessages */


/* debug [class_header]: Header for MTRBaseClusterMessages */
// The class instance for the [MTRBaseClusterMessages] class.
var (
	MTRBaseClusterMessagesClass     _MTRBaseClusterMessagesClass
	MTRBaseClusterMessagesClassOnce sync.Once
)

func getMTRBaseClusterMessagesClass() _MTRBaseClusterMessagesClass {
	MTRBaseClusterMessagesClassOnce.Do(func() {
		MTRBaseClusterMessagesClass = _MTRBaseClusterMessagesClass{objc.GetClass("MTRBaseClusterMessages")}
	})
	return MTRBaseClusterMessagesClass
}

type _MTRBaseClusterMessagesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterMessages */
// An interface definition for the [MTRBaseClusterMessages] class.
type IMTRBaseClusterMessages interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterMessages */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterMessages */
	// methods:
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterMessages */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMessagesClass) Alloc() MTRBaseClusterMessages {
	rv := objc.Send[MTRBaseClusterMessages](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterMessagesClass) New() MTRBaseClusterMessages {
	rv := objc.Send[MTRBaseClusterMessages](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMessages) Init() MTRBaseClusterMessages {
	rv := objc.Send[MTRBaseClusterMessages](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMessages) Autorelease() MTRBaseClusterMessages {
	rv := objc.Send[MTRBaseClusterMessages](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMessages creates a new MTRBaseClusterMessages instance.
func NewMTRBaseClusterMessages() MTRBaseClusterMessages {
	return getMTRBaseClusterMessagesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterMessages */
// Cluster Messages
//
// This cluster provides an interface for passing messages to be presented by a device.


// Cluster Messages
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMessages
type MTRBaseClusterMessages struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMessagesFrom constructs a [MTRBaseClusterMessages] from an unsafe.Pointer.
//
// Cluster Messages
func MTRBaseClusterMessagesFrom(ptr unsafe.Pointer) MTRBaseClusterMessages {
	return MTRBaseClusterMessages{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterMessages *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterMessages */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterMessages */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterMessages */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMessages/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterMessages) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterMessages */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterMessages */



