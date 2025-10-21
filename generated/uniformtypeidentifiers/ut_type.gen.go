// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

package uniformtypeidentifiers

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UTType] class.
type IUTType interface {
	objectivec.IObject
	ConformsToType(type_ unsafe.Pointer) bool
	IsSubtypeOfType(type_ unsafe.Pointer) bool
	IsSupertypeOfType(type_ unsafe.Pointer) bool
}

// An object that represents a type of data to load, send, or receive.
//
// The object may represent files on disk, abstract data types with no on-disk representation, or entirely unrelated hierarchical classification systems, such as hardware. Each instance has a unique , and helpful properties, and . The object may provide additional information related to the type. For example, it may include a localized user-facing description, a reference URL to technical documentation about the type, or its version number. You can look up types by their conformance to get either a type or a list of types that are relevant to your use case. To define your own types in your app’s , see .
//
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

// Alloc allocates a new instance without initialization.
func (uc _UTTypeClass) Alloc() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

// Creates a type that represents the specified filename extension.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
func NewUTTypeWithFilenameExtension(filenameExtension string) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:"), objc.String(filenameExtension))
	return rv
}

// Creates a type that represents the specified filename extension and conforms to an existing type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
func NewUTTypeWithFilenameExtensionConformingToType(filenameExtension string, supertype unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:conformingToType:"), objc.String(filenameExtension), supertype)
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:conformingTo:)
func NewUTTypeImportedTypeWithIdentifierConformingToType(identifier string, parentType unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return rv
}

// Creates a type based on a MIME type and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
func NewUTTypeWithMIMETypeConformingToType(mimeType string, supertype unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:conformingToType:"), objc.String(mimeType), supertype)
	return rv
}

// Creates a type that represents the specified tag and tag class and which conforms to an existing type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
func NewUTTypeWithTagTagClassConformingToType(tag string, tagClass string, supertype unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return rv
}

// Creates a type your app owns based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:)
func NewUTTypeExportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type your app owns based on an identifier and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:conformingTo:)
func NewUTTypeExportedTypeWithIdentifierConformingToType(identifier string, parentType unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:)
func NewUTTypeImportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type based on a MIME type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:)
func NewUTTypeWithMIMEType(mimeType string) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:"), objc.String(mimeType))
	return rv
}

// Creates a type based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(_:)
func NewUTTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[UTType](objc.ID(getUTTypeClass().class), objc.Sel("typeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(_:)
func (uc _UTTypeClass) TypeWithIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type your app owns based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:)
func (uc _UTTypeClass) ExportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("exportedTypeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type your app owns based on an identifier and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:conformingTo:)
func (uc _UTTypeClass) ExportedTypeWithIdentifierConformingToType(identifier string, parentType unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("exportedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return rv
}

// Creates a type that represents the specified filename extension.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
func (uc _UTTypeClass) TypeWithFilenameExtension(filenameExtension string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithFilenameExtension:"), objc.String(filenameExtension))
	return rv
}

// Creates a type that represents the specified filename extension and conforms to an existing type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
func (uc _UTTypeClass) TypeWithFilenameExtensionConformingToType(filenameExtension string, supertype unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithFilenameExtension:conformingToType:"), objc.String(filenameExtension), supertype)
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:)
func (uc _UTTypeClass) ImportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("importedTypeWithIdentifier:"), objc.String(identifier))
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:conformingTo:)
func (uc _UTTypeClass) ImportedTypeWithIdentifierConformingToType(identifier string, parentType unsafe.Pointer) UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("importedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return rv
}

// Creates a type based on a MIME type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:)
func (uc _UTTypeClass) TypeWithMIMEType(mimeType string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithMIMEType:"), objc.String(mimeType))
	return rv
}

// Creates a type based on a MIME type and a supertype that it conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
func (uc _UTTypeClass) TypeWithMIMETypeConformingToType(mimeType string, supertype unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithMIMEType:conformingToType:"), objc.String(mimeType), supertype)
	return rv
}

// Creates a type that represents the specified tag and tag class and which conforms to an existing type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
func (uc _UTTypeClass) TypeWithTagTagClassConformingToType(tag string, tagClass string, supertype unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("typeWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return rv
}

// Returns an array of types from the provided tag and tag class.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/types(tag:tagClass:conformingTo:)
func (uc _UTTypeClass) TypesWithTagTagClassConformingToType(tag string, tagClass string, supertype unsafe.Pointer) []UTType {
	rv := objc.Send[[]UTType](objc.ID(uc.class), objc.Sel("typesWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return rv
}

// A type representing the @c SHCustomCatalog file format with the .shazamcatalog extension
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHCustomCatalogContentType
func (uc _UTTypeClass) SHCustomCatalogContentType() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("SHCustomCatalogContentType"))
	return rv
}

// A type representing the @c SHSignature file format with the .shazamsignature extension
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHSignatureContentType
func (uc _UTTypeClass) SHSignatureContentType() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("SHSignatureContentType"))
	return rv
}

// Returns a Boolean value that indicates whether a type conforms to the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/conforms(to:)
func (u_ UTType) ConformsToType(type_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("conformsToType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSubtype(of:)
func (u_ UTType) IsSubtypeOfType(type_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isSubtypeOfType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSupertype(of:)
func (u_ UTType) IsSupertypeOfType(type_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isSupertypeOfType:"), type_)
	return rv
}

// A type representing the @c SHCustomCatalog file format with the .shazamcatalog extension
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHCustomCatalogContentType
func (u_ UTType) SHCustomCatalogContentType() UTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("SHCustomCatalogContentType"))
	return rv
}

// A type representing the @c SHSignature file format with the .shazamsignature extension
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHSignatureContentType
func (u_ UTType) SHSignatureContentType() UTType {
	rv := objc.Send[UTType](u_.ID, objc.Sel("SHSignatureContentType"))
	return rv
}

// The string that represents the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/identifier
func (u_ UTType) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether the system declares the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDeclared
func (u_ UTType) Declared() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("declared"))
	return rv
}

// A Boolean value that indicates whether the system generates the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDynamic
func (u_ UTType) Dynamic() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("dynamic"))
	return rv
}

// A Boolean value that indicates whether the type is in the public domain.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isPublic
func (u_ UTType) PublicType() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("publicType"))
	return rv
}

// A localized description of the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/localizedDescription
func (u_ UTType) LocalizedDescription() string {
	rv := objc.Send[string](u_.ID, objc.Sel("localizedDescription"))
	return rv
}

// The preferred filename extension for the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredFilenameExtension
func (u_ UTType) PreferredFilenameExtension() string {
	rv := objc.Send[string](u_.ID, objc.Sel("preferredFilenameExtension"))
	return rv
}

// The preferred MIME type for the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredMIMEType
func (u_ UTType) PreferredMIMEType() string {
	rv := objc.Send[string](u_.ID, objc.Sel("preferredMIMEType"))
	return rv
}

// The reference URL for the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/referenceURL
func (u_ UTType) ReferenceURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("referenceURL"))
	return rv
}

// The set of types the type directly or indirectly conforms to.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/supertypes
func (u_ UTType) Supertypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("supertypes"))
	return rv
}

// The tag specification dictionary of the type.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/tags
func (u_ UTType) Tags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("tags"))
	return rv
}

// The type’s version, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/version
func (u_ UTType) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("version"))
	return rv
}
