// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import "github.com/ebitengine/purego/objc"

// didReceiveMetricPayloadsProtocol is the didReceiveMetricPayloads: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didReceiveMetricPayloads:.
var didReceiveMetricPayloadsProtocol *objc.Protocol

func init() {
	didReceiveMetricPayloadsProtocol = objc.GetProtocol("didReceiveMetricPayloads:")
}
