//go:build darwin && ios

// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DocumentCameraScan


// Requests the image of a page at a specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan/imageOfPage(at:)
func (d_ DocumentCameraScan) ImageOfPageAtIndex(index uint) appkit.Image {
	rv := objc.Send[appkit.Image](d_.ID, objc.Sel("imageOfPageAtIndex:"), index)
	return rv
}

// iOS-only properties

// The number of pages in the scanned document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan/pageCount
func (d_ DocumentCameraScan) PageCount() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("pageCount"))
	return rv
}

// The title of the scanned document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan/title
func (d_ DocumentCameraScan) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("title"))
	return rv
}





