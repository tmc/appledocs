// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import "github.com/ebitengine/purego/objc"

// widgetSearchTermClearedProtocol is the widgetSearchTermCleared: protocol.
//
// Availability:
//   - macOS 10.10+ (Deprecated in 11.0)
//
// Use this protocol when registering custom classes that conform to widgetSearchTermCleared:.
var widgetSearchTermClearedProtocol *objc.Protocol

func init() {
	widgetSearchTermClearedProtocol = objc.GetProtocol("widgetSearchTermCleared:")
}
