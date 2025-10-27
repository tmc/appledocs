// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [GraphCreateSparseOpDescriptor] class.
type IGraphCreateSparseOpDescriptor interface {
	IGraphObject
	

	// properties:
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	SparseStorageType() GraphSparseStorageType
	SetSparseStorageType(value GraphSparseStorageType)


	

	// methods:


}





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










// Creates a descriptor for a sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseDescriptor(descriptorWithStorageType:dataType:)
func (gc _GraphCreateSparseOpDescriptorClass) DescriptorWithStorageTypeDataType(sparseStorageType GraphSparseStorageType, dataType DataType /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStorageType:dataType:"), sparseStorageType, dataType)
	return rv
}

















// Defines the datatype of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/dataType
func (g_ GraphCreateSparseOpDescriptor) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}


// Defines the datatype of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/dataType
func (g_ GraphCreateSparseOpDescriptor) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}


// Defines the storage format of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SparseStorageType() GraphSparseStorageType {
	rv := objc.Send[GraphSparseStorageType](g_.ID, objc.Sel("sparseStorageType"))
	return rv
}


// Defines the storage format of the sparse tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SetSparseStorageType(value GraphSparseStorageType) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSparseStorageType:"), value)
}








