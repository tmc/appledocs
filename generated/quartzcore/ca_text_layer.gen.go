// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

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

// An interface definition for the [TextLayer] class.
type ITextLayer interface {
	ILayer
}

// A layer that provides simple text layout and rendering of plain or attributed strings.
//
// The first line is aligned to the top of the layer.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.Send[TextLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/alignmentMode
func (t_ TextLayer) AlignmentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("alignmentMode"))
	return rv
}


// SetAlignmentMode sets the value of the alignmentMode property.
// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/alignmentMode
func (t_ TextLayer) SetAlignmentMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignmentMode:"), value)
}

// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/allowsFontSubpixelQuantization
func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsFontSubpixelQuantization"))
	return rv
}


// SetAllowsFontSubpixelQuantization sets the value of the allowsFontSubpixelQuantization property.
// Determines whether to allow subpixel quantization for the graphics context used for text rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/allowsFontSubpixelQuantization
func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsFontSubpixelQuantization:"), value)
}

// The font used to render the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/font
func (t_ TextLayer) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font used to render the receiver’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/font
func (t_ TextLayer) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}

// The font size used to render the receiver’s text. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/fontSize
func (t_ TextLayer) FontSize() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fontSize"))
	return rv
}


// SetFontSize sets the value of the fontSize property.
// The font size used to render the receiver’s text. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/fontSize
func (t_ TextLayer) SetFontSize(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFontSize:"), value)
}

// The color used to render the receiver’s text. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/foregroundColor
func (t_ TextLayer) ForegroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](t_.ID, objc.Sel("foregroundColor"))
	return rv
}


// SetForegroundColor sets the value of the foregroundColor property.
// The color used to render the receiver’s text. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/foregroundColor
func (t_ TextLayer) SetForegroundColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}

// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/isWrapped
func (t_ TextLayer) Wrapped() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("wrapped"))
	return rv
}


// SetWrapped sets the value of the wrapped property.
// Determines whether the text is wrapped to fit within the receiver’s bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/isWrapped
func (t_ TextLayer) SetWrapped(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWrapped:"), value)
}

// The text to be rendered by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/string
func (t_ TextLayer) String() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("string"))
	return rv
}


// SetString sets the value of the string property.
// The text to be rendered by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/string
func (t_ TextLayer) SetString(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}

// Determines how the text is truncated to fit within the receiver’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) TruncationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("truncationMode"))
	return rv
}


// SetTruncationMode sets the value of the truncationMode property.
// Determines how the text is truncated to fit within the receiver’s bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) SetTruncationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTruncationMode:"), value)
}



