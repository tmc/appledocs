// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CpuCacheMode() CPUCacheMode /* not a class type */
	SetCpuCacheMode(value CPUCacheMode /* not a class type */)
	DataType() TensorDataType /* not a class type */
	SetDataType(value TensorDataType /* not a class type */)
	Dimensions() objc.IObject /* cross-framework: TensorExtents */
	SetDimensions(value objc.IObject /* cross-framework: TensorExtents */)
	HazardTrackingMode() HazardTrackingMode /* not a class type */
	SetHazardTrackingMode(value HazardTrackingMode /* not a class type */)
	ResourceOptions() ResourceOptions /* not a class type */
	SetResourceOptions(value ResourceOptions /* not a class type */)
	StorageMode() StorageMode /* not a class type */
	SetStorageMode(value StorageMode /* not a class type */)
	Strides() objc.IObject /* cross-framework: TensorExtents */
	SetStrides(value objc.IObject /* cross-framework: TensorExtents */)
	Usage() TensorUsage /* not a class type */
	SetUsage(value TensorUsage /* not a class type */)
	MTLTensorDomain() objc.IObject /* cross-framework: NSString */
	MTL_TENSOR_MAX_RANK() unsafe.Pointer
	SetMTL_TENSOR_MAX_RANK(value unsafe.Pointer)
	// methods:
}

// A configuration type for creating new tensor instances.


// A configuration type for creating new tensor instances.
//
// [Full Topic]
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



// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/cpucachemode
func (t_ TensorDescriptor) CpuCacheMode() CPUCacheMode /* not a class type */ {
	rv := objc.Send[CPUCacheMode](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}


// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/cpucachemode
func (t_ TensorDescriptor) SetCpuCacheMode(value CPUCacheMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}


// A data format for the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/datatype
func (t_ TensorDescriptor) DataType() TensorDataType /* not a class type */ {
	rv := objc.Send[TensorDataType](t_.ID, objc.Sel("dataType"))
	return rv
}


// A data format for the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/datatype
func (t_ TensorDescriptor) SetDataType(value TensorDataType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataType:"), value)
}


// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/dimensions
func (t_ TensorDescriptor) Dimensions() objc.IObject /* cross-framework: TensorExtents */ {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("dimensions"))
	return rv
}


// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/dimensions
func (t_ TensorDescriptor) SetDimensions(value objc.IObject /* cross-framework: TensorExtents */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDimensions:"), value)
}


// A value that configures the hazard tracking of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/hazardtrackingmode
func (t_ TensorDescriptor) HazardTrackingMode() HazardTrackingMode /* not a class type */ {
	rv := objc.Send[HazardTrackingMode](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}


// A value that configures the hazard tracking of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/hazardtrackingmode
func (t_ TensorDescriptor) SetHazardTrackingMode(value HazardTrackingMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}


// A packed set of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/resourceoptions
func (t_ TensorDescriptor) ResourceOptions() ResourceOptions /* not a class type */ {
	rv := objc.Send[ResourceOptions](t_.ID, objc.Sel("resourceOptions"))
	return rv
}


// A packed set of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/resourceoptions
func (t_ TensorDescriptor) SetResourceOptions(value ResourceOptions /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}


// A value that configures the memory location and access permissions of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/storagemode
func (t_ TensorDescriptor) StorageMode() StorageMode /* not a class type */ {
	rv := objc.Send[StorageMode](t_.ID, objc.Sel("storageMode"))
	return rv
}


// A value that configures the memory location and access permissions of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/storagemode
func (t_ TensorDescriptor) SetStorageMode(value StorageMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}


// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/strides
func (t_ TensorDescriptor) Strides() objc.IObject /* cross-framework: TensorExtents */ {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("strides"))
	return rv
}


// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/strides
func (t_ TensorDescriptor) SetStrides(value objc.IObject /* cross-framework: TensorExtents */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStrides:"), value)
}


// A set of contexts in which you can use tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/usage
func (t_ TensorDescriptor) Usage() TensorUsage /* not a class type */ {
	rv := objc.Send[TensorUsage](t_.ID, objc.Sel("usage"))
	return rv
}


// A set of contexts in which you can use tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordescriptor/usage
func (t_ TensorDescriptor) SetUsage(value TensorUsage /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}


// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorDescriptor) MTLTensorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) MTL_TENSOR_MAX_RANK() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) SetMTL_TENSOR_MAX_RANK(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}



