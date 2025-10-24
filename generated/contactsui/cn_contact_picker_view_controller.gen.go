// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNContactPickerViewController] class.
var (
	CNContactPickerViewControllerClass     _CNContactPickerViewControllerClass
	CNContactPickerViewControllerClassOnce sync.Once
)

func getCNContactPickerViewControllerClass() _CNContactPickerViewControllerClass {
	CNContactPickerViewControllerClassOnce.Do(func() {
		CNContactPickerViewControllerClass = _CNContactPickerViewControllerClass{objc.GetClass("CNContactPickerViewController")}
	})
	return CNContactPickerViewControllerClass
}

type _CNContactPickerViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [CNContactPickerViewController] class.
type ICNContactPickerViewController interface {
	appkit.IViewController
	// properties:
	// methods:
}

// A view controller that displays an interface for picking contacts.
//
// A allows the user to select one or more contacts (or their properties) from the list of contacts displayed in the contact view controller ( ). The picker supports both single selection and multiselection of the contacts. The app using contact picker view does not need access to the user’s contacts and the user will not be prompted for “grant permission” access. The app has access only to the user’s final selection. There are predefined predicates in this class that let you control the user selection of the contact. Changing the predicates only take effect before the view is presented.


// A view controller that displays an interface for picking contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController
type CNContactPickerViewController struct {
	appkit.ViewController
}

// CNContactPickerViewControllerFrom constructs a [CNContactPickerViewController] from an unsafe.Pointer.
//
// A view controller that displays an interface for picking contacts.
func CNContactPickerViewControllerFrom(ptr unsafe.Pointer) CNContactPickerViewController {
	return CNContactPickerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactPickerViewControllerClass) Alloc() CNContactPickerViewController {
	rv := objc.Send[CNContactPickerViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactPickerViewControllerClass) New() CNContactPickerViewController {
	rv := objc.Send[CNContactPickerViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactPickerViewController) Init() CNContactPickerViewController {
	rv := objc.Send[CNContactPickerViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactPickerViewController) Autorelease() CNContactPickerViewController {
	rv := objc.Send[CNContactPickerViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactPickerViewController creates a new CNContactPickerViewController instance.
func NewCNContactPickerViewController() CNContactPickerViewController {
	return getCNContactPickerViewControllerClass().New()
}



