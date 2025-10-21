// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// rateProtocol is the rate protocol.
//
// Use this protocol when registering custom classes that conform to rate.
var rateProtocol *objc.Protocol

func init() {
	rateProtocol = objc.GetProtocol("rate")
}
