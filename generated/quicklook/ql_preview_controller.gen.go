// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CurrentPreviewItem() PreviewItem /* not a class type */
	SetCurrentPreviewItem(value PreviewItem /* not a class type */)
	CurrentPreviewItemIndex() int
	SetCurrentPreviewItemIndex(value int)
	DataSource() PreviewControllerDataSource /* not a class type */
	SetDataSource(value PreviewControllerDataSource /* not a class type */)
	Delegate() PreviewControllerDelegate /* not a class type */
	SetDelegate(value PreviewControllerDelegate /* not a class type */)
	PreviewItemTitle() objc.IObject /* cross-framework: NSString */
	SetPreviewItemTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A specialized view controller for previewing an item.
//
// A can display previews for many common file types, including the following: iWork documents Microsoft Office documents Rich text format, or RTF, documents PDF files Images Text files with a uniform type identifier that conforms to the type. To learn more, see . Comma-separated values, or CSV, files 3D models in the USDZ format with both standalone and AR views for viewing the model


// A specialized view controller for previewing an item.
//
// [Full Topic]
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



// The item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitem
func (p_ PreviewController) CurrentPreviewItem() PreviewItem /* not a class type */ {
	rv := objc.Send[PreviewItem](p_.ID, objc.Sel("currentPreviewItem"))
	return rv
}


// The item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitem
func (p_ PreviewController) SetCurrentPreviewItem(value PreviewItem /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItem:"), value)
}


// The index within the preview item navigation list of the item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitemindex
func (p_ PreviewController) CurrentPreviewItemIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPreviewItemIndex"))
	return rv
}


// The index within the preview item navigation list of the item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/currentpreviewitemindex
func (p_ PreviewController) SetCurrentPreviewItemIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItemIndex:"), value)
}


// The preview controller’s data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/datasource
func (p_ PreviewController) DataSource() PreviewControllerDataSource /* not a class type */ {
	rv := objc.Send[PreviewControllerDataSource](p_.ID, objc.Sel("dataSource"))
	return rv
}


// The preview controller’s data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/datasource
func (p_ PreviewController) SetDataSource(value PreviewControllerDataSource /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}


// The preview controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/delegate
func (p_ PreviewController) Delegate() PreviewControllerDelegate /* not a class type */ {
	rv := objc.Send[PreviewControllerDelegate](p_.ID, objc.Sel("delegate"))
	return rv
}


// The preview controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewcontroller/delegate
func (p_ PreviewController) SetDelegate(value PreviewControllerDelegate /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) PreviewItemTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("previewItemTitle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) SetPreviewItemTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreviewItemTitle:"), value)
}


