// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XMLParser] class.
var (
	XMLParserClass     _XMLParserClass
	XMLParserClassOnce sync.Once
)

func getXMLParserClass() _XMLParserClass {
	XMLParserClassOnce.Do(func() {
		XMLParserClass = _XMLParserClass{objc.GetClass("NSXMLParser")}
	})
	return XMLParserClass
}

type _XMLParserClass struct {
	class objc.Class
}

// An interface definition for the [XMLParser] class.
type IXMLParser interface {
	objectivec.IObject
	Parse() bool
	AllowedExternalEntityURLs() unsafe.Pointer
	SetAllowedExternalEntityURLs(value unsafe.Pointer)
	LineNumber() int
	ParserError() NSError
	PublicID() string
	ShouldReportNamespacePrefixes() bool
	SetShouldReportNamespacePrefixes(value bool)
	ColumnNumber() int
	SetColumnNumber(value int)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ExternalEntityResolvingPolicy() unsafe.Pointer
	SetExternalEntityResolvingPolicy(value unsafe.Pointer)
	ShouldProcessNamespaces() bool
	SetShouldProcessNamespaces(value bool)
	ShouldResolveExternalEntities() bool
	SetShouldResolveExternalEntities(value bool)
	SystemID() string
	SetSystemID(value string)
}

// An event driven parser of XML documents (including DTD declarations).
//
// An notifies its delegate about the items (elements, attributes, CDATA blocks, comments, and so on) that it encounters as it processes an XML document. It does not itself do anything with those parsed items except report them. It also reports parsing errors. For convenience, an object in the following descriptions is sometimes referred to as a parser object. Unless used in a callback, the is a thread-safe class as long as any given instance is only used in one thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser
type XMLParser struct {
	objectivec.Object
}

// XMLParserFrom constructs a [XMLParser] from an unsafe.Pointer.
//
// An event driven parser of XML documents (including DTD declarations).
func XMLParserFrom(ptr unsafe.Pointer) XMLParser {
	return XMLParser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XMLParserClass) Alloc() XMLParser {
	rv := objc.Send[XMLParser](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLParserClass) New() XMLParser {
	rv := objc.Send[XMLParser](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLParser) Init() XMLParser {
	rv := objc.Send[XMLParser](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLParser) Autorelease() XMLParser {
	rv := objc.Send[XMLParser](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLParser creates a new XMLParser instance.
func NewXMLParser() XMLParser {
	return getXMLParserClass().New()
}





// Initializes a parser with the XML contents encapsulated in a given data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/init(data:)

func NewXMLParserWithData(data IData) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[XMLParser](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}




// Initializes a parser with the XML contents from the specified stream and parses it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/init(stream:)

func NewXMLParserWithStream(stream IInputStream) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[XMLParser](instance.ID, objc.Sel("initWithStream:"), stream)
	rv.Autorelease()
	return rv
}


// Starts the event-driven parsing operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/parse()
func (x_ XMLParser) Parse() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("parse"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/allowedExternalEntityURLs
func (x_ XMLParser) AllowedExternalEntityURLs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("allowedExternalEntityURLs"))
	return rv
}


// SetAllowedExternalEntityURLs sets the value of the allowedExternalEntityURLs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/allowedExternalEntityURLs
func (x_ XMLParser) SetAllowedExternalEntityURLs(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAllowedExternalEntityURLs:"), value)
}

// The line number of the XML document being processed by the parser.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/lineNumber
func (x_ XMLParser) LineNumber() int {
	rv := objc.Send[int](x_.ID, objc.Sel("lineNumber"))
	return rv
}

// An object from which you can obtain information about a parsing error.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/parserError
func (x_ XMLParser) ParserError() NSError {
	rv := objc.Send[NSError](x_.ID, objc.Sel("parserError"))
	return rv
}

// The public identifier of the external entity referenced in the XML document.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/publicID
func (x_ XMLParser) PublicID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("publicID"))
	return rv
}

// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldReportNamespacePrefixes
func (x_ XMLParser) ShouldReportNamespacePrefixes() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldReportNamespacePrefixes"))
	return rv
}


// SetShouldReportNamespacePrefixes sets the value of the shouldReportNamespacePrefixes property.
// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldReportNamespacePrefixes
func (x_ XMLParser) SetShouldReportNamespacePrefixes(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldReportNamespacePrefixes:"), value)
}

// The column number of the XML document being processed by the parser.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/columnnumber
func (x_ XMLParser) ColumnNumber() int {
	rv := objc.Send[int](x_.ID, objc.Sel("columnNumber"))
	return rv
}


// SetColumnNumber sets the value of the columnNumber property.
// The column number of the XML document being processed by the parser.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/columnnumber
func (x_ XMLParser) SetColumnNumber(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setColumnNumber:"), value)
}

// A delegate object that receives messages about the parsing process.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/delegate
func (x_ XMLParser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate object that receives messages about the parsing process.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/delegate
func (x_ XMLParser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDelegate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/externalentityresolvingpolicy-swift.property
func (x_ XMLParser) ExternalEntityResolvingPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("externalEntityResolvingPolicy"))
	return rv
}


// SetExternalEntityResolvingPolicy sets the value of the externalEntityResolvingPolicy property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/externalentityresolvingpolicy-swift.property
func (x_ XMLParser) SetExternalEntityResolvingPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExternalEntityResolvingPolicy:"), value)
}

// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/shouldprocessnamespaces
func (x_ XMLParser) ShouldProcessNamespaces() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldProcessNamespaces"))
	return rv
}


// SetShouldProcessNamespaces sets the value of the shouldProcessNamespaces property.
// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/shouldprocessnamespaces
func (x_ XMLParser) SetShouldProcessNamespaces(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldProcessNamespaces:"), value)
}

// A Boolean value that determines whether the parser reports declarations of external entities.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/shouldresolveexternalentities
func (x_ XMLParser) ShouldResolveExternalEntities() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldResolveExternalEntities"))
	return rv
}


// SetShouldResolveExternalEntities sets the value of the shouldResolveExternalEntities property.
// A Boolean value that determines whether the parser reports declarations of external entities.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/shouldresolveexternalentities
func (x_ XMLParser) SetShouldResolveExternalEntities(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldResolveExternalEntities:"), value)
}

// The system identifier of the external entity referenced in the XML document.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/systemid
func (x_ XMLParser) SystemID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("systemID"))
	return rv
}


// SetSystemID sets the value of the systemID property.
// The system identifier of the external entity referenced in the XML document.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlparser/systemid
func (x_ XMLParser) SetSystemID(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), objc.String(value))
}


