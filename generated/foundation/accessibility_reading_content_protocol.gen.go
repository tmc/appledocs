// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// AccessibilityReadingContentProtocol is the UIAccessibilityReadingContent protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to UIAccessibilityReadingContent.
var AccessibilityReadingContentProtocol *objc.Protocol

func init() {
	AccessibilityReadingContentProtocol = objc.GetProtocol("UIAccessibilityReadingContent")
}
