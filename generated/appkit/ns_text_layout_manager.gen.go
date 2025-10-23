// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutManager] class.
var (
	TextLayoutManagerClass     _TextLayoutManagerClass
	TextLayoutManagerClassOnce sync.Once
)

func getTextLayoutManagerClass() _TextLayoutManagerClass {
	TextLayoutManagerClassOnce.Do(func() {
		TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}
	})
	return TextLayoutManagerClass
}

type _TextLayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextLayoutManager] class.
type ITextLayoutManager interface {
	objectivec.IObject
	// properties:
	Delegate() TextLayoutManagerDelegate /* not a class type */
	SetDelegate(value TextLayoutManagerDelegate /* not a class type */)
	LayoutQueue() objc.IObject /* cross-framework: OperationQueue */
	SetLayoutQueue(value objc.IObject /* cross-framework: OperationQueue */)
	LimitsLayoutForSuspiciousContents() bool /* primitive/slice/pointer. */
	SetLimitsLayoutForSuspiciousContents(value bool /* primitive/slice/pointer. */)
	RenderingAttributesValidator() unsafe.Pointer
	SetRenderingAttributesValidator(value unsafe.Pointer)
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool /* primitive/slice/pointer. */
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool /* primitive/slice/pointer. */)
	TextContainer() ITextContainer
	SetTextContainer(value ITextContainer)
	TextContentManager() objc.IObject /* cross-framework: TextContentManager */
	SetTextContentManager(value objc.IObject /* cross-framework: TextContentManager */)
	TextSelectionNavigation() objc.IObject /* cross-framework: TextSelectionNavigation */
	SetTextSelectionNavigation(value objc.IObject /* cross-framework: TextSelectionNavigation */)
	TextSelections() objc.IObject /* cross-framework: TextSelection */
	SetTextSelections(value objc.IObject /* cross-framework: TextSelection */)
	TextViewportLayoutController() objc.IObject /* cross-framework: TextViewportLayoutController */
	SetTextViewportLayoutController(value objc.IObject /* cross-framework: TextViewportLayoutController */)
	UsageBoundsForTextContainer() objc.IObject /* cross-framework: Rect */
	SetUsageBoundsForTextContainer(value objc.IObject /* cross-framework: Rect */)
	UsesFontLeading() bool /* primitive/slice/pointer. */
	SetUsesFontLeading(value bool /* primitive/slice/pointer. */)
	UsesHyphenation() bool /* primitive/slice/pointer. */
	SetUsesHyphenation(value bool /* primitive/slice/pointer. */)
	// methods:
	SetRenderingAttributesForTextRange(renderingAttributes foundation.IDictionary /* already interface */, textRange ITextRange)
}

// The primary class that you use to manage text layout and presentation for custom text displays.
//
// is the centerpiece of the TextKit object network that maintains the layout geometry through an array of objects. It lays out results using and objects vended from a that participates in the content layout process.


// The primary class that you use to manage text layout and presentation for custom text displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager
type TextLayoutManager struct {
	objectivec.Object
}

// TextLayoutManagerFrom constructs a [TextLayoutManager] from an unsafe.Pointer.
//
// The primary class that you use to manage text layout and presentation for custom text displays.
func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayoutManager) Autorelease() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayoutManager creates a new TextLayoutManager instance.
func NewTextLayoutManager() TextLayoutManager {
	return getTextLayoutManagerClass().New()
}



// Sets the rendering attributes for the range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/setRenderingAttributes(_:for:)
func (t_ TextLayoutManager) SetRenderingAttributesForTextRange(renderingAttributes foundation.IDictionary /* already interface */, textRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingAttributes:forTextRange:"), renderingAttributes, textRange)
}


// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/delegate
func (t_ TextLayoutManager) Delegate() TextLayoutManagerDelegate /* not a class type */ {
	rv := objc.Send[TextLayoutManagerDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/delegate
func (t_ TextLayoutManager) SetDelegate(value TextLayoutManagerDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// The queue that the framework dispatches layout operations on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/layoutqueue
func (t_ TextLayoutManager) LayoutQueue() objc.IObject /* cross-framework: OperationQueue */ {
	rv := objc.Send[OperationQueue](t_.ID, objc.Sel("layoutQueue"))
	return rv
}


// The queue that the framework dispatches layout operations on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/layoutqueue
func (t_ TextLayoutManager) SetLayoutQueue(value objc.IObject /* cross-framework: OperationQueue */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutQueue:"), value)
}


// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/limitslayoutforsuspiciouscontents
func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}


// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/limitslayoutforsuspiciouscontents
func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}


// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/renderingattributesvalidator
func (t_ TextLayoutManager) RenderingAttributesValidator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("renderingAttributesValidator"))
	return rv
}


// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/renderingattributesvalidator
func (t_ TextLayoutManager) SetRenderingAttributesValidator(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingAttributesValidator:"), value)
}


// Specifies the behavior for resolving
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextLayoutManager) ResolvesNaturalAlignmentWithBaseWritingDirection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}


// Specifies the behavior for resolving
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextLayoutManager) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
}


// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textcontainer
func (t_ TextLayoutManager) TextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textcontainer
func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}


// Returns the text content manager associated with this text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textcontentmanager
func (t_ TextLayoutManager) TextContentManager() objc.IObject /* cross-framework: TextContentManager */ {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("textContentManager"))
	return rv
}


// Returns the text content manager associated with this text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textcontentmanager
func (t_ TextLayoutManager) SetTextContentManager(value objc.IObject /* cross-framework: TextContentManager */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentManager:"), value)
}


// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textselectionnavigation
func (t_ TextLayoutManager) TextSelectionNavigation() objc.IObject /* cross-framework: TextSelectionNavigation */ {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("textSelectionNavigation"))
	return rv
}


// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textselectionnavigation
func (t_ TextLayoutManager) SetTextSelectionNavigation(value objc.IObject /* cross-framework: TextSelectionNavigation */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelectionNavigation:"), value)
}


// An array of text selections associated by the text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textselections
func (t_ TextLayoutManager) TextSelections() objc.IObject /* cross-framework: TextSelection */ {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("textSelections"))
	return rv
}


// An array of text selections associated by the text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textselections
func (t_ TextLayoutManager) SetTextSelections(value objc.IObject /* cross-framework: TextSelection */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelections:"), value)
}


// The text viewport layout controller associated with the layout manager’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textviewportlayoutcontroller
func (t_ TextLayoutManager) TextViewportLayoutController() objc.IObject /* cross-framework: TextViewportLayoutController */ {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("textViewportLayoutController"))
	return rv
}


// The text viewport layout controller associated with the layout manager’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textviewportlayoutcontroller
func (t_ TextLayoutManager) SetTextViewportLayoutController(value objc.IObject /* cross-framework: TextViewportLayoutController */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextViewportLayoutController:"), value)
}


// Returns the usage bounds for the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usageboundsfortextcontainer
func (t_ TextLayoutManager) UsageBoundsForTextContainer() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("usageBoundsForTextContainer"))
	return rv
}


// Returns the usage bounds for the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usageboundsfortextcontainer
func (t_ TextLayoutManager) SetUsageBoundsForTextContainer(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsageBoundsForTextContainer:"), value)
}


// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usesfontleading
func (t_ TextLayoutManager) UsesFontLeading() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usesfontleading
func (t_ TextLayoutManager) SetUsesFontLeading(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}


// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/useshyphenation
func (t_ TextLayoutManager) UsesHyphenation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesHyphenation"))
	return rv
}


// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/useshyphenation
func (t_ TextLayoutManager) SetUsesHyphenation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesHyphenation:"), value)
}



