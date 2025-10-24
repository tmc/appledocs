// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"
)

// PSFAddToHomeScreenActivityItem is the SFAddToHomeScreenActivityItem protocol interface.
//
// A protocol that describes a bookmark someone can add to their Home Screen.
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - visionOS 1.1+
//
// See: doc://com.apple.safariservices/documentation/SafariServices/SFAddToHomeScreenActivityItem
type PSFAddToHomeScreenActivityItem interface {
	// Optional methods
	GetHomeScreenWebAppInfoWithCompletionHandler(completionHandler unsafe.Pointer)
	HasGetHomeScreenWebAppInfoWithCompletionHandler() bool
	GetWebAppManifestWithCompletionHandler(completionHandler unsafe.Pointer)
	HasGetWebAppManifestWithCompletionHandler() bool
}
