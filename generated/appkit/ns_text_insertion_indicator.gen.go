// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextInsertionIndicator */


/* debug [class_header]: Header for NSTextInsertionIndicator */
// The class instance for the [TextInsertionIndicator] class.
var (
	TextInsertionIndicatorClass     _TextInsertionIndicatorClass
	TextInsertionIndicatorClassOnce sync.Once
)

func getTextInsertionIndicatorClass() _TextInsertionIndicatorClass {
	TextInsertionIndicatorClassOnce.Do(func() {
		TextInsertionIndicatorClass = _TextInsertionIndicatorClass{objc.GetClass("NSTextInsertionIndicator")}
	})
	return TextInsertionIndicatorClass
}

type _TextInsertionIndicatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextInsertionIndicator */
// An interface definition for the [TextInsertionIndicator] class.
type ITextInsertionIndicator interface {
	IView
	
/* debug [class_interface_properties]: Properties for TextInsertionIndicator */
	// properties:
	DisplayMode() TextInsertionIndicatorDisplayMode
	SetDisplayMode(value TextInsertionIndicatorDisplayMode)
	AutomaticModeOptions() objectivec.IObject
	SetAutomaticModeOptions(value objectivec.IObject)
	Color() IColor
	SetColor(value IColor)
	EffectsViewInserter() objectivec.IObject
	SetEffectsViewInserter(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextInsertionIndicator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextInsertionIndicator */
// Alloc allocates a new instance without initialization.
func (tc _TextInsertionIndicatorClass) Alloc() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextInsertionIndicatorClass) New() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextInsertionIndicator) Init() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextInsertionIndicator) Autorelease() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextInsertionIndicator creates a new TextInsertionIndicator instance.
func NewTextInsertionIndicator() TextInsertionIndicator {
	return getTextInsertionIndicatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextInsertionIndicator */
// A view that represents the insertion indicator in text.
//
// and both use to display the insertion indicator. You can use this indicator if you have your own text engine or need to display an indicator elsewhere. To use the indicator, instantiate an , then add the view to your view hierarchy. Set the indicator view’s frame to where you want to display a text insertion indicator. The indicator has the same height as the indicator view’s frame, and centers horizontally within the indicator view’s frame. The specifies whether the indicator hides, remains visible, or blinks (automatic). When set to , the indicator stops blinking when you set the frame. The indicator starts blinking when the frame doesn’t change for a period of time. When the user dictates, the indicator displays a trailing glow when it is moved. Set the to when your custom view becomes the first responder. When your custom view resigns first responder, set the to to indicate that key events aren’t sent to your view. By default the indicator’s color is . You can set a different color.


// A view that represents the insertion indicator in text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator
type TextInsertionIndicator struct {
	View
}

// TextInsertionIndicatorFrom constructs a [TextInsertionIndicator] from an unsafe.Pointer.
//
// A view that represents the insertion indicator in text.
func TextInsertionIndicatorFrom(ptr unsafe.Pointer) TextInsertionIndicator {
	return TextInsertionIndicator{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextInsertionIndicator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextInsertionIndicator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextInsertionIndicator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextInsertionIndicator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextInsertionIndicator */

// A value that describes the display mode of an indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) DisplayMode() TextInsertionIndicatorDisplayMode {
	rv := objc.Send[TextInsertionIndicatorDisplayMode](t_.ID, objc.Sel("displayMode"))
	return rv
}/* debug [instance_properties/getter]: displayMode */


// A value that describes the display mode of an indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) SetDisplayMode(value TextInsertionIndicatorDisplayMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplayMode:"), value)
}/* debug [instance_properties/setter]: displayMode */


// Options that affect the automatic display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/automaticmodeoptions-swift.property
func (t_ TextInsertionIndicator) AutomaticModeOptions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("automaticModeOptions"))
	return rv
}/* debug [instance_properties/getter]: automaticModeOptions */


// Options that affect the automatic display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/automaticmodeoptions-swift.property
func (t_ TextInsertionIndicator) SetAutomaticModeOptions(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticModeOptions:"), value)
}/* debug [instance_properties/setter]: automaticModeOptions */


// The color of this indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/color
func (t_ TextInsertionIndicator) Color() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The color of this indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/color
func (t_ TextInsertionIndicator) SetColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// An optional closure the system calls during dictation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/effectsviewinserter
func (t_ TextInsertionIndicator) EffectsViewInserter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("effectsViewInserter"))
	return rv
}/* debug [instance_properties/getter]: effectsViewInserter */


// An optional closure the system calls during dictation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinsertionindicator/effectsviewinserter
func (t_ TextInsertionIndicator) SetEffectsViewInserter(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEffectsViewInserter:"), value)
}/* debug [instance_properties/setter]: effectsViewInserter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextInsertionIndicator */



