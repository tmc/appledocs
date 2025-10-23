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
	// properties:
	Characters() []TextStorage /* primitive/slice/pointer. */
	SetCharacters(value []TextStorage /* primitive/slice/pointer. */)
	FixesAttributesLazily() bool /* primitive/slice/pointer. */
	Words() []TextStorage /* primitive/slice/pointer. */
	SetWords(value []TextStorage /* primitive/slice/pointer. */)
	AttributeRuns() ITextStorage
	SetAttributeRuns(value ITextStorage)
	ChangeInLength() int /* primitive/slice/pointer. */
	SetChangeInLength(value int /* primitive/slice/pointer. */)
	Delegate() TextStorageDelegate /* not a class type */
	SetDelegate(value TextStorageDelegate /* not a class type */)
	EditedMask() TextStorageEditActions /* not a class type */
	SetEditedMask(value TextStorageEditActions /* not a class type */)
	EditedRange() objc.IObject /* cross-framework: Range */
	SetEditedRange(value objc.IObject /* cross-framework: Range */)
	Font() IFont
	SetFont(value IFont)
	ForegroundColor() IColor
	SetForegroundColor(value IColor)
	LayoutManagers() objc.IObject /* cross-framework: LayoutManager */
	SetLayoutManagers(value objc.IObject /* cross-framework: LayoutManager */)
	Paragraphs() ITextStorage
	SetParagraphs(value ITextStorage)
	TextStorageObserver() TextStorageObserving /* not a class type */
	SetTextStorageObserver(value TextStorageObserving /* not a class type */)
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
	// methods:
	ProcessEditing()
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



// Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/processEditing()
func (t_ TextStorage) ProcessEditing() {
	objc.Send[objc.ID](t_.ID, objc.Sel("processEditing"))
}


// The text storage contents as an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) Characters() []TextStorage /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("characters"))
	return rv
}


// The text storage contents as an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) SetCharacters(value []TextStorage /* primitive/slice/pointer. */) {
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


// A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/fixesAttributesLazily
func (t_ TextStorage) FixesAttributesLazily() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("fixesAttributesLazily"))
	return rv
}


// The text storage contents as an array of words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t_ TextStorage) Words() []TextStorage /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("words"))
	return rv
}


// The text storage contents as an array of words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t_ TextStorage) SetWords(value []TextStorage /* primitive/slice/pointer. */) {
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setWords:"), nsArray)
}


// The text storage contents as an array of attribute runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/attributeruns
func (t_ TextStorage) AttributeRuns() ITextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("attributeRuns"))
	return rv
}


// The text storage contents as an array of attribute runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/attributeruns
func (t_ TextStorage) SetAttributeRuns(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributeRuns:"), value)
}


// The difference between the current length of the edited range and its length before editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/changeinlength
func (t_ TextStorage) ChangeInLength() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("changeInLength"))
	return rv
}


// The difference between the current length of the edited range and its length before editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/changeinlength
func (t_ TextStorage) SetChangeInLength(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChangeInLength:"), value)
}


// The delegate for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/delegate
func (t_ TextStorage) Delegate() TextStorageDelegate /* not a class type */ {
	rv := objc.Send[TextStorageDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/delegate
func (t_ TextStorage) SetDelegate(value TextStorageDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A mask that describes the kinds of edits pending for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedmask
func (t_ TextStorage) EditedMask() TextStorageEditActions /* not a class type */ {
	rv := objc.Send[TextStorageEditActions](t_.ID, objc.Sel("editedMask"))
	return rv
}


// A mask that describes the kinds of edits pending for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedmask
func (t_ TextStorage) SetEditedMask(value TextStorageEditActions /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedMask:"), value)
}


// The range of text that contains changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedrange
func (t_ TextStorage) EditedRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("editedRange"))
	return rv
}


// The range of text that contains changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/editedrange
func (t_ TextStorage) SetEditedRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedRange:"), value)
}


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/font
func (t_ TextStorage) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/font
func (t_ TextStorage) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/foregroundcolor
func (t_ TextStorage) ForegroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("foregroundColor"))
	return rv
}


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/foregroundcolor
func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}


// The layout managers for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/layoutmanagers
func (t_ TextStorage) LayoutManagers() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManagers"))
	return rv
}


// The layout managers for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/layoutmanagers
func (t_ TextStorage) SetLayoutManagers(value objc.IObject /* cross-framework: LayoutManager */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManagers:"), value)
}


// The text storage contents as an array of paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/paragraphs
func (t_ TextStorage) Paragraphs() ITextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("paragraphs"))
	return rv
}


// The text storage contents as an array of paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/paragraphs
func (t_ TextStorage) SetParagraphs(value ITextStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphs:"), value)
}


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/textstorageobserver
func (t_ TextStorage) TextStorageObserver() TextStorageObserving /* not a class type */ {
	rv := objc.Send[TextStorageObserving](t_.ID, objc.Sel("textStorageObserver"))
	return rv
}


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/textstorageobserver
func (t_ TextStorage) SetTextStorageObserver(value TextStorageObserving /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStorageObserver:"), value)
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}



