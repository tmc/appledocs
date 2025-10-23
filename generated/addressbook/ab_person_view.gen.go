// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [ABPersonView] class.
type IABPersonView interface {
	appkit.IView
	// properties:
	Editing() bool /* primitive/slice/pointer. */
	SetEditing(value bool /* primitive/slice/pointer. */)
	Person() IABPerson
	SetPerson(value IABPerson)
	ShouldShowLinkedPeople() bool /* primitive/slice/pointer. */
	SetShouldShowLinkedPeople(value bool /* primitive/slice/pointer. */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _ABPersonViewClass) Alloc() ABPersonView {
	rv := objc.Send[ABPersonView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether the person view is in editing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/editing
func (a_ ABPersonView) Editing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("editing"))
	return rv
}


// A Boolean value that indicates whether the person view is in editing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/editing
func (a_ ABPersonView) SetEditing(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEditing:"), value)
}


// The contact record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/person
func (a_ ABPersonView) Person() IABPerson {
	rv := objc.Send[ABPerson](a_.ID, objc.Sel("person"))
	return rv
}


// The contact record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/person
func (a_ ABPersonView) SetPerson(value IABPerson) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerson:"), value)
}


// Indicates whether the person view should display data from person records that are linked with the person record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/shouldShowLinkedPeople
func (a_ ABPersonView) ShouldShowLinkedPeople() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldShowLinkedPeople"))
	return rv
}


// Indicates whether the person view should display data from person records that are linked with the person record being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonView/shouldShowLinkedPeople
func (a_ ABPersonView) SetShouldShowLinkedPeople(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldShowLinkedPeople:"), value)
}



