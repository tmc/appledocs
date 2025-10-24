// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphCreateSparseOpDescriptor */


/* debug [class_header]: Header for MPSGraphCreateSparseOpDescriptor */
// The class instance for the [GraphCreateSparseOpDescriptor] class.
var (
	GraphCreateSparseOpDescriptorClass     _GraphCreateSparseOpDescriptorClass
	GraphCreateSparseOpDescriptorClassOnce sync.Once
)

func getGraphCreateSparseOpDescriptorClass() _GraphCreateSparseOpDescriptorClass {
	GraphCreateSparseOpDescriptorClassOnce.Do(func() {
		GraphCreateSparseOpDescriptorClass = _GraphCreateSparseOpDescriptorClass{objc.GetClass("MPSGraphCreateSparseOpDescriptor")}
	})
	return GraphCreateSparseOpDescriptorClass
}

type _GraphCreateSparseOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphCreateSparseOpDescriptor */
// An interface definition for the [GraphCreateSparseOpDescriptor] class.
type IGraphCreateSparseOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphCreateSparseOpDescriptor */
	// properties:
	DataType() objc.IObject /* cross-framework: DataType */
	SetDataType(value objc.IObject /* cross-framework: DataType */)
	SparseStorageType() GraphSparseStorageType
	SetSparseStorageType(value GraphSparseStorageType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphCreateSparseOpDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphCreateSparseOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphCreateSparseOpDescriptorClass) Alloc() GraphCreateSparseOpDescriptor {
	rv := objc.Send[GraphCreateSparseOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphCreateSparseOpDescriptorClass) New() GraphCreateSparseOpDescriptor {
	rv := objc.Send[GraphCreateSparseOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphCreateSparseOpDescriptor) Init() GraphCreateSparseOpDescriptor {
	rv := objc.Send[GraphCreateSparseOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphCreateSparseOpDescriptor) Autorelease() GraphCreateSparseOpDescriptor {
	rv := objc.Send[GraphCreateSparseOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphCreateSparseOpDescriptor creates a new GraphCreateSparseOpDescriptor instance.
func NewGraphCreateSparseOpDescriptor() GraphCreateSparseOpDescriptor {
	return getGraphCreateSparseOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphCreateSparseOpDescriptor */
// A class that describes the properties of a create sparse operation.


// A class that describes the properties of a create sparse operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor
type GraphCreateSparseOpDescriptor struct {
	GraphObject
}

// GraphCreateSparseOpDescriptorFrom constructs a [GraphCreateSparseOpDescriptor] from an unsafe.Pointer.
//
// A class that describes the properties of a create sparse operation.
func GraphCreateSparseOpDescriptorFrom(ptr unsafe.Pointer) GraphCreateSparseOpDescriptor {
	return GraphCreateSparseOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphCreateSparseOpDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphCreateSparseOpDescriptor */

// Creates a descriptor for a sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseDescriptor(descriptorWithStorageType:dataType:)
func (gc _GraphCreateSparseOpDescriptorClass) DescriptorWithStorageTypeDataType(sparseStorageType GraphSparseStorageType, dataType objc.IObject /* cross-framework: DataType */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStorageType:dataType:"), sparseStorageType, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithStorageTypeDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphCreateSparseOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphCreateSparseOpDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphCreateSparseOpDescriptor */

// Defines the datatype of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/dataType
func (g_ GraphCreateSparseOpDescriptor) DataType() objc.IObject /* cross-framework: DataType */ {
	rv := objc.Send[metalperformanceshaders.DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// Defines the datatype of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/dataType
func (g_ GraphCreateSparseOpDescriptor) SetDataType(value objc.IObject /* cross-framework: DataType */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// Defines the storage format of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SparseStorageType() GraphSparseStorageType {
	rv := objc.Send[GraphSparseStorageType](g_.ID, objc.Sel("sparseStorageType"))
	return rv
}/* debug [instance_properties/getter]: sparseStorageType */


// Defines the storage format of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SetSparseStorageType(value GraphSparseStorageType) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSparseStorageType:"), value)
}/* debug [instance_properties/setter]: sparseStorageType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphCreateSparseOpDescriptor */



