// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import "github.com/ebitengine/purego/objc"

// paymentQueueShouldShowPriceConsentProtocol is the paymentQueueShouldShowPriceConsent: protocol.
//
// Availability:
//   - Mac Catalyst 13.4+ (Deprecated in 18.0)
//   - iOS 13.4+ (Deprecated in 18.0)
//   - iPadOS 13.4+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to paymentQueueShouldShowPriceConsent:.
var paymentQueueShouldShowPriceConsentProtocol *objc.Protocol

func init() {
	paymentQueueShouldShowPriceConsentProtocol = objc.GetProtocol("paymentQueueShouldShowPriceConsent:")
}
