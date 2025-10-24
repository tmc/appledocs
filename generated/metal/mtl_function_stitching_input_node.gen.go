// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionStitchingInputNode */


/* debug [class_header]: Header for MTLFunctionStitchingInputNode */
// The class instance for the [FunctionStitchingInputNode] class.
var (
	FunctionStitchingInputNodeClass     _FunctionStitchingInputNodeClass
	FunctionStitchingInputNodeClassOnce sync.Once
)

func getFunctionStitchingInputNodeClass() _FunctionStitchingInputNodeClass {
	FunctionStitchingInputNodeClassOnce.Do(func() {
		FunctionStitchingInputNodeClass = _FunctionStitchingInputNodeClass{objc.GetClass("MTLFunctionStitchingInputNode")}
	})
	return FunctionStitchingInputNodeClass
}

type _FunctionStitchingInputNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionStitchingInputNode */
// An interface definition for the [FunctionStitchingInputNode] class.
type IFunctionStitchingInputNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionStitchingInputNode */
	// properties:
	ArgumentIndex() uint
	SetArgumentIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionStitchingInputNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionStitchingInputNode */
// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingInputNodeClass) Alloc() FunctionStitchingInputNode {
	rv := objc.Send[FunctionStitchingInputNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionStitchingInputNodeClass) New() FunctionStitchingInputNode {
	rv := objc.Send[FunctionStitchingInputNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionStitchingInputNode) Init() FunctionStitchingInputNode {
	rv := objc.Send[FunctionStitchingInputNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionStitchingInputNode) Autorelease() FunctionStitchingInputNode {
	rv := objc.Send[FunctionStitchingInputNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionStitchingInputNode creates a new FunctionStitchingInputNode instance.
func NewFunctionStitchingInputNode() FunctionStitchingInputNode {
	return getFunctionStitchingInputNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionStitchingInputNode */
// A call graph node that describes an input to the call graph.
//
// An input node contains data from one of the stitched function’s parameters. The output data type of an input node has the same type as the matching parameter.


// A call graph node that describes an input to the call graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode
type FunctionStitchingInputNode struct {
	objectivec.Object
}

// FunctionStitchingInputNodeFrom constructs a [FunctionStitchingInputNode] from an unsafe.Pointer.
//
// A call graph node that describes an input to the call graph.
func FunctionStitchingInputNodeFrom(ptr unsafe.Pointer) FunctionStitchingInputNode {
	return FunctionStitchingInputNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionStitchingInputNode */

// Creates a new input node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/init(argumentIndex:)
func NewFunctionStitchingInputNodeWithArgumentIndex(argument uint) FunctionStitchingInputNode {
	instance := getFunctionStitchingInputNodeClass().Alloc()
	rv := objc.Send[FunctionStitchingInputNode](instance.ID, objc.Sel("initWithArgumentIndex:"), argument)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFunctionStitchingInputNodeWithArgumentIndex */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionStitchingInputNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionStitchingInputNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionStitchingInputNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionStitchingInputNode */

// The index in the command’s buffer argument table that declares which data to read for this input node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/argumentIndex
func (f_ FunctionStitchingInputNode) ArgumentIndex() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("argumentIndex"))
	return rv
}/* debug [instance_properties/getter]: argumentIndex */


// The index in the command’s buffer argument table that declares which data to read for this input node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/argumentIndex
func (f_ FunctionStitchingInputNode) SetArgumentIndex(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setArgumentIndex:"), value)
}/* debug [instance_properties/setter]: argumentIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionStitchingInputNode */


