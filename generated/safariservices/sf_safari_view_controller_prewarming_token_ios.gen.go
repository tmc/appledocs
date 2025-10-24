//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SFSafariViewControllerPrewarmingToken


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/PrewarmingToken/invalidate()
func (s_ SFSafariViewControllerPrewarmingToken) Invalidate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidate"))
}

// iOS-only properties





