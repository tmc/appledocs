// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ABSearchElement] class.
type IABSearchElement interface {
	objectivec.IObject
	MatchesRecord(record unsafe.Pointer) bool
}

// An object you use to specify a search query for records in the Address Book database.
//
// The class is “toll-free bridged” with its procedural C opaque-type counterpart. This means that the type is interchangeable in function or method calls with instances of the class.
//
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

// Alloc allocates a new instance without initialization.
func (ac _ABSearchElementClass) Alloc() ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns a compound search element, created by combining the search elements in an array with the given conjunction.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/init(forConjunction:children:)
func NewABSearchElementForConjunctionChildren(conjuction unsafe.Pointer, children objc.ID) ABSearchElement {
	rv := objc.Send[ABSearchElement](objc.ID(getABSearchElementClass().class), objc.Sel("searchElementForConjunction:children:"), conjuction, children)
	return rv
}


// Returns a compound search element, created by combining the search elements in an array with the given conjunction.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/init(forConjunction:children:)
func (ac _ABSearchElementClass) SearchElementForConjunctionChildren(conjuction unsafe.Pointer, children objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("searchElementForConjunction:children:"), conjuction, children)
	return rv
}

// Tests whether or not a record matches a search element.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABSearchElement/matchesRecord(_:)
func (a_ ABSearchElement) MatchesRecord(record unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("matchesRecord:"), record)
	return rv
}


