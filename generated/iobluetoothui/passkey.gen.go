// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class passkey */


/* debug [class_header]: Header for passkey */
// The class instance for the [passkey] class.
var (
	PasskeyClass     _passkeyClass
	PasskeyClassOnce sync.Once
)

func getpasskeyClass() _passkeyClass {
	PasskeyClassOnce.Do(func() {
		PasskeyClass = _passkeyClass{objc.GetClass("passkey")}
	})
	return PasskeyClass
}

type _passkeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for passkey */
// An interface definition for the [passkey] class.
type Ipasskey interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for passkey */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for passkey */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for passkey */
// Alloc allocates a new instance without initialization.
func (pc _passkeyClass) Alloc() passkey {
	rv := objc.Send[passkey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _passkeyClass) New() passkey {
	rv := objc.Send[passkey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ passkey) Init() passkey {
	rv := objc.Send[passkey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ passkey) Autorelease() passkey {
	rv := objc.Send[passkey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// Newpasskey creates a new passkey instance.
func Newpasskey() passkey {
	return getpasskeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for passkey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/passkey-c.ivar
type passkey struct {
	objectivec.Object
}

// passkeyFrom constructs a [passkey] from an unsafe.Pointer.
func passkeyFrom(ptr unsafe.Pointer) passkey {
	return passkey{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for passkey *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for passkey */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for passkey */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for passkey */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for passkey */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class passkey */



