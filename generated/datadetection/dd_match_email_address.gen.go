// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchEmailAddress */


/* debug [class_header]: Header for DDMatchEmailAddress */
// The class instance for the [DDMatchEmailAddress] class.
var (
	DDMatchEmailAddressClass     _DDMatchEmailAddressClass
	DDMatchEmailAddressClassOnce sync.Once
)

func getDDMatchEmailAddressClass() _DDMatchEmailAddressClass {
	DDMatchEmailAddressClassOnce.Do(func() {
		DDMatchEmailAddressClass = _DDMatchEmailAddressClass{objc.GetClass("DDMatchEmailAddress")}
	})
	return DDMatchEmailAddressClass
}

type _DDMatchEmailAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchEmailAddress */
// An interface definition for the [DDMatchEmailAddress] class.
type IDDMatchEmailAddress interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchEmailAddress */
	// properties:
	EmailAddress() objc.IObject /* cross-framework: NSString */
	Label() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchEmailAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchEmailAddress */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchEmailAddressClass) Alloc() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchEmailAddressClass) New() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchEmailAddress) Init() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchEmailAddress) Autorelease() DDMatchEmailAddress {
	rv := objc.Send[DDMatchEmailAddress](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchEmailAddress creates a new DDMatchEmailAddress instance.
func NewDDMatchEmailAddress() DDMatchEmailAddress {
	return getDDMatchEmailAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchEmailAddress */
// An object that contains an email address that the data detection system matches.
//
// The DataDetection framework returns an email match in a object, which includes an email address, and optionally a label that categorizes the email address.


// An object that contains an email address that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress
type DDMatchEmailAddress struct {
	DDMatch
}

// DDMatchEmailAddressFrom constructs a [DDMatchEmailAddress] from an unsafe.Pointer.
//
// An object that contains an email address that the data detection system matches.
func DDMatchEmailAddressFrom(ptr unsafe.Pointer) DDMatchEmailAddress {
	return DDMatchEmailAddress{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchEmailAddress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchEmailAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchEmailAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchEmailAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchEmailAddress */

// A string that represents an email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress/emailAddress
func (d_ DDMatchEmailAddress) EmailAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("emailAddress"))
	return rv
}/* debug [instance_properties/getter]: emailAddress */


// A string that categorizes an email address, such as Home or Work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchEmailAddress/label
func (d_ DDMatchEmailAddress) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchEmailAddress */



