// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CATextLayer */


/* debug [class_header]: Header for CATextLayer */
// The class instance for the [TextLayer] class.
var (
	TextLayerClass     _TextLayerClass
	TextLayerClassOnce sync.Once
)

func getTextLayerClass() _TextLayerClass {
	TextLayerClassOnce.Do(func() {
		TextLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}
	})
	return TextLayerClass
}

type _TextLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextLayer */
// An interface definition for the [TextLayer] class.
type ITextLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for TextLayer */
	// properties:
	AlignmentMode() TextLayerAlignmentMode /* typedef */
	SetAlignmentMode(value TextLayerAlignmentMode /* typedef */)
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)
	Font() TypeRef /* not a class type */
	SetFont(value TypeRef /* not a class type */)
	FontSize() float64
	SetFontSize(value float64)
	ForegroundColor() ColorRef /* not a class type */
	SetForegroundColor(value ColorRef /* not a class type */)
	Wrapped() bool
	SetWrapped(value bool)
	String() objc.ID
	SetString(value objc.ID)
	TruncationMode() TextLayerTruncationMode /* typedef */
	SetTruncationMode(value TextLayerTruncationMode /* typedef */)
	IsWrapped() bool
	SetIsWrapped(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextLayer */
// Alloc allocates a new instance without initialization.
func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.Send[TextLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextLayerClass) New() TextLayer {
	rv := objc.Send[TextLayer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayer) Init() TextLayer {
	rv := objc.Send[TextLayer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayer) Autorelease() TextLayer {
	rv := objc.Send[TextLayer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayer creates a new TextLayer instance.
func NewTextLayer() TextLayer {
	return getTextLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextLayer */
// A layer that provides simple text layout and rendering of plain or attributed strings.
//
// The first line is aligned to the top of the layer.


// A layer that provides simple text layout and rendering of plain or attributed strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer
type TextLayer struct {
	Layer
}

// TextLayerFrom constructs a [TextLayer] from an unsafe.Pointer.
//
// A layer that provides simple text layout and rendering of plain or attributed strings.
func TextLayerFrom(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: LayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextLayer */

// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/alignmentMode
func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("alignmentMode"))
	return rv
}/* debug [instance_properties/getter]: alignmentMode */


// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/alignmentMode
func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignmentMode:"), value)
}/* debug [instance_properties/setter]: alignmentMode */


// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/allowsFontSubpixelQuantization
func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsFontSubpixelQuantization"))
	return rv
}/* debug [instance_properties/getter]: allowsFontSubpixelQuantization */


// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/allowsFontSubpixelQuantization
func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsFontSubpixelQuantization:"), value)
}/* debug [instance_properties/setter]: allowsFontSubpixelQuantization */


// The font used to render the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/font
func (t_ TextLayer) Font() TypeRef /* not a class type */ {
	rv := objc.Send[TypeRef](t_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font used to render the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/font
func (t_ TextLayer) SetFont(value TypeRef /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The font size used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/fontSize
func (t_ TextLayer) FontSize() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fontSize"))
	return rv
}/* debug [instance_properties/getter]: fontSize */


// The font size used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/fontSize
func (t_ TextLayer) SetFontSize(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFontSize:"), value)
}/* debug [instance_properties/setter]: fontSize */


// The color used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/foregroundColor
func (t_ TextLayer) ForegroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](t_.ID, objc.Sel("foregroundColor"))
	return rv
}/* debug [instance_properties/getter]: foregroundColor */


// The color used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/foregroundColor
func (t_ TextLayer) SetForegroundColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}/* debug [instance_properties/setter]: foregroundColor */


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/isWrapped
func (t_ TextLayer) Wrapped() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("wrapped"))
	return rv
}/* debug [instance_properties/getter]: wrapped */


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/isWrapped
func (t_ TextLayer) SetWrapped(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWrapped:"), value)
}/* debug [instance_properties/setter]: wrapped */


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/string
func (t_ TextLayer) String() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/string
func (t_ TextLayer) SetString(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */


// Determines how the text is truncated to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) TruncationMode() TextLayerTruncationMode /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("truncationMode"))
	return rv
}/* debug [instance_properties/getter]: truncationMode */


// Determines how the text is truncated to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTruncationMode:"), value)
}/* debug [instance_properties/setter]: truncationMode */


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/iswrapped
func (t_ TextLayer) IsWrapped() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWrapped"))
	return rv
}/* debug [instance_properties/getter]: isWrapped */


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/iswrapped
func (t_ TextLayer) SetIsWrapped(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWrapped:"), value)
}/* debug [instance_properties/setter]: isWrapped */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CATextLayer */



