// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation
import (
	"unsafe"
)


// C struct types
// CFAllocatorContext - A structure that defines the context or operating environment for an allocator (CFAllocator) object. Every Core Foundation allocator object must have a context defined for it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorContext
type CFAllocatorContext struct {
	Allocate AllocatorAllocateCallBack // A prototype for a function callback that allocates memory of a requested size. In implementing this function, allocate a block of memory of at least   bytes and return a pointer to the start of the block. The   argument is a bitfield that you should currently not use (that is, assign 0). The   parameter should always be greater than 0. If it is not, or if problems in allocation occur, return  . This function pointer may not be assigned  .
	CopyDescription AllocatorCopyDescriptionCallBack // A prototype for a function callback that provides a description of the data pointed to by the   field. In implementing this function, return a reference to a CFString object that describes your allocator, particularly some characteristics of your program-defined data. You may set this function pointer to  , in which case Core Foundation will provide a rudimentary description.
	Deallocate AllocatorDeallocateCallBack // A prototype for a function callback that deallocates a given block of memory. In implementing this function, make the block of memory pointed to by   available for subsequent reuse by the allocator but unavailable for continued use by the program. The   parameter cannot be   and if the   parameter is not a block of memory that has been previously allocated by the allocator, the results are undefined; abnormal program termination can occur. You can set this callback to  , in which case the   function has no effect.
	Info unsafe.Pointer // An untyped pointer to program-defined data. Allocate memory for this data and assign a pointer to it. This data is often control information for the allocator. You may assign  .
	PreferredSize AllocatorPreferredSizeCallBack // A prototype for a function callback that determines whether there is enough free memory to satisfy a request. In implementing this function, return the actual size the allocator is likely to allocate given a request for a block of memory of size  . The   argument is a bitfield that you should currently not use.
	Reallocate AllocatorReallocateCallBack // A prototype for a function callback that reallocates memory of a requested size for an existing block of memory.
	Release AllocatorReleaseCallBack // A prototype for a function callback that releases the data pointed to by the   field. In implementing this function, release (or free) the data you have defined for the allocator context. You may set this function pointer to  , but doing so might result in memory leaks.
	Retain AllocatorRetainCallBack // A prototype for a function callback that retains the data pointed to by the   field. In implementing this function, retain the data you have defined for the allocator context in this field. (This might make sense only if the data is a Core Foundation object.) You may set this function pointer to  .
	Version Index // An integer of type  . Assign the version number of the allocator. Currently the only valid value is 0.
}

// CFArrayCallBacks - Structure containing the callbacks of a CFArray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCallBacks
type CFArrayCallBacks struct {
	CopyDescription ArrayCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If  , the collection will create a simple description of each value. See   for a description of this callback.
	Equal ArrayEqualCallBack // The callback used to compare values in the array for equality for some operations. If  , the collection will use pointer equality to compare values in the collection. See   for a description of this callback.
	Release ArrayReleaseCallBack // The callback used to release values as they are removed from the collection. If  , values are not released. See   for a description of this callback.
	Retain ArrayRetainCallBack // The callback used to retain each value as they are added to the collection. If  , values are not retained. See   for a description of this callback.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
}

// CFBagCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values of a CFBag object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCallBacks
type CFBagCallBacks struct {
	CopyDescription BagCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If  , the collection will create a simple description of each value. See   for a description of this callback.
	Equal BagEqualCallBack // The callback used to compare values in the collection for equality for some operations. If  , the collection will use pointer equality to compare values in the collection. See   for a description of this callback.
	Hash BagHashCallBack // The callback used to compute a hash code for values in a collection. If  , the collection computes a hash code by converting the pointer value to an integer. See   for a description of this callback.
	Release BagReleaseCallBack // The callback used to release values as they are removed from the collection. If  , values are not released. See   for a description of this callback.
	Retain BagRetainCallBack // The callback used to retain each value as they are added to the collection. If  , values are not retained. See   for a descriptions of this function’s parameters.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
}

