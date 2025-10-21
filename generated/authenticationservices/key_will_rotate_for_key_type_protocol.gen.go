// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// keyWillRotateForKeyTypeProtocol is the keyWillRotateForKeyType: protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to keyWillRotateForKeyType:.
var keyWillRotateForKeyTypeProtocol *objc.Protocol

func init() {
	keyWillRotateForKeyTypeProtocol = objc.GetProtocol("keyWillRotateForKeyType:")
}
