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
}

// A view controller that displays an interface for picking contacts.
//
// A allows the user to select one or more contacts (or their properties) from the list of contacts displayed in the contact view controller ( ). The picker supports both single selection and multiselection of the contacts. The app using contact picker view does not need access to the user’s contacts and the user will not be prompted for “grant permission” access. The app has access only to the user’s final selection. There are predefined predicates in this class that let you control the user selection of the contact. Changing the predicates only take effect before the view is presented.
//
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


// The delegate to be notified when the user selects a contact or a property.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/delegate
func (c_ CNContactPickerViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate to be notified when the user selects a contact or a property.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/delegate
func (c_ CNContactPickerViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The property keys to display in the contact detail card.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/displayedPropertyKeys
func (c_ CNContactPickerViewController) DisplayedPropertyKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("displayedPropertyKeys"))
	return rv
}


// SetDisplayedPropertyKeys sets the value of the displayedPropertyKeys property.
// The property keys to display in the contact detail card.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/displayedPropertyKeys
func (c_ CNContactPickerViewController) SetDisplayedPropertyKeys(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayedPropertyKeys:"), nsArray)
}

// A predicate to determine the contact selectability in the list of contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForEnablingContact
func (c_ CNContactPickerViewController) PredicateForEnablingContact() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForEnablingContact"))
	return rv
}


// SetPredicateForEnablingContact sets the value of the predicateForEnablingContact property.
// A predicate to determine the contact selectability in the list of contacts.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForEnablingContact
func (c_ CNContactPickerViewController) SetPredicateForEnablingContact(value foundation.IPredicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicateForEnablingContact:"), value)
}

// A predicate to control the return of the selected contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfContact
func (c_ CNContactPickerViewController) PredicateForSelectionOfContact() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForSelectionOfContact"))
	return rv
}


// SetPredicateForSelectionOfContact sets the value of the predicateForSelectionOfContact property.
// A predicate to control the return of the selected contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfContact
func (c_ CNContactPickerViewController) SetPredicateForSelectionOfContact(value foundation.IPredicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicateForSelectionOfContact:"), value)
}

// A predicate to control the properties of the selected contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfProperty
func (c_ CNContactPickerViewController) PredicateForSelectionOfProperty() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForSelectionOfProperty"))
	return rv
}


// SetPredicateForSelectionOfProperty sets the value of the predicateForSelectionOfProperty property.
// A predicate to control the properties of the selected contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfProperty
func (c_ CNContactPickerViewController) SetPredicateForSelectionOfProperty(value foundation.IPredicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicateForSelectionOfProperty:"), value)
}



