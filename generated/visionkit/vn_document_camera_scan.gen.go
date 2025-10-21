// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DocumentCameraScan] class.
var (
	DocumentCameraScanClass     _DocumentCameraScanClass
	DocumentCameraScanClassOnce sync.Once
)

func getDocumentCameraScanClass() _DocumentCameraScanClass {
	DocumentCameraScanClassOnce.Do(func() {
		DocumentCameraScanClass = _DocumentCameraScanClass{objc.GetClass("VNDocumentCameraScan")}
	})
	return DocumentCameraScanClass
}

type _DocumentCameraScanClass struct {
	class objc.Class
}

// An interface definition for the [DocumentCameraScan] class.
type IDocumentCameraScan interface {
	objectivec.IObject
}

// A single document scanned in the document camera.
//
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan
type DocumentCameraScan struct {
	objectivec.Object
}

// DocumentCameraScanFrom constructs a [DocumentCameraScan] from an unsafe.Pointer.
//
// A single document scanned in the document camera.
func DocumentCameraScanFrom(ptr unsafe.Pointer) DocumentCameraScan {
	return DocumentCameraScan{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentCameraScanClass) Alloc() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DocumentCameraScanClass) New() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentCameraScan) Init() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentCameraScan) Autorelease() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentCameraScan creates a new DocumentCameraScan instance.
func NewDocumentCameraScan() DocumentCameraScan {
	return getDocumentCameraScanClass().New()
}


// The number of pages in the scanned document.
//
// [Full Topic]: https://developer.apple.com/documentation/visionkit/vndocumentcamerascan/pagecount
func (d_ DocumentCameraScan) PageCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("pageCount"))
	return rv
}


// SetPageCount sets the value of the pageCount property.
// The number of pages in the scanned document.

//
// [Full Topic]: https://developer.apple.com/documentation/visionkit/vndocumentcamerascan/pagecount
func (d_ DocumentCameraScan) SetPageCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPageCount:"), value)
}

// The title of the scanned document.
//
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan/title
func (d_ DocumentCameraScan) Title() string {
	rv := objc.Send[string](d_.ID, objc.Sel("title"))
	return rv
}



