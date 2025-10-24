// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import "github.com/ebitengine/purego/objc"

// assessmentSessionDidUpdateProtocol is the assessmentSessionDidUpdate: protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 17.5+
//   - iPadOS 17.5+
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to assessmentSessionDidUpdate:.
var assessmentSessionDidUpdateProtocol *objc.Protocol

func init() {
	assessmentSessionDidUpdateProtocol = objc.GetProtocol("assessmentSessionDidUpdate:")
}

