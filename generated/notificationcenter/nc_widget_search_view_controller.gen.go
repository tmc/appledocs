// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)





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





// An interface definition for the [NCWidgetSearchViewController] class.
type INCWidgetSearchViewController interface {
	appkit.IViewController
	

	// properties:
	SearchDescription() foundation.foundation.INSString
	SetSearchDescription(value foundation.foundation.INSString)
	SearchResultKeyPath() foundation.foundation.INSString
	SetSearchResultKeyPath(value foundation.foundation.INSString)
	SearchResults() []objc.ID
	SetSearchResults(value []objc.ID)
	SearchResultsPlaceholderString() foundation.foundation.INSString
	SetSearchResultsPlaceholderString(value foundation.foundation.INSString)


	

	// methods:


}





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

























// A localized description of the nature of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchDescription
func (n_ NCWidgetSearchViewController) SearchDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchDescription"))
	return rv
}


// A localized description of the nature of the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchDescription
func (n_ NCWidgetSearchViewController) SetSearchDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchDescription:"), value)
}


// A key path for the string property to display for each object in the search results array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultKeyPath
func (n_ NCWidgetSearchViewController) SearchResultKeyPath() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchResultKeyPath"))
	return rv
}


// A key path for the string property to display for each object in the search results array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultKeyPath
func (n_ NCWidgetSearchViewController) SetSearchResultKeyPath(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchResultKeyPath:"), value)
}


// An array of search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResults
func (n_ NCWidgetSearchViewController) SearchResults() []objc.ID {
	rv := objc.Send[[]objc.ID](n_.ID, objc.Sel("searchResults"))
	return rv
}


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
}


// A localized phrase displayed in the results list when no search results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultsPlaceholderString
func (n_ NCWidgetSearchViewController) SearchResultsPlaceholderString() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchResultsPlaceholderString"))
	return rv
}


// A localized phrase displayed in the results list when no search results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetSearchViewController/searchResultsPlaceholderString
func (n_ NCWidgetSearchViewController) SetSearchResultsPlaceholderString(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchResultsPlaceholderString:"), value)
}










