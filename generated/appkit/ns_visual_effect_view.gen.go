// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSVisualEffectView */


/* debug [class_header]: Header for NSVisualEffectView */
// The class instance for the [VisualEffectView] class.
var (
	VisualEffectViewClass     _VisualEffectViewClass
	VisualEffectViewClassOnce sync.Once
)

func getVisualEffectViewClass() _VisualEffectViewClass {
	VisualEffectViewClassOnce.Do(func() {
		VisualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}
	})
	return VisualEffectViewClass
}

type _VisualEffectViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VisualEffectView */
// An interface definition for the [VisualEffectView] class.
type IVisualEffectView interface {
	IView
	
/* debug [class_interface_properties]: Properties for VisualEffectView */
	// properties:
	BlendingMode() VisualEffectBlendingMode
	SetBlendingMode(value VisualEffectBlendingMode)
	InteriorBackgroundStyle() BackgroundStyle
	Emphasized() bool
	SetEmphasized(value bool)
	MaskImage() IImage
	SetMaskImage(value IImage)
	Material() VisualEffectMaterial
	SetMaterial(value VisualEffectMaterial)
	State() VisualEffectState
	SetState(value VisualEffectState)
	AllowsVibrancy() bool
	SetAllowsVibrancy(value bool)
	IsEmphasized() bool
	SetIsEmphasized(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VisualEffectView */
	// methods:
	ViewDidMoveToWindow()
	ViewWillMoveToWindow(newWindow IWindow)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VisualEffectView */
// Alloc allocates a new instance without initialization.
func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VisualEffectView) Init() VisualEffectView {
	rv := objc.Send[VisualEffectView](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VisualEffectView) Autorelease() VisualEffectView {
	rv := objc.Send[VisualEffectView](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisualEffectView creates a new VisualEffectView instance.
func NewVisualEffectView() VisualEffectView {
	return getVisualEffectViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VisualEffectView */
// A view that adds translucency and vibrancy effects to the views in your interface.
//
// Use visual effect views primarily as background views for your app’s content. A visual effect view makes your foreground content more prominent by employing the following effects: and the blurring of background content adds depth to your interface. is a subtle blending of foreground and background colors to increase the contrast and make the foreground content stand out visually. The material and blending mode you assign determines the exact appearance of the visual effect. Not all materials support transparency, and materials apply vibrancy in different ways. The appearance and behavior of materials can also change based on system settings, so always pick a material based on its intended use. For example, use the material when your view serves as the background of your window’s sidebar. Don’t select materials based on the apparent colors they impart on your interface. AppKit creates visual effect views automatically for window titlebars, popovers, and source list table views. You don’t need to add visual effect views to those elements of your interface.


// A view that adds translucency and vibrancy effects to the views in your interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView
type VisualEffectView struct {
	View
}

// VisualEffectViewFrom constructs a [VisualEffectView] from an unsafe.Pointer.
//
// A view that adds translucency and vibrancy effects to the views in your interface.
func VisualEffectViewFrom(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VisualEffectView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VisualEffectView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VisualEffectView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VisualEffectView */

// Notifies the view that it moved to a new window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}/* debug [instance_methods/method]: ViewDidMoveToWindow */


// Notifies the view immediately before it moves to a new window (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow IWindow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}/* debug [instance_methods/method]: ViewWillMoveToWindow */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VisualEffectView */

// A value indicating how the view’s contents blend with the surrounding content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) BlendingMode() VisualEffectBlendingMode {
	rv := objc.Send[VisualEffectBlendingMode](v_.ID, objc.Sel("blendingMode"))
	return rv
}/* debug [instance_properties/getter]: blendingMode */


// A value indicating how the view’s contents blend with the surrounding content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) SetBlendingMode(value VisualEffectBlendingMode) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBlendingMode:"), value)
}/* debug [instance_properties/setter]: blendingMode */


// The view’s interior background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/interiorBackgroundStyle
func (v_ VisualEffectView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](v_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: interiorBackgroundStyle */


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) Emphasized() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("emphasized"))
	return rv
}/* debug [instance_properties/getter]: emphasized */


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) SetEmphasized(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEmphasized:"), value)
}/* debug [instance_properties/setter]: emphasized */


// An image whose alpha channel masks the visual effect view’s material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) MaskImage() IImage {
	rv := objc.Send[Image](v_.ID, objc.Sel("maskImage"))
	return rv
}/* debug [instance_properties/getter]: maskImage */


// An image whose alpha channel masks the visual effect view’s material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) SetMaskImage(value IImage) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaskImage:"), value)
}/* debug [instance_properties/setter]: maskImage */


// The material shown by the visual effect view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) Material() VisualEffectMaterial {
	rv := objc.Send[VisualEffectMaterial](v_.ID, objc.Sel("material"))
	return rv
}/* debug [instance_properties/getter]: material */


// The material shown by the visual effect view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) SetMaterial(value VisualEffectMaterial) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaterial:"), value)
}/* debug [instance_properties/setter]: material */


// A value that indicates whether a view has a visual effect applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) State() VisualEffectState {
	rv := objc.Send[VisualEffectState](v_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A value that indicates whether a view has a visual effect applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) SetState(value VisualEffectState) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/allowsvibrancy
func (v_ VisualEffectView) AllowsVibrancy() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("allowsVibrancy"))
	return rv
}/* debug [instance_properties/getter]: allowsVibrancy */


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/allowsvibrancy
func (v_ VisualEffectView) SetAllowsVibrancy(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAllowsVibrancy:"), value)
}/* debug [instance_properties/setter]: allowsVibrancy */


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/isemphasized
func (v_ VisualEffectView) IsEmphasized() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isEmphasized"))
	return rv
}/* debug [instance_properties/getter]: isEmphasized */


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/isemphasized
func (v_ VisualEffectView) SetIsEmphasized(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsEmphasized:"), value)
}/* debug [instance_properties/setter]: isEmphasized */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSVisualEffectView */



