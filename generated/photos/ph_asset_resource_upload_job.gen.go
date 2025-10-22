// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHAssetResourceUploadJob] class.
var (
	PHAssetResourceUploadJobClass     _PHAssetResourceUploadJobClass
	PHAssetResourceUploadJobClassOnce sync.Once
)

func getPHAssetResourceUploadJobClass() _PHAssetResourceUploadJobClass {
	PHAssetResourceUploadJobClassOnce.Do(func() {
		PHAssetResourceUploadJobClass = _PHAssetResourceUploadJobClass{objc.GetClass("PHAssetResourceUploadJob")}
	})
	return PHAssetResourceUploadJobClass
}

type _PHAssetResourceUploadJobClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResourceUploadJob] class.
type IPHAssetResourceUploadJob interface {
	IPHObject
	Destination() foundation.URLRequest
	Resource() PHAssetResource
	State() PHAssetResourceUploadJobState
}

// Represents a request to upload a
//
// Used within an application’s extension to represent a request to upload a to a destination . When the extensions principal class receives a call to background uploads, it can create new s using and any existing upload jobs can be fetched using and handled by updating their state using a to mark them as acknowledged, or to be retried. The maximum number of jobs that can be a in flight is limited to the .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob
type PHAssetResourceUploadJob struct {
	PHObject
}

// PHAssetResourceUploadJobFrom constructs a [PHAssetResourceUploadJob] from an unsafe.Pointer.
//
// Represents a request to upload a
func PHAssetResourceUploadJobFrom(ptr unsafe.Pointer) PHAssetResourceUploadJob {
	return PHAssetResourceUploadJob{
		PHObject: PHObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceUploadJobClass) Alloc() PHAssetResourceUploadJob {
	rv := objc.Send[PHAssetResourceUploadJob](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceUploadJobClass) New() PHAssetResourceUploadJob {
	rv := objc.Send[PHAssetResourceUploadJob](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResourceUploadJob) Init() PHAssetResourceUploadJob {
	rv := objc.Send[PHAssetResourceUploadJob](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResourceUploadJob) Autorelease() PHAssetResourceUploadJob {
	rv := objc.Send[PHAssetResourceUploadJob](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResourceUploadJob creates a new PHAssetResourceUploadJob instance.
func NewPHAssetResourceUploadJob() PHAssetResourceUploadJob {
	return getPHAssetResourceUploadJobClass().New()
}


// Returns all asset resource upload jobs applicable for a given action.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/fetchJobs(action:options:)
func (pc _PHAssetResourceUploadJobClass) FetchJobsWithActionOptions(action IPHAssetResourceUploadJobAction, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchJobsWithAction:options:"), action, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/jobLimit
func (pc _PHAssetResourceUploadJobClass) JobLimit() int {
	rv := objc.Send[int](objc.ID(pc.class), objc.Sel("jobLimit"))
	return rv
}
// The asset resource this upload job represents.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/destination
func (p_ PHAssetResourceUploadJob) Destination() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](p_.ID, objc.Sel("destination"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/jobLimit
func (p_ PHAssetResourceUploadJob) JobLimit() int {
	rv := objc.Send[int](p_.ID, objc.Sel("jobLimit"))
	return rv
}

// The maximum number of unacknowledged upload jobs allowed, this includes registered, pending, succeeded and failed jobs.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/resource
func (p_ PHAssetResourceUploadJob) Resource() PHAssetResource {
	rv := objc.Send[PHAssetResource](p_.ID, objc.Sel("resource"))
	return rv
}

// The destination to send this asset resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/state-swift.property
func (p_ PHAssetResourceUploadJob) State() PHAssetResourceUploadJobState {
	rv := objc.Send[PHAssetResourceUploadJobState](p_.ID, objc.Sel("state"))
	return rv
}



