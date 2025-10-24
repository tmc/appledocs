// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABSearchElement */


/* debug [class_header]: Header for ABSearchElement */
// The class instance for the [ABSearchElement] class.
var (
	ABSearchElementClass     _ABSearchElementClass
	ABSearchElementClassOnce sync.Once
)

func getABSearchElementClass() _ABSearchElementClass {
	ABSearchElementClassOnce.Do(func() {
		ABSearchElementClass = _ABSearchElementClass{objc.GetClass("ABSearchElement")}
	})
	return ABSearchElementClass
}

type _ABSearchElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABSearchElement */
// An interface definition for the [ABSearchElement] class.
type IABSearchElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ABSearchElement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABSearchElement */
	// methods:
	MatchesRecord(record IABRecord) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABSearchElement */
// Alloc allocates a new instance without initialization.
func (ac _ABSearchElementClass) Alloc() ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ABSearchElementClass) New() ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABSearchElement) Init() ABSearchElement {
	rv := objc.Send[ABSearchElement](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABSearchElement) Autorelease() ABSearchElement {
	rv := objc.Send[ABSearchElement](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABSearchElement creates a new ABSearchElement instance.
func NewABSearchElement() ABSearchElement {
	return getABSearchElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABSearchElement */
// An object you use to specify a search query for records in the Address Book database.
//
// The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.


// An object you use to specify a search query for records in the Address Book database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement
type ABSearchElement struct {
	objectivec.Object
}

// ABSearchElementFrom constructs a [ABSearchElement] from an unsafe.Pointer.
//
// An object you use to specify a search query for records in the Address Book database.
func ABSearchElementFrom(ptr unsafe.Pointer) ABSearchElement {
	return ABSearchElement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABSearchElement */

// Returns a compound search element, created by combining the search elements in an array with the given conjunction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/init(forConjunction:children:)
func NewABSearchElementForConjunctionChildren(conjuction ABSearchConjunction /* typedef */, children objc.IObject /* cross-framework: NSArray */) ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(getABSearchElementClass().class), objc.Sel("searchElementForConjunction:children:"), conjuction, children)
	return rv
}/* debug [class_init_methods/constructor]: NewABSearchElementForConjunctionChildren */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABSearchElement */

// Returns a compound search element, created by combining the search elements in an array with the given conjunction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/init(forConjunction:children:)
func (ac _ABSearchElementClass) SearchElementForConjunctionChildren(conjuction ABSearchConjunction /* typedef */, children objc.IObject /* cross-framework: NSArray */) ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("searchElementForConjunction:children:"), conjuction, children)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SearchElementForConjunctionChildren) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABSearchElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABSearchElement */

// Tests whether or not a record matches a search element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/matchesRecord(_:)
func (a_ ABSearchElement) MatchesRecord(record IABRecord) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("matchesRecord:"), record)
	return rv
}/* debug [instance_methods/method]: MatchesRecord */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABSearchElement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABSearchElement */


