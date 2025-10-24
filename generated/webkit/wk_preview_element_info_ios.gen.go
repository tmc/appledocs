//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PreviewElementInfo


// iOS-only properties

// The link for the webpage to be previewed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreviewElementInfo/linkURL
func (p_ PreviewElementInfo) LinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("linkURL"))
	return rv
}





