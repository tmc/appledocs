// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import "github.com/ebitengine/purego/objc"

// assessmentSessionDidBeginProtocol is the assessmentSessionDidBegin: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.4+
//   - iPadOS 13.4+
//   - macOS 10.15.4+
//
// Use this protocol when registering custom classes that conform to assessmentSessionDidBegin:.
var assessmentSessionDidBeginProtocol *objc.Protocol

func init() {
	assessmentSessionDidBeginProtocol = objc.GetProtocol("assessmentSessionDidBegin:")
}

