// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTable] class.
var textTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	class objc.Class
}

// An interface definition for the [TextTable] class.
type ITextTable interface {
	ITextBlock
}

// An object that represents a text table as a whole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable

type TextTable struct {
	TextBlock
}

// TextTableFrom constructs a [TextTable] from an unsafe.Pointer.
//
// An object that represents a text table as a whole.
func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: TextBlockFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextTableClass) New() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTable) Init() TextTable {
	rv := objc.Send[TextTable](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTable) Autorelease() TextTable {
	rv := objc.Send[TextTable](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTable creates a new TextTable instance.
func NewTextTable() TextTable {
	return textTableClass.New()
}




