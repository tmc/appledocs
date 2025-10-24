// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PAssetReaderCaptionValidationHandling is the AVAssetReaderCaptionValidationHandling protocol interface.
//
// A protocol that defines the methods for caption validation events.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 12.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVAssetReaderCaptionValidationHandling
type PAssetReaderCaptionValidationHandling interface {
	// Optional methods
	CaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements(adaptor IAVAssetReaderOutputCaptionAdaptor, caption IAVCaption, syntaxElements []string)
	HasCaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements() bool
}
