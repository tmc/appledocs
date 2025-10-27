// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [PersistableContentKeyRequest] class.
var (
	PersistableContentKeyRequestClass     _PersistableContentKeyRequestClass
	PersistableContentKeyRequestClassOnce sync.Once
)

func getPersistableContentKeyRequestClass() _PersistableContentKeyRequestClass {
	PersistableContentKeyRequestClassOnce.Do(func() {
		PersistableContentKeyRequestClass = _PersistableContentKeyRequestClass{objc.GetClass("AVPersistableContentKeyRequest")}
	})
	return PersistableContentKeyRequestClass
}

type _PersistableContentKeyRequestClass struct {
	class objc.Class
}





// An interface definition for the [PersistableContentKeyRequest] class.
type IPersistableContentKeyRequest interface {
	IContentKeyRequest
	

	// properties:


	

	// methods:
	PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse foundation.foundation.INSData, options foundation.IDictionary, outError foundation.foundation.INSError) foundation.Data


}





// Alloc allocates a new instance without initialization.
func (pc _PersistableContentKeyRequestClass) Alloc() PersistableContentKeyRequest {
	rv := objc.Send[PersistableContentKeyRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PersistableContentKeyRequestClass) New() PersistableContentKeyRequest {
	rv := objc.Send[PersistableContentKeyRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistableContentKeyRequest) Init() PersistableContentKeyRequest {
	rv := objc.Send[PersistableContentKeyRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistableContentKeyRequest) Autorelease() PersistableContentKeyRequest {
	rv := objc.Send[PersistableContentKeyRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistableContentKeyRequest creates a new PersistableContentKeyRequest instance.
func NewPersistableContentKeyRequest() PersistableContentKeyRequest {
	return getPersistableContentKeyRequestClass().New()
}





// An object that encapsulates information about a persistable content decryption key request issued from a content key session.
//
// This class allows clients to create and use persistable content keys.


// An object that encapsulates information about a persistable content decryption key request issued from a content key session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest
type PersistableContentKeyRequest struct {
	ContentKeyRequest
}

// PersistableContentKeyRequestFrom constructs a [PersistableContentKeyRequest] from an unsafe.Pointer.
//
// An object that encapsulates information about a persistable content decryption key request issued from a content key session.
func PersistableContentKeyRequestFrom(ptr unsafe.Pointer) PersistableContentKeyRequest {
	return PersistableContentKeyRequest{
		ContentKeyRequest: ContentKeyRequestFrom(ptr),
	}
}




















// Creates a persistable content key from the content key context data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest/persistableContentKey(fromKeyVendorResponse:options:)
func (p_ PersistableContentKeyRequest) PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse foundation.foundation.INSData, options foundation.IDictionary, outError foundation.foundation.INSError) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("persistableContentKeyFromKeyVendorResponse:options:error:"), keyVendorResponse, options, outError)
	return rv
}













