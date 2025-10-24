// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTTelephonyNetworkInfo */


/* debug [class_header]: Header for CTTelephonyNetworkInfo */
// The class instance for the [TelephonyNetworkInfo] class.
var (
	TelephonyNetworkInfoClass     _TelephonyNetworkInfoClass
	TelephonyNetworkInfoClassOnce sync.Once
)

func getTelephonyNetworkInfoClass() _TelephonyNetworkInfoClass {
	TelephonyNetworkInfoClassOnce.Do(func() {
		TelephonyNetworkInfoClass = _TelephonyNetworkInfoClass{objc.GetClass("CTTelephonyNetworkInfo")}
	})
	return TelephonyNetworkInfoClass
}

type _TelephonyNetworkInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TelephonyNetworkInfo */
// An interface definition for the [TelephonyNetworkInfo] class.
type ITelephonyNetworkInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TelephonyNetworkInfo */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TelephonyNetworkInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TelephonyNetworkInfo */
// Alloc allocates a new instance without initialization.
func (tc _TelephonyNetworkInfoClass) Alloc() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TelephonyNetworkInfoClass) New() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TelephonyNetworkInfo) Init() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TelephonyNetworkInfo) Autorelease() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTelephonyNetworkInfo creates a new TelephonyNetworkInfo instance.
func NewTelephonyNetworkInfo() TelephonyNetworkInfo {
	return getTelephonyNetworkInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TelephonyNetworkInfo */
// An object that provides notifications of changes to the user’s cellular service provider.
//
// Your app should be able to handle changes to the user’s cellular service provider. For example, the user could swap the device’s SIM card with one from another provider while your app is running. This class also gives you access to the object, which contains information about the user’s home cellular service provider.


// An object that provides notifications of changes to the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo
type TelephonyNetworkInfo struct {
	objectivec.Object
}

// TelephonyNetworkInfoFrom constructs a [TelephonyNetworkInfo] from an unsafe.Pointer.
//
// An object that provides notifications of changes to the user’s cellular service provider.
func TelephonyNetworkInfoFrom(ptr unsafe.Pointer) TelephonyNetworkInfo {
	return TelephonyNetworkInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TelephonyNetworkInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TelephonyNetworkInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TelephonyNetworkInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TelephonyNetworkInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TelephonyNetworkInfo */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTTelephonyNetworkInfo */


