// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// occlusionProtocol is the occlusion protocol.
//
// Use this protocol when registering custom classes that conform to occlusion.
var occlusionProtocol *objc.Protocol

func init() {
	occlusionProtocol = objc.GetProtocol("occlusion")
}
