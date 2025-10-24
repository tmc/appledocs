
// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

// Package pdfkit provides Go bindings for the PDFKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PDFKit without requiring cgo.
package pdfkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PDFKit.framework/PDFKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

