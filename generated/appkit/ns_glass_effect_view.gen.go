// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSGlassEffectView */


/* debug [class_header]: Header for NSGlassEffectView */
// The class instance for the [GlassEffectView] class.
var (
	GlassEffectViewClass     _GlassEffectViewClass
	GlassEffectViewClassOnce sync.Once
)

func getGlassEffectViewClass() _GlassEffectViewClass {
	GlassEffectViewClassOnce.Do(func() {
		GlassEffectViewClass = _GlassEffectViewClass{objc.GetClass("NSGlassEffectView")}
	})
	return GlassEffectViewClass
}

type _GlassEffectViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GlassEffectView */
// An interface definition for the [GlassEffectView] class.
type IGlassEffectView interface {
	IView
	
/* debug [class_interface_properties]: Properties for GlassEffectView */
	// properties:
	ContentView() IView
	SetContentView(value IView)
	CornerRadius() float64
	SetCornerRadius(value float64)
	Style() GlassEffectViewStyle
	SetStyle(value GlassEffectViewStyle)
	TintColor() IColor
	SetTintColor(value IColor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GlassEffectView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GlassEffectView */
// Alloc allocates a new instance without initialization.
func (gc _GlassEffectViewClass) Alloc() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GlassEffectViewClass) New() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlassEffectView) Init() GlassEffectView {
	rv := objc.Send[GlassEffectView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlassEffectView) Autorelease() GlassEffectView {
	rv := objc.Send[GlassEffectView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlassEffectView creates a new GlassEffectView instance.
func NewGlassEffectView() GlassEffectView {
	return getGlassEffectViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GlassEffectView */
// A view that embeds its content view in a dynamic glass effect.


// A view that embeds its content view in a dynamic glass effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView
type GlassEffectView struct {
	View
}

// GlassEffectViewFrom constructs a [GlassEffectView] from an unsafe.Pointer.
//
// A view that embeds its content view in a dynamic glass effect.
func GlassEffectViewFrom(ptr unsafe.Pointer) GlassEffectView {
	return GlassEffectView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GlassEffectView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GlassEffectView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GlassEffectView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GlassEffectView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GlassEffectView */

// The view to embed in glass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) ContentView() IView {
	rv := objc.Send[View](g_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The view to embed in glass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) SetContentView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */


// The amount of curvature for all corners of the glass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) CornerRadius() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("cornerRadius"))
	return rv
}/* debug [instance_properties/getter]: cornerRadius */


// The amount of curvature for all corners of the glass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) SetCornerRadius(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCornerRadius:"), value)
}/* debug [instance_properties/setter]: cornerRadius */


// The style of glass this view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) Style() GlassEffectViewStyle {
	rv := objc.Send[GlassEffectViewStyle](g_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// The style of glass this view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) SetStyle(value GlassEffectViewStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// The color the glass effect view uses to tint the background and glass effect toward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) TintColor() IColor {
	rv := objc.Send[Color](g_.ID, objc.Sel("tintColor"))
	return rv
}/* debug [instance_properties/getter]: tintColor */


// The color the glass effect view uses to tint the background and glass effect toward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) SetTintColor(value IColor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTintColor:"), value)
}/* debug [instance_properties/setter]: tintColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGlassEffectView */



