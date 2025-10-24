// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactsUserDefaults */


/* debug [class_header]: Header for CNContactsUserDefaults */
// The class instance for the [CNContactsUserDefaults] class.
var (
	CNContactsUserDefaultsClass     _CNContactsUserDefaultsClass
	CNContactsUserDefaultsClassOnce sync.Once
)

func getCNContactsUserDefaultsClass() _CNContactsUserDefaultsClass {
	CNContactsUserDefaultsClassOnce.Do(func() {
		CNContactsUserDefaultsClass = _CNContactsUserDefaultsClass{objc.GetClass("CNContactsUserDefaults")}
	})
	return CNContactsUserDefaultsClass
}

type _CNContactsUserDefaultsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactsUserDefaults */
// An interface definition for the [CNContactsUserDefaults] class.
type ICNContactsUserDefaults interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContactsUserDefaults */
	// properties:
	CountryCode() objc.IObject /* cross-framework: NSString */
	SortOrder() CNContactSortOrder
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactsUserDefaults */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactsUserDefaults */
// Alloc allocates a new instance without initialization.
func (cc _CNContactsUserDefaultsClass) Alloc() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactsUserDefaultsClass) New() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactsUserDefaults) Init() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactsUserDefaults) Autorelease() CNContactsUserDefaults {
	rv := objc.Send[CNContactsUserDefaults](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactsUserDefaults creates a new CNContactsUserDefaults instance.
func NewCNContactsUserDefaults() CNContactsUserDefaults {
	return getCNContactsUserDefaultsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactsUserDefaults */
// An object that defines the default options to use when displaying contacts.


// An object that defines the default options to use when displaying contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults
type CNContactsUserDefaults struct {
	objectivec.Object
}

// CNContactsUserDefaultsFrom constructs a [CNContactsUserDefaults] from an unsafe.Pointer.
//
// An object that defines the default options to use when displaying contacts.
func CNContactsUserDefaultsFrom(ptr unsafe.Pointer) CNContactsUserDefaults {
	return CNContactsUserDefaults{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactsUserDefaults *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactsUserDefaults */

// The singleton contacts user defaults object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/shared()
func (cc _CNContactsUserDefaultsClass) SharedDefaults() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("sharedDefaults"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedDefaults) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactsUserDefaults */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactsUserDefaults */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactsUserDefaults */

// An ISO country code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/countryCode
func (c_ CNContactsUserDefaults) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_properties/getter]: countryCode */


// Default sorting order by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactsUserDefaults/sortOrder
func (c_ CNContactsUserDefaults) SortOrder() CNContactSortOrder {
	rv := objc.Send[CNContactSortOrder](c_.ID, objc.Sel("sortOrder"))
	return rv
}/* debug [instance_properties/getter]: sortOrder */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactsUserDefaults */



