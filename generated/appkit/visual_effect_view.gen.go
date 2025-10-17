
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [VisualEffectView] class.
var VisualEffectViewClass _VisualEffectViewClass

func init() {
	VisualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}
}

type _VisualEffectViewClass struct {
	objc.Class
}

// An interface definition for the [VisualEffectView] class.
type IVisualEffectView interface {
	ID() objc.ID
	ViewDidMoveToWindow()
	ViewWillMoveToWindow(newWindow unsafe.Pointer)
}

type VisualEffectView struct {
	id objc.ID
}

func VisualEffectViewFrom(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ VisualEffectView) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewVisualEffectView creates and returns a new initialized instance.
func NewVisualEffectView() VisualEffectView {
	return VisualEffectViewClass.New()
}

// Init initializes the instance.
func (v_ VisualEffectView) Init() VisualEffectView {
	rv := objc.Send[VisualEffectView](v_.ID(), selInit)
	return rv
}
// Notifies the view that it moved to a new window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidMoveToWindow"))
}
// Notifies the view immediately before it moves to a new window (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewWillMoveToWindow:"), newWindow)
}
// A value indicating how the view’s contents blend with the surrounding content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) BlendingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("blendingMode"))
	return rv
}
// SetBlendingMode sets the value of the blendingMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/blendingMode-swift.property
func (v_ VisualEffectView) SetBlendingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBlendingMode:"), value)
}
// The view’s interior background style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/interiorBackgroundStyle
func (v_ VisualEffectView) InteriorBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("interiorBackgroundStyle"))
	return rv
}
// A Boolean value indicating whether to emphasize the look of the material. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) Emphasized() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("emphasized"))
	return rv
}
// SetEmphasized sets the value of the emphasized property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/isEmphasized
func (v_ VisualEffectView) SetEmphasized(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setEmphasized:"), value)
}
// An image whose alpha channel masks the visual effect view’s material. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) MaskImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("maskImage"))
	return rv
}
// SetMaskImage sets the value of the maskImage property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/maskImage
func (v_ VisualEffectView) SetMaskImage(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setMaskImage:"), value)
}
// The material shown by the visual effect view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) Material() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("material"))
	return rv
}
// SetMaterial sets the value of the material property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/material-swift.property
func (v_ VisualEffectView) SetMaterial(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setMaterial:"), value)
}
// A value that indicates whether a view has a visual effect applied. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("state"))
	return rv
}
// SetState sets the value of the state property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/state-swift.property
func (v_ VisualEffectView) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setState:"), value)
}
