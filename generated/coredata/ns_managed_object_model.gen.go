// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectModel] class.
var (
	ManagedObjectModelClass     _ManagedObjectModelClass
	ManagedObjectModelClassOnce sync.Once
)

func getManagedObjectModelClass() _ManagedObjectModelClass {
	ManagedObjectModelClassOnce.Do(func() {
		ManagedObjectModelClass = _ManagedObjectModelClass{objc.GetClass("NSManagedObjectModel")}
	})
	return ManagedObjectModelClass
}

type _ManagedObjectModelClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectModel] class.
type IManagedObjectModel interface {
	objectivec.IObject
	FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables unsafe.Pointer) FetchRequest
	FetchRequestTemplateForName(name string) FetchRequest
	IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool
	SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string)
	Configurations() string
	SetConfigurations(value string)
	Entities() NSEntityDescription
	SetEntities(value IEntityDescription)
	EntitiesByName() NSEntityDescription
	SetEntitiesByName(value IEntityDescription)
	EntityVersionHashesByName() foundation.Data
	SetEntityVersionHashesByName(value foundation.IData)
	FetchRequestTemplatesByName() unsafe.Pointer
	SetFetchRequestTemplatesByName(value unsafe.Pointer)
	LocalizationDictionary() string
	SetLocalizationDictionary(value string)
	VersionChecksum() string
	SetVersionChecksum(value string)
	VersionIdentifiers() unsafe.Pointer
	SetVersionIdentifiers(value unsafe.Pointer)
}

// A programmatic representation of the file describing your objects.
//
// The model contains one or more objects representing the entities in the schema. Each object has property description objects (instances of subclasses of ) that represent the properties (or fields) of the entity in the schema. The Core Data framework uses this description in several ways: Constraining UI creation in Interface Builder Validating attribute and relationship values at runtime Mapping between your managed objects and a database or file-based schema for object persistence A managed object model maintains a mapping between each of its entity objects and a corresponding managed object class for use with the persistent storage mechanisms in the Core Data framework. You can determine the entity for a particular managed object with the method. You typically create managed object models using the data modeling tool in Xcode, but it’s possible to build a model programmatically if needed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel
type ManagedObjectModel struct {
	objectivec.Object
}

// ManagedObjectModelFrom constructs a [ManagedObjectModel] from an unsafe.Pointer.
//
// A programmatic representation of the file describing your objects.
func ManagedObjectModelFrom(ptr unsafe.Pointer) ManagedObjectModel {
	return ManagedObjectModel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectModelClass) Alloc() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectModelClass) New() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObjectModel) Init() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObjectModel) Autorelease() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObjectModel creates a new ManagedObjectModel instance.
func NewManagedObjectModel() ManagedObjectModel {
	return getManagedObjectModelClass().New()
}


// Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestFromTemplate(withName:substitutionVariables:)
func (m_ ManagedObjectModel) FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables unsafe.Pointer) FetchRequest {
	rv := objc.Send[FetchRequest](m_.ID, objc.Sel("fetchRequestFromTemplateWithName:substitutionVariables:"), objc.String(name), variables)
	return rv
}

// Returns the fetch request with a specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestTemplate(forName:)
func (m_ ManagedObjectModel) FetchRequestTemplateForName(name string) FetchRequest {
	rv := objc.Send[FetchRequest](m_.ID, objc.Sel("fetchRequestTemplateForName:"), objc.String(name))
	return rv
}

// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/isConfiguration(withName:compatibleWithStoreMetadata:)
func (m_ ManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), objc.String(configuration), metadata)
	return rv
}

// Associates the specified fetch request with the receiver using the given name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setFetchRequestTemplate(_:forName:)
func (m_ ManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplate:forName:"), fetchRequestTemplate, objc.String(name))
}

// All the available configuration names of the model.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/configurations
func (m_ ManagedObjectModel) Configurations() string {
	rv := objc.Send[string](m_.ID, objc.Sel("configurations"))
	return rv
}


// SetConfigurations sets the value of the configurations property.
// All the available configuration names of the model.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/configurations
func (m_ ManagedObjectModel) SetConfigurations(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfigurations:"), objc.String(value))
}

// The entities in the model.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entities
func (m_ ManagedObjectModel) Entities() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](m_.ID, objc.Sel("entities"))
	return rv
}


// SetEntities sets the value of the entities property.
// The entities in the model.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entities
func (m_ ManagedObjectModel) SetEntities(value IEntityDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntities:"), value)
}

// The entities of the model, keyed by name.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entitiesbyname
func (m_ ManagedObjectModel) EntitiesByName() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](m_.ID, objc.Sel("entitiesByName"))
	return rv
}


// SetEntitiesByName sets the value of the entitiesByName property.
// The entities of the model, keyed by name.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entitiesbyname
func (m_ ManagedObjectModel) SetEntitiesByName(value IEntityDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntitiesByName:"), value)
}

// The dictionary of the model’s entity names and their corresponding version hashes.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entityversionhashesbyname
func (m_ ManagedObjectModel) EntityVersionHashesByName() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("entityVersionHashesByName"))
	return rv
}


// SetEntityVersionHashesByName sets the value of the entityVersionHashesByName property.
// The dictionary of the model’s entity names and their corresponding version hashes.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entityversionhashesbyname
func (m_ ManagedObjectModel) SetEntityVersionHashesByName(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntityVersionHashesByName:"), value)
}

// A dictionary of the receiver’s fetch request templates, keyed by name.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/fetchrequesttemplatesbyname
func (m_ ManagedObjectModel) FetchRequestTemplatesByName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fetchRequestTemplatesByName"))
	return rv
}


// SetFetchRequestTemplatesByName sets the value of the fetchRequestTemplatesByName property.
// A dictionary of the receiver’s fetch request templates, keyed by name.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/fetchrequesttemplatesbyname
func (m_ ManagedObjectModel) SetFetchRequestTemplatesByName(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplatesByName:"), value)
}

// The localization dictionary of the model.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/localizationdictionary
func (m_ ManagedObjectModel) LocalizationDictionary() string {
	rv := objc.Send[string](m_.ID, objc.Sel("localizationDictionary"))
	return rv
}


// SetLocalizationDictionary sets the value of the localizationDictionary property.
// The localization dictionary of the model.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/localizationdictionary
func (m_ ManagedObjectModel) SetLocalizationDictionary(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalizationDictionary:"), objc.String(value))
}

// The Base64-encoded 128-bit model version hash.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionchecksum
func (m_ ManagedObjectModel) VersionChecksum() string {
	rv := objc.Send[string](m_.ID, objc.Sel("versionChecksum"))
	return rv
}


// SetVersionChecksum sets the value of the versionChecksum property.
// The Base64-encoded 128-bit model version hash.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionchecksum
func (m_ ManagedObjectModel) SetVersionChecksum(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersionChecksum:"), objc.String(value))
}

// The set of developer-defined version identifiers for the object model.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionidentifiers
func (m_ ManagedObjectModel) VersionIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("versionIdentifiers"))
	return rv
}


// SetVersionIdentifiers sets the value of the versionIdentifiers property.
// The set of developer-defined version identifiers for the object model.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionidentifiers
func (m_ ManagedObjectModel) SetVersionIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersionIdentifiers:"), value)
}



