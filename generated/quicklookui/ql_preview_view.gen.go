// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class QLPreviewView */


/* debug [class_header]: Header for QLPreviewView */
// The class instance for the [PreviewView] class.
var (
	PreviewViewClass     _PreviewViewClass
	PreviewViewClassOnce sync.Once
)

func getPreviewViewClass() _PreviewViewClass {
	PreviewViewClassOnce.Do(func() {
		PreviewViewClass = _PreviewViewClass{objc.GetClass("QLPreviewView")}
	})
	return PreviewViewClass
}

type _PreviewViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewView */
// An interface definition for the [PreviewView] class.
type IPreviewView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for PreviewView */
	// properties:
	Autostarts() bool
	SetAutostarts(value bool)
	DisplayState() objc.ID
	SetDisplayState(value objc.ID)
	PreviewItem() unsafe.Pointer
	SetPreviewItem(value unsafe.Pointer)
	ShouldCloseWithWindow() bool
	SetShouldCloseWithWindow(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewView */
	// methods:
	Close()
	RefreshPreviewItem()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewView */
// Alloc allocates a new instance without initialization.
func (pc _PreviewViewClass) Alloc() PreviewView {
	rv := objc.Send[PreviewView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewViewClass) New() PreviewView {
	rv := objc.Send[PreviewView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewView) Init() PreviewView {
	rv := objc.Send[PreviewView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewView) Autorelease() PreviewView {
	rv := objc.Send[PreviewView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewView creates a new PreviewView instance.
func NewPreviewView() PreviewView {
	return getPreviewViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewView */
// A Quick Look preview of an item that you can embed into your view hierarchy.


// A Quick Look preview of an item that you can embed into your view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView
type PreviewView struct {
	appkit.View
}

// PreviewViewFrom constructs a [PreviewView] from an unsafe.Pointer.
//
// A Quick Look preview of an item that you can embed into your view hierarchy.
func PreviewViewFrom(ptr unsafe.Pointer) PreviewView {
	return PreviewView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewView */

// Creates a preview view with the provided frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/init(frame:)
func NewPreviewViewWithFrame(frame Rect /* not a class type */) PreviewView {
	instance := getPreviewViewClass().Alloc()
	rv := objc.Send[PreviewView](instance.ID, objc.Sel("initWithFrame:"), frame)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewViewWithFrame */


// Creates a preview view with the provided frame and style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/init(frame:style:)
func NewPreviewViewWithFrameStyle(frame Rect /* not a class type */, style PreviewViewStyle) PreviewView {
	instance := getPreviewViewClass().Alloc()
	rv := objc.Send[PreviewView](instance.ID, objc.Sel("initWithFrame:style:"), frame, style)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewViewWithFrameStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewView */

// Closes the view, releasing the current preview item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/close()
func (p_ PreviewView) Close() {
	objc.Send[objc.ID](p_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Updates the preview to display the currently previewed item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/refreshPreviewItem()
func (p_ PreviewView) RefreshPreviewItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("refreshPreviewItem"))
}/* debug [instance_methods/method]: RefreshPreviewItem */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewView */

// A Boolean value that determines whether the preview starts automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/autostarts
func (p_ PreviewView) Autostarts() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autostarts"))
	return rv
}/* debug [instance_properties/getter]: autostarts */


// A Boolean value that determines whether the preview starts automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/autostarts
func (p_ PreviewView) SetAutostarts(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutostarts:"), value)
}/* debug [instance_properties/setter]: autostarts */


// The current display state of the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/displayState
func (p_ PreviewView) DisplayState() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("displayState"))
	return rv
}/* debug [instance_properties/getter]: displayState */


// The current display state of the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/displayState
func (p_ PreviewView) SetDisplayState(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayState:"), value)
}/* debug [instance_properties/setter]: displayState */


// The item to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/previewItem
func (p_ PreviewView) PreviewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("previewItem"))
	return rv
}/* debug [instance_properties/getter]: previewItem */


// The item to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/previewItem
func (p_ PreviewView) SetPreviewItem(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreviewItem:"), value)
}/* debug [instance_properties/setter]: previewItem */


// A Boolean value that determines whether the preview should close when its window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/shouldCloseWithWindow
func (p_ PreviewView) ShouldCloseWithWindow() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldCloseWithWindow"))
	return rv
}/* debug [instance_properties/getter]: shouldCloseWithWindow */


// A Boolean value that determines whether the preview should close when its window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/shouldCloseWithWindow
func (p_ PreviewView) SetShouldCloseWithWindow(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldCloseWithWindow:"), value)
}/* debug [instance_properties/setter]: shouldCloseWithWindow */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewView */


