// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// ignoreProtocol is the ignore protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to ignore.
var ignoreProtocol *objc.Protocol

func init() {
	ignoreProtocol = objc.GetProtocol("ignore")
}
