// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSSystemExtensionsWorkspace] class.
var (
	OSSystemExtensionsWorkspaceClass     _OSSystemExtensionsWorkspaceClass
	OSSystemExtensionsWorkspaceClassOnce sync.Once
)

func getOSSystemExtensionsWorkspaceClass() _OSSystemExtensionsWorkspaceClass {
	OSSystemExtensionsWorkspaceClassOnce.Do(func() {
		OSSystemExtensionsWorkspaceClass = _OSSystemExtensionsWorkspaceClass{objc.GetClass("OSSystemExtensionsWorkspace")}
	})
	return OSSystemExtensionsWorkspaceClass
}

type _OSSystemExtensionsWorkspaceClass struct {
	class objc.Class
}

// An interface definition for the [OSSystemExtensionsWorkspace] class.
type IOSSystemExtensionsWorkspace interface {
	objectivec.IObject
	AddObserverError(observer objectivec.IObject, error_ unsafe.Pointer) bool
	RemoveObserver(observer objectivec.IObject)
	SystemExtensionsForApplicationWithBundleIDError(bundleID string, out_error unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace
type OSSystemExtensionsWorkspace struct {
	objectivec.Object
}

// OSSystemExtensionsWorkspaceFrom constructs a [OSSystemExtensionsWorkspace] from an unsafe.Pointer.
func OSSystemExtensionsWorkspaceFrom(ptr unsafe.Pointer) OSSystemExtensionsWorkspace {
	return OSSystemExtensionsWorkspace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionsWorkspaceClass) Alloc() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSSystemExtensionsWorkspaceClass) New() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionsWorkspace) Init() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionsWorkspace) Autorelease() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionsWorkspace creates a new OSSystemExtensionsWorkspace instance.
func NewOSSystemExtensionsWorkspace() OSSystemExtensionsWorkspace {
	return getOSSystemExtensionsWorkspaceClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/shared
func (oc _OSSystemExtensionsWorkspaceClass) SharedWorkspace() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("sharedWorkspace"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/addObserver(_:)
func (o_ OSSystemExtensionsWorkspace) AddObserverError(observer objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addObserver:error:"), observer, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/removeObserver(_:)
func (o_ OSSystemExtensionsWorkspace) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:"), observer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/systemExtensions(forApplicationWithBundleID:)
func (o_ OSSystemExtensionsWorkspace) SystemExtensionsForApplicationWithBundleIDError(bundleID string, out_error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("systemExtensionsForApplicationWithBundleID:error:"), objc.String(bundleID), out_error)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/shared
func (o_ OSSystemExtensionsWorkspace) SharedWorkspace() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o_.ID, objc.Sel("sharedWorkspace"))
	return rv
}




