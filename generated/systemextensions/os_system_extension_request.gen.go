// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSSystemExtensionRequest] class.
var (
	OSSystemExtensionRequestClass     _OSSystemExtensionRequestClass
	OSSystemExtensionRequestClassOnce sync.Once
)

func getOSSystemExtensionRequestClass() _OSSystemExtensionRequestClass {
	OSSystemExtensionRequestClassOnce.Do(func() {
		OSSystemExtensionRequestClass = _OSSystemExtensionRequestClass{objc.GetClass("OSSystemExtensionRequest")}
	})
	return OSSystemExtensionRequestClass
}

type _OSSystemExtensionRequestClass struct {
	class objc.Class
}

// An interface definition for the [OSSystemExtensionRequest] class.
type IOSSystemExtensionRequest interface {
	objectivec.IObject
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Identifier() string
}

// A request to activate or deactivate a system extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest
type OSSystemExtensionRequest struct {
	objectivec.Object
}

// OSSystemExtensionRequestFrom constructs a [OSSystemExtensionRequest] from an unsafe.Pointer.
//
// A request to activate or deactivate a system extension.
func OSSystemExtensionRequestFrom(ptr unsafe.Pointer) OSSystemExtensionRequest {
	return OSSystemExtensionRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionRequestClass) Alloc() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSSystemExtensionRequestClass) New() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionRequest) Init() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionRequest) Autorelease() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionRequest creates a new OSSystemExtensionRequest instance.
func NewOSSystemExtensionRequest() OSSystemExtensionRequest {
	return getOSSystemExtensionRequestClass().New()
}


// Creates a request to activate a System Extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/activationRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) ActivationRequestForExtensionQueue(identifier string, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("activationRequestForExtension:queue:"), objc.String(identifier), queue)
	return rv
}

// Creates a request to deactivate a System Extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/deactivationRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) DeactivationRequestForExtensionQueue(identifier string, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("deactivationRequestForExtension:queue:"), objc.String(identifier), queue)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/propertiesRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) PropertiesRequestForExtensionQueue(identifier string, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("propertiesRequestForExtension:queue:"), objc.String(identifier), queue)
	return rv
}

// A delegate to receive updates about the progress of a request.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/delegate
func (o_ OSSystemExtensionRequest) Delegate() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate to receive updates about the progress of a request.

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/delegate
func (o_ OSSystemExtensionRequest) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}

// The bundle identifier of the target extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/identifier
func (o_ OSSystemExtensionRequest) Identifier() string {
	rv := objc.Send[string](o_.ID, objc.Sel("identifier"))
	return rv
}



