// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [WebsiteDataStore] class.
type IWebsiteDataStore interface {
	objectivec.IObject
	WebsiteDataStore() WKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	HttpCookieStore() WKHTTPCookieStore
	SetHttpCookieStore(value IWKHTTPCookieStore)
	Identifier() foundation.UUID
	SetIdentifier(value foundation.IUUID)
	IsPersistent() bool
	SetIsPersistent(value bool)
	ProxyConfigurations() unsafe.Pointer
	SetProxyConfigurations(value unsafe.Pointer)
}

// An object that manages cookies, disk and memory caches, and other types of data for a web view.
//
// Use a object to configure and manage web site data. Specifically, use this object to: Manage cookies that your web site uses Learn about the types of data that websites store Remove unwanted web site data Create a data store object and assign it to the property of a object before you create your web view. By default, uses the default data store returned by the method, which saves website data persistently to disk. To implement private browsing, create a nonpersistent data store using the method instead. To implement profile browsing, create a persistent data store using the method, passing an identifier that you use to identify the data store.
//
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

// Alloc allocates a new instance without initialization.
func (wc _WebsiteDataStoreClass) Alloc() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the default data store, which stores data persistently to disk.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/default()
func (wc _WebsiteDataStoreClass) DefaultDataStore() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("defaultDataStore"))
	return rv
}

// Creates a new data store object that stores website data in memory, and doesn’t write that data to disk.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/nonPersistent()
func (wc _WebsiteDataStoreClass) NonPersistentDataStore() WebsiteDataStore {
	rv := objc.Send[WebsiteDataStore](objc.ID(wc.class), objc.Sel("nonPersistentDataStore"))
	return rv
}

// The object you use to get and set the site’s cookies and to track the cached data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/websitedatastore
func (w_ WebsiteDataStore) WebsiteDataStore() WKWebsiteDataStore {
	rv := objc.Send[WKWebsiteDataStore](w_.ID, objc.Sel("websiteDataStore"))
	return rv
}


// SetWebsiteDataStore sets the value of the websiteDataStore property.
// The object you use to get and set the site’s cookies and to track the cached data objects.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/websitedatastore
func (w_ WebsiteDataStore) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebsiteDataStore:"), value)
}

// The object that manages the HTTP cookies for your website.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/httpcookiestore
func (w_ WebsiteDataStore) HttpCookieStore() WKHTTPCookieStore {
	rv := objc.Send[WKHTTPCookieStore](w_.ID, objc.Sel("httpCookieStore"))
	return rv
}


// SetHttpCookieStore sets the value of the httpCookieStore property.
// The object that manages the HTTP cookies for your website.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/httpcookiestore
func (w_ WebsiteDataStore) SetHttpCookieStore(value IWKHTTPCookieStore) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHttpCookieStore:"), value)
}

// An identifier that uniquely identifies a data store.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/identifier
func (w_ WebsiteDataStore) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](w_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// An identifier that uniquely identifies a data store.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/identifier
func (w_ WebsiteDataStore) SetIdentifier(value foundation.IUUID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIdentifier:"), value)
}

// A Boolean value that indicates whether this object stores data to disk.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/ispersistent
func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isPersistent"))
	return rv
}


// SetIsPersistent sets the value of the isPersistent property.
// A Boolean value that indicates whether this object stores data to disk.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/ispersistent
func (w_ WebsiteDataStore) SetIsPersistent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsPersistent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/proxyconfigurations-cdc1
func (w_ WebsiteDataStore) ProxyConfigurations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("proxyConfigurations"))
	return rv
}


// SetProxyConfigurations sets the value of the proxyConfigurations property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/proxyconfigurations-cdc1
func (w_ WebsiteDataStore) SetProxyConfigurations(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProxyConfigurations:"), value)
}



