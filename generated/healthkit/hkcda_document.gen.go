// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKCDADocument] class.
var (
	HKCDADocumentClass     _HKCDADocumentClass
	HKCDADocumentClassOnce sync.Once
)

func getHKCDADocumentClass() _HKCDADocumentClass {
	HKCDADocumentClassOnce.Do(func() {
		HKCDADocumentClass = _HKCDADocumentClass{objc.GetClass("HKCDADocument")}
	})
	return HKCDADocumentClass
}

type _HKCDADocumentClass struct {
	class objc.Class
}

// An interface definition for the [HKCDADocument] class.
type IHKCDADocument interface {
	objectivec.IObject
	// properties:
	AuthorName() objc.IObject /* cross-framework: NSString */
	SetAuthorName(value objc.IObject /* cross-framework: NSString */)
	CustodianName() objc.IObject /* cross-framework: NSString */
	SetCustodianName(value objc.IObject /* cross-framework: NSString */)
	DocumentData() objc.IObject /* cross-framework: Data */
	SetDocumentData(value objc.IObject /* cross-framework: Data */)
	PatientName() objc.IObject /* cross-framework: NSString */
	SetPatientName(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Document() IHKCDADocument
	SetDocument(value IHKCDADocument)
	// methods:
}

// An object representing a Clinical Document Architecture (CDA) document in HealthKit.
//
// CDA documents use XML to encode clinical documents so that they can be easily exchanged. For more information on the CDA document format, see the standard. Do not instantiate objects directly. Instead, create a new object by calling the method, and passing the CDA’s XML data. HealthKit creates a object for the XML, and assigns it to the sample’s property. objects are immutable. When you create a new document sample, HealthKit parses the title, patient name, author name, and custodian name from the XML to populates the document object’s properties. These properties cannot be changed. Like many HealthKit classes, the class should not be subclassed.


// An object representing a Clinical Document Architecture (CDA) document in HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument
type HKCDADocument struct {
	objectivec.Object
}

// HKCDADocumentFrom constructs a [HKCDADocument] from an unsafe.Pointer.
//
// An object representing a Clinical Document Architecture (CDA) document in HealthKit.
func HKCDADocumentFrom(ptr unsafe.Pointer) HKCDADocument {
	return HKCDADocument{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCDADocumentClass) Alloc() HKCDADocument {
	rv := objc.Send[HKCDADocument](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCDADocumentClass) New() HKCDADocument {
	rv := objc.Send[HKCDADocument](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCDADocument) Init() HKCDADocument {
	rv := objc.Send[HKCDADocument](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCDADocument) Autorelease() HKCDADocument {
	rv := objc.Send[HKCDADocument](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCDADocument creates a new HKCDADocument instance.
func NewHKCDADocument() HKCDADocument {
	return getHKCDADocumentClass().New()
}



// The document’s author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/authorname
func (h_ HKCDADocument) AuthorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("authorName"))
	return rv
}


// The document’s author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/authorname
func (h_ HKCDADocument) SetAuthorName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAuthorName:"), value)
}


// The name of the organization responsible for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/custodianname
func (h_ HKCDADocument) CustodianName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("custodianName"))
	return rv
}


// The name of the organization responsible for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/custodianname
func (h_ HKCDADocument) SetCustodianName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCustodianName:"), value)
}


// The CDA document stored as XML data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/documentdata
func (h_ HKCDADocument) DocumentData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("documentData"))
	return rv
}


// The CDA document stored as XML data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/documentdata
func (h_ HKCDADocument) SetDocumentData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDocumentData:"), value)
}


// The patient’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/patientname
func (h_ HKCDADocument) PatientName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("patientName"))
	return rv
}


// The patient’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/patientname
func (h_ HKCDADocument) SetPatientName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPatientName:"), value)
}


// The document’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/title
func (h_ HKCDADocument) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("title"))
	return rv
}


// The document’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocument/title
func (h_ HKCDADocument) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setTitle:"), value)
}


// The CDA document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocumentsample/document
func (h_ HKCDADocument) Document() IHKCDADocument {
	rv := objc.Send[HKCDADocument](h_.ID, objc.Sel("document"))
	return rv
}


// The CDA document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocumentsample/document
func (h_ HKCDADocument) SetDocument(value IHKCDADocument) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDocument:"), value)
}



