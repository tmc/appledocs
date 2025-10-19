// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StatusBar] class.
var (
	statusBarClass     _StatusBarClass
	statusBarClassOnce sync.Once
)

func getStatusBarClass() _StatusBarClass {
	statusBarClassOnce.Do(func() {
		statusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}
	})
	return statusBarClass
}

type _StatusBarClass struct {
	class objc.Class
}

// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	objectivec.IObject
}

// An object that manages a collection of status items displayed within the system-wide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar
type StatusBar struct {
	objectivec.Object
}

// StatusBarFrom constructs a [StatusBar] from an unsafe.Pointer.
//
// An object that manages a collection of status items displayed within the system-wide menu bar.
func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StatusBarClass) Alloc() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StatusBarClass) New() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusBar) Init() StatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusBar) Autorelease() StatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusBar creates a new StatusBar instance.
func NewStatusBar() StatusBar {
	return getStatusBarClass().New()
}




