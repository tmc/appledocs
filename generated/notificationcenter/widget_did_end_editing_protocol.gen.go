// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import "github.com/ebitengine/purego/objc"

// widgetDidEndEditingProtocol is the widgetDidEndEditing protocol.
//
// Availability:
//   - macOS 10.10+ (Deprecated in 11.0)
//
// Use this protocol when registering custom classes that conform to widgetDidEndEditing.
var widgetDidEndEditingProtocol *objc.Protocol

func init() {
	widgetDidEndEditingProtocol = objc.GetProtocol("widgetDidEndEditing")
}
