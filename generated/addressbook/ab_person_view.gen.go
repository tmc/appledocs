// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class ABPersonView */


/* debug [class_header]: Header for ABPersonView */
// The class instance for the [ABPersonView] class.
var (
	ABPersonViewClass     _ABPersonViewClass
	ABPersonViewClassOnce sync.Once
)

func getABPersonViewClass() _ABPersonViewClass {
	ABPersonViewClassOnce.Do(func() {
		ABPersonViewClass = _ABPersonViewClass{objc.GetClass("ABPersonView")}
	})
	return ABPersonViewClass
}

type _ABPersonViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABPersonView */
// An interface definition for the [ABPersonView] class.
type IABPersonView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for ABPersonView */
	// properties:
	Editing() bool
	SetEditing(value bool)
	Person() IABPerson
	SetPerson(value IABPerson)
	ShouldShowLinkedPeople() bool
	SetShouldShowLinkedPeople(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABPersonView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABPersonView */
// Alloc allocates a new instance without initialization.
func (ac _ABPersonViewClass) Alloc() ABPersonView {
	rv := objc.Send[ABPersonView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ABPersonViewClass) New() ABPersonView {
	rv := objc.Send[ABPersonView](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABPersonView) Init() ABPersonView {
	rv := objc.Send[ABPersonView](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABPersonView) Autorelease() ABPersonView {
	rv := objc.Send[ABPersonView](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABPersonView creates a new ABPersonView instance.
func NewABPersonView() ABPersonView {
	return getABPersonViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABPersonView */
// An object that provides a view for displaying and editing contacts.


// An object that provides a view for displaying and editing contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView
type ABPersonView struct {
	appkit.View
}

// ABPersonViewFrom constructs a [ABPersonView] from an unsafe.Pointer.
//
// An object that provides a view for displaying and editing contacts.
func ABPersonViewFrom(ptr unsafe.Pointer) ABPersonView {
	return ABPersonView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABPersonView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABPersonView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABPersonView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABPersonView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABPersonView */

// A Boolean value that indicates whether the person view is in editing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/editing
func (a_ ABPersonView) Editing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("editing"))
	return rv
}/* debug [instance_properties/getter]: editing */


// A Boolean value that indicates whether the person view is in editing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/editing
func (a_ ABPersonView) SetEditing(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEditing:"), value)
}/* debug [instance_properties/setter]: editing */


// The contact record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/person
func (a_ ABPersonView) Person() IABPerson {
	rv := objc.Send[ABPerson](a_.ID, objc.Sel("person"))
	return rv
}/* debug [instance_properties/getter]: person */


// The contact record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/person
func (a_ ABPersonView) SetPerson(value IABPerson) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerson:"), value)
}/* debug [instance_properties/setter]: person */


// Indicates whether the person view should display data from person records that are linked with the person record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/shouldShowLinkedPeople
func (a_ ABPersonView) ShouldShowLinkedPeople() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldShowLinkedPeople"))
	return rv
}/* debug [instance_properties/getter]: shouldShowLinkedPeople */


// Indicates whether the person view should display data from person records that are linked with the person record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/shouldShowLinkedPeople
func (a_ ABPersonView) SetShouldShowLinkedPeople(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldShowLinkedPeople:"), value)
}/* debug [instance_properties/setter]: shouldShowLinkedPeople */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABPersonView */



