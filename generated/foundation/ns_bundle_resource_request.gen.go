// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BundleResourceRequest] class.
var (
	BundleResourceRequestClass     _BundleResourceRequestClass
	BundleResourceRequestClassOnce sync.Once
)

func getBundleResourceRequestClass() _BundleResourceRequestClass {
	BundleResourceRequestClassOnce.Do(func() {
		BundleResourceRequestClass = _BundleResourceRequestClass{objc.GetClass("NSBundleResourceRequest")}
	})
	return BundleResourceRequestClass
}

type _BundleResourceRequestClass struct {
	class objc.Class
}

// An interface definition for the [BundleResourceRequest] class.
type IBundleResourceRequest interface {
	objectivec.IObject
	// properties:
	NSBundleErrorMaximum() int /* primitive/slice/pointer. */
	SetNSBundleErrorMaximum(value int /* primitive/slice/pointer. */)
	NSBundleErrorMinimum() int /* primitive/slice/pointer. */
	SetNSBundleErrorMinimum(value int /* primitive/slice/pointer. */)
	NSBundleOnDemandResourceExceededMaximumSizeError() int /* primitive/slice/pointer. */
	SetNSBundleOnDemandResourceExceededMaximumSizeError(value int /* primitive/slice/pointer. */)
	NSBundleOnDemandResourceInvalidTagError() int /* primitive/slice/pointer. */
	SetNSBundleOnDemandResourceInvalidTagError(value int /* primitive/slice/pointer. */)
	NSBundleOnDemandResourceOutOfSpaceError() int /* primitive/slice/pointer. */
	SetNSBundleOnDemandResourceOutOfSpaceError(value int /* primitive/slice/pointer. */)
	Bundle() IBundle
	SetBundle(value IBundle)
	LoadingPriority() float64 /* primitive/slice/pointer. */
	SetLoadingPriority(value float64 /* primitive/slice/pointer. */)
	Progress() IProgress
	SetProgress(value IProgress)
	Tags() IString
	SetTags(value IString)
	NSBundleResourceRequestLoadingPriorityUrgent() float64 /* primitive/slice/pointer. */
	// methods:
}

// A resource manager you use to download content hosted on the App Store at the time your app needs it.
//
// You identify on-demand resources during development by creating string identifiers known as tags and assigning one or more tags to each resource. An object manages the resources marked by one or more tags. You use the resource request to inform the system when the managed tags are needed and when you have finished accessing them. The resource request manages the downloading of any resources marked with the managed tags that are not already on the device and informs your app when the resources are ready for use. The system will not attempt to purge the resources marked with a tag from on-device storage as long as at least one object is managing the tag. Apps can access resources after the completion handler of either or is called successfully. Management ends after a call to or after the resource request object is deallocated. Other properties and methods let you track the progress of a download, change the priority of a download, and check whether the resources marked by a set of tags are already on the device. Methods in indicate to the system the relative importance of preserving a tag in memory after it is no longer in use. For more information, see and .


// A resource manager you use to download content hosted on the App Store at the time your app needs it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest
type BundleResourceRequest struct {
	objectivec.Object
}

// BundleResourceRequestFrom constructs a [BundleResourceRequest] from an unsafe.Pointer.
//
// A resource manager you use to download content hosted on the App Store at the time your app needs it.
func BundleResourceRequestFrom(ptr unsafe.Pointer) BundleResourceRequest {
	return BundleResourceRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BundleResourceRequestClass) Alloc() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BundleResourceRequestClass) New() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BundleResourceRequest) Init() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BundleResourceRequest) Autorelease() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBundleResourceRequest creates a new BundleResourceRequest instance.
func NewBundleResourceRequest() BundleResourceRequest {
	return getBundleResourceRequestClass().New()
}



// The end of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrormaximum-swift.var
func (b_ BundleResourceRequest) NSBundleErrorMaximum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrormaximum-swift.var
func (b_ BundleResourceRequest) SetNSBundleErrorMaximum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleErrorMaximum:"), value)
}


// The start of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrorminimum-swift.var
func (b_ BundleResourceRequest) NSBundleErrorMinimum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrorminimum-swift.var
func (b_ BundleResourceRequest) SetNSBundleErrorMinimum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleErrorMinimum:"), value)
}


// The application exceeded the amount of on-demand resources content in use at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceexceededmaximumsizeerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceExceededMaximumSizeError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceExceededMaximumSizeError"))
	return rv
}


// The application exceeded the amount of on-demand resources content in use at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceexceededmaximumsizeerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceExceededMaximumSizeError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceExceededMaximumSizeError:"), value)
}


// The application specified a tag that the system couldn’t find in the application tag manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceinvalidtagerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceInvalidTagError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceInvalidTagError"))
	return rv
}


// The application specified a tag that the system couldn’t find in the application tag manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceinvalidtagerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceInvalidTagError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceInvalidTagError:"), value)
}


// Insufficient space available to download the requested on-demand resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceoutofspaceerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceOutOfSpaceError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceOutOfSpaceError"))
	return rv
}


// Insufficient space available to download the requested on-demand resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceoutofspaceerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceOutOfSpaceError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceOutOfSpaceError:"), value)
}


// A reference to the bundle used for storing the downloaded resources. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/bundle
func (b_ BundleResourceRequest) Bundle() IBundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("bundle"))
	return rv
}


// A reference to the bundle used for storing the downloaded resources. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/bundle
func (b_ BundleResourceRequest) SetBundle(value IBundle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBundle:"), value)
}


// A hint to the system of the relative priority of the resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/loadingpriority
func (b_ BundleResourceRequest) LoadingPriority() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("loadingPriority"))
	return rv
}


// A hint to the system of the relative priority of the resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/loadingpriority
func (b_ BundleResourceRequest) SetLoadingPriority(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLoadingPriority:"), value)
}


// A reference to the progress object associated with the specified resource request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/progress
func (b_ BundleResourceRequest) Progress() IProgress {
	rv := objc.Send[Progress](b_.ID, objc.Sel("progress"))
	return rv
}


// A reference to the progress object associated with the specified resource request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/progress
func (b_ BundleResourceRequest) SetProgress(value IProgress) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setProgress:"), value)
}


// A set of strings, with each string specifying a tag used to mark on-demand resources managed by the request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/tags
func (b_ BundleResourceRequest) Tags() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("tags"))
	return rv
}


// A set of strings, with each string specifying a tag used to mark on-demand resources managed by the request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/tags
func (b_ BundleResourceRequest) SetTags(value IString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTags:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequestloadingpriorityurgent
func (b_ BundleResourceRequest) NSBundleResourceRequestLoadingPriorityUrgent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("NSBundleResourceRequestLoadingPriorityUrgent"))
	return rv
}



