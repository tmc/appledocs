// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SearchFieldCell] class.
var (
	SearchFieldCellClass     _SearchFieldCellClass
	SearchFieldCellClassOnce sync.Once
)

func getSearchFieldCellClass() _SearchFieldCellClass {
	SearchFieldCellClassOnce.Do(func() {
		SearchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}
	})
	return SearchFieldCellClass
}

type _SearchFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SearchFieldCell] class.
type ISearchFieldCell interface {
	ITextFieldCell
	// properties:
	CancelButtonCell() IButtonCell
	SetCancelButtonCell(value IButtonCell)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentSearches() []string
	SetRecentSearches(value []string)
	RecentsAutosaveName() objc.IObject /* cross-framework: SearchFieldRecentsAutosaveName */
	SetRecentsAutosaveName(value objc.IObject /* cross-framework: SearchFieldRecentsAutosaveName */)
	SearchButtonCell() IButtonCell
	SetSearchButtonCell(value IButtonCell)
	SearchMenuTemplate() IMenu
	SetSearchMenuTemplate(value IMenu)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	// methods:
	CancelButtonRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */
	ResetCancelButtonCell()
	ResetSearchButtonCell()
	SearchButtonRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */
	SearchTextRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */
}

// The programmatic interface for text fields that are used for text-based searches.
//
// The class defines the programmatic interface for text fields that are optimized for text-based searches. An object is “wrapped” by an control object, which directly inherits from the class. The search field implemented by these classes presents a standard user interface for searches, including a search button, a cancel button, and a pop-up icon menu for listing recent search strings and custom search categories. When the user types and then pauses, the cell’s action message is sent to its target. You can query the cell’s string value for the current text to search for. Do not rely on the sender of the action to be an object because the menu may change. If you need to change the menu, modify the search menu template and update the value in the property.


// The programmatic interface for text fields that are used for text-based searches.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/init(textCell:)
func NewSearchFieldCellTextCell(string_ objc.IObject /* cross-framework: NSString */) SearchFieldCell {
	instance := getSearchFieldCellClass().Alloc()
	rv := objc.Send[SearchFieldCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/init(coder:)
func NewSearchFieldCellWithCoder(coder foundation.Coder) SearchFieldCell {
	instance := getSearchFieldCellClass().Alloc()
	rv := objc.Send[SearchFieldCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Modifies the bounding rectangle for the cancel button cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/cancelButtonRect(forBounds:)
func (s_ SearchFieldCell) CancelButtonRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("cancelButtonRectForBounds:"), rect)
	return rv
}


// Resets the cancel button cell to its default attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/resetCancelButtonCell()
func (s_ SearchFieldCell) ResetCancelButtonCell() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetCancelButtonCell"))
}


// Resets the search button cell to its default attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/resetSearchButtonCell()
func (s_ SearchFieldCell) ResetSearchButtonCell() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetSearchButtonCell"))
}


// Modifies the bounding rectangle for the search button cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchButtonRect(forBounds:)
func (s_ SearchFieldCell) SearchButtonRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("searchButtonRectForBounds:"), rect)
	return rv
}


// Modifies the bounding rectangle for the search-text field cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchTextRect(forBounds:)
func (s_ SearchFieldCell) SearchTextRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("searchTextRectForBounds:"), rect)
	return rv
}


// The button cell used to display the cancel-button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/cancelButtonCell
func (s_ SearchFieldCell) CancelButtonCell() IButtonCell {
	rv := objc.Send[ButtonCell](s_.ID, objc.Sel("cancelButtonCell"))
	return rv
}


// The button cell used to display the cancel-button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/cancelButtonCell
func (s_ SearchFieldCell) SetCancelButtonCell(value IButtonCell) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCancelButtonCell:"), value)
}


// The maximum number of search strings that can appear in the search menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/maximumRecents
func (s_ SearchFieldCell) MaximumRecents() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumRecents"))
	return rv
}


// The maximum number of search strings that can appear in the search menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/maximumRecents
func (s_ SearchFieldCell) SetMaximumRecents(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumRecents:"), value)
}


// An array of the recent search strings to display in the pop-up icon menu of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentSearches
func (s_ SearchFieldCell) RecentSearches() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("recentSearches"))
	return rv
}


// An array of the recent search strings to display in the pop-up icon menu of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentSearches
func (s_ SearchFieldCell) SetRecentSearches(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecentSearches:"), nsArray)
}


// The autosave name under which the search field automatically saves the list of recent search strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentsAutosaveName
func (s_ SearchFieldCell) RecentsAutosaveName() objc.IObject /* cross-framework: SearchFieldRecentsAutosaveName */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("recentsAutosaveName"))
	return rv
}


// The autosave name under which the search field automatically saves the list of recent search strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentsAutosaveName
func (s_ SearchFieldCell) SetRecentsAutosaveName(value objc.IObject /* cross-framework: SearchFieldRecentsAutosaveName */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecentsAutosaveName:"), value)
}


// The button cell used to display the search-button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchButtonCell
func (s_ SearchFieldCell) SearchButtonCell() IButtonCell {
	rv := objc.Send[ButtonCell](s_.ID, objc.Sel("searchButtonCell"))
	return rv
}


// The button cell used to display the search-button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchButtonCell
func (s_ SearchFieldCell) SetSearchButtonCell(value IButtonCell) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchButtonCell:"), value)
}


// The menu object used to dynamically construct the search field’s pop-up icon menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchMenuTemplate
func (s_ SearchFieldCell) SearchMenuTemplate() IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("searchMenuTemplate"))
	return rv
}


// The menu object used to dynamically construct the search field’s pop-up icon menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchMenuTemplate
func (s_ SearchFieldCell) SetSearchMenuTemplate(value IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchMenuTemplate:"), value)
}


// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsSearchStringImmediately
func (s_ SearchFieldCell) SendsSearchStringImmediately() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsSearchStringImmediately"))
	return rv
}


// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsSearchStringImmediately
func (s_ SearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsSearchStringImmediately:"), value)
}


// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsWholeSearchString
func (s_ SearchFieldCell) SendsWholeSearchString() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsWholeSearchString"))
	return rv
}


// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsWholeSearchString
func (s_ SearchFieldCell) SetSendsWholeSearchString(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsWholeSearchString:"), value)
}


