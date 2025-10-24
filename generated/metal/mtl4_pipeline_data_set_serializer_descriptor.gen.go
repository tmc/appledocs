// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4PipelineDataSetSerializerDescriptor */


/* debug [class_header]: Header for MTL4PipelineDataSetSerializerDescriptor */
// The class instance for the [MTL4PipelineDataSetSerializerDescriptor] class.
var (
	MTL4PipelineDataSetSerializerDescriptorClass     _MTL4PipelineDataSetSerializerDescriptorClass
	MTL4PipelineDataSetSerializerDescriptorClassOnce sync.Once
)

func getMTL4PipelineDataSetSerializerDescriptorClass() _MTL4PipelineDataSetSerializerDescriptorClass {
	MTL4PipelineDataSetSerializerDescriptorClassOnce.Do(func() {
		MTL4PipelineDataSetSerializerDescriptorClass = _MTL4PipelineDataSetSerializerDescriptorClass{objc.GetClass("MTL4PipelineDataSetSerializerDescriptor")}
	})
	return MTL4PipelineDataSetSerializerDescriptorClass
}

type _MTL4PipelineDataSetSerializerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4PipelineDataSetSerializerDescriptor */
// An interface definition for the [MTL4PipelineDataSetSerializerDescriptor] class.
type IMTL4PipelineDataSetSerializerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4PipelineDataSetSerializerDescriptor */
	// properties:
	Configuration() MTL4PipelineDataSetSerializerConfiguration
	SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4PipelineDataSetSerializerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4PipelineDataSetSerializerDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineDataSetSerializerDescriptorClass) Alloc() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineDataSetSerializerDescriptorClass) New() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineDataSetSerializerDescriptor) Init() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineDataSetSerializerDescriptor) Autorelease() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDataSetSerializerDescriptor creates a new MTL4PipelineDataSetSerializerDescriptor instance.
func NewMTL4PipelineDataSetSerializerDescriptor() MTL4PipelineDataSetSerializerDescriptor {
	return getMTL4PipelineDataSetSerializerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4PipelineDataSetSerializerDescriptor */
// Groups together properties to create a pipeline data set serializer.


// Groups together properties to create a pipeline data set serializer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor
type MTL4PipelineDataSetSerializerDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDataSetSerializerDescriptorFrom constructs a [MTL4PipelineDataSetSerializerDescriptor] from an unsafe.Pointer.
//
// Groups together properties to create a pipeline data set serializer.
func MTL4PipelineDataSetSerializerDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineDataSetSerializerDescriptor {
	return MTL4PipelineDataSetSerializerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4PipelineDataSetSerializerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4PipelineDataSetSerializerDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4PipelineDataSetSerializerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4PipelineDataSetSerializerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4PipelineDataSetSerializerDescriptor */

// Specifies the configuration of the serialization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor/configuration
func (m_ MTL4PipelineDataSetSerializerDescriptor) Configuration() MTL4PipelineDataSetSerializerConfiguration {
	rv := objc.Send[MTL4PipelineDataSetSerializerConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// Specifies the configuration of the serialization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor/configuration
func (m_ MTL4PipelineDataSetSerializerDescriptor) SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4PipelineDataSetSerializerDescriptor */



