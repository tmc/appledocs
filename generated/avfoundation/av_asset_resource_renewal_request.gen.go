// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AssetResourceRenewalRequest] class.
var (
	AssetResourceRenewalRequestClass     _AssetResourceRenewalRequestClass
	AssetResourceRenewalRequestClassOnce sync.Once
)

func getAssetResourceRenewalRequestClass() _AssetResourceRenewalRequestClass {
	AssetResourceRenewalRequestClassOnce.Do(func() {
		AssetResourceRenewalRequestClass = _AssetResourceRenewalRequestClass{objc.GetClass("AVAssetResourceRenewalRequest")}
	})
	return AssetResourceRenewalRequestClass
}

type _AssetResourceRenewalRequestClass struct {
	class objc.Class
}





// An interface definition for the [AssetResourceRenewalRequest] class.
type IAssetResourceRenewalRequest interface {
	IAssetResourceLoadingRequest
	

	// properties:
	RenewalDate() foundation.Date
	SetRenewalDate(value foundation.Date)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetResourceRenewalRequestClass) Alloc() AssetResourceRenewalRequest {
	rv := objc.Send[AssetResourceRenewalRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetResourceRenewalRequestClass) New() AssetResourceRenewalRequest {
	rv := objc.Send[AssetResourceRenewalRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceRenewalRequest) Init() AssetResourceRenewalRequest {
	rv := objc.Send[AssetResourceRenewalRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceRenewalRequest) Autorelease() AssetResourceRenewalRequest {
	rv := objc.Send[AssetResourceRenewalRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceRenewalRequest creates a new AssetResourceRenewalRequest instance.
func NewAssetResourceRenewalRequest() AssetResourceRenewalRequest {
	return getAssetResourceRenewalRequestClass().New()
}





// An object that encapsulates information about a resource request from a resource loader to renew a previously issued request.
//
// When an needs to renew a resource, because the has been set on a previous loading request, it asks its object to assist. The resource loader encapsulates the request information by creating an instance of this object, which it then hands to its delegate for processing. The delegate uses the information in this object to perform the request and report on the success or failure of the operation. The class is a subclass of .


// An object that encapsulates information about a resource request from a resource loader to renew a previously issued request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceRenewalRequest
type AssetResourceRenewalRequest struct {
	AssetResourceLoadingRequest
}

// AssetResourceRenewalRequestFrom constructs a [AssetResourceRenewalRequest] from an unsafe.Pointer.
//
// An object that encapsulates information about a resource request from a resource loader to renew a previously issued request.
func AssetResourceRenewalRequestFrom(ptr unsafe.Pointer) AssetResourceRenewalRequest {
	return AssetResourceRenewalRequest{
		AssetResourceLoadingRequest: AssetResourceLoadingRequestFrom(ptr),
	}
}

























// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/renewaldate
func (a_ AssetResourceRenewalRequest) RenewalDate() foundation.Date {
	rv := objc.Send[foundation.Date](a_.ID, objc.Sel("renewalDate"))
	return rv
}


// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/renewaldate
func (a_ AssetResourceRenewalRequest) SetRenewalDate(value foundation.Date) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRenewalDate:"), value)
}








