// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ContentKeyResponse] class.
type IContentKeyResponse interface {
	objectivec.IObject
}

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

// Alloc allocates a new instance without initialization.
func (cc _ContentKeyResponseClass) Alloc() ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a new key response object for key data and initialization vector sent in the clear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(clearKeyData:initializationVector:)
func NewContentKeyResponseWithClearKeyDataInitializationVector(keyData foundation.NSData, initializationVector foundation.NSData) ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(getContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return rv
}


// Creates a content key response with an encrypted key response data blob when FairPlay Streaming is the key delivery method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(fairPlayStreamingKeyResponseData:)
func NewContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData foundation.NSData) ContentKeyResponse {
	rv := objc.Send[ContentKeyResponse](objc.ID(getContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return rv
}



// Creates a new key response object for key data and initialization vector sent in the clear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(clearKeyData:initializationVector:)
func (cc _ContentKeyResponseClass) ContentKeyResponseWithClearKeyDataInitializationVector(keyData foundation.NSData, initializationVector foundation.NSData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return rv
}


// Creates a content key response with an encrypted key response data blob when FairPlay Streaming is the key delivery method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(fairPlayStreamingKeyResponseData:)
func (cc _ContentKeyResponseClass) ContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData foundation.NSData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return rv
}


