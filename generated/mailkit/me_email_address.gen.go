// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEEmailAddress */


/* debug [class_header]: Header for MEEmailAddress */
// The class instance for the [MEEmailAddress] class.
var (
	MEEmailAddressClass     _MEEmailAddressClass
	MEEmailAddressClassOnce sync.Once
)

func getMEEmailAddressClass() _MEEmailAddressClass {
	MEEmailAddressClassOnce.Do(func() {
		MEEmailAddressClass = _MEEmailAddressClass{objc.GetClass("MEEmailAddress")}
	})
	return MEEmailAddressClass
}

type _MEEmailAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEEmailAddress */
// An interface definition for the [MEEmailAddress] class.
type IMEEmailAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEEmailAddress */
	// properties:
	AddressString() objc.IObject /* cross-framework: NSString */
	RawString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEEmailAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEEmailAddress */
// Alloc allocates a new instance without initialization.
func (mc _MEEmailAddressClass) Alloc() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEEmailAddressClass) New() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEEmailAddress) Init() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEEmailAddress) Autorelease() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEmailAddress creates a new MEEmailAddress instance.
func NewMEEmailAddress() MEEmailAddress {
	return getMEEmailAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEEmailAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress
type MEEmailAddress struct {
	objectivec.Object
}

// MEEmailAddressFrom constructs a [MEEmailAddress] from an unsafe.Pointer.
func MEEmailAddressFrom(ptr unsafe.Pointer) MEEmailAddress {
	return MEEmailAddress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEEmailAddress */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/init(rawString:)
func NewMEEmailAddressWithRawString(rawString objc.IObject /* cross-framework: NSString */) MEEmailAddress {
	instance := getMEEmailAddressClass().Alloc()
	rv := objc.Send[MEEmailAddress](instance.ID, objc.Sel("initWithRawString:"), rawString)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEEmailAddressWithRawString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEEmailAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEEmailAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEEmailAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEEmailAddress */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/addressString
func (m_ MEEmailAddress) AddressString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("addressString"))
	return rv
}/* debug [instance_properties/getter]: addressString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/rawString
func (m_ MEEmailAddress) RawString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("rawString"))
	return rv
}/* debug [instance_properties/getter]: rawString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEEmailAddress */


