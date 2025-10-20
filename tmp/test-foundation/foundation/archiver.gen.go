// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ArchiverClass _ArchiverClass

func init() {
	ArchiverClass = _ArchiverClass{objc.GetClass("NSArchiver")}
}

type _ArchiverClass struct {
	class objc.Class
}

type Archiver struct {
	objc.ID
}

func ArchiverFrom(ptr unsafe.Pointer) Archiver {
	return Archiver{
		ID: objc.ID(ptr),
	}
}




