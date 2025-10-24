// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLHeapDescriptor */


/* debug [class_header]: Header for MTLHeapDescriptor */
// The class instance for the [HeapDescriptor] class.
var (
	HeapDescriptorClass     _HeapDescriptorClass
	HeapDescriptorClassOnce sync.Once
)

func getHeapDescriptorClass() _HeapDescriptorClass {
	HeapDescriptorClassOnce.Do(func() {
		HeapDescriptorClass = _HeapDescriptorClass{objc.GetClass("MTLHeapDescriptor")}
	})
	return HeapDescriptorClass
}

type _HeapDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HeapDescriptor */
// An interface definition for the [HeapDescriptor] class.
type IHeapDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HeapDescriptor */
	// properties:
	CpuCacheMode() CPUCacheMode
	SetCpuCacheMode(value CPUCacheMode)
	HazardTrackingMode() HazardTrackingMode
	SetHazardTrackingMode(value HazardTrackingMode)
	MaxCompatiblePlacementSparsePageSize() SparsePageSize
	SetMaxCompatiblePlacementSparsePageSize(value SparsePageSize)
	ResourceOptions() ResourceOptions
	SetResourceOptions(value ResourceOptions)
	Size() uint
	SetSize(value uint)
	SparsePageSize() SparsePageSize
	SetSparsePageSize(value SparsePageSize)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	Type() HeapType
	SetType(value HeapType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HeapDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HeapDescriptor */
// Alloc allocates a new instance without initialization.
func (hc _HeapDescriptorClass) Alloc() HeapDescriptor {
	rv := objc.Send[HeapDescriptor](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HeapDescriptorClass) New() HeapDescriptor {
	rv := objc.Send[HeapDescriptor](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HeapDescriptor) Init() HeapDescriptor {
	rv := objc.Send[HeapDescriptor](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HeapDescriptor) Autorelease() HeapDescriptor {
	rv := objc.Send[HeapDescriptor](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeapDescriptor creates a new HeapDescriptor instance.
func NewHeapDescriptor() HeapDescriptor {
	return getHeapDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HeapDescriptor */
// A configuration that customizes the behavior for a Metal memory heap.
//
// Create an by configuring an instance’s properties and passing it to the method of an . Each new heap inherits the descriptor’s configuration as you create it, which means you can modify and reuse a descriptor to create other heaps.


// A configuration that customizes the behavior for a Metal memory heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor
type HeapDescriptor struct {
	objectivec.Object
}

// HeapDescriptorFrom constructs a [HeapDescriptor] from an unsafe.Pointer.
//
// A configuration that customizes the behavior for a Metal memory heap.
func HeapDescriptorFrom(ptr unsafe.Pointer) HeapDescriptor {
	return HeapDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HeapDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HeapDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HeapDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HeapDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HeapDescriptor */

// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/cpuCacheMode
func (h_ HeapDescriptor) CpuCacheMode() CPUCacheMode {
	rv := objc.Send[CPUCacheMode](h_.ID, objc.Sel("cpuCacheMode"))
	return rv
}/* debug [instance_properties/getter]: cpuCacheMode */


// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/cpuCacheMode
func (h_ HeapDescriptor) SetCpuCacheMode(value CPUCacheMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCpuCacheMode:"), value)
}/* debug [instance_properties/setter]: cpuCacheMode */


// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/hazardTrackingMode
func (h_ HeapDescriptor) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Send[HazardTrackingMode](h_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}/* debug [instance_properties/getter]: hazardTrackingMode */


// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/hazardTrackingMode
func (h_ HeapDescriptor) SetHazardTrackingMode(value HazardTrackingMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHazardTrackingMode:"), value)
}/* debug [instance_properties/setter]: hazardTrackingMode */


// Specifies the largest sparse page size that the Metal heap supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/maxCompatiblePlacementSparsePageSize
func (h_ HeapDescriptor) MaxCompatiblePlacementSparsePageSize() SparsePageSize {
	rv := objc.Send[SparsePageSize](h_.ID, objc.Sel("maxCompatiblePlacementSparsePageSize"))
	return rv
}/* debug [instance_properties/getter]: maxCompatiblePlacementSparsePageSize */


// Specifies the largest sparse page size that the Metal heap supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/maxCompatiblePlacementSparsePageSize
func (h_ HeapDescriptor) SetMaxCompatiblePlacementSparsePageSize(value SparsePageSize) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMaxCompatiblePlacementSparsePageSize:"), value)
}/* debug [instance_properties/setter]: maxCompatiblePlacementSparsePageSize */


// The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/resourceOptions
func (h_ HeapDescriptor) ResourceOptions() ResourceOptions {
	rv := objc.Send[ResourceOptions](h_.ID, objc.Sel("resourceOptions"))
	return rv
}/* debug [instance_properties/getter]: resourceOptions */


// The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/resourceOptions
func (h_ HeapDescriptor) SetResourceOptions(value ResourceOptions) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResourceOptions:"), value)
}/* debug [instance_properties/setter]: resourceOptions */


// The total amount of memory, in bytes, for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/size
func (h_ HeapDescriptor) Size() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The total amount of memory, in bytes, for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/size
func (h_ HeapDescriptor) SetSize(value uint) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// The page size for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/sparsePageSize
func (h_ HeapDescriptor) SparsePageSize() SparsePageSize {
	rv := objc.Send[SparsePageSize](h_.ID, objc.Sel("sparsePageSize"))
	return rv
}/* debug [instance_properties/getter]: sparsePageSize */


// The page size for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/sparsePageSize
func (h_ HeapDescriptor) SetSparsePageSize(value SparsePageSize) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSparsePageSize:"), value)
}/* debug [instance_properties/setter]: sparsePageSize */


// The storage mode for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/storageMode
func (h_ HeapDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](h_.ID, objc.Sel("storageMode"))
	return rv
}/* debug [instance_properties/getter]: storageMode */


// The storage mode for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/storageMode
func (h_ HeapDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStorageMode:"), value)
}/* debug [instance_properties/setter]: storageMode */


// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/type
func (h_ HeapDescriptor) Type() HeapType {
	rv := objc.Send[HeapType](h_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/type
func (h_ HeapDescriptor) SetType(value HeapType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLHeapDescriptor */



