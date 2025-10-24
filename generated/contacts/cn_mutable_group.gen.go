// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNMutableGroup */


/* debug [class_header]: Header for CNMutableGroup */
// The class instance for the [CNMutableGroup] class.
var (
	CNMutableGroupClass     _CNMutableGroupClass
	CNMutableGroupClassOnce sync.Once
)

func getCNMutableGroupClass() _CNMutableGroupClass {
	CNMutableGroupClassOnce.Do(func() {
		CNMutableGroupClass = _CNMutableGroupClass{objc.GetClass("CNMutableGroup")}
	})
	return CNMutableGroupClass
}

type _CNMutableGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNMutableGroup */
// An interface definition for the [CNMutableGroup] class.
type ICNMutableGroup interface {
	ICNGroup
	
/* debug [class_interface_properties]: Properties for CNMutableGroup */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNMutableGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNMutableGroup */
// Alloc allocates a new instance without initialization.
func (cc _CNMutableGroupClass) Alloc() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNMutableGroupClass) New() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNMutableGroup) Init() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNMutableGroup) Autorelease() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNMutableGroup creates a new CNMutableGroup instance.
func NewCNMutableGroup() CNMutableGroup {
	return getCNMutableGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNMutableGroup */
// A mutable object that represents a group of contacts.
//
// Contacts may be members of one or more groups, depending upon the accounts they come from. The class is not a thread-safe class.


// A mutable object that represents a group of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableGroup
type CNMutableGroup struct {
	CNGroup
}

// CNMutableGroupFrom constructs a [CNMutableGroup] from an unsafe.Pointer.
//
// A mutable object that represents a group of contacts.
func CNMutableGroupFrom(ptr unsafe.Pointer) CNMutableGroup {
	return CNMutableGroup{
		CNGroup: CNGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNMutableGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNMutableGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNMutableGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNMutableGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNMutableGroup */

// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableGroup/name
func (c_ CNMutableGroup) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableGroup/name
func (c_ CNMutableGroup) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNMutableGroup */



