// Code generated from Apple documentation for LocalAuthenticationEmbeddedUI. DO NOT EDIT.

package localauthenticationembeddedui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreimage"
)

/* debug [class.gen.go]: Generating class LAAuthenticationView */


/* debug [class_header]: Header for LAAuthenticationView */
// The class instance for the [AuthenticationView] class.
var (
	AuthenticationViewClass     _AuthenticationViewClass
	AuthenticationViewClassOnce sync.Once
)

func getAuthenticationViewClass() _AuthenticationViewClass {
	AuthenticationViewClassOnce.Do(func() {
		AuthenticationViewClass = _AuthenticationViewClass{objc.GetClass("LAAuthenticationView")}
	})
	return AuthenticationViewClass
}

type _AuthenticationViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthenticationView */
// An interface definition for the [AuthenticationView] class.
type IAuthenticationView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for AuthenticationView */
	// properties:
	Context() coreimage.Context
	ControlSize() ControlSize /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthenticationView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthenticationView */
// Alloc allocates a new instance without initialization.
func (ac _AuthenticationViewClass) Alloc() AuthenticationView {
	rv := objc.Send[AuthenticationView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthenticationViewClass) New() AuthenticationView {
	rv := objc.Send[AuthenticationView](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthenticationView) Init() AuthenticationView {
	rv := objc.Send[AuthenticationView](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthenticationView) Autorelease() AuthenticationView {
	rv := objc.Send[AuthenticationView](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthenticationView creates a new AuthenticationView instance.
func NewAuthenticationView() AuthenticationView {
	return getAuthenticationViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthenticationView */
// A graphical representation of the state of biometric authentication.
//
// In the view that you use to manage authentication, add a local authentication view as a subview and provide it with an instance. For example, you can do this in the doc://com.apple.documentation/documentation/appkit/nsviewcontroller/1434405-loadview method of your view controller: When the view appears, call the context’s method to initiate the authentication: The local authentication view displays an icon that depends on the type of authentication you request, and the types of authentication that the system supports. For example, for a device that supports Touch ID, if you request the policy, like in the example above, the view displays the familiar finger print icon: In the case above, if the user has a connected Apple Watch, that authentication mechanism works as well. If you limit the authentication to the policy, the icon shows an Apple Watch in profile: You can include other content around this icon that suits your app. The system also displays a message on the Touch Bar or on the user’s Apple Watch, if appropriate. When the evaluation succeeds, the icon transitions into a checkmark: If you call the evaluation without first attaching it to a local authentication view, the system shows a standard authentication alert instead.


// A graphical representation of the state of biometric authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI/LAAuthenticationView
type AuthenticationView struct {
	appkit.View
}

// AuthenticationViewFrom constructs a [AuthenticationView] from an unsafe.Pointer.
//
// A graphical representation of the state of biometric authentication.
func AuthenticationViewFrom(ptr unsafe.Pointer) AuthenticationView {
	return AuthenticationView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthenticationView */

// Creates a new authentication icon that reflects the current authentication state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI/LAAuthenticationView/init(context:)
func NewAuthenticationViewWithContext(context coreimage.Context) AuthenticationView {
	instance := getAuthenticationViewClass().Alloc()
	rv := objc.Send[AuthenticationView](instance.ID, objc.Sel("initWithContext:"), context)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthenticationViewWithContext */


// Creates a new authentication icon that reflects the current authentication state, using a specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI/LAAuthenticationView/init(context:controlSize:)
func NewAuthenticationViewWithContextControlSize(context coreimage.Context, controlSize ControlSize /* not a class type */) AuthenticationView {
	instance := getAuthenticationViewClass().Alloc()
	rv := objc.Send[AuthenticationView](instance.ID, objc.Sel("initWithContext:controlSize:"), context, controlSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthenticationViewWithContextControlSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthenticationView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthenticationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthenticationView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthenticationView */

// The local authentication context associated with the authentication view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI/LAAuthenticationView/context
func (a_ AuthenticationView) Context() coreimage.Context {
	rv := objc.Send[coreimage.Context](a_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// The size of the local authentication view user interface element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI/LAAuthenticationView/controlSize
func (a_ AuthenticationView) ControlSize() ControlSize /* not a class type */ {
	rv := objc.Send[ControlSize](a_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAAuthenticationView */


