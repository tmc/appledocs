// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [TextStorage] class.
var (
	TextStorageClass     _TextStorageClass
	TextStorageClassOnce sync.Once
)

func getTextStorageClass() _TextStorageClass {
	TextStorageClassOnce.Do(func() {
		TextStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}
	})
	return TextStorageClass
}

type _TextStorageClass struct {
	class objc.Class
}





// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	IMutableAttributedString
	

	// properties:
	AttributeRuns() []TextStorage
	SetAttributeRuns(value []TextStorage)
	ChangeInLength() int
	Characters() []TextStorage
	SetCharacters(value []TextStorage)
	EditedMask() TextStorageEditActions
	EditedRange() foundation.Range
	FixesAttributesLazily() bool
	Font() IFont
	SetFont(value IFont)
	ForegroundColor() IColor
	SetForegroundColor(value IColor)
	LayoutManagers() []LayoutManager
	Paragraphs() []TextStorage
	SetParagraphs(value []TextStorage)
	TextStorageObserver() unsafe.Pointer
	SetTextStorageObserver(value unsafe.Pointer)
	Words() []TextStorage
	SetWords(value []TextStorage)
	String() foundation.foundation.INSString
	SetString(value foundation.foundation.INSString)


	

	// methods:
	AddLayoutManager(aLayoutManager ILayoutManager)
	EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	EnsureAttributesAreFixedInRange(range_ foundation.Range)
	InvalidateAttributesInRange(range_ foundation.Range)
	ProcessEditing()
	RemoveLayoutManager(aLayoutManager ILayoutManager)


}





// Alloc allocates a new instance without initialization.
func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextStorageClass) New() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextStorage) Init() TextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextStorage) Autorelease() TextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextStorage creates a new TextStorage instance.
func NewTextStorage() TextStorage {
	return getTextStorageClass().New()
}





// The fundamental storage mechanism of TextKit that contains the text managed by the system.
//
// is a semi-concrete subclass of that adds behavior for managing a set of client objects. A text storage object notifies its layout managers of changes to its characters or attributes, which lets the layout managers redisplay the text as needed. You can access a text storage object from any thread of your app, but your app must guarantee access from only one thread at a time. In macOS, this class also defines properties for getting and setting scriptable attributes of objects. Unless you’re dealing with scriptability, you shouldn’t access these properties directly. In particular, using the , , or properties is an inefficient way to manipulate the text storage, since accessing these properties involves the creation of many objects. Instead, use the text access methods defined by , , , and to perform character-level manipulation.


// The fundamental storage mechanism of TextKit that contains the text managed by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage
type TextStorage struct {
	MutableAttributedString
}

// TextStorageFrom constructs a [TextStorage] from an unsafe.Pointer.
//
// The fundamental storage mechanism of TextKit that contains the text managed by the system.
func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: MutableAttributedStringFrom(ptr),
	}
}




















// Adds a layout manager to the text storage object’s set of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/addLayoutManager(_:)
func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addLayoutManager:"), aLayoutManager)
}


// Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/edited(_:range:changeInLength:)
func (t_ TextStorage) EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("edited:range:changeInLength:"), editedMask, editedRange, delta)
}


// Ensures that attribute fixing occurs in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/ensureAttributesAreFixed(in:)
func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("ensureAttributesAreFixedInRange:"), range_)
}


// Invalidates attributes in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/invalidateAttributes(in:)
func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateAttributesInRange:"), range_)
}


// Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/processEditing()
func (t_ TextStorage) ProcessEditing() {
	objc.Send[objc.ID](t_.ID, objc.Sel("processEditing"))
}


// Removes a layout manager from the text storage object’s set of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/removeLayoutManager(_:)
func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeLayoutManager:"), aLayoutManager)
}







// The text storage contents as an array of attribute runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/attributeRuns
func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("attributeRuns"))
	return rv
}


// The text storage contents as an array of attribute runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/attributeRuns
func (t_ TextStorage) SetAttributeRuns(value []TextStorage) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributeRuns:"), nsArray)
}


// The difference between the current length of the edited range and its length before editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/changeInLength
func (t_ TextStorage) ChangeInLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("changeInLength"))
	return rv
}


// The text storage contents as an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("characters"))
	return rv
}


// The text storage contents as an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) SetCharacters(value []TextStorage) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setCharacters:"), nsArray)
}


// A mask that describes the kinds of edits pending for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedMask
func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.Send[TextStorageEditActions](t_.ID, objc.Sel("editedMask"))
	return rv
}


// The range of text that contains changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedRange
func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("editedRange"))
	return rv
}


// A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/fixesAttributesLazily
func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fixesAttributesLazily"))
	return rv
}


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/font
func (t_ TextStorage) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/font
func (t_ TextStorage) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/foregroundColor
func (t_ TextStorage) ForegroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("foregroundColor"))
	return rv
}


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/foregroundColor
func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}


// The layout managers for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/layoutManagers
func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.Send[[]LayoutManager](t_.ID, objc.Sel("layoutManagers"))
	return rv
}


// The text storage contents as an array of paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/paragraphs
func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("paragraphs"))
	return rv
}


// The text storage contents as an array of paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/paragraphs
func (t_ TextStorage) SetParagraphs(value []TextStorage) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphs:"), nsArray)
}


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) TextStorageObserver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textStorageObserver"))
	return rv
}


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) SetTextStorageObserver(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStorageObserver:"), value)
}


// The text storage contents as an array of words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t_ TextStorage) Words() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("words"))
	return rv
}


// The text storage contents as an array of words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t_ TextStorage) SetWords(value []TextStorage) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setWords:"), nsArray)
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) String() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) SetString(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}