// CFBinaryHeapCallBacks - Structure containing the callbacks for values for a 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCallBacks
type CFBinaryHeapCallBacks struct {
	Compare unsafe.Pointer // The callback used to compare values in the binary heap in some operations. This field cannot be  .
	CopyDescription unsafe.Pointer // Callback function used to get a description of a value in a binary heap.
	Release unsafe.Pointer // Callback function used to release a value before it is removed from a binary heap.
	Retain unsafe.Pointer // Callback function used to retain a value being added to a binary heap.
	Version Index // The version number of the structure type being passed in as a parameter to the   creation functions. This structure is version  .
}

// CFBinaryHeapCompareContext - Not used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCompareContext
type CFBinaryHeapCompareContext struct {
	CopyDescription unsafe.Pointer
	Info unsafe.Pointer
	Release unsafe.Pointer
	Retain unsafe.Pointer
	Version Index
}

// CFDictionaryKeyCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the keys in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryKeyCallBacks
type CFDictionaryKeyCallBacks struct {
	CopyDescription DictionaryCopyDescriptionCallBack // The callback used to create a descriptive string representation of each key in the dictionary. If  , the collection will create a simple description of each key. See   for a description of this callback.
	Equal DictionaryEqualCallBack // The callback used to compare keys in the dictionary for equality. If  , the collection will use pointer equality to compare keys in the collection. See   for a description of this callback.
	Hash DictionaryHashCallBack // The callback used to compute a hash code for keys as they are used to access, add, or remove values in the dictionary. If  , the collection computes a hash code by converting the pointer value to an integer. See   for a description of this callback.
	Release DictionaryReleaseCallBack // The callback used to release keys as they are removed from the dictionary. If  , keys are not released. See   for a description of this callback.
	Retain DictionaryRetainCallBack // The callback used to retain each key as they are added to the collection. This callback returns the value to use as the key in the dictionary, which is usually the value parameter passed to this callback, but may be a different value if a different value should be used as the key. If  , keys are not retained. See   for a descriptions of this function’s parameters.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
}

// CFDictionaryValueCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryValueCallBacks
type CFDictionaryValueCallBacks struct {
	CopyDescription DictionaryCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the dictionary. If  , the collection will create a simple description of each value. See   for a description of this callback.
	Equal DictionaryEqualCallBack // The callback used to compare values in the dictionary for equality. If  , the collection will use pointer equality to compare values in the collection. See   for a description of this callback.
	Release DictionaryReleaseCallBack // The callback used to release values as they are removed from the dictionary. If  , values are not released. See   for a description of this callback.
	Retain DictionaryRetainCallBack // The callback used to retain each value as they are added to the collection. This callback returns the value to use as the value in the dictionary, which is usually the value parameter passed to this callback, but may be a different value if a different value should be used as the value. If  , values are not retained. See   for a descriptions of this function’s parameters.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
}

// CFFileDescriptorContext - Defines a structure for the context of a CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorContext
type CFFileDescriptorContext struct {
	CopyDescription unsafe.Pointer // The callback used to create a descriptive string representation of the CFFileDescriptor.
	Info unsafe.Pointer
	Release unsafe.Pointer // The release callback used by the CFFileDescriptor.
	Retain unsafe.Pointer // The retain callback used by the CFFileDescriptor.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is 0.
}

// CFGregorianDate - Structure used to represent a point in time using the Gregorian calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDate
type CFGregorianDate struct {
	Day unsafe.Pointer
	Hour unsafe.Pointer
	Minute unsafe.Pointer
	Month unsafe.Pointer
	Second float64
	Year unsafe.Pointer
}

// CFGregorianUnits - Structure used to represent a time interval in Gregorian units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnits
type CFGregorianUnits struct {
	Days unsafe.Pointer
	Hours unsafe.Pointer
	Minutes unsafe.Pointer
	Months unsafe.Pointer
	Seconds float64
	Years unsafe.Pointer
}

// CFMachPortContext - A structure that contains program-defined data and callbacks with which you can configure a CFMachPort object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortContext
type CFMachPortContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the CFMachPort object at creation time. This pointer is passed to all the callbacks defined in the context.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be  .
}

// CFMessagePortContext - A structure that contains program-defined data and callbacks with which you can configure a CFMessagePort object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortContext
type CFMessagePortContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the message port at creation time. This pointer is passed to all the callbacks defined in the context.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be  .
}

// CFRange - A structure representing a range of sequential items in a container, such as characters in a buffer or elements in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRange
type CFRange struct {
	Length Index // An integer representing the number of items in the range. For type compatibility with the rest of the system,   is the maximum value you should use for length.
	Location Index // An integer representing the starting location of the range. For type compatibility with the rest of the system,   is the maximum value you should use for location.
}

// CFRunLoopObserverContext - A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopObserver object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverContext
type CFRunLoopObserverContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the run loop observer at creation time. This pointer is passed to all the callbacks defined in the context.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be  .
}

// CFRunLoopSourceContext - A structure that contains program-defined data and callbacks with which you can configure a version 0 CFRunLoopSource’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceContext
type CFRunLoopSourceContext struct {
	Cancel unsafe.Pointer
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Equal unsafe.Pointer // An equality test callback for your program-defined   pointer. Can be  .
	Hash unsafe.Pointer // A hash calculation callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the CFRunLoopSource at creation time. This pointer is passed to all the callbacks defined in the context.
	Perform unsafe.Pointer // A perform callback for the run loop source. This callback is called when the source has fired.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Schedule unsafe.Pointer // A scheduling callback for the run loop source. This callback is called when the source is added to a run loop mode. Can be  .
	Version Index // Version number of the structure. Must be 0.
}

// CFRunLoopSourceContext1 - A structure that contains program-defined data and callbacks with which you can configure a version 1 CFRunLoopSource’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceContext1
type CFRunLoopSourceContext1 struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Equal unsafe.Pointer // An equality test callback for your program-defined   pointer. Can be  .
	GetPort unsafe.Pointer // A callback to retrieve the native Mach port represented by the source. This callback is called when the source is either added to or removed from a run loop mode.
	Hash unsafe.Pointer // A hash calculation callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the run loop source at creation time. This pointer is passed to all the callbacks defined in the context.
	Perform unsafe.Pointer // A perform callback for the run loop source. This callback is called when the source has fired.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be 1.
}

// CFRunLoopTimerContext - A structure that contains program-defined data and callbacks with which you can configure a CFRunLoopTimer’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerContext
type CFRunLoopTimerContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the run loop timer at creation time. This pointer is passed to all the callbacks defined in the context.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be 0.
}

// CFSetCallBacks - This structure contains the callbacks used to retain, release, describe, and compare the values of a CFSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCallBacks
type CFSetCallBacks struct {
	CopyDescription SetCopyDescriptionCallBack // The callback used to create a descriptive string representation of each value in the collection. If  , the collection will create a simple description of each value. See   for a description of this callback.
	Equal SetEqualCallBack // The callback used to compare values in the collection for equality for some operations. If  , the collection will use pointer equality to compare values in the collection. See   for a description of this callback.
	Hash SetHashCallBack // The callback used to compute a hash code for values in a collection. If  , the collection computes a hash code by converting the pointer value to an integer. See   for a description of this callback.
	Release SetReleaseCallBack // The callback used to release values as they are removed from the collection. If  , values are not released. See   for a description of this callback.
	Retain SetRetainCallBack // The callback used to retain each value as they are added to the collection. If  , values are not retained. See   for a descriptions of this function’s parameters.
	Version Index // The version number of this structure. If not one of the defined version numbers for this opaque type, the behavior is undefined. The current version of this structure is  .
}

// CFSocketContext - A structure that contains program-defined data and callbacks with which you can configure a CFSocket object’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketContext
type CFSocketContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the CFSocket object at creation time. This pointer is passed to all the callbacks defined in the context.
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be  .
}

// CFSocketSignature - A structure that fully specifies the communication protocol and connection address of a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSignature
type CFSocketSignature struct {
	Address DataRef // A CFData object holding the contents of a   appropriate for the given protocol family (  or  , for example), identifying the address of the socket.
	Protocol unsafe.Pointer // The protocol type of the socket.
	ProtocolFamily unsafe.Pointer // The protocol family of the socket.
	SocketType unsafe.Pointer // The socket type of the socket.
}

