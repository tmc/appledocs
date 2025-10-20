// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// tableViewProtocol is the tableView: protocol.
//
// Availability:
//   - macOS 10.0+
//
// Use this protocol when registering custom classes that conform to tableView:.
var tableViewProtocol *objc.Protocol

func init() {
	tableViewProtocol = objc.GetProtocol("tableView:")
}
