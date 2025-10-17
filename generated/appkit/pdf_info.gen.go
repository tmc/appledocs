// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFInfo] class.
var PDFInfoClass objc.Class

func init() {
	PDFInfoClass = objc.GetClass("NSPDFInfo")
}

type PDFInfo struct {
	objc.ID
}

func PDFInfoFrom(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{
		ID: objc.ID(ptr),
	}
}



