// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import "github.com/ebitengine/purego/objc"

// safariViewControllerWillOpenInBrowserProtocol is the safariViewControllerWillOpenInBrowser: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//
// Use this protocol when registering custom classes that conform to safariViewControllerWillOpenInBrowser:.
var safariViewControllerWillOpenInBrowserProtocol *objc.Protocol

func init() {
	safariViewControllerWillOpenInBrowserProtocol = objc.GetProtocol("safariViewControllerWillOpenInBrowser:")
}
