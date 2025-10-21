// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import "github.com/ebitengine/purego/objc"

// PaymentQueueDelegateProtocol is the SKPaymentQueueDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 13.0+ (Deprecated in 18.0)
//   - iPadOS 13.0+ (Deprecated in 18.0)
//   - macOS 10.15+ (Deprecated in 15.0)
//   - tvOS 13.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.2+ (Deprecated in 11.0)
//
// Use this protocol when registering custom classes that conform to SKPaymentQueueDelegate.
var PaymentQueueDelegateProtocol *objc.Protocol

func init() {
	PaymentQueueDelegateProtocol = objc.GetProtocol("SKPaymentQueueDelegate")
}


