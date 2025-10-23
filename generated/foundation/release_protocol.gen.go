// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// releaseProtocol is the release protocol.
//
// Use this protocol when registering custom classes that conform to release.
var releaseProtocol *objc.Protocol

func init() {
	releaseProtocol = objc.GetProtocol("release")
}
