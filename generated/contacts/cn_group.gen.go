// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNGroup */


/* debug [class_header]: Header for CNGroup */
// The class instance for the [CNGroup] class.
var (
	CNGroupClass     _CNGroupClass
	CNGroupClassOnce sync.Once
)

func getCNGroupClass() _CNGroupClass {
	CNGroupClassOnce.Do(func() {
		CNGroupClass = _CNGroupClass{objc.GetClass("CNGroup")}
	})
	return CNGroupClass
}

type _CNGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNGroup */
// An interface definition for the [CNGroup] class.
type ICNGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNGroup */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	CNGroupIdentifierKey() objc.IObject /* cross-framework: NSString */
	CNGroupNameKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNGroup */
// Alloc allocates a new instance without initialization.
func (cc _CNGroupClass) Alloc() CNGroup {
	rv := objc.Send[CNGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNGroupClass) New() CNGroup {
	rv := objc.Send[CNGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNGroup) Init() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNGroup) Autorelease() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNGroup creates a new CNGroup instance.
func NewCNGroup() CNGroup {
	return getCNGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNGroup */
// An immutable object that represents a group of contacts.
//
// Contacts may be members of one or more groups, depending upon their accounts. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents a group of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup
type CNGroup struct {
	objectivec.Object
}

// CNGroupFrom constructs a [CNGroup] from an unsafe.Pointer.
//
// An immutable object that represents a group of contacts.
func CNGroupFrom(ptr unsafe.Pointer) CNGroup {
	return CNGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNGroup */

// Returns a predicate to find groups with the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup/predicateForGroups(withIdentifiers:)
func (cc _CNGroupClass) PredicateForGroupsWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForGroupsWithIdentifiers:"), identifiers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForGroupsWithIdentifiers) */


// Returns a predicate to find groups in the specified container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup/predicateForGroupsInContainer(withIdentifier:)
func (cc _CNGroupClass) PredicateForGroupsInContainerWithIdentifier(containerIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForGroupsInContainerWithIdentifier:"), containerIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForGroupsInContainerWithIdentifier) */


// Returns a predicate to find subgroups in the specified parent group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup/predicateForSubgroupsInGroup(withIdentifier:)
func (cc _CNGroupClass) PredicateForSubgroupsInGroupWithIdentifier(parentGroupIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForSubgroupsInGroupWithIdentifier:"), parentGroupIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForSubgroupsInGroupWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNGroup */

// The unique identifier for a group on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup/identifier
func (c_ CNGroup) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup/name
func (c_ CNGroup) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The identifier of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroupidentifierkey
func (c_ CNGroup) CNGroupIdentifierKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNGroupIdentifierKey"))
	return rv
}/* debug [instance_properties/getter]: CNGroupIdentifierKey */


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroupnamekey
func (c_ CNGroup) CNGroupNameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNGroupNameKey"))
	return rv
}/* debug [instance_properties/getter]: CNGroupNameKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNGroup */



