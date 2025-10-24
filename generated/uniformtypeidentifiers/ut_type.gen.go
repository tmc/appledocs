// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

package uniformtypeidentifiers

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UTType */

/* debug [class_header]: Header for UTType */
// The class instance for the [UTType] class.
var (
	UTTypeClass     _UTTypeClass
	UTTypeClassOnce sync.Once
)

func getUTTypeClass() _UTTypeClass {
	UTTypeClassOnce.Do(func() {
		UTTypeClass = _UTTypeClass{objc.GetClass("UTType")}
	})
	return UTTypeClass
}

type _UTTypeClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UTType */
// An interface definition for the [UTType] class.
type IUTType interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UTType */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Declared() bool
	Dynamic() bool
	PublicType() bool
	LocalizedDescription() objc.IObject       /* cross-framework: NSString */
	PreferredFilenameExtension() objc.IObject /* cross-framework: NSString */
	PreferredMIMEType() objc.IObject          /* cross-framework: NSString */
	ReferenceURL() objc.IObject               /* cross-framework: NSURL */
	Supertypes() unsafe.Pointer
	Tags() foundation.IDictionary
	Version() objc.IObject /* cross-framework: NSNumber */
	IsDeclared() bool
	SetIsDeclared(value bool)
	IsDynamic() bool
	SetIsDynamic(value bool)
	IsPublic() bool
	SetIsPublic(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UTType */
	// methods:
	ConformsToType(type_ IUTType) bool
	IsSubtypeOfType(type_ IUTType) bool
	IsSupertypeOfType(type_ IUTType) bool
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UTType */
// Alloc allocates a new instance without initialization.
func (uc _UTTypeClass) Alloc() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UTTypeClass) New() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UTType) Init() UTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UTType) Autorelease() UTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUTType creates a new UTType instance.
func NewUTType() UTType {
	return getUTTypeClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UTType */
// An object that represents a type of data to load, send, or receive.
//
// The object may represent files on disk, abstract data types with no on-disk representation, or entirely unrelated hierarchical classification systems, such as hardware. Each instance has a unique , and helpful properties, and . The object may provide additional information related to the type. For example, it may include a localized user-facing description, a reference URL to technical documentation about the type, or its version number. You can look up types by their conformance to get either a type or a list of types that are relevant to your use case. To define your own types in your app’s , see .

// An object that represents a type of data to load, send, or receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference
type UTType struct {
	objectivec.Object
}

// UTTypeFrom constructs a [UTType] from an unsafe.Pointer.
//
// An object that represents a type of data to load, send, or receive.
func UTTypeFrom(ptr unsafe.Pointer) UTType {
	return UTType{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UTType */

// Creates a type your app owns based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:)
func NewUTTypeExportedTypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:"), identifier)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeExportedTypeWithIdentifier */

// Creates a type your app owns based on an identifier and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:conformingTo:)
func NewUTTypeExportedTypeWithIdentifierConformingToType(identifier objc.IObject /* cross-framework: NSString */, parentType IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:conformingToType:"), identifier, parentType)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeExportedTypeWithIdentifierConformingToType */

// Creates a type your app uses, but doesn’t own, based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:)
func NewUTTypeImportedTypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:"), identifier)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeImportedTypeWithIdentifier */

// Creates a type your app uses, but doesn’t own, based on an identifier and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:conformingTo:)
func NewUTTypeImportedTypeWithIdentifierConformingToType(identifier objc.IObject /* cross-framework: NSString */, parentType IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:conformingToType:"), identifier, parentType)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeImportedTypeWithIdentifierConformingToType */

// Creates a type that represents the specified filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
func NewUTTypeWithFilenameExtension(filenameExtension objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:"), filenameExtension)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithFilenameExtension */

// Creates a type that represents the specified filename extension and conforms to an existing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
func NewUTTypeWithFilenameExtensionConformingToType(filenameExtension objc.IObject /* cross-framework: NSString */, supertype IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:conformingToType:"), filenameExtension, supertype)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithFilenameExtensionConformingToType */

// Creates a type based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(_:)
func NewUTTypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithIdentifier:"), identifier)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithIdentifier */

// Creates a type based on a MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:)
func NewUTTypeWithMIMEType(mimeType objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:"), mimeType)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithMIMEType */

// Creates a type based on a MIME type and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
func NewUTTypeWithMIMETypeConformingToType(mimeType objc.IObject /* cross-framework: NSString */, supertype IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:conformingToType:"), mimeType, supertype)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithMIMETypeConformingToType */

// Creates a type that represents the specified tag and tag class and which conforms to an existing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
func NewUTTypeWithTagTagClassConformingToType(tag objc.IObject /* cross-framework: NSString */, tagClass objc.IObject /* cross-framework: NSString */, supertype IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithTag:tagClass:conformingToType:"), tag, tagClass, supertype)
	return rv
} /* debug [class_init_methods/constructor]: NewUTTypeWithTagTagClassConformingToType */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UTType */

// Creates a type based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(_:)
func (uc _UTTypeClass) TypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithIdentifier:"), identifier)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithIdentifier) */

// Creates a type your app owns based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:)
func (uc _UTTypeClass) ExportedTypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("exportedTypeWithIdentifier:"), identifier)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ExportedTypeWithIdentifier) */

// Creates a type your app owns based on an identifier and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:conformingTo:)
func (uc _UTTypeClass) ExportedTypeWithIdentifierConformingToType(identifier objc.IObject /* cross-framework: NSString */, parentType IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("exportedTypeWithIdentifier:conformingToType:"), identifier, parentType)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ExportedTypeWithIdentifierConformingToType) */

// Creates a type that represents the specified filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
func (uc _UTTypeClass) TypeWithFilenameExtension(filenameExtension objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithFilenameExtension:"), filenameExtension)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithFilenameExtension) */

// Creates a type that represents the specified filename extension and conforms to an existing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
func (uc _UTTypeClass) TypeWithFilenameExtensionConformingToType(filenameExtension objc.IObject /* cross-framework: NSString */, supertype IUTType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithFilenameExtension:conformingToType:"), filenameExtension, supertype)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithFilenameExtensionConformingToType) */

// Creates a type your app uses, but doesn’t own, based on an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:)
func (uc _UTTypeClass) ImportedTypeWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("importedTypeWithIdentifier:"), identifier)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ImportedTypeWithIdentifier) */

// Creates a type your app uses, but doesn’t own, based on an identifier and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:conformingTo:)
func (uc _UTTypeClass) ImportedTypeWithIdentifierConformingToType(identifier objc.IObject /* cross-framework: NSString */, parentType IUTType) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("importedTypeWithIdentifier:conformingToType:"), identifier, parentType)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ImportedTypeWithIdentifierConformingToType) */

// Creates a type based on a MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:)
func (uc _UTTypeClass) TypeWithMIMEType(mimeType objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithMIMEType:"), mimeType)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithMIMEType) */

// Creates a type based on a MIME type and a supertype that it conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
func (uc _UTTypeClass) TypeWithMIMETypeConformingToType(mimeType objc.IObject /* cross-framework: NSString */, supertype IUTType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithMIMEType:conformingToType:"), mimeType, supertype)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithMIMETypeConformingToType) */

// Creates a type that represents the specified tag and tag class and which conforms to an existing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
func (uc _UTTypeClass) TypeWithTagTagClassConformingToType(tag objc.IObject /* cross-framework: NSString */, tagClass objc.IObject /* cross-framework: NSString */, supertype IUTType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithTag:tagClass:conformingToType:"), tag, tagClass, supertype)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypeWithTagTagClassConformingToType) */

// Returns an array of types from the provided tag and tag class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/types(tag:tagClass:conformingTo:)
func (uc _UTTypeClass) TypesWithTagTagClassConformingToType(tag objc.IObject /* cross-framework: NSString */, tagClass objc.IObject /* cross-framework: NSString */, supertype IUTType) []UTType {
	rv := objc.Send[[]UTType](objc.ID(uc.class), objc.Sel("typesWithTag:tagClass:conformingToType:"), tag, tagClass, supertype)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TypesWithTagTagClassConformingToType) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UTType */

// A type representing the @c SHCustomCatalog file format with the .shazamcatalog extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHCustomCatalogContentType
func (uc _UTTypeClass) SHCustomCatalogContentType() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("SHCustomCatalogContentType"))
	return rv
} /* debug [class_properties_class/property]: SHCustomCatalogContentType */

// A type representing the @c SHSignature file format with the .shazamsignature extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHSignatureContentType
func (uc _UTTypeClass) SHSignatureContentType() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("SHSignatureContentType"))
	return rv
} /* debug [class_properties_class/property]: SHSignatureContentType */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UTType */

// Returns a Boolean value that indicates whether a type conforms to the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/conforms(to:)
func (u_ UTType) ConformsToType(type_ IUTType) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("conformsToType:"), type_)
	return rv
} /* debug [instance_methods/method]: ConformsToType */

// Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSubtype(of:)
func (u_ UTType) IsSubtypeOfType(type_ IUTType) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isSubtypeOfType:"), type_)
	return rv
} /* debug [instance_methods/method]: IsSubtypeOfType */

// Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSupertype(of:)
func (u_ UTType) IsSupertypeOfType(type_ IUTType) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isSupertypeOfType:"), type_)
	return rv
} /* debug [instance_methods/method]: IsSupertypeOfType */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UTType */

// The string that represents the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/identifier
func (u_ UTType) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// A Boolean value that indicates whether the system declares the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDeclared
func (u_ UTType) Declared() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("declared"))
	return rv
} /* debug [instance_properties/getter]: declared */

// A Boolean value that indicates whether the system generates the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDynamic
func (u_ UTType) Dynamic() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("dynamic"))
	return rv
} /* debug [instance_properties/getter]: dynamic */

// A Boolean value that indicates whether the type is in the public domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isPublic
func (u_ UTType) PublicType() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("publicType"))
	return rv
} /* debug [instance_properties/getter]: publicType */

// A localized description of the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/localizedDescription
func (u_ UTType) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("localizedDescription"))
	return rv
} /* debug [instance_properties/getter]: localizedDescription */

// The preferred filename extension for the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredFilenameExtension
func (u_ UTType) PreferredFilenameExtension() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("preferredFilenameExtension"))
	return rv
} /* debug [instance_properties/getter]: preferredFilenameExtension */

// The preferred MIME type for the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredMIMEType
func (u_ UTType) PreferredMIMEType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("preferredMIMEType"))
	return rv
} /* debug [instance_properties/getter]: preferredMIMEType */

