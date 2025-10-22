// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Characters() []TextStorage
	SetCharacters(value []TextStorage)
	TextStorageObserver() objc.ID
	SetTextStorageObserver(value objc.ID)
	AttributeRuns() NSTextStorage
	SetAttributeRuns(value ITextStorage)
	ChangeInLength() int
	SetChangeInLength(value int)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	EditedMask() unsafe.Pointer
	SetEditedMask(value unsafe.Pointer)
	EditedRange() foundation.Range
	SetEditedRange(value foundation.IRange)
	FixesAttributesLazily() bool
	SetFixesAttributesLazily(value bool)
	Font() NSFont
	SetFont(value IFont)
	ForegroundColor() NSColor
	SetForegroundColor(value IColor)
	LayoutManagers() NSLayoutManager
	SetLayoutManagers(value ILayoutManager)
	Paragraphs() NSTextStorage
	SetParagraphs(value ITextStorage)
	Words() NSTextStorage
	SetWords(value ITextStorage)
	String() string
	SetString(value string)
}

// The fundamental storage mechanism of TextKit that contains the text managed by the system.
//
// is a semi-concrete subclass of that adds behavior for managing a set of client objects. A text storage object notifies its layout managers of changes to its characters or attributes, which lets the layout managers redisplay the text as needed. You can access a text storage object from any thread of your app, but your app must guarantee access from only one thread at a time. In macOS, this class also defines properties for getting and setting scriptable attributes of objects. Unless you’re dealing with scriptability, you shouldn’t access these properties directly. In particular, using the , , or properties is an inefficient way to manipulate the text storage, since accessing these properties involves the creation of many objects. Instead, use the text access methods defined by , , , and to perform character-level manipulation.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The text storage contents as an array of characters.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("characters"))
	return rv
}


// SetCharacters sets the value of the characters property.
// The text storage contents as an array of characters.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) SetCharacters(value []TextStorage) {
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setCharacters:"), nsArray)
}

// The observer for the text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) TextStorageObserver() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("textStorageObserver"))
	return rv
}


// SetTextStorageObserver sets the value of the textStorageObserver property.
// The observer for the text storage object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) SetTextStorageObserver(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStorageObserver:"), value)
}

// The text storage contents as an array of attribute runs.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/attributeruns
func (t_ TextStorage) AttributeRuns() NSTextStorage {
	rv := objc.Send[NSTextStorage](t_.ID, objc.Sel("attributeRuns"))
	return rv
}


// SetAttributeRuns sets the value of the attributeRuns property.
// The text storage contents as an array of attribute runs.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/attributeruns
func (t_ TextStorage) SetAttributeRuns(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributeRuns:"), value)
}

// The difference between the current length of the edited range and its length before editing.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/changeinlength
func (t_ TextStorage) ChangeInLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("changeInLength"))
	return rv
}


// SetChangeInLength sets the value of the changeInLength property.
// The difference between the current length of the edited range and its length before editing.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/changeinlength
func (t_ TextStorage) SetChangeInLength(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChangeInLength:"), value)
}

// The delegate for the text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/delegate
func (t_ TextStorage) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the text storage object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/delegate
func (t_ TextStorage) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A mask that describes the kinds of edits pending for the text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedmask
func (t_ TextStorage) EditedMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("editedMask"))
	return rv
}


// SetEditedMask sets the value of the editedMask property.
// A mask that describes the kinds of edits pending for the text storage object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedmask
func (t_ TextStorage) SetEditedMask(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedMask:"), value)
}

// The range of text that contains changes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedrange
func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("editedRange"))
	return rv
}


// SetEditedRange sets the value of the editedRange property.
// The range of text that contains changes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedrange
func (t_ TextStorage) SetEditedRange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedRange:"), value)
}

// A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/fixesattributeslazily
func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fixesAttributesLazily"))
	return rv
}


// SetFixesAttributesLazily sets the value of the fixesAttributesLazily property.
// A Boolean value that indicates whether the text storage object fixes attributes lazily.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/fixesattributeslazily
func (t_ TextStorage) SetFixesAttributesLazily(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFixesAttributesLazily:"), value)
}

// The font for the text storage.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/font
func (t_ TextStorage) Font() NSFont {
	rv := objc.Send[NSFont](t_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font for the text storage.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/font
func (t_ TextStorage) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}

// The color for the text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/foregroundcolor
func (t_ TextStorage) ForegroundColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("foregroundColor"))
	return rv
}


// SetForegroundColor sets the value of the foregroundColor property.
// The color for the text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/foregroundcolor
func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}

// The layout managers for the text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/layoutmanagers
func (t_ TextStorage) LayoutManagers() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](t_.ID, objc.Sel("layoutManagers"))
	return rv
}


// SetLayoutManagers sets the value of the layoutManagers property.
// The layout managers for the text storage object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/layoutmanagers
func (t_ TextStorage) SetLayoutManagers(value ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManagers:"), value)
}

// The text storage contents as an array of paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/paragraphs
func (t_ TextStorage) Paragraphs() NSTextStorage {
	rv := objc.Send[NSTextStorage](t_.ID, objc.Sel("paragraphs"))
	return rv
}


// SetParagraphs sets the value of the paragraphs property.
// The text storage contents as an array of paragraphs.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/paragraphs
func (t_ TextStorage) SetParagraphs(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphs:"), value)
}

// The text storage contents as an array of words.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/words
func (t_ TextStorage) Words() NSTextStorage {
	rv := objc.Send[NSTextStorage](t_.ID, objc.Sel("words"))
	return rv
}


// SetWords sets the value of the words property.
// The text storage contents as an array of words.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/words
func (t_ TextStorage) SetWords(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWords:"), value)
}

// The character contents of the attributed string as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) String() string {
	rv := objc.Send[string](t_.ID, objc.Sel("string"))
	return rv
}


// SetString sets the value of the string property.
// The character contents of the attributed string as a string.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) SetString(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), objc.String(value))
}



