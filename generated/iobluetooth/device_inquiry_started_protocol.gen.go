// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// deviceInquiryStartedProtocol is the deviceInquiryStarted: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to deviceInquiryStarted:.
var deviceInquiryStartedProtocol *objc.Protocol

func init() {
	deviceInquiryStartedProtocol = objc.GetProtocol("deviceInquiryStarted:")
}

