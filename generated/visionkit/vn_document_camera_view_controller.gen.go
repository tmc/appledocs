// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DocumentCameraViewController] class.
var (
	DocumentCameraViewControllerClass     _DocumentCameraViewControllerClass
	DocumentCameraViewControllerClassOnce sync.Once
)

func getDocumentCameraViewControllerClass() _DocumentCameraViewControllerClass {
	DocumentCameraViewControllerClassOnce.Do(func() {
		DocumentCameraViewControllerClass = _DocumentCameraViewControllerClass{objc.GetClass("VNDocumentCameraViewController")}
	})
	return DocumentCameraViewControllerClass
}

type _DocumentCameraViewControllerClass struct {
	class objc.Class
}





// An interface definition for the [DocumentCameraViewController] class.
type IDocumentCameraViewController interface {
	IViewController
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DocumentCameraViewControllerClass) Alloc() DocumentCameraViewController {
	rv := objc.Send[DocumentCameraViewController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DocumentCameraViewControllerClass) New() DocumentCameraViewController {
	rv := objc.Send[DocumentCameraViewController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentCameraViewController) Init() DocumentCameraViewController {
	rv := objc.Send[DocumentCameraViewController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentCameraViewController) Autorelease() DocumentCameraViewController {
	rv := objc.Send[DocumentCameraViewController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentCameraViewController creates a new DocumentCameraViewController instance.
func NewDocumentCameraViewController() DocumentCameraViewController {
	return getDocumentCameraViewControllerClass().New()
}





// An object that presents UI for a camera pass-through that helps people scan physical documents.


// An object that presents UI for a camera pass-through that helps people scan physical documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraViewController
type DocumentCameraViewController struct {
	ViewController
}

// DocumentCameraViewControllerFrom constructs a [DocumentCameraViewController] from an unsafe.Pointer.
//
// An object that presents UI for a camera pass-through that helps people scan physical documents.
func DocumentCameraViewControllerFrom(ptr unsafe.Pointer) DocumentCameraViewController {
	return DocumentCameraViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}















// A Boolean variable that indicates whether or not the current device supports document scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraViewController/isSupported
func (dc _DocumentCameraViewControllerClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("supported"))
	return rv
}
















