// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// zoneProtocol is the zone protocol.
//
// Use this protocol when registering custom classes that conform to zone.
var zoneProtocol *objc.Protocol

func init() {
	zoneProtocol = objc.GetProtocol("zone")
}
