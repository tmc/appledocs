// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import "github.com/ebitengine/purego/objc"

// queryProtocol is the query: protocol.
//
// Availability:
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to query:.
var queryProtocol *objc.Protocol

func init() {
	queryProtocol = objc.GetProtocol("query:")
}

