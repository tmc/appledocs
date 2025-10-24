// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4PipelineDescriptor */


/* debug [class_header]: Header for MTL4PipelineDescriptor */
// The class instance for the [MTL4PipelineDescriptor] class.
var (
	MTL4PipelineDescriptorClass     _MTL4PipelineDescriptorClass
	MTL4PipelineDescriptorClassOnce sync.Once
)

func getMTL4PipelineDescriptorClass() _MTL4PipelineDescriptorClass {
	MTL4PipelineDescriptorClassOnce.Do(func() {
		MTL4PipelineDescriptorClass = _MTL4PipelineDescriptorClass{objc.GetClass("MTL4PipelineDescriptor")}
	})
	return MTL4PipelineDescriptorClass
}

type _MTL4PipelineDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4PipelineDescriptor */
// An interface definition for the [MTL4PipelineDescriptor] class.
type IMTL4PipelineDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4PipelineDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Options() IMTL4PipelineOptions
	SetOptions(value IMTL4PipelineOptions)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4PipelineDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4PipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineDescriptorClass) Alloc() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineDescriptorClass) New() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineDescriptor) Init() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineDescriptor) Autorelease() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDescriptor creates a new MTL4PipelineDescriptor instance.
func NewMTL4PipelineDescriptor() MTL4PipelineDescriptor {
	return getMTL4PipelineDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4PipelineDescriptor */
// Base type for descriptors you use for building pipeline state objects.


// Base type for descriptors you use for building pipeline state objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor
type MTL4PipelineDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDescriptorFrom constructs a [MTL4PipelineDescriptor] from an unsafe.Pointer.
//
// Base type for descriptors you use for building pipeline state objects.
func MTL4PipelineDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineDescriptor {
	return MTL4PipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4PipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4PipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4PipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4PipelineDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4PipelineDescriptor */

// Assigns an optional string that uniquely identifies a pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/label
func (m_ MTL4PipelineDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional string that uniquely identifies a pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/label
func (m_ MTL4PipelineDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Provides compile-time options when you build the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/options
func (m_ MTL4PipelineDescriptor) Options() IMTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Provides compile-time options when you build the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/options
func (m_ MTL4PipelineDescriptor) SetOptions(value IMTL4PipelineOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4PipelineDescriptor */



