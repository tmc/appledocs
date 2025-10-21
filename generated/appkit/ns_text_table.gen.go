// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTable] class.
var (
	TextTableClass     _TextTableClass
	TextTableClassOnce sync.Once
)

func getTextTableClass() _TextTableClass {
	TextTableClassOnce.Do(func() {
		TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}
	})
	return TextTableClass
}

type _TextTableClass struct {
	class objc.Class
}

// An interface definition for the [TextTable] class.
type ITextTable interface {
	ITextBlock
}

// An object that represents a text table as a whole.
//
// A text table is responsible for laying out and drawing the text table blocks it contains, and it maintains the basic parameters of the table.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getTextTableClass().New()
}




