// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/securityfoundation"
)

/* debug [class.gen.go]: Generating class SFAuthorizationView */


/* debug [class_header]: Header for SFAuthorizationView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFAuthorizationView */
// An interface definition for the [SFAuthorizationView] class.
type ISFAuthorizationView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for SFAuthorizationView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFAuthorizationView */
	// methods:
	Authorization() securityfoundation.SFAuthorization
	AuthorizationRights() unsafe.Pointer
	AuthorizationState() unsafe.Pointer
	Authorize(inSender objc.IObject) bool
	Deauthorize(inSender objc.IObject) bool
	Delegate() objc.ID
	IsEnabled() bool
	SetAuthorizationRights(authorizationRights unsafe.Pointer)
	SetAutoupdate(autoupdate bool)
	SetAutoupdateInterval(autoupdate bool, interval float64)
	SetDelegate(delegate objc.IObject)
	SetEnabled(enabled bool)
	SetFlags(flags unsafe.Pointer)
	SetString(authorizationString unsafe.Pointer)
	UpdateStatus(inSender objc.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFAuthorizationView */
// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationViewClass) Alloc() SFAuthorizationView {
	rv := objc.Send[SFAuthorizationView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFAuthorizationView */
// The class responsible for displaying a lock icon that can be used to indicate that a user interface has restricted access.
//
// The lock appears locked when the user must be authorized and appears open when the user has been authorized. The closed and open lock icons of the authorization view are shown in the following figure. When you add an authorization view as a custom view to a window or dialog box, you must initialize it before it displays correctly. To initialize the view, use the method to create a default rights structure (containing a prompt string) or the method to specify a rights structure. You must also either specify automatic updates ( or ) or perform a manual update ( ) to set the lock icon to its initial state. You can implement delegate methods that are invoked when the authorization view changes state. You can optionally implement the delegate methods to obtain the state of the authorization object when you are using an authorization view. When the user clicks a locked authorization view icon, the Security Server displays an authentication dialog (to request a user name and password, for example). When the user provides the requested credentials, the lock icon unlocks and the user is considered preauthorized to perform the functions specified by the authorization rights structure. You can call the method to determine whether the user has been preauthorized: this method returns if the view is in the unlocked state, otherwise . Before committing changes or performing actions that require authorization, you should check the user’s authorization again, even if they are preauthorized. The default behavior of this view is to preauthorize rights; if this is not possible it unlocks and waits for authorization to be checked when explicitly required.


// The class responsible for displaying a lock icon that can be used to indicate that a user interface has restricted access.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFAuthorizationView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFAuthorizationView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFAuthorizationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFAuthorizationView */

// Returns the authorization object associated with this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorization()
func (s_ SFAuthorizationView) Authorization() securityfoundation.SFAuthorization {
	rv := objc.Send[securityfoundation.SFAuthorization](s_.ID, objc.Sel("authorization"))
	return rv
}/* debug [instance_methods/method]: Authorization */


// Returns the authorization rights for this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorizationRights()
func (s_ SFAuthorizationView) AuthorizationRights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationRights"))
	return rv
}/* debug [instance_methods/method]: AuthorizationRights */


// Returns the current state of the authorization view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorizationState()
func (s_ SFAuthorizationView) AuthorizationState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationState"))
	return rv
}/* debug [instance_methods/method]: AuthorizationState */


// Attempts to unlock the lock icon in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/authorize(_:)
func (s_ SFAuthorizationView) Authorize(inSender objc.IObject) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorize:"), inSender)
	return rv
}/* debug [instance_methods/method]: Authorize */


// Sets the authorization state to unauthorized and locks the lock icon in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/deauthorize(_:)
func (s_ SFAuthorizationView) Deauthorize(inSender objc.IObject) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("deauthorize:"), inSender)
	return rv
}/* debug [instance_methods/method]: Deauthorize */


// Returns the delegate for this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/delegate()
func (s_ SFAuthorizationView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// Indicates whether the authorization view is enabled ( ) or disabled ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/isEnabled()
func (s_ SFAuthorizationView) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_methods/method]: IsEnabled */


// Sets the authorization rights for this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAuthorizationRights(_:)
func (s_ SFAuthorizationView) SetAuthorizationRights(authorizationRights unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuthorizationRights:"), authorizationRights)
}/* debug [instance_methods/method]: SetAuthorizationRights */


// Sets the authorization view to update itself automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAutoupdate(_:)
func (s_ SFAuthorizationView) SetAutoupdate(autoupdate bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutoupdate:"), autoupdate)
}/* debug [instance_methods/method]: SetAutoupdate */


// Sets the authorization view to update itself at a specific interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setAutoupdate(_:interval:)
func (s_ SFAuthorizationView) SetAutoupdateInterval(autoupdate bool, interval float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutoupdate:interval:"), autoupdate, interval)
}/* debug [instance_methods/method]: SetAutoupdateInterval */


// Sets the delegate for this authorization view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setDelegate(_:)
func (s_ SFAuthorizationView) SetDelegate(delegate objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), delegate)
}/* debug [instance_methods/method]: SetDelegate */


// Sets the current state of the authorization view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setEnabled(_:)
func (s_ SFAuthorizationView) SetEnabled(enabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), enabled)
}/* debug [instance_methods/method]: SetEnabled */


// Sets the current authorization flags for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setFlags(_:)
func (s_ SFAuthorizationView) SetFlags(flags unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFlags:"), flags)
}/* debug [instance_methods/method]: SetFlags */


// Sets the requested-right string to use with the default authorization rights set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/setString(_:)
func (s_ SFAuthorizationView) SetString(authorizationString unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setString:"), authorizationString)
}/* debug [instance_methods/method]: SetString */


// Manually updates the authorization view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationView/updateStatus(_:)
func (s_ SFAuthorizationView) UpdateStatus(inSender objc.IObject) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("updateStatus:"), inSender)
	return rv
}/* debug [instance_methods/method]: UpdateStatus */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFAuthorizationView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFAuthorizationView */



