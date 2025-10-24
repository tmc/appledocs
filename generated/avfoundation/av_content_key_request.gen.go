// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVContentKeyRequest */


/* debug [class_header]: Header for AVContentKeyRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentKeyRequest */
// An interface definition for the [ContentKeyRequest] class.
type IContentKeyRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentKeyRequest */
	// properties:
	CanProvidePersistableContentKey() bool
	ContentKey() IAVContentKey
	ContentKeySpecifier() IAVContentKeySpecifier
	Error() Error
	Identifier() objc.ID
	InitializationData() objc.IObject /* cross-framework: NSData */
	Options() foundation.IDictionary
	OriginatingRecipient() unsafe.Pointer
	RenewsExpiringResponseData() bool
	Status() ContentKeyRequestStatus
	AVContentKeyRequestProtocolVersionsKey() objc.IObject /* cross-framework: NSString */
	AVContentKeyRequestRandomDeviceIdentifierSeedKey() objc.IObject /* cross-framework: NSString */
	AVContentKeyRequestRequiresValidationDataInSecureTokenKey() objc.IObject /* cross-framework: NSString */
	AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentKeyRequest */
	// methods:
	MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier objc.IObject /* cross-framework: NSData */, contentIdentifier objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer)
	ProcessContentKeyResponse(keyResponse IAVContentKeyResponse)
	ProcessContentKeyResponseError(error_ Error)
	RespondByRequestingPersistableContentKeyRequestAndReturnError(outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentKeyRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentKeyRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentKeyRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentKeyRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentKeyRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentKeyRequest */

// Obtains encrypted key request data for a specific combination of app and content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/makeStreamingContentKeyRequestData(forApp:contentIdentifier:options:completionHandler:)
func (c_ ContentKeyRequest) MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier objc.IObject /* cross-framework: NSData */, contentIdentifier objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("makeStreamingContentKeyRequestDataForApp:contentIdentifier:options:completionHandler:"), appIdentifier, contentIdentifier, options, handler)
}/* debug [instance_methods/method]: MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler */


// Sends the specified content key response to the receiver for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponse(_:)
func (c_ ContentKeyRequest) ProcessContentKeyResponse(keyResponse IAVContentKeyResponse) {
	objc.Send[objc.ID](c_.ID, objc.Sel("processContentKeyResponse:"), keyResponse)
}/* debug [instance_methods/method]: ProcessContentKeyResponse */


// Tells the receiver that the app was unable to obtain a content key response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponseError(_:)
func (c_ ContentKeyRequest) ProcessContentKeyResponseError(error_ Error) {
	objc.Send[objc.ID](c_.ID, objc.Sel("processContentKeyResponseError:"), error_)
}/* debug [instance_methods/method]: ProcessContentKeyResponseError */


// Tells the receiver that the app requires a persistable content key request object for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/respondByRequestingPersistableContentKeyRequest()-7i2pw
func (c_ ContentKeyRequest) RespondByRequestingPersistableContentKeyRequestAndReturnError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("respondByRequestingPersistableContentKeyRequestAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: RespondByRequestingPersistableContentKeyRequestAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentKeyRequest */

// The content key request used to create a persistable content key or respond to a previous request with a persistable content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/canProvidePersistableContentKey
func (c_ ContentKeyRequest) CanProvidePersistableContentKey() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canProvidePersistableContentKey"))
	return rv
}/* debug [instance_properties/getter]: canProvidePersistableContentKey */


// The generated content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKey
func (c_ ContentKeyRequest) ContentKey() IAVContentKey {
	rv := objc.Send[ContentKey](c_.ID, objc.Sel("contentKey"))
	return rv
}/* debug [instance_properties/getter]: contentKey */


// The requested content key specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKeySpecifier
func (c_ ContentKeyRequest) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](c_.ID, objc.Sel("contentKeySpecifier"))
	return rv
}/* debug [instance_properties/getter]: contentKeySpecifier */


// The error description for a failed key request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/error
func (c_ ContentKeyRequest) Error() Error {
	rv := objc.Send[Error](c_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The identifier for the content key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/identifier
func (c_ ContentKeyRequest) Identifier() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The data used to obtain a key response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/initializationData
func (c_ ContentKeyRequest) InitializationData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("initializationData"))
	return rv
}/* debug [instance_properties/getter]: initializationData */


// A dictionary of options used to initialize key loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/options
func (c_ ContentKeyRequest) Options() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The AVContentKeyRecipient which initiated this request, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/originatingRecipient
func (c_ ContentKeyRequest) OriginatingRecipient() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("originatingRecipient"))
	return rv
}/* debug [instance_properties/getter]: originatingRecipient */


// A Boolean value that indicates whether the content key request renews previously provided response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/renewsExpiringResponseData
func (c_ ContentKeyRequest) RenewsExpiringResponseData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("renewsExpiringResponseData"))
	return rv
}/* debug [instance_properties/getter]: renewsExpiringResponseData */


// The current state of the content key request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/status-swift.property
func (c_ ContentKeyRequest) Status() ContentKeyRequestStatus {
	rv := objc.Send[ContentKeyRequestStatus](c_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// A key that specifies the versions of the content protection protocol supported by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestprotocolversionskey
func (c_ ContentKeyRequest) AVContentKeyRequestProtocolVersionsKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestProtocolVersionsKey"))
	return rv
}/* debug [instance_properties/getter]: AVContentKeyRequestProtocolVersionsKey */


// Value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrandomdeviceidentifierseedkey
func (c_ ContentKeyRequest) AVContentKeyRequestRandomDeviceIdentifierSeedKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestRandomDeviceIdentifierSeedKey"))
	return rv
}/* debug [instance_properties/getter]: AVContentKeyRequestRandomDeviceIdentifierSeedKey */


// A key that requires the secure token to have extended validation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrequiresvalidationdatainsecuretokenkey
func (c_ ContentKeyRequest) AVContentKeyRequestRequiresValidationDataInSecureTokenKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestRequiresValidationDataInSecureTokenKey"))
	return rv
}/* debug [instance_properties/getter]: AVContentKeyRequestRequiresValidationDataInSecureTokenKey */


// Value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestshouldrandomizedeviceidentifierkey
func (c_ ContentKeyRequest) AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVContentKeyRequestShouldRandomizeDeviceIdentifierKey"))
	return rv
}/* debug [instance_properties/getter]: AVContentKeyRequestShouldRandomizeDeviceIdentifierKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContentKeyRequest */


