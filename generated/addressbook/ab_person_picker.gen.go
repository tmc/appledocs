// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ABPersonPicker] class.
var (
	ABPersonPickerClass     _ABPersonPickerClass
	ABPersonPickerClassOnce sync.Once
)

func getABPersonPickerClass() _ABPersonPickerClass {
	ABPersonPickerClassOnce.Do(func() {
		ABPersonPickerClass = _ABPersonPickerClass{objc.GetClass("ABPersonPicker")}
	})
	return ABPersonPickerClass
}

type _ABPersonPickerClass struct {
	class objc.Class
}

// An interface definition for the [ABPersonPicker] class.
type IABPersonPicker interface {
	objectivec.IObject
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect, positioningView unsafe.Pointer, preferredEdge RectEdge)
}

// A picker object that you display when you want the user to select contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker
type ABPersonPicker struct {
	objectivec.Object
}

// ABPersonPickerFrom constructs a [ABPersonPicker] from an unsafe.Pointer.
//
// A picker object that you display when you want the user to select contacts.
func ABPersonPickerFrom(ptr unsafe.Pointer) ABPersonPicker {
	return ABPersonPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ABPersonPickerClass) Alloc() ABPersonPicker {
	rv := objc.Send[ABPersonPicker](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABPersonPickerClass) New() ABPersonPicker {
	rv := objc.Send[ABPersonPicker](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABPersonPicker) Init() ABPersonPicker {
	rv := objc.Send[ABPersonPicker](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABPersonPicker) Autorelease() ABPersonPicker {
	rv := objc.Send[ABPersonPicker](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABPersonPicker creates a new ABPersonPicker instance.
func NewABPersonPicker() ABPersonPicker {
	return getABPersonPickerClass().New()
}


// Closes the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/close
func (a_ ABPersonPicker) Close() {
	objc.Send[objc.ID](a_.ID, objc.Sel("close"))
}

// Shows the picker in a popover relative to a view.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/showRelativeToRect:ofView:preferredEdge:
func (a_ ABPersonPicker) ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect, positioningView unsafe.Pointer, preferredEdge RectEdge) {
	objc.Send[objc.ID](a_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}

// An object the system notifies as the user interacts with the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/delegate
func (a_ ABPersonPicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object the system notifies as the user interacts with the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/delegate
func (a_ ABPersonPicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// An array of properties to display in the picker when the user selects a person.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/properties
func (a_ ABPersonPicker) Properties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("properties"))
	return rv
}


// SetProperties sets the value of the properties property.
// An array of properties to display in the picker when the user selects a person.

//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/properties
func (a_ ABPersonPicker) SetProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProperties:"), value)
}



