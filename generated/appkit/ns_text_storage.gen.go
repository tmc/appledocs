// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class NSTextStorage */


/* debug [class_header]: Header for NSTextStorage */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextStorage */
// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	IMutableAttributedString
	
/* debug [class_interface_properties]: Properties for TextStorage */
	// properties:
	AttributeRuns() []TextStorage
	SetAttributeRuns(value []TextStorage)
	ChangeInLength() int
	Characters() []TextStorage
	SetCharacters(value []TextStorage)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	EditedMask() TextStorageEditActions
	EditedRange() corefoundation.Range
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
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextStorage */
	// methods:
	AddLayoutManager(aLayoutManager ILayoutManager)
	EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int)
	EnsureAttributesAreFixedInRange(range_ corefoundation.Range)
	InvalidateAttributesInRange(range_ corefoundation.Range)
	ProcessEditing()
	RemoveLayoutManager(aLayoutManager ILayoutManager)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextStorage */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextStorage */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextStorage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextStorage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextStorage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextStorage */

// Adds a layout manager to the text storage object’s set of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/addLayoutManager(_:)
func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addLayoutManager:"), aLayoutManager)
}/* debug [instance_methods/method]: AddLayoutManager */


// Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/edited(_:range:changeInLength:)
func (t_ TextStorage) EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange corefoundation.Range, delta int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("edited:range:changeInLength:"), editedMask, editedRange, delta)
}/* debug [instance_methods/method]: EditedRangeChangeInLength */


// Ensures that attribute fixing occurs in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/ensureAttributesAreFixed(in:)
func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("ensureAttributesAreFixedInRange:"), range_)
}/* debug [instance_methods/method]: EnsureAttributesAreFixedInRange */


// Invalidates attributes in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/invalidateAttributes(in:)
func (t_ TextStorage) InvalidateAttributesInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateAttributesInRange:"), range_)
}/* debug [instance_methods/method]: InvalidateAttributesInRange */


// Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/processEditing()
func (t_ TextStorage) ProcessEditing() {
	objc.Send[objc.ID](t_.ID, objc.Sel("processEditing"))
}/* debug [instance_methods/method]: ProcessEditing */


// Removes a layout manager from the text storage object’s set of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/removeLayoutManager(_:)
func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeLayoutManager:"), aLayoutManager)
}/* debug [instance_methods/method]: RemoveLayoutManager */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextStorage */

// The text storage contents as an array of attribute runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/attributeRuns
func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("attributeRuns"))
	return rv
}/* debug [instance_properties/getter]: attributeRuns */


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
}/* debug [instance_properties/setter]: attributeRuns */


// The difference between the current length of the edited range and its length before editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/changeInLength
func (t_ TextStorage) ChangeInLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("changeInLength"))
	return rv
}/* debug [instance_properties/getter]: changeInLength */


// The text storage contents as an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("characters"))
	return rv
}/* debug [instance_properties/getter]: characters */


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
}/* debug [instance_properties/setter]: characters */


// The delegate for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/delegate
func (t_ TextStorage) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/delegate
func (t_ TextStorage) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A mask that describes the kinds of edits pending for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedMask
func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.Send[TextStorageEditActions](t_.ID, objc.Sel("editedMask"))
	return rv
}/* debug [instance_properties/getter]: editedMask */


// The range of text that contains changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedRange
func (t_ TextStorage) EditedRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("editedRange"))
	return rv
}/* debug [instance_properties/getter]: editedRange */


// A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/fixesAttributesLazily
func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fixesAttributesLazily"))
	return rv
}/* debug [instance_properties/getter]: fixesAttributesLazily */


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/font
func (t_ TextStorage) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font for the text storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/font
func (t_ TextStorage) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/foregroundColor
func (t_ TextStorage) ForegroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("foregroundColor"))
	return rv
}/* debug [instance_properties/getter]: foregroundColor */


// The color for the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/foregroundColor
func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setForegroundColor:"), value)
}/* debug [instance_properties/setter]: foregroundColor */


// The layout managers for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/layoutManagers
func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.Send[[]LayoutManager](t_.ID, objc.Sel("layoutManagers"))
	return rv
}/* debug [instance_properties/getter]: layoutManagers */


// The text storage contents as an array of paragraphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/paragraphs
func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("paragraphs"))
	return rv
}/* debug [instance_properties/getter]: paragraphs */


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
}/* debug [instance_properties/setter]: paragraphs */


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) TextStorageObserver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textStorageObserver"))
	return rv
}/* debug [instance_properties/getter]: textStorageObserver */


// The observer for the text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t_ TextStorage) SetTextStorageObserver(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextStorageObserver:"), value)
}/* debug [instance_properties/setter]: textStorageObserver */


// The text storage contents as an array of words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t_ TextStorage) Words() []TextStorage {
	rv := objc.Send[[]TextStorage](t_.ID, objc.Sel("words"))
	return rv
}/* debug [instance_properties/getter]: words */


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
}/* debug [instance_properties/setter]: words */


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/string
func (t_ TextStorage) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextStorage */



