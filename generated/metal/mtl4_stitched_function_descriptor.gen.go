// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4StitchedFunctionDescriptor */


/* debug [class_header]: Header for MTL4StitchedFunctionDescriptor */
// The class instance for the [MTL4StitchedFunctionDescriptor] class.
var (
	MTL4StitchedFunctionDescriptorClass     _MTL4StitchedFunctionDescriptorClass
	MTL4StitchedFunctionDescriptorClassOnce sync.Once
)

func getMTL4StitchedFunctionDescriptorClass() _MTL4StitchedFunctionDescriptorClass {
	MTL4StitchedFunctionDescriptorClassOnce.Do(func() {
		MTL4StitchedFunctionDescriptorClass = _MTL4StitchedFunctionDescriptorClass{objc.GetClass("MTL4StitchedFunctionDescriptor")}
	})
	return MTL4StitchedFunctionDescriptorClass
}

type _MTL4StitchedFunctionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4StitchedFunctionDescriptor */
// An interface definition for the [MTL4StitchedFunctionDescriptor] class.
type IMTL4StitchedFunctionDescriptor interface {
	IMTL4FunctionDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4StitchedFunctionDescriptor */
	// properties:
	FunctionDescriptors() []MTL4FunctionDescriptor
	SetFunctionDescriptors(value []MTL4FunctionDescriptor)
	FunctionGraph() IMTLFunctionStitchingGraph
	SetFunctionGraph(value IMTLFunctionStitchingGraph)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4StitchedFunctionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4StitchedFunctionDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4StitchedFunctionDescriptorClass) Alloc() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4StitchedFunctionDescriptorClass) New() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4StitchedFunctionDescriptor) Init() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4StitchedFunctionDescriptor) Autorelease() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4StitchedFunctionDescriptor creates a new MTL4StitchedFunctionDescriptor instance.
func NewMTL4StitchedFunctionDescriptor() MTL4StitchedFunctionDescriptor {
	return getMTL4StitchedFunctionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4StitchedFunctionDescriptor */
// Groups together properties that describe a shader function suitable for stitching.


// Groups together properties that describe a shader function suitable for stitching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor
type MTL4StitchedFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4StitchedFunctionDescriptorFrom constructs a [MTL4StitchedFunctionDescriptor] from an unsafe.Pointer.
//
// Groups together properties that describe a shader function suitable for stitching.
func MTL4StitchedFunctionDescriptorFrom(ptr unsafe.Pointer) MTL4StitchedFunctionDescriptor {
	return MTL4StitchedFunctionDescriptor{
		MTL4FunctionDescriptor: MTL4FunctionDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4StitchedFunctionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4StitchedFunctionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4StitchedFunctionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4StitchedFunctionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4StitchedFunctionDescriptor */

// Configures an array of function descriptors with references to functions that contribute to the stitching process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionDescriptors
func (m_ MTL4StitchedFunctionDescriptor) FunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]MTL4FunctionDescriptor](m_.ID, objc.Sel("functionDescriptors"))
	return rv
}/* debug [instance_properties/getter]: functionDescriptors */


// Configures an array of function descriptors with references to functions that contribute to the stitching process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionDescriptors
func (m_ MTL4StitchedFunctionDescriptor) SetFunctionDescriptors(value []MTL4FunctionDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionDescriptors:"), nsArray)
}/* debug [instance_properties/setter]: functionDescriptors */


// Sets the graph representing how to stitch functions together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionGraph
func (m_ MTL4StitchedFunctionDescriptor) FunctionGraph() IMTLFunctionStitchingGraph {
	rv := objc.Send[FunctionStitchingGraph](m_.ID, objc.Sel("functionGraph"))
	return rv
}/* debug [instance_properties/getter]: functionGraph */


// Sets the graph representing how to stitch functions together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionGraph
func (m_ MTL4StitchedFunctionDescriptor) SetFunctionGraph(value IMTLFunctionStitchingGraph) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionGraph:"), value)
}/* debug [instance_properties/setter]: functionGraph */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4StitchedFunctionDescriptor */



