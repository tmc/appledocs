// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSearchField */


/* debug [class_header]: Header for NSSearchField */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SearchField */
// An interface definition for the [SearchField] class.
type ISearchField interface {
	ITextField
	
/* debug [class_interface_properties]: Properties for SearchField */
	// properties:
	CancelButtonBounds() Rect /* not a class type */
	CentersPlaceholder() bool
	SetCentersPlaceholder(value bool)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName /* typedef */
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName /* typedef */)
	RecentSearches() []string
	SetRecentSearches(value []string)
	SearchButtonBounds() Rect /* not a class type */
	SearchMenuTemplate() IMenu
	SetSearchMenuTemplate(value IMenu)
	SearchTextBounds() Rect /* not a class type */
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SearchField */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SearchField */
// Alloc allocates a new instance without initialization.
func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SearchField */
// A text field optimized for performing text-based searches.
//
// provides a customized text field for entering search data. The class also provides a search button, a cancel button, and a pop-up icon menu for listing recent search strings and custom search categories. An object wraps an object. The cell provides access to most search field attributes and a comprehensive programmatic interface for manipulating the search field. You can use an object to manipulate some aspects of the search field. For additional information about search fields and how to implement them, see the class.


// A text field optimized for performing text-based searches.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SearchField *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SearchField */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SearchField */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SearchField */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SearchField */

// The rectangle for the cancel button within the bounds of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/cancelButtonBounds
func (s_ SearchField) CancelButtonBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("cancelButtonBounds"))
	return rv
}/* debug [instance_properties/getter]: cancelButtonBounds */


// A Boolean value that determines whether the search field’s components are centered within the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/centersPlaceholder
func (s_ SearchField) CentersPlaceholder() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("centersPlaceholder"))
	return rv
}/* debug [instance_properties/getter]: centersPlaceholder */


// A Boolean value that determines whether the search field’s components are centered within the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/centersPlaceholder
func (s_ SearchField) SetCentersPlaceholder(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCentersPlaceholder:"), value)
}/* debug [instance_properties/setter]: centersPlaceholder */


// The delegate for the search field, or if the search field doesn’t have a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/delegate
func (s_ SearchField) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the search field, or if the search field doesn’t have a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/delegate
func (s_ SearchField) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The maximum number of search strings that can appear in the search menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/maximumRecents
func (s_ SearchField) MaximumRecents() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumRecents"))
	return rv
}/* debug [instance_properties/getter]: maximumRecents */


// The maximum number of search strings that can appear in the search menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/maximumRecents
func (s_ SearchField) SetMaximumRecents(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumRecents:"), value)
}/* debug [instance_properties/setter]: maximumRecents */


// The name under which the search field automatically archives the list of recent search strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsAutosaveName-swift.property
func (s_ SearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("recentsAutosaveName"))
	return rv
}/* debug [instance_properties/getter]: recentsAutosaveName */


// The name under which the search field automatically archives the list of recent search strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsAutosaveName-swift.property
func (s_ SearchField) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecentsAutosaveName:"), value)
}/* debug [instance_properties/setter]: recentsAutosaveName */


// The list of recent search strings for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentSearches
func (s_ SearchField) RecentSearches() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("recentSearches"))
	return rv
}/* debug [instance_properties/getter]: recentSearches */


// The list of recent search strings for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/recentSearches
func (s_ SearchField) SetRecentSearches(value []string) {
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
}/* debug [instance_properties/setter]: recentSearches */


// The rectangle for the search button within the bounds of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchButtonBounds
func (s_ SearchField) SearchButtonBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("searchButtonBounds"))
	return rv
}/* debug [instance_properties/getter]: searchButtonBounds */


// The menu object used to dynamically construct the search field’s pop-up icon menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchMenuTemplate
func (s_ SearchField) SearchMenuTemplate() IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("searchMenuTemplate"))
	return rv
}/* debug [instance_properties/getter]: searchMenuTemplate */


// The menu object used to dynamically construct the search field’s pop-up icon menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchMenuTemplate
func (s_ SearchField) SetSearchMenuTemplate(value IMenu) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSearchMenuTemplate:"), value)
}/* debug [instance_properties/setter]: searchMenuTemplate */


// The rectangle for the search text within the bounds of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/searchTextBounds
func (s_ SearchField) SearchTextBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("searchTextBounds"))
	return rv
}/* debug [instance_properties/getter]: searchTextBounds */


// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsSearchStringImmediately
func (s_ SearchField) SendsSearchStringImmediately() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsSearchStringImmediately"))
	return rv
}/* debug [instance_properties/getter]: sendsSearchStringImmediately */


// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsSearchStringImmediately
func (s_ SearchField) SetSendsSearchStringImmediately(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsSearchStringImmediately:"), value)
}/* debug [instance_properties/setter]: sendsSearchStringImmediately */


// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsWholeSearchString
func (s_ SearchField) SendsWholeSearchString() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("sendsWholeSearchString"))
	return rv
}/* debug [instance_properties/getter]: sendsWholeSearchString */


// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsWholeSearchString
func (s_ SearchField) SetSendsWholeSearchString(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendsWholeSearchString:"), value)
}/* debug [instance_properties/setter]: sendsWholeSearchString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSearchField */



