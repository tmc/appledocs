// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	Shape() objc.IObject /* cross-framework: NSNumber */
	SetShape(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// The shaped type class for types on tensors with a shape and data type.


// The shaped type class for types on tensors with a shape and data type.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/init(shape:dataType:)
func NewGraphShapedTypeWithShapeDataType(shape Shape /* not a class type */, dataType DataType /* not a class type */) GraphShapedType {
	instance := getGraphShapedTypeClass().Alloc()
	rv := objc.Send[GraphShapedType](instance.ID, objc.Sel("initWithShape:dataType:"), shape, dataType)
	rv.Autorelease()
	return rv
}



// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}


// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/shape
func (g_ GraphShapedType) Shape() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("shape"))
	return rv
}


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/shape
func (g_ GraphShapedType) SetShape(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}


