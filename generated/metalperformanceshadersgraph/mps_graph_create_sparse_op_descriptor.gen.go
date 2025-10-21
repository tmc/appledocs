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
}

// A class that describes the properties of a create sparse operation.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphCreateSparseOpDescriptorClass) Alloc() GraphCreateSparseOpDescriptor {
	rv := objc.Send[GraphCreateSparseOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a descriptor for a sparse tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseDescriptor(descriptorWithStorageType:dataType:)
func (gc _GraphCreateSparseOpDescriptorClass) DescriptorWithStorageTypeDataType(sparseStorageType unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStorageType:dataType:"), sparseStorageType, dataType)
	return rv
}

// Defines the storage format of the sparse tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SparseStorageType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("sparseStorageType"))
	return rv
}


// SetSparseStorageType sets the value of the sparseStorageType property.
// Defines the storage format of the sparse tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g_ GraphCreateSparseOpDescriptor) SetSparseStorageType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSparseStorageType:"), value)
}



