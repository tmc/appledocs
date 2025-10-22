// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [NCWidgetListViewController] class.
var (
	NCWidgetListViewControllerClass     _NCWidgetListViewControllerClass
	NCWidgetListViewControllerClassOnce sync.Once
)

func getNCWidgetListViewControllerClass() _NCWidgetListViewControllerClass {
	NCWidgetListViewControllerClassOnce.Do(func() {
		NCWidgetListViewControllerClass = _NCWidgetListViewControllerClass{objc.GetClass("NCWidgetListViewController")}
	})
	return NCWidgetListViewControllerClass
}

type _NCWidgetListViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [NCWidgetListViewController] class.
type INCWidgetListViewController interface {
	appkit.IViewController
	RowForViewController(viewController appkit.IViewController) uint
	ViewControllerAtRowMakeIfNecessary(row uint, makeIfNecesary bool) appkit.ViewController
	Contents() []objc.ID
	SetContents(value []objc.ID)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Editing() bool
	SetEditing(value bool)
	HasDividerLines() bool
	SetHasDividerLines(value bool)
	MinimumVisibleRowCount() uint
	SetMinimumVisibleRowCount(value uint)
	ShowsAddButtonWhenEditing() bool
	SetShowsAddButtonWhenEditing(value bool)
}

// An object that provides a list view for displaying content in a macOS Today widget.
//
// The class provides a list view for displaying content in a Today widget. A list view controller works together with its delegate to display content and support user interaction with the list. To learn about the list view controller delegate methods, see . You store the contents of a widget as an array of objects in the list view controller’s property. To display the objects, you use a object, which provides a custom view controller for each object in . A list view controller also provides properties that make it easy to specify aspects of the list’s appearance and behavior, such as the number of rows to display, the presence of divider lines, and the ability to edit the list.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController
type NCWidgetListViewController struct {
	appkit.ViewController
}

// NCWidgetListViewControllerFrom constructs a [NCWidgetListViewController] from an unsafe.Pointer.
//
// An object that provides a list view for displaying content in a macOS Today widget.
func NCWidgetListViewControllerFrom(ptr unsafe.Pointer) NCWidgetListViewController {
	return NCWidgetListViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NCWidgetListViewControllerClass) Alloc() NCWidgetListViewController {
	rv := objc.Send[NCWidgetListViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NCWidgetListViewControllerClass) New() NCWidgetListViewController {
	rv := objc.Send[NCWidgetListViewController](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NCWidgetListViewController) Init() NCWidgetListViewController {
	rv := objc.Send[NCWidgetListViewController](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NCWidgetListViewController) Autorelease() NCWidgetListViewController {
	rv := objc.Send[NCWidgetListViewController](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNCWidgetListViewController creates a new NCWidgetListViewController instance.
func NewNCWidgetListViewController() NCWidgetListViewController {
	return getNCWidgetListViewControllerClass().New()
}


// Returns the row represented by the specified content view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/row(for:)
func (n_ NCWidgetListViewController) RowForViewController(viewController appkit.IViewController) uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("rowForViewController:"), viewController)
	return rv
}

// Returns the content view controller associated with the specified row, or a new content view controller if desired.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/viewController(atRow:makeIfNecessary:)
func (n_ NCWidgetListViewController) ViewControllerAtRowMakeIfNecessary(row uint, makeIfNecesary bool) appkit.ViewController {
	rv := objc.Send[appkit.ViewController](n_.ID, objc.Sel("viewControllerAtRow:makeIfNecessary:"), row, makeIfNecesary)
	return rv
}

// An array of objects to display in the list view.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/contents
func (n_ NCWidgetListViewController) Contents() []objc.ID {
	rv := objc.Send[[]objc.ID](n_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// An array of objects to display in the list view.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/contents
func (n_ NCWidgetListViewController) SetContents(value []objc.ID) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setContents:"), nsArray)
}

// The list view controller’s delegate or if the receiver doesn’t have a delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/delegate
func (n_ NCWidgetListViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The list view controller’s delegate or if the receiver doesn’t have a delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/delegate
func (n_ NCWidgetListViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the list is in editing mode.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/editing
func (n_ NCWidgetListViewController) Editing() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("editing"))
	return rv
}


// SetEditing sets the value of the editing property.
// A Boolean value that indicates whether the list is in editing mode.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/editing
func (n_ NCWidgetListViewController) SetEditing(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEditing:"), value)
}

// A Boolean value that indicates whether list displays divider lines between rows.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/hasDividerLines
func (n_ NCWidgetListViewController) HasDividerLines() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasDividerLines"))
	return rv
}


// SetHasDividerLines sets the value of the hasDividerLines property.
// A Boolean value that indicates whether list displays divider lines between rows.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/hasDividerLines
func (n_ NCWidgetListViewController) SetHasDividerLines(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasDividerLines:"), value)
}

// The minimum number of visible rows to display.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/minimumVisibleRowCount
func (n_ NCWidgetListViewController) MinimumVisibleRowCount() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumVisibleRowCount"))
	return rv
}


// SetMinimumVisibleRowCount sets the value of the minimumVisibleRowCount property.
// The minimum number of visible rows to display.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/minimumVisibleRowCount
func (n_ NCWidgetListViewController) SetMinimumVisibleRowCount(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumVisibleRowCount:"), value)
}

// A Boolean value that indicates whether an Add (+) button is displayed while the list is in editing mode.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/showsAddButtonWhenEditing
func (n_ NCWidgetListViewController) ShowsAddButtonWhenEditing() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("showsAddButtonWhenEditing"))
	return rv
}


// SetShowsAddButtonWhenEditing sets the value of the showsAddButtonWhenEditing property.
// A Boolean value that indicates whether an Add (+) button is displayed while the list is in editing mode.

//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetListViewController/showsAddButtonWhenEditing
func (n_ NCWidgetListViewController) SetShowsAddButtonWhenEditing(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShowsAddButtonWhenEditing:"), value)
}



