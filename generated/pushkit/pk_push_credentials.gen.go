// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKPushCredentials */


/* debug [class_header]: Header for PKPushCredentials */
// The class instance for the [PushCredentials] class.
var (
	PushCredentialsClass     _PushCredentialsClass
	PushCredentialsClassOnce sync.Once
)

func getPushCredentialsClass() _PushCredentialsClass {
	PushCredentialsClassOnce.Do(func() {
		PushCredentialsClass = _PushCredentialsClass{objc.GetClass("PKPushCredentials")}
	})
	return PushCredentialsClass
}

type _PushCredentialsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PushCredentials */
// An interface definition for the [PushCredentials] class.
type IPushCredentials interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PushCredentials */
	// properties:
	Token() objc.IObject /* cross-framework: NSData */
	Type() PushType /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PushCredentials */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PushCredentials */
// Alloc allocates a new instance without initialization.
func (pc _PushCredentialsClass) Alloc() PushCredentials {
	rv := objc.Send[PushCredentials](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PushCredentialsClass) New() PushCredentials {
	rv := objc.Send[PushCredentials](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PushCredentials) Init() PushCredentials {
	rv := objc.Send[PushCredentials](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PushCredentials) Autorelease() PushCredentials {
	rv := objc.Send[PushCredentials](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPushCredentials creates a new PushCredentials instance.
func NewPushCredentials() PushCredentials {
	return getPushCredentialsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PushCredentials */
// An object that encapsulates the device token you use to deliver push notifications to your app.
//
// When registering your app’s push types, PushKit creates a object for each type your app supports and delivers it to your delegate’s method. Don’t create objects yourself.


// An object that encapsulates the device token you use to deliver push notifications to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushCredentials
type PushCredentials struct {
	objectivec.Object
}

// PushCredentialsFrom constructs a [PushCredentials] from an unsafe.Pointer.
//
// An object that encapsulates the device token you use to deliver push notifications to your app.
func PushCredentialsFrom(ptr unsafe.Pointer) PushCredentials {
	return PushCredentials{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PushCredentials *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PushCredentials */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PushCredentials */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PushCredentials */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PushCredentials */

// A unique device token to use when sending push notifications to the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushCredentials/token
func (p_ PushCredentials) Token() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("token"))
	return rv
}/* debug [instance_properties/getter]: token */


// The push type constant associated with the token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushCredentials/type
func (p_ PushCredentials) Type() PushType /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKPushCredentials */



