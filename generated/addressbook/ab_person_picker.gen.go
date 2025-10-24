// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABPersonPicker */


/* debug [class_header]: Header for ABPersonPicker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABPersonPicker */
// An interface definition for the [ABPersonPicker] class.
type IABPersonPicker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ABPersonPicker */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Properties() objc.IObject /* cross-framework: NSArray */
	SetProperties(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABPersonPicker */
	// methods:
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect /* not a class type */, positioningView appkit.View, preferredEdge RectEdge /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABPersonPicker */
// Alloc allocates a new instance without initialization.
func (ac _ABPersonPickerClass) Alloc() ABPersonPicker {
	rv := objc.Send[ABPersonPicker](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABPersonPicker */
// A picker object that you display when you want the user to select contacts.


// A picker object that you display when you want the user to select contacts.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABPersonPicker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABPersonPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABPersonPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABPersonPicker */

// Closes the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/close
func (a_ ABPersonPicker) Close() {
	objc.Send[objc.ID](a_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Shows the picker in a popover relative to a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/showRelativeToRect:ofView:preferredEdge:
func (a_ ABPersonPicker) ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect /* not a class type */, positioningView appkit.View, preferredEdge RectEdge /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}/* debug [instance_methods/method]: ShowRelativeToRectOfViewPreferredEdge */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABPersonPicker */

// An object the system notifies as the user interacts with the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/delegate
func (a_ ABPersonPicker) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object the system notifies as the user interacts with the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/delegate
func (a_ ABPersonPicker) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// An array of properties to display in the picker when the user selects a person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/properties
func (a_ ABPersonPicker) Properties() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// An array of properties to display in the picker when the user selects a person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPersonPicker/properties
func (a_ ABPersonPicker) SetProperties(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProperties:"), value)
}/* debug [instance_properties/setter]: properties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABPersonPicker */



