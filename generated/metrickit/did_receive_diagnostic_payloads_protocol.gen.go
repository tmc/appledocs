// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import "github.com/ebitengine/purego/objc"

// didReceiveDiagnosticPayloadsProtocol is the didReceiveDiagnosticPayloads: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didReceiveDiagnosticPayloads:.
var didReceiveDiagnosticPayloadsProtocol *objc.Protocol

func init() {
	didReceiveDiagnosticPayloadsProtocol = objc.GetProtocol("didReceiveDiagnosticPayloads:")
}
