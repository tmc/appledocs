// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchToolbarItem] class.
var (
	searchToolbarItemClass     _SearchToolbarItemClass
	searchToolbarItemClassOnce sync.Once
)

func getSearchToolbarItemClass() _SearchToolbarItemClass {
	searchToolbarItemClassOnce.Do(func() {
		searchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}
	})
	return searchToolbarItemClass
}

type _SearchToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	IToolbarItem
}

// A toolbar item that contains a search field optimized for performing text-based searches.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem
type SearchToolbarItem struct {
	ToolbarItem
}

// SearchToolbarItemFrom constructs a [SearchToolbarItem] from an unsafe.Pointer.
//
// A toolbar item that contains a search field optimized for performing text-based searches.
func SearchToolbarItemFrom(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SearchToolbarItem) Autorelease() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSearchToolbarItem creates a new SearchToolbarItem instance.
func NewSearchToolbarItem() SearchToolbarItem {
	return getSearchToolbarItemClass().New()
}




