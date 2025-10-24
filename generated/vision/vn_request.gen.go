// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNRequest */


/* debug [class_header]: Header for VNRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Request */
// An interface definition for the [Request] class.
type IRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Request */
	// properties:
	CompletionHandler() RequestCompletionHandler /* not a class type */
	PreferBackgroundProcessing() bool
	SetPreferBackgroundProcessing(value bool)
	Results() []Observation
	Revision() uint
	SetRevision(value uint)
	UsesCPUOnly() bool
	SetUsesCPUOnly(value bool)
	SupportedComputeStageDevices() ComputeDevice /* not a class type */
	SetSupportedComputeStageDevices(value ComputeDevice /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Request */
	// methods:
	Cancel()
	ComputeDeviceForComputeStage(computeStage ComputeStage /* typedef */) unsafe.Pointer
	SetComputeDeviceForComputeStage(computeDevice unsafe.Pointer, computeStage ComputeStage /* typedef */)
	SupportedComputeStageDevicesAndReturnError(error_ objectivec.IObject) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Request */
// Alloc allocates a new instance without initialization.
func (rc _RequestClass) Alloc() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Request */
// The abstract superclass for analysis requests.
//
// Other Vision request handlers that perform image analysis inherit from this abstract base class. Instantiate one of its subclasses to perform image analysis.


// The abstract superclass for analysis requests.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Request */

// Creates a new Vision request with an optional completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) Request {
	instance := getRequestClass().Alloc()
	rv := objc.Send[Request](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRequestWithCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Request */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Request */

// The current revison supported by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/currentRevision
func (rc _RequestClass) CurrentRevision() uint {
	rv := objc.Send[uint](objc.ID(rc.class), objc.Sel("currentRevision"))
	return rv
}/* debug [class_properties_class/property]: currentRevision */

// The revision of the latest request for the particular SDK linked with the client application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/defaultRevision
func (rc _RequestClass) DefaultRevision() uint {
	rv := objc.Send[uint](objc.ID(rc.class), objc.Sel("defaultRevision"))
	return rv
}/* debug [class_properties_class/property]: defaultRevision */

// The collection of currently-supported algorithm versions for the class of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedRevisions
func (rc _RequestClass) SupportedRevisions() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](objc.ID(rc.class), objc.Sel("supportedRevisions"))
	return rv
}/* debug [class_properties_class/property]: supportedRevisions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Request */

// Cancels the request before it can finish executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/cancel()
func (r_ Request) Cancel() {
	objc.Send[objc.ID](r_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Returns the compute device for a compute stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/computeDeviceForComputeStage:
func (r_ Request) ComputeDeviceForComputeStage(computeStage ComputeStage /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("computeDeviceForComputeStage:"), computeStage)
	return rv
}/* debug [instance_methods/method]: ComputeDeviceForComputeStage */


// Assigns a compute device for a compute stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/setComputeDevice:forComputeStage:
func (r_ Request) SetComputeDeviceForComputeStage(computeDevice unsafe.Pointer, computeStage ComputeStage /* typedef */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setComputeDevice:forComputeStage:"), computeDevice, computeStage)
}/* debug [instance_methods/method]: SetComputeDeviceForComputeStage */


// The collection of compute devices per stage that a request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedComputeStageDevicesAndReturnError:
func (r_ Request) SupportedComputeStageDevicesAndReturnError(error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("supportedComputeStageDevicesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedComputeStageDevicesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Request */

// The completion handler the system invokes after the request finishes processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/completionHandler
func (r_ Request) CompletionHandler() RequestCompletionHandler /* not a class type */ {
	rv := objc.Send[RequestCompletionHandler](r_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// The current revison supported by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/currentRevision
func (r_ Request) CurrentRevision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("currentRevision"))
	return rv
}/* debug [instance_properties/getter]: currentRevision */


// The revision of the latest request for the particular SDK linked with the client application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/defaultRevision
func (r_ Request) DefaultRevision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("defaultRevision"))
	return rv
}/* debug [instance_properties/getter]: defaultRevision */


// A hint to minimize the resource burden of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/preferBackgroundProcessing
func (r_ Request) PreferBackgroundProcessing() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("preferBackgroundProcessing"))
	return rv
}/* debug [instance_properties/getter]: preferBackgroundProcessing */


// A hint to minimize the resource burden of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/preferBackgroundProcessing
func (r_ Request) SetPreferBackgroundProcessing(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreferBackgroundProcessing:"), value)
}/* debug [instance_properties/setter]: preferBackgroundProcessing */


// The collection of observation results generated by request processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/results
func (r_ Request) Results() []Observation {
	rv := objc.Send[[]Observation](r_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The specific algorithm or implementation revision that’s used to perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/revision
func (r_ Request) Revision() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("revision"))
	return rv
}/* debug [instance_properties/getter]: revision */


// The specific algorithm or implementation revision that’s used to perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/revision
func (r_ Request) SetRevision(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRevision:"), value)
}/* debug [instance_properties/setter]: revision */


// The collection of currently-supported algorithm versions for the class of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/supportedRevisions
func (r_ Request) SupportedRevisions() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](r_.ID, objc.Sel("supportedRevisions"))
	return rv
}/* debug [instance_properties/getter]: supportedRevisions */


// A Boolean signifying that the Vision request should execute exclusively on the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/usesCPUOnly
func (r_ Request) UsesCPUOnly() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("usesCPUOnly"))
	return rv
}/* debug [instance_properties/getter]: usesCPUOnly */


// A Boolean signifying that the Vision request should execute exclusively on the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequest/usesCPUOnly
func (r_ Request) SetUsesCPUOnly(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUsesCPUOnly:"), value)
}/* debug [instance_properties/setter]: usesCPUOnly */


// The collection of compute devices per stage that a request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/supportedcomputestagedevices
func (r_ Request) SupportedComputeStageDevices() ComputeDevice /* not a class type */ {
	rv := objc.Send[ComputeDevice](r_.ID, objc.Sel("supportedComputeStageDevices"))
	return rv
}/* debug [instance_properties/getter]: supportedComputeStageDevices */


// The collection of compute devices per stage that a request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/supportedcomputestagedevices
func (r_ Request) SetSupportedComputeStageDevices(value ComputeDevice /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportedComputeStageDevices:"), value)
}/* debug [instance_properties/setter]: supportedComputeStageDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRequest */


