// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKImageEditPanel] class.
var (
	IKImageEditPanelClass     _IKImageEditPanelClass
	IKImageEditPanelClassOnce sync.Once
)

func getIKImageEditPanelClass() _IKImageEditPanelClass {
	IKImageEditPanelClassOnce.Do(func() {
		IKImageEditPanelClass = _IKImageEditPanelClass{objc.GetClass("IKImageEditPanel")}
	})
	return IKImageEditPanelClass
}

type _IKImageEditPanelClass struct {
	class objc.Class
}

// An interface definition for the [IKImageEditPanel] class.
type IIKImageEditPanel interface {
	appkit.IPanel
}

// The class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel
type IKImageEditPanel struct {
	appkit.Panel
}

// IKImageEditPanelFrom constructs a [IKImageEditPanel] from an unsafe.Pointer.
//
// The class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing.
func IKImageEditPanelFrom(ptr unsafe.Pointer) IKImageEditPanel {
	return IKImageEditPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKImageEditPanelClass) Alloc() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKImageEditPanelClass) New() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageEditPanel) Init() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageEditPanel) Autorelease() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageEditPanel creates a new IKImageEditPanel instance.
func NewIKImageEditPanel() IKImageEditPanel {
	return getIKImageEditPanelClass().New()
}




