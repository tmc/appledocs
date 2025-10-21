// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import "github.com/ebitengine/purego/objc"

// authorizedToStartStreamForClientProtocol is the authorizedToStartStreamForClient: protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to authorizedToStartStreamForClient:.
var authorizedToStartStreamForClientProtocol *objc.Protocol

func init() {
	authorizedToStartStreamForClientProtocol = objc.GetProtocol("authorizedToStartStreamForClient:")
}
