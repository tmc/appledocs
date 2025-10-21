// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Request] class.
var (
	RequestClass     _RequestClass
	RequestClassOnce sync.Once
)

func getRequestClass() _RequestClass {
	RequestClassOnce.Do(func() {
		RequestClass = _RequestClass{objc.GetClass("VNRequest")}
	})
	return RequestClass
}

type _RequestClass struct {
	class objc.Class
}

// An interface definition for the [Request] class.
type IRequest interface {
	objectivec.IObject
	Cancel()
	ComputeDeviceForComputeStage(computeStage unsafe.Pointer) objc.ID
	SetComputeDeviceForComputeStage(computeDevice objc.ID, computeStage unsafe.Pointer)
	SupportedComputeStageDevicesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
}

// The abstract superclass for analysis requests.
//
// Other Vision request handlers that perform image analysis inherit from this abstract base class. Instantiate one of its subclasses to perform image analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest
type Request struct {
	objectivec.Object
}

// RequestFrom constructs a [Request] from an unsafe.Pointer.
//
// The abstract superclass for analysis requests.
func RequestFrom(ptr unsafe.Pointer) Request {
	return Request{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RequestClass) Alloc() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RequestClass) New() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Request) Init() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Request) Autorelease() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRequest creates a new Request instance.
func NewRequest() Request {
	return getRequestClass().New()
}


// Creates a new Vision request with an optional completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewRequestWithCompletionHandler(completionHandler unsafe.Pointer) Request {
	instance := getRequestClass().Alloc()
	rv := objc.Send[Request](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}


// The current revison supported by the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/currentRevision
func (rc _RequestClass) CurrentRevision() uint {
	rv := objc.Send[uint](objc.ID(rc.class), objc.Sel("currentRevision"))
	return rv
}
// The revision of the latest request for the particular SDK linked with the client application.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/defaultRevision
func (rc _RequestClass) DefaultRevision() uint {
	rv := objc.Send[uint](objc.ID(rc.class), objc.Sel("defaultRevision"))
	return rv
}
// The collection of currently-supported algorithm versions for the class of request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedRevisions
func (rc _RequestClass) SupportedRevisions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("supportedRevisions"))
	return rv
}
// Cancels the request before it can finish executing.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/cancel()
func (r_ Request) Cancel() {
	objc.Send[objc.ID](r_.ID, objc.Sel("cancel"))
}

// Returns the compute device for a compute stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/computeDeviceForComputeStage:
func (r_ Request) ComputeDeviceForComputeStage(computeStage unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("computeDeviceForComputeStage:"), computeStage)
	return rv
}

// Assigns a compute device for a compute stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/setComputeDevice:forComputeStage:
func (r_ Request) SetComputeDeviceForComputeStage(computeDevice objc.ID, computeStage unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setComputeDevice:forComputeStage:"), computeDevice, computeStage)
}

// The collection of compute devices per stage that a request supports.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedComputeStageDevicesAndReturnError:
func (r_ Request) SupportedComputeStageDevicesAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("supportedComputeStageDevicesAndReturnError:"), error_)
	return rv
}

// The completion handler the system invokes after the request finishes processing.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/completionHandler
func (r_ Request) CompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("completionHandler"))
	return rv
}

// The current revison supported by the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/currentRevision
func (r_ Request) CurrentRevision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("currentRevision"))
	return rv
}

// The revision of the latest request for the particular SDK linked with the client application.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/defaultRevision
func (r_ Request) DefaultRevision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("defaultRevision"))
	return rv
}

// A hint to minimize the resource burden of the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/preferBackgroundProcessing
func (r_ Request) PreferBackgroundProcessing() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("preferBackgroundProcessing"))
	return rv
}


// SetPreferBackgroundProcessing sets the value of the preferBackgroundProcessing property.
// A hint to minimize the resource burden of the request.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/preferBackgroundProcessing
func (r_ Request) SetPreferBackgroundProcessing(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreferBackgroundProcessing:"), value)
}
// The collection of observation results generated by request processing.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/results
func (r_ Request) Results() []Observation {
	rv := objc.Send[[]Observation](r_.ID, objc.Sel("results"))
	return rv
}

// The specific algorithm or implementation revision that’s used to perform the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/revision
func (r_ Request) Revision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("revision"))
	return rv
}


// SetRevision sets the value of the revision property.
// The specific algorithm or implementation revision that’s used to perform the request.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/revision
func (r_ Request) SetRevision(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRevision:"), value)
}
// The collection of currently-supported algorithm versions for the class of request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedRevisions
func (r_ Request) SupportedRevisions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("supportedRevisions"))
	return rv
}

// A Boolean signifying that the Vision request should execute exclusively on the CPU.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/usesCPUOnly
func (r_ Request) UsesCPUOnly() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("usesCPUOnly"))
	return rv
}


// SetUsesCPUOnly sets the value of the usesCPUOnly property.
// A Boolean signifying that the Vision request should execute exclusively on the CPU.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/usesCPUOnly
func (r_ Request) SetUsesCPUOnly(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUsesCPUOnly:"), value)
}

