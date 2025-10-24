// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PContentSharingPickerObserver is the SCContentSharingPickerObserver protocol interface.
//
// An observer protocol your app implements to receive messages from the operating system’s content picker.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 14.0+
//
// See: doc://com.apple.screencapturekit/documentation/ScreenCaptureKit/SCContentSharingPickerObserver
type PContentSharingPickerObserver interface {
	// Required methods
	ContentSharingPickerDidCancelForStream(picker ISCContentSharingPicker, stream ISCStream)/* debug [protocol_interface/required_method]: ContentSharingPickerDidCancelForStream */
	ContentSharingPickerDidUpdateWithFilterForStream(picker ISCContentSharingPicker, filter ISCContentFilter, stream ISCStream)/* debug [protocol_interface/required_method]: ContentSharingPickerDidUpdateWithFilterForStream */
	ContentSharingPickerStartDidFailWithError(error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: ContentSharingPickerStartDidFailWithError */
}
