// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BackgroundExtensionView] class.
var (
	backgroundExtensionViewClass     _BackgroundExtensionViewClass
	backgroundExtensionViewClassOnce sync.Once
)

func getBackgroundExtensionViewClass() _BackgroundExtensionViewClass {
	backgroundExtensionViewClassOnce.Do(func() {
		backgroundExtensionViewClass = _BackgroundExtensionViewClass{objc.GetClass("NSBackgroundExtensionView")}
	})
	return backgroundExtensionViewClass
}

type _BackgroundExtensionViewClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundExtensionView] class.
type IBackgroundExtensionView interface {
	IView
}

// A view that extends content to fill its own bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView
type BackgroundExtensionView struct {
	View
}

// BackgroundExtensionViewFrom constructs a [BackgroundExtensionView] from an unsafe.Pointer.
//
// A view that extends content to fill its own bounds.
func BackgroundExtensionViewFrom(ptr unsafe.Pointer) BackgroundExtensionView {
	return BackgroundExtensionView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundExtensionViewClass) Alloc() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackgroundExtensionViewClass) New() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackgroundExtensionView) Init() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackgroundExtensionView) Autorelease() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackgroundExtensionView creates a new BackgroundExtensionView instance.
func NewBackgroundExtensionView() BackgroundExtensionView {
	return getBackgroundExtensionViewClass().New()
}




