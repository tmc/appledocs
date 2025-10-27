// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PObject is the NSObject protocol interface.
//
// The group of methods that are fundamental to all Objective-C objects.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.objectivec/documentation/ObjectiveC/NSObjectProtocol
type PObject interface {
	// Required methods
	Autorelease() IObject
	Class() objc.Class
	Release()
	Retain() IObject
	RetainCount() uint
	Zone() IObject
	ConformsToProtocol(aProtocol IProtocol) bool
	IsEqual(object IObject) bool
	IsKindOfClass(aClass objc.Class) bool
	IsMemberOfClass(aClass objc.Class) bool
	IsProxy() bool
	PerformSelector(aSelector objc.SEL) IObject
	PerformSelectorWithObject(aSelector objc.SEL, object IObject) IObject
	PerformSelectorWithObjectWithObject(aSelector objc.SEL, object1 IObject, object2 IObject) IObject
	RespondsToSelector(aSelector objc.SEL) bool
	Self() IObject
}
