// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebsiteDataStore */


/* debug [class_header]: Header for WKWebsiteDataStore */
// The class instance for the [WebsiteDataStore] class.
var (
	WebsiteDataStoreClass     _WebsiteDataStoreClass
	WebsiteDataStoreClassOnce sync.Once
)

func getWebsiteDataStoreClass() _WebsiteDataStoreClass {
	WebsiteDataStoreClassOnce.Do(func() {
		WebsiteDataStoreClass = _WebsiteDataStoreClass{objc.GetClass("WKWebsiteDataStore")}
	})
	return WebsiteDataStoreClass
}

type _WebsiteDataStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebsiteDataStore */
// An interface definition for the [WebsiteDataStore] class.
type IWebsiteDataStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebsiteDataStore */
	// properties:
	HttpCookieStore() IWKHTTPCookieStore
	Identifier() foundation.UUID
	Persistent() bool
	ProxyConfigurations() []objc.IObject /* cross-framework: Object */
	SetProxyConfigurations(value []objc.IObject /* cross-framework: Object */)
	WebsiteDataStore() IWKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	IsPersistent() bool
	SetIsPersistent(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebsiteDataStore */
	// methods:
	FetchDataOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	FetchDataRecordsOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler func([]unsafe.Pointer))
	RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes unsafe.Pointer, dataRecords []WebsiteDataRecord, completionHandler func())
	RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes unsafe.Pointer, date objc.IObject /* cross-framework: NSDate */, completionHandler func())
	RestoreDataCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebsiteDataStore */
// Alloc allocates a new instance without initialization.
func (wc _WebsiteDataStoreClass) Alloc() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebsiteDataStoreClass) New() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebsiteDataStore) Init() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebsiteDataStore) Autorelease() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebsiteDataStore creates a new WebsiteDataStore instance.
func NewWebsiteDataStore() WebsiteDataStore {
	return getWebsiteDataStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebsiteDataStore */
// An object that manages cookies, disk and memory caches, and other types of data for a web view.
//
// Use a object to configure and manage web site data. Specifically, use this object to: Manage cookies that your web site uses Learn about the types of data that websites store Remove unwanted web site data Create a data store object and assign it to the property of a object before you create your web view. By default, uses the default data store returned by the method, which saves website data persistently to disk. To implement private browsing, create a nonpersistent data store using the method instead. To implement profile browsing, create a persistent data store using the method, passing an identifier that you use to identify the data store.


// An object that manages cookies, disk and memory caches, and other types of data for a web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore
type WebsiteDataStore struct {
	objectivec.Object
}

// WebsiteDataStoreFrom constructs a [WebsiteDataStore] from an unsafe.Pointer.
//
// An object that manages cookies, disk and memory caches, and other types of data for a web view.
func WebsiteDataStoreFrom(ptr unsafe.Pointer) WebsiteDataStore {
	return WebsiteDataStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebsiteDataStore */

// Returns the persistent data store with the unique identifier you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/init(forIdentifier:)
func NewWebsiteDataStoreForIdentifier(identifier foundation.UUID) WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(getWebsiteDataStoreClass().class), objc.Sel("dataStoreForIdentifier:"), identifier)
	return rv
}/* debug [class_init_methods/constructor]: NewWebsiteDataStoreForIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebsiteDataStore */

// Returns the set of all the available data types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/allWebsiteDataTypes()
func (wc _WebsiteDataStoreClass) AllWebsiteDataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("allWebsiteDataTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllWebsiteDataTypes) */


// Returns the default data store, which stores data persistently to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/default()
func (wc _WebsiteDataStoreClass) DefaultDataStore() IWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("defaultDataStore"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultDataStore) */


// Fetches an array of identifiers from existing data stores that have identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/fetchAllDataStoreIdentifiers(_:)
func (wc _WebsiteDataStoreClass) FetchAllDataStoreIdentifiers(completionHandler func([]unsafe.Pointer)) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("fetchAllDataStoreIdentifiers:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FetchAllDataStoreIdentifiers) */


// Returns the persistent data store with the unique identifier you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/init(forIdentifier:)
func (wc _WebsiteDataStoreClass) DataStoreForIdentifier(identifier foundation.UUID) IWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("dataStoreForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataStoreForIdentifier) */


// Creates a new data store object that stores website data in memory, and doesn’t write that data to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/nonPersistent()
func (wc _WebsiteDataStoreClass) NonPersistentDataStore() IWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("nonPersistentDataStore"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NonPersistentDataStore) */


// Removes the data store that matches the identifier you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/remove(forIdentifier:completionHandler:)
func (wc _WebsiteDataStoreClass) RemoveDataStoreForIdentifierCompletionHandler(identifier foundation.UUID, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("removeDataStoreForIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveDataStoreForIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebsiteDataStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebsiteDataStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/fetchData(of:completionHandler:)
func (w_ WebsiteDataStore) FetchDataOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataOfTypes:completionHandler:"), dataTypes, completionHandler)
}/* debug [instance_methods/method]: FetchDataOfTypesCompletionHandler */


// Fetches the specified types of records from the data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/fetchDataRecords(ofTypes:completionHandler:)
func (w_ WebsiteDataStore) FetchDataRecordsOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler func([]unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataRecordsOfTypes:completionHandler:"), dataTypes, completionHandler)
}/* debug [instance_methods/method]: FetchDataRecordsOfTypesCompletionHandler */


// Removes the specified types of website data from one or more data records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/removeData(ofTypes:for:completionHandler:)
func (w_ WebsiteDataStore) RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes unsafe.Pointer, dataRecords []WebsiteDataRecord, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeDataOfTypes:forDataRecords:completionHandler:"), dataTypes, dataRecords, completionHandler)
}/* debug [instance_methods/method]: RemoveDataOfTypesForDataRecordsCompletionHandler */


// Removes website data that changed after the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/removeData(ofTypes:modifiedSince:completionHandler:)
func (w_ WebsiteDataStore) RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes unsafe.Pointer, date objc.IObject /* cross-framework: NSDate */, completionHandler func()) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeDataOfTypes:modifiedSince:completionHandler:"), dataTypes, date, completionHandler)
}/* debug [instance_methods/method]: RemoveDataOfTypesModifiedSinceCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/restoreData(_:completionHandler:)
func (w_ WebsiteDataStore) RestoreDataCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](w_.ID, objc.Sel("restoreData:completionHandler:"), data, completionHandler)
}/* debug [instance_methods/method]: RestoreDataCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebsiteDataStore */

// The object that manages the HTTP cookies for your website.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/httpCookieStore
func (w_ WebsiteDataStore) HttpCookieStore() IWKHTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](w_.ID, objc.Sel("httpCookieStore"))
	return rv
}/* debug [instance_properties/getter]: httpCookieStore */


// An identifier that uniquely identifies a data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/identifier
func (w_ WebsiteDataStore) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](w_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether this object stores data to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/isPersistent
func (w_ WebsiteDataStore) Persistent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("persistent"))
	return rv
}/* debug [instance_properties/getter]: persistent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/proxyConfigurations-6g21z
func (w_ WebsiteDataStore) ProxyConfigurations() []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](w_.ID, objc.Sel("proxyConfigurations"))
	return rv
}/* debug [instance_properties/getter]: proxyConfigurations */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/proxyConfigurations-6g21z
func (w_ WebsiteDataStore) SetProxyConfigurations(value []objc.IObject /* cross-framework: Object */) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](w_.ID, objc.Sel("setProxyConfigurations:"), nsArray)
}/* debug [instance_properties/setter]: proxyConfigurations */


// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/websitedatastore
func (w_ WebsiteDataStore) WebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](w_.ID, objc.Sel("websiteDataStore"))
	return rv
}/* debug [instance_properties/getter]: websiteDataStore */


// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/websitedatastore
func (w_ WebsiteDataStore) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebsiteDataStore:"), value)
}/* debug [instance_properties/setter]: websiteDataStore */


// A Boolean value that indicates whether this object stores data to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/ispersistent
func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isPersistent"))
	return rv
}/* debug [instance_properties/getter]: isPersistent */


// A Boolean value that indicates whether this object stores data to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/ispersistent
func (w_ WebsiteDataStore) SetIsPersistent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsPersistent:"), value)
}/* debug [instance_properties/setter]: isPersistent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebsiteDataStore */


