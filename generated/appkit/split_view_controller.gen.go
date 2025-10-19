// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SplitViewController] class.
var (
	splitViewControllerClass     _SplitViewControllerClass
	splitViewControllerClassOnce sync.Once
)

func getSplitViewControllerClass() _SplitViewControllerClass {
	splitViewControllerClassOnce.Do(func() {
		splitViewControllerClass = _SplitViewControllerClass{objc.GetClass("NSSplitViewController")}
	})
	return splitViewControllerClass
}

type _SplitViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [SplitViewController] class.
type ISplitViewController interface {
	IViewController
}

// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController

type SplitViewController struct {
	ViewController
}

// SplitViewControllerFrom constructs a [SplitViewController] from an unsafe.Pointer.
//
// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views.
func SplitViewControllerFrom(ptr unsafe.Pointer) SplitViewController {
	return SplitViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SplitViewControllerClass) Alloc() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SplitViewControllerClass) New() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewController) Init() SplitViewController {
	rv := objc.Send[SplitViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewController) Autorelease() SplitViewController {
	rv := objc.Send[SplitViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewController creates a new SplitViewController instance.
func NewSplitViewController() SplitViewController {
	return getSplitViewControllerClass().New()
}




