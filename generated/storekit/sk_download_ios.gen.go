//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Download


// iOS-only properties

// The current state of the download object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/downloadState
func (d_ Download) DownloadState() DownloadState {
	rv := objc.Send[DownloadState](d_.ID, objc.Sel("downloadState"))
	return rv
}





