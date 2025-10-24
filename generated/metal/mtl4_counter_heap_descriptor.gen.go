// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CounterHeapDescriptor */


/* debug [class_header]: Header for MTL4CounterHeapDescriptor */
// The class instance for the [MTL4CounterHeapDescriptor] class.
var (
	MTL4CounterHeapDescriptorClass     _MTL4CounterHeapDescriptorClass
	MTL4CounterHeapDescriptorClassOnce sync.Once
)

func getMTL4CounterHeapDescriptorClass() _MTL4CounterHeapDescriptorClass {
	MTL4CounterHeapDescriptorClassOnce.Do(func() {
		MTL4CounterHeapDescriptorClass = _MTL4CounterHeapDescriptorClass{objc.GetClass("MTL4CounterHeapDescriptor")}
	})
	return MTL4CounterHeapDescriptorClass
}

type _MTL4CounterHeapDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CounterHeapDescriptor */
// An interface definition for the [MTL4CounterHeapDescriptor] class.
type IMTL4CounterHeapDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CounterHeapDescriptor */
	// properties:
	Count() uint
	SetCount(value uint)
	Type() MTL4CounterHeapType
	SetType(value MTL4CounterHeapType)
	MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CounterHeapDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CounterHeapDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CounterHeapDescriptorClass) Alloc() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CounterHeapDescriptorClass) New() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CounterHeapDescriptor) Init() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CounterHeapDescriptor) Autorelease() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CounterHeapDescriptor creates a new MTL4CounterHeapDescriptor instance.
func NewMTL4CounterHeapDescriptor() MTL4CounterHeapDescriptor {
	return getMTL4CounterHeapDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CounterHeapDescriptor */
// Groups together parameters for configuring a counter heap object at creation time.


// Groups together parameters for configuring a counter heap object at creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor
type MTL4CounterHeapDescriptor struct {
	objectivec.Object
}

// MTL4CounterHeapDescriptorFrom constructs a [MTL4CounterHeapDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for configuring a counter heap object at creation time.
func MTL4CounterHeapDescriptorFrom(ptr unsafe.Pointer) MTL4CounterHeapDescriptor {
	return MTL4CounterHeapDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CounterHeapDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CounterHeapDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CounterHeapDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CounterHeapDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CounterHeapDescriptor */

// Assigns the number of entries in the heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/count
func (m_ MTL4CounterHeapDescriptor) Count() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// Assigns the number of entries in the heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/count
func (m_ MTL4CounterHeapDescriptor) SetCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}/* debug [instance_properties/setter]: count */


// Assigns the type of data that the heap contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/type
func (m_ MTL4CounterHeapDescriptor) Type() MTL4CounterHeapType {
	rv := objc.Send[MTL4CounterHeapType](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// Assigns the type of data that the heap contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/type
func (m_ MTL4CounterHeapDescriptor) SetType(value MTL4CounterHeapType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CounterHeapDescriptor) MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTL4CommandQueueErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CounterHeapDescriptor */



