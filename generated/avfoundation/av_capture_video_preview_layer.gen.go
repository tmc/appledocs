// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureVideoPreviewLayer] class.
var aVCaptureVideoPreviewLayerClass = _AVCaptureVideoPreviewLayerClass{objc.GetClass("AVCaptureVideoPreviewLayer")}

type _AVCaptureVideoPreviewLayerClass struct {
	class objc.Class
}

// A Core Animation layer that displays video from a camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer

type AVCaptureVideoPreviewLayer struct {
	Layer
}

// AVCaptureVideoPreviewLayerFrom constructs a [AVCaptureVideoPreviewLayer] from an unsafe.Pointer.
//
// A Core Animation layer that displays video from a camera device.
func AVCaptureVideoPreviewLayerFrom(ptr unsafe.Pointer) AVCaptureVideoPreviewLayer {
	return AVCaptureVideoPreviewLayer{
		Layer: LayerFrom(ptr),
	}
}



