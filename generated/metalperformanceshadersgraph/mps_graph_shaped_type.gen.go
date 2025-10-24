// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphShapedType */


/* debug [class_header]: Header for MPSGraphShapedType */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphShapedType */
// An interface definition for the [GraphShapedType] class.
type IGraphShapedType interface {
	IGraphType
	
/* debug [class_interface_properties]: Properties for GraphShapedType */
	// properties:
	DataType() objc.IObject /* cross-framework: DataType */
	SetDataType(value objc.IObject /* cross-framework: DataType */)
	Shape() Shape /* not a class type */
	SetShape(value Shape /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphShapedType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphShapedType */
// Alloc allocates a new instance without initialization.
func (gc _GraphShapedTypeClass) Alloc() GraphShapedType {
	rv := objc.Send[GraphShapedType](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphShapedType */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphShapedType */

// Initializes a shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/init(shape:dataType:)
func NewGraphShapedTypeWithShapeDataType(shape Shape /* not a class type */, dataType objc.IObject /* cross-framework: DataType */) GraphShapedType {
	instance := getGraphShapedTypeClass().Alloc()
	rv := objc.Send[GraphShapedType](instance.ID, objc.Sel("initWithShape:dataType:"), shape, dataType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphShapedTypeWithShapeDataType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphShapedType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphShapedType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphShapedType */

// Checks if shapes and element data type are the same as the input shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/isEqual(to:)
func (g_ GraphShapedType) IsEqualTo(object IMPSGraphShapedType) bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isEqualTo:"), object)
	return rv
}/* debug [instance_methods/method]: IsEqualTo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphShapedType */

// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) DataType() objc.IObject /* cross-framework: DataType */ {
	rv := objc.Send[metalperformanceshaders.DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g_ GraphShapedType) SetDataType(value objc.IObject /* cross-framework: DataType */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/shape
func (g_ GraphShapedType) Shape() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/shape
func (g_ GraphShapedType) SetShape(value Shape /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphShapedType */


