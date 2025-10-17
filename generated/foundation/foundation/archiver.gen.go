// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Archiver] class.
var ArchiverClass objc.Class

func init() {
	ArchiverClass = objc.GetClass("NSArchiver")
}

type Archiver struct {
	objc.ID
}

func ArchiverFrom(ptr unsafe.Pointer) Archiver {
	return Archiver{
		ID: objc.ID(ptr),
	}
}




