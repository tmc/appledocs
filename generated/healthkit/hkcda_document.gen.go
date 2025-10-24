// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKCDADocument */


/* debug [class_header]: Header for HKCDADocument */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCDADocument */
// An interface definition for the [HKCDADocument] class.
type IHKCDADocument interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKCDADocument */
	// properties:
	AuthorName() objc.IObject /* cross-framework: NSString */
	CustodianName() objc.IObject /* cross-framework: NSString */
	DocumentData() objc.IObject /* cross-framework: NSData */
	PatientName() objc.IObject /* cross-framework: NSString */
	Title() objc.IObject /* cross-framework: NSString */
	Document() IHKCDADocument
	SetDocument(value IHKCDADocument)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCDADocument */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCDADocument */
// Alloc allocates a new instance without initialization.
func (hc _HKCDADocumentClass) Alloc() HKCDADocument {
	rv := objc.Send[HKCDADocument](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCDADocument */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCDADocument *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCDADocument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCDADocument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCDADocument */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCDADocument */

// The document’s author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument/authorName
func (h_ HKCDADocument) AuthorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("authorName"))
	return rv
}/* debug [instance_properties/getter]: authorName */


// The name of the organization responsible for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument/custodianName
func (h_ HKCDADocument) CustodianName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("custodianName"))
	return rv
}/* debug [instance_properties/getter]: custodianName */


// The CDA document stored as XML data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument/documentData
func (h_ HKCDADocument) DocumentData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](h_.ID, objc.Sel("documentData"))
	return rv
}/* debug [instance_properties/getter]: documentData */


// The patient’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument/patientName
func (h_ HKCDADocument) PatientName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("patientName"))
	return rv
}/* debug [instance_properties/getter]: patientName */


// The document’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocument/title
func (h_ HKCDADocument) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The CDA document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocumentsample/document
func (h_ HKCDADocument) Document() IHKCDADocument {
	rv := objc.Send[HKCDADocument](h_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// The CDA document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcdadocumentsample/document
func (h_ HKCDADocument) SetDocument(value IHKCDADocument) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDocument:"), value)
}/* debug [instance_properties/setter]: document */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCDADocument */



