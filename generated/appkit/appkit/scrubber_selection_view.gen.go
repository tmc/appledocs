// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberSelectionView] class.
var ScrubberSelectionViewClass objc.Class

func init() {
	ScrubberSelectionViewClass = objc.GetClass("NSScrubberSelectionView")
}

type ScrubberSelectionView struct {
	objc.ID
}

func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		ID: objc.ID(ptr),
	}
}



