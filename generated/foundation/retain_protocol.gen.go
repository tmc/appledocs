// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// retainProtocol is the retain protocol.
//
// Use this protocol when registering custom classes that conform to retain.
var retainProtocol *objc.Protocol

func init() {
	retainProtocol = objc.GetProtocol("retain")
}
