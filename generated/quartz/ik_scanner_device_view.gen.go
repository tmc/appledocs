// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKScannerDeviceView] class.
var (
	IKScannerDeviceViewClass     _IKScannerDeviceViewClass
	IKScannerDeviceViewClassOnce sync.Once
)

func getIKScannerDeviceViewClass() _IKScannerDeviceViewClass {
	IKScannerDeviceViewClassOnce.Do(func() {
		IKScannerDeviceViewClass = _IKScannerDeviceViewClass{objc.GetClass("IKScannerDeviceView")}
	})
	return IKScannerDeviceViewClass
}

type _IKScannerDeviceViewClass struct {
	class objc.Class
}

// An interface definition for the [IKScannerDeviceView] class.
type IIKScannerDeviceView interface {
	appkit.IView
}

// The class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView
type IKScannerDeviceView struct {
	appkit.View
}

// IKScannerDeviceViewFrom constructs a [IKScannerDeviceView] from an unsafe.Pointer.
//
// The class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the protocol.
func IKScannerDeviceViewFrom(ptr unsafe.Pointer) IKScannerDeviceView {
	return IKScannerDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKScannerDeviceViewClass) Alloc() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKScannerDeviceViewClass) New() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKScannerDeviceView) Init() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKScannerDeviceView) Autorelease() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKScannerDeviceView creates a new IKScannerDeviceView instance.
func NewIKScannerDeviceView() IKScannerDeviceView {
	return getIKScannerDeviceViewClass().New()
}




