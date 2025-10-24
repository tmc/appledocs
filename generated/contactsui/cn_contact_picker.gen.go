// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactPicker */


/* debug [class_header]: Header for CNContactPicker */
// The class instance for the [CNContactPicker] class.
var (
	CNContactPickerClass     _CNContactPickerClass
	CNContactPickerClassOnce sync.Once
)

func getCNContactPickerClass() _CNContactPickerClass {
	CNContactPickerClassOnce.Do(func() {
		CNContactPickerClass = _CNContactPickerClass{objc.GetClass("CNContactPicker")}
	})
	return CNContactPickerClass
}

type _CNContactPickerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactPicker */
// An interface definition for the [CNContactPicker] class.
type ICNContactPicker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContactPicker */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplayedKeys() []string
	SetDisplayedKeys(value []string)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactPicker */
	// methods:
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect /* not a class type */, positioningView appkit.View, preferredEdge RectEdge /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactPicker */
// Alloc allocates a new instance without initialization.
func (cc _CNContactPickerClass) Alloc() CNContactPicker {
	rv := objc.Send[CNContactPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactPickerClass) New() CNContactPicker {
	rv := objc.Send[CNContactPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactPicker) Init() CNContactPicker {
	rv := objc.Send[CNContactPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactPicker) Autorelease() CNContactPicker {
	rv := objc.Send[CNContactPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactPicker creates a new CNContactPicker instance.
func NewCNContactPicker() CNContactPicker {
	return getCNContactPickerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactPicker */
// A popover-based interface for selecting a contact.
//
// Before displaying the popover, configure the property with the information you want to display in the interface.


// A popover-based interface for selecting a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker
type CNContactPicker struct {
	objectivec.Object
}

// CNContactPickerFrom constructs a [CNContactPicker] from an unsafe.Pointer.
//
// A popover-based interface for selecting a contact.
func CNContactPickerFrom(ptr unsafe.Pointer) CNContactPicker {
	return CNContactPicker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactPicker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactPicker */

// Closes the popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/close()
func (c_ CNContactPicker) Close() {
	objc.Send[objc.ID](c_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Shows the picker popover anchored to the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/showRelative(to:of:preferredEdge:)
func (c_ CNContactPicker) ShowRelativeToRectOfViewPreferredEdge(positioningRect Rect /* not a class type */, positioningView appkit.View, preferredEdge RectEdge /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}/* debug [instance_methods/method]: ShowRelativeToRectOfViewPreferredEdge */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactPicker */

// The picker delegate to be notified when the user chooses a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/delegate
func (c_ CNContactPicker) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The picker delegate to be notified when the user chooses a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/delegate
func (c_ CNContactPicker) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The keys to be displayed when a contact is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/displayedKeys
func (c_ CNContactPicker) DisplayedKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("displayedKeys"))
	return rv
}/* debug [instance_properties/getter]: displayedKeys */


// The keys to be displayed when a contact is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/displayedKeys
func (c_ CNContactPicker) SetDisplayedKeys(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayedKeys:"), nsArray)
}/* debug [instance_properties/setter]: displayedKeys */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactPicker */



