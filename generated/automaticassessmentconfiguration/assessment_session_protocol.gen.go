// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import "github.com/ebitengine/purego/objc"

// assessmentSessionProtocol is the assessmentSession: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.4+
//   - iPadOS 13.4+
//   - macOS 10.15.4+
//
// Use this protocol when registering custom classes that conform to assessmentSession:.
var assessmentSessionProtocol *objc.Protocol

func init() {
	assessmentSessionProtocol = objc.GetProtocol("assessmentSession:")
}

