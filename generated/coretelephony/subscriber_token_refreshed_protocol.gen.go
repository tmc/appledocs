// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import "github.com/ebitengine/purego/objc"

// subscriberTokenRefreshedProtocol is the subscriberTokenRefreshed: protocol.
//
// Availability:
//   - iOS 12.1+
//   - iPadOS 12.1+
//
// Use this protocol when registering custom classes that conform to subscriberTokenRefreshed:.
var subscriberTokenRefreshedProtocol *objc.Protocol

func init() {
	subscriberTokenRefreshedProtocol = objc.GetProtocol("subscriberTokenRefreshed:")
}
