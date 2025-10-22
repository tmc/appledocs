// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ICCameraFolder] class.
var (
	ICCameraFolderClass     _ICCameraFolderClass
	ICCameraFolderClassOnce sync.Once
)

func getICCameraFolderClass() _ICCameraFolderClass {
	ICCameraFolderClassOnce.Do(func() {
		ICCameraFolderClass = _ICCameraFolderClass{objc.GetClass("ICCameraFolder")}
	})
	return ICCameraFolderClass
}

type _ICCameraFolderClass struct {
	class objc.Class
}

// An interface definition for the [ICCameraFolder] class.
type IICCameraFolder interface {
	IICCameraItem
	Contents() []ICCameraItem
}

// An object that represents a folder on a camera.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder
type ICCameraFolder struct {
	ICCameraItem
}

// ICCameraFolderFrom constructs a [ICCameraFolder] from an unsafe.Pointer.
//
// An object that represents a folder on a camera.
func ICCameraFolderFrom(ptr unsafe.Pointer) ICCameraFolder {
	return ICCameraFolder{
		ICCameraItem: ICCameraItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ICCameraFolderClass) Alloc() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICCameraFolderClass) New() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraFolder) Init() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraFolder) Autorelease() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFolder creates a new ICCameraFolder instance.
func NewICCameraFolder() ICCameraFolder {
	return getICCameraFolderClass().New()
}


// A list of items that this folder contains.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder/contents
func (i_ ICCameraFolder) Contents() []ICCameraItem {
	rv := objc.Send[[]ICCameraItem](i_.ID, objc.Sel("contents"))
	return rv
}



