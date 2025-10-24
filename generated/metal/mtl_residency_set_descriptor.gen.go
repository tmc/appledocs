// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLResidencySetDescriptor */


/* debug [class_header]: Header for MTLResidencySetDescriptor */
// The class instance for the [ResidencySetDescriptor] class.
var (
	ResidencySetDescriptorClass     _ResidencySetDescriptorClass
	ResidencySetDescriptorClassOnce sync.Once
)

func getResidencySetDescriptorClass() _ResidencySetDescriptorClass {
	ResidencySetDescriptorClassOnce.Do(func() {
		ResidencySetDescriptorClass = _ResidencySetDescriptorClass{objc.GetClass("MTLResidencySetDescriptor")}
	})
	return ResidencySetDescriptorClass
}

type _ResidencySetDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResidencySetDescriptor */
// An interface definition for the [ResidencySetDescriptor] class.
type IResidencySetDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResidencySetDescriptor */
	// properties:
	InitialCapacity() uint
	SetInitialCapacity(value uint)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResidencySetDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResidencySetDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _ResidencySetDescriptorClass) Alloc() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResidencySetDescriptorClass) New() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResidencySetDescriptor) Init() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResidencySetDescriptor) Autorelease() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResidencySetDescriptor creates a new ResidencySetDescriptor instance.
func NewResidencySetDescriptor() ResidencySetDescriptor {
	return getResidencySetDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResidencySetDescriptor */
// A configuration that customizes the behavior for a residency set.
//
// Make an by creating and configuring an instance and pass it to the method of an instance. See for more information.


// A configuration that customizes the behavior for a residency set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor
type ResidencySetDescriptor struct {
	objectivec.Object
}

// ResidencySetDescriptorFrom constructs a [ResidencySetDescriptor] from an unsafe.Pointer.
//
// A configuration that customizes the behavior for a residency set.
func ResidencySetDescriptorFrom(ptr unsafe.Pointer) ResidencySetDescriptor {
	return ResidencySetDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResidencySetDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResidencySetDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResidencySetDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResidencySetDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResidencySetDescriptor */

// The number of allocations a new residency set can store without reallocating memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/initialCapacity
func (r_ ResidencySetDescriptor) InitialCapacity() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("initialCapacity"))
	return rv
}/* debug [instance_properties/getter]: initialCapacity */


// The number of allocations a new residency set can store without reallocating memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/initialCapacity
func (r_ ResidencySetDescriptor) SetInitialCapacity(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInitialCapacity:"), value)
}/* debug [instance_properties/setter]: initialCapacity */


// An optional name that can help you identify a residency set you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/label
func (r_ ResidencySetDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// An optional name that can help you identify a residency set you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/label
func (r_ ResidencySetDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLResidencySetDescriptor */



