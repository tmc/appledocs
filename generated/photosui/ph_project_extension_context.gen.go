// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/photos"
)

// The class instance for the [PHProjectExtensionContext] class.
var (
	PHProjectExtensionContextClass     _PHProjectExtensionContextClass
	PHProjectExtensionContextClassOnce sync.Once
)

func getPHProjectExtensionContextClass() _PHProjectExtensionContextClass {
	PHProjectExtensionContextClassOnce.Do(func() {
		PHProjectExtensionContextClass = _PHProjectExtensionContextClass{objc.GetClass("PHProjectExtensionContext")}
	})
	return PHProjectExtensionContextClass
}

type _PHProjectExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectExtensionContext] class.
type IPHProjectExtensionContext interface {
	IExtensionContext
	ShowEditorForAsset(asset photos.IPHAsset)
	UpdatedProjectInfoFromProjectInfoCompletion(existingProjectInfo IPHProjectInfo, completion unsafe.Pointer) foundation.Progress
}

// An object that provides Photos project extensions with access to the underlying project, as well as to the user’s photo library for editing.
//
// When a Photos project extension is initialized, it is handed an extension context object. This object provides the extension with access to the underlying project, as well as the photo library from which assets are fetched and edited.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectExtensionContext
type PHProjectExtensionContext struct {
	ExtensionContext
}

// PHProjectExtensionContextFrom constructs a [PHProjectExtensionContext] from an unsafe.Pointer.
//
// An object that provides Photos project extensions with access to the underlying project, as well as to the user’s photo library for editing.
func PHProjectExtensionContextFrom(ptr unsafe.Pointer) PHProjectExtensionContext {
	return PHProjectExtensionContext{
		ExtensionContext: ExtensionContextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectExtensionContextClass) Alloc() PHProjectExtensionContext {
	rv := objc.Send[PHProjectExtensionContext](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectExtensionContextClass) New() PHProjectExtensionContext {
	rv := objc.Send[PHProjectExtensionContext](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectExtensionContext) Init() PHProjectExtensionContext {
	rv := objc.Send[PHProjectExtensionContext](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectExtensionContext) Autorelease() PHProjectExtensionContext {
	rv := objc.Send[PHProjectExtensionContext](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectExtensionContext creates a new PHProjectExtensionContext instance.
func NewPHProjectExtensionContext() PHProjectExtensionContext {
	return getPHProjectExtensionContextClass().New()
}


// Invokes the built-in photo editor for the given asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectExtensionContext/showEditor(for:)
func (p_ PHProjectExtensionContext) ShowEditorForAsset(asset photos.IPHAsset) {
	objc.Send[objc.ID](p_.ID, objc.Sel("showEditorForAsset:"), asset)
}

// Creates an updated instance from existing project information and current assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectExtensionContext/updatedProjectInfo(from:completion:)
func (p_ PHProjectExtensionContext) UpdatedProjectInfoFromProjectInfoCompletion(existingProjectInfo IPHProjectInfo, completion unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](p_.ID, objc.Sel("updatedProjectInfoFromProjectInfo:completion:"), existingProjectInfo, completion)
	return rv
}

// A read-only version of the photo library being modified.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectExtensionContext/photoLibrary
func (p_ PHProjectExtensionContext) PhotoLibrary() photos.PHPhotoLibrary {
	rv := objc.Send[photos.PHPhotoLibrary](p_.ID, objc.Sel("photoLibrary"))
	return rv
}

// A read-only version of the project being edited.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectExtensionContext/project
func (p_ PHProjectExtensionContext) Project() photos.PHProject {
	rv := objc.Send[photos.PHProject](p_.ID, objc.Sel("project"))
	return rv
}



