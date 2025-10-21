// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSProbeResult] class.
var (
	FSProbeResultClass     _FSProbeResultClass
	FSProbeResultClassOnce sync.Once
)

func getFSProbeResultClass() _FSProbeResultClass {
	FSProbeResultClassOnce.Do(func() {
		FSProbeResultClass = _FSProbeResultClass{objc.GetClass("FSProbeResult")}
	})
	return FSProbeResultClass
}

type _FSProbeResultClass struct {
	class objc.Class
}

// An interface definition for the [FSProbeResult] class.
type IFSProbeResult interface {
	objectivec.IObject
}

// An object that represents the results of a specific probe.
//
// For any value other than , ensure the and values are non- . When a container or volume format doesn’t use a name, return an empty string. Also use an empty string in the case in which the format supports a name, but the value isn’t set yet. Some container or volume formats may lack a durable UUID on which to base a container identifier. This situation is only valid for unary file systems. In such a case, return a random UUID. With a block device resource, a probe operation may successfully get a result but encounter an error reading the name or UUID. If this happens, use whatever information is available, and provide an empty string or random UUID for the name or container ID, respectively.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult
type FSProbeResult struct {
	objectivec.Object
}

// FSProbeResultFrom constructs a [FSProbeResult] from an unsafe.Pointer.
//
// An object that represents the results of a specific probe.
func FSProbeResultFrom(ptr unsafe.Pointer) FSProbeResult {
	return FSProbeResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSProbeResultClass) Alloc() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSProbeResultClass) New() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSProbeResult) Init() FSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSProbeResult) Autorelease() FSProbeResult {
	rv := objc.Send[FSProbeResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSProbeResult creates a new FSProbeResult instance.
func NewFSProbeResult() FSProbeResult {
	return getFSProbeResultClass().New()
}


// Creates a probe result for a recognized file system.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/recognized(name:containerID:)
func (fc _FSProbeResultClass) RecognizedProbeResultWithNameContainerID(name string, containerID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("recognizedProbeResultWithName:containerID:"), objc.String(name), containerID)
	return rv
}

// Creates a probe result for a recognized and usable file system.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usable(name:containerID:)
func (fc _FSProbeResultClass) UsableProbeResultWithNameContainerID(name string, containerID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("usableProbeResultWithName:containerID:"), objc.String(name), containerID)
	return rv
}

// Creates a probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited(name:containerID:)
func (fc _FSProbeResultClass) UsableButLimitedProbeResultWithNameContainerID(name string, containerID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("usableButLimitedProbeResultWithName:containerID:"), objc.String(name), containerID)
	return rv
}

// A probe result for an unrecognized file system.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/notRecognized
func (fc _FSProbeResultClass) NotRecognizedProbeResult() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("notRecognizedProbeResult"))
	return rv
}
// A probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited
func (fc _FSProbeResultClass) UsableButLimitedProbeResult() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("usableButLimitedProbeResult"))
	return rv
}
// The container identifier, as found during the probe operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/containerID
func (f_ FSProbeResult) ContainerID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("containerID"))
	return rv
}

// The resource name, as found during the probe operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/name
func (f_ FSProbeResult) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("name"))
	return rv
}

// A probe result for an unrecognized file system.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/notRecognized
func (f_ FSProbeResult) NotRecognizedProbeResult() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("notRecognizedProbeResult"))
	return rv
}

// The match result, representing the recognition and usability of a probed resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/result
func (f_ FSProbeResult) Result() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("result"))
	return rv
}

// A probe result for a recognized file system that is usable, but with limited capabilities.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited
func (f_ FSProbeResult) UsableButLimitedProbeResult() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("usableButLimitedProbeResult"))
	return rv
}



