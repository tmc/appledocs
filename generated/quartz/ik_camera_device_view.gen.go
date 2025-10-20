// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKCameraDeviceView] class.
var (
	IKCameraDeviceViewClass     _IKCameraDeviceViewClass
	IKCameraDeviceViewClassOnce sync.Once
)

func getIKCameraDeviceViewClass() _IKCameraDeviceViewClass {
	IKCameraDeviceViewClassOnce.Do(func() {
		IKCameraDeviceViewClass = _IKCameraDeviceViewClass{objc.GetClass("IKCameraDeviceView")}
	})
	return IKCameraDeviceViewClass
}

type _IKCameraDeviceViewClass struct {
	class objc.Class
}

// An interface definition for the [IKCameraDeviceView] class.
type IIKCameraDeviceView interface {
	appkit.IView
}

// The class displays the contents of the selected camera.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView
type IKCameraDeviceView struct {
	appkit.View
}

// IKCameraDeviceViewFrom constructs a [IKCameraDeviceView] from an unsafe.Pointer.
//
// The class displays the contents of the selected camera.
func IKCameraDeviceViewFrom(ptr unsafe.Pointer) IKCameraDeviceView {
	return IKCameraDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKCameraDeviceViewClass) Alloc() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKCameraDeviceViewClass) New() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKCameraDeviceView) Init() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKCameraDeviceView) Autorelease() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKCameraDeviceView creates a new IKCameraDeviceView instance.
func NewIKCameraDeviceView() IKCameraDeviceView {
	return getIKCameraDeviceViewClass().New()
}




