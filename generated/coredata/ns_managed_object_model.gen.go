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
	// properties:
	Configurations() string /* primitive/slice/pointer. */
	SetConfigurations(value string /* primitive/slice/pointer. */)
	Entities() IEntityDescription
	SetEntities(value IEntityDescription)
	EntitiesByName() IEntityDescription
	SetEntitiesByName(value IEntityDescription)
	EntityVersionHashesByName() foundation.objc.IObject /* cross-framework: Data */
	SetEntityVersionHashesByName(value foundation.objc.IObject /* cross-framework: Data */)
	FetchRequestTemplatesByName() FetchRequestResult /* not a class type */
	SetFetchRequestTemplatesByName(value FetchRequestResult /* not a class type */)
	LocalizationDictionary() string /* primitive/slice/pointer. */
	SetLocalizationDictionary(value string /* primitive/slice/pointer. */)
	VersionChecksum() string /* primitive/slice/pointer. */
	SetVersionChecksum(value string /* primitive/slice/pointer. */)
	VersionIdentifiers() unsafe.Pointer
	SetVersionIdentifiers(value unsafe.Pointer)
	// methods:
	IsConfigurationCompatibleWithStoreMetadata(configuration string /* primitive/slice/pointer. */, metadata foundation.IDictionary /* already interface */) bool /* primitive/slice/pointer. */
	SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string /* primitive/slice/pointer. */)
}

// A programmatic representation of the file describing your objects.
//
// The model contains one or more objects representing the entities in the schema. Each object has property description objects (instances of subclasses of ) that represent the properties (or fields) of the entity in the schema. The Core Data framework uses this description in several ways: Constraining UI creation in Interface Builder Validating attribute and relationship values at runtime Mapping between your managed objects and a database or file-based schema for object persistence A managed object model maintains a mapping between each of its entity objects and a corresponding managed object class for use with the persistent storage mechanisms in the Core Data framework. You can determine the entity for a particular managed object with the method. You typically create managed object models using the data modeling tool in Xcode, but it’s possible to build a model programmatically if needed.


// A programmatic representation of the file describing your objects.
//
// [Full Topic]
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



// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/isConfiguration(withName:compatibleWithStoreMetadata:)
func (m_ ManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string /* primitive/slice/pointer. */, metadata foundation.IDictionary /* already interface */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), objc.String(configuration), metadata)
	return rv
}


// Associates the specified fetch request with the receiver using the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setFetchRequestTemplate(_:forName:)
func (m_ ManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplate:forName:"), fetchRequestTemplate, objc.String(name))
}


// All the available configuration names of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/configurations
func (m_ ManagedObjectModel) Configurations() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("configurations"))
	return rv
}


// All the available configuration names of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/configurations
func (m_ ManagedObjectModel) SetConfigurations(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfigurations:"), objc.String(value))
}


// The entities in the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entities
func (m_ ManagedObjectModel) Entities() IEntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("entities"))
	return rv
}


// The entities in the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entities
func (m_ ManagedObjectModel) SetEntities(value IEntityDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntities:"), value)
}


// The entities of the model, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entitiesbyname
func (m_ ManagedObjectModel) EntitiesByName() IEntityDescription {
	rv := objc.Send[EntityDescription](m_.ID, objc.Sel("entitiesByName"))
	return rv
}


// The entities of the model, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entitiesbyname
func (m_ ManagedObjectModel) SetEntitiesByName(value IEntityDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntitiesByName:"), value)
}


// The dictionary of the model’s entity names and their corresponding version hashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entityversionhashesbyname
func (m_ ManagedObjectModel) EntityVersionHashesByName() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("entityVersionHashesByName"))
	return rv
}


// The dictionary of the model’s entity names and their corresponding version hashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/entityversionhashesbyname
func (m_ ManagedObjectModel) SetEntityVersionHashesByName(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntityVersionHashesByName:"), value)
}


// A dictionary of the receiver’s fetch request templates, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/fetchrequesttemplatesbyname
func (m_ ManagedObjectModel) FetchRequestTemplatesByName() FetchRequestResult /* not a class type */ {
	rv := objc.Send[FetchRequestResult](m_.ID, objc.Sel("fetchRequestTemplatesByName"))
	return rv
}


// A dictionary of the receiver’s fetch request templates, keyed by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/fetchrequesttemplatesbyname
func (m_ ManagedObjectModel) SetFetchRequestTemplatesByName(value FetchRequestResult /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplatesByName:"), value)
}


// The localization dictionary of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/localizationdictionary
func (m_ ManagedObjectModel) LocalizationDictionary() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("localizationDictionary"))
	return rv
}


// The localization dictionary of the model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/localizationdictionary
func (m_ ManagedObjectModel) SetLocalizationDictionary(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalizationDictionary:"), objc.String(value))
}


// The Base64-encoded 128-bit model version hash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionchecksum
func (m_ ManagedObjectModel) VersionChecksum() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("versionChecksum"))
	return rv
}


// The Base64-encoded 128-bit model version hash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionchecksum
func (m_ ManagedObjectModel) SetVersionChecksum(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersionChecksum:"), objc.String(value))
}


// The set of developer-defined version identifiers for the object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionidentifiers
func (m_ ManagedObjectModel) VersionIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("versionIdentifiers"))
	return rv
}


// The set of developer-defined version identifiers for the object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/versionidentifiers
func (m_ ManagedObjectModel) SetVersionIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersionIdentifiers:"), value)
}



