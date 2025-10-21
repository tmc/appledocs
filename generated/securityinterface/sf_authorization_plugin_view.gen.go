// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [SFAuthorizationPluginView] class.
type ISFAuthorizationPluginView interface {
	objectivec.IObject
	ButtonPressed(inButtonType unsafe.Pointer)
	Callbacks() unsafe.Pointer
	DidActivate()
	DidDeactivate()
	DisplayView()
	EngineRef() unsafe.Pointer
	FirstKeyView() unsafe.Pointer
	FirstResponder() unsafe.Pointer
	LastError() unsafe.Pointer
	LastKeyView() unsafe.Pointer
	SetButtonEnabled(inButtonType unsafe.Pointer, inEnabled bool)
	SetEnabled(inEnabled bool)
	UpdateView()
	ViewForType(inType unsafe.Pointer) unsafe.Pointer
	WillActivateWithUser(inUserInformation objc.ID)
}

// Allows authorization plug-in developers to create a custom view their plug-in can display.
//
// If you’re developing an authorization plug-in, you can subclass the class to create views that provide a custom user interface for your plug-in. By subclassing the class, you avoid changing or duplicating the Apple-provided authentication or login window dialogs to display your custom view. To instantiate your subclass, you need the callbacks structure containing entry points to the Security Server that you receive in your plug-in’s function and the authorization engine handle you receive in your plug-in’s function. Your custom subclass of must override the following methods:
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFAuthorizationPluginViewClass) Alloc() SFAuthorizationPluginView {
	rv := objc.Send[SFAuthorizationPluginView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes a new authorization plug-in view with the specified callbacks and authorization engine handle.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/init(callbacks:andEngineRef:)
func NewSFAuthorizationPluginViewWithCallbacksAndEngineRef(callbacks unsafe.Pointer, engineRef unsafe.Pointer) SFAuthorizationPluginView {
	instance := getSFAuthorizationPluginViewClass().Alloc()
	rv := objc.Send[SFAuthorizationPluginView](instance.ID, objc.Sel("initWithCallbacks:andEngineRef:"), callbacks, engineRef)
	rv.Autorelease()
	return rv
}


// Tells the authorization plug-in that the user pressed a button in the custom view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/buttonPressed(_:)
func (s_ SFAuthorizationPluginView) ButtonPressed(inButtonType unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("buttonPressed:"), inButtonType)
}

// Returns the authorization callbacks structure with which this instance was initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/callbacks()
func (s_ SFAuthorizationPluginView) Callbacks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("callbacks"))
	return rv
}

// Tells the authorization plug-in when its user interface has become active.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/didActivate()
func (s_ SFAuthorizationPluginView) DidActivate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didActivate"))
}

// Tells the authorization plug-in that its user interface has been deactivated.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/didDeactivate()
func (s_ SFAuthorizationPluginView) DidDeactivate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didDeactivate"))
}

// Displays the user interface provided by the authorization plug-in view subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/display()
func (s_ SFAuthorizationPluginView) DisplayView() {
	objc.Send[objc.ID](s_.ID, objc.Sel("displayView"))
}

// Returns the authorization engine handle with which this instance was initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/engineRef()
func (s_ SFAuthorizationPluginView) EngineRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("engineRef"))
	return rv
}

// Returns the first view in the keyboard loop of the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/firstKeyView()
func (s_ SFAuthorizationPluginView) FirstKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("firstKeyView"))
	return rv
}

// Returns the view that should get focus for keyboard events.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/firstResponder()
func (s_ SFAuthorizationPluginView) FirstResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("firstResponder"))
	return rv
}

// Returns the last error that occurred during evaluation.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/lastError()
func (s_ SFAuthorizationPluginView) LastError() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lastError"))
	return rv
}

// Returns the last view in the keyboard loop of the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/lastKeyView()
func (s_ SFAuthorizationPluginView) LastKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lastKeyView"))
	return rv
}

// Enables or disables a button in the authorization plug-in’s user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/setButton(_:enabled:)
func (s_ SFAuthorizationPluginView) SetButtonEnabled(inButtonType unsafe.Pointer, inEnabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButton:enabled:"), inButtonType, inEnabled)
}

// Enables or disables the controls in the authorization plug-in’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/setEnabled(_:)
func (s_ SFAuthorizationPluginView) SetEnabled(inEnabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), inEnabled)
}

// Tells the authorization plug-in to get and display the appropriate view in the authorization plug-in’s user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/update()
func (s_ SFAuthorizationPluginView) UpdateView() {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateView"))
}

// Returns the appropriate view object for the specified view type.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/view(for:)
func (s_ SFAuthorizationPluginView) ViewForType(inType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("viewForType:"), inType)
	return rv
}

// Tells the authorization plug-in that its user interface is about to be made active by the Apple-provided Security Agent.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFAuthorizationPluginView/willActivate(withUser:)
func (s_ SFAuthorizationPluginView) WillActivateWithUser(inUserInformation objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("willActivateWithUser:"), inUserInformation)
}


