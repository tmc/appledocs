// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import "github.com/ebitengine/purego/objc"

// TelephonyNetworkInfoDelegateProtocol is the CTTelephonyNetworkInfoDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//
// Use this protocol when registering custom classes that conform to CTTelephonyNetworkInfoDelegate.
var TelephonyNetworkInfoDelegateProtocol *objc.Protocol

func init() {
	TelephonyNetworkInfoDelegateProtocol = objc.GetProtocol("CTTelephonyNetworkInfoDelegate")
}

