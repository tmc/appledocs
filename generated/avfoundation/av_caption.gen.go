// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaption */


/* debug [class_header]: Header for AVCaption */
// The class instance for the [Caption] class.
var (
	CaptionClass     _CaptionClass
	CaptionClassOnce sync.Once
)

func getCaptionClass() _CaptionClass {
	CaptionClassOnce.Do(func() {
		CaptionClass = _CaptionClass{objc.GetClass("AVCaption")}
	})
	return CaptionClass
}

type _CaptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Caption */
// An interface definition for the [Caption] class.
type ICaption interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Caption */
	// properties:
	Animation() CaptionAnimation
	Region() IAVCaptionRegion
	Text() objc.IObject /* cross-framework: NSString */
	TextAlignment() CaptionTextAlignment
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Caption */
	// methods:
	BackgroundColorAtIndexRange(index int, outRange corefoundation.Range) ColorRef /* not a class type */
	DecorationAtIndexRange(index int, outRange corefoundation.Range) CaptionDecoration
	FontStyleAtIndexRange(index int, outRange corefoundation.Range) CaptionFontStyle
	FontWeightAtIndexRange(index int, outRange corefoundation.Range) CaptionFontWeight
	RubyAtIndexRange(index int, outRange corefoundation.Range) ICaptionRuby
	TextColorAtIndexRange(index int, outRange corefoundation.Range) ColorRef /* not a class type */
	TextCombineAtIndexRange(index int, outRange corefoundation.Range) CaptionTextCombine
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Caption */
// Alloc allocates a new instance without initialization.
func (cc _CaptionClass) Alloc() Caption {
	rv := objc.Send[Caption](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionClass) New() Caption {
	rv := objc.Send[Caption](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Caption) Init() Caption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Caption) Autorelease() Caption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaption creates a new Caption instance.
func NewCaption() Caption {
	return getCaptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Caption */
// An object that represents text to present over a time range.
//
// A caption contains a cue, which is a single sentence or paragraph of text for a time range in the video timeline. Within the active range, the caption may animate (for example, Karaoke lyrics) by rolling-up, changing visibility, or using other dynamic styling.


// An object that represents text to present over a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption
type Caption struct {
	objectivec.Object
}

// CaptionFrom constructs a [Caption] from an unsafe.Pointer.
//
// An object that represents text to present over a time range.
func CaptionFrom(ptr unsafe.Pointer) Caption {
	return Caption{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Caption */

// Creates a caption that contains text and a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/init(_:timeRange:)
func NewCaptionWithTextTimeRange(text objc.IObject /* cross-framework: NSString */, timeRange TimeRange /* not a class type */) Caption {
	instance := getCaptionClass().Alloc()
	rv := objc.Send[Caption](instance.ID, objc.Sel("initWithText:timeRange:"), text, timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionWithTextTimeRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Caption */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Caption */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Caption */

// Returns the background color at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/backgroundColorAtIndex:range:
func (c_ Caption) BackgroundColorAtIndexRange(index int, outRange corefoundation.Range) ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](c_.ID, objc.Sel("backgroundColorAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: BackgroundColorAtIndexRange */


// Returns the text decoration at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/decorationAtIndex:range:
func (c_ Caption) DecorationAtIndexRange(index int, outRange corefoundation.Range) CaptionDecoration {
	rv := objc.Send[CaptionDecoration](c_.ID, objc.Sel("decorationAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: DecorationAtIndexRange */


// Returns the font style and range at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/fontStyleAtIndex:range:
func (c_ Caption) FontStyleAtIndexRange(index int, outRange corefoundation.Range) CaptionFontStyle {
	rv := objc.Send[CaptionFontStyle](c_.ID, objc.Sel("fontStyleAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: FontStyleAtIndexRange */


// Returns the font weight and range at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/fontWeightAtIndex:range:
func (c_ Caption) FontWeightAtIndexRange(index int, outRange corefoundation.Range) CaptionFontWeight {
	rv := objc.Send[CaptionFontWeight](c_.ID, objc.Sel("fontWeightAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: FontWeightAtIndexRange */


// Returns the ruby text at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/rubyAtIndex:range:
func (c_ Caption) RubyAtIndexRange(index int, outRange corefoundation.Range) ICaptionRuby {
	rv := objc.Send[CaptionRuby](c_.ID, objc.Sel("rubyAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: RubyAtIndexRange */


// Returns the text color at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/textColorAtIndex:range:
func (c_ Caption) TextColorAtIndexRange(index int, outRange corefoundation.Range) ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](c_.ID, objc.Sel("textColorAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: TextColorAtIndexRange */


// Returns the text combine at the index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/textCombineAtIndex:range:
func (c_ Caption) TextCombineAtIndexRange(index int, outRange corefoundation.Range) CaptionTextCombine {
	rv := objc.Send[CaptionTextCombine](c_.ID, objc.Sel("textCombineAtIndex:range:"), index, outRange)
	return rv
}/* debug [instance_methods/method]: TextCombineAtIndexRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Caption */

// The animation that the system applies to this caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/animation-swift.property
func (c_ Caption) Animation() CaptionAnimation {
	rv := objc.Send[CaptionAnimation](c_.ID, objc.Sel("animation"))
	return rv
}/* debug [instance_properties/getter]: animation */


// The region in which the caption exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/region
func (c_ Caption) Region() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/text
func (c_ Caption) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("text"))
	return rv
}/* debug [instance_properties/getter]: text */


// The alignment for the caption text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/textAlignment-swift.property
func (c_ Caption) TextAlignment() CaptionTextAlignment {
	rv := objc.Send[CaptionTextAlignment](c_.ID, objc.Sel("textAlignment"))
	return rv
}/* debug [instance_properties/getter]: textAlignment */


// The time range over which the system presents the caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/timeRange
func (c_ Caption) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaption */


