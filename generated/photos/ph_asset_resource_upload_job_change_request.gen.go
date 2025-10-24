// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHAssetResourceUploadJobChangeRequest] class.
var (
	PHAssetResourceUploadJobChangeRequestClass     _PHAssetResourceUploadJobChangeRequestClass
	PHAssetResourceUploadJobChangeRequestClassOnce sync.Once
)

func getPHAssetResourceUploadJobChangeRequestClass() _PHAssetResourceUploadJobChangeRequestClass {
	PHAssetResourceUploadJobChangeRequestClassOnce.Do(func() {
		PHAssetResourceUploadJobChangeRequestClass = _PHAssetResourceUploadJobChangeRequestClass{objc.GetClass("PHAssetResourceUploadJobChangeRequest")}
	})
	return PHAssetResourceUploadJobChangeRequestClass
}

type _PHAssetResourceUploadJobChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResourceUploadJobChangeRequest] class.
type IPHAssetResourceUploadJobChangeRequest interface {
	IPHChangeRequest
	Acknowledge()
	RetryWithDestination(destination foundation.IURLRequest)
}

// Used within an application’s extension to create and manage records
//
// When the extensions principal class receives a call to background uploads, it can create new s through calls to perform changes on a PHPhotoLibrary using and any in-flight upload jobs can be handled by updating their state to mark them as acknowledged, or to be retried. The maximum number of jobs that can be in flight is limited to the . can only be created or used within a photo library change block. For details on change blocks, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest
type PHAssetResourceUploadJobChangeRequest struct {
	PHChangeRequest
}

// PHAssetResourceUploadJobChangeRequestFrom constructs a [PHAssetResourceUploadJobChangeRequest] from an unsafe.Pointer.
//
// Used within an application’s extension to create and manage records
func PHAssetResourceUploadJobChangeRequestFrom(ptr unsafe.Pointer) PHAssetResourceUploadJobChangeRequest {
	return PHAssetResourceUploadJobChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceUploadJobChangeRequestClass) Alloc() PHAssetResourceUploadJobChangeRequest {
	rv := objc.Send[PHAssetResourceUploadJobChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceUploadJobChangeRequestClass) New() PHAssetResourceUploadJobChangeRequest {
	rv := objc.Send[PHAssetResourceUploadJobChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResourceUploadJobChangeRequest) Init() PHAssetResourceUploadJobChangeRequest {
	rv := objc.Send[PHAssetResourceUploadJobChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResourceUploadJobChangeRequest) Autorelease() PHAssetResourceUploadJobChangeRequest {
	rv := objc.Send[PHAssetResourceUploadJobChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResourceUploadJobChangeRequest creates a new PHAssetResourceUploadJobChangeRequest instance.
func NewPHAssetResourceUploadJobChangeRequest() PHAssetResourceUploadJobChangeRequest {
	return getPHAssetResourceUploadJobChangeRequestClass().New()
}

// Creates a request for modifying the specified upload job.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest/init(for:)
func NewPHAssetResourceUploadJobChangeRequestForUploadJob(job IPHAssetResourceUploadJob) PHAssetResourceUploadJobChangeRequest {
	rv := objc.Send[PHAssetResourceUploadJobChangeRequest](objc.ID(getPHAssetResourceUploadJobChangeRequestClass().class), objc.Sel("changeRequestForUploadJob:"), job)
	return rv
}

// Used to create an asset resource upload job.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest/createJob(destination:resource:)
func (pc _PHAssetResourceUploadJobChangeRequestClass) CreateJobWithDestinationResource(destination foundation.IURLRequest, resource IPHAssetResource) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("createJobWithDestination:resource:"), destination, resource)
}

// Creates a request for modifying the specified upload job.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest/init(for:)
func (pc _PHAssetResourceUploadJobChangeRequestClass) ChangeRequestForUploadJob(job IPHAssetResourceUploadJob) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("changeRequestForUploadJob:"), job)
	return rv
}

// Acknowledges a successful or failed job. Jobs must be acknowledged to free up space for .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest/acknowledge()
func (p_ PHAssetResourceUploadJobChangeRequest) Acknowledge() {
	objc.Send[objc.ID](p_.ID, objc.Sel("acknowledge"))
}

// Retries a job that is failed, unacknowledged, and has not been retried before. Successful retries also free up space for .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJobChangeRequest/retry(destination:)
func (p_ PHAssetResourceUploadJobChangeRequest) RetryWithDestination(destination foundation.IURLRequest) {
	objc.Send[objc.ID](p_.ID, objc.Sel("retryWithDestination:"), destination)
}
