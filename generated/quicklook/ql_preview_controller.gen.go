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

// The index within the preview item navigation list of the item displaying in the Quick Look preview controller.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitemindex
func (p_ PreviewController) CurrentPreviewItemIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPreviewItemIndex"))
	return rv
}


// SetCurrentPreviewItemIndex sets the value of the currentPreviewItemIndex property.
// The index within the preview item navigation list of the item displaying in the Quick Look preview controller.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitemindex
func (p_ PreviewController) SetCurrentPreviewItemIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItemIndex:"), value)
}

// The preview controller’s data source.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/datasource
func (p_ PreviewController) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The preview controller’s data source.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/datasource
func (p_ PreviewController) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}

// The preview controller’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/delegate
func (p_ PreviewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The preview controller’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/delegate
func (p_ PreviewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) PreviewItemTitle() string {
	rv := objc.Send[string](p_.ID, objc.Sel("previewItemTitle"))
	return rv
}


// SetPreviewItemTitle sets the value of the previewItemTitle property.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) SetPreviewItemTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreviewItemTitle:"), objc.String(value))
}



