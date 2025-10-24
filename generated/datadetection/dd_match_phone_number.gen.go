// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchPhoneNumber */


/* debug [class_header]: Header for DDMatchPhoneNumber */
// The class instance for the [DDMatchPhoneNumber] class.
var (
	DDMatchPhoneNumberClass     _DDMatchPhoneNumberClass
	DDMatchPhoneNumberClassOnce sync.Once
)

func getDDMatchPhoneNumberClass() _DDMatchPhoneNumberClass {
	DDMatchPhoneNumberClassOnce.Do(func() {
		DDMatchPhoneNumberClass = _DDMatchPhoneNumberClass{objc.GetClass("DDMatchPhoneNumber")}
	})
	return DDMatchPhoneNumberClass
}

type _DDMatchPhoneNumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchPhoneNumber */
// An interface definition for the [DDMatchPhoneNumber] class.
type IDDMatchPhoneNumber interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchPhoneNumber */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	PhoneNumber() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchPhoneNumber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchPhoneNumber */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchPhoneNumberClass) Alloc() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchPhoneNumberClass) New() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchPhoneNumber) Init() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchPhoneNumber) Autorelease() DDMatchPhoneNumber {
	rv := objc.Send[DDMatchPhoneNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchPhoneNumber creates a new DDMatchPhoneNumber instance.
func NewDDMatchPhoneNumber() DDMatchPhoneNumber {
	return getDDMatchPhoneNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchPhoneNumber */
// An object that contains a phone number that the data detection system matches.
//
// The DataDetection framework returns a phone number match in a object, which contains a phone number, and optionally a label that categorizes the phone number.


// An object that contains a phone number that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber
type DDMatchPhoneNumber struct {
	DDMatch
}

// DDMatchPhoneNumberFrom constructs a [DDMatchPhoneNumber] from an unsafe.Pointer.
//
// An object that contains a phone number that the data detection system matches.
func DDMatchPhoneNumberFrom(ptr unsafe.Pointer) DDMatchPhoneNumber {
	return DDMatchPhoneNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchPhoneNumber *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchPhoneNumber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchPhoneNumber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchPhoneNumber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchPhoneNumber */

// A string that categorizes a phone number, such as Home or Work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber/label
func (d_ DDMatchPhoneNumber) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that represents a phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPhoneNumber/phoneNumber
func (d_ DDMatchPhoneNumber) PhoneNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("phoneNumber"))
	return rv
}/* debug [instance_properties/getter]: phoneNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchPhoneNumber */



