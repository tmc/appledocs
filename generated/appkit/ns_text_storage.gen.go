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
	foundation.IMutableAttributedString
}

// The fundamental storage mechanism of TextKit that contains the text managed by the system.
//
// is a semi-concrete subclass of that adds behavior for managing a set of client objects. A text storage object notifies its layout managers of changes to its characters or attributes, which lets the layout managers redisplay the text as needed. You can access a text storage object from any thread of your app, but your app must guarantee access from only one thread at a time. In macOS, this class also defines properties for getting and setting scriptable attributes of objects. Unless you’re dealing with scriptability, you shouldn’t access these properties directly. In particular, using the , , or properties is an inefficient way to manipulate the text storage, since accessing these properties involves the creation of many objects. Instead, use the text access methods defined by , , , and to perform character-level manipulation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage
type TextStorage struct {
	foundation.MutableAttributedString
}

// TextStorageFrom constructs a [TextStorage] from an unsafe.Pointer.
//
// The fundamental storage mechanism of TextKit that contains the text managed by the system.
func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: foundation.MutableAttributedStringFrom(ptr),
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setCharacters:"), value)
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


