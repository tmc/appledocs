// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




