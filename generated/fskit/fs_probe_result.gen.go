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
	Result() FSMatchResult
	ContainerID() IFSContainerIdentifier
	SetContainerID(value IFSContainerIdentifier)
	Name() string
	SetName(value string)
}

// An object that represents the results of a specific probe.
//
// For any value other than , ensure the and values are non- . When a container or volume format doesn’t use a name, return an empty string. Also use an empty string in the case in which the format supports a name, but the value isn’t set yet. Some container or volume formats may lack a durable UUID on which to base a container identifier. This situation is only valid for unary file systems. In such a case, return a random UUID. With a block device resource, a probe operation may successfully get a result but encounter an error reading the name or UUID. If this happens, use whatever information is available, and provide an empty string or random UUID for the name or container ID, respectively.


// An object that represents the results of a specific probe.
//
// [Full Topic]
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



// The match result, representing the recognition and usability of a probed resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSProbeResult/result
func (f_ FSProbeResult) Result() FSMatchResult {
	rv := objc.Send[FSMatchResult](f_.ID, objc.Sel("result"))
	return rv
}


// The container identifier, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsproberesult/containerid
func (f_ FSProbeResult) ContainerID() IFSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](f_.ID, objc.Sel("containerID"))
	return rv
}


// The container identifier, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsproberesult/containerid
func (f_ FSProbeResult) SetContainerID(value IFSContainerIdentifier) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setContainerID:"), value)
}


// The resource name, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsproberesult/name
func (f_ FSProbeResult) Name() string {
	rv := objc.Send[string](f_.ID, objc.Sel("name"))
	return rv
}


// The resource name, as found during the probe operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsproberesult/name
func (f_ FSProbeResult) SetName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), objc.String(value))
}



