// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKPictureTaker] class.
var (
	IKPictureTakerClass     _IKPictureTakerClass
	IKPictureTakerClassOnce sync.Once
)

func getIKPictureTakerClass() _IKPictureTakerClass {
	IKPictureTakerClassOnce.Do(func() {
		IKPictureTakerClass = _IKPictureTakerClass{objc.GetClass("IKPictureTaker")}
	})
	return IKPictureTakerClass
}

type _IKPictureTakerClass struct {
	class objc.Class
}

// An interface definition for the [IKPictureTaker] class.
type IIKPictureTaker interface {
	appkit.IPanel
}

// The class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker
type IKPictureTaker struct {
	appkit.Panel
}

// IKPictureTakerFrom constructs a [IKPictureTaker] from an unsafe.Pointer.
//
// The class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera.
func IKPictureTakerFrom(ptr unsafe.Pointer) IKPictureTaker {
	return IKPictureTaker{
		Panel: appkit.PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKPictureTakerClass) Alloc() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKPictureTakerClass) New() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKPictureTaker) Init() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKPictureTaker) Autorelease() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKPictureTaker creates a new IKPictureTaker instance.
func NewIKPictureTaker() IKPictureTaker {
	return getIKPictureTakerClass().New()
}




