// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKContentRuleListStore */

/* debug [class_header]: Header for WKContentRuleListStore */
// The class instance for the [ContentRuleListStore] class.
var (
	ContentRuleListStoreClass     _ContentRuleListStoreClass
	ContentRuleListStoreClassOnce sync.Once
)

func getContentRuleListStoreClass() _ContentRuleListStoreClass {
	ContentRuleListStoreClassOnce.Do(func() {
		ContentRuleListStoreClass = _ContentRuleListStoreClass{objc.GetClass("WKContentRuleListStore")}
	})
	return ContentRuleListStoreClass
}

type _ContentRuleListStoreClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for ContentRuleListStore */
// An interface definition for the [ContentRuleListStore] class.
type IContentRuleListStore interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for ContentRuleListStore */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for ContentRuleListStore */
	// methods:
	CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, encodedContentRuleList objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	GetAvailableContentRuleListIdentifiers(completionHandler unsafe.Pointer)
	LookUpContentRuleListForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	RemoveContentRuleListForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for ContentRuleListStore */
// Alloc allocates a new instance without initialization.
func (cc _ContentRuleListStoreClass) Alloc() ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentRuleListStoreClass) New() ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentRuleListStore) Init() ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentRuleListStore) Autorelease() ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentRuleListStore creates a new ContentRuleListStore instance.
func NewContentRuleListStore() ContentRuleListStore {
	return getContentRuleListStoreClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for ContentRuleListStore */
// An object that contains the rules for how to load and filter content in the web view.
//
// Use a to compile and manage rules for filtering content in a web view. Rule lists act as content blockers inside your app. You use them to prevent the web view from loading specific content, either based on the original location of that content or other criteria you specify. For example, a corporate app might use rules to prevent the web view from loading content that originates from outside the corporate network. Fetch the default object or create a custom store object and use it to compile or access the available rules. Each store object stores its existing rules persistently in the file system and loads those rules at creation time. A store object doesn’t automatically apply any of its rules to a particular web view. To apply a rule to a web view, add it to the object of the web view’s configuration object.

// An object that contains the rules for how to load and filter content in the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore
type ContentRuleListStore struct {
	objectivec.Object
}

// ContentRuleListStoreFrom constructs a [ContentRuleListStore] from an unsafe.Pointer.
//
// An object that contains the rules for how to load and filter content in the web view.
func ContentRuleListStoreFrom(ptr unsafe.Pointer) ContentRuleListStore {
	return ContentRuleListStore{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for ContentRuleListStore */

// Creates a new content rule list store in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/init(url:)
func NewContentRuleListStoreWithURL(url objc.IObject /* cross-framework: NSURL */) ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](objc.ID(getContentRuleListStoreClass().class), objc.Sel("storeWithURL:"), url)
	return rv
} /* debug [class_init_methods/constructor]: NewContentRuleListStoreWithURL */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for ContentRuleListStore */

// Returns the default content rule list store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/default()
func (cc _ContentRuleListStoreClass) DefaultStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("defaultStore"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultStore) */

// Creates a new content rule list store in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/init(url:)
func (cc _ContentRuleListStoreClass) StoreWithURL(url objc.IObject /* cross-framework: NSURL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("storeWithURL:"), url)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=StoreWithURL) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for ContentRuleListStore */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for ContentRuleListStore */

// Compiles the specified JSON content into a new rule list and adds it to the current data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/compileContentRuleList(forIdentifier:encodedContentRuleList:completionHandler:)
func (c_ ContentRuleListStore) CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, encodedContentRuleList objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("compileContentRuleListForIdentifier:encodedContentRuleList:completionHandler:"), identifier, encodedContentRuleList, completionHandler)
} /* debug [instance_methods/method]: CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler */

// Fetches the identifiers for all rule lists in the store asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/getAvailableContentRuleListIdentifiers(_:)
func (c_ ContentRuleListStore) GetAvailableContentRuleListIdentifiers(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getAvailableContentRuleListIdentifiers:"), completionHandler)
} /* debug [instance_methods/method]: GetAvailableContentRuleListIdentifiers */

// Searches asynchronously for a specific rule list in the data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/lookUpContentRuleList(forIdentifier:completionHandler:)
func (c_ ContentRuleListStore) LookUpContentRuleListForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("lookUpContentRuleListForIdentifier:completionHandler:"), identifier, completionHandler)
} /* debug [instance_methods/method]: LookUpContentRuleListForIdentifierCompletionHandler */

// Removes a rule list from the current data store asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/removeContentRuleList(forIdentifier:completionHandler:)
func (c_ ContentRuleListStore) RemoveContentRuleListForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeContentRuleListForIdentifier:completionHandler:"), identifier, completionHandler)
} /* debug [instance_methods/method]: RemoveContentRuleListForIdentifierCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for ContentRuleListStore */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKContentRuleListStore */
