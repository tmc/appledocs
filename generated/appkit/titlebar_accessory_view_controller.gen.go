// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TitlebarAccessoryViewController] class.
var (
	titlebarAccessoryViewControllerClass     _TitlebarAccessoryViewControllerClass
	titlebarAccessoryViewControllerClassOnce sync.Once
)

func getTitlebarAccessoryViewControllerClass() _TitlebarAccessoryViewControllerClass {
	titlebarAccessoryViewControllerClassOnce.Do(func() {
		titlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}
	})
	return titlebarAccessoryViewControllerClass
}

type _TitlebarAccessoryViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TitlebarAccessoryViewController] class.
type ITitlebarAccessoryViewController interface {
	IViewController
}

// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController

type TitlebarAccessoryViewController struct {
	ViewController
}

// TitlebarAccessoryViewControllerFrom constructs a [TitlebarAccessoryViewController] from an unsafe.Pointer.
//
// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window.
func TitlebarAccessoryViewControllerFrom(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TitlebarAccessoryViewController) Autorelease() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTitlebarAccessoryViewController creates a new TitlebarAccessoryViewController instance.
func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return getTitlebarAccessoryViewControllerClass().New()
}




