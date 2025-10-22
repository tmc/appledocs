// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKFilterBrowserView] class.
var (
	IKFilterBrowserViewClass     _IKFilterBrowserViewClass
	IKFilterBrowserViewClassOnce sync.Once
)

func getIKFilterBrowserViewClass() _IKFilterBrowserViewClass {
	IKFilterBrowserViewClassOnce.Do(func() {
		IKFilterBrowserViewClass = _IKFilterBrowserViewClass{objc.GetClass("IKFilterBrowserView")}
	})
	return IKFilterBrowserViewClass
}

type _IKFilterBrowserViewClass struct {
	class objc.Class
}

// An interface definition for the [IKFilterBrowserView] class.
type IIKFilterBrowserView interface {
	appkit.IView
}

// The class is used as a container for the elements of an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserView
type IKFilterBrowserView struct {
	appkit.View
}

// IKFilterBrowserViewFrom constructs a [IKFilterBrowserView] from an unsafe.Pointer.
//
// The class is used as a container for the elements of an object.
func IKFilterBrowserViewFrom(ptr unsafe.Pointer) IKFilterBrowserView {
	return IKFilterBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKFilterBrowserViewClass) Alloc() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKFilterBrowserViewClass) New() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterBrowserView) Init() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterBrowserView) Autorelease() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterBrowserView creates a new IKFilterBrowserView instance.
func NewIKFilterBrowserView() IKFilterBrowserView {
	return getIKFilterBrowserViewClass().New()
}




