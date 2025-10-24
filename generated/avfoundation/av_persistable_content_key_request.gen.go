// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPersistableContentKeyRequest */


/* debug [class_header]: Header for AVPersistableContentKeyRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PersistableContentKeyRequest */
// An interface definition for the [PersistableContentKeyRequest] class.
type IPersistableContentKeyRequest interface {
	IContentKeyRequest
	
/* debug [class_interface_properties]: Properties for PersistableContentKeyRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PersistableContentKeyRequest */
	// methods:
	PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, outError objectivec.IObject) foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PersistableContentKeyRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PersistableContentKeyRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PersistableContentKeyRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PersistableContentKeyRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PersistableContentKeyRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PersistableContentKeyRequest */

// Creates a persistable content key from the content key context data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest/persistableContentKey(fromKeyVendorResponse:options:)
func (p_ PersistableContentKeyRequest) PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, outError objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("persistableContentKeyFromKeyVendorResponse:options:error:"), keyVendorResponse, options, outError)
	return rv
}/* debug [instance_methods/method]: PersistableContentKeyFromKeyVendorResponseOptionsError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PersistableContentKeyRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPersistableContentKeyRequest */



