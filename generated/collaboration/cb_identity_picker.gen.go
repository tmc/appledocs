// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBIdentityPicker */


/* debug [class_header]: Header for CBIdentityPicker */
// The class instance for the [CBIdentityPicker] class.
var (
	CBIdentityPickerClass     _CBIdentityPickerClass
	CBIdentityPickerClassOnce sync.Once
)

func getCBIdentityPickerClass() _CBIdentityPickerClass {
	CBIdentityPickerClassOnce.Do(func() {
		CBIdentityPickerClass = _CBIdentityPickerClass{objc.GetClass("CBIdentityPicker")}
	})
	return CBIdentityPickerClass
}

type _CBIdentityPickerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBIdentityPicker */
// An interface definition for the [CBIdentityPicker] class.
type ICBIdentityPicker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBIdentityPicker */
	// properties:
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	Identities() []CBIdentity
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBIdentityPicker */
	// methods:
	RunModal() int
	RunModalForWindowCompletionHandler(window appkit.Window, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBIdentityPicker */
// Alloc allocates a new instance without initialization.
func (cc _CBIdentityPickerClass) Alloc() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBIdentityPickerClass) New() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBIdentityPicker) Init() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBIdentityPicker) Autorelease() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBIdentityPicker creates a new CBIdentityPicker instance.
func NewCBIdentityPicker() CBIdentityPicker {
	return getCBIdentityPickerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBIdentityPicker */
// A object allows a user to select identities—for example, user or group objects—that it wants one or more services or shared resources to have access to. An identity picker can be displayed either as an application-modal dialog or as a sheet attached to a document window. An identity picker returns the selected records to be added to access control lists using Collaboration. If a selected record is not a user or group identity, then an identity picker prompts the user for additional information—such as a password—to promote that record to a sharing account.


// A object allows a user to select identities—for example, user or group objects—that it wants one or more services or shared resources to have access to. An identity picker can be displayed either as an application-modal dialog or as a sheet attached to a document window. An identity picker returns the selected records to be added to access control lists using Collaboration. If a selected record is not a user or group identity, then an identity picker prompts the user for additional information—such as a password—to promote that record to a sharing account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker
type CBIdentityPicker struct {
	objectivec.Object
}

// CBIdentityPickerFrom constructs a [CBIdentityPicker] from an unsafe.Pointer.
//
// A object allows a user to select identities—for example, user or group objects—that it wants one or more services or shared resources to have access to. An identity picker can be displayed either as an application-modal dialog or as a sheet attached to a document window. An identity picker returns the selected records to be added to access control lists using Collaboration. If a selected record is not a user or group identity, then an identity picker prompts the user for additional information—such as a password—to promote that record to a sharing account.
func CBIdentityPickerFrom(ptr unsafe.Pointer) CBIdentityPicker {
	return CBIdentityPicker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBIdentityPicker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBIdentityPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBIdentityPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBIdentityPicker */

// Runs the receiver as an application-modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/runModal()
func (c_ CBIdentityPicker) RunModal() int {
	rv := objc.Send[int](c_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Runs the identity picker modally as a sheet attached to a specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/runModal(for:completionHandler:)
func (c_ CBIdentityPicker) RunModalForWindowCompletionHandler(window appkit.Window, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("runModalForWindow:completionHandler:"), window, completionHandler)
}/* debug [instance_methods/method]: RunModalForWindowCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBIdentityPicker */

// A Boolean value indicating whether the user is allowed to select multiple identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/allowsMultipleSelection
func (c_ CBIdentityPicker) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsMultipleSelection */


// A Boolean value indicating whether the user is allowed to select multiple identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/allowsMultipleSelection
func (c_ CBIdentityPicker) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}/* debug [instance_properties/setter]: allowsMultipleSelection */


// The array of identities (represented by objects) selected using the identity picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/identities
func (c_ CBIdentityPicker) Identities() []CBIdentity {
	rv := objc.Send[[]CBIdentity](c_.ID, objc.Sel("identities"))
	return rv
}/* debug [instance_properties/getter]: identities */


// The title of the identity picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/title
func (c_ CBIdentityPicker) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the identity picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/title
func (c_ CBIdentityPicker) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBIdentityPicker */



