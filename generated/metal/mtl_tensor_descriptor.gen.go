// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTensorDescriptor */


/* debug [class_header]: Header for MTLTensorDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TensorDescriptor */
// An interface definition for the [TensorDescriptor] class.
type ITensorDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TensorDescriptor */
	// properties:
	CpuCacheMode() CPUCacheMode
	SetCpuCacheMode(value CPUCacheMode)
	DataType() TensorDataType
	SetDataType(value TensorDataType)
	Dimensions() IMTLTensorExtents
	SetDimensions(value IMTLTensorExtents)
	HazardTrackingMode() HazardTrackingMode
	SetHazardTrackingMode(value HazardTrackingMode)
	ResourceOptions() ResourceOptions
	SetResourceOptions(value ResourceOptions)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	Strides() IMTLTensorExtents
	SetStrides(value IMTLTensorExtents)
	Usage() TensorUsage
	SetUsage(value TensorUsage)
	MTLTensorDomain() objc.IObject /* cross-framework: NSString */
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TensorDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TensorDescriptor */
// Alloc allocates a new instance without initialization.
func (tc _TensorDescriptorClass) Alloc() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TensorDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TensorDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TensorDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TensorDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TensorDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TensorDescriptor */

// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/cpuCacheMode
func (t_ TensorDescriptor) CpuCacheMode() CPUCacheMode {
	rv := objc.Send[CPUCacheMode](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}/* debug [instance_properties/getter]: cpuCacheMode */


// A value that configures the cache mode of CPU mapping of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/cpuCacheMode
func (t_ TensorDescriptor) SetCpuCacheMode(value CPUCacheMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}/* debug [instance_properties/setter]: cpuCacheMode */


// A data format for the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dataType
func (t_ TensorDescriptor) DataType() TensorDataType {
	rv := objc.Send[TensorDataType](t_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// A data format for the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dataType
func (t_ TensorDescriptor) SetDataType(value TensorDataType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) Dimensions() IMTLTensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("dimensions"))
	return rv
}/* debug [instance_properties/getter]: dimensions */


// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) SetDimensions(value IMTLTensorExtents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDimensions:"), value)
}/* debug [instance_properties/setter]: dimensions */


// A value that configures the hazard tracking of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/hazardTrackingMode
func (t_ TensorDescriptor) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Send[HazardTrackingMode](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}/* debug [instance_properties/getter]: hazardTrackingMode */


// A value that configures the hazard tracking of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/hazardTrackingMode
func (t_ TensorDescriptor) SetHazardTrackingMode(value HazardTrackingMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}/* debug [instance_properties/setter]: hazardTrackingMode */


// A packed set of the , and properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/resourceOptions
func (t_ TensorDescriptor) ResourceOptions() ResourceOptions {
	rv := objc.Send[ResourceOptions](t_.ID, objc.Sel("resourceOptions"))
	return rv
}/* debug [instance_properties/getter]: resourceOptions */


// A packed set of the , and properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/resourceOptions
func (t_ TensorDescriptor) SetResourceOptions(value ResourceOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}/* debug [instance_properties/setter]: resourceOptions */


// A value that configures the memory location and access permissions of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/storageMode
func (t_ TensorDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](t_.ID, objc.Sel("storageMode"))
	return rv
}/* debug [instance_properties/getter]: storageMode */


// A value that configures the memory location and access permissions of tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/storageMode
func (t_ TensorDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}/* debug [instance_properties/setter]: storageMode */


// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/strides
func (t_ TensorDescriptor) Strides() IMTLTensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// An array of strides, in elements, one for each dimension in the tensors you create with this descriptor, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/strides
func (t_ TensorDescriptor) SetStrides(value IMTLTensorExtents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStrides:"), value)
}/* debug [instance_properties/setter]: strides */


// A set of contexts in which you can use tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/usage
func (t_ TensorDescriptor) Usage() TensorUsage {
	rv := objc.Send[TensorUsage](t_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// A set of contexts in which you can use tensors you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/usage
func (t_ TensorDescriptor) SetUsage(value TensorUsage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */


// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorDescriptor) MTLTensorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTLTensorDomain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}/* debug [instance_properties/getter]: MTL_TENSOR_MAX_RANK */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorDescriptor) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}/* debug [instance_properties/setter]: MTL_TENSOR_MAX_RANK */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTensorDescriptor */



