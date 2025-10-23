// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [AuthorizationAppleIDButton] class.
var (
	AuthorizationAppleIDButtonClass     _AuthorizationAppleIDButtonClass
	AuthorizationAppleIDButtonClassOnce sync.Once
)

func getAuthorizationAppleIDButtonClass() _AuthorizationAppleIDButtonClass {
	AuthorizationAppleIDButtonClassOnce.Do(func() {
		AuthorizationAppleIDButtonClass = _AuthorizationAppleIDButtonClass{objc.GetClass("ASAuthorizationAppleIDButton")}
	})
	return AuthorizationAppleIDButtonClass
}

type _AuthorizationAppleIDButtonClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationAppleIDButton] class.
type IAuthorizationAppleIDButton interface {
	appkit.IControl
	// properties:
	CornerRadius() float64 /* primitive/slice/pointer. */
	SetCornerRadius(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// A control you add to your interface that enables users to initiate the Sign In with Apple flow.
//
// Choose one of the built-in button styles and types, and change the corner radius of the button by setting the property, but don’t otherwise modify the style of the button. Don’t use an Apple ID authorization button for any purpose other than to initiate the Sign In with Apple flow. After the user taps the button, create a request using the provider, and then use an instance of to execute the request. For more information about which Sign in with Apple buttons are available on different Apple platforms, see .


// A control you add to your interface that enables users to initiate the Sign In with Apple flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton
type AuthorizationAppleIDButton struct {
	appkit.Control
}

// AuthorizationAppleIDButtonFrom constructs a [AuthorizationAppleIDButton] from an unsafe.Pointer.
//
// A control you add to your interface that enables users to initiate the Sign In with Apple flow.
func AuthorizationAppleIDButtonFrom(ptr unsafe.Pointer) AuthorizationAppleIDButton {
	return AuthorizationAppleIDButton{
		Control: appkit.ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDButtonClass) Alloc() AuthorizationAppleIDButton {
	rv := objc.Send[AuthorizationAppleIDButton](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationAppleIDButtonClass) New() AuthorizationAppleIDButton {
	rv := objc.Send[AuthorizationAppleIDButton](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDButton) Init() AuthorizationAppleIDButton {
	rv := objc.Send[AuthorizationAppleIDButton](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDButton) Autorelease() AuthorizationAppleIDButton {
	rv := objc.Send[AuthorizationAppleIDButton](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDButton creates a new AuthorizationAppleIDButton instance.
func NewAuthorizationAppleIDButton() AuthorizationAppleIDButton {
	return getAuthorizationAppleIDButtonClass().New()
}



// Creates a new Sign In with Apple authorization button with the given type and style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/init(authorizationButtonType:authorizationButtonStyle:)
func NewAuthorizationAppleIDButtonWithAuthorizationButtonTypeAuthorizationButtonStyle(type_ AuthorizationAppleIDButtonType, style AuthorizationAppleIDButtonStyle) AuthorizationAppleIDButton {
	instance := getAuthorizationAppleIDButtonClass().Alloc()
	rv := objc.Send[AuthorizationAppleIDButton](instance.ID, objc.Sel("initWithAuthorizationButtonType:authorizationButtonStyle:"), type_, style)
	rv.Autorelease()
	return rv
}


// Creates a new Sign In with Apple authorization button with the given type and style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/init(type:style:)
func NewAuthorizationAppleIDButtonWithTypeStyle(type_ AuthorizationAppleIDButtonType, style AuthorizationAppleIDButtonStyle) AuthorizationAppleIDButton {
	rv := objc.Send[AuthorizationAppleIDButton](objc.ID(getAuthorizationAppleIDButtonClass().class), objc.Sel("buttonWithType:style:"), type_, style)
	return rv
}



// Creates a new Sign In with Apple authorization button with the given type and style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/init(type:style:)
func (ac _AuthorizationAppleIDButtonClass) ButtonWithTypeStyle(type_ AuthorizationAppleIDButtonType, style AuthorizationAppleIDButtonStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("buttonWithType:style:"), type_, style)
	return rv
}


// The radius, in points, for the rounded corners on the Apple ID sign-in button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/cornerRadius
func (a_ AuthorizationAppleIDButton) CornerRadius() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("cornerRadius"))
	return rv
}


// The radius, in points, for the rounded corners on the Apple ID sign-in button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/cornerRadius
func (a_ AuthorizationAppleIDButton) SetCornerRadius(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCornerRadius:"), value)
}


