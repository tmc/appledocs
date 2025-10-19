// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RectangleFeature] class.
var rectangleFeatureClass = _RectangleFeatureClass{objc.GetClass("CIRectangleFeature")}

type _RectangleFeatureClass struct {
	class objc.Class
}

// Information about a rectangular region detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature

type RectangleFeature struct {
	Feature
}

// RectangleFeatureFrom constructs a [RectangleFeature] from an unsafe.Pointer.
//
// Information about a rectangular region detected in a still or video image.
func RectangleFeatureFrom(ptr unsafe.Pointer) RectangleFeature {
	return RectangleFeature{
		Feature: FeatureFrom(ptr),
	}
}



