// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DocumentController] class.
var (
	DocumentControllerClass     _DocumentControllerClass
	DocumentControllerClassOnce sync.Once
)

func getDocumentControllerClass() _DocumentControllerClass {
	DocumentControllerClassOnce.Do(func() {
		DocumentControllerClass = _DocumentControllerClass{objc.GetClass("NSDocumentController")}
	})
	return DocumentControllerClass
}

type _DocumentControllerClass struct {
	class objc.Class
}

// An interface definition for the [DocumentController] class.
type IDocumentController interface {
	objectivec.IObject
	BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer)
	MakeDocumentWithContentsOfURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) unsafe.Pointer
	OpenDocumentWithContentsOfURLDisplayCompletionHandler(url unsafe.Pointer, displayDocument bool, completionHandler unsafe.Pointer)
}

// An object that manages an app’s documents.
//
// As the first-responder target of New and Open menu commands, creates and opens documents and tracks them throughout a session of the app. When opening documents, a document controller runs and manages the modal Open panel. objects also maintain and manage the mappings of document types, extensions, and subclasses as specified in the property loaded from the information property list ( ). You can use various methods to get a list of the current documents, get the current document (which is the document whose window is currently key), get documents based on a given filename or window, and find out about a document’s extension, type, display name, and document class. In some situations, it’s worthwhile to subclass in non- -based apps to get some of its features. For example, the management of the Open Recent menu is useful in apps that don’t use subclasses of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController
type DocumentController struct {
	objectivec.Object
}

// DocumentControllerFrom constructs a [DocumentController] from an unsafe.Pointer.
//
// An object that manages an app’s documents.
func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentControllerClass) Alloc() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DocumentControllerClass) New() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentController) Init() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentController) Autorelease() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentController creates a new DocumentController instance.
func NewDocumentController() DocumentController {
	return getDocumentControllerClass().New()
}

// Presents an Open dialog and delivers the results to a completion handler as an array of URLs for the chosen files, or nil.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(completionHandler:)
func (d_ DocumentController) BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanelWithCompletionHandler:"), completionHandler)
}

// Instantiates a document located by a URL, of a specified type, and returns it if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentWithContentsOfURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(withContentsOf:display:completionHandler:)
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplayCompletionHandler(url unsafe.Pointer, displayDocument bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:completionHandler:"), url, displayDocument, completionHandler)
}

// The list of recent-document URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/recentDocumentURLs
func (d_ DocumentController) RecentDocumentURLs() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](d_.ID, objc.Sel("recentDocumentURLs"))
	return rv
}
