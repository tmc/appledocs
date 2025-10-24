// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// authenticatedContextProtocol is the authenticatedContext protocol.
//
// Availability:
//   - macOS 13.3+
//
// Use this protocol when registering custom classes that conform to authenticatedContext.
var authenticatedContextProtocol *objc.Protocol

func init() {
	authenticatedContextProtocol = objc.GetProtocol("authenticatedContext")
}

