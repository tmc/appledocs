//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for PreviewController


// Asks the Quick Look preview controller to recompute the display of the current preview item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/refreshCurrentPreviewItem()
func (p_ PreviewController) RefreshCurrentPreviewItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("refreshCurrentPreviewItem"))
}

// Asks the preview controller to reload its data from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/reloadData()
func (p_ PreviewController) ReloadData() {
	objc.Send[objc.ID](p_.ID, objc.Sel("reloadData"))
}

// iOS-only properties

// The item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/currentPreviewItem
func (p_ PreviewController) CurrentPreviewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentPreviewItem"))
	return rv
}

// The index within the preview item navigation list of the item displaying in the Quick Look preview controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/currentPreviewItemIndex
func (p_ PreviewController) CurrentPreviewItemIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPreviewItemIndex"))
	return rv
}
func (p_ PreviewController) SetCurrentPreviewItemIndex(value int) {
	p_.ID.Send(objc.RegisterName("setCurrentPreviewItemIndex:"), value)
}

// The preview controller’s data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/dataSource
func (p_ PreviewController) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}
func (p_ PreviewController) SetDataSource(value unsafe.Pointer) {
	p_.ID.Send(objc.RegisterName("setDataSource:"), value)
}

// The preview controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/delegate
func (p_ PreviewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}
func (p_ PreviewController) SetDelegate(value unsafe.Pointer) {
	p_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





