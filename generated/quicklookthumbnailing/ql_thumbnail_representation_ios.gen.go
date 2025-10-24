//go:build darwin && ios

// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ThumbnailRepresentation


// iOS-only properties

// A thumbnail in the form of a UIKit image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/uiImage
func (t_ ThumbnailRepresentation) UIImage() appkit.Image {
	rv := objc.Send[appkit.Image](t_.ID, objc.Sel("UIImage"))
	return rv
}








