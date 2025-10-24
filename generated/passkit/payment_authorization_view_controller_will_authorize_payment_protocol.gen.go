// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// paymentAuthorizationViewControllerWillAuthorizePaymentProtocol is the paymentAuthorizationViewControllerWillAuthorizePayment: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.3+
//   - iPadOS 8.3+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to paymentAuthorizationViewControllerWillAuthorizePayment:.
var paymentAuthorizationViewControllerWillAuthorizePaymentProtocol *objc.Protocol

func init() {
	paymentAuthorizationViewControllerWillAuthorizePaymentProtocol = objc.GetProtocol("paymentAuthorizationViewControllerWillAuthorizePayment:")
}

