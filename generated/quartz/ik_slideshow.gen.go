// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AutoPlayDelay() float64
	SetAutoPlayDelay(value float64)
	// methods:
}

// The class encapsulates a data source and options for a slideshow.


// The class encapsulates a data source and options for a slideshow.
//
// [Full Topic]
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



// Controls the interval of time before a slideshow starts to play automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/autoplaydelay
func (i_ IKSlideshow) AutoPlayDelay() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("autoPlayDelay"))
	return rv
}


// Controls the interval of time before a slideshow starts to play automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshow/autoplaydelay
func (i_ IKSlideshow) SetAutoPlayDelay(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutoPlayDelay:"), value)
}



