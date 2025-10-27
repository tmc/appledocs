// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ContentKeyRequest] class.
var (
	ContentKeyRequestClass     _ContentKeyRequestClass
	ContentKeyRequestClassOnce sync.Once
)

func getContentKeyRequestClass() _ContentKeyRequestClass {
	ContentKeyRequestClassOnce.Do(func() {
		ContentKeyRequestClass = _ContentKeyRequestClass{objc.GetClass("AVContentKeyRequest")}
	})
	return ContentKeyRequestClass
}

type _ContentKeyRequestClass struct {
	class objc.Class
}





// An interface definition for the [ContentKeyRequest] class.
type IContentKeyRequest interface {
	objectivec.IObject
	

	// properties:
	CanProvidePersistableContentKey() bool
	ContentKey() IAVContentKey
	ContentKeySpecifier() IAVContentKeySpecifier
	Error() foundation.foundation.INSError
	Identifier() objc.ID
	InitializationData() foundation.foundation.INSData
	Options() foundation.IDictionary
	OriginatingRecipient() unsafe.Pointer
	RenewsExpiringResponseData() bool
	Status() ContentKeyRequestStatus
	AVContentKeyRequestProtocolVersionsKey() foundation.foundation.INSString
	AVContentKeyRequestRandomDeviceIdentifierSeedKey() foundation.foundation.INSString
	AVContentKeyRequestRequiresValidationDataInSecureTokenKey() foundation.foundation.INSString
	AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() foundation.foundation.INSString


	

	// methods:
	MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier foundation.foundation.INSData, contentIdentifier foundation.foundation.INSData, options foundation.IDictionary, handler unsafe.Pointer)
	ProcessContentKeyResponse(keyResponse IAVContentKeyResponse)
	ProcessContentKeyResponseError(error_ foundation.foundation.INSError)
	RespondByRequestingPersistableContentKeyRequestAndReturnError(outError foundation.foundation.INSError) bool


}





// Alloc allocates a new instance without initialization.
func (cc _ContentKeyRequestClass) Alloc() ContentKeyRequest {
	rv := objc.Send[ContentKeyRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentKeyRequestClass) New() ContentKeyRequest {
	rv := objc.Send[ContentKeyRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentKeyRequest) Init() ContentKeyRequest {
	rv := objc.Send[ContentKeyRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentKeyRequest) Autorelease() ContentKeyRequest {
	rv := objc.Send[ContentKeyRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentKeyRequest creates a new ContentKeyRequest instance.
func NewContentKeyRequest() ContentKeyRequest {
	return getContentKeyRequestClass().New()
}





// An object that encapsulates information about a content decryption key request issued from a content key session object.


// An object that encapsulates information about a content decryption key request issued from a content key session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest
type ContentKeyRequest struct {
	objectivec.Object
}

// ContentKeyRequestFrom constructs a [ContentKeyRequest] from an unsafe.Pointer.
//
// An object that encapsulates information about a content decryption key request issued from a content key session object.
func ContentKeyRequestFrom(ptr unsafe.Pointer) ContentKeyRequest {
	return ContentKeyRequest{objectivec.Object{objc.ID(ptr)}}
}




















// Obtains encrypted key request data for a specific combination of app and content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/makeStreamingContentKeyRequestData(forApp:contentIdentifier:options:completionHandler:)
func (c_ ContentKeyRequest) MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier foundation.foundation.INSData, contentIdentifier foundation.foundation.INSData, options foundation.IDictionary, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("makeStreamingContentKeyRequestDataForApp:contentIdentifier:options:completionHandler:"), appIdentifier, contentIdentifier, options, handler)
}


// Sends the specified content key response to the receiver for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponse(_:)
func (c_ ContentKeyRequest) ProcessContentKeyResponse(keyResponse IAVContentKeyResponse) {
	objc.Send[objc.ID](c_.ID, objc.Sel("processContentKeyResponse:"), keyResponse)
}


// Tells the receiver that the app was unable to obtain a content key response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponseError(_:)
func (c_ ContentKeyRequest) ProcessContentKeyResponseError(error_ foundation.foundation.INSError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("processContentKeyResponseError:"), error_)
}


// Tells the receiver that the app requires a persistable content key request object for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/respondByRequestingPersistableContentKeyRequest()-7i2pw
func (c_ ContentKeyRequest) RespondByRequestingPersistableContentKeyRequestAndReturnError(outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("respondByRequestingPersistableContentKeyRequestAndReturnError:"), outError)
	return rv
}







// The content key request used to create a persistable content key or respond to a previous request with a persistable content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/canProvidePersistableContentKey
func (c_ ContentKeyRequest) CanProvidePersistableContentKey() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canProvidePersistableContentKey"))
	return rv
}


// The generated content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKey
func (c_ ContentKeyRequest) ContentKey() IAVContentKey {
	rv := objc.Send[ContentKey](c_.ID, objc.Sel("contentKey"))
	return rv
}


// The requested content key specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKeySpecifier
func (c_ ContentKeyRequest) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](c_.ID, objc.Sel("contentKeySpecifier"))
	return rv
}


// The error description for a failed key request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/error
func (c_ ContentKeyRequest) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](c_.ID, objc.Sel("error"))
	return rv
}


// The identifier for the content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/identifier
func (c_ ContentKeyRequest) Identifier() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("identifier"))
	return rv
}


// The data used to obtain a key response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/initializationData
func (c_ ContentKeyRequest) InitializationData() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("initializationData"))
	return rv
}


// A dictionary of options used to initialize key loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/options
func (c_ ContentKeyRequest) Options() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("options"))
	return rv
}


// The AVContentKeyRecipient which initiated this request, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/originatingRecipient
func (c_ ContentKeyRequest) OriginatingRecipient() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("originatingRecipient"))
	return rv
}


// A Boolean value that indicates whether the content key request renews previously provided response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/renewsExpiringResponseData
func (c_ ContentKeyRequest) RenewsExpiringResponseData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("renewsExpiringResponseData"))
	return rv
}


// The current state of the content key request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/status-swift.property
func (c_ ContentKeyRequest) Status() ContentKeyRequestStatus {
	rv := objc.Send[ContentKeyRequestStatus](c_.ID, objc.Sel("status"))
	return rv
}


// A key that specifies the versions of the content protection protocol supported by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestprotocolversionskey
func (c_ ContentKeyRequest) AVContentKeyRequestProtocolVersionsKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestProtocolVersionsKey"))
	return rv
}


// Value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrandomdeviceidentifierseedkey
func (c_ ContentKeyRequest) AVContentKeyRequestRandomDeviceIdentifierSeedKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestRandomDeviceIdentifierSeedKey"))
	return rv
}


// A key that requires the secure token to have extended validation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrequiresvalidationdatainsecuretokenkey
func (c_ ContentKeyRequest) AVContentKeyRequestRequiresValidationDataInSecureTokenKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestRequiresValidationDataInSecureTokenKey"))
	return rv
}


// Value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestshouldrandomizedeviceidentifierkey
func (c_ ContentKeyRequest) AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestShouldRandomizeDeviceIdentifierKey"))
	return rv
}







