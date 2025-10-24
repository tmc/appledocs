// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSPerson */


/* debug [class_header]: Header for CSPerson */
// The class instance for the [CSPerson] class.
var (
	CSPersonClass     _CSPersonClass
	CSPersonClassOnce sync.Once
)

func getCSPersonClass() _CSPersonClass {
	CSPersonClassOnce.Do(func() {
		CSPersonClass = _CSPersonClass{objc.GetClass("CSPerson")}
	})
	return CSPersonClass
}

type _CSPersonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSPerson */
// An interface definition for the [CSPerson] class.
type ICSPerson interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSPerson */
	// properties:
	ContactIdentifier() objc.IObject /* cross-framework: NSString */
	SetContactIdentifier(value objc.IObject /* cross-framework: NSString */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	HandleIdentifier() objc.IObject /* cross-framework: NSString */
	Handles() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSPerson */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSPerson */
// Alloc allocates a new instance without initialization.
func (cc _CSPersonClass) Alloc() CSPerson {
	rv := objc.Send[CSPerson](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSPersonClass) New() CSPerson {
	rv := objc.Send[CSPerson](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSPerson) Init() CSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSPerson) Autorelease() CSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSPerson creates a new CSPerson instance.
func NewCSPerson() CSPerson {
	return getCSPersonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSPerson */
// An object that represents a person in the context of search results.
//
// A object represents a person in the context of search results. You can create a object when you have a display name and a contact handle of some kind, such as an email address or phone number. If you create a object to represent a specific contact, you can use the value of the contact’s identifier property for the person object’s property. Using the same value lets you avoid using names or phone numbers to look up the contact that’s associated with a person.


// An object that represents a person in the context of search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson
type CSPerson struct {
	objectivec.Object
}

// CSPersonFrom constructs a [CSPerson] from an unsafe.Pointer.
//
// An object that represents a person in the context of search results.
func CSPersonFrom(ptr unsafe.Pointer) CSPerson {
	return CSPerson{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSPerson */

// Returns a new object initialized with the specified display name and contact attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(displayName:handles:handleIdentifier:)
func NewCSPersonWithDisplayNameHandlesHandleIdentifier(displayName objc.IObject /* cross-framework: NSString */, handles []string, handleIdentifier objc.IObject /* cross-framework: NSString */) CSPerson {
	instance := getCSPersonClass().Alloc()
	rv := objc.Send[CSPerson](instance.ID, objc.Sel("initWithDisplayName:handles:handleIdentifier:"), displayName, handles, handleIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSPersonWithDisplayNameHandlesHandleIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSPerson */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSPerson */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSPerson */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSPerson */

// The identifier for the contact associated with the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/contactIdentifier
func (c_ CSPerson) ContactIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}/* debug [instance_properties/getter]: contactIdentifier */


// The identifier for the contact associated with the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/contactIdentifier
func (c_ CSPerson) SetContactIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifier:"), value)
}/* debug [instance_properties/setter]: contactIdentifier */


// A display name for the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/displayName
func (c_ CSPerson) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A key that identifies the type of contact property represented by the person object’s handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handleIdentifier
func (c_ CSPerson) HandleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("handleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: handleIdentifier */


// An array of contact handles related to the person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handles
func (c_ CSPerson) Handles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("handles"))
	return rv
}/* debug [instance_properties/getter]: handles */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSPerson */


