// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// classProtocol is the class protocol.
//
// Use this protocol when registering custom classes that conform to class.
var classProtocol *objc.Protocol

func init() {
	classProtocol = objc.GetProtocol("class")
}
