// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CapturePhotoSettings] class.
var (
	CapturePhotoSettingsClass     _CapturePhotoSettingsClass
	CapturePhotoSettingsClassOnce sync.Once
)

func getCapturePhotoSettingsClass() _CapturePhotoSettingsClass {
	CapturePhotoSettingsClassOnce.Do(func() {
		CapturePhotoSettingsClass = _CapturePhotoSettingsClass{objc.GetClass("AVCapturePhotoSettings")}
	})
	return CapturePhotoSettingsClass
}

type _CapturePhotoSettingsClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type CapturePhotoSettings struct {
	objectivec.Object
}

// CapturePhotoSettingsFrom constructs a [CapturePhotoSettings] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func CapturePhotoSettingsFrom(ptr unsafe.Pointer) CapturePhotoSettings {
	return CapturePhotoSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoSettingsClass) Alloc() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoSettingsClass) New() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoSettings) Init() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoSettings) Autorelease() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoSettings creates a new CapturePhotoSettings instance.
func NewCapturePhotoSettings() CapturePhotoSettings {
	return getCapturePhotoSettingsClass().New()
}




