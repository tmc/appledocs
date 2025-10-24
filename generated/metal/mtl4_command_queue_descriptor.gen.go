// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CommandQueueDescriptor */


/* debug [class_header]: Header for MTL4CommandQueueDescriptor */
// The class instance for the [MTL4CommandQueueDescriptor] class.
var (
	MTL4CommandQueueDescriptorClass     _MTL4CommandQueueDescriptorClass
	MTL4CommandQueueDescriptorClassOnce sync.Once
)

func getMTL4CommandQueueDescriptorClass() _MTL4CommandQueueDescriptorClass {
	MTL4CommandQueueDescriptorClassOnce.Do(func() {
		MTL4CommandQueueDescriptorClass = _MTL4CommandQueueDescriptorClass{objc.GetClass("MTL4CommandQueueDescriptor")}
	})
	return MTL4CommandQueueDescriptorClass
}

type _MTL4CommandQueueDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CommandQueueDescriptor */
// An interface definition for the [MTL4CommandQueueDescriptor] class.
type IMTL4CommandQueueDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CommandQueueDescriptor */
	// properties:
	FeedbackQueue() objectivec.IObject
	SetFeedbackQueue(value objectivec.IObject)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CommandQueueDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CommandQueueDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandQueueDescriptorClass) Alloc() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CommandQueueDescriptorClass) New() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandQueueDescriptor) Init() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandQueueDescriptor) Autorelease() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandQueueDescriptor creates a new MTL4CommandQueueDescriptor instance.
func NewMTL4CommandQueueDescriptor() MTL4CommandQueueDescriptor {
	return getMTL4CommandQueueDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CommandQueueDescriptor */
// Groups together parameters for the creation of a new command queue.


// Groups together parameters for the creation of a new command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor
type MTL4CommandQueueDescriptor struct {
	objectivec.Object
}

// MTL4CommandQueueDescriptorFrom constructs a [MTL4CommandQueueDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for the creation of a new command queue.
func MTL4CommandQueueDescriptorFrom(ptr unsafe.Pointer) MTL4CommandQueueDescriptor {
	return MTL4CommandQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CommandQueueDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CommandQueueDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CommandQueueDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CommandQueueDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CommandQueueDescriptor */

// Assigns a dispatch queue to which Metal submits feedback notification blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/feedbackQueue
func (m_ MTL4CommandQueueDescriptor) FeedbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("feedbackQueue"))
	return rv
}/* debug [instance_properties/getter]: feedbackQueue */


// Assigns a dispatch queue to which Metal submits feedback notification blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/feedbackQueue
func (m_ MTL4CommandQueueDescriptor) SetFeedbackQueue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeedbackQueue:"), value)
}/* debug [instance_properties/setter]: feedbackQueue */


// Assigns an optional label to the command queue instance for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/label
func (m_ MTL4CommandQueueDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional label to the command queue instance for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/label
func (m_ MTL4CommandQueueDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommandQueueDescriptor) MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTL4CommandQueueErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CommandQueueDescriptor */



