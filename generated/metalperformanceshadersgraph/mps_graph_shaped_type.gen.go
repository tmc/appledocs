// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	Shape() objc.IObject /* cross-framework: NSNumber */
	SetShape(value objc.IObject /* cross-framework: NSNumber */)
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



/* debug [class_init_methods]: Init methods for GraphShapedType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphShapedType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphShapedType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphShapedType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphShapedType */

// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/datatype
func (g_ GraphShapedType) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/datatype
func (g_ GraphShapedType) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/shape
func (g_ GraphShapedType) Shape() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The Shape of the shaped type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphshapedtype/shape
func (g_ GraphShapedType) SetShape(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphShapedType */



