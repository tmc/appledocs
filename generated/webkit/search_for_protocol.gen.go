// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// searchForProtocol is the searchFor: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to searchFor:.
var searchForProtocol *objc.Protocol

func init() {
	searchForProtocol = objc.GetProtocol("searchFor:")
}
