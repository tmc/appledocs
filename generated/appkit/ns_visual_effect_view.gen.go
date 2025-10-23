// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VisualEffectView] class.
type IVisualEffectView interface {
	IView
	BlendingMode() NSVisualEffectBlendingMode
	SetBlendingMode(value NSVisualEffectBlendingMode)
	InteriorBackgroundStyle() NSBackgroundStyle
	Emphasized() bool
	SetEmphasized(value bool)
	MaskImage() IImage
	SetMaskImage(value IImage)
	Material() NSVisualEffectMaterial
	SetMaterial(value NSVisualEffectMaterial)
	State() NSVisualEffectState
	SetState(value NSVisualEffectState)
	AllowsVibrancy() bool
	SetAllowsVibrancy(value bool)
	IsEmphasized() bool
	SetIsEmphasized(value bool)
	ViewDidMoveToWindow()
	ViewWillMoveToWindow(newWindow IWindow)
}

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

// Alloc allocates a new instance without initialization.
func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Notifies the view that it moved to a new window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}


// Notifies the view immediately before it moves to a new window (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow IWindow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}


// A value indicating how the view’s contents blend with the surrounding content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) BlendingMode() NSVisualEffectBlendingMode {
	rv := objc.Send[NSVisualEffectBlendingMode](v_.ID, objc.Sel("blendingMode"))
	return rv
}


// A value indicating how the view’s contents blend with the surrounding content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) SetBlendingMode(value NSVisualEffectBlendingMode) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBlendingMode:"), value)
}


// The view’s interior background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/interiorBackgroundStyle
func (v_ VisualEffectView) InteriorBackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](v_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) Emphasized() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("emphasized"))
	return rv
}


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) SetEmphasized(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEmphasized:"), value)
}


// An image whose alpha channel masks the visual effect view’s material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) MaskImage() IImage {
	rv := objc.Send[Image](v_.ID, objc.Sel("maskImage"))
	return rv
}


// An image whose alpha channel masks the visual effect view’s material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) SetMaskImage(value IImage) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaskImage:"), value)
}


// The material shown by the visual effect view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) Material() NSVisualEffectMaterial {
	rv := objc.Send[NSVisualEffectMaterial](v_.ID, objc.Sel("material"))
	return rv
}


// The material shown by the visual effect view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) SetMaterial(value NSVisualEffectMaterial) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaterial:"), value)
}


// A value that indicates whether a view has a visual effect applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) State() NSVisualEffectState {
	rv := objc.Send[NSVisualEffectState](v_.ID, objc.Sel("state"))
	return rv
}


// A value that indicates whether a view has a visual effect applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) SetState(value NSVisualEffectState) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setState:"), value)
}


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/allowsvibrancy
func (v_ VisualEffectView) AllowsVibrancy() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("allowsVibrancy"))
	return rv
}


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/allowsvibrancy
func (v_ VisualEffectView) SetAllowsVibrancy(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAllowsVibrancy:"), value)
}


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/isemphasized
func (v_ VisualEffectView) IsEmphasized() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isEmphasized"))
	return rv
}


// A Boolean value indicating whether to emphasize the look of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/isemphasized
func (v_ VisualEffectView) SetIsEmphasized(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsEmphasized:"), value)
}



