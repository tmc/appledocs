// Code generated from Apple documentation for ThreadNetwork. DO NOT EDIT.

package threadnetwork

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class THClient */

/* debug [class_header]: Header for THClient */
// The class instance for the [THClient] class.
var (
	THClientClass     _THClientClass
	THClientClassOnce sync.Once
)

func getTHClientClass() _THClientClass {
	THClientClassOnce.Do(func() {
		THClientClass = _THClientClass{objc.GetClass("THClient")}
	})
	return THClientClass
}

type _THClientClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for THClient */
// An interface definition for the [THClient] class.
type ITHClient interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for THClient */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for THClient */
	// methods:
	CheckPreferredNetworkForActiveOperationalDatasetCompletion(activeOperationalDataSet objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	DeleteCredentialsForBorderAgentCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	IsPreferredNetworkAvailableWithCompletion(completion unsafe.Pointer)
	RetrieveAllActiveCredentials(completion unsafe.Pointer)
	RetrieveAllCredentials(completion unsafe.Pointer)
	RetrieveCredentialsForBorderAgentCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	RetrieveCredentialsForExtendedPANIDCompletion(extendedPANID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	RetrievePreferredCredentials(completion unsafe.Pointer)
	StoreCredentialsForBorderAgentActiveOperationalDataSetCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, activeOperationalDataSet objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for THClient */
// Alloc allocates a new instance without initialization.
func (tc _THClientClass) Alloc() THClient {
	rv := objc.Send[THClient](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _THClientClass) New() THClient {
	rv := objc.Send[THClient](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ THClient) Init() THClient {
	rv := objc.Send[THClient](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ THClient) Autorelease() THClient {
	rv := objc.Send[THClient](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTHClient creates a new THClient instance.
func NewTHClient() THClient {
	return getTHClientClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for THClient */
// A class that supports safely sharing Thread credentials between multiple clients.
//
// Request credentials for either a specific Thread network or for the using . The preferred network is the default Thread network chosen by the framework for a home. The ThreadNetwork framework maintains a database of network credentials. The class allows clients to store, list, and delete credentials for a given network from the database. Some methods in use the , a string that you store in your application’s . The ThreadNetwork framework uses the team ID to preserve the privacy of the Thread network credentials across different clients. For example, credentials stored by one client can’t be deleted or modified by another client.

// A class that supports safely sharing Thread credentials between multiple clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient
type THClient struct {
	objectivec.Object
}

// THClientFrom constructs a [THClient] from an unsafe.Pointer.
//
// A class that supports safely sharing Thread credentials between multiple clients.
func THClientFrom(ptr unsafe.Pointer) THClient {
	return THClient{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for THClient */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for THClient */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for THClient */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for THClient */

// Determines if the essential operating parameters match the preferred network’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/checkPreferredNetwork(forActiveOperationalDataset:completion:)
func (t_ THClient) CheckPreferredNetworkForActiveOperationalDatasetCompletion(activeOperationalDataSet objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkPreferredNetworkForActiveOperationalDataset:completion:"), activeOperationalDataSet, completion)
} /* debug [instance_methods/method]: CheckPreferredNetworkForActiveOperationalDatasetCompletion */

// Deletes Thread network credentials from the framework database for a Border Agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/deleteCredentials(forBorderAgent:completion:)
func (t_ THClient) DeleteCredentialsForBorderAgentCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deleteCredentialsForBorderAgent:completion:"), borderAgentID, completion)
} /* debug [instance_methods/method]: DeleteCredentialsForBorderAgentCompletion */

// Indicates whether a preferred network is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/isPreferredNetworkAvailable(completion:)
func (t_ THClient) IsPreferredNetworkAvailableWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("isPreferredNetworkAvailableWithCompletion:"), completion)
} /* debug [instance_methods/method]: IsPreferredNetworkAvailableWithCompletion */

// Returns a set of the active credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/retrieveAllActiveCredentials(_:)
func (t_ THClient) RetrieveAllActiveCredentials(completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("retrieveAllActiveCredentials:"), completion)
} /* debug [instance_methods/method]: RetrieveAllActiveCredentials */

// Requests all Thread credentials from the framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/retrieveAllCredentials(_:)
func (t_ THClient) RetrieveAllCredentials(completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("retrieveAllCredentials:"), completion)
} /* debug [instance_methods/method]: RetrieveAllCredentials */

// Requests Thread credentials for a Border Agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/retrieveCredentials(forBorderAgent:completion:)
func (t_ THClient) RetrieveCredentialsForBorderAgentCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("retrieveCredentialsForBorderAgent:completion:"), borderAgentID, completion)
} /* debug [instance_methods/method]: RetrieveCredentialsForBorderAgentCompletion */

// Requests Thread credentials for an extended Personal Area Network (PAN) ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/retrieveCredentials(forExtendedPANID:completion:)
func (t_ THClient) RetrieveCredentialsForExtendedPANIDCompletion(extendedPANID objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("retrieveCredentialsForExtendedPANID:completion:"), extendedPANID, completion)
} /* debug [instance_methods/method]: RetrieveCredentialsForExtendedPANIDCompletion */

// Requests Thread credentials for the preferred network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/retrievePreferredCredentials(_:)
func (t_ THClient) RetrievePreferredCredentials(completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("retrievePreferredCredentials:"), completion)
} /* debug [instance_methods/method]: RetrievePreferredCredentials */

// Stores Thread network credentials into the framework database that a Border Agent provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THClient/storeCredentials(forBorderAgent:activeOperationalDataSet:completion:)
func (t_ THClient) StoreCredentialsForBorderAgentActiveOperationalDataSetCompletion(borderAgentID objc.IObject /* cross-framework: NSData */, activeOperationalDataSet objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("storeCredentialsForBorderAgent:activeOperationalDataSet:completion:"), borderAgentID, activeOperationalDataSet, completion)
} /* debug [instance_methods/method]: StoreCredentialsForBorderAgentActiveOperationalDataSetCompletion */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for THClient */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class THClient */
