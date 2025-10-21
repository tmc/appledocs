// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextList] class.
var (
	TextListClass     _TextListClass
	TextListClassOnce sync.Once
)

func getTextListClass() _TextListClass {
	TextListClassOnce.Do(func() {
		TextListClass = _TextListClass{objc.GetClass("NSTextList")}
	})
	return TextListClass
}

type _TextListClass struct {
	class objc.Class
}

// An interface definition for the [TextList] class.
type ITextList interface {
	objectivec.IObject
}

// A section of text that forms a single list.
//
// The visible elements of the list, including list markers, appear in the text as they do for lists created by hand. The list object, however, allows the list to be recognized as such by the text system. This enables automatic creation of markers and spacing. Text lists are used in HTML import and export. Text lists appear as attributes on paragraphs, as part of the paragraph style. An may have an array of text lists, representing the nested lists containing the paragraph, in order from outermost to innermost. For example, if list1 contains four paragraphs, the middle two of which are also in the inner list2, then the text lists array for the first and fourth paragraphs is (list1), while the text lists array for the second and third paragraphs is (list1, list2). The methods implementing this are on , and on . In addition, has convenience methods for lists, such as , which determines the range covered by a list, and , which determines the ordinal position within a list of a particular item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextList
type TextList struct {
	objectivec.Object
}

// TextListFrom constructs a [TextList] from an unsafe.Pointer.
//
// A section of text that forms a single list.
func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextListClass) Alloc() TextList {
	rv := objc.Send[TextList](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextListClass) New() TextList {
	rv := objc.Send[TextList](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextList) Init() TextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextList) Autorelease() TextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextList creates a new TextList instance.
func NewTextList() TextList {
	return getTextListClass().New()
}


// The text lists that contain the paragraph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (t_ TextList) TextLists() NSTextList {
	rv := objc.Send[NSTextList](t_.ID, objc.Sel("textLists"))
	return rv
}


// SetTextLists sets the value of the textLists property.
// The text lists that contain the paragraph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/textlists
func (t_ TextList) SetTextLists(value ITextList) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLists:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/isordered
func (t_ TextList) IsOrdered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isOrdered"))
	return rv
}


// SetIsOrdered sets the value of the isOrdered property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/isordered
func (t_ TextList) SetIsOrdered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsOrdered:"), value)
}

// Returns the list options mask value of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/listoptions
func (t_ TextList) ListOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("listOptions"))
	return rv
}


// SetListOptions sets the value of the listOptions property.
// Returns the list options mask value of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/listoptions
func (t_ TextList) SetListOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setListOptions:"), value)
}

// Returns the marker format string used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/markerformat-swift.property
func (t_ TextList) MarkerFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("markerFormat"))
	return rv
}


// SetMarkerFormat sets the value of the markerFormat property.
// Returns the marker format string used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/markerformat-swift.property
func (t_ TextList) SetMarkerFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkerFormat:"), value)
}

// Sets the starting item number for the text list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/startingitemnumber
func (t_ TextList) StartingItemNumber() int {
	rv := objc.Send[int](t_.ID, objc.Sel("startingItemNumber"))
	return rv
}


// SetStartingItemNumber sets the value of the startingItemNumber property.
// Sets the starting item number for the text list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/startingitemnumber
func (t_ TextList) SetStartingItemNumber(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartingItemNumber:"), value)
}



