// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TensorDescriptor] class.
var (
	TensorDescriptorClass     _TensorDescriptorClass
	TensorDescriptorClassOnce sync.Once
)

func getTensorDescriptorClass() _TensorDescriptorClass {
	TensorDescriptorClassOnce.Do(func() {
		TensorDescriptorClass = _TensorDescriptorClass{objc.GetClass("MTLTensorDescriptor")}
	})
	return TensorDescriptorClass
}

type _TensorDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [TensorDescriptor] class.
type ITensorDescriptor interface {
	objectivec.IObject
}

// A configuration type for creating new tensor instances.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor
type TensorDescriptor struct {
	objectivec.Object
}

// TensorDescriptorFrom constructs a [TensorDescriptor] from an unsafe.Pointer.
//
// A configuration type for creating new tensor instances.
func TensorDescriptorFrom(ptr unsafe.Pointer) TensorDescriptor {
	return TensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TensorDescriptorClass) Alloc() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TensorDescriptorClass) New() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TensorDescriptor) Init() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TensorDescriptor) Autorelease() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTensorDescriptor creates a new TensorDescriptor instance.
func NewTensorDescriptor() TensorDescriptor {
	return getTensorDescriptorClass().New()
}


// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorDescriptor) MTLTensorDomain() string {
	rv := objc.Send[string](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}

// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/cpucachemode
func (t_ TensorDescriptor) CpuCacheMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}


// SetCpuCacheMode sets the value of the cpuCacheMode property.
// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/cpucachemode
func (t_ TensorDescriptor) SetCpuCacheMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) MTL_TENSOR_MAX_RANK() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// SetMTL_TENSOR_MAX_RANK sets the value of the MTL_TENSOR_MAX_RANK property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) SetMTL_TENSOR_MAX_RANK(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}

// A value that configures the hazard tracking of tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/hazardtrackingmode
func (t_ TensorDescriptor) HazardTrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}


// SetHazardTrackingMode sets the value of the hazardTrackingMode property.
// A value that configures the hazard tracking of tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/hazardtrackingmode
func (t_ TensorDescriptor) SetHazardTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}

// A set of contexts in which you can use tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/usage
func (t_ TensorDescriptor) Usage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("usage"))
	return rv
}


// SetUsage sets the value of the usage property.
// A set of contexts in which you can use tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/usage
func (t_ TensorDescriptor) SetUsage(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}

// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/strides
func (t_ TensorDescriptor) Strides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("strides"))
	return rv
}


// SetStrides sets the value of the strides property.
// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/strides
func (t_ TensorDescriptor) SetStrides(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStrides:"), value)
}

// A packed set of the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/resourceoptions
func (t_ TensorDescriptor) ResourceOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resourceOptions"))
	return rv
}


// SetResourceOptions sets the value of the resourceOptions property.
// A packed set of the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/resourceoptions
func (t_ TensorDescriptor) SetResourceOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}

// A value that configures the memory location and access permissions of tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/storagemode
func (t_ TensorDescriptor) StorageMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("storageMode"))
	return rv
}


// SetStorageMode sets the value of the storageMode property.
// A value that configures the memory location and access permissions of tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/storagemode
func (t_ TensorDescriptor) SetStorageMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}

// A data format for the tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/datatype
func (t_ TensorDescriptor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// A data format for the tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/datatype
func (t_ TensorDescriptor) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataType:"), value)
}

// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) Dimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("dimensions"))
	return rv
}


// SetDimensions sets the value of the dimensions property.
// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) SetDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDimensions:"), value)
}



