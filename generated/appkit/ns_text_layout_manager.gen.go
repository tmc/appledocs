// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
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
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange unsafe.Pointer, type_ unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer)
	ReplaceTextContentManager(textContentManager unsafe.Pointer)
	ReplaceContentsInRangeWithAttributedString(range_ unsafe.Pointer, attributedString unsafe.Pointer)
	ReplaceContentsInRangeWithTextElements(range_ unsafe.Pointer, textElements unsafe.Pointer)
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
func (t_ TextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange unsafe.Pointer, type_ unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTextSegmentsInRange:type:options:usingBlock:"), textRange, type_, options, block)
}

// Replaces the current text content manager with a new one you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replace(_:)
func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceTextContentManager:"), textContentManager)
}

// Replaces content at the location you specify with an attributed string you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-2elb
func (t_ TextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ unsafe.Pointer, attributedString unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withAttributedString:"), range_, attributedString)
}

// Replaces content at the location you specify with the text elements string you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/replaceContents(in:with:)-80j0b
func (t_ TextLayoutManager) ReplaceContentsInRangeWithTextElements(range_ unsafe.Pointer, textElements unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, textElements)
}

// The text container object that provides geometric information for the layout destination.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) TextContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textContainer"))
	return rv
}

// SetTextContainer sets the value of the textContainer property.
// The text container object that provides geometric information for the layout destination.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContainer
func (t_ TextLayoutManager) SetTextContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}

// Returns the text content manager associated with this text layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textContentManager
func (t_ TextLayoutManager) TextContentManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textContentManager"))
	return rv
}

// Returns a text selection manager configured to have the text layout manager as its data source.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) TextSelectionNavigation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textSelectionNavigation"))
	return rv
}

// SetTextSelectionNavigation sets the value of the textSelectionNavigation property.
// Returns a text selection manager configured to have the text layout manager as its data source.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/textSelectionNavigation
func (t_ TextLayoutManager) SetTextSelectionNavigation(value unsafe.Pointer) {
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
