// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccountAuthenticationModificationController] class.
var (
	AccountAuthenticationModificationControllerClass     _AccountAuthenticationModificationControllerClass
	AccountAuthenticationModificationControllerClassOnce sync.Once
)

func getAccountAuthenticationModificationControllerClass() _AccountAuthenticationModificationControllerClass {
	AccountAuthenticationModificationControllerClassOnce.Do(func() {
		AccountAuthenticationModificationControllerClass = _AccountAuthenticationModificationControllerClass{objc.GetClass("ASAccountAuthenticationModificationController")}
	})
	return AccountAuthenticationModificationControllerClass
}

type _AccountAuthenticationModificationControllerClass struct {
	class objc.Class
}

// An interface definition for the [AccountAuthenticationModificationController] class.
type IAccountAuthenticationModificationController interface {
	objectivec.IObject
}

// An object that performs a request to modify an account’s authentication properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController
type AccountAuthenticationModificationController struct {
	objectivec.Object
}

// AccountAuthenticationModificationControllerFrom constructs a [AccountAuthenticationModificationController] from an unsafe.Pointer.
//
// An object that performs a request to modify an account’s authentication properties.
func AccountAuthenticationModificationControllerFrom(ptr unsafe.Pointer) AccountAuthenticationModificationController {
	return AccountAuthenticationModificationController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationControllerClass) Alloc() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccountAuthenticationModificationControllerClass) New() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationController) Init() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationController) Autorelease() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationController creates a new AccountAuthenticationModificationController instance.
func NewAccountAuthenticationModificationController() AccountAuthenticationModificationController {
	return getAccountAuthenticationModificationControllerClass().New()
}


// An object that receives notifications about the request’s status.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/delegate
func (a_ AccountAuthenticationModificationController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that receives notifications about the request’s status.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/delegate
func (a_ AccountAuthenticationModificationController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// An object that provides a presentation context for the account modification request’s user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asaccountauthenticationmodificationcontroller/presentationcontextprovider
func (a_ AccountAuthenticationModificationController) PresentationContextProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationContextProvider"))
	return rv
}


// SetPresentationContextProvider sets the value of the presentationContextProvider property.
// An object that provides a presentation context for the account modification request’s user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asaccountauthenticationmodificationcontroller/presentationcontextprovider
func (a_ AccountAuthenticationModificationController) SetPresentationContextProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationContextProvider:"), value)
}



