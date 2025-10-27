// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FunctionStitchingGraph] class.
var (
	FunctionStitchingGraphClass     _FunctionStitchingGraphClass
	FunctionStitchingGraphClassOnce sync.Once
)

func getFunctionStitchingGraphClass() _FunctionStitchingGraphClass {
	FunctionStitchingGraphClassOnce.Do(func() {
		FunctionStitchingGraphClass = _FunctionStitchingGraphClass{objc.GetClass("MTLFunctionStitchingGraph")}
	})
	return FunctionStitchingGraphClass
}

type _FunctionStitchingGraphClass struct {
	class objc.Class
}





// An interface definition for the [FunctionStitchingGraph] class.
type IFunctionStitchingGraph interface {
	objectivec.IObject
	

	// properties:
	Attributes() []objc.ID
	SetAttributes(value []objc.ID)
	FunctionName() foundation.foundation.INSString
	SetFunctionName(value foundation.foundation.INSString)
	Nodes() []FunctionStitchingFunctionNode
	SetNodes(value []FunctionStitchingFunctionNode)
	OutputNode() IMTLFunctionStitchingFunctionNode
	SetOutputNode(value IMTLFunctionStitchingFunctionNode)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingGraphClass) Alloc() FunctionStitchingGraph {
	rv := objc.Send[FunctionStitchingGraph](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionStitchingGraphClass) New() FunctionStitchingGraph {
	rv := objc.Send[FunctionStitchingGraph](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionStitchingGraph) Init() FunctionStitchingGraph {
	rv := objc.Send[FunctionStitchingGraph](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionStitchingGraph) Autorelease() FunctionStitchingGraph {
	rv := objc.Send[FunctionStitchingGraph](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionStitchingGraph creates a new FunctionStitchingGraph instance.
func NewFunctionStitchingGraph() FunctionStitchingGraph {
	return getFunctionStitchingGraphClass().New()
}





// A description of a new stitched function.
//
// An instance describes the function graph for a stitched function. A is a visible function you create by composing other Metal shader functions together in a function graph. A function stitching graph contains nodes for the function’s arguments and any functions it calls in the implementation. Data flows from the arguments to the end of the graph until the stitched function evaluates all of the graph’s nodes. The graph in the figure below constructs a new function that adds numbers from two source arrays, storing the result in a third array. The function’s parameters are pointers to the source and destination arrays, and an index for performing the array lookup. The graph uses three separate MSL functions to construct the stitched function: a function to look up a value from an array, a function that adds two numbers together, and a function that stores a value to an array. Create an instance for each stitched function you want to create. Configure its properties to describe the new function and the nodes that define its behavior, as described below. To create a new library with these stitched functions, see .


// A description of a new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph
type FunctionStitchingGraph struct {
	objectivec.Object
}

// FunctionStitchingGraphFrom constructs a [FunctionStitchingGraph] from an unsafe.Pointer.
//
// A description of a new stitched function.
func FunctionStitchingGraphFrom(ptr unsafe.Pointer) FunctionStitchingGraph {
	return FunctionStitchingGraph{objectivec.Object{objc.ID(ptr)}}
}






// Creates a description of a new function call graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/init(functionName:nodes:outputNode:attributes:)
func NewFunctionStitchingGraphWithFunctionNameNodesOutputNodeAttributes(functionName foundation.foundation.INSString, nodes []FunctionStitchingFunctionNode, outputNode IMTLFunctionStitchingFunctionNode, attributes []objc.ID) FunctionStitchingGraph {
	instance := getFunctionStitchingGraphClass().Alloc()
	rv := objc.Send[FunctionStitchingGraph](instance.ID, objc.Sel("initWithFunctionName:nodes:outputNode:attributes:"), functionName, nodes, outputNode, attributes)
	rv.Autorelease()
	return rv
}






















// A list of attributes to configure how the Metal device object generates the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/attributes
func (f_ FunctionStitchingGraph) Attributes() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("attributes"))
	return rv
}


// A list of attributes to configure how the Metal device object generates the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/attributes
func (f_ FunctionStitchingGraph) SetAttributes(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributes:"), nsArray)
}


// The name of the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/functionName
func (f_ FunctionStitchingGraph) FunctionName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("functionName"))
	return rv
}


// The name of the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/functionName
func (f_ FunctionStitchingGraph) SetFunctionName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFunctionName:"), value)
}


// The nodes in the function’s call graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/nodes
func (f_ FunctionStitchingGraph) Nodes() []FunctionStitchingFunctionNode {
	rv := objc.Send[[]FunctionStitchingFunctionNode](f_.ID, objc.Sel("nodes"))
	return rv
}


// The nodes in the function’s call graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/nodes
func (f_ FunctionStitchingGraph) SetNodes(value []FunctionStitchingFunctionNode) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setNodes:"), nsArray)
}


// The node with the output that’s the output of the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/outputNode
func (f_ FunctionStitchingGraph) OutputNode() IMTLFunctionStitchingFunctionNode {
	rv := objc.Send[FunctionStitchingFunctionNode](f_.ID, objc.Sel("outputNode"))
	return rv
}


// The node with the output that’s the output of the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/outputNode
func (f_ FunctionStitchingGraph) SetOutputNode(value IMTLFunctionStitchingFunctionNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputNode:"), value)
}







