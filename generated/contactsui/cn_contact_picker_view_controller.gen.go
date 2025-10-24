// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNContactPickerViewController */


/* debug [class_header]: Header for CNContactPickerViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactPickerViewController */
// An interface definition for the [CNContactPickerViewController] class.
type ICNContactPickerViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for CNContactPickerViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactPickerViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactPickerViewController */
// Alloc allocates a new instance without initialization.
func (cc _CNContactPickerViewControllerClass) Alloc() CNContactPickerViewController {
	rv := objc.Send[CNContactPickerViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactPickerViewController */
// A view controller that displays an interface for picking contacts.
//
// A allows the user to select one or more contacts (or their properties) from the list of contacts displayed in the contact view controller ( ). The picker supports both single selection and multiselection of the contacts. The app using contact picker view does not need access to the user’s contacts and the user will not be prompted for “grant permission” access. The app has access only to the user’s final selection. There are predefined predicates in this class that let you control the user selection of the contact. Changing the predicates only take effect before the view is presented.


// A view controller that displays an interface for picking contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController
type CNContactPickerViewController struct {
	ViewController
}

// CNContactPickerViewControllerFrom constructs a [CNContactPickerViewController] from an unsafe.Pointer.
//
// A view controller that displays an interface for picking contacts.
func CNContactPickerViewControllerFrom(ptr unsafe.Pointer) CNContactPickerViewController {
	return CNContactPickerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactPickerViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactPickerViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactPickerViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactPickerViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactPickerViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactPickerViewController */


