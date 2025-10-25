// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetResourceLoadingRequest */


/* debug [class_header]: Header for AVAssetResourceLoadingRequest */
// The class instance for the [AssetResourceLoadingRequest] class.
var (
	AssetResourceLoadingRequestClass     _AssetResourceLoadingRequestClass
	AssetResourceLoadingRequestClassOnce sync.Once
)

func getAssetResourceLoadingRequestClass() _AssetResourceLoadingRequestClass {
	AssetResourceLoadingRequestClassOnce.Do(func() {
		AssetResourceLoadingRequestClass = _AssetResourceLoadingRequestClass{objc.GetClass("AVAssetResourceLoadingRequest")}
	})
	return AssetResourceLoadingRequestClass
}

type _AssetResourceLoadingRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetResourceLoadingRequest */
// An interface definition for the [AssetResourceLoadingRequest] class.
type IAssetResourceLoadingRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetResourceLoadingRequest */
	// properties:
	ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest
	DataRequest() IAVAssetResourceLoadingDataRequest
	Cancelled() bool
	Finished() bool
	Redirect() foundation.URLRequest
	SetRedirect(value foundation.URLRequest)
	Request() foundation.URLRequest
	Requestor() IAVAssetResourceLoadingRequestor
	Response() foundation.URLResponse
	SetResponse(value foundation.URLResponse)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsFinished() bool
	SetIsFinished(value bool)
	AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetResourceLoadingRequest */
	// methods:
	FinishLoading()
	FinishLoadingWithError(error_ Error)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetResourceLoadingRequest */
// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoadingRequestClass) Alloc() AssetResourceLoadingRequest {
	rv := objc.Send[AssetResourceLoadingRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetResourceLoadingRequestClass) New() AssetResourceLoadingRequest {
	rv := objc.Send[AssetResourceLoadingRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceLoadingRequest) Init() AssetResourceLoadingRequest {
	rv := objc.Send[AssetResourceLoadingRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceLoadingRequest) Autorelease() AssetResourceLoadingRequest {
	rv := objc.Send[AssetResourceLoadingRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceLoadingRequest creates a new AssetResourceLoadingRequest instance.
func NewAssetResourceLoadingRequest() AssetResourceLoadingRequest {
	return getAssetResourceLoadingRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetResourceLoadingRequest */
// An object that encapsulates information about a resource request from a resource loader object.
//
// When an object needs help loading a resource, it asks its object to assist. The resource loader encapsulates the request information by creating an instance of this object, which it then hands to its delegate object for processing. The delegate uses the information in this object to perform the request and report on the success or failure of the operation.


// An object that encapsulates information about a resource request from a resource loader object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest
type AssetResourceLoadingRequest struct {
	objectivec.Object
}

// AssetResourceLoadingRequestFrom constructs a [AssetResourceLoadingRequest] from an unsafe.Pointer.
//
// An object that encapsulates information about a resource request from a resource loader object.
func AssetResourceLoadingRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingRequest {
	return AssetResourceLoadingRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetResourceLoadingRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetResourceLoadingRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetResourceLoadingRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetResourceLoadingRequest */

// Causes the receiver to treat the processing of the request as complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading()
func (a_ AssetResourceLoadingRequest) FinishLoading() {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishLoading"))
}/* debug [instance_methods/method]: FinishLoading */


// Causes the receiver to handle the failure to load a resource for which a resource loader’s delegate took responsibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading(with:)
func (a_ AssetResourceLoadingRequest) FinishLoadingWithError(error_ Error) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishLoadingWithError:"), error_)
}/* debug [instance_methods/method]: FinishLoadingWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetResourceLoadingRequest */

// The information for a requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/contentInformationRequest
func (a_ AssetResourceLoadingRequest) ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](a_.ID, objc.Sel("contentInformationRequest"))
	return rv
}/* debug [instance_properties/getter]: contentInformationRequest */


// The range of requested resource data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/dataRequest
func (a_ AssetResourceLoadingRequest) DataRequest() IAVAssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](a_.ID, objc.Sel("dataRequest"))
	return rv
}/* debug [instance_properties/getter]: dataRequest */


// A Boolean value that indicates whether the request has been cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/isCancelled
func (a_ AssetResourceLoadingRequest) Cancelled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates whether loading of the resource has finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/isFinished
func (a_ AssetResourceLoadingRequest) Finished() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("finished"))
	return rv
}/* debug [instance_properties/getter]: finished */


// An URL request instance if the loading request was redirected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/redirect
func (a_ AssetResourceLoadingRequest) Redirect() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](a_.ID, objc.Sel("redirect"))
	return rv
}/* debug [instance_properties/getter]: redirect */


// An URL request instance if the loading request was redirected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/redirect
func (a_ AssetResourceLoadingRequest) SetRedirect(value foundation.URLRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRedirect:"), value)
}/* debug [instance_properties/setter]: redirect */


// The URL request object for the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/request
func (a_ AssetResourceLoadingRequest) Request() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](a_.ID, objc.Sel("request"))
	return rv
}/* debug [instance_properties/getter]: request */


// The asset resource requestor that made the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/requestor
func (a_ AssetResourceLoadingRequest) Requestor() IAVAssetResourceLoadingRequestor {
	rv := objc.Send[AssetResourceLoadingRequestor](a_.ID, objc.Sel("requestor"))
	return rv
}/* debug [instance_properties/getter]: requestor */


// The URL response for the loading request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/response
func (a_ AssetResourceLoadingRequest) Response() foundation.URLResponse {
	rv := objc.Send[foundation.URLResponse](a_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The URL response for the loading request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/response
func (a_ AssetResourceLoadingRequest) SetResponse(value foundation.URLResponse) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setResponse:"), value)
}/* debug [instance_properties/setter]: response */


// A Boolean value that indicates whether the request has been cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/iscancelled
func (a_ AssetResourceLoadingRequest) IsCancelled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates whether the request has been cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/iscancelled
func (a_ AssetResourceLoadingRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates whether loading of the resource has finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/isfinished
func (a_ AssetResourceLoadingRequest) IsFinished() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isFinished"))
	return rv
}/* debug [instance_properties/getter]: isFinished */


// A Boolean value that indicates whether loading of the resource has finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/isfinished
func (a_ AssetResourceLoadingRequest) SetIsFinished(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsFinished:"), value)
}/* debug [instance_properties/setter]: isFinished */


// Specifies whether the content key request requires a persistable key to be returned from the key vendor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequeststreamingcontentkeyrequestrequirespersistentkey
func (a_ AssetResourceLoadingRequest) AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetResourceLoadingRequest */



