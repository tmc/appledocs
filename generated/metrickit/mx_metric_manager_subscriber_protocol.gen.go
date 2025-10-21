// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import "github.com/ebitengine/purego/objc"

// MXMetricManagerSubscriberProtocol is the MXMetricManagerSubscriber protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MXMetricManagerSubscriber.
var MXMetricManagerSubscriberProtocol *objc.Protocol

func init() {
	MXMetricManagerSubscriberProtocol = objc.GetProtocol("MXMetricManagerSubscriber")
}
