// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// hashProtocol is the hash protocol.
//
// Use this protocol when registering custom classes that conform to hash.
var hashProtocol *objc.Protocol

func init() {
	hashProtocol = objc.GetProtocol("hash")
}
