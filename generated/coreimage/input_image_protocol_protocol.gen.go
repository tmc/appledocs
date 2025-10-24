// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

// PinputImage is the inputImage protocol interface.
//
// The input image whose red channel defines a mask. If the red channel pixel value is greater than 0.5 then the point is considered in the mask and output pixel will be zero. Otherwise the output pixel will be a value between zero and one.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.4+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.coreimage/documentation/CoreImage/CIDistanceGradientFromRedMask/inputImage
type PinputImage interface {
}
