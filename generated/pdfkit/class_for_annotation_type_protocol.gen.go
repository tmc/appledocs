// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import "github.com/ebitengine/purego/objc"

// classForAnnotationTypeProtocol is the classForAnnotationType: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to classForAnnotationType:.
var classForAnnotationTypeProtocol *objc.Protocol

func init() {
	classForAnnotationTypeProtocol = objc.GetProtocol("classForAnnotationType:")
}
