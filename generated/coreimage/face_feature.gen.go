// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FaceFeature] class.
var faceFeatureClass = _FaceFeatureClass{objc.GetClass("CIFaceFeature")}

type _FaceFeatureClass struct {
	class objc.Class
}

// Information about a face detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature

type FaceFeature struct {
	Feature
}

// FaceFeatureFrom constructs a [FaceFeature] from an unsafe.Pointer.
//
// Information about a face detected in a still or video image.
func FaceFeatureFrom(ptr unsafe.Pointer) FaceFeature {
	return FaceFeature{
		Feature: FeatureFrom(ptr),
	}
}



