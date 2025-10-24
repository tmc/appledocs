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
	

	// properties:
	AllowedExternalEntityURLs() unsafe.Pointer
	SetAllowedExternalEntityURLs(value unsafe.Pointer)
	ColumnNumber() int
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ExternalEntityResolvingPolicy() XMLParserExternalEntityResolvingPolicy
	SetExternalEntityResolvingPolicy(value XMLParserExternalEntityResolvingPolicy)
	LineNumber() int
	ParserError() IError
	PublicID() IString
	ShouldProcessNamespaces() bool
	SetShouldProcessNamespaces(value bool)
	ShouldReportNamespacePrefixes() bool
	SetShouldReportNamespacePrefixes(value bool)
	ShouldResolveExternalEntities() bool
	SetShouldResolveExternalEntities(value bool)
	SystemID() IString


	

	// methods:
	AbortParsing()
	Parse() bool


}





// Alloc allocates a new instance without initialization.
func (xc _XMLParserClass) Alloc() XMLParser {
	rv := objc.Send[XMLParser](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An event driven parser of XML documents (including DTD declarations).
//
// An notifies its delegate about the items (elements, attributes, CDATA blocks, comments, and so on) that it encounters as it processes an XML document. It does not itself do anything with those parsed items except report them. It also reports parsing errors. For convenience, an object in the following descriptions is sometimes referred to as a parser object. Unless used in a callback, the is a thread-safe class as long as any given instance is only used in one thread.


// An event driven parser of XML documents (including DTD declarations).
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






// Initializes a parser with the XML content referenced by the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/init(contentsOf:)
func NewXMLParserWithContentsOfURL(url IURL) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[XMLParser](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
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

















// Stops the parser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/abortParsing()
func (x_ XMLParser) AbortParsing() {
	objc.Send[objc.ID](x_.ID, objc.Sel("abortParsing"))
}


// Starts the event-driven parsing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/parse()
func (x_ XMLParser) Parse() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("parse"))
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/allowedExternalEntityURLs
func (x_ XMLParser) AllowedExternalEntityURLs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("allowedExternalEntityURLs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/allowedExternalEntityURLs
func (x_ XMLParser) SetAllowedExternalEntityURLs(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAllowedExternalEntityURLs:"), value)
}


// The column number of the XML document being processed by the parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/columnNumber
func (x_ XMLParser) ColumnNumber() int {
	rv := objc.Send[int](x_.ID, objc.Sel("columnNumber"))
	return rv
}


// A delegate object that receives messages about the parsing process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/delegate
func (x_ XMLParser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object that receives messages about the parsing process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/delegate
func (x_ XMLParser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDelegate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/externalEntityResolvingPolicy-swift.property
func (x_ XMLParser) ExternalEntityResolvingPolicy() XMLParserExternalEntityResolvingPolicy {
	rv := objc.Send[XMLParserExternalEntityResolvingPolicy](x_.ID, objc.Sel("externalEntityResolvingPolicy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/externalEntityResolvingPolicy-swift.property
func (x_ XMLParser) SetExternalEntityResolvingPolicy(value XMLParserExternalEntityResolvingPolicy) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExternalEntityResolvingPolicy:"), value)
}


// The line number of the XML document being processed by the parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/lineNumber
func (x_ XMLParser) LineNumber() int {
	rv := objc.Send[int](x_.ID, objc.Sel("lineNumber"))
	return rv
}


// An object from which you can obtain information about a parsing error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/parserError
func (x_ XMLParser) ParserError() IError {
	rv := objc.Send[Error](x_.ID, objc.Sel("parserError"))
	return rv
}


// The public identifier of the external entity referenced in the XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/publicID
func (x_ XMLParser) PublicID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("publicID"))
	return rv
}


// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldProcessNamespaces
func (x_ XMLParser) ShouldProcessNamespaces() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldProcessNamespaces"))
	return rv
}


// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldProcessNamespaces
func (x_ XMLParser) SetShouldProcessNamespaces(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldProcessNamespaces:"), value)
}


// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldReportNamespacePrefixes
func (x_ XMLParser) ShouldReportNamespacePrefixes() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldReportNamespacePrefixes"))
	return rv
}


// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldReportNamespacePrefixes
func (x_ XMLParser) SetShouldReportNamespacePrefixes(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldReportNamespacePrefixes:"), value)
}


// A Boolean value that determines whether the parser reports declarations of external entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldResolveExternalEntities
func (x_ XMLParser) ShouldResolveExternalEntities() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("shouldResolveExternalEntities"))
	return rv
}


// A Boolean value that determines whether the parser reports declarations of external entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/shouldResolveExternalEntities
func (x_ XMLParser) SetShouldResolveExternalEntities(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setShouldResolveExternalEntities:"), value)
}


// The system identifier of the external entity referenced in the XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLParser/systemID
func (x_ XMLParser) SystemID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("systemID"))
	return rv
}







