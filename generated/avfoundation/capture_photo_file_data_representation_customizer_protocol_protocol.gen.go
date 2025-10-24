// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PCapturePhotoFileDataRepresentationCustomizer is the AVCapturePhotoFileDataRepresentationCustomizer protocol interface.
//
// A protocol that defines the methods to implement to customize the packaging of photo data.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCapturePhotoFileDataRepresentationCustomizer
type PCapturePhotoFileDataRepresentationCustomizer interface {
	// Optional methods
	ReplacementAppleProRAWCompressionSettingsForPhotoDefaultSettingsMaximumBitDepth(photo IAVCapturePhoto, defaultSettings foundation.IDictionary, maximumBitDepth int) foundation.IDictionary
	HasReplacementAppleProRAWCompressionSettingsForPhotoDefaultSettingsMaximumBitDepth() bool
	ReplacementDepthDataForPhoto(photo IAVCapturePhoto) DepthData
	HasReplacementDepthDataForPhoto() bool
	ReplacementEmbeddedThumbnailPixelBufferWithPhotoFormatForPhoto(replacementEmbeddedThumbnailPhotoFormatOut foundation.IDictionary, photo IAVCapturePhoto) PixelBufferRef
	HasReplacementEmbeddedThumbnailPixelBufferWithPhotoFormatForPhoto() bool
	ReplacementMetadataForPhoto(photo IAVCapturePhoto) foundation.IDictionary
	HasReplacementMetadataForPhoto() bool
	ReplacementPortraitEffectsMatteForPhoto(photo IAVCapturePhoto) PortraitEffectsMatte
	HasReplacementPortraitEffectsMatteForPhoto() bool
	ReplacementSemanticSegmentationMatteOfTypeForPhoto(semanticSegmentationMatteType SemanticSegmentationMatteType /* typedef */, photo IAVCapturePhoto) SemanticSegmentationMatte
	HasReplacementSemanticSegmentationMatteOfTypeForPhoto() bool
}
