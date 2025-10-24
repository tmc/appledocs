//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ContextMenuElementInfo


// iOS-only properties

// The URL of the link that the user clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContextMenuElementInfo/linkURL
func (c_ ContextMenuElementInfo) LinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("linkURL"))
	return rv
}





