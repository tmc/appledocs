// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [UserDefaults] class.
var (
	UserDefaultsClass     _UserDefaultsClass
	UserDefaultsClassOnce sync.Once
)

func getUserDefaultsClass() _UserDefaultsClass {
	UserDefaultsClassOnce.Do(func() {
		UserDefaultsClass = _UserDefaultsClass{objc.GetClass("NSUserDefaults")}
	})
	return UserDefaultsClass
}

type _UserDefaultsClass struct {
	class objc.Class
}





// An interface definition for the [UserDefaults] class.
type IUserDefaults interface {
	objectivec.IObject
	

	// properties:
	VolatileDomainNames() []string


	

	// methods:
	AddSuiteNamed(suiteName IString)
	ArrayForKey(defaultName IString) IArray
	BoolForKey(defaultName IString) bool
	DataForKey(defaultName IString) IData
	DictionaryForKey(defaultName IString) IDictionary
	DictionaryRepresentation() IDictionary
	DoubleForKey(defaultName IString) float64
	FloatForKey(defaultName IString) float32
	IntegerForKey(defaultName IString) int
	ObjectForKey(defaultName IString) objc.ID
	ObjectIsForcedForKey(key IString) bool
	ObjectIsForcedForKeyInDomain(key IString, domain IString) bool
	PersistentDomainForName(domainName IString) IDictionary
	RegisterDefaults(registrationDictionary IDictionary)
	RemoveObjectForKey(defaultName IString)
	RemovePersistentDomainForName(domainName IString)
	RemoveSuiteNamed(suiteName IString)
	RemoveVolatileDomainForName(domainName IString)
	SetFloatForKey(value float32, defaultName IString)
	SetURLForKey(url IURL, defaultName IString)
	SetDoubleForKey(value float64, defaultName IString)
	SetBoolForKey(value bool, defaultName IString)
	SetIntegerForKey(value int, defaultName IString)
	SetObjectForKey(value objectivec.IObject, defaultName IString)
	SetPersistentDomainForName(domain IDictionary, domainName IString)
	SetVolatileDomainForName(domain IDictionary, domainName IString)
	StringForKey(defaultName IString) IString
	StringArrayForKey(defaultName IString) []string
	URLForKey(defaultName IString) IURL
	VolatileDomainForName(domainName IString) IDictionary


}





// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsClass) Alloc() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UserDefaultsClass) New() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserDefaults) Init() UserDefaults {
	rv := objc.Send[UserDefaults](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserDefaults) Autorelease() UserDefaults {
	rv := objc.Send[UserDefaults](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserDefaults creates a new UserDefaults instance.
func NewUserDefaults() UserDefaults {
	return getUserDefaultsClass().New()
}





// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
//
// The class provides a programmatic interface for interacting with the defaults system. The defaults system allows an app to customize its behavior to match a user’s preferences. For example, you can allow users to specify their preferred units of measurement or media playback speed. Apps store these preferences by assigning values to a set of parameters in a user’s defaults database. The parameters are referred to as because they’re commonly used to determine an app’s default state at startup or the way it acts by default. At runtime, you use objects to read the defaults that your app uses from a user’s defaults database. caches the information to avoid having to open the user’s defaults database each time you need a default value. When you set a default value, it’s changed synchronously within your process, and asynchronously to persistent storage and other processes. With the exception of managed devices in educational institutions, a user’s defaults are stored locally on a single device, and persisted for backup and restore. To synchronize preferences and other data across a user’s connected devices, use instead.


// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults
type UserDefaults struct {
	objectivec.Object
}

// UserDefaultsFrom constructs a [UserDefaults] from an unsafe.Pointer.
//
// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
func UserDefaultsFrom(ptr unsafe.Pointer) UserDefaults {
	return UserDefaults{objectivec.Object{objc.ID(ptr)}}
}






// Creates a user defaults object initialized with the defaults for the specified database name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/init(suiteName:)
func NewUserDefaultsWithSuiteName(suitename IString) UserDefaults {
	instance := getUserDefaultsClass().Alloc()
	rv := objc.Send[UserDefaults](instance.ID, objc.Sel("initWithSuiteName:"), suitename)
	rv.Autorelease()
	return rv
}












// Returns the shared defaults object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/standard
func (uc _UserDefaultsClass) StandardUserDefaults() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("standardUserDefaults"))
	return rv
}






// Inserts the specified domain name into the receiver’s search list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/addSuite(named:)
func (u_ UserDefaults) AddSuiteNamed(suiteName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addSuiteNamed:"), suiteName)
}


// Returns the array associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/array(forKey:)
func (u_ UserDefaults) ArrayForKey(defaultName IString) IArray {
	rv := objc.Send[Array](u_.ID, objc.Sel("arrayForKey:"), defaultName)
	return rv
}


// Returns the Boolean value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/bool(forKey:)
func (u_ UserDefaults) BoolForKey(defaultName IString) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("boolForKey:"), defaultName)
	return rv
}


// Returns the data object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/data(forKey:)
func (u_ UserDefaults) DataForKey(defaultName IString) IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("dataForKey:"), defaultName)
	return rv
}


// Returns the dictionary object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/dictionary(forKey:)
func (u_ UserDefaults) DictionaryForKey(defaultName IString) IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("dictionaryForKey:"), defaultName)
	return rv
}


// Returns a dictionary that contains a union of all key-value pairs in the domains in the search list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/dictionaryRepresentation()
func (u_ UserDefaults) DictionaryRepresentation() IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}


