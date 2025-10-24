// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCallDirectoryProvider */


/* debug [class_header]: Header for CXCallDirectoryProvider */
// The class instance for the [CXCallDirectoryProvider] class.
var (
	CXCallDirectoryProviderClass     _CXCallDirectoryProviderClass
	CXCallDirectoryProviderClassOnce sync.Once
)

func getCXCallDirectoryProviderClass() _CXCallDirectoryProviderClass {
	CXCallDirectoryProviderClassOnce.Do(func() {
		CXCallDirectoryProviderClass = _CXCallDirectoryProviderClass{objc.GetClass("CXCallDirectoryProvider")}
	})
	return CXCallDirectoryProviderClass
}

type _CXCallDirectoryProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallDirectoryProvider */
// An interface definition for the [CXCallDirectoryProvider] class.
type ICXCallDirectoryProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCallDirectoryProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallDirectoryProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallDirectoryProvider */
// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryProviderClass) Alloc() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallDirectoryProviderClass) New() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryProvider) Init() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryProvider) Autorelease() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryProvider creates a new CXCallDirectoryProvider instance.
func NewCXCallDirectoryProvider() CXCallDirectoryProvider {
	return getCXCallDirectoryProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallDirectoryProvider */
// The principal object for a Call Directory app extension for a host app.


// The principal object for a Call Directory app extension for a host app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryProvider
type CXCallDirectoryProvider struct {
	objectivec.Object
}

// CXCallDirectoryProviderFrom constructs a [CXCallDirectoryProvider] from an unsafe.Pointer.
//
// The principal object for a Call Directory app extension for a host app.
func CXCallDirectoryProviderFrom(ptr unsafe.Pointer) CXCallDirectoryProvider {
	return CXCallDirectoryProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallDirectoryProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallDirectoryProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallDirectoryProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallDirectoryProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallDirectoryProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallDirectoryProvider */


