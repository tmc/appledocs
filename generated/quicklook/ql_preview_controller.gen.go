// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PreviewController] class.
var (
	PreviewControllerClass     _PreviewControllerClass
	PreviewControllerClassOnce sync.Once
)

func getPreviewControllerClass() _PreviewControllerClass {
	PreviewControllerClassOnce.Do(func() {
		PreviewControllerClass = _PreviewControllerClass{objc.GetClass("QLPreviewController")}
	})
	return PreviewControllerClass
}

type _PreviewControllerClass struct {
	class objc.Class
}

// An interface definition for the [PreviewController] class.
type IPreviewController interface {
	appkit.IViewController
	RefreshCurrentPreviewItem()
}

// A specialized view controller for previewing an item.
//
// A can display previews for many common file types, including the following: iWork documents Microsoft Office documents Rich text format, or RTF, documents PDF files Images Text files with a uniform type identifier that conforms to the type. To learn more, see . Comma-separated values, or CSV, files 3D models in the USDZ format with both standalone and AR views for viewing the model
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController
type PreviewController struct {
	appkit.ViewController
}

// PreviewControllerFrom constructs a [PreviewController] from an unsafe.Pointer.
//
// A specialized view controller for previewing an item.
func PreviewControllerFrom(ptr unsafe.Pointer) PreviewController {
	return PreviewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewControllerClass) Alloc() PreviewController {
	rv := objc.Send[PreviewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewControllerClass) New() PreviewController {
	rv := objc.Send[PreviewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewController) Init() PreviewController {
	rv := objc.Send[PreviewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewController) Autorelease() PreviewController {
	rv := objc.Send[PreviewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewController creates a new PreviewController instance.
func NewPreviewController() PreviewController {
	return getPreviewControllerClass().New()
}


// Asks the Quick Look preview controller to recompute the display of the current preview item.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/refreshCurrentPreviewItem()
func (p_ PreviewController) RefreshCurrentPreviewItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("refreshCurrentPreviewItem"))
}

// The item displaying in the Quick Look preview controller.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/currentPreviewItem
func (p_ PreviewController) CurrentPreviewItem() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("currentPreviewItem"))
	return rv
}



