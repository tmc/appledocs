// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVContentKeyResponse */


/* debug [class_header]: Header for AVContentKeyResponse */
// The class instance for the [ContentKeyResponse] class.
var (
	ContentKeyResponseClass     _ContentKeyResponseClass
	ContentKeyResponseClassOnce sync.Once
)

func getContentKeyResponseClass() _ContentKeyResponseClass {
	ContentKeyResponseClassOnce.Do(func() {
		ContentKeyResponseClass = _ContentKeyResponseClass{objc.GetClass("AVContentKeyResponse")}
	})
	return ContentKeyResponseClass
}

type _ContentKeyResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentKeyResponse */
// An interface definition for the [ContentKeyResponse] class.
type IContentKeyResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentKeyResponse */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentKeyResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentKeyResponse */
// Alloc allocates a new instance without initialization.
func (cc _ContentKeyResponseClass) Alloc() ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentKeyResponseClass) New() ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentKeyResponse) Init() ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentKeyResponse) Autorelease() ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentKeyResponse creates a new ContentKeyResponse instance.
func NewContentKeyResponse() ContentKeyResponse {
	return getContentKeyResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentKeyResponse */
// An object that encapsulates information about a response to a content decryption key request.


// An object that encapsulates information about a response to a content decryption key request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse
type ContentKeyResponse struct {
	objectivec.Object
}

// ContentKeyResponseFrom constructs a [ContentKeyResponse] from an unsafe.Pointer.
//
// An object that encapsulates information about a response to a content decryption key request.
func ContentKeyResponseFrom(ptr unsafe.Pointer) ContentKeyResponse {
	return ContentKeyResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentKeyResponse */

// Creates a content key response with an authorization token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(authorizationTokenData:)
func NewContentKeyResponseWithAuthorizationTokenData(authorizationTokenData objc.IObject /* cross-framework: NSData */) ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(getContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithAuthorizationTokenData:"), authorizationTokenData)
	return rv
}/* debug [class_init_methods/constructor]: NewContentKeyResponseWithAuthorizationTokenData */


// Creates a new key response object for key data and initialization vector sent in the clear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(clearKeyData:initializationVector:)
func NewContentKeyResponseWithClearKeyDataInitializationVector(keyData objc.IObject /* cross-framework: NSData */, initializationVector objc.IObject /* cross-framework: NSData */) ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(getContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return rv
}/* debug [class_init_methods/constructor]: NewContentKeyResponseWithClearKeyDataInitializationVector */


// Creates a content key response with an encrypted key response data blob when FairPlay Streaming is the key delivery method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(fairPlayStreamingKeyResponseData:)
func NewContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData objc.IObject /* cross-framework: NSData */) ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(getContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return rv
}/* debug [class_init_methods/constructor]: NewContentKeyResponseWithFairPlayStreamingKeyResponseData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentKeyResponse */

// Creates a content key response with an authorization token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(authorizationTokenData:)
func (cc _ContentKeyResponseClass) ContentKeyResponseWithAuthorizationTokenData(authorizationTokenData objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeyResponseWithAuthorizationTokenData:"), authorizationTokenData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentKeyResponseWithAuthorizationTokenData) */


// Creates a new key response object for key data and initialization vector sent in the clear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(clearKeyData:initializationVector:)
func (cc _ContentKeyResponseClass) ContentKeyResponseWithClearKeyDataInitializationVector(keyData objc.IObject /* cross-framework: NSData */, initializationVector objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentKeyResponseWithClearKeyDataInitializationVector) */


// Creates a content key response with an encrypted key response data blob when FairPlay Streaming is the key delivery method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(fairPlayStreamingKeyResponseData:)
func (cc _ContentKeyResponseClass) ContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentKeyResponseWithFairPlayStreamingKeyResponseData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentKeyResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentKeyResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentKeyResponse */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContentKeyResponse */


