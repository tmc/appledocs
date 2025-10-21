// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import "github.com/ebitengine/purego/objc"

// INRequestPaymentIntentHandlingProtocol is the INRequestPaymentIntentHandling protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.2+
//
// Use this protocol when registering custom classes that conform to INRequestPaymentIntentHandling.
var INRequestPaymentIntentHandlingProtocol *objc.Protocol

func init() {
	INRequestPaymentIntentHandlingProtocol = objc.GetProtocol("INRequestPaymentIntentHandling")
}
