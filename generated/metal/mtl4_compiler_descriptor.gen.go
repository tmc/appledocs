// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CompilerDescriptor */


/* debug [class_header]: Header for MTL4CompilerDescriptor */
// The class instance for the [MTL4CompilerDescriptor] class.
var (
	MTL4CompilerDescriptorClass     _MTL4CompilerDescriptorClass
	MTL4CompilerDescriptorClassOnce sync.Once
)

func getMTL4CompilerDescriptorClass() _MTL4CompilerDescriptorClass {
	MTL4CompilerDescriptorClassOnce.Do(func() {
		MTL4CompilerDescriptorClass = _MTL4CompilerDescriptorClass{objc.GetClass("MTL4CompilerDescriptor")}
	})
	return MTL4CompilerDescriptorClass
}

type _MTL4CompilerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CompilerDescriptor */
// An interface definition for the [MTL4CompilerDescriptor] class.
type IMTL4CompilerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CompilerDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	PipelineDataSetSerializer() unsafe.Pointer
	SetPipelineDataSetSerializer(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CompilerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CompilerDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CompilerDescriptorClass) Alloc() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CompilerDescriptorClass) New() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CompilerDescriptor) Init() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CompilerDescriptor) Autorelease() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CompilerDescriptor creates a new MTL4CompilerDescriptor instance.
func NewMTL4CompilerDescriptor() MTL4CompilerDescriptor {
	return getMTL4CompilerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CompilerDescriptor */
// Groups together properties for creating a compiler context.


// Groups together properties for creating a compiler context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor
type MTL4CompilerDescriptor struct {
	objectivec.Object
}

// MTL4CompilerDescriptorFrom constructs a [MTL4CompilerDescriptor] from an unsafe.Pointer.
//
// Groups together properties for creating a compiler context.
func MTL4CompilerDescriptorFrom(ptr unsafe.Pointer) MTL4CompilerDescriptor {
	return MTL4CompilerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CompilerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CompilerDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CompilerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CompilerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CompilerDescriptor */

// Assigns an optional descriptor label to the compiler for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/label
func (m_ MTL4CompilerDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional descriptor label to the compiler for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/label
func (m_ MTL4CompilerDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Assigns a pipeline data set serializer into which this compiler stores data for all pipelines it creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/pipelineDataSetSerializer
func (m_ MTL4CompilerDescriptor) PipelineDataSetSerializer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pipelineDataSetSerializer"))
	return rv
}/* debug [instance_properties/getter]: pipelineDataSetSerializer */


// Assigns a pipeline data set serializer into which this compiler stores data for all pipelines it creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/pipelineDataSetSerializer
func (m_ MTL4CompilerDescriptor) SetPipelineDataSetSerializer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPipelineDataSetSerializer:"), value)
}/* debug [instance_properties/setter]: pipelineDataSetSerializer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CompilerDescriptor */



