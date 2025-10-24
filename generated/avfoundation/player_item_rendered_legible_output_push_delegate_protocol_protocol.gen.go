// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"
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
	RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(output IAVPlayerItemRenderedLegibleOutput, captionImages []RenderedCaptionImage, itemTime objc.IObject /* cross-framework: Time */)
	HasRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime() bool
}

// PlayerItemRenderedLegibleOutputPushDelegate is a delegate implementation builder for the PPlayerItemRenderedLegibleOutputPushDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerItemRenderedLegibleOutputPushDelegate struct {
	_RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime func(output IAVPlayerItemRenderedLegibleOutput, captionImages []RenderedCaptionImage, itemTime objc.IObject /* cross-framework: Time */)
}

// SetRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime sets the handler for the RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime delegate method.
//
// Tells the delegate that new rendered caption images are available.
func (d *PlayerItemRenderedLegibleOutputPushDelegate) SetRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(f func(output IAVPlayerItemRenderedLegibleOutput, captionImages []RenderedCaptionImage, itemTime objc.IObject /* cross-framework: Time */)) {
	d._RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime = f
}

// RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime implements the PPlayerItemRenderedLegibleOutputPushDelegate interface.
func (d *PlayerItemRenderedLegibleOutputPushDelegate) RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(output IAVPlayerItemRenderedLegibleOutput, captionImages []RenderedCaptionImage, itemTime objc.IObject /* cross-framework: Time */) {
	if d._RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime != nil {
		d._RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime(output, captionImages, itemTime)
	}
}

// HasRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime returns true if a handler for RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime has been set.
func (d *PlayerItemRenderedLegibleOutputPushDelegate) HasRenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime() bool {
	return d._RenderedLegibleOutputDidOutputRenderedCaptionImagesForItemTime != nil
}
