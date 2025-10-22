// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSSystemExtensionManager] class.
var (
	OSSystemExtensionManagerClass     _OSSystemExtensionManagerClass
	OSSystemExtensionManagerClassOnce sync.Once
)

func getOSSystemExtensionManagerClass() _OSSystemExtensionManagerClass {
	OSSystemExtensionManagerClassOnce.Do(func() {
		OSSystemExtensionManagerClass = _OSSystemExtensionManagerClass{objc.GetClass("OSSystemExtensionManager")}
	})
	return OSSystemExtensionManagerClass
}

type _OSSystemExtensionManagerClass struct {
	class objc.Class
}

// An interface definition for the [OSSystemExtensionManager] class.
type IOSSystemExtensionManager interface {
	objectivec.IObject
	SubmitRequest(request IOSSystemExtensionRequest)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
}

// A type that facilitates activation and deactivation of system extensions.
//
// Create an instance of with the class methods on that type, and submit it to the shared instance of the extension manager with . Set the on the request to receive the result of the activation or deactivation. The delegate also receives notifications if the user needs to authorize the extension or if a version conflict occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager
type OSSystemExtensionManager struct {
	objectivec.Object
}

// OSSystemExtensionManagerFrom constructs a [OSSystemExtensionManager] from an unsafe.Pointer.
//
// A type that facilitates activation and deactivation of system extensions.
func OSSystemExtensionManagerFrom(ptr unsafe.Pointer) OSSystemExtensionManager {
	return OSSystemExtensionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionManagerClass) Alloc() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSSystemExtensionManagerClass) New() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionManager) Init() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionManager) Autorelease() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionManager creates a new OSSystemExtensionManager instance.
func NewOSSystemExtensionManager() OSSystemExtensionManager {
	return getOSSystemExtensionManagerClass().New()
}


// The shared instance of the extension manager.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/shared
func (oc _OSSystemExtensionManagerClass) SharedManager() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("sharedManager"))
	return rv
}
// Submits a system extension request to the manager.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/submitRequest(_:)
func (o_ OSSystemExtensionManager) SubmitRequest(request IOSSystemExtensionRequest) {
	objc.Send[objc.ID](o_.ID, objc.Sel("submitRequest:"), request)
}

// The shared instance of the extension manager.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/shared
func (o_ OSSystemExtensionManager) SharedManager() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("sharedManager"))
	return rv
}

// A delegate to receive updates about the progress of a request.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequest/delegate
func (o_ OSSystemExtensionManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate to receive updates about the progress of a request.

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequest/delegate
func (o_ OSSystemExtensionManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}



