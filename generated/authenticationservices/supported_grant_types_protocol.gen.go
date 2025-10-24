// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// supportedGrantTypesProtocol is the supportedGrantTypes protocol.
//
// Availability:
//   - macOS 14.0+
//
// Use this protocol when registering custom classes that conform to supportedGrantTypes.
var supportedGrantTypesProtocol *objc.Protocol

func init() {
	supportedGrantTypesProtocol = objc.GetProtocol("supportedGrantTypes")
}

