// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.OperationQueue)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	RenderingAttributesValidator() unsafe.Pointer
	SetRenderingAttributesValidator(value unsafe.Pointer)
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool)
	TextContainer() ITextContainer
	SetTextContainer(value ITextContainer)
	TextContentManager() ITextContentManager
	TextSelectionNavigation() ITextSelectionNavigation
	SetTextSelectionNavigation(value ITextSelectionNavigation)
	TextSelections() []TextSelection
	SetTextSelections(value []TextSelection)
	TextViewportLayoutController() ITextViewportLayoutController
	UsageBoundsForTextContainer() corefoundation.CGRect
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	UsesHyphenation() bool
	SetUsesHyphenation(value bool)


	

	// methods:
	AddRenderingAttributeValueForTextRange(renderingAttribute AttributedStringKey /* not a class type */, value objectivec.IObject, textRange ITextRange)
	EnsureLayoutForRange(range_ ITextRange)
	EnsureLayoutForBounds(bounds corefoundation.CGRect)
	EnumerateRenderingAttributesFromLocationReverseUsingBlock(location unsafe.Pointer, reverse bool, block unsafe.Pointer)
	EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location unsafe.Pointer, options TextLayoutFragmentEnumerationOptions, block unsafe.Pointer) unsafe.Pointer
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block unsafe.Pointer)
	InvalidateLayoutForRange(range_ ITextRange)
	InvalidateRenderingAttributesForTextRange(textRange ITextRange)
	RemoveRenderingAttributeForTextRange(renderingAttribute AttributedStringKey /* not a class type */, textRange ITextRange)
	RenderingAttributesForLinkAtLocation(link objectivec.IObject, location unsafe.Pointer) foundation.IDictionary
	ReplaceTextContentManager(textContentManager ITextContentManager)
	ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.foundation.INSAttributedString)
	ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement)
	SetRenderingAttributesForTextRange(renderingAttributes foundation.IDictionary, textRange ITextRange)
	TextLayoutFragmentForPosition(position corefoundation.CGPoint) ITextLayoutFragment
	TextLayoutFragmentForLocation(location unsafe.Pointer) ITextLayoutFragment


}





// Alloc allocates a new instance without initialization.
func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a new text layout manager with the coder you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/init(coder:)
func NewTextLayoutManagerWithCoder(coder foundation.foundation.INSCoder) TextLayoutManager {
	instance := getTextLayoutManagerClass().Alloc()
	rv := objc.Send[TextLayoutManager](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}












// Returns the default set of attributes for rendering a link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/linkRenderingAttributes
func (tc _TextLayoutManagerClass) LinkRenderingAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(tc.class), objc.Sel("linkRenderingAttributes"))
	return rv
}






// Sets the rendering attribute for the value and range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/addRenderingAttribute(_:value:for:)
func (t_ TextLayoutManager) AddRenderingAttributeValueForTextRange(renderingAttribute AttributedStringKey /* not a class type */, value objectivec.IObject, textRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addRenderingAttribute:value:forTextRange:"), renderingAttribute, value, textRange)
}


// Performs the layout for specified text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/ensureLayout(for:)-3duae
func (t_ TextLayoutManager) EnsureLayoutForRange(range_ ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("ensureLayoutForRange:"), range_)
}


// Performs the layout for filling the bounds you specify inside the last text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/ensureLayout(for:)-6ptsj
func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds corefoundation.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("ensureLayoutForBounds:"), bounds)
}


// Enumerates the rendering attributes from a location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateRenderingAttributes(from:reverse:using:)
func (t_ TextLayoutManager) EnumerateRenderingAttributesFromLocationReverseUsingBlock(location unsafe.Pointer, reverse bool, block unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateRenderingAttributesFromLocation:reverse:usingBlock:"), location, reverse, block)
}


// Enumerates the text layout fragments starting at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateTextLayoutFragments(from:options:using:)
func (t_ TextLayoutManager) EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location unsafe.Pointer, options TextLayoutFragmentEnumerationOptions, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("enumerateTextLayoutFragmentsFromLocation:options:usingBlock:"), location, options, block)
	return rv
}


// Enumerates text segments of a specific type and in the text range you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/enumerateTextSegments(in:type:options:using:)
func (t_ TextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTextSegmentsInRange:type:options:usingBlock:"), textRange, type_, options, block)
}


// Invalidates the layout information for specified text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/invalidateLayout(for:)
func (t_ TextLayoutManager) InvalidateLayoutForRange(range_ ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateLayoutForRange:"), range_)
}