// The reference URL for the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/referenceURL
func (u_ UTType) ReferenceURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](u_.ID, objc.Sel("referenceURL"))
	return rv
} /* debug [instance_properties/getter]: referenceURL */

// The set of types the type directly or indirectly conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/supertypes
func (u_ UTType) Supertypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("supertypes"))
	return rv
} /* debug [instance_properties/getter]: supertypes */

// The tag specification dictionary of the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/tags
func (u_ UTType) Tags() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](u_.ID, objc.Sel("tags"))
	return rv
} /* debug [instance_properties/getter]: tags */

// The type’s version, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/version
func (u_ UTType) Version() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](u_.ID, objc.Sel("version"))
	return rv
} /* debug [instance_properties/getter]: version */

// A type representing the @c SHCustomCatalog file format with the .shazamcatalog extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHCustomCatalogContentType
func (u_ UTType) SHCustomCatalogContentType() IUTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("SHCustomCatalogContentType"))
	return rv
} /* debug [instance_properties/getter]: SHCustomCatalogContentType */

// A type representing the @c SHSignature file format with the .shazamsignature extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHSignatureContentType
func (u_ UTType) SHSignatureContentType() IUTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("SHSignatureContentType"))
	return rv
} /* debug [instance_properties/getter]: SHSignatureContentType */

// A Boolean value that indicates whether the system declares the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/isdeclared
func (u_ UTType) IsDeclared() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isDeclared"))
	return rv
} /* debug [instance_properties/getter]: isDeclared */

// A Boolean value that indicates whether the system declares the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/isdeclared
func (u_ UTType) SetIsDeclared(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsDeclared:"), value)
} /* debug [instance_properties/setter]: isDeclared */

// A Boolean value that indicates whether the system generates the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/isdynamic
func (u_ UTType) IsDynamic() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isDynamic"))
	return rv
} /* debug [instance_properties/getter]: isDynamic */

// A Boolean value that indicates whether the system generates the type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/isdynamic
func (u_ UTType) SetIsDynamic(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsDynamic:"), value)
} /* debug [instance_properties/setter]: isDynamic */

// A Boolean value that indicates whether the type is in the public domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/ispublic
func (u_ UTType) IsPublic() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPublic"))
	return rv
} /* debug [instance_properties/getter]: isPublic */

// A Boolean value that indicates whether the type is in the public domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttypereference/ispublic
func (u_ UTType) SetIsPublic(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPublic:"), value)
} /* debug [instance_properties/setter]: isPublic */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UTType */
