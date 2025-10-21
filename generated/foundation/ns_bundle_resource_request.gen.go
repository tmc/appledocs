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
	BeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer)
	ConditionallyBeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer)
	EndAccessingResources()
}

// A resource manager you use to download content hosted on the App Store at the time your app needs it.
//
// You identify on-demand resources during development by creating string identifiers known as tags and assigning one or more tags to each resource. An object manages the resources marked by one or more tags. You use the resource request to inform the system when the managed tags are needed and when you have finished accessing them. The resource request manages the downloading of any resources marked with the managed tags that are not already on the device and informs your app when the resources are ready for use. The system will not attempt to purge the resources marked with a tag from on-device storage as long as at least one object is managing the tag. Apps can access resources after the completion handler of either or is called successfully. Management ends after a call to or after the resource request object is deallocated. Other properties and methods let you track the progress of a download, change the priority of a download, and check whether the resources marked by a set of tags are already on the device. Methods in indicate to the system the relative importance of preserving a tag in memory after it is no longer in use. For more information, see and .
//
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

// Initializes a resource request for managing the on-demand resources marked with any of the set of specified tags. The managed resources are loaded into the main bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/init(tags:)
func NewBundleResourceRequestWithTags(tags unsafe.Pointer) BundleResourceRequest {
	instance := getBundleResourceRequestClass().Alloc()
	rv := objc.Send[BundleResourceRequest](instance.ID, objc.Sel("initWithTags:"), tags)
	rv.Autorelease()
	return rv
}

// Requests access to the resources marked with the managed tags. If any of the resources are not on the device, they are requested from the App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/beginAccessingResources(completionHandler:)
func (b_ BundleResourceRequest) BeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("beginAccessingResourcesWithCompletionHandler:"), completionHandler)
}

// Checks whether the resources marked with the tags managed by the request are already on the device. If all of the resources are on the device, you can begin accessing those resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/conditionallyBeginAccessingResources(completionHandler:)
func (b_ BundleResourceRequest) ConditionallyBeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("conditionallyBeginAccessingResourcesWithCompletionHandler:"), completionHandler)
}

// Informs the system that you have finished accessing the resources marked with the tags managed by the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/endAccessingResources()
func (b_ BundleResourceRequest) EndAccessingResources() {
	objc.Send[objc.ID](b_.ID, objc.Sel("endAccessingResources"))
}

// A reference to the progress object associated with the specified resource request. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/progress
func (b_ BundleResourceRequest) Progress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("progress"))
	return rv
}
