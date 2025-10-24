// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NCWidgetSearchViewController */


/* debug [class_header]: Header for NCWidgetSearchViewController */
// The class instance for the [NCWidgetSearchViewController] class.
var (
	NCWidgetSearchViewControllerClass     _NCWidgetSearchViewControllerClass
	NCWidgetSearchViewControllerClassOnce sync.Once
)

func getNCWidgetSearchViewControllerClass() _NCWidgetSearchViewControllerClass {
	NCWidgetSearchViewControllerClassOnce.Do(func() {
		NCWidgetSearchViewControllerClass = _NCWidgetSearchViewControllerClass{objc.GetClass("NCWidgetSearchViewController")}
	})
	return NCWidgetSearchViewControllerClass
}

type _NCWidgetSearchViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NCWidgetSearchViewController */
// An interface definition for the [NCWidgetSearchViewController] class.
type INCWidgetSearchViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for NCWidgetSearchViewController */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	SearchDescription() objc.IObject /* cross-framework: NSString */
	SetSearchDescription(value objc.IObject /* cross-framework: NSString */)
	SearchResultKeyPath() objc.IObject /* cross-framework: NSString */
	SetSearchResultKeyPath(value objc.IObject /* cross-framework: NSString */)
	SearchResults() []objc.ID
	SetSearchResults(value []objc.ID)
	SearchResultsPlaceholderString() objc.IObject /* cross-framework: NSString */
	SetSearchResultsPlaceholderString(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NCWidgetSearchViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NCWidgetSearchViewController */
// Alloc allocates a new instance without initialization.
func (nc _NCWidgetSearchViewControllerClass) Alloc() NCWidgetSearchViewController {
	rv := objc.Send[NCWidgetSearchViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NCWidgetSearchViewControllerClass) New() NCWidgetSearchViewController {
	rv := objc.Send[NCWidgetSearchViewController](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NCWidgetSearchViewController) Init() NCWidgetSearchViewController {
	rv := objc.Send[NCWidgetSearchViewController](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NCWidgetSearchViewController) Autorelease() NCWidgetSearchViewController {
	rv := objc.Send[NCWidgetSearchViewController](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNCWidgetSearchViewController creates a new NCWidgetSearchViewController instance.
func NewNCWidgetSearchViewController() NCWidgetSearchViewController {
	return getNCWidgetSearchViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NCWidgetSearchViewController */
// An object that provides a default search view within a macOS Today widget.
//
// The class provides a default search view within a Today widget. A search view controller works together with its delegate to perform searches on the user’s input and display results from which a user can choose. To learn about the search view controller delegate methods, see . When a widget is in editing mode, it can enable search for new content by instantiating an object and presenting it using . The search view controller displays the default search field and a list of results. It uses its to perform the search itself.


// An object that provides a default search view within a macOS Today widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController
type NCWidgetSearchViewController struct {
	appkit.ViewController
}

// NCWidgetSearchViewControllerFrom constructs a [NCWidgetSearchViewController] from an unsafe.Pointer.
//
// An object that provides a default search view within a macOS Today widget.
func NCWidgetSearchViewControllerFrom(ptr unsafe.Pointer) NCWidgetSearchViewController {
	return NCWidgetSearchViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NCWidgetSearchViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NCWidgetSearchViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NCWidgetSearchViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NCWidgetSearchViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NCWidgetSearchViewController */

// The search view controller’s delegate or if the receiver doesn’t have a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/delegate
func (n_ NCWidgetSearchViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The search view controller’s delegate or if the receiver doesn’t have a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/delegate
func (n_ NCWidgetSearchViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A localized description of the nature of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchDescription
func (n_ NCWidgetSearchViewController) SearchDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchDescription"))
	return rv
}/* debug [instance_properties/getter]: searchDescription */


// A localized description of the nature of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchDescription
func (n_ NCWidgetSearchViewController) SetSearchDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchDescription:"), value)
}/* debug [instance_properties/setter]: searchDescription */


// A key path for the string property to display for each object in the search results array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultKeyPath
func (n_ NCWidgetSearchViewController) SearchResultKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchResultKeyPath"))
	return rv
}/* debug [instance_properties/getter]: searchResultKeyPath */


// A key path for the string property to display for each object in the search results array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultKeyPath
func (n_ NCWidgetSearchViewController) SetSearchResultKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchResultKeyPath:"), value)
}/* debug [instance_properties/setter]: searchResultKeyPath */


// An array of search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResults
func (n_ NCWidgetSearchViewController) SearchResults() []objc.ID {
	rv := objc.Send[[]objc.ID](n_.ID, objc.Sel("searchResults"))
	return rv
}/* debug [instance_properties/getter]: searchResults */


// An array of search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResults
func (n_ NCWidgetSearchViewController) SetSearchResults(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchResults:"), nsArray)
}/* debug [instance_properties/setter]: searchResults */


// A localized phrase displayed in the results list when no search results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultsPlaceholderString
func (n_ NCWidgetSearchViewController) SearchResultsPlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchResultsPlaceholderString"))
	return rv
}/* debug [instance_properties/getter]: searchResultsPlaceholderString */


// A localized phrase displayed in the results list when no search results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultsPlaceholderString
func (n_ NCWidgetSearchViewController) SetSearchResultsPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchResultsPlaceholderString:"), value)
}/* debug [instance_properties/setter]: searchResultsPlaceholderString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NCWidgetSearchViewController */





