// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSSystemExtensionsWorkspace */

/* debug [class_header]: Header for OSSystemExtensionsWorkspace */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OSSystemExtensionsWorkspace */
// An interface definition for the [OSSystemExtensionsWorkspace] class.
type IOSSystemExtensionsWorkspace interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for OSSystemExtensionsWorkspace */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OSSystemExtensionsWorkspace */
	// methods:
	AddObserverError(observer unsafe.Pointer, error_ unsafe.Pointer) bool
	RemoveObserver(observer unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OSSystemExtensionsWorkspace */
// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionsWorkspaceClass) Alloc() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OSSystemExtensionsWorkspace */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace
type OSSystemExtensionsWorkspace struct {
	objectivec.Object
}

// OSSystemExtensionsWorkspaceFrom constructs a [OSSystemExtensionsWorkspace] from an unsafe.Pointer.
func OSSystemExtensionsWorkspaceFrom(ptr unsafe.Pointer) OSSystemExtensionsWorkspace {
	return OSSystemExtensionsWorkspace{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OSSystemExtensionsWorkspace */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OSSystemExtensionsWorkspace */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OSSystemExtensionsWorkspace */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/shared
func (oc _OSSystemExtensionsWorkspaceClass) SharedWorkspace() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("sharedWorkspace"))
	return rv
} /* debug [class_properties_class/property]: sharedWorkspace */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OSSystemExtensionsWorkspace */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/addObserver(_:)
func (o_ OSSystemExtensionsWorkspace) AddObserverError(observer unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addObserver:error:"), observer, error_)
	return rv
} /* debug [instance_methods/method]: AddObserverError */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/removeObserver(_:)
func (o_ OSSystemExtensionsWorkspace) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:"), observer)
} /* debug [instance_methods/method]: RemoveObserver */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OSSystemExtensionsWorkspace */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/shared
func (o_ OSSystemExtensionsWorkspace) SharedWorkspace() IOSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o_.ID, objc.Sel("sharedWorkspace"))
	return rv
} /* debug [instance_properties/getter]: sharedWorkspace */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class OSSystemExtensionsWorkspace */
