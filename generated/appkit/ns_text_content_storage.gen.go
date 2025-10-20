// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextContentStorage] class.
var (
	TextContentStorageClass     _TextContentStorageClass
	TextContentStorageClassOnce sync.Once
)

func getTextContentStorageClass() _TextContentStorageClass {
	TextContentStorageClassOnce.Do(func() {
		TextContentStorageClass = _TextContentStorageClass{objc.GetClass("NSTextContentStorage")}
	})
	return TextContentStorageClass
}

type _TextContentStorageClass struct {
	class objc.Class
}

// An interface definition for the [TextContentStorage] class.
type ITextContentStorage interface {
	ITextContentManager
	AdjustedRangeFromRangeForEditingTextSelection(textRange unsafe.Pointer, forEditingTextSelection bool) unsafe.Pointer
	LocationFromLocationWithOffset(location objc.ID, offset int) objc.ID
	OffsetFromLocationToLocation(from objc.ID, to objc.ID) int
	TextElementForAttributedString(attributedString unsafe.Pointer) unsafe.Pointer
}

// A concrete object for managing your view’s text content and generating the text elements necessary for layout.
//
// An object provides the backing store for a view that contains text. This object stores the text in an attributed string object, and defaults to using an object. It also maps portions of the text to objects to organize the text into paragraphs, lists, and other common element types found in text content. During layout, TextKit uses these elements to lay out and render the text in your view. The standard system views use an object to manage their text content. When building a custom text view, use this type to store the text for your view. works with an associated to lay out your view’s text. When someone inserts new text or edits the existing text, call the method and use a block to modify the contents of the property. Wrapping your edits in an edit transaction lets the rest of the text system respond to those changes. TextKit uses the abstract protocol to identify locations within text. manager provides its own implementation of this protocol to represent locations within its storage object. To get the start and end locations, access the object’s property and use them to create new location objects. If you provide your own implementation of the protocol to manage locations in your content, subclass and implement your own storage object to support those locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage
type TextContentStorage struct {
	TextContentManager
}

// TextContentStorageFrom constructs a [TextContentStorage] from an unsafe.Pointer.
//
// A concrete object for managing your view’s text content and generating the text elements necessary for layout.
func TextContentStorageFrom(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		TextContentManager: TextContentManagerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextContentStorageClass) Alloc() TextContentStorage {
	rv := objc.Send[TextContentStorage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextContentStorageClass) New() TextContentStorage {
	rv := objc.Send[TextContentStorage](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextContentStorage) Init() TextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextContentStorage) Autorelease() TextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextContentStorage creates a new TextContentStorage instance.
func NewTextContentStorage() TextContentStorage {
	return getTextContentStorageClass().New()
}

// Returns the text range, if any, in the backing store that required manual adjustment after editing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/adjustedRange(from:forEditingTextSelection:)
func (t_ TextContentStorage) AdjustedRangeFromRangeForEditingTextSelection(textRange unsafe.Pointer, forEditingTextSelection bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"), textRange, forEditingTextSelection)
	return rv
}

// Returns a new text location object based on an existing location and offset you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/location(_:offsetBy:)
func (t_ TextContentStorage) LocationFromLocationWithOffset(location objc.ID, offset int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("locationFromLocation:withOffset:"), location, offset)
	return rv
}

// Returns the number of characters between the specified locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/offset(from:to:)
func (t_ TextContentStorage) OffsetFromLocationToLocation(from objc.ID, to objc.ID) int {
	rv := objc.Send[int](t_.ID, objc.Sel("offsetFromLocation:toLocation:"), from, to)
	return rv
}

// Returns the text element corresponding to object’s attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/textElement(for:)
func (t_ TextContentStorage) TextElementForAttributedString(attributedString unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textElementForAttributedString:"), attributedString)
	return rv
}

// An attributed string that contains the contents of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/attributedString
func (t_ TextContentStorage) AttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attributedString"))
	return rv
}

// SetAttributedString sets the value of the attributedString property.
// An attributed string that contains the contents of the document.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/attributedString
func (t_ TextContentStorage) SetAttributedString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/includesTextListMarkers
func (t_ TextContentStorage) IncludesTextListMarkers() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("includesTextListMarkers"))
	return rv
}

// SetIncludesTextListMarkers sets the value of the includesTextListMarkers property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/includesTextListMarkers
func (t_ TextContentStorage) SetIncludesTextListMarkers(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncludesTextListMarkers:"), value)
}
