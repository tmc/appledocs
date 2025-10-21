// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserDefaultsController] class.
var (
	UserDefaultsControllerClass     _UserDefaultsControllerClass
	UserDefaultsControllerClassOnce sync.Once
)

func getUserDefaultsControllerClass() _UserDefaultsControllerClass {
	UserDefaultsControllerClassOnce.Do(func() {
		UserDefaultsControllerClass = _UserDefaultsControllerClass{objc.GetClass("NSUserDefaultsController")}
	})
	return UserDefaultsControllerClass
}

type _UserDefaultsControllerClass struct {
	class objc.Class
}

// An interface definition for the [UserDefaultsController] class.
type IUserDefaultsController interface {
	IController
	Revert(sender objc.ID)
}

// A controller that accesses user preference information for your app from the user’s defaults database.
//
// is a Cocoa bindings–compatible controller class. Properties of the shared instance of this class can be bound to user interface elements to access and modify values stored in .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController
type UserDefaultsController struct {
	Controller
}

// UserDefaultsControllerFrom constructs a [UserDefaultsController] from an unsafe.Pointer.
//
// A controller that accesses user preference information for your app from the user’s defaults database.
func UserDefaultsControllerFrom(ptr unsafe.Pointer) UserDefaultsController {
	return UserDefaultsController{
		Controller: ControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsControllerClass) Alloc() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserDefaultsControllerClass) New() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserDefaultsController) Init() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserDefaultsController) Autorelease() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserDefaultsController creates a new UserDefaultsController instance.
func NewUserDefaultsController() UserDefaultsController {
	return getUserDefaultsControllerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(coder:)
func NewUserDefaultsControllerWithCoder(coder unsafe.Pointer) UserDefaultsController {
	instance := getUserDefaultsControllerClass().Alloc()
	rv := objc.Send[UserDefaultsController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/revert(_:)
func (u_ UserDefaultsController) Revert(sender objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("revert:"), sender)
}

// Returns a key value coding compliant object that is used to access the user default properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/values
func (u_ UserDefaultsController) Values() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("values"))
	return rv
}

// Returns whether any changes made to bound user default properties are saved immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/appliesimmediately
func (u_ UserDefaultsController) AppliesImmediately() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("appliesImmediately"))
	return rv
}


// SetAppliesImmediately sets the value of the appliesImmediately property.
// Returns whether any changes made to bound user default properties are saved immediately.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/appliesimmediately
func (u_ UserDefaultsController) SetAppliesImmediately(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppliesImmediately:"), value)
}

// Returns the instance of NSUserDefaults in use by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/defaults
func (u_ UserDefaultsController) Defaults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("defaults"))
	return rv
}


// SetDefaults sets the value of the defaults property.
// Returns the instance of NSUserDefaults in use by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/defaults
func (u_ UserDefaultsController) SetDefaults(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDefaults:"), value)
}

// Returns whether the receiver has user default values that have not been saved to NSUserDefaults.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/hasunappliedchanges
func (u_ UserDefaultsController) HasUnappliedChanges() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasUnappliedChanges"))
	return rv
}


// SetHasUnappliedChanges sets the value of the hasUnappliedChanges property.
// Returns whether the receiver has user default values that have not been saved to NSUserDefaults.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/hasunappliedchanges
func (u_ UserDefaultsController) SetHasUnappliedChanges(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasUnappliedChanges:"), value)
}

// Returns a dictionary containing the receiver’s initial default values.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/initialvalues
func (u_ UserDefaultsController) InitialValues() string {
	rv := objc.Send[string](u_.ID, objc.Sel("initialValues"))
	return rv
}


// SetInitialValues sets the value of the initialValues property.
// Returns a dictionary containing the receiver’s initial default values.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/initialvalues
func (u_ UserDefaultsController) SetInitialValues(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInitialValues:"), objc.String(value))
}


