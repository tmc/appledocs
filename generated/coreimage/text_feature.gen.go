// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextFeature] class.
var textFeatureClass = _TextFeatureClass{objc.GetClass("CITextFeature")}

type _TextFeatureClass struct {
	class objc.Class
}

// Information about a text that was detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature

type TextFeature struct {
	Feature
}

// TextFeatureFrom constructs a [TextFeature] from an unsafe.Pointer.
//
// Information about a text that was detected in a still or video image.
func TextFeatureFrom(ptr unsafe.Pointer) TextFeature {
	return TextFeature{
		Feature: FeatureFrom(ptr),
	}
}



