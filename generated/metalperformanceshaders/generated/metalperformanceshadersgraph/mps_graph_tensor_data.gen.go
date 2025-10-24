// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphTensorData */


/* debug [class_header]: Header for MPSGraphTensorData */
// The class instance for the [GraphTensorData] class.
var (
	GraphTensorDataClass     _GraphTensorDataClass
	GraphTensorDataClassOnce sync.Once
)

func getGraphTensorDataClass() _GraphTensorDataClass {
	GraphTensorDataClassOnce.Do(func() {
		GraphTensorDataClass = _GraphTensorDataClass{objc.GetClass("MPSGraphTensorData")}
	})
	return GraphTensorDataClass
}

type _GraphTensorDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphTensorData */
// An interface definition for the [GraphTensorData] class.
type IGraphTensorData interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphTensorData */
	// properties:
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	Device() IMPSGraphDevice
	SetDevice(value IMPSGraphDevice)
	Shape() objc.IObject /* cross-framework: NSNumber */
	SetShape(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphTensorData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphTensorData */
// Alloc allocates a new instance without initialization.
func (gc _GraphTensorDataClass) Alloc() GraphTensorData {
	rv := objc.Send[GraphTensorData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphTensorDataClass) New() GraphTensorData {
	rv := objc.Send[GraphTensorData](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphTensorData) Init() GraphTensorData {
	rv := objc.Send[GraphTensorData](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphTensorData) Autorelease() GraphTensorData {
	rv := objc.Send[GraphTensorData](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphTensorData creates a new GraphTensorData instance.
func NewGraphTensorData() GraphTensorData {
	return getGraphTensorDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphTensorData */
// The representation of a compute data type.
//
// Pass data to a graph using a tensor data, a reference will be taken to your data and used just in time when the graph is run.


// The representation of a compute data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData
type GraphTensorData struct {
	GraphObject
}

// GraphTensorDataFrom constructs a [GraphTensorData] from an unsafe.Pointer.
//
// The representation of a compute data type.
func GraphTensorDataFrom(ptr unsafe.Pointer) GraphTensorData {
	return GraphTensorData{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphTensorData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphTensorData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphTensorData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphTensorData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphTensorData */

// The data type of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/datatype
func (g_ GraphTensorData) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/datatype
func (g_ GraphTensorData) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The device of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/device
func (g_ GraphTensorData) Device() IMPSGraphDevice {
	rv := objc.Send[GraphDevice](g_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/device
func (g_ GraphTensorData) SetDevice(value IMPSGraphDevice) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// The shape of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/shape
func (g_ GraphTensorData) Shape() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The shape of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/shape
func (g_ GraphTensorData) SetShape(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphTensorData */



