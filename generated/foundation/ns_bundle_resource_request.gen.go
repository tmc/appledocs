// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSBundleResourceRequest */


/* debug [class_header]: Header for NSBundleResourceRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BundleResourceRequest */
// An interface definition for the [BundleResourceRequest] class.
type IBundleResourceRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BundleResourceRequest */
	// properties:
	NSBundleErrorMaximum() int
	SetNSBundleErrorMaximum(value int)
	NSBundleErrorMinimum() int
	SetNSBundleErrorMinimum(value int)
	NSBundleOnDemandResourceExceededMaximumSizeError() int
	SetNSBundleOnDemandResourceExceededMaximumSizeError(value int)
	NSBundleOnDemandResourceInvalidTagError() int
	SetNSBundleOnDemandResourceInvalidTagError(value int)
	NSBundleOnDemandResourceOutOfSpaceError() int
	SetNSBundleOnDemandResourceOutOfSpaceError(value int)
	Bundle() IBundle
	SetBundle(value IBundle)
	LoadingPriority() float64
	SetLoadingPriority(value float64)
	Progress() IProgress
	SetProgress(value IProgress)
	Tags() IString
	SetTags(value IString)
	NSBundleResourceRequestLoadingPriorityUrgent() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BundleResourceRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BundleResourceRequest */
// Alloc allocates a new instance without initialization.
func (bc _BundleResourceRequestClass) Alloc() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BundleResourceRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BundleResourceRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BundleResourceRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BundleResourceRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BundleResourceRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BundleResourceRequest */

// The end of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrormaximum-swift.var
func (b_ BundleResourceRequest) NSBundleErrorMaximum() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSBundleErrorMaximum */


// The end of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrormaximum-swift.var
func (b_ BundleResourceRequest) SetNSBundleErrorMaximum(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSBundleErrorMaximum */


// The start of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrorminimum-swift.var
func (b_ BundleResourceRequest) NSBundleErrorMinimum() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSBundleErrorMinimum */


// The start of the range of error codes reserved for bundle errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleerrorminimum-swift.var
func (b_ BundleResourceRequest) SetNSBundleErrorMinimum(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSBundleErrorMinimum */


// The application exceeded the amount of on-demand resources content in use at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceexceededmaximumsizeerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceExceededMaximumSizeError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceExceededMaximumSizeError"))
	return rv
}/* debug [instance_properties/getter]: NSBundleOnDemandResourceExceededMaximumSizeError */


// The application exceeded the amount of on-demand resources content in use at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceexceededmaximumsizeerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceExceededMaximumSizeError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceExceededMaximumSizeError:"), value)
}/* debug [instance_properties/setter]: NSBundleOnDemandResourceExceededMaximumSizeError */


// The application specified a tag that the system couldn’t find in the application tag manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceinvalidtagerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceInvalidTagError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceInvalidTagError"))
	return rv
}/* debug [instance_properties/getter]: NSBundleOnDemandResourceInvalidTagError */


// The application specified a tag that the system couldn’t find in the application tag manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceinvalidtagerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceInvalidTagError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceInvalidTagError:"), value)
}/* debug [instance_properties/setter]: NSBundleOnDemandResourceInvalidTagError */


// Insufficient space available to download the requested on-demand resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceoutofspaceerror-swift.var
func (b_ BundleResourceRequest) NSBundleOnDemandResourceOutOfSpaceError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSBundleOnDemandResourceOutOfSpaceError"))
	return rv
}/* debug [instance_properties/getter]: NSBundleOnDemandResourceOutOfSpaceError */


// Insufficient space available to download the requested on-demand resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleondemandresourceoutofspaceerror-swift.var
func (b_ BundleResourceRequest) SetNSBundleOnDemandResourceOutOfSpaceError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSBundleOnDemandResourceOutOfSpaceError:"), value)
}/* debug [instance_properties/setter]: NSBundleOnDemandResourceOutOfSpaceError */


// A reference to the bundle used for storing the downloaded resources. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/bundle
func (b_ BundleResourceRequest) Bundle() IBundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("bundle"))
	return rv
}/* debug [instance_properties/getter]: bundle */


// A reference to the bundle used for storing the downloaded resources. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/bundle
func (b_ BundleResourceRequest) SetBundle(value IBundle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBundle:"), value)
}/* debug [instance_properties/setter]: bundle */


// A hint to the system of the relative priority of the resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/loadingpriority
func (b_ BundleResourceRequest) LoadingPriority() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("loadingPriority"))
	return rv
}/* debug [instance_properties/getter]: loadingPriority */


// A hint to the system of the relative priority of the resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/loadingpriority
func (b_ BundleResourceRequest) SetLoadingPriority(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLoadingPriority:"), value)
}/* debug [instance_properties/setter]: loadingPriority */


// A reference to the progress object associated with the specified resource request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/progress
func (b_ BundleResourceRequest) Progress() IProgress {
	rv := objc.Send[Progress](b_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */


// A reference to the progress object associated with the specified resource request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/progress
func (b_ BundleResourceRequest) SetProgress(value IProgress) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setProgress:"), value)
}/* debug [instance_properties/setter]: progress */


// A set of strings, with each string specifying a tag used to mark on-demand resources managed by the request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/tags
func (b_ BundleResourceRequest) Tags() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("tags"))
	return rv
}/* debug [instance_properties/getter]: tags */


// A set of strings, with each string specifying a tag used to mark on-demand resources managed by the request. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequest/tags
func (b_ BundleResourceRequest) SetTags(value IString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTags:"), value)
}/* debug [instance_properties/setter]: tags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundleresourcerequestloadingpriorityurgent
func (b_ BundleResourceRequest) NSBundleResourceRequestLoadingPriorityUrgent() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("NSBundleResourceRequestLoadingPriorityUrgent"))
	return rv
}/* debug [instance_properties/getter]: NSBundleResourceRequestLoadingPriorityUrgent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBundleResourceRequest */


