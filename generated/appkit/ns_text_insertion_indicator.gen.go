// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [TextInsertionIndicator] class.
type ITextInsertionIndicator interface {
	IView
	AutomaticModeOptions() TextInsertionIndicatorAutomaticModeOptions
	SetAutomaticModeOptions(value TextInsertionIndicatorAutomaticModeOptions)
	Color() NSColor
	SetColor(value IColor)
	DisplayMode() TextInsertionIndicatorDisplayMode
	SetDisplayMode(value TextInsertionIndicatorDisplayMode)
	EffectsViewInserter() unsafe.Pointer
	SetEffectsViewInserter(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (tc _TextInsertionIndicatorClass) Alloc() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Options that affect the automatic display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/automaticModeOptions-swift.property
func (t_ TextInsertionIndicator) AutomaticModeOptions() TextInsertionIndicatorAutomaticModeOptions {
	rv := objc.Send[TextInsertionIndicatorAutomaticModeOptions](t_.ID, objc.Sel("automaticModeOptions"))
	return rv
}


// Options that affect the automatic display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/automaticModeOptions-swift.property
func (t_ TextInsertionIndicator) SetAutomaticModeOptions(value TextInsertionIndicatorAutomaticModeOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticModeOptions:"), value)
}


// The color of this indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/color
func (t_ TextInsertionIndicator) Color() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("color"))
	return rv
}


// The color of this indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/color
func (t_ TextInsertionIndicator) SetColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}


// A value that describes the display mode of an indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) DisplayMode() TextInsertionIndicatorDisplayMode {
	rv := objc.Send[TextInsertionIndicatorDisplayMode](t_.ID, objc.Sel("displayMode"))
	return rv
}


// A value that describes the display mode of an indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) SetDisplayMode(value TextInsertionIndicatorDisplayMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplayMode:"), value)
}


// An optional closure the system calls during dictation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/effectsViewInserter
func (t_ TextInsertionIndicator) EffectsViewInserter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("effectsViewInserter"))
	return rv
}


// An optional closure the system calls during dictation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/effectsViewInserter
func (t_ TextInsertionIndicator) SetEffectsViewInserter(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEffectsViewInserter:"), value)
}



