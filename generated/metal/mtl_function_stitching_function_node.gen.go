// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionStitchingFunctionNode */


/* debug [class_header]: Header for MTLFunctionStitchingFunctionNode */
// The class instance for the [FunctionStitchingFunctionNode] class.
var (
	FunctionStitchingFunctionNodeClass     _FunctionStitchingFunctionNodeClass
	FunctionStitchingFunctionNodeClassOnce sync.Once
)

func getFunctionStitchingFunctionNodeClass() _FunctionStitchingFunctionNodeClass {
	FunctionStitchingFunctionNodeClassOnce.Do(func() {
		FunctionStitchingFunctionNodeClass = _FunctionStitchingFunctionNodeClass{objc.GetClass("MTLFunctionStitchingFunctionNode")}
	})
	return FunctionStitchingFunctionNodeClass
}

type _FunctionStitchingFunctionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionStitchingFunctionNode */
// An interface definition for the [FunctionStitchingFunctionNode] class.
type IFunctionStitchingFunctionNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionStitchingFunctionNode */
	// properties:
	Arguments() []objc.ID
	SetArguments(value []objc.ID)
	ControlDependencies() []FunctionStitchingFunctionNode
	SetControlDependencies(value []FunctionStitchingFunctionNode)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionStitchingFunctionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionStitchingFunctionNode */
// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingFunctionNodeClass) Alloc() FunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionStitchingFunctionNodeClass) New() FunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionStitchingFunctionNode) Init() FunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionStitchingFunctionNode) Autorelease() FunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionStitchingFunctionNode creates a new FunctionStitchingFunctionNode instance.
func NewFunctionStitchingFunctionNode() FunctionStitchingFunctionNode {
	return getFunctionStitchingFunctionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionStitchingFunctionNode */
// A call graph node that describes a function call and its inputs.
//
// When the Metal device object evaluates the function graph to compile the stitched function, it evaluates the nodes stored in the property that it hasn’t already evaluated, and then calls the function specified by to generate the node’s output. If the function has side effects on the input data, use the property on other nodes to specify whether the Metal device object must evaluate this node first.


// A call graph node that describes a function call and its inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode
type FunctionStitchingFunctionNode struct {
	objectivec.Object
}

// FunctionStitchingFunctionNodeFrom constructs a [FunctionStitchingFunctionNode] from an unsafe.Pointer.
//
// A call graph node that describes a function call and its inputs.
func FunctionStitchingFunctionNodeFrom(ptr unsafe.Pointer) FunctionStitchingFunctionNode {
	return FunctionStitchingFunctionNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionStitchingFunctionNode */

// Creates a new function node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/init(name:arguments:controlDependencies:)
func NewFunctionStitchingFunctionNodeWithNameArgumentsControlDependencies(name objc.IObject /* cross-framework: NSString */, arguments []objc.ID, controlDependencies []FunctionStitchingFunctionNode) FunctionStitchingFunctionNode {
	instance := getFunctionStitchingFunctionNodeClass().Alloc()
	rv := objc.Send[FunctionStitchingFunctionNode](instance.ID, objc.Sel("initWithName:arguments:controlDependencies:"), name, arguments, controlDependencies)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFunctionStitchingFunctionNodeWithNameArgumentsControlDependencies */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionStitchingFunctionNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionStitchingFunctionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionStitchingFunctionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionStitchingFunctionNode */

// An ordered list of the nodes that provide the function’s arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/arguments
func (f_ FunctionStitchingFunctionNode) Arguments() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// An ordered list of the nodes that provide the function’s arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/arguments
func (f_ FunctionStitchingFunctionNode) SetArguments(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setArguments:"), nsArray)
}/* debug [instance_properties/setter]: arguments */


// The list of nodes that must execute before executing the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/controlDependencies
func (f_ FunctionStitchingFunctionNode) ControlDependencies() []FunctionStitchingFunctionNode {
	rv := objc.Send[[]FunctionStitchingFunctionNode](f_.ID, objc.Sel("controlDependencies"))
	return rv
}/* debug [instance_properties/getter]: controlDependencies */


// The list of nodes that must execute before executing the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/controlDependencies
func (f_ FunctionStitchingFunctionNode) SetControlDependencies(value []FunctionStitchingFunctionNode) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setControlDependencies:"), nsArray)
}/* debug [instance_properties/setter]: controlDependencies */


// The name of the function to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/name
func (f_ FunctionStitchingFunctionNode) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the function to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/name
func (f_ FunctionStitchingFunctionNode) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionStitchingFunctionNode */


