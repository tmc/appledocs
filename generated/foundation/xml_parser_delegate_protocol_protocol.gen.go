// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PXMLParserDelegate is the NSXMLParserDelegate protocol interface.
//
// The interface an XML parser uses to inform its delegate about the content of the parsed document.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/XMLParserDelegate
type PXMLParserDelegate interface {
	// Optional methods
	ParserDidEndElementNamespaceURIQualifiedName(parser IXMLParser, elementName IString, namespaceURI IString, qName IString)
	HasParserDidEndElementNamespaceURIQualifiedName() bool
	ParserDidEndMappingPrefix(parser IXMLParser, prefix IString)
	HasParserDidEndMappingPrefix() bool
	ParserDidStartElementNamespaceURIQualifiedNameAttributes(parser IXMLParser, elementName IString, namespaceURI IString, qName IString, attributeDict IDictionary)
	HasParserDidStartElementNamespaceURIQualifiedNameAttributes() bool
	ParserDidStartMappingPrefixToURI(parser IXMLParser, prefix IString, namespaceURI IString)
	HasParserDidStartMappingPrefixToURI() bool
	ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue(parser IXMLParser, attributeName IString, elementName IString, type_ IString, defaultValue IString)
	HasParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue() bool
	ParserFoundCDATA(parser IXMLParser, CDATABlock IData)
	HasParserFoundCDATA() bool
	ParserFoundCharacters(parser IXMLParser, string_ IString)
	HasParserFoundCharacters() bool
	ParserFoundComment(parser IXMLParser, comment IString)
	HasParserFoundComment() bool
	ParserFoundElementDeclarationWithNameModel(parser IXMLParser, elementName IString, model IString)
	HasParserFoundElementDeclarationWithNameModel() bool
	ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID(parser IXMLParser, name IString, publicID IString, systemID IString)
	HasParserFoundExternalEntityDeclarationWithNamePublicIDSystemID() bool
	ParserFoundIgnorableWhitespace(parser IXMLParser, whitespaceString IString)
	HasParserFoundIgnorableWhitespace() bool
	ParserFoundInternalEntityDeclarationWithNameValue(parser IXMLParser, name IString, value IString)
	HasParserFoundInternalEntityDeclarationWithNameValue() bool
	ParserFoundNotationDeclarationWithNamePublicIDSystemID(parser IXMLParser, name IString, publicID IString, systemID IString)
	HasParserFoundNotationDeclarationWithNamePublicIDSystemID() bool
	ParserFoundProcessingInstructionWithTargetData(parser IXMLParser, target IString, data IString)
	HasParserFoundProcessingInstructionWithTargetData() bool
	ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName(parser IXMLParser, name IString, publicID IString, systemID IString, notationName IString)
	HasParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName() bool
	ParserParseErrorOccurred(parser IXMLParser, parseError IError)
	HasParserParseErrorOccurred() bool
	ParserResolveExternalEntityNameSystemID(parser IXMLParser, name IString, systemID IString) Data
	HasParserResolveExternalEntityNameSystemID() bool
	ParserValidationErrorOccurred(parser IXMLParser, validationError IError)
	HasParserValidationErrorOccurred() bool
	ParserDidEndDocument(parser IXMLParser)
	HasParserDidEndDocument() bool
	ParserDidStartDocument(parser IXMLParser)
	HasParserDidStartDocument() bool
}

// XMLParserDelegate is a delegate implementation builder for the PXMLParserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type XMLParserDelegate struct {
	_ParserDidEndElementNamespaceURIQualifiedName func(parser IXMLParser, elementName IString, namespaceURI IString, qName IString)
	_ParserDidEndMappingPrefix func(parser IXMLParser, prefix IString)
	_ParserDidStartElementNamespaceURIQualifiedNameAttributes func(parser IXMLParser, elementName IString, namespaceURI IString, qName IString, attributeDict IDictionary)
	_ParserDidStartMappingPrefixToURI func(parser IXMLParser, prefix IString, namespaceURI IString)
	_ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue func(parser IXMLParser, attributeName IString, elementName IString, type_ IString, defaultValue IString)
	_ParserFoundCDATA func(parser IXMLParser, CDATABlock IData)
	_ParserFoundCharacters func(parser IXMLParser, string_ IString)
	_ParserFoundComment func(parser IXMLParser, comment IString)
	_ParserFoundElementDeclarationWithNameModel func(parser IXMLParser, elementName IString, model IString)
	_ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID func(parser IXMLParser, name IString, publicID IString, systemID IString)
	_ParserFoundIgnorableWhitespace func(parser IXMLParser, whitespaceString IString)
	_ParserFoundInternalEntityDeclarationWithNameValue func(parser IXMLParser, name IString, value IString)
	_ParserFoundNotationDeclarationWithNamePublicIDSystemID func(parser IXMLParser, name IString, publicID IString, systemID IString)
	_ParserFoundProcessingInstructionWithTargetData func(parser IXMLParser, target IString, data IString)
	_ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName func(parser IXMLParser, name IString, publicID IString, systemID IString, notationName IString)
	_ParserParseErrorOccurred func(parser IXMLParser, parseError IError)
	_ParserResolveExternalEntityNameSystemID func(parser IXMLParser, name IString, systemID IString) Data
	_ParserValidationErrorOccurred func(parser IXMLParser, validationError IError)
	_ParserDidEndDocument func(parser IXMLParser)
	_ParserDidStartDocument func(parser IXMLParser)
}

// SetParserDidEndElementNamespaceURIQualifiedName sets the handler for the ParserDidEndElementNamespaceURIQualifiedName delegate method.
//
// Sent by a parser object to its delegate when it encounters an end tag for a specific element.
func (d *XMLParserDelegate) SetParserDidEndElementNamespaceURIQualifiedName(f func(parser IXMLParser, elementName IString, namespaceURI IString, qName IString)) {
	d._ParserDidEndElementNamespaceURIQualifiedName = f
}

// SetParserDidEndMappingPrefix sets the handler for the ParserDidEndMappingPrefix delegate method.
//
// Sent by a parser object to its delegate when a given namespace prefix goes out of scope.
func (d *XMLParserDelegate) SetParserDidEndMappingPrefix(f func(parser IXMLParser, prefix IString)) {
	d._ParserDidEndMappingPrefix = f
}

// SetParserDidStartElementNamespaceURIQualifiedNameAttributes sets the handler for the ParserDidStartElementNamespaceURIQualifiedNameAttributes delegate method.
//
// Sent by a parser object to its delegate when it encounters a start tag for a given element.
func (d *XMLParserDelegate) SetParserDidStartElementNamespaceURIQualifiedNameAttributes(f func(parser IXMLParser, elementName IString, namespaceURI IString, qName IString, attributeDict IDictionary)) {
	d._ParserDidStartElementNamespaceURIQualifiedNameAttributes = f
}

// SetParserDidStartMappingPrefixToURI sets the handler for the ParserDidStartMappingPrefixToURI delegate method.
//
// Sent by a parser object to its delegate the first time it encounters a given namespace prefix, which is mapped to a URI.
func (d *XMLParserDelegate) SetParserDidStartMappingPrefixToURI(f func(parser IXMLParser, prefix IString, namespaceURI IString)) {
	d._ParserDidStartMappingPrefixToURI = f
}

// SetParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue sets the handler for the ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue delegate method.
//
// Sent by a parser object to its delegate when it encounters a declaration of an attribute that is associated with a specific element.
func (d *XMLParserDelegate) SetParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue(f func(parser IXMLParser, attributeName IString, elementName IString, type_ IString, defaultValue IString)) {
	d._ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue = f
}

// SetParserFoundCDATA sets the handler for the ParserFoundCDATA delegate method.
//
// Sent by a parser object to its delegate when it encounters a CDATA block.
func (d *XMLParserDelegate) SetParserFoundCDATA(f func(parser IXMLParser, CDATABlock IData)) {
	d._ParserFoundCDATA = f
}

// SetParserFoundCharacters sets the handler for the ParserFoundCharacters delegate method.
//
// Sent by a parser object to provide its delegate with a string representing all or part of the characters of the current element.
func (d *XMLParserDelegate) SetParserFoundCharacters(f func(parser IXMLParser, string_ IString)) {
	d._ParserFoundCharacters = f
}

// SetParserFoundComment sets the handler for the ParserFoundComment delegate method.
//
// Sent by a parser object to its delegate when it encounters a comment in the XML.
func (d *XMLParserDelegate) SetParserFoundComment(f func(parser IXMLParser, comment IString)) {
	d._ParserFoundComment = f
}

// SetParserFoundElementDeclarationWithNameModel sets the handler for the ParserFoundElementDeclarationWithNameModel delegate method.
//
// Sent by a parser object to its delegate when it encounters a declaration of an element with a given model.
func (d *XMLParserDelegate) SetParserFoundElementDeclarationWithNameModel(f func(parser IXMLParser, elementName IString, model IString)) {
	d._ParserFoundElementDeclarationWithNameModel = f
}

// SetParserFoundExternalEntityDeclarationWithNamePublicIDSystemID sets the handler for the ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID delegate method.
//
// Sent by a parser object to its delegate when it encounters an external entity declaration.
func (d *XMLParserDelegate) SetParserFoundExternalEntityDeclarationWithNamePublicIDSystemID(f func(parser IXMLParser, name IString, publicID IString, systemID IString)) {
	d._ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID = f
}

// SetParserFoundIgnorableWhitespace sets the handler for the ParserFoundIgnorableWhitespace delegate method.
//
// Reported by a parser object to provide its delegate with a string representing all or part of the ignorable whitespace characters of the current element.
func (d *XMLParserDelegate) SetParserFoundIgnorableWhitespace(f func(parser IXMLParser, whitespaceString IString)) {
	d._ParserFoundIgnorableWhitespace = f
}

// SetParserFoundInternalEntityDeclarationWithNameValue sets the handler for the ParserFoundInternalEntityDeclarationWithNameValue delegate method.
//
// Sent by a parser object to the delegate when it encounters an internal entity declaration.
func (d *XMLParserDelegate) SetParserFoundInternalEntityDeclarationWithNameValue(f func(parser IXMLParser, name IString, value IString)) {
	d._ParserFoundInternalEntityDeclarationWithNameValue = f
}

// SetParserFoundNotationDeclarationWithNamePublicIDSystemID sets the handler for the ParserFoundNotationDeclarationWithNamePublicIDSystemID delegate method.
//
// Sent by a parser object to its delegate when it encounters a notation declaration.
func (d *XMLParserDelegate) SetParserFoundNotationDeclarationWithNamePublicIDSystemID(f func(parser IXMLParser, name IString, publicID IString, systemID IString)) {
	d._ParserFoundNotationDeclarationWithNamePublicIDSystemID = f
}

// SetParserFoundProcessingInstructionWithTargetData sets the handler for the ParserFoundProcessingInstructionWithTargetData delegate method.
//
// Sent by a parser object to its delegate when it encounters a processing instruction.
func (d *XMLParserDelegate) SetParserFoundProcessingInstructionWithTargetData(f func(parser IXMLParser, target IString, data IString)) {
	d._ParserFoundProcessingInstructionWithTargetData = f
}

// SetParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName sets the handler for the ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName delegate method.
//
// Sent by a parser object to its delegate when it encounters an unparsed entity declaration.
func (d *XMLParserDelegate) SetParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName(f func(parser IXMLParser, name IString, publicID IString, systemID IString, notationName IString)) {
	d._ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName = f
}

// SetParserParseErrorOccurred sets the handler for the ParserParseErrorOccurred delegate method.
//
// Sent by a parser object to its delegate when it encounters a fatal error.
func (d *XMLParserDelegate) SetParserParseErrorOccurred(f func(parser IXMLParser, parseError IError)) {
	d._ParserParseErrorOccurred = f
}

// SetParserResolveExternalEntityNameSystemID sets the handler for the ParserResolveExternalEntityNameSystemID delegate method.
//
// Sent by a parser object to its delegate when it encounters a given external entity with a specific system ID.
func (d *XMLParserDelegate) SetParserResolveExternalEntityNameSystemID(f func(parser IXMLParser, name IString, systemID IString) Data) {
	d._ParserResolveExternalEntityNameSystemID = f
}

// SetParserValidationErrorOccurred sets the handler for the ParserValidationErrorOccurred delegate method.
//
// Sent by a parser object to its delegate when it encounters a fatal validation error.   currently does not invoke this method and does not perform validation.
func (d *XMLParserDelegate) SetParserValidationErrorOccurred(f func(parser IXMLParser, validationError IError)) {
	d._ParserValidationErrorOccurred = f
}

// SetParserDidEndDocument sets the handler for the ParserDidEndDocument delegate method.
//
// Sent by the parser object to the delegate when it has successfully completed parsing.
func (d *XMLParserDelegate) SetParserDidEndDocument(f func(parser IXMLParser)) {
	d._ParserDidEndDocument = f
}

// SetParserDidStartDocument sets the handler for the ParserDidStartDocument delegate method.
//
// Sent by the parser object to the delegate when it begins parsing a document.
func (d *XMLParserDelegate) SetParserDidStartDocument(f func(parser IXMLParser)) {
	d._ParserDidStartDocument = f
}

// ParserDidEndElementNamespaceURIQualifiedName implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidEndElementNamespaceURIQualifiedName(parser IXMLParser, elementName IString, namespaceURI IString, qName IString) {
	if d._ParserDidEndElementNamespaceURIQualifiedName != nil {
		d._ParserDidEndElementNamespaceURIQualifiedName(parser, elementName, namespaceURI, qName)
	}
}

// HasParserDidEndElementNamespaceURIQualifiedName returns true if a handler for ParserDidEndElementNamespaceURIQualifiedName has been set.
func (d *XMLParserDelegate) HasParserDidEndElementNamespaceURIQualifiedName() bool {
	return d._ParserDidEndElementNamespaceURIQualifiedName != nil
}

// ParserDidEndMappingPrefix implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidEndMappingPrefix(parser IXMLParser, prefix IString) {
	if d._ParserDidEndMappingPrefix != nil {
		d._ParserDidEndMappingPrefix(parser, prefix)
	}
}

// HasParserDidEndMappingPrefix returns true if a handler for ParserDidEndMappingPrefix has been set.
func (d *XMLParserDelegate) HasParserDidEndMappingPrefix() bool {
	return d._ParserDidEndMappingPrefix != nil
}

// ParserDidStartElementNamespaceURIQualifiedNameAttributes implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidStartElementNamespaceURIQualifiedNameAttributes(parser IXMLParser, elementName IString, namespaceURI IString, qName IString, attributeDict IDictionary) {
	if d._ParserDidStartElementNamespaceURIQualifiedNameAttributes != nil {
		d._ParserDidStartElementNamespaceURIQualifiedNameAttributes(parser, elementName, namespaceURI, qName, attributeDict)
	}
}

// HasParserDidStartElementNamespaceURIQualifiedNameAttributes returns true if a handler for ParserDidStartElementNamespaceURIQualifiedNameAttributes has been set.
func (d *XMLParserDelegate) HasParserDidStartElementNamespaceURIQualifiedNameAttributes() bool {
	return d._ParserDidStartElementNamespaceURIQualifiedNameAttributes != nil
}

// ParserDidStartMappingPrefixToURI implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidStartMappingPrefixToURI(parser IXMLParser, prefix IString, namespaceURI IString) {
	if d._ParserDidStartMappingPrefixToURI != nil {
		d._ParserDidStartMappingPrefixToURI(parser, prefix, namespaceURI)
	}
}

// HasParserDidStartMappingPrefixToURI returns true if a handler for ParserDidStartMappingPrefixToURI has been set.
func (d *XMLParserDelegate) HasParserDidStartMappingPrefixToURI() bool {
	return d._ParserDidStartMappingPrefixToURI != nil
}

// ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue(parser IXMLParser, attributeName IString, elementName IString, type_ IString, defaultValue IString) {
	if d._ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue != nil {
		d._ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue(parser, attributeName, elementName, type_, defaultValue)
	}
}

// HasParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue returns true if a handler for ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue has been set.
func (d *XMLParserDelegate) HasParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue() bool {
	return d._ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue != nil
}

// ParserFoundCDATA implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundCDATA(parser IXMLParser, CDATABlock IData) {
	if d._ParserFoundCDATA != nil {
		d._ParserFoundCDATA(parser, CDATABlock)
	}
}

// HasParserFoundCDATA returns true if a handler for ParserFoundCDATA has been set.
func (d *XMLParserDelegate) HasParserFoundCDATA() bool {
	return d._ParserFoundCDATA != nil
}

// ParserFoundCharacters implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundCharacters(parser IXMLParser, string_ IString) {
	if d._ParserFoundCharacters != nil {
		d._ParserFoundCharacters(parser, string_)
	}
}

// HasParserFoundCharacters returns true if a handler for ParserFoundCharacters has been set.
func (d *XMLParserDelegate) HasParserFoundCharacters() bool {
	return d._ParserFoundCharacters != nil
}

// ParserFoundComment implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundComment(parser IXMLParser, comment IString) {
	if d._ParserFoundComment != nil {
		d._ParserFoundComment(parser, comment)
	}
}

// HasParserFoundComment returns true if a handler for ParserFoundComment has been set.
func (d *XMLParserDelegate) HasParserFoundComment() bool {
	return d._ParserFoundComment != nil
}

// ParserFoundElementDeclarationWithNameModel implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundElementDeclarationWithNameModel(parser IXMLParser, elementName IString, model IString) {
	if d._ParserFoundElementDeclarationWithNameModel != nil {
		d._ParserFoundElementDeclarationWithNameModel(parser, elementName, model)
	}
}

// HasParserFoundElementDeclarationWithNameModel returns true if a handler for ParserFoundElementDeclarationWithNameModel has been set.
func (d *XMLParserDelegate) HasParserFoundElementDeclarationWithNameModel() bool {
	return d._ParserFoundElementDeclarationWithNameModel != nil
}

// ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID(parser IXMLParser, name IString, publicID IString, systemID IString) {
	if d._ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID != nil {
		d._ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID(parser, name, publicID, systemID)
	}
}

// HasParserFoundExternalEntityDeclarationWithNamePublicIDSystemID returns true if a handler for ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID has been set.
func (d *XMLParserDelegate) HasParserFoundExternalEntityDeclarationWithNamePublicIDSystemID() bool {
	return d._ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID != nil
}

// ParserFoundIgnorableWhitespace implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundIgnorableWhitespace(parser IXMLParser, whitespaceString IString) {
	if d._ParserFoundIgnorableWhitespace != nil {
		d._ParserFoundIgnorableWhitespace(parser, whitespaceString)
	}
}

// HasParserFoundIgnorableWhitespace returns true if a handler for ParserFoundIgnorableWhitespace has been set.
func (d *XMLParserDelegate) HasParserFoundIgnorableWhitespace() bool {
	return d._ParserFoundIgnorableWhitespace != nil
}

// ParserFoundInternalEntityDeclarationWithNameValue implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundInternalEntityDeclarationWithNameValue(parser IXMLParser, name IString, value IString) {
	if d._ParserFoundInternalEntityDeclarationWithNameValue != nil {
		d._ParserFoundInternalEntityDeclarationWithNameValue(parser, name, value)
	}
}

// HasParserFoundInternalEntityDeclarationWithNameValue returns true if a handler for ParserFoundInternalEntityDeclarationWithNameValue has been set.
func (d *XMLParserDelegate) HasParserFoundInternalEntityDeclarationWithNameValue() bool {
	return d._ParserFoundInternalEntityDeclarationWithNameValue != nil
}

// ParserFoundNotationDeclarationWithNamePublicIDSystemID implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundNotationDeclarationWithNamePublicIDSystemID(parser IXMLParser, name IString, publicID IString, systemID IString) {
	if d._ParserFoundNotationDeclarationWithNamePublicIDSystemID != nil {
		d._ParserFoundNotationDeclarationWithNamePublicIDSystemID(parser, name, publicID, systemID)
	}
}

// HasParserFoundNotationDeclarationWithNamePublicIDSystemID returns true if a handler for ParserFoundNotationDeclarationWithNamePublicIDSystemID has been set.
func (d *XMLParserDelegate) HasParserFoundNotationDeclarationWithNamePublicIDSystemID() bool {
	return d._ParserFoundNotationDeclarationWithNamePublicIDSystemID != nil
}

// ParserFoundProcessingInstructionWithTargetData implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundProcessingInstructionWithTargetData(parser IXMLParser, target IString, data IString) {
	if d._ParserFoundProcessingInstructionWithTargetData != nil {
		d._ParserFoundProcessingInstructionWithTargetData(parser, target, data)
	}
}

// HasParserFoundProcessingInstructionWithTargetData returns true if a handler for ParserFoundProcessingInstructionWithTargetData has been set.
func (d *XMLParserDelegate) HasParserFoundProcessingInstructionWithTargetData() bool {
	return d._ParserFoundProcessingInstructionWithTargetData != nil
}

// ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName(parser IXMLParser, name IString, publicID IString, systemID IString, notationName IString) {
	if d._ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName != nil {
		d._ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName(parser, name, publicID, systemID, notationName)
	}
}

// HasParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName returns true if a handler for ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName has been set.
func (d *XMLParserDelegate) HasParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName() bool {
	return d._ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName != nil
}

// ParserParseErrorOccurred implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserParseErrorOccurred(parser IXMLParser, parseError IError) {
	if d._ParserParseErrorOccurred != nil {
		d._ParserParseErrorOccurred(parser, parseError)
	}
}

// HasParserParseErrorOccurred returns true if a handler for ParserParseErrorOccurred has been set.
func (d *XMLParserDelegate) HasParserParseErrorOccurred() bool {
	return d._ParserParseErrorOccurred != nil
}

// ParserResolveExternalEntityNameSystemID implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserResolveExternalEntityNameSystemID(parser IXMLParser, name IString, systemID IString) Data {
	if d._ParserResolveExternalEntityNameSystemID != nil {
		return d._ParserResolveExternalEntityNameSystemID(parser, name, systemID)
	}
	var zero Data
	return zero
}

// HasParserResolveExternalEntityNameSystemID returns true if a handler for ParserResolveExternalEntityNameSystemID has been set.
func (d *XMLParserDelegate) HasParserResolveExternalEntityNameSystemID() bool {
	return d._ParserResolveExternalEntityNameSystemID != nil
}

// ParserValidationErrorOccurred implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserValidationErrorOccurred(parser IXMLParser, validationError IError) {
	if d._ParserValidationErrorOccurred != nil {
		d._ParserValidationErrorOccurred(parser, validationError)
	}
}

// HasParserValidationErrorOccurred returns true if a handler for ParserValidationErrorOccurred has been set.
func (d *XMLParserDelegate) HasParserValidationErrorOccurred() bool {
	return d._ParserValidationErrorOccurred != nil
}

// ParserDidEndDocument implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidEndDocument(parser IXMLParser) {
	if d._ParserDidEndDocument != nil {
		d._ParserDidEndDocument(parser)
	}
}

// HasParserDidEndDocument returns true if a handler for ParserDidEndDocument has been set.
func (d *XMLParserDelegate) HasParserDidEndDocument() bool {
	return d._ParserDidEndDocument != nil
}

// ParserDidStartDocument implements the PXMLParserDelegate interface.
func (d *XMLParserDelegate) ParserDidStartDocument(parser IXMLParser) {
	if d._ParserDidStartDocument != nil {
		d._ParserDidStartDocument(parser)
	}
}

// HasParserDidStartDocument returns true if a handler for ParserDidStartDocument has been set.
func (d *XMLParserDelegate) HasParserDidStartDocument() bool {
	return d._ParserDidStartDocument != nil
}
