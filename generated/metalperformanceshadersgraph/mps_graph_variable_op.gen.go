// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphVariableOp] class.
var (
	GraphVariableOpClass     _GraphVariableOpClass
	GraphVariableOpClassOnce sync.Once
)

func getGraphVariableOpClass() _GraphVariableOpClass {
	GraphVariableOpClassOnce.Do(func() {
		GraphVariableOpClass = _GraphVariableOpClass{objc.GetClass("MPSGraphVariableOp")}
	})
	return GraphVariableOpClass
}

type _GraphVariableOpClass struct {
	class objc.Class
}

// An interface definition for the [GraphVariableOp] class.
type IGraphVariableOp interface {
	IGraphOperation
}

// The class that defines the parameters for a variable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp
type GraphVariableOp struct {
	GraphOperation
}

// GraphVariableOpFrom constructs a [GraphVariableOp] from an unsafe.Pointer.
//
// The class that defines the parameters for a variable.
func GraphVariableOpFrom(ptr unsafe.Pointer) GraphVariableOp {
	return GraphVariableOp{
		GraphOperation: GraphOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphVariableOpClass) Alloc() GraphVariableOp {
	rv := objc.Send[GraphVariableOp](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphVariableOpClass) New() GraphVariableOp {
	rv := objc.Send[GraphVariableOp](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphVariableOp) Init() GraphVariableOp {
	rv := objc.Send[GraphVariableOp](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphVariableOp) Autorelease() GraphVariableOp {
	rv := objc.Send[GraphVariableOp](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphVariableOp creates a new GraphVariableOp instance.
func NewGraphVariableOp() GraphVariableOp {
	return getGraphVariableOpClass().New()
}


// The shape of the variable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp/shape
func (g_ GraphVariableOp) Shape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("shape"))
	return rv
}

// The data type of the variable.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/datatype
func (g_ GraphVariableOp) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The data type of the variable.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/datatype
func (g_ GraphVariableOp) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}




