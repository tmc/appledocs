// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PNCWidgetProviding is the NCWidgetProviding protocol interface.
//
// The interface for customizing the appearance and behavior of a Today widget.
//
// Availability:
//   - Mac Catalyst 10.0+ (Deprecated in 14.0)
//   - iOS 10.0+ (Deprecated in 14.0)
//   - iPadOS 10.0+ (Deprecated in 14.0)
//   - macOS 10.10+ (Deprecated in 11.0)
//
// See: doc://com.apple.notificationcenter/documentation/NotificationCenter/NCWidgetProviding
type PNCWidgetProviding interface {
	// Optional methods
	WidgetActiveDisplayModeDidChangeWithMaximumSize(activeDisplayMode NCWidgetDisplayMode, maxSize corefoundation.CGSize)
	HasWidgetActiveDisplayModeDidChangeWithMaximumSize() bool
	WidgetDidBeginEditing()
	HasWidgetDidBeginEditing() bool
	WidgetDidEndEditing()
	HasWidgetDidEndEditing() bool
	WidgetMarginInsetsForProposedMarginInsets(defaultMarginInsets foundation.EdgeInsets) foundation.EdgeInsets
	HasWidgetMarginInsetsForProposedMarginInsets() bool
	WidgetPerformUpdateWithCompletionHandler(completionHandler unsafe.Pointer)
	HasWidgetPerformUpdateWithCompletionHandler() bool
}
