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

// An interface definition for the [TextFeature] class.
type ITextFeature interface {
	IFeature
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
// Alloc allocates a new instance without initialization.
func (tc _TextFeatureClass) Alloc() TextFeature {
	rv := objc.Send[TextFeature](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextFeatureClass) New() TextFeature {
	rv := objc.Send[TextFeature](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextFeature) Init() TextFeature {
	rv := objc.Send[TextFeature](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextFeature) Autorelease() TextFeature {
	rv := objc.Send[TextFeature](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextFeature creates a new TextFeature instance.
func NewTextFeature() TextFeature {
	return textFeatureClass.New()
}




