// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pasteboard] class.
var (
	PasteboardClass     _PasteboardClass
	PasteboardClassOnce sync.Once
)

func getPasteboardClass() _PasteboardClass {
	PasteboardClassOnce.Do(func() {
		PasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}
	})
	return PasteboardClass
}

type _PasteboardClass struct {
	class objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	objectivec.IObject
	ClearContents() int
	DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	IndexOfPasteboardItem(pasteboardItem unsafe.Pointer) uint
	ReadObjectsForClassesOptions(classArray unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer
	SetDataForType(data unsafe.Pointer, dataType unsafe.Pointer) bool
	SetPropertyListForType(plist objc.ID, dataType unsafe.Pointer) bool
	SetStringForType(string_ string, dataType unsafe.Pointer) bool
	StringForType(dataType unsafe.Pointer) unsafe.Pointer
	WriteObjects(objects unsafe.Pointer) bool
}

// An object that transfers data to and from the pasteboard server.
//
// The pasteboard server is shared by all running apps. It contains data that the user has cut or copied, as well as other data that one application wants to transfer to another. objects are an application’s sole interface to the server and to all pasteboard operations. An object is also used to transfer data between apps and service providers listed in each application’s Services menu. The drag pasteboard is used to transfer data that is being dragged by the user. A pasteboard can contain multiple items. You can directly write or read any object that implements the or   respectively. This allows you to write and read common items such as URLs, colors, images, strings, attributed strings, and sounds without an intermediary object. Your custom classes can also implement these protocols for use with the pasteboard. Writing methods such as provide a convenient means of writing to the first pasteboard item, without having to create the first pasteboard item. You can use code like this, for example: The general pasteboard, available by way of the class method, automatically participates with the Universal Clipboard feature in macOS 10.12 and later and in iOS 10.0 and later. There is no macOS API for interacting with this feature.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard
type Pasteboard struct {
	objectivec.Object
}

// PasteboardFrom constructs a [Pasteboard] from an unsafe.Pointer.
//
// An object that transfers data to and from the pasteboard server.
func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pasteboard) Autorelease() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasteboard creates a new Pasteboard instance.
func NewPasteboard() Pasteboard {
	return getPasteboardClass().New()
}


// Clears the existing contents of the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/clearContents()
func (p_ Pasteboard) ClearContents() int {
	rv := objc.Send[int](p_.ID, objc.Sel("clearContents"))
	return rv
}

// Determines whether the first pasteboard item matches the specified patterns, without notifying the person using the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/detectPatternsForPatterns:completionHandler:
func (p_ Pasteboard) DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectPatternsForPatterns:completionHandler:"), patterns, completionHandler)
}

// Determines whether the first pasteboard item matches the specified patterns, reading the contents if it finds a match.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/detectValuesForPatterns:completionHandler:
func (p_ Pasteboard) DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectValuesForPatterns:completionHandler:"), patterns, completionHandler)
}

// Returns the index of the specified pasteboard item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/index(of:)
func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem unsafe.Pointer) uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("indexOfPasteboardItem:"), pasteboardItem)
	return rv
}

// Reads from the receiver objects that best match the specified array of classes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/readObjects(forClasses:options:)
func (p_ Pasteboard) ReadObjectsForClassesOptions(classArray unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("readObjectsForClasses:options:"), classArray, options)
	return rv
}

// Sets the data as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setData(_:forType:)
func (p_ Pasteboard) SetDataForType(data unsafe.Pointer, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setData:forType:"), data, dataType)
	return rv
}

// Sets the given property list as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setPropertyList(_:forType:)
func (p_ Pasteboard) SetPropertyListForType(plist objc.ID, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setPropertyList:forType:"), plist, dataType)
	return rv
}

// Sets the given string as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string_ string, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), objc.String(string_), dataType)
	return rv
}

// Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/string(forType:)
func (p_ Pasteboard) StringForType(dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringForType:"), dataType)
	return rv
}

// Writes an array of objects to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/writeObjects(_:)
func (p_ Pasteboard) WriteObjects(objects unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeObjects:"), objects)
	return rv
}

// The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/accessBehavior-86972
func (p_ Pasteboard) AccessBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("accessBehavior"))
	return rv
}

// An array that contains all the items held by the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/pasteboardItems
func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := objc.Send[[]PasteboardItem](p_.ID, objc.Sel("pasteboardItems"))
	return rv
}



