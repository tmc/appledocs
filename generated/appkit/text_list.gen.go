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
	textListClass     _TextListClass
	textListClassOnce sync.Once
)

func getTextListClass() _TextListClass {
	textListClassOnce.Do(func() {
		textListClass = _TextListClass{objc.GetClass("NSTextList")}
	})
	return textListClass
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




