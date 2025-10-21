// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ContentRuleListStore] class.
type IContentRuleListStore interface {
	objectivec.IObject
	RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer)
}

// An object that contains the rules for how to load and filter content in the web view.
//
// Use a to compile and manage rules for filtering content in a web view. Rule lists act as content blockers inside your app. You use them to prevent the web view from loading specific content, either based on the original location of that content or other criteria you specify. For example, a corporate app might use rules to prevent the web view from loading content that originates from outside the corporate network. Fetch the default object or create a custom store object and use it to compile or access the available rules. Each store object stores its existing rules persistently in the file system and loads those rules at creation time. A store object doesn’t automatically apply any of its rules to a particular web view. To apply a rule to a web view, add it to the object of the web view’s configuration object.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ContentRuleListStoreClass) Alloc() ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new content rule list store in the specified directory.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/init(url:)
func NewContentRuleListStoreWithURL(url unsafe.Pointer) ContentRuleListStore {
	rv := objc.Send[ContentRuleListStore](objc.ID(getContentRuleListStoreClass().class), objc.Sel("storeWithURL:"), url)
	return rv
}


// Creates a new content rule list store in the specified directory.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/init(url:)
func (cc _ContentRuleListStoreClass) StoreWithURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("storeWithURL:"), url)
	return rv
}

// Removes a rule list from the current data store asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/removeContentRuleList(forIdentifier:completionHandler:)
func (c_ ContentRuleListStore) RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeContentRuleListForIdentifier:completionHandler:"), objc.String(identifier), completionHandler)
}


