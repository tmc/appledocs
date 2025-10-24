// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

// PMXMetricManagerSubscriber is the MXMetricManagerSubscriber protocol interface.
//
// A protocol defining a method for receiving a daily metrics report.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metrickit/documentation/MetricKit/MXMetricManagerSubscriber
type PMXMetricManagerSubscriber interface {
	// Optional methods
	DidReceiveMetricPayloads(payloads []MXMetricPayload)
	HasDidReceiveMetricPayloads() bool
	DidReceiveDiagnosticPayloads(payloads []MXDiagnosticPayload)
	HasDidReceiveDiagnosticPayloads() bool
}
