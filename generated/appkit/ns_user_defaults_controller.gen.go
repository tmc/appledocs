// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AppliesImmediately() bool
	SetAppliesImmediately(value bool)
	Defaults() foundation.UserDefaults
	HasUnappliedChanges() bool
	InitialValues() foundation.IDictionary
	SetInitialValues(value foundation.IDictionary)
	Values() objc.ID


	

	// methods:
	Revert(sender objectivec.IObject)
	RevertToInitialValues(sender objectivec.IObject)
	Save(sender objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsControllerClass) Alloc() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A controller that accesses user preference information for your app from the user’s defaults database.
//
// is a Cocoa bindings–compatible controller class. Properties of the shared instance of this class can be bound to user interface elements to access and modify values stored in .


// A controller that accesses user preference information for your app from the user’s defaults database.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(coder:)
func NewUserDefaultsControllerWithCoder(coder foundation.foundation.INSCoder) UserDefaultsController {
	instance := getUserDefaultsControllerClass().Alloc()
	rv := objc.Send[UserDefaultsController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns an initialized NSUserDefaultsController object using the NSUserDefaults instance specified in and the initial default values contained in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(defaults:initialValues:)
func NewUserDefaultsControllerWithDefaultsInitialValues(defaults foundation.UserDefaults, initialValues foundation.IDictionary) UserDefaultsController {
	instance := getUserDefaultsControllerClass().Alloc()
	rv := objc.Send[UserDefaultsController](instance.ID, objc.Sel("initWithDefaults:initialValues:"), defaults, initialValues)
	rv.Autorelease()
	return rv
}












// Returns the shared instance of NSUserDefaultsController, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/shared
func (uc _UserDefaultsControllerClass) SharedUserDefaultsController() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("sharedUserDefaultsController"))
	return rv
}






// Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/revert(_:)
func (u_ UserDefaultsController) Revert(sender objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("revert:"), sender)
}


// Causes the receiver to discard all edits and replace the values of all the user default properties with any corresponding values in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/revertToInitialValues(_:)
func (u_ UserDefaultsController) RevertToInitialValues(sender objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("revertToInitialValues:"), sender)
}


// Saves the values of the receiver’s user default properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/save(_:)
func (u_ UserDefaultsController) Save(sender objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("save:"), sender)
}







// Returns whether any changes made to bound user default properties are saved immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/appliesImmediately
func (u_ UserDefaultsController) AppliesImmediately() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("appliesImmediately"))
	return rv
}


// Returns whether any changes made to bound user default properties are saved immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/appliesImmediately
func (u_ UserDefaultsController) SetAppliesImmediately(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAppliesImmediately:"), value)
}


// Returns the instance of NSUserDefaults in use by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/defaults
func (u_ UserDefaultsController) Defaults() foundation.UserDefaults {
	rv := objc.Send[foundation.UserDefaults](u_.ID, objc.Sel("defaults"))
	return rv
}


// Returns whether the receiver has user default values that have not been saved to NSUserDefaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/hasUnappliedChanges
func (u_ UserDefaultsController) HasUnappliedChanges() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasUnappliedChanges"))
	return rv
}


// Returns a dictionary containing the receiver’s initial default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/initialValues
func (u_ UserDefaultsController) InitialValues() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](u_.ID, objc.Sel("initialValues"))
	return rv
}


// Returns a dictionary containing the receiver’s initial default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/initialValues
func (u_ UserDefaultsController) SetInitialValues(value foundation.IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInitialValues:"), value)
}


// Returns the shared instance of NSUserDefaultsController, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/shared
func (u_ UserDefaultsController) SharedUserDefaultsController() IUserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID, objc.Sel("sharedUserDefaultsController"))
	return rv
}


// Returns a key value coding compliant object that is used to access the user default properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/values
func (u_ UserDefaultsController) Values() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("values"))
	return rv
}