// Returns the double value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/double(forKey:)
func (u_ UserDefaults) DoubleForKey(defaultName IString) float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("doubleForKey:"), defaultName)
	return rv
}


// Returns the float value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/float(forKey:)
func (u_ UserDefaults) FloatForKey(defaultName IString) float32 {
	rv := objc.Send[float32](u_.ID, objc.Sel("floatForKey:"), defaultName)
	return rv
}


// Returns the integer value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/integer(forKey:)
func (u_ UserDefaults) IntegerForKey(defaultName IString) int {
	rv := objc.Send[int](u_.ID, objc.Sel("integerForKey:"), defaultName)
	return rv
}


// Returns the object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/object(forKey:)
func (u_ UserDefaults) ObjectForKey(defaultName IString) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("objectForKey:"), defaultName)
	return rv
}


// Returns a Boolean value indicating whether the specified key is managed by an administrator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/objectIsForced(forKey:)
func (u_ UserDefaults) ObjectIsForcedForKey(key IString) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("objectIsForcedForKey:"), key)
	return rv
}


// Returns a Boolean value indicating whether the key in the specified domain is managed by an administrator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/objectIsForced(forKey:inDomain:)
func (u_ UserDefaults) ObjectIsForcedForKeyInDomain(key IString, domain IString) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("objectIsForcedForKey:inDomain:"), key, domain)
	return rv
}


// Returns a dictionary representation of the defaults for the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/persistentDomain(forName:)
func (u_ UserDefaults) PersistentDomainForName(domainName IString) IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("persistentDomainForName:"), domainName)
	return rv
}


// Adds the contents of the specified dictionary to the registration domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/register(defaults:)
func (u_ UserDefaults) RegisterDefaults(registrationDictionary IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("registerDefaults:"), registrationDictionary)
}


// Removes the value of the specified default key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/removeObject(forKey:)
func (u_ UserDefaults) RemoveObjectForKey(defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeObjectForKey:"), defaultName)
}


// Removes the contents of the specified persistent domain from the user’s defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/removePersistentDomain(forName:)
func (u_ UserDefaults) RemovePersistentDomainForName(domainName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removePersistentDomainForName:"), domainName)
}


// Removes the specified domain name from the receiver’s search list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/removeSuite(named:)
func (u_ UserDefaults) RemoveSuiteNamed(suiteName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeSuiteNamed:"), suiteName)
}


// Removes the specified volatile domain from the user’s defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/removeVolatileDomain(forName:)
func (u_ UserDefaults) RemoveVolatileDomainForName(domainName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeVolatileDomainForName:"), domainName)
}


// Sets the value of the specified default key to the specified float value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-1t5ec
func (u_ UserDefaults) SetFloatForKey(value float32, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFloat:forKey:"), value, defaultName)
}


// Sets the value of the specified default key to the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-2bqjt
func (u_ UserDefaults) SetURLForKey(url IURL, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setURL:forKey:"), url, defaultName)
}


// Sets the value of the specified default key to the double value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-2w22f
func (u_ UserDefaults) SetDoubleForKey(value float64, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDouble:forKey:"), value, defaultName)
}


// Sets the value of the specified default key to the specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-3nn5m
func (u_ UserDefaults) SetBoolForKey(value bool, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBool:forKey:"), value, defaultName)
}


// Sets the value of the specified default key to the specified integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-3v852
func (u_ UserDefaults) SetIntegerForKey(value int, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInteger:forKey:"), value, defaultName)
}


// Sets the value of the specified default key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-8ab6d
func (u_ UserDefaults) SetObjectForKey(value objectivec.IObject, defaultName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setObject:forKey:"), value, defaultName)
}


// Sets a dictionary for the specified persistent domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/setPersistentDomain(_:forName:)
func (u_ UserDefaults) SetPersistentDomainForName(domain IDictionary, domainName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPersistentDomain:forName:"), domain, domainName)
}


// Sets the dictionary for the specified volatile domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/setVolatileDomain(_:forName:)
func (u_ UserDefaults) SetVolatileDomainForName(domain IDictionary, domainName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setVolatileDomain:forName:"), domain, domainName)
}


// Returns the string associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/string(forKey:)
func (u_ UserDefaults) StringForKey(defaultName IString) IString {
	rv := objc.Send[String](u_.ID, objc.Sel("stringForKey:"), defaultName)
	return rv
}


// Returns the array of strings associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/stringArray(forKey:)
func (u_ UserDefaults) StringArrayForKey(defaultName IString) []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("stringArrayForKey:"), defaultName)
	return rv
}


// Returns the URL associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/url(forKey:)
func (u_ UserDefaults) URLForKey(defaultName IString) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLForKey:"), defaultName)
	return rv
}


// Returns the dictionary for the specified volatile domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/volatileDomain(forName:)
func (u_ UserDefaults) VolatileDomainForName(domainName IString) IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("volatileDomainForName:"), domainName)
	return rv
}







// Returns the shared defaults object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/standard
func (u_ UserDefaults) StandardUserDefaults() IUserDefaults {
	rv := objc.Send[UserDefaults](u_.ID, objc.Sel("standardUserDefaults"))
	return rv
}


// The current volatile domain names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/volatileDomainNames
func (u_ UserDefaults) VolatileDomainNames() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("volatileDomainNames"))
	return rv
}