// CFStreamClientContext - A structure that contains program-defined data and callbacks with which you can configure a stream’s client behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamClientContext
type CFStreamClientContext struct {
	CopyDescription unsafe.Pointer // A copy description callback for your program-defined   pointer. Can be  .
	Info unsafe.Pointer // An arbitrary pointer to program-defined data, which can be associated with the client. This pointer is passed to the callbacks defined in the context and to the client callback function  .
	Release unsafe.Pointer // A release callback for your program-defined   pointer. Can be  .
	Retain unsafe.Pointer // A retain callback for your program-defined   pointer. Can be  .
	Version Index // Version number of the structure. Must be  .
}

// CFStreamError - The structure returned by 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamError
type CFStreamError struct {
	Domain Index // The error domain that should be used to interpret the error. See   for possible values.
	Error unsafe.Pointer // The error code.
}

// CFStringInlineBuffer - Defines the buffer and related fields used for in-line buffer access of characters in CFString objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringInlineBuffer
type CFStringInlineBuffer struct {
	Buffer unsafe.Pointer
	BufferedRangeEnd Index
	BufferedRangeStart Index
	DirectCStringBuffer unsafe.Pointer
	DirectUniCharBuffer unsafe.Pointer
	RangeToBuffer unsafe.Pointer
	TheString StringRef
}

// CFSwappedFloat32 - Structure holding a 32-bit float value in a platform-independentbyte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSwappedFloat32
type CFSwappedFloat32 struct {
	V uint32 // A 32-bit float value stored with a platform-independentbyte order.
}

// CFSwappedFloat64 - Structure holding a 64-bit float value in a platform-independentbyte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSwappedFloat64
type CFSwappedFloat64 struct {
	V uint64 // A 64-bit float value stored with a platform-independentbyte order.
}

// CFTreeContext - Structure containing program-defined data and callbacks for a CFTree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeContext
type CFTreeContext struct {
	CopyDescription TreeCopyDescriptionCallBack // The callback used to provide a description of the   field.
	Info unsafe.Pointer // A C pointer to a program-defined block of data, referred to as the information pointer.
	Release TreeReleaseCallBack // The callback used to release a previously retained   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. This value may be  .
	Retain TreeRetainCallBack // The callback used to retain the   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. This value may be  .
	Version Index // The version number of the structure type being passed in as a parameter to a CFTree creation function. This structure is version  .
}

// CFUUIDBytes - A 128-bit struct that represents a UUID as raw bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDBytes
type CFUUIDBytes struct {
	Byte0 unsafe.Pointer // The first byte.
	Byte1 unsafe.Pointer // The second byte.
	Byte10 unsafe.Pointer // The eleventh byte.
	Byte11 unsafe.Pointer // The twelfth byte.
	Byte12 unsafe.Pointer // The thirteenth byte.
	Byte13 unsafe.Pointer // The fourteenth byte.
	Byte14 unsafe.Pointer // The fifteenth byte.
	Byte15 unsafe.Pointer // The sixteenth byte.
	Byte2 unsafe.Pointer // The third byte.
	Byte3 unsafe.Pointer // The fourth byte.
	Byte4 unsafe.Pointer // The fifth byte.
	Byte5 unsafe.Pointer // The sixth byte.
	Byte6 unsafe.Pointer // The seventh byte.
	Byte7 unsafe.Pointer // The eighth byte.
	Byte8 unsafe.Pointer // The ninth byte.
	Byte9 unsafe.Pointer // The tenth byte.
}

// CFXMLAttributeDeclarationInfo - Contains information about an element attribute definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLAttributeDeclarationInfo
type CFXMLAttributeDeclarationInfo struct {
	AttributeName StringRef // The name of the attribute.
	DefaultString StringRef // The attribute’s default value.
	TypeString StringRef // Describes the declaration of a single attribute.
}

// CFXMLAttributeListDeclarationInfo - Contains a list of the attributes associated with an element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLAttributeListDeclarationInfo
type CFXMLAttributeListDeclarationInfo struct {
	Attributes unsafe.Pointer // A C array of attributes.
	NumberOfAttributes Index // The number of attributes in the array.
}

// CFXMLDocumentInfo - Contains the source URL and text encoding information for the XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLDocumentInfo
type CFXMLDocumentInfo struct {
	Encoding StringEncoding // The text encoding of the XML document.
	SourceURL URLRef // The source URL of the XML document.
}

// CFXMLDocumentTypeInfo - Contains the external ID of the DTD.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLDocumentTypeInfo
type CFXMLDocumentTypeInfo struct {
	ExternalID unsafe.Pointer // The external ID of the DTD.
}

