// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	TruncationMode() TextLayerTruncationMode /* not a class type */
	SetTruncationMode(value TextLayerTruncationMode /* not a class type */)
	AlignmentMode() TextLayerAlignmentMode /* not a class type */
	SetAlignmentMode(value TextLayerAlignmentMode /* not a class type */)
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)
	Font() TypeRef /* not a class type */
	SetFont(value TypeRef /* not a class type */)
	FontSize() float64
	SetFontSize(value float64)
	ForegroundColor() objectivec.IObject
	SetForegroundColor(value objectivec.IObject)
	IsWrapped() bool
	SetIsWrapped(value bool)
	String() unsafe.Pointer
	SetString(value unsafe.Pointer)
	// methods:
}

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



// Determines how the text is truncated to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) TruncationMode() TextLayerTruncationMode /* not a class type */ {
	rv := objc.Send[TextLayerTruncationMode](t_.ID, objc.Sel("truncationMode"))
	return rv
}


// Determines how the text is truncated to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer/truncationMode
func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTruncationMode:"), value)
}


// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/alignmentmode
func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode /* not a class type */ {
	rv := objc.Send[TextLayerAlignmentMode](t_.ID, objc.Sel("alignmentMode"))
	return rv
}


// Determines how individual lines of text are horizontally aligned within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/alignmentmode
func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignmentMode:"), value)
}


// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/allowsfontsubpixelquantization
func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsFontSubpixelQuantization"))
	return rv
}


// Determines whether to allow subpixel quantization for the graphics context used for text rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/allowsfontsubpixelquantization
func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsFontSubpixelQuantization:"), value)
}


// The font used to render the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/font
func (t_ TextLayer) Font() TypeRef /* not a class type */ {
	rv := objc.Send[TypeRef](t_.ID, objc.Sel("font"))
	return rv
}


// The font used to render the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/font
func (t_ TextLayer) SetFont(value TypeRef /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}


// The font size used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/fontsize
func (t_ TextLayer) FontSize() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fontSize"))
	return rv
}


// The font size used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/fontsize
func (t_ TextLayer) SetFontSize(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFontSize:"), value)
}


// The color used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/foregroundcolor
func (t_ TextLayer) ForegroundColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("foregroundColor"))
	return rv
}


// The color used to render the receiver’s text. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/foregroundcolor
func (t_ TextLayer) SetForegroundColor(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/iswrapped
func (t_ TextLayer) IsWrapped() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWrapped"))
	return rv
}


// Determines whether the text is wrapped to fit within the receiver’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/iswrapped
func (t_ TextLayer) SetIsWrapped(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWrapped:"), value)
}


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ TextLayer) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("string"))
	return rv
}


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ TextLayer) SetString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}



