// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [SFAuthorizationView] class.
var (
	SFAuthorizationViewClass     _SFAuthorizationViewClass
	SFAuthorizationViewClassOnce sync.Once
)

func getSFAuthorizationViewClass() _SFAuthorizationViewClass {
	SFAuthorizationViewClassOnce.Do(func() {
		SFAuthorizationViewClass = _SFAuthorizationViewClass{objc.GetClass("SFAuthorizationView")}
	})
	return SFAuthorizationViewClass
}

type _SFAuthorizationViewClass struct {
	class objc.Class
}

// An interface definition for the [SFAuthorizationView] class.
type ISFAuthorizationView interface {
	appkit.IView
	Authorization() unsafe.Pointer
	AuthorizationRights() unsafe.Pointer
	AuthorizationState() unsafe.Pointer
	Authorize(inSender objc.ID) bool
	Deauthorize(inSender objc.ID) bool
	Delegate() objc.ID
	IsEnabled() bool
	SetAuthorizationRights(authorizationRights unsafe.Pointer)
	SetAutoupdate(autoupdate bool)
	SetAutoupdateInterval(autoupdate bool, interval TimeInterval)
	SetDelegate(delegate objc.ID)
	SetEnabled(enabled bool)
	SetFlags(flags unsafe.Pointer)
	SetString(authorizationString unsafe.Pointer)
	UpdateStatus(inSender objc.ID) bool
}

// The class responsible for displaying a lock icon that can be used to indicate that a user interface has restricted access.
//
// The lock appears locked when the user must be authorized and appears open when the user has been authorized. The closed and open lock icons of the authorization view are shown in the following figure. When you add an authorization view as a custom view to a window or dialog box, you must initialize it before it displays correctly. To initialize the view, use the method to create a default rights structure (containing a prompt string) or the method to specify a rights structure. You must also either specify automatic updates ( or ) or perform a manual update ( ) to set the lock icon to its initial state. You can implement delegate methods that are invoked when the authorization view changes state. You can optionally implement the delegate methods to obtain the state of the authorization object when you are using an authorization view. When the user clicks a locked authorization view icon, the Security Server displays an authentication dialog (to request a user name and password, for example). When the user provides the requested credentials, the lock icon unlocks and the user is considered preauthorized to perform the functions specified by the authorization rights structure. You can call the method to determine whether the user has been preauthorized: this method returns if the view is in the unlocked state, otherwise . Before committing changes or performing actions that require authorization, you should check the user’s authorization again, even if they are preauthorized. The default behavior of this view is to preauthorize rights; if this is not possible it unlocks and waits for authorization to be checked when explicitly required.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView
type SFAuthorizationView struct {
	appkit.View
}

// SFAuthorizationViewFrom constructs a [SFAuthorizationView] from an unsafe.Pointer.
//
// The class responsible for displaying a lock icon that can be used to indicate that a user interface has restricted access.
func SFAuthorizationViewFrom(ptr unsafe.Pointer) SFAuthorizationView {
	return SFAuthorizationView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationViewClass) Alloc() SFAuthorizationView {
	rv := objc.Send[SFAuthorizationView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFAuthorizationViewClass) New() SFAuthorizationView {
	rv := objc.Send[SFAuthorizationView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAuthorizationView) Init() SFAuthorizationView {
	rv := objc.Send[SFAuthorizationView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAuthorizationView) Autorelease() SFAuthorizationView {
	rv := objc.Send[SFAuthorizationView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAuthorizationView creates a new SFAuthorizationView instance.
func NewSFAuthorizationView() SFAuthorizationView {
	return getSFAuthorizationViewClass().New()
}


// Returns the authorization object associated with this view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorization()
func (s_ SFAuthorizationView) Authorization() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorization"))
	return rv
}

// Returns the authorization rights for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorizationRights()
func (s_ SFAuthorizationView) AuthorizationRights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationRights"))
	return rv
}

// Returns the current state of the authorization view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorizationState()
func (s_ SFAuthorizationView) AuthorizationState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationState"))
	return rv
}

// Attempts to unlock the lock icon in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorize(_:)
func (s_ SFAuthorizationView) Authorize(inSender objc.ID) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorize:"), inSender)
	return rv
}

// Sets the authorization state to unauthorized and locks the lock icon in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/deauthorize(_:)
func (s_ SFAuthorizationView) Deauthorize(inSender objc.ID) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("deauthorize:"), inSender)
	return rv
}

// Returns the delegate for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/delegate()
func (s_ SFAuthorizationView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}

// Indicates whether the authorization view is enabled ( ) or disabled ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/isEnabled()
func (s_ SFAuthorizationView) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}

// Sets the authorization rights for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAuthorizationRights(_:)
func (s_ SFAuthorizationView) SetAuthorizationRights(authorizationRights unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuthorizationRights:"), authorizationRights)
}

// Sets the authorization view to update itself automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAutoupdate(_:)
func (s_ SFAuthorizationView) SetAutoupdate(autoupdate bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutoupdate:"), autoupdate)
}

// Sets the authorization view to update itself at a specific interval.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAutoupdate(_:interval:)
func (s_ SFAuthorizationView) SetAutoupdateInterval(autoupdate bool, interval TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutoupdate:interval:"), autoupdate, interval)
}

// Sets the delegate for this authorization view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setDelegate(_:)
func (s_ SFAuthorizationView) SetDelegate(delegate objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), delegate)
}

// Sets the current state of the authorization view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setEnabled(_:)
func (s_ SFAuthorizationView) SetEnabled(enabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), enabled)
}

// Sets the current authorization flags for the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setFlags(_:)
func (s_ SFAuthorizationView) SetFlags(flags unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFlags:"), flags)
}

// Sets the requested-right string to use with the default authorization rights set.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setString(_:)
func (s_ SFAuthorizationView) SetString(authorizationString unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setString:"), authorizationString)
}

// Manually updates the authorization view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/updateStatus(_:)
func (s_ SFAuthorizationView) UpdateStatus(inSender objc.ID) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("updateStatus:"), inSender)
	return rv
}



