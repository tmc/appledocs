// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TabViewController] class.
var (
	tabViewControllerClass     _TabViewControllerClass
	tabViewControllerClassOnce sync.Once
)

func getTabViewControllerClass() _TabViewControllerClass {
	tabViewControllerClassOnce.Do(func() {
		tabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}
	})
	return tabViewControllerClass
}

type _TabViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	IViewController
}

// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController

type TabViewController struct {
	ViewController
}

// TabViewControllerFrom constructs a [TabViewController] from an unsafe.Pointer.
//
// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time.
func TabViewControllerFrom(ptr unsafe.Pointer) TabViewController {
	return TabViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TabViewControllerClass) Alloc() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TabViewControllerClass) New() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabViewController) Init() TabViewController {
	rv := objc.Send[TabViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabViewController) Autorelease() TabViewController {
	rv := objc.Send[TabViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabViewController creates a new TabViewController instance.
func NewTabViewController() TabViewController {
	return getTabViewControllerClass().New()
}




