// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphVariableOp */


/* debug [class_header]: Header for MPSGraphVariableOp */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphVariableOp */
// An interface definition for the [GraphVariableOp] class.
type IGraphVariableOp interface {
	IGraphOperation
	
/* debug [class_interface_properties]: Properties for GraphVariableOp */
	// properties:
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	Shape() objc.IObject /* cross-framework: NSNumber */
	SetShape(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphVariableOp */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphVariableOp */
// Alloc allocates a new instance without initialization.
func (gc _GraphVariableOpClass) Alloc() GraphVariableOp {
	rv := objc.Send[GraphVariableOp](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphVariableOp */
// The class that defines the parameters for a variable.


// The class that defines the parameters for a variable.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphVariableOp *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphVariableOp */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphVariableOp */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphVariableOp */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphVariableOp */

// The data type of the variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/datatype
func (g_ GraphVariableOp) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/datatype
func (g_ GraphVariableOp) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The shape of the variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/shape
func (g_ GraphVariableOp) Shape() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The shape of the variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/shape
func (g_ GraphVariableOp) SetShape(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphVariableOp */






