// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ICCameraItem] class.
var (
	ICCameraItemClass     _ICCameraItemClass
	ICCameraItemClassOnce sync.Once
)

func getICCameraItemClass() _ICCameraItemClass {
	ICCameraItemClassOnce.Do(func() {
		ICCameraItemClass = _ICCameraItemClass{objc.GetClass("ICCameraItem")}
	})
	return ICCameraItemClass
}

type _ICCameraItemClass struct {
	class objc.Class
}

// An interface definition for the [ICCameraItem] class.
type IICCameraItem interface {
	objectivec.IObject
}

// An abstract class that represents a camera item.
//
// The ImageCaptureCore framework defines two concrete subclasses of camera items: and .
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem
type ICCameraItem struct {
	objectivec.Object
}

// ICCameraItemFrom constructs a [ICCameraItem] from an unsafe.Pointer.
//
// An abstract class that represents a camera item.
func ICCameraItemFrom(ptr unsafe.Pointer) ICCameraItem {
	return ICCameraItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICCameraItemClass) Alloc() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICCameraItemClass) New() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraItem) Init() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraItem) Autorelease() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraItem creates a new ICCameraItem instance.
func NewICCameraItem() ICCameraItem {
	return getICCameraItemClass().New()
}


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isRaw
func (i_ ICCameraItem) Raw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("raw"))
	return rv
}

// The item’s modification date, usually the same as its modification date.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/modificationDate
func (i_ ICCameraItem) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("modificationDate"))
	return rv
}

// The item’s thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/thumbnail
func (i_ ICCameraItem) Thumbnail() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](i_.ID, objc.Sel("thumbnail"))
	return rv
}

// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/uti
func (i_ ICCameraItem) UTI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("UTI"))
	return rv
}



