// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ICCameraFile] class.
var (
	ICCameraFileClass     _ICCameraFileClass
	ICCameraFileClassOnce sync.Once
)

func getICCameraFileClass() _ICCameraFileClass {
	ICCameraFileClassOnce.Do(func() {
		ICCameraFileClass = _ICCameraFileClass{objc.GetClass("ICCameraFile")}
	})
	return ICCameraFileClass
}

type _ICCameraFileClass struct {
	class objc.Class
}

// An interface definition for the [ICCameraFile] class.
type IICCameraFile interface {
	IICCameraItem
	RequestSecurityScopedURLWithCompletion(completion unsafe.Pointer)
}

// An object that represents a file on a camera.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile
type ICCameraFile struct {
	ICCameraItem
}

// ICCameraFileFrom constructs a [ICCameraFile] from an unsafe.Pointer.
//
// An object that represents a file on a camera.
func ICCameraFileFrom(ptr unsafe.Pointer) ICCameraFile {
	return ICCameraFile{
		ICCameraItem: ICCameraItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ICCameraFileClass) Alloc() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICCameraFileClass) New() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraFile) Init() ICCameraFile {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraFile) Autorelease() ICCameraFile {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFile creates a new ICCameraFile instance.
func NewICCameraFile() ICCameraFile {
	return getICCameraFileClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestSecurityScopedURL(completion:)
func (i_ ICCameraFile) RequestSecurityScopedURLWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSecurityScopedURLWithCompletion:"), completion)
}

// The burst UUID of the file if it is in a burst.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstUUID
func (i_ ICCameraFile) BurstUUID() string {
	rv := objc.Send[string](i_.ID, objc.Sel("burstUUID"))
	return rv
}

// The duration, in seconds, of an audio or video file.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/duration
func (i_ ICCameraFile) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("duration"))
	return rv
}

// The width of an image or movie frame.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/width
func (i_ ICCameraFile) Width() int {
	rv := objc.Send[int](i_.ID, objc.Sel("width"))
	return rv
}



