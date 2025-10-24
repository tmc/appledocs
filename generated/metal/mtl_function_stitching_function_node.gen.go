// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FunctionStitchingFunctionNode] class.
type IFunctionStitchingFunctionNode interface {
	objectivec.IObject
	// properties:
	Arguments() FunctionStitchingNode /* not a class type */
	SetArguments(value FunctionStitchingNode /* not a class type */)
	ControlDependencies() IMTLFunctionStitchingFunctionNode
	SetControlDependencies(value IMTLFunctionStitchingFunctionNode)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingFunctionNodeClass) Alloc() FunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An ordered list of the nodes that provide the function’s arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/arguments
func (f_ FunctionStitchingFunctionNode) Arguments() FunctionStitchingNode /* not a class type */ {
	rv := objc.Send[FunctionStitchingNode](f_.ID, objc.Sel("arguments"))
	return rv
}


// An ordered list of the nodes that provide the function’s arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/arguments
func (f_ FunctionStitchingFunctionNode) SetArguments(value FunctionStitchingNode /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setArguments:"), value)
}


// The list of nodes that must execute before executing the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/controldependencies
func (f_ FunctionStitchingFunctionNode) ControlDependencies() IMTLFunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](f_.ID, objc.Sel("controlDependencies"))
	return rv
}


// The list of nodes that must execute before executing the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/controldependencies
func (f_ FunctionStitchingFunctionNode) SetControlDependencies(value IMTLFunctionStitchingFunctionNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setControlDependencies:"), value)
}


// The name of the function to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/name
func (f_ FunctionStitchingFunctionNode) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}


// The name of the function to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingfunctionnode/name
func (f_ FunctionStitchingFunctionNode) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}