// Invalidates the rendering attributes of the specified text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/invalidateRenderingAttributes(for:)
func (t_ TextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateRenderingAttributesForTextRange:"), textRange)
}


// Removes the rendering attribute from the specified text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/removeRenderingAttribute(_:for:)
func (t_ TextLayoutManager) RemoveRenderingAttributeForTextRange(renderingAttribute AttributedStringKey /* not a class type */, textRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeRenderingAttribute:forTextRange:"), renderingAttribute, textRange)
}


// Returns a dictionary of rendering attributes for rendering a link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/renderingAttributes(forLink:at:)
func (t_ TextLayoutManager) RenderingAttributesForLinkAtLocation(link objectivec.IObject, location unsafe.Pointer) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("renderingAttributesForLink:atLocation:"), link, location)
	return rv
}


// Replaces the current text content manager with a new one you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replace(_:)
func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceTextContentManager:"), textContentManager)
}


// Replaces content at the location you specify with an attributed string you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-2elb
func (t_ TextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withAttributedString:"), range_, attributedString)
}


// Replaces content at the location you specify with the text elements string you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-80j0b
func (t_ TextLayoutManager) ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []TextElement) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, textElements)
}


// Sets the rendering attributes for the range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/setRenderingAttributes(_:for:)
func (t_ TextLayoutManager) SetRenderingAttributesForTextRange(renderingAttributes foundation.IDictionary, textRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingAttributes:forTextRange:"), renderingAttributes, textRange)
}


// Returns the text layout fragment at the position you specify in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textLayoutFragment(for:)-4dhrx
func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position corefoundation.CGPoint) ITextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("textLayoutFragmentForPosition:"), position)
	return rv
}


// Returns the text layout fragment from the document at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textLayoutFragment(for:)-68dez
func (t_ TextLayoutManager) TextLayoutFragmentForLocation(location unsafe.Pointer) ITextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("textLayoutFragmentForLocation:"), location)
	return rv
}







// The queue that the framework dispatches layout operations on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/layoutQueue
func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := objc.Send[foundation.OperationQueue](t_.ID, objc.Sel("layoutQueue"))
	return rv
}


// The queue that the framework dispatches layout operations on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/layoutQueue
func (t_ TextLayoutManager) SetLayoutQueue(value foundation.OperationQueue) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutQueue:"), value)
}


// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/limitsLayoutForSuspiciousContents
func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}


// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/limitsLayoutForSuspiciousContents
func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}


// Returns the default set of attributes for rendering a link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/linkRenderingAttributes
func (t_ TextLayoutManager) LinkRenderingAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("linkRenderingAttributes"))
	return rv
}


// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/renderingAttributesValidator
func (t_ TextLayoutManager) RenderingAttributesValidator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("renderingAttributesValidator"))
	return rv
}


// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/renderingAttributesValidator
func (t_ TextLayoutManager) SetRenderingAttributesValidator(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRenderingAttributesValidator:"), value)
}


// Specifies the behavior for resolving to the visual alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/resolvesNaturalAlignmentWithBaseWritingDirection
func (t_ TextLayoutManager) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}


// Specifies the behavior for resolving to the visual alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/resolvesNaturalAlignmentWithBaseWritingDirection
func (t_ TextLayoutManager) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
}


// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) TextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}


// Returns the text content manager associated with this text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContentManager
func (t_ TextLayoutManager) TextContentManager() ITextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("textContentManager"))
	return rv
}


// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) TextSelectionNavigation() ITextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("textSelectionNavigation"))
	return rv
}


// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelectionNavigation:"), value)
}


// An array of text selections associated by the text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelections
func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := objc.Send[[]TextSelection](t_.ID, objc.Sel("textSelections"))
	return rv
}


// An array of text selections associated by the text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelections
func (t_ TextLayoutManager) SetTextSelections(value []TextSelection) {
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


// The text viewport layout controller associated with the layout manager’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textViewportLayoutController
func (t_ TextLayoutManager) TextViewportLayoutController() ITextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("textViewportLayoutController"))
	return rv
}


// Returns the usage bounds for the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usageBoundsForTextContainer
func (t_ TextLayoutManager) UsageBoundsForTextContainer() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("usageBoundsForTextContainer"))
	return rv
}


// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesFontLeading
func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesFontLeading
func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}


// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesHyphenation
func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesHyphenation"))
	return rv
}


// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/usesHyphenation
func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesHyphenation:"), value)
}







