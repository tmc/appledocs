// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SearchField] class.
var (
	SearchFieldClass     _SearchFieldClass
	SearchFieldClassOnce sync.Once
)

func getSearchFieldClass() _SearchFieldClass {
	SearchFieldClassOnce.Do(func() {
		SearchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}
	})
	return SearchFieldClass
}

type _SearchFieldClass struct {
	class objc.Class
}

// An interface definition for the [SearchField] class.
type ISearchField interface {
	ITextField
	RectForCancelButtonWhenCentered(isCentered bool) coregraphics.CGRect
	RectForSearchButtonWhenCentered(isCentered bool) coregraphics.CGRect
	RectForSearchTextWhenCentered(isCentered bool) coregraphics.CGRect
	CancelButtonBounds() coregraphics.CGRect
	CentersPlaceholder() bool
	SetCentersPlaceholder(value bool)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentSearches() []string
	SetRecentSearches(value []string)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value ISearchFieldRecentsAutosaveName)
	SearchButtonBounds() coregraphics.CGRect
	SearchMenuTemplate() NSMenu
	SetSearchMenuTemplate(value IMenu)
	SearchTextBounds() coregraphics.CGRect
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
}

// A text field optimized for performing text-based searches.
//
// provides a customized text field for entering search data. The class also provides a search button, a cancel button, and a pop-up icon menu for listing recent search strings and custom search categories. An object wraps an object. The cell provides access to most search field attributes and a comprehensive programmatic interface for manipulating the search field. You can use an object to manipulate some aspects of the search field. For additional information about search fields and how to implement them, see the class.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField
type SearchField struct {
	TextField
}

// SearchFieldFrom constructs a [SearchField] from an unsafe.Pointer.
//
// A text field optimized for performing text-based searches.
func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		TextField: TextFieldFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SearchFieldClass) New() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SearchField) Init() SearchField {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SearchField) Autorelease() SearchField {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSearchField creates a new SearchField instance.
func NewSearchField() SearchField {
	return getSearchFieldClass().New()
}


// The rectangle for the cancel button within the bounds of the search field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/rectForCancelButton(whenCentered:)
func (s_ SearchField) RectForCancelButtonWhenCentered(isCentered bool) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("rectForCancelButtonWhenCentered:"), isCentered)
	return rv
}

// The rectangle for the search button within the bounds of the search field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/rectForSearchButton(whenCentered:)
func (s_ SearchField) RectForSearchButtonWhenCentered(isCentered bool) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("rectForSearchButtonWhenCentered:"), isCentered)
	return rv
}

// The rectangle for the search text within the bounds of the field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/rectForSearchText(whenCentered:)
func (s_ SearchField) RectForSearchTextWhenCentered(isCentered bool) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("rectForSearchTextWhenCentered:"), isCentered)
	return rv
}

// The rectangle for the cancel button within the bounds of the search field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/cancelButtonBounds
func (s_ SearchField) CancelButtonBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("cancelButtonBounds"))
	return rv
}

// A Boolean value that determines whether the search field’s components are centered within the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/centersPlaceholder
func (s_ SearchField) CentersPlaceholder() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("centersPlaceholder"))
	return rv
}


// SetCentersPlaceholder sets the value of the centersPlaceholder property.
// A Boolean value that determines whether the search field’s components are centered within the control.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/centersPlaceholder
func (s_ SearchField) SetCentersPlaceholder(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCentersPlaceholder:"), value)
}

// The delegate for the search field, or if the search field doesn’t have a delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/delegate
func (s_ SearchField) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the search field, or if the search field doesn’t have a delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/delegate
func (s_ SearchField) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// The maximum number of search strings that can appear in the search menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/maximumRecents
func (s_ SearchField) MaximumRecents() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumRecents"))
	return rv
}


// SetMaximumRecents sets the value of the maximumRecents property.
// The maximum number of search strings that can appear in the search menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/maximumRecents
func (s_ SearchField) SetMaximumRecents(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumRecents:"), value)
}

// The list of recent search strings for the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentSearches
func (s_ SearchField) RecentSearches() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("recentSearches"))
	return rv
}


// SetRecentSearches sets the value of the recentSearches property.
// The list of recent search strings for the control.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentSearches
func (s_ SearchField) SetRecentSearches(value []string) {
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

// The name under which the search field automatically archives the list of recent search strings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsAutosaveName-swift.property
func (s_ SearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.Send[SearchFieldRecentsAutosaveName](s_.ID, objc.Sel("recentsAutosaveName"))
	return rv
}


// SetRecentsAutosaveName sets the value of the recentsAutosaveName property.
// The name under which the search field automatically archives the list of recent search strings.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsAutosaveName-swift.property
func (s_ SearchField) SetRecentsAutosaveName(value ISearchFieldRecentsAutosaveName) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecentsAutosaveName:"), value)
}

// The rectangle for the search button within the bounds of the search field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchButtonBounds
func (s_ SearchField) SearchButtonBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("searchButtonBounds"))
	return rv
}

// The menu object used to dynamically construct the search field’s pop-up icon menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchMenuTemplate
func (s_ SearchField) SearchMenuTemplate() NSMenu {
	rv := objc.Send[NSMenu](s_.ID, objc.Sel("searchMenuTemplate"))
	return rv
}


// SetSearchMenuTemplate sets the value of the searchMenuTemplate property.
// The menu object used to dynamically construct the search field’s pop-up icon menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchMenuTemplate
func (s_ SearchField) SetSearchMenuTemplate(value IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchMenuTemplate:"), value)
}

// The rectangle for the search text within the bounds of the search field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchTextBounds
func (s_ SearchField) SearchTextBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("searchTextBounds"))
	return rv
}

// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsSearchStringImmediately
func (s_ SearchField) SendsSearchStringImmediately() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsSearchStringImmediately"))
	return rv
}


// SetSendsSearchStringImmediately sets the value of the sendsSearchStringImmediately property.
// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsSearchStringImmediately
func (s_ SearchField) SetSendsSearchStringImmediately(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsSearchStringImmediately:"), value)
}

// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsWholeSearchString
func (s_ SearchField) SendsWholeSearchString() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsWholeSearchString"))
	return rv
}


// SetSendsWholeSearchString sets the value of the sendsWholeSearchString property.
// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsWholeSearchString
func (s_ SearchField) SetSendsWholeSearchString(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsWholeSearchString:"), value)
}



