// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CapturePhoto] class.
var (
	CapturePhotoClass     _CapturePhotoClass
	CapturePhotoClassOnce sync.Once
)

func getCapturePhotoClass() _CapturePhotoClass {
	CapturePhotoClassOnce.Do(func() {
		CapturePhotoClass = _CapturePhotoClass{objc.GetClass("AVCapturePhoto")}
	})
	return CapturePhotoClass
}

type _CapturePhotoClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhoto] class.
type ICapturePhoto interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type CapturePhoto struct {
	objectivec.Object
}

// CapturePhotoFrom constructs a [CapturePhoto] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func CapturePhotoFrom(ptr unsafe.Pointer) CapturePhoto {
	return CapturePhoto{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoClass) Alloc() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoClass) New() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhoto) Init() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhoto) Autorelease() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhoto creates a new CapturePhoto instance.
func NewCapturePhoto() CapturePhoto {
	return getCapturePhotoClass().New()
}




