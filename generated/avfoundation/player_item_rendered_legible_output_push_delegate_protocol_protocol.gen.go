// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PPlayerItemRenderedLegibleOutputPushDelegate is the AVPlayerItemRenderedLegibleOutputPushDelegate protocol interface.
//
// A delegate that handles the rendered pixel buffers produced by a rendered legible output object.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutputPushDelegate
type PPlayerItemRenderedLegibleOutputPushDelegate interface {
	// Optional methods
	RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(output IAVPlayerItemRenderedLegibleOutput, captionImages []RenderedCaptionImage, itemTime objectivec.IObject)
	HasRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime() bool
}
