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

// The programmatic interface for text fields that are used for text-based searches.
//
// The class defines the programmatic interface for text fields that are optimized for text-based searches. An object is “wrapped” by an control object, which directly inherits from the class. The search field implemented by these classes presents a standard user interface for searches, including a search button, a cancel button, and a pop-up icon menu for listing recent search strings and custom search categories. When the user types and then pauses, the cell’s action message is sent to its target. You can query the cell’s string value for the current text to search for. Do not rely on the sender of the action to be an object because the menu may change. If you need to change the menu, modify the search menu template and update the value in the property.
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




