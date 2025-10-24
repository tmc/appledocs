// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePlanDeviceUsage] class.
var (
	ComputePlanDeviceUsageClass     _ComputePlanDeviceUsageClass
	ComputePlanDeviceUsageClassOnce sync.Once
)

func getComputePlanDeviceUsageClass() _ComputePlanDeviceUsageClass {
	ComputePlanDeviceUsageClassOnce.Do(func() {
		ComputePlanDeviceUsageClass = _ComputePlanDeviceUsageClass{objc.GetClass("MLComputePlanDeviceUsage")}
	})
	return ComputePlanDeviceUsageClass
}

type _ComputePlanDeviceUsageClass struct {
	class objc.Class
}

// An interface definition for the [ComputePlanDeviceUsage] class.
type IComputePlanDeviceUsage interface {
	objectivec.IObject
	// properties:
	PreferredComputeDevice() objc.ID
	SupportedComputeDevices() []objc.ID
	// methods:
}

// The anticipated compute devices to use for executing a layer or operation.


// The anticipated compute devices to use for executing a layer or operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type ComputePlanDeviceUsage struct {
	objectivec.Object
}

// ComputePlanDeviceUsageFrom constructs a [ComputePlanDeviceUsage] from an unsafe.Pointer.
//
// The anticipated compute devices to use for executing a layer or operation.
func ComputePlanDeviceUsageFrom(ptr unsafe.Pointer) ComputePlanDeviceUsage {
	return ComputePlanDeviceUsage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePlanDeviceUsageClass) Alloc() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePlanDeviceUsageClass) New() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePlanDeviceUsage) Init() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePlanDeviceUsage) Autorelease() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePlanDeviceUsage creates a new ComputePlanDeviceUsage instance.
func NewComputePlanDeviceUsage() ComputePlanDeviceUsage {
	return getComputePlanDeviceUsageClass().New()
}



// The compute device that the framework prefers to execute the layer/operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/preferredComputeDevice
func (c_ ComputePlanDeviceUsage) PreferredComputeDevice() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("preferredComputeDevice"))
	return rv
}


// The compute devices that can execute the layer/operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/supportedComputeDevices
func (c_ ComputePlanDeviceUsage) SupportedComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("supportedComputeDevices"))
	return rv
}



