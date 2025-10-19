// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StatusBarButton] class.
var (
	statusBarButtonClass     _StatusBarButtonClass
	statusBarButtonClassOnce sync.Once
)

func getStatusBarButtonClass() _StatusBarButtonClass {
	statusBarButtonClassOnce.Do(func() {
		statusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}
	})
	return statusBarButtonClass
}

type _StatusBarButtonClass struct {
	class objc.Class
}

// An interface definition for the [StatusBarButton] class.
type IStatusBarButton interface {
	IButton
}

// The appearance and behavior of an item in the systemwide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton

type StatusBarButton struct {
	Button
}

// StatusBarButtonFrom constructs a [StatusBarButton] from an unsafe.Pointer.
//
// The appearance and behavior of an item in the systemwide menu bar.
func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: ButtonFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusBarButton) Autorelease() StatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusBarButton creates a new StatusBarButton instance.
func NewStatusBarButton() StatusBarButton {
	return getStatusBarButtonClass().New()
}




