// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Feature] class.
var featureClass = _FeatureClass{objc.GetClass("CIFeature")}

type _FeatureClass struct {
	class objc.Class
}

// The abstract superclass for objects representing notable features detected in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFeature

type Feature struct {
	objectivec.Object
}

// FeatureFrom constructs a [Feature] from an unsafe.Pointer.
//
// The abstract superclass for objects representing notable features detected in an image.
func FeatureFrom(ptr unsafe.Pointer) Feature {
	return Feature{objectivec.Object{objc.ID(ptr)}}
}



