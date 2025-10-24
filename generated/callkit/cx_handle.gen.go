// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXHandle */


/* debug [class_header]: Header for CXHandle */
// The class instance for the [CXHandle] class.
var (
	CXHandleClass     _CXHandleClass
	CXHandleClassOnce sync.Once
)

func getCXHandleClass() _CXHandleClass {
	CXHandleClassOnce.Do(func() {
		CXHandleClass = _CXHandleClass{objc.GetClass("CXHandle")}
	})
	return CXHandleClass
}

type _CXHandleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXHandle */
// An interface definition for the [CXHandle] class.
type ICXHandle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXHandle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXHandle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXHandle */
// Alloc allocates a new instance without initialization.
func (cc _CXHandleClass) Alloc() CXHandle {
	rv := objc.Send[CXHandle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXHandleClass) New() CXHandle {
	rv := objc.Send[CXHandle](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXHandle) Init() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXHandle) Autorelease() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXHandle creates a new CXHandle instance.
func NewCXHandle() CXHandle {
	return getCXHandleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXHandle */
// A way to reach a call recipient, such as a phone number or email address.
//
// When the telephony provider receives an incoming call or the user starts an outgoing call, the other caller is identified by a object. For a caller identified by a phone number, the handle type is and the value is a sequence of digits. For a caller identified by an email address, the handle type is and the value is an email address. For a caller identified in any other way, the handle type is and the value typically follows some domain-specific format, such as a username, numeric ID, or URL.


// A way to reach a call recipient, such as a phone number or email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle
type CXHandle struct {
	objectivec.Object
}

// CXHandleFrom constructs a [CXHandle] from an unsafe.Pointer.
//
// A way to reach a call recipient, such as a phone number or email address.
func CXHandleFrom(ptr unsafe.Pointer) CXHandle {
	return CXHandle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXHandle */

// Initializes a new handle of a given type with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/init(type:value:)
func NewCXHandleWithTypeValue(type_ CXHandleType, value objc.IObject /* cross-framework: NSString */) CXHandle {
	instance := getCXHandleClass().Alloc()
	rv := objc.Send[CXHandle](instance.ID, objc.Sel("initWithType:value:"), type_, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXHandleWithTypeValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXHandle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXHandle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXHandle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXHandle */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXHandle */


