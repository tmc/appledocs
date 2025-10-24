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
	Autorelease() IObject/* debug [protocol_interface/required_method]: Autorelease */
	Class() objc.Class/* debug [protocol_interface/required_method]: Class */
	Release()/* debug [protocol_interface/required_method]: Release */
	Retain() IObject/* debug [protocol_interface/required_method]: Retain */
	RetainCount() uint/* debug [protocol_interface/required_method]: RetainCount */
	Zone() IObject/* debug [protocol_interface/required_method]: Zone */
	ConformsToProtocol(aProtocol IProtocol) bool/* debug [protocol_interface/required_method]: ConformsToProtocol */
	IsEqual(object IObject) bool/* debug [protocol_interface/required_method]: IsEqual */
	IsKindOfClass(aClass objc.Class) bool/* debug [protocol_interface/required_method]: IsKindOfClass */
	IsMemberOfClass(aClass objc.Class) bool/* debug [protocol_interface/required_method]: IsMemberOfClass */
	IsProxy() bool/* debug [protocol_interface/required_method]: IsProxy */
	PerformSelector(aSelector objc.SEL) objc.ID/* debug [protocol_interface/required_method]: PerformSelector */
	PerformSelectorWithObject(aSelector objc.SEL, object IObject) objc.ID/* debug [protocol_interface/required_method]: PerformSelectorWithObject */
	PerformSelectorWithObjectWithObject(aSelector objc.SEL, object1 IObject, object2 IObject) objc.ID/* debug [protocol_interface/required_method]: PerformSelectorWithObjectWithObject */
	RespondsToSelector(aSelector objc.SEL) bool/* debug [protocol_interface/required_method]: RespondsToSelector */
	Self() IObject/* debug [protocol_interface/required_method]: Self */
}
