// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import "github.com/ebitengine/purego/objc"

// URLProtocol is the URL protocol.
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - visionOS 1.1+
//
// Use this protocol when registering custom classes that conform to URL.
var URLProtocol *objc.Protocol

func init() {
	URLProtocol = objc.GetProtocol("URL")
}


