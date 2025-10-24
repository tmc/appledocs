// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchToolbarItem] class.
var (
	SearchToolbarItemClass     _SearchToolbarItemClass
	SearchToolbarItemClassOnce sync.Once
)

func getSearchToolbarItemClass() _SearchToolbarItemClass {
	SearchToolbarItemClassOnce.Do(func() {
		SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}
	})
	return SearchToolbarItemClass
}

type _SearchToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	IToolbarItem
	// properties:
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	SearchField() objc.IObject /* cross-framework: SearchField */
	SetSearchField(value objc.IObject /* cross-framework: SearchField */)
	// methods:
	BeginSearchInteraction()
}

// A toolbar item that contains a search field optimized for performing text-based searches.
//
// automatically resizes to accommodate typing when the focus switches to the toolbar item. When the toolbar is low on space, the system may collapse the search item into a button representation, which then expands to a full search field when the user clicks on it.


// A toolbar item that contains a search field optimized for performing text-based searches.
//
// [Full Topic]
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



// Starts a search interaction and moves the keyboard focus to the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/beginSearchInteraction()
func (s_ SearchToolbarItem) BeginSearchInteraction() {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSearchInteraction"))
}


// The preferred width for the toolbar item when it has keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/preferredwidthforsearchfield
func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preferredWidthForSearchField"))
	return rv
}


// The preferred width for the toolbar item when it has keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/preferredwidthforsearchfield
func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredWidthForSearchField:"), value)
}


// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/resignsfirstresponderwithcancel
func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("resignsFirstResponderWithCancel"))
	return rv
}


// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/resignsfirstresponderwithcancel
func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResignsFirstResponderWithCancel:"), value)
}


// The search field inside the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/searchfield
func (s_ SearchToolbarItem) SearchField() objc.IObject /* cross-framework: SearchField */ {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("searchField"))
	return rv
}


// The search field inside the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/searchfield
func (s_ SearchToolbarItem) SetSearchField(value objc.IObject /* cross-framework: SearchField */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchField:"), value)
}



