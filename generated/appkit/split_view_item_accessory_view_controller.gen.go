// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SplitViewItemAccessoryViewController] class.
var splitViewItemAccessoryViewControllerClass = _SplitViewItemAccessoryViewControllerClass{objc.GetClass("NSSplitViewItemAccessoryViewController")}

type _SplitViewItemAccessoryViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [SplitViewItemAccessoryViewController] class.
type ISplitViewItemAccessoryViewController interface {
	IViewController
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController

type SplitViewItemAccessoryViewController struct {
	ViewController
}

// SplitViewItemAccessoryViewControllerFrom constructs a [SplitViewItemAccessoryViewController] from an unsafe.Pointer.
func SplitViewItemAccessoryViewControllerFrom(ptr unsafe.Pointer) SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemAccessoryViewControllerClass) Alloc() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SplitViewItemAccessoryViewControllerClass) New() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewItemAccessoryViewController) Init() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewItemAccessoryViewController) Autorelease() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewItemAccessoryViewController creates a new SplitViewItemAccessoryViewController instance.
func NewSplitViewItemAccessoryViewController() SplitViewItemAccessoryViewController {
	return splitViewItemAccessoryViewControllerClass.New()
}




