// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import "github.com/ebitengine/purego/objc"

// messageReceivedWithNameProtocol is the messageReceivedWithName: protocol.
//
// Availability:
//   - macOS 10.12+
//
// Use this protocol when registering custom classes that conform to messageReceivedWithName:.
var messageReceivedWithNameProtocol *objc.Protocol

func init() {
	messageReceivedWithNameProtocol = objc.GetProtocol("messageReceivedWithName:")
}
