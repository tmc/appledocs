// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StatusBarButton] class.
var (
	StatusBarButtonClass     _StatusBarButtonClass
	StatusBarButtonClassOnce sync.Once
)

func getStatusBarButtonClass() _StatusBarButtonClass {
	StatusBarButtonClassOnce.Do(func() {
		StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}
	})
	return StatusBarButtonClass
}

type _StatusBarButtonClass struct {
	class objc.Class
}

// An interface definition for the [StatusBarButton] class.
type IStatusBarButton interface {
	IButton
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

// The appearance and behavior of an item in the systemwide menu bar.


// The appearance and behavior of an item in the systemwide menu bar.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton/appearsDisabled
func (s_ StatusBarButton) AppearsDisabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("appearsDisabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton/appearsDisabled
func (s_ StatusBarButton) SetAppearsDisabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppearsDisabled:"), value)
}



