// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextAttachmentCell] class.
var (
	TextAttachmentCellClass     _TextAttachmentCellClass
	TextAttachmentCellClassOnce sync.Once
)

func getTextAttachmentCellClass() _TextAttachmentCellClass {
	TextAttachmentCellClassOnce.Do(func() {
		TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}
	})
	return TextAttachmentCellClass
}

type _TextAttachmentCellClass struct {
	class objc.Class
}

// An interface definition for the [TextAttachmentCell] class.
type ITextAttachmentCell interface {
	ICell
}

// An object that implements the functionality of the text attachment cell protocol.
//
// This specification describes only those methods whose implementations have features that are particular to this class. For a general discussion of the protocol’s methods, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type TextAttachmentCell struct {
	Cell
}

// TextAttachmentCellFrom constructs a [TextAttachmentCell] from an unsafe.Pointer.
//
// An object that implements the functionality of the text attachment cell protocol.
func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: CellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAttachmentCell) Autorelease() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAttachmentCell creates a new TextAttachmentCell instance.
func NewTextAttachmentCell() TextAttachmentCell {
	return getTextAttachmentCellClass().New()
}




