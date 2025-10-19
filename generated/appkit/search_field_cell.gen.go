// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchFieldCell] class.
var (
	searchFieldCellClass     _SearchFieldCellClass
	searchFieldCellClassOnce sync.Once
)

func getSearchFieldCellClass() _SearchFieldCellClass {
	searchFieldCellClassOnce.Do(func() {
		searchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}
	})
	return searchFieldCellClass
}

type _SearchFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SearchFieldCell] class.
type ISearchFieldCell interface {
	ITextFieldCell
}

// The programmatic interface for text fields that are used for text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell

type SearchFieldCell struct {
	TextFieldCell
}

// SearchFieldCellFrom constructs a [SearchFieldCell] from an unsafe.Pointer.
//
// The programmatic interface for text fields that are used for text-based searches.
func SearchFieldCellFrom(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SearchFieldCell) Autorelease() SearchFieldCell {
	rv := objc.Send[SearchFieldCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSearchFieldCell creates a new SearchFieldCell instance.
func NewSearchFieldCell() SearchFieldCell {
	return getSearchFieldCellClass().New()
}




