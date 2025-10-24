// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
	appkit.IViewController
	// properties:
	Delegate() DocumentCameraViewControllerDelegate /* not a class type */
	SetDelegate(value DocumentCameraViewControllerDelegate /* not a class type */)
	// methods:
}

// An object that presents UI for a camera pass-through that helps people scan physical documents.


// An object that presents UI for a camera pass-through that helps people scan physical documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraViewController
type DocumentCameraViewController struct {
	appkit.ViewController
}

// DocumentCameraViewControllerFrom constructs a [DocumentCameraViewController] from an unsafe.Pointer.
//
// An object that presents UI for a camera pass-through that helps people scan physical documents.
func DocumentCameraViewControllerFrom(ptr unsafe.Pointer) DocumentCameraViewController {
	return DocumentCameraViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentCameraViewControllerClass) Alloc() DocumentCameraViewController {
	rv := objc.Send[DocumentCameraViewController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean variable that indicates whether or not the current device supports document scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraViewController/isSupported
func (dc _DocumentCameraViewControllerClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("supported"))
	return rv
}

// The delegate to be notified when the user saves or cancels the document scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/visionkit/vndocumentcameraviewcontroller/delegate
func (d_ DocumentCameraViewController) Delegate() DocumentCameraViewControllerDelegate /* not a class type */ {
	rv := objc.Send[DocumentCameraViewControllerDelegate](d_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate to be notified when the user saves or cancels the document scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/visionkit/vndocumentcameraviewcontroller/delegate
func (d_ DocumentCameraViewController) SetDelegate(value DocumentCameraViewControllerDelegate /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}


