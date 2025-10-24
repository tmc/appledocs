//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SFSafariViewControllerDataStore


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DataStore/clearWebsiteData(completionHandler:)
func (s_ SFSafariViewControllerDataStore) ClearWebsiteDataWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("clearWebsiteDataWithCompletionHandler:"), completion)
}

// iOS-only properties