// CFXMLElementInfo - Contains a list of element attributes packaged as CFDictionary key/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLElementInfo
type CFXMLElementInfo struct {
	AttributeOrder ArrayRef // An array specifying the order in which the attributes appeared in the XML document.
	Attributes DictionaryRef // The dictionary of attribute values.
	IsEmpty unsafe.Pointer // A flag indicating whether the element was expressed in closed form.
}

// CFXMLElementTypeDeclarationInfo - Contains a description of the element type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLElementTypeDeclarationInfo
type CFXMLElementTypeDeclarationInfo struct {
	ContentDescription StringRef // A textual description of the element type.
}

// CFXMLEntityInfo - Contains information describing an XML entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityInfo
type CFXMLEntityInfo struct {
	EntityID unsafe.Pointer //  will be   if   is internal.
	EntityType unsafe.Pointer // The entity type code.
	NotationName StringRef //  if   is parsed.
	ReplacementText StringRef //  if   is external or unparsed, otherwise the text that the entity should be replaced with.
}

// CFXMLEntityReferenceInfo - Contains information describing an XML entity reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityReferenceInfo
type CFXMLEntityReferenceInfo struct {
	EntityType unsafe.Pointer // The entity type code.
}

// CFXMLExternalID - Contains the system and public IDs for an external entity reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLExternalID
type CFXMLExternalID struct {
	PublicID StringRef // The publicID string.
	SystemID URLRef // The systemID URL.
}

// CFXMLNotationInfo - Contains the external ID of the notation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNotationInfo
type CFXMLNotationInfo struct {
	ExternalID unsafe.Pointer // The external ID of the notation.
}

// CFXMLParserCallBacks - Contains version information and function pointers to callbacks needed when parsing XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCallBacks
type CFXMLParserCallBacks struct {
	AddChild XMLParserAddChildCallBack // Called when a child is added.
	CreateXMLStructure XMLParserCreateXMLStructureCallBack // Called when an XML structure is created.
	EndXMLStructure XMLParserEndXMLStructureCallBack // Called when an XML structure has ended.
	HandleError XMLParserHandleErrorCallBack // Called when a parse error needs to be handled.
	ResolveExternalEntity XMLParserResolveExternalEntityCallBack // Called when an external entity needs to be resolved.
	Version Index // Version number. Must be  .
}

// CFXMLParserContext - Contains version information and function pointers to callbacks used when handling a program-defined context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserContext
type CFXMLParserContext struct {
	CopyDescription XMLParserCopyDescriptionCallBack // A copy description callback for your program-defined context data. Optional.
	Info unsafe.Pointer // An arbitrary program-defined value passed to all the callbacks in this structure and in the   structure.
	Release XMLParserReleaseCallBack // A release callback for your program-defined context data. Optional.
	Retain XMLParserRetainCallBack // A retain callback for your program-defined context data. Optional.
	Version Index // Version number of this structure. Must be 0.
}

// CFXMLProcessingInstructionInfo - Contains the text of the processing instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLProcessingInstructionInfo
type CFXMLProcessingInstructionInfo struct {
	DataString StringRef // The text of the processing instruction.
}

// CGAffineTransform
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
type CGAffineTransform struct {
	A float64
	B float64
	C float64
	D float64
	Tx float64
	Ty float64
}

// CGAffineTransformComponents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransformComponents
type CGAffineTransformComponents struct {
	HorizontalShear float64
	Rotation float64
	Scale CGSize
	Translation CGVector
}

// CGPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
type CGPoint struct {
	X float64
	Y float64
}

// CGRect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect
type CGRect struct {
	Origin CGPoint
	Size CGSize
}

// CGSize - A structure that contains width and height values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGSize
type CGSize struct {
	Height float64 // A height value.
	Width float64 // A width value.
}

// CGVector - A structure that contains a two-dimensional vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGVector
type CGVector struct {
	Dx float64 // The x component of the vector.
	Dy float64 // The y component of the vector.
}

// IUnknownVTbl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/IUnknownVTbl
type IUnknownVTbl struct {
	AddRef unsafe.Pointer
	QueryInterface unsafe.Pointer
	Release unsafe.Pointer
}





