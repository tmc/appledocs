// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IKSlideshow] class.
var (
	IKSlideshowClass     _IKSlideshowClass
	IKSlideshowClassOnce sync.Once
)

func getIKSlideshowClass() _IKSlideshowClass {
	IKSlideshowClassOnce.Do(func() {
		IKSlideshowClass = _IKSlideshowClass{objc.GetClass("IKSlideshow")}
	})
	return IKSlideshowClass
}

type _IKSlideshowClass struct {
	class objc.Class
}

// An interface definition for the [IKSlideshow] class.
type IIKSlideshow interface {
	objectivec.IObject
}

// The class encapsulates a data source and options for a slideshow.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow
type IKSlideshow struct {
	objectivec.Object
}

// IKSlideshowFrom constructs a [IKSlideshow] from an unsafe.Pointer.
//
// The class encapsulates a data source and options for a slideshow.
func IKSlideshowFrom(ptr unsafe.Pointer) IKSlideshow {
	return IKSlideshow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IKSlideshowClass) Alloc() IKSlideshow {
	rv := objc.Send[IKSlideshow](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKSlideshowClass) New() IKSlideshow {
	rv := objc.Send[IKSlideshow](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKSlideshow) Init() IKSlideshow {
	rv := objc.Send[IKSlideshow](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKSlideshow) Autorelease() IKSlideshow {
	rv := objc.Send[IKSlideshow](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKSlideshow creates a new IKSlideshow instance.
func NewIKSlideshow() IKSlideshow {
	return getIKSlideshowClass().New()
}




