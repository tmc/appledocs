// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContainer */


/* debug [class_header]: Header for CNContainer */
// The class instance for the [CNContainer] class.
var (
	CNContainerClass     _CNContainerClass
	CNContainerClassOnce sync.Once
)

func getCNContainerClass() _CNContainerClass {
	CNContainerClassOnce.Do(func() {
		CNContainerClass = _CNContainerClass{objc.GetClass("CNContainer")}
	})
	return CNContainerClass
}

type _CNContainerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContainer */
// An interface definition for the [CNContainer] class.
type ICNContainer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContainer */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	Type() CNContainerType
	CNContainerIdentifierKey() objc.IObject /* cross-framework: NSString */
	CNContainerNameKey() objc.IObject /* cross-framework: NSString */
	CNContainerTypeKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContainer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContainer */
// Alloc allocates a new instance without initialization.
func (cc _CNContainerClass) Alloc() CNContainer {
	rv := objc.Send[CNContainer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContainerClass) New() CNContainer {
	rv := objc.Send[CNContainer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContainer) Init() CNContainer {
	rv := objc.Send[CNContainer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContainer) Autorelease() CNContainer {
	rv := objc.Send[CNContainer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContainer creates a new CNContainer instance.
func NewCNContainer() CNContainer {
	return getCNContainerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContainer */
// An immutable object that represents a collection of contacts.
//
// A contact can be in only one container. CardDAV accounts usually have only one container whereas Exchange accounts may have multiple containers, where each container represents an Exchange folder. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents a collection of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer
type CNContainer struct {
	objectivec.Object
}

// CNContainerFrom constructs a [CNContainer] from an unsafe.Pointer.
//
// An immutable object that represents a collection of contacts.
func CNContainerFrom(ptr unsafe.Pointer) CNContainer {
	return CNContainer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContainer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContainer */

// Returns a predicate to find the container of the specified contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainerOfContact(withIdentifier:)
func (cc _CNContainerClass) PredicateForContainerOfContactWithIdentifier(contactIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainerOfContactWithIdentifier:"), contactIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContainerOfContactWithIdentifier) */


// Returns a predicate to find the container of the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainerOfGroup(withIdentifier:)
func (cc _CNContainerClass) PredicateForContainerOfGroupWithIdentifier(groupIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainerOfGroupWithIdentifier:"), groupIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContainerOfGroupWithIdentifier) */


// Returns a predicate to find the containers with the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainers(withIdentifiers:)
func (cc _CNContainerClass) PredicateForContainersWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainersWithIdentifiers:"), identifiers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContainersWithIdentifiers) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContainer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContainer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContainer */

// The unique identifier for a contacts container on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/identifier
func (c_ CNContainer) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The name of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/name
func (c_ CNContainer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The type of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/type
func (c_ CNContainer) Type() CNContainerType {
	rv := objc.Send[CNContainerType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The identifier key of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontaineridentifierkey
func (c_ CNContainer) CNContainerIdentifierKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContainerIdentifierKey"))
	return rv
}/* debug [instance_properties/getter]: CNContainerIdentifierKey */


// The name of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainernamekey
func (c_ CNContainer) CNContainerNameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContainerNameKey"))
	return rv
}/* debug [instance_properties/getter]: CNContainerNameKey */


// The type of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainertypekey
func (c_ CNContainer) CNContainerTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContainerTypeKey"))
	return rv
}/* debug [instance_properties/getter]: CNContainerTypeKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContainer */



