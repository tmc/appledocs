// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNContactPicker] class.
type ICNContactPicker interface {
	objectivec.IObject
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.IRect, positioningView appkit.IView, preferredEdge foundation.IRectEdge)
}

// A popover-based interface for selecting a contact.
//
// Before displaying the popover, configure the property with the information you want to display in the interface.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CNContactPickerClass) Alloc() CNContactPicker {
	rv := objc.Send[CNContactPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Closes the popover.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/close()
func (c_ CNContactPicker) Close() {
	objc.Send[objc.ID](c_.ID, objc.Sel("close"))
}

// Shows the picker popover anchored to the specified view.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/showRelative(to:of:preferredEdge:)
func (c_ CNContactPicker) ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.IRect, positioningView appkit.IView, preferredEdge foundation.IRectEdge) {
	objc.Send[objc.ID](c_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}

// The picker delegate to be notified when the user chooses a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/delegate
func (c_ CNContactPicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The picker delegate to be notified when the user chooses a contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/delegate
func (c_ CNContactPicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The keys to be displayed when a contact is expanded.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/displayedKeys
func (c_ CNContactPicker) DisplayedKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("displayedKeys"))
	return rv
}


// SetDisplayedKeys sets the value of the displayedKeys property.
// The keys to be displayed when a contact is expanded.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPicker/displayedKeys
func (c_ CNContactPicker) SetDisplayedKeys(value []string) {
	// Convert Go slice to NSArray
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
}



