// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFConfiguration] class.
var (
	PDFConfigurationClass     _PDFConfigurationClass
	PDFConfigurationClassOnce sync.Once
)

func getPDFConfigurationClass() _PDFConfigurationClass {
	PDFConfigurationClassOnce.Do(func() {
		PDFConfigurationClass = _PDFConfigurationClass{objc.GetClass("WKPDFConfiguration")}
	})
	return PDFConfigurationClass
}

type _PDFConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PDFConfiguration] class.
type IPDFConfiguration interface {
	objectivec.IObject
}

// The configuration data to use when generating a PDF representation of a web view’s contents.
//
// Create a object when you want to generate a PDF version of your web view’s content. Use this object to specify the portion of the web view to capture. To generate the PDF content, pass the configuration object to the method of , which returns the PDF data for you to use.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration
type PDFConfiguration struct {
	objectivec.Object
}

// PDFConfigurationFrom constructs a [PDFConfiguration] from an unsafe.Pointer.
//
// The configuration data to use when generating a PDF representation of a web view’s contents.
func PDFConfigurationFrom(ptr unsafe.Pointer) PDFConfiguration {
	return PDFConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFConfigurationClass) Alloc() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFConfigurationClass) New() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFConfiguration) Init() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFConfiguration) Autorelease() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFConfiguration creates a new PDFConfiguration instance.
func NewPDFConfiguration() PDFConfiguration {
	return getPDFConfigurationClass().New()
}


// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpdfconfiguration/rect-2a0vp
func (p_ PDFConfiguration) Rect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("rect"))
	return rv
}


// SetRect sets the value of the rect property.
// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpdfconfiguration/rect-2a0vp
func (p_ PDFConfiguration) SetRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRect:"), value)
}

// A Boolean value that indicates whether the PDF may have a transparent background.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/allowTransparentBackground
func (p_ PDFConfiguration) AllowTransparentBackground() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowTransparentBackground"))
	return rv
}


// SetAllowTransparentBackground sets the value of the allowTransparentBackground property.
// A Boolean value that indicates whether the PDF may have a transparent background.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/allowTransparentBackground
func (p_ PDFConfiguration) SetAllowTransparentBackground(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowTransparentBackground:"), value)
}



