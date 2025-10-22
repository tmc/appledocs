// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphShapedType] class.
var (
	GraphShapedTypeClass     _GraphShapedTypeClass
	GraphShapedTypeClassOnce sync.Once
)

func getGraphShapedTypeClass() _GraphShapedTypeClass {
	GraphShapedTypeClassOnce.Do(func() {
		GraphShapedTypeClass = _GraphShapedTypeClass{objc.GetClass("MPSGraphShapedType")}
	})
	return GraphShapedTypeClass
}

type _GraphShapedTypeClass struct {
	class objc.Class
}

// An interface definition for the [GraphShapedType] class.
type IGraphShapedType interface {
	IGraphType
	DataType() unsafe.Pointer
	SetDataType(value unsafe.Pointer)
	Shape() unsafe.Pointer
	SetShape(value unsafe.Pointer)
}

// The shaped type class for types on tensors with a shape and data type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType
type GraphShapedType struct {
	GraphType
}

// GraphShapedTypeFrom constructs a [GraphShapedType] from an unsafe.Pointer.
//
// The shaped type class for types on tensors with a shape and data type.
func GraphShapedTypeFrom(ptr unsafe.Pointer) GraphShapedType {
	return GraphShapedType{
		GraphType: GraphTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphShapedTypeClass) Alloc() GraphShapedType {
	rv := objc.Send[GraphShapedType](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphShapedTypeClass) New() GraphShapedType {
	rv := objc.Send[GraphShapedType](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphShapedType) Init() GraphShapedType {
	rv := objc.Send[GraphShapedType](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphShapedType) Autorelease() GraphShapedType {
	rv := objc.Send[GraphShapedType](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphShapedType creates a new GraphShapedType instance.
func NewGraphShapedType() GraphShapedType {
	return getGraphShapedTypeClass().New()
}




// Initializes a shaped type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/init(shape:dataType:)
func NewGraphShapedTypeWithShapeDataType(shape unsafe.Pointer, dataType unsafe.Pointer) GraphShapedType {
	instance := getGraphShapedTypeClass().Alloc()
	rv := objc.Send[GraphShapedType](instance.ID, objc.Sel("initWithShape:dataType:"), shape, dataType)
	rv.Autorelease()
	return rv
}


// Checks if shapes and element data type are the same as the input shaped type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/isEqual(to:)
func (g_ GraphShapedType) IsEqualTo(object MPSGraphShapedType) bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isEqualTo:"), object)
	return rv
}

// The data type of the shaped type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The data type of the shaped type.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}

// The Shape of the shaped type.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/shape
func (g_ GraphShapedType) Shape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("shape"))
	return rv
}


// SetShape sets the value of the shape property.
// The Shape of the shaped type.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/shape
func (g_ GraphShapedType) SetShape(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}


