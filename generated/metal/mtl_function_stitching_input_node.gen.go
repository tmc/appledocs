// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FunctionStitchingInputNode] class.
type IFunctionStitchingInputNode interface {
	objectivec.IObject
	ArgumentIndex() int
	SetArgumentIndex(value int)
}

// A call graph node that describes an input to the call graph.
//
// An input node contains data from one of the stitched function’s parameters. The output data type of an input node has the same type as the matching parameter.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingInputNodeClass) Alloc() FunctionStitchingInputNode {
	rv := objc.Send[FunctionStitchingInputNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The index in the command’s buffer argument table that declares which data to read for this input node.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode/argumentindex
func (f_ FunctionStitchingInputNode) ArgumentIndex() int {
	rv := objc.Send[int](f_.ID, objc.Sel("argumentIndex"))
	return rv
}


// SetArgumentIndex sets the value of the argumentIndex property.
// The index in the command’s buffer argument table that declares which data to read for this input node.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinginputnode/argumentindex
func (f_ FunctionStitchingInputNode) SetArgumentIndex(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setArgumentIndex:"), value)
}



