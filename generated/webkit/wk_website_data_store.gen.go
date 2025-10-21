// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (wc _WebsiteDataStoreClass) DefaultDataStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("defaultDataStore"))
	return rv
}

// Creates a new data store object that stores website data in memory, and doesn’t write that data to disk.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/nonPersistent()
func (wc _WebsiteDataStoreClass) NonPersistentDataStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("nonPersistentDataStore"))
	return rv
}



