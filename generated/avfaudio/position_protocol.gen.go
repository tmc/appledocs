// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// positionProtocol is the position protocol.
//
// Use this protocol when registering custom classes that conform to position.
var positionProtocol *objc.Protocol

func init() {
	positionProtocol = objc.GetProtocol("position")
}
