// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetResourceLoadingContentInformationRequest */


/* debug [class_header]: Header for AVAssetResourceLoadingContentInformationRequest */
// The class instance for the [AssetResourceLoadingContentInformationRequest] class.
var (
	AssetResourceLoadingContentInformationRequestClass     _AssetResourceLoadingContentInformationRequestClass
	AssetResourceLoadingContentInformationRequestClassOnce sync.Once
)

func getAssetResourceLoadingContentInformationRequestClass() _AssetResourceLoadingContentInformationRequestClass {
	AssetResourceLoadingContentInformationRequestClassOnce.Do(func() {
		AssetResourceLoadingContentInformationRequestClass = _AssetResourceLoadingContentInformationRequestClass{objc.GetClass("AVAssetResourceLoadingContentInformationRequest")}
	})
	return AssetResourceLoadingContentInformationRequestClass
}

type _AssetResourceLoadingContentInformationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetResourceLoadingContentInformationRequest */
// An interface definition for the [AssetResourceLoadingContentInformationRequest] class.
type IAssetResourceLoadingContentInformationRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetResourceLoadingContentInformationRequest */
	// properties:
	AllowedContentTypes() []string
	ContentLength() objectivec.IObject
	SetContentLength(value objectivec.IObject)
	ContentType() objc.IObject /* cross-framework: NSString */
	SetContentType(value objc.IObject /* cross-framework: NSString */)
	ByteRangeAccessSupported() bool
	SetByteRangeAccessSupported(value bool)
	EntireLengthAvailableOnDemand() bool
	SetEntireLengthAvailableOnDemand(value bool)
	RenewalDate() objc.IObject /* cross-framework: NSDate */
	SetRenewalDate(value objc.IObject /* cross-framework: NSDate */)
	IsByteRangeAccessSupported() bool
	SetIsByteRangeAccessSupported(value bool)
	IsEntireLengthAvailableOnDemand() bool
	SetIsEntireLengthAvailableOnDemand(value bool)
	ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest
	SetContentInformationRequest(value IAVAssetResourceLoadingContentInformationRequest)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetResourceLoadingContentInformationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetResourceLoadingContentInformationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoadingContentInformationRequestClass) Alloc() AssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetResourceLoadingContentInformationRequestClass) New() AssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceLoadingContentInformationRequest) Init() AssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceLoadingContentInformationRequest) Autorelease() AssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceLoadingContentInformationRequest creates a new AssetResourceLoadingContentInformationRequest instance.
func NewAssetResourceLoadingContentInformationRequest() AssetResourceLoadingContentInformationRequest {
	return getAssetResourceLoadingContentInformationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetResourceLoadingContentInformationRequest */
// A query for retrieving essential information about a resource that an asset resource-loading request references.
//
// When a resource loading delegate, which must implement the  protocol, receives an instance of  when the is invoked and accepts responsibility for loading the resource, it must check whether the  property of the  is not . Whenever the value is not , the request includes a query for the information that  encapsulates. In response to such queries, the resource loading delegate should set the values of the content information request’s properties appropriately before invoking the method . When is invoked, the values of the properties of its  property will, in part, determine how the requested resource is processed. For example, if the requested resource’s URL is the URL of an and is set by the resource loading delegate to a value that the underlying media system doesn’t recognize as a supported media file type, operations on the , such as playback, are likely to fail.


// A query for retrieving essential information about a resource that an asset resource-loading request references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest
type AssetResourceLoadingContentInformationRequest struct {
	objectivec.Object
}

// AssetResourceLoadingContentInformationRequestFrom constructs a [AssetResourceLoadingContentInformationRequest] from an unsafe.Pointer.
//
// A query for retrieving essential information about a resource that an asset resource-loading request references.
func AssetResourceLoadingContentInformationRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingContentInformationRequest {
	return AssetResourceLoadingContentInformationRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetResourceLoadingContentInformationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetResourceLoadingContentInformationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetResourceLoadingContentInformationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetResourceLoadingContentInformationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetResourceLoadingContentInformationRequest */

// The types of data that are accepted as a valid response for the requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/allowedContentTypes
func (a_ AssetResourceLoadingContentInformationRequest) AllowedContentTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("allowedContentTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedContentTypes */


// The length, in bytes, of the requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentLength
func (a_ AssetResourceLoadingContentInformationRequest) ContentLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("contentLength"))
	return rv
}/* debug [instance_properties/getter]: contentLength */


// The length, in bytes, of the requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentLength
func (a_ AssetResourceLoadingContentInformationRequest) SetContentLength(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentLength:"), value)
}/* debug [instance_properties/setter]: contentLength */


// The UTI that specifies the type of data contained by the requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentType
func (a_ AssetResourceLoadingContentInformationRequest) ContentType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The UTI that specifies the type of data contained by the requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/contentType
func (a_ AssetResourceLoadingContentInformationRequest) SetContentType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentType:"), value)
}/* debug [instance_properties/setter]: contentType */


// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isByteRangeAccessSupported
func (a_ AssetResourceLoadingContentInformationRequest) ByteRangeAccessSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("byteRangeAccessSupported"))
	return rv
}/* debug [instance_properties/getter]: byteRangeAccessSupported */


// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isByteRangeAccessSupported
func (a_ AssetResourceLoadingContentInformationRequest) SetByteRangeAccessSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setByteRangeAccessSupported:"), value)
}/* debug [instance_properties/setter]: byteRangeAccessSupported */


// A Boolean value that indicates whether asset data loading can expect data immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isEntireLengthAvailableOnDemand
func (a_ AssetResourceLoadingContentInformationRequest) EntireLengthAvailableOnDemand() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("entireLengthAvailableOnDemand"))
	return rv
}/* debug [instance_properties/getter]: entireLengthAvailableOnDemand */


// A Boolean value that indicates whether asset data loading can expect data immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/isEntireLengthAvailableOnDemand
func (a_ AssetResourceLoadingContentInformationRequest) SetEntireLengthAvailableOnDemand(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEntireLengthAvailableOnDemand:"), value)
}/* debug [instance_properties/setter]: entireLengthAvailableOnDemand */


// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/renewalDate
func (a_ AssetResourceLoadingContentInformationRequest) RenewalDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](a_.ID, objc.Sel("renewalDate"))
	return rv
}/* debug [instance_properties/getter]: renewalDate */


// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingContentInformationRequest/renewalDate
func (a_ AssetResourceLoadingContentInformationRequest) SetRenewalDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenewalDate:"), value)
}/* debug [instance_properties/setter]: renewalDate */


// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/isbyterangeaccesssupported
func (a_ AssetResourceLoadingContentInformationRequest) IsByteRangeAccessSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isByteRangeAccessSupported"))
	return rv
}/* debug [instance_properties/getter]: isByteRangeAccessSupported */


// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/isbyterangeaccesssupported
func (a_ AssetResourceLoadingContentInformationRequest) SetIsByteRangeAccessSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsByteRangeAccessSupported:"), value)
}/* debug [instance_properties/setter]: isByteRangeAccessSupported */


// A Boolean value that indicates whether asset data loading can expect data immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/isentirelengthavailableondemand
func (a_ AssetResourceLoadingContentInformationRequest) IsEntireLengthAvailableOnDemand() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEntireLengthAvailableOnDemand"))
	return rv
}/* debug [instance_properties/getter]: isEntireLengthAvailableOnDemand */


// A Boolean value that indicates whether asset data loading can expect data immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/isentirelengthavailableondemand
func (a_ AssetResourceLoadingContentInformationRequest) SetIsEntireLengthAvailableOnDemand(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEntireLengthAvailableOnDemand:"), value)
}/* debug [instance_properties/setter]: isEntireLengthAvailableOnDemand */


// The information for a requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/contentinformationrequest
func (a_ AssetResourceLoadingContentInformationRequest) ContentInformationRequest() IAVAssetResourceLoadingContentInformationRequest {
	rv := objc.Send[AssetResourceLoadingContentInformationRequest](a_.ID, objc.Sel("contentInformationRequest"))
	return rv
}/* debug [instance_properties/getter]: contentInformationRequest */


// The information for a requested resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/contentinformationrequest
func (a_ AssetResourceLoadingContentInformationRequest) SetContentInformationRequest(value IAVAssetResourceLoadingContentInformationRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentInformationRequest:"), value)
}/* debug [instance_properties/setter]: contentInformationRequest */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetResourceLoadingContentInformationRequest */



