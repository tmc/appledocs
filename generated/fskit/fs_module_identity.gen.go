// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSModuleIdentity */


/* debug [class_header]: Header for FSModuleIdentity */
// The class instance for the [FSModuleIdentity] class.
var (
	FSModuleIdentityClass     _FSModuleIdentityClass
	FSModuleIdentityClassOnce sync.Once
)

func getFSModuleIdentityClass() _FSModuleIdentityClass {
	FSModuleIdentityClassOnce.Do(func() {
		FSModuleIdentityClass = _FSModuleIdentityClass{objc.GetClass("FSModuleIdentity")}
	})
	return FSModuleIdentityClass
}

type _FSModuleIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSModuleIdentity */
// An interface definition for the [FSModuleIdentity] class.
type IFSModuleIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSModuleIdentity */
	// properties:
	BundleIdentifier() objc.IObject /* cross-framework: NSString */
	Enabled() bool
	Url() objc.IObject /* cross-framework: NSURL */
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSModuleIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSModuleIdentity */
// Alloc allocates a new instance without initialization.
func (fc _FSModuleIdentityClass) Alloc() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSModuleIdentityClass) New() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSModuleIdentity) Init() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSModuleIdentity) Autorelease() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSModuleIdentity creates a new FSModuleIdentity instance.
func NewFSModuleIdentity() FSModuleIdentity {
	return getFSModuleIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSModuleIdentity */
// An installed file system module.


// An installed file system module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity
type FSModuleIdentity struct {
	objectivec.Object
}

// FSModuleIdentityFrom constructs a [FSModuleIdentity] from an unsafe.Pointer.
//
// An installed file system module.
func FSModuleIdentityFrom(ptr unsafe.Pointer) FSModuleIdentity {
	return FSModuleIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSModuleIdentity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSModuleIdentity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSModuleIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSModuleIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSModuleIdentity */

// The module’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/bundleIdentifier
func (f_ FSModuleIdentity) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("bundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: bundleIdentifier */


// A Boolean value that indicates if the module is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/isEnabled
func (f_ FSModuleIdentity) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// The module’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/url
func (f_ FSModuleIdentity) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A Boolean value that indicates if the module is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsmoduleidentity/isenabled
func (f_ FSModuleIdentity) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates if the module is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsmoduleidentity/isenabled
func (f_ FSModuleIdentity) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSModuleIdentity */



