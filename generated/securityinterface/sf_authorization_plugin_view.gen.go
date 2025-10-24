// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFAuthorizationPluginView */


/* debug [class_header]: Header for SFAuthorizationPluginView */
// The class instance for the [SFAuthorizationPluginView] class.
var (
	SFAuthorizationPluginViewClass     _SFAuthorizationPluginViewClass
	SFAuthorizationPluginViewClassOnce sync.Once
)

func getSFAuthorizationPluginViewClass() _SFAuthorizationPluginViewClass {
	SFAuthorizationPluginViewClassOnce.Do(func() {
		SFAuthorizationPluginViewClass = _SFAuthorizationPluginViewClass{objc.GetClass("SFAuthorizationPluginView")}
	})
	return SFAuthorizationPluginViewClass
}

type _SFAuthorizationPluginViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFAuthorizationPluginView */
// An interface definition for the [SFAuthorizationPluginView] class.
type ISFAuthorizationPluginView interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFAuthorizationPluginView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFAuthorizationPluginView */
	// methods:
	ButtonPressed(inButtonType unsafe.Pointer)
	Callbacks() unsafe.Pointer
	DidActivate()
	DidDeactivate()
	DisplayView()
	EngineRef() unsafe.Pointer
	FirstKeyView() appkit.View
	FirstResponder() appkit.Responder
	LastError() objc.IObject /* cross-framework: Error */
	LastKeyView() appkit.View
	SetButtonEnabled(inButtonType unsafe.Pointer, inEnabled bool)
	SetEnabled(inEnabled bool)
	UpdateView()
	ViewForType(inType unsafe.Pointer) appkit.View
	WillActivateWithUser(inUserInformation objc.IObject /* cross-framework: NSDictionary */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFAuthorizationPluginView */
// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationPluginViewClass) Alloc() SFAuthorizationPluginView {
	rv := objc.Send[SFAuthorizationPluginView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFAuthorizationPluginViewClass) New() SFAuthorizationPluginView {
	rv := objc.Send[SFAuthorizationPluginView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAuthorizationPluginView) Init() SFAuthorizationPluginView {
	rv := objc.Send[SFAuthorizationPluginView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAuthorizationPluginView) Autorelease() SFAuthorizationPluginView {
	rv := objc.Send[SFAuthorizationPluginView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAuthorizationPluginView creates a new SFAuthorizationPluginView instance.
func NewSFAuthorizationPluginView() SFAuthorizationPluginView {
	return getSFAuthorizationPluginViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFAuthorizationPluginView */
// Allows authorization plug-in developers to create a custom view their plug-in can display.
//
// If you’re developing an authorization plug-in, you can subclass the class to create views that provide a custom user interface for your plug-in. By subclassing the class, you avoid changing or duplicating the Apple-provided authentication or login window dialogs to display your custom view. To instantiate your subclass, you need the callbacks structure containing entry points to the Security Server that you receive in your plug-in’s function and the authorization engine handle you receive in your plug-in’s function. Your custom subclass of must override the following methods:


// Allows authorization plug-in developers to create a custom view their plug-in can display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView
type SFAuthorizationPluginView struct {
	objectivec.Object
}

// SFAuthorizationPluginViewFrom constructs a [SFAuthorizationPluginView] from an unsafe.Pointer.
//
// Allows authorization plug-in developers to create a custom view their plug-in can display.
func SFAuthorizationPluginViewFrom(ptr unsafe.Pointer) SFAuthorizationPluginView {
	return SFAuthorizationPluginView{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFAuthorizationPluginView */

// Initializes a new authorization plug-in view with the specified callbacks and authorization engine handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/init(callbacks:andEngineRef:)
func NewSFAuthorizationPluginViewWithCallbacksAndEngineRef(callbacks unsafe.Pointer, engineRef unsafe.Pointer) SFAuthorizationPluginView {
	instance := getSFAuthorizationPluginViewClass().Alloc()
	rv := objc.Send[SFAuthorizationPluginView](instance.ID, objc.Sel("initWithCallbacks:andEngineRef:"), callbacks, engineRef)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFAuthorizationPluginViewWithCallbacksAndEngineRef */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFAuthorizationPluginView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFAuthorizationPluginView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFAuthorizationPluginView */

// Tells the authorization plug-in that the user pressed a button in the custom view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/buttonPressed(_:)
func (s_ SFAuthorizationPluginView) ButtonPressed(inButtonType unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("buttonPressed:"), inButtonType)
}/* debug [instance_methods/method]: ButtonPressed */


// Returns the authorization callbacks structure with which this instance was initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/callbacks()
func (s_ SFAuthorizationPluginView) Callbacks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("callbacks"))
	return rv
}/* debug [instance_methods/method]: Callbacks */


// Tells the authorization plug-in when its user interface has become active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/didActivate()
func (s_ SFAuthorizationPluginView) DidActivate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didActivate"))
}/* debug [instance_methods/method]: DidActivate */


// Tells the authorization plug-in that its user interface has been deactivated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/didDeactivate()
func (s_ SFAuthorizationPluginView) DidDeactivate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didDeactivate"))
}/* debug [instance_methods/method]: DidDeactivate */


// Displays the user interface provided by the authorization plug-in view subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/display()
func (s_ SFAuthorizationPluginView) DisplayView() {
	objc.Send[objc.ID](s_.ID, objc.Sel("displayView"))
}/* debug [instance_methods/method]: DisplayView */


// Returns the authorization engine handle with which this instance was initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/engineRef()
func (s_ SFAuthorizationPluginView) EngineRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("engineRef"))
	return rv
}/* debug [instance_methods/method]: EngineRef */


// Returns the first view in the keyboard loop of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/firstKeyView()
func (s_ SFAuthorizationPluginView) FirstKeyView() appkit.View {
	rv := objc.Send[appkit.View](s_.ID, objc.Sel("firstKeyView"))
	return rv
}/* debug [instance_methods/method]: FirstKeyView */


// Returns the view that should get focus for keyboard events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/firstResponder()
func (s_ SFAuthorizationPluginView) FirstResponder() appkit.Responder {
	rv := objc.Send[appkit.Responder](s_.ID, objc.Sel("firstResponder"))
	return rv
}/* debug [instance_methods/method]: FirstResponder */


// Returns the last error that occurred during evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/lastError()
func (s_ SFAuthorizationPluginView) LastError() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](s_.ID, objc.Sel("lastError"))
	return rv
}/* debug [instance_methods/method]: LastError */


// Returns the last view in the keyboard loop of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/lastKeyView()
func (s_ SFAuthorizationPluginView) LastKeyView() appkit.View {
	rv := objc.Send[appkit.View](s_.ID, objc.Sel("lastKeyView"))
	return rv
}/* debug [instance_methods/method]: LastKeyView */


// Enables or disables a button in the authorization plug-in’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/setButton(_:enabled:)
func (s_ SFAuthorizationPluginView) SetButtonEnabled(inButtonType unsafe.Pointer, inEnabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButton:enabled:"), inButtonType, inEnabled)
}/* debug [instance_methods/method]: SetButtonEnabled */


// Enables or disables the controls in the authorization plug-in’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/setEnabled(_:)
func (s_ SFAuthorizationPluginView) SetEnabled(inEnabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), inEnabled)
}/* debug [instance_methods/method]: SetEnabled */


// Tells the authorization plug-in to get and display the appropriate view in the authorization plug-in’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/update()
func (s_ SFAuthorizationPluginView) UpdateView() {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateView"))
}/* debug [instance_methods/method]: UpdateView */


// Returns the appropriate view object for the specified view type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/view(for:)
func (s_ SFAuthorizationPluginView) ViewForType(inType unsafe.Pointer) appkit.View {
	rv := objc.Send[appkit.View](s_.ID, objc.Sel("viewForType:"), inType)
	return rv
}/* debug [instance_methods/method]: ViewForType */


// Tells the authorization plug-in that its user interface is about to be made active by the Apple-provided Security Agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/willActivate(withUser:)
func (s_ SFAuthorizationPluginView) WillActivateWithUser(inUserInformation objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("willActivateWithUser:"), inUserInformation)
}/* debug [instance_methods/method]: WillActivateWithUser */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFAuthorizationPluginView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFAuthorizationPluginView */


