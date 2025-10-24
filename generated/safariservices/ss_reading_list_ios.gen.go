//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SSReadingList


// Adds an item to the Reading List.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList/addItem(with:title:previewText:)
func (s_ SSReadingList) AddReadingListItemWithURLTitlePreviewTextError(URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */, previewText objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addReadingListItemWithURL:title:previewText:error:"), URL, title, previewText, error_)
	return rv
}

// iOS-only properties






