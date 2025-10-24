// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSearchToolbarItem */


/* debug [class_header]: Header for NSSearchToolbarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SearchToolbarItem */
// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	IToolbarItem
	
/* debug [class_interface_properties]: Properties for SearchToolbarItem */
	// properties:
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	SearchField() ISearchField
	SetSearchField(value ISearchField)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SearchToolbarItem */
	// methods:
	BeginSearchInteraction()
	EndSearchInteraction()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SearchToolbarItem */
// Alloc allocates a new instance without initialization.
func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SearchToolbarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SearchToolbarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SearchToolbarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SearchToolbarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SearchToolbarItem */

// Starts a search interaction and moves the keyboard focus to the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/beginSearchInteraction()
func (s_ SearchToolbarItem) BeginSearchInteraction() {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSearchInteraction"))
}/* debug [instance_methods/method]: BeginSearchInteraction */


// Ends a search interaction by giving up the first responder and adjusting the size of the search field to the available width for the toolbar item if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/endSearchInteraction()
func (s_ SearchToolbarItem) EndSearchInteraction() {
	objc.Send[objc.ID](s_.ID, objc.Sel("endSearchInteraction"))
}/* debug [instance_methods/method]: EndSearchInteraction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SearchToolbarItem */

// The preferred width for the toolbar item when it has keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/preferredWidthForSearchField
func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preferredWidthForSearchField"))
	return rv
}/* debug [instance_properties/getter]: preferredWidthForSearchField */


// The preferred width for the toolbar item when it has keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/preferredWidthForSearchField
func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredWidthForSearchField:"), value)
}/* debug [instance_properties/setter]: preferredWidthForSearchField */


// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/resignsFirstResponderWithCancel
func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("resignsFirstResponderWithCancel"))
	return rv
}/* debug [instance_properties/getter]: resignsFirstResponderWithCancel */


// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/resignsFirstResponderWithCancel
func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResignsFirstResponderWithCancel:"), value)
}/* debug [instance_properties/setter]: resignsFirstResponderWithCancel */


// The search field inside the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/searchField
func (s_ SearchToolbarItem) SearchField() ISearchField {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("searchField"))
	return rv
}/* debug [instance_properties/getter]: searchField */


// The search field inside the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/searchField
func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchField:"), value)
}/* debug [instance_properties/setter]: searchField */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSearchToolbarItem */



