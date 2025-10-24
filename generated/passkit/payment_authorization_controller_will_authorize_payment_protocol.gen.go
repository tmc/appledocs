// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// paymentAuthorizationControllerWillAuthorizePaymentProtocol is the paymentAuthorizationControllerWillAuthorizePayment: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// Use this protocol when registering custom classes that conform to paymentAuthorizationControllerWillAuthorizePayment:.
var paymentAuthorizationControllerWillAuthorizePaymentProtocol *objc.Protocol

func init() {
	paymentAuthorizationControllerWillAuthorizePaymentProtocol = objc.GetProtocol("paymentAuthorizationControllerWillAuthorizePayment:")
}

