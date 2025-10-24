// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXCallDirectoryExtensionContext */


/* debug [class_header]: Header for CXCallDirectoryExtensionContext */
// The class instance for the [CXCallDirectoryExtensionContext] class.
var (
	CXCallDirectoryExtensionContextClass     _CXCallDirectoryExtensionContextClass
	CXCallDirectoryExtensionContextClassOnce sync.Once
)

func getCXCallDirectoryExtensionContextClass() _CXCallDirectoryExtensionContextClass {
	CXCallDirectoryExtensionContextClassOnce.Do(func() {
		CXCallDirectoryExtensionContextClass = _CXCallDirectoryExtensionContextClass{objc.GetClass("CXCallDirectoryExtensionContext")}
	})
	return CXCallDirectoryExtensionContextClass
}

type _CXCallDirectoryExtensionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallDirectoryExtensionContext */
// An interface definition for the [CXCallDirectoryExtensionContext] class.
type ICXCallDirectoryExtensionContext interface {
	foundation.IExtensionContext
	
/* debug [class_interface_properties]: Properties for CXCallDirectoryExtensionContext */
	// properties:
	IsIncremental() bool
	SetIsIncremental(value bool)
	CXCallDirectoryPhoneNumberMax() CXCallDirectoryPhoneNumber /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallDirectoryExtensionContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallDirectoryExtensionContext */
// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryExtensionContextClass) Alloc() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallDirectoryExtensionContextClass) New() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryExtensionContext) Init() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryExtensionContext) Autorelease() CXCallDirectoryExtensionContext {
	rv := objc.Send[CXCallDirectoryExtensionContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryExtensionContext creates a new CXCallDirectoryExtensionContext instance.
func NewCXCallDirectoryExtensionContext() CXCallDirectoryExtensionContext {
	return getCXCallDirectoryExtensionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallDirectoryExtensionContext */
// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
//
// The system doesn’t initialize objects directly, but instead passes them as arguments to the instance method .


// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryExtensionContext
type CXCallDirectoryExtensionContext struct {
	foundation.ExtensionContext
}

// CXCallDirectoryExtensionContextFrom constructs a [CXCallDirectoryExtensionContext] from an unsafe.Pointer.
//
// A programmatic interface for adding identification and blocking entries to a Call Directory app extension.
func CXCallDirectoryExtensionContextFrom(ptr unsafe.Pointer) CXCallDirectoryExtensionContext {
	return CXCallDirectoryExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallDirectoryExtensionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallDirectoryExtensionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallDirectoryExtensionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallDirectoryExtensionContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallDirectoryExtensionContext */

// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) IsIncremental() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isIncremental"))
	return rv
}/* debug [instance_properties/getter]: isIncremental */


// A Boolean value that indicates whether the request provides data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryextensioncontext/isincremental
func (c_ CXCallDirectoryExtensionContext) SetIsIncremental(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsIncremental:"), value)
}/* debug [instance_properties/setter]: isIncremental */


// The maximum allowable value for a phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcalldirectoryphonenumbermax
func (c_ CXCallDirectoryExtensionContext) CXCallDirectoryPhoneNumberMax() CXCallDirectoryPhoneNumber /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("CXCallDirectoryPhoneNumberMax"))
	return rv
}/* debug [instance_properties/getter]: CXCallDirectoryPhoneNumberMax */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallDirectoryExtensionContext */


