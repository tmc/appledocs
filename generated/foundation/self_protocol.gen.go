// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// selfProtocol is the self protocol.
//
// Use this protocol when registering custom classes that conform to self.
var selfProtocol *objc.Protocol

func init() {
	selfProtocol = objc.GetProtocol("self")
}
