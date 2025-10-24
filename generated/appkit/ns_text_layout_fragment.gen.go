// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextLayoutFragment */


/* debug [class_header]: Header for NSTextLayoutFragment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextLayoutFragment */
// An interface definition for the [TextLayoutFragment] class.
type ITextLayoutFragment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextLayoutFragment */
	// properties:
	BottomMargin() float64
	LayoutFragmentFrame() corefoundation.CGRect
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.OperationQueue)
	LeadingPadding() float64
	RangeInElement() ITextRange
	RenderingSurfaceBounds() corefoundation.CGRect
	State() TextLayoutFragmentState
	TextAttachmentViewProviders() []TextAttachmentViewProvider
	TextElement() ITextElement
	TextLayoutManager() ITextLayoutManager
	TextLineFragments() []TextLineFragment
	TopMargin() float64
	TrailingPadding() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextLayoutFragment */
	// methods:
	DrawAtPointInContext(point corefoundation.CGPoint, context ContextRef /* not a class type */)
	FrameForTextAttachmentAtLocation(location unsafe.Pointer) corefoundation.CGRect
	InvalidateLayout()
	TextLineFragmentForTextLocationIsUpstreamAffinity(textLocation unsafe.Pointer, isUpstreamAffinity bool) ITextLineFragment
	TextLineFragmentForVerticalOffsetRequiresExactMatch(verticalOffset float64, requiresExactMatch bool) ITextLineFragment
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextLayoutFragment */
// Alloc allocates a new instance without initialization.
func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextLayoutFragment */
// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.


// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextLayoutFragment */

// Creates a new layout fragment with the coder you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(coder:)
func NewTextLayoutFragmentWithCoder(coder foundation.Coder) TextLayoutFragment {
	instance := getTextLayoutFragmentClass().Alloc()
	rv := objc.Send[TextLayoutFragment](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextLayoutFragmentWithCoder */


// Create a new layout fragment using the provided text element and range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/init(textElement:range:)
func NewTextLayoutFragmentWithTextElementRange(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	instance := getTextLayoutFragmentClass().Alloc()
	rv := objc.Send[TextLayoutFragment](instance.ID, objc.Sel("initWithTextElement:range:"), textElement, rangeInElement)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextLayoutFragmentWithTextElementRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextLayoutFragment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextLayoutFragment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextLayoutFragment */

// Renders the visual representation of this element in the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/draw(at:in:)
func (t_ TextLayoutFragment) DrawAtPointInContext(point corefoundation.CGPoint, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawAtPoint:inContext:"), point, context)
}/* debug [instance_methods/method]: DrawAtPointInContext */


// Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/frameForTextAttachment(at:)
func (t_ TextLayoutFragment) FrameForTextAttachmentAtLocation(location unsafe.Pointer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("frameForTextAttachmentAtLocation:"), location)
	return rv
}/* debug [instance_methods/method]: FrameForTextAttachmentAtLocation */


// Invalidates any layout information associated with the text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/invalidateLayout()
func (t_ TextLayoutFragment) InvalidateLayout() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateLayout"))
}/* debug [instance_methods/method]: InvalidateLayout */


// Returns a text line fragment from a specific text location in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragment(for:isUpstreamAffinity:)
func (t_ TextLayoutFragment) TextLineFragmentForTextLocationIsUpstreamAffinity(textLocation unsafe.Pointer, isUpstreamAffinity bool) ITextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("textLineFragmentForTextLocation:isUpstreamAffinity:"), textLocation, isUpstreamAffinity)
	return rv
}/* debug [instance_methods/method]: TextLineFragmentForTextLocationIsUpstreamAffinity */


// Returns the text line fragment for the vertical offset you provide, or the closest text line fragment beyond the vertical offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragment(forVerticalOffset:requiresExactMatch:)
func (t_ TextLayoutFragment) TextLineFragmentForVerticalOffsetRequiresExactMatch(verticalOffset float64, requiresExactMatch bool) ITextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("textLineFragmentForVerticalOffset:requiresExactMatch:"), verticalOffset, requiresExactMatch)
	return rv
}/* debug [instance_methods/method]: TextLineFragmentForVerticalOffsetRequiresExactMatch */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextLayoutFragment */

// The amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/bottomMargin
func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("bottomMargin"))
	return rv
}/* debug [instance_properties/getter]: bottomMargin */


// The rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/layoutFragmentFrame
func (t_ TextLayoutFragment) LayoutFragmentFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("layoutFragmentFrame"))
	return rv
}/* debug [instance_properties/getter]: layoutFragmentFrame */


// The queue on which the framework dispatches layout operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/layoutQueue
func (t_ TextLayoutFragment) LayoutQueue() foundation.OperationQueue {
	rv := objc.Send[foundation.OperationQueue](t_.ID, objc.Sel("layoutQueue"))
	return rv
}/* debug [instance_properties/getter]: layoutQueue */


// The queue on which the framework dispatches layout operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/layoutQueue
func (t_ TextLayoutFragment) SetLayoutQueue(value foundation.OperationQueue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutQueue:"), value)
}/* debug [instance_properties/setter]: layoutQueue */


// The amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/leadingPadding
func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("leadingPadding"))
	return rv
}/* debug [instance_properties/getter]: leadingPadding */


// The range inside the text element relative to the document origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/rangeInElement
func (t_ TextLayoutFragment) RangeInElement() ITextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("rangeInElement"))
	return rv
}/* debug [instance_properties/getter]: rangeInElement */


// The bounds defining the area required for rendering the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/renderingSurfaceBounds
func (t_ TextLayoutFragment) RenderingSurfaceBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("renderingSurfaceBounds"))
	return rv
}/* debug [instance_properties/getter]: renderingSurfaceBounds */


// The layout information state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/state-swift.property
func (t_ TextLayoutFragment) State() TextLayoutFragmentState {
	rv := objc.Send[TextLayoutFragmentState](t_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The attachment view provider associated with the text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textAttachmentViewProviders
func (t_ TextLayoutFragment) TextAttachmentViewProviders() []TextAttachmentViewProvider {
	rv := objc.Send[[]TextAttachmentViewProvider](t_.ID, objc.Sel("textAttachmentViewProviders"))
	return rv
}/* debug [instance_properties/getter]: textAttachmentViewProviders */


// The parent text element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textElement
func (t_ TextLayoutFragment) TextElement() ITextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("textElement"))
	return rv
}/* debug [instance_properties/getter]: textElement */


// The layout manager for this text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLayoutManager
func (t_ TextLayoutFragment) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}/* debug [instance_properties/getter]: textLayoutManager */


// An array of text line fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/textLineFragments
func (t_ TextLayoutFragment) TextLineFragments() []TextLineFragment {
	rv := objc.Send[[]TextLineFragment](t_.ID, objc.Sel("textLineFragments"))
	return rv
}/* debug [instance_properties/getter]: textLineFragments */


// The amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/topMargin
func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("topMargin"))
	return rv
}/* debug [instance_properties/getter]: topMargin */


// The amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/trailingPadding
func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("trailingPadding"))
	return rv
}/* debug [instance_properties/getter]: trailingPadding */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextLayoutFragment */


