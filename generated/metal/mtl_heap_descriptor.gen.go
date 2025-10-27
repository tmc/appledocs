// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [HeapDescriptor] class.
type IHeapDescriptor interface {
	objectivec.IObject
	

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


	

	// methods:


}





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

























// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/cpuCacheMode
func (h_ HeapDescriptor) CpuCacheMode() CPUCacheMode {
	rv := objc.Send[CPUCacheMode](h_.ID, objc.Sel("cpuCacheMode"))
	return rv
}


// The CPU cache behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/cpuCacheMode
func (h_ HeapDescriptor) SetCpuCacheMode(value CPUCacheMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCpuCacheMode:"), value)
}


// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/hazardTrackingMode
func (h_ HeapDescriptor) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Send[HazardTrackingMode](h_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}


// The hazard tracking behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/hazardTrackingMode
func (h_ HeapDescriptor) SetHazardTrackingMode(value HazardTrackingMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHazardTrackingMode:"), value)
}


// Specifies the largest sparse page size that the Metal heap supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/maxCompatiblePlacementSparsePageSize
func (h_ HeapDescriptor) MaxCompatiblePlacementSparsePageSize() SparsePageSize {
	rv := objc.Send[SparsePageSize](h_.ID, objc.Sel("maxCompatiblePlacementSparsePageSize"))
	return rv
}


// Specifies the largest sparse page size that the Metal heap supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/maxCompatiblePlacementSparsePageSize
func (h_ HeapDescriptor) SetMaxCompatiblePlacementSparsePageSize(value SparsePageSize) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMaxCompatiblePlacementSparsePageSize:"), value)
}


// The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/resourceOptions
func (h_ HeapDescriptor) ResourceOptions() ResourceOptions {
	rv := objc.Send[ResourceOptions](h_.ID, objc.Sel("resourceOptions"))
	return rv
}


// The combined behavior for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/resourceOptions
func (h_ HeapDescriptor) SetResourceOptions(value ResourceOptions) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResourceOptions:"), value)
}


// The total amount of memory, in bytes, for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/size
func (h_ HeapDescriptor) Size() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("size"))
	return rv
}


// The total amount of memory, in bytes, for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/size
func (h_ HeapDescriptor) SetSize(value uint) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSize:"), value)
}


// The page size for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/sparsePageSize
func (h_ HeapDescriptor) SparsePageSize() SparsePageSize {
	rv := objc.Send[SparsePageSize](h_.ID, objc.Sel("sparsePageSize"))
	return rv
}


// The page size for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/sparsePageSize
func (h_ HeapDescriptor) SetSparsePageSize(value SparsePageSize) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSparsePageSize:"), value)
}


// The storage mode for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/storageMode
func (h_ HeapDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](h_.ID, objc.Sel("storageMode"))
	return rv
}


// The storage mode for the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/storageMode
func (h_ HeapDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStorageMode:"), value)
}


// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/type
func (h_ HeapDescriptor) Type() HeapType {
	rv := objc.Send[HeapType](h_.ID, objc.Sel("type"))
	return rv
}


// The memory placement strategy for any resources you allocate from the heaps you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapDescriptor/type
func (h_ HeapDescriptor) SetType(value HeapType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setType:"), value)
}








