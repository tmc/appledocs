// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QRCodeFeature] class.
var qRCodeFeatureClass = _QRCodeFeatureClass{objc.GetClass("CIQRCodeFeature")}

type _QRCodeFeatureClass struct {
	class objc.Class
}

// Information about a Quick Response code detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature

type QRCodeFeature struct {
	Feature
}

// QRCodeFeatureFrom constructs a [QRCodeFeature] from an unsafe.Pointer.
//
// Information about a Quick Response code detected in a still or video image.
func QRCodeFeatureFrom(ptr unsafe.Pointer) QRCodeFeature {
	return QRCodeFeature{
		Feature: FeatureFrom(ptr),
	}
}



