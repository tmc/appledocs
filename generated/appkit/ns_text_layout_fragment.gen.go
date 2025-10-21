// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutFragment] class.
var (
	TextLayoutFragmentClass     _TextLayoutFragmentClass
	TextLayoutFragmentClassOnce sync.Once
)

func getTextLayoutFragmentClass() _TextLayoutFragmentClass {
	TextLayoutFragmentClassOnce.Do(func() {
		TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}
	})
	return TextLayoutFragmentClass
}

type _TextLayoutFragmentClass struct {
	class objc.Class
}

// An interface definition for the [TextLayoutFragment] class.
type ITextLayoutFragment interface {
	objectivec.IObject
}

// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment
type TextLayoutFragment struct {
	objectivec.Object
}

// TextLayoutFragmentFrom constructs a [TextLayoutFragment] from an unsafe.Pointer.
//
// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayoutFragment) Autorelease() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayoutFragment creates a new TextLayoutFragment instance.
func NewTextLayoutFragment() TextLayoutFragment {
	return getTextLayoutFragmentClass().New()
}


// The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/bottommargin
func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("bottomMargin"))
	return rv
}


// SetBottomMargin sets the value of the bottomMargin property.
// The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/bottommargin
func (t_ TextLayoutFragment) SetBottomMargin(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBottomMargin:"), value)
}

// The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/layoutfragmentframe
func (t_ TextLayoutFragment) LayoutFragmentFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("layoutFragmentFrame"))
	return rv
}


// SetLayoutFragmentFrame sets the value of the layoutFragmentFrame property.
// The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/layoutfragmentframe
func (t_ TextLayoutFragment) SetLayoutFragmentFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutFragmentFrame:"), value)
}

// The queue on which the framework dispatches layout operations.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/layoutqueue
func (t_ TextLayoutFragment) LayoutQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("layoutQueue"))
	return rv
}


// SetLayoutQueue sets the value of the layoutQueue property.
// The queue on which the framework dispatches layout operations.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/layoutqueue
func (t_ TextLayoutFragment) SetLayoutQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutQueue:"), value)
}

// The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/leadingpadding
func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("leadingPadding"))
	return rv
}


// SetLeadingPadding sets the value of the leadingPadding property.
// The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/leadingpadding
func (t_ TextLayoutFragment) SetLeadingPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLeadingPadding:"), value)
}

// The range inside the text element relative to the document origin.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/rangeinelement
func (t_ TextLayoutFragment) RangeInElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("rangeInElement"))
	return rv
}


// SetRangeInElement sets the value of the rangeInElement property.
// The range inside the text element relative to the document origin.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/rangeinelement
func (t_ TextLayoutFragment) SetRangeInElement(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeInElement:"), value)
}

// The bounds defining the area required for rendering the contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/renderingsurfacebounds
func (t_ TextLayoutFragment) RenderingSurfaceBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("renderingSurfaceBounds"))
	return rv
}


// SetRenderingSurfaceBounds sets the value of the renderingSurfaceBounds property.
// The bounds defining the area required for rendering the contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/renderingsurfacebounds
func (t_ TextLayoutFragment) SetRenderingSurfaceBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingSurfaceBounds:"), value)
}

// The layout information state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/state-swift.property
func (t_ TextLayoutFragment) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The layout information state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/state-swift.property
func (t_ TextLayoutFragment) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setState:"), value)
}

// The attachment view provider associated with the text layout fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textattachmentviewproviders
func (t_ TextLayoutFragment) TextAttachmentViewProviders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textAttachmentViewProviders"))
	return rv
}


// SetTextAttachmentViewProviders sets the value of the textAttachmentViewProviders property.
// The attachment view provider associated with the text layout fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textattachmentviewproviders
func (t_ TextLayoutFragment) SetTextAttachmentViewProviders(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextAttachmentViewProviders:"), value)
}

// The parent text element.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textelement
func (t_ TextLayoutFragment) TextElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textElement"))
	return rv
}


// SetTextElement sets the value of the textElement property.
// The parent text element.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textelement
func (t_ TextLayoutFragment) SetTextElement(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextElement:"), value)
}

// The layout manager for this text layout fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textlayoutmanager
func (t_ TextLayoutFragment) TextLayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// SetTextLayoutManager sets the value of the textLayoutManager property.
// The layout manager for this text layout fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textlayoutmanager
func (t_ TextLayoutFragment) SetTextLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}

// An array of text line fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textlinefragments
func (t_ TextLayoutFragment) TextLineFragments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textLineFragments"))
	return rv
}


// SetTextLineFragments sets the value of the textLineFragments property.
// An array of text line fragments.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/textlinefragments
func (t_ TextLayoutFragment) SetTextLineFragments(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLineFragments:"), value)
}

// The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/topmargin
func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("topMargin"))
	return rv
}


// SetTopMargin sets the value of the topMargin property.
// The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/topmargin
func (t_ TextLayoutFragment) SetTopMargin(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTopMargin:"), value)
}

// The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/trailingpadding
func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("trailingPadding"))
	return rv
}


// SetTrailingPadding sets the value of the trailingPadding property.
// The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutfragment/trailingpadding
func (t_ TextLayoutFragment) SetTrailingPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrailingPadding:"), value)
}



