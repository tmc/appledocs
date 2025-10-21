// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block unsafe.Pointer)
	ReplaceTextContentManager(textContentManager ITextContentManager)
	ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString)
	ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement)
}

// The primary class that you use to manage text layout and presentation for custom text displays.
//
// is the centerpiece of the TextKit object network that maintains the layout geometry through an array of objects. It lays out results using and objects vended from a that participates in the content layout process.
//
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


// Enumerates text segments of a specific type and in the text range you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateTextSegments(in:type:options:using:)
func (t_ TextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTextSegmentsInRange:type:options:usingBlock:"), textRange, type_, options, block)
}

// Replaces the current text content manager with a new one you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replace(_:)
func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceTextContentManager:"), textContentManager)
}

// Replaces content at the location you specify with an attributed string you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-2elb
func (t_ TextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withAttributedString:"), range_, attributedString)
}

// Replaces content at the location you specify with the text elements string you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-80j0b
func (t_ TextLayoutManager) ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, textElements)
}

// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) TextContainer() NSTextContainer {
	rv := objc.Send[NSTextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// SetTextContainer sets the value of the textContainer property.
// The text container object that provides geometric information for the layout destination.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}

// Returns the text content manager associated with this text layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContentManager
func (t_ TextLayoutManager) TextContentManager() NSTextContentManager {
	rv := objc.Send[NSTextContentManager](t_.ID, objc.Sel("textContentManager"))
	return rv
}

// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) TextSelectionNavigation() NSTextSelectionNavigation {
	rv := objc.Send[NSTextSelectionNavigation](t_.ID, objc.Sel("textSelectionNavigation"))
	return rv
}


// SetTextSelectionNavigation sets the value of the textSelectionNavigation property.
// Returns a text selection manager configured to have the text layout manager as its data source.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelectionNavigation:"), value)
}

// An array of text selections associated by the text layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelections
func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := objc.Send[[]TextSelection](t_.ID, objc.Sel("textSelections"))
	return rv
}


// SetTextSelections sets the value of the textSelections property.
// An array of text selections associated by the text layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelections
func (t_ TextLayoutManager) SetTextSelections(value []TextSelection) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelections:"), nsArray)
}

// Returns the usage bounds for the text container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usageBoundsForTextContainer
func (t_ TextLayoutManager) UsageBoundsForTextContainer() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("usageBoundsForTextContainer"))
	return rv
}

// The delegate for the text layout manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/delegate
func (t_ TextLayoutManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the text layout manager object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/delegate
func (t_ TextLayoutManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// The queue that the framework dispatches layout operations on.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/layoutqueue
func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := objc.Send[foundation.OperationQueue](t_.ID, objc.Sel("layoutQueue"))
	return rv
}


// SetLayoutQueue sets the value of the layoutQueue property.
// The queue that the framework dispatches layout operations on.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/layoutqueue
func (t_ TextLayoutManager) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutQueue:"), value)
}

// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/limitslayoutforsuspiciouscontents
func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}


// SetLimitsLayoutForSuspiciousContents sets the value of the limitsLayoutForSuspiciousContents property.
// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/limitslayoutforsuspiciouscontents
func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}

// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/renderingattributesvalidator
func (t_ TextLayoutManager) RenderingAttributesValidator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("renderingAttributesValidator"))
	return rv
}


// SetRenderingAttributesValidator sets the value of the renderingAttributesValidator property.
// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/renderingattributesvalidator
func (t_ TextLayoutManager) SetRenderingAttributesValidator(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingAttributesValidator:"), value)
}

// Specifies the behavior for resolving
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextLayoutManager) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}


// SetResolvesNaturalAlignmentWithBaseWritingDirection sets the value of the resolvesNaturalAlignmentWithBaseWritingDirection property.
// Specifies the behavior for resolving

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextLayoutManager) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
}

// The text viewport layout controller associated with the layout manager’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textviewportlayoutcontroller
func (t_ TextLayoutManager) TextViewportLayoutController() NSTextViewportLayoutController {
	rv := objc.Send[NSTextViewportLayoutController](t_.ID, objc.Sel("textViewportLayoutController"))
	return rv
}


// SetTextViewportLayoutController sets the value of the textViewportLayoutController property.
// The text viewport layout controller associated with the layout manager’s text container.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/textviewportlayoutcontroller
func (t_ TextLayoutManager) SetTextViewportLayoutController(value ITextViewportLayoutController) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextViewportLayoutController:"), value)
}

// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usesfontleading
func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// SetUsesFontLeading sets the value of the usesFontLeading property.
// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/usesfontleading
func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}

// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/useshyphenation
func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesHyphenation"))
	return rv
}


// SetUsesHyphenation sets the value of the usesHyphenation property.
// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlayoutmanager/useshyphenation
func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesHyphenation:"), value)
}



