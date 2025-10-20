// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKDeviceBrowserView] class.
var (
	IKDeviceBrowserViewClass     _IKDeviceBrowserViewClass
	IKDeviceBrowserViewClassOnce sync.Once
)

func getIKDeviceBrowserViewClass() _IKDeviceBrowserViewClass {
	IKDeviceBrowserViewClassOnce.Do(func() {
		IKDeviceBrowserViewClass = _IKDeviceBrowserViewClass{objc.GetClass("IKDeviceBrowserView")}
	})
	return IKDeviceBrowserViewClass
}

type _IKDeviceBrowserViewClass struct {
	class objc.Class
}

// An interface definition for the [IKDeviceBrowserView] class.
type IIKDeviceBrowserView interface {
	appkit.IView
}

// The allows you to select a camera or scanner from a list of the available devices.
//
// The delegate must conform to the protocol. The delegate provides methods to inform you of selection changes in the browser as well as errors encountered when creating the browser list.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView
type IKDeviceBrowserView struct {
	appkit.View
}

// IKDeviceBrowserViewFrom constructs a [IKDeviceBrowserView] from an unsafe.Pointer.
//
// The allows you to select a camera or scanner from a list of the available devices.
func IKDeviceBrowserViewFrom(ptr unsafe.Pointer) IKDeviceBrowserView {
	return IKDeviceBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKDeviceBrowserViewClass) Alloc() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKDeviceBrowserViewClass) New() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKDeviceBrowserView) Init() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKDeviceBrowserView) Autorelease() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKDeviceBrowserView creates a new IKDeviceBrowserView instance.
func NewIKDeviceBrowserView() IKDeviceBrowserView {
	return getIKDeviceBrowserViewClass().New()
}




