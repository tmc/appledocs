// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NetService] class.
var (
	NetServiceClass     _NetServiceClass
	NetServiceClassOnce sync.Once
)

func getNetServiceClass() _NetServiceClass {
	NetServiceClassOnce.Do(func() {
		NetServiceClass = _NetServiceClass{objc.GetClass("NSNetService")}
	})
	return NetServiceClass
}

type _NetServiceClass struct {
	class objc.Class
}





// An interface definition for the [NetService] class.
type INetService interface {
	objectivec.IObject
	

	// properties:
	Addresses() []Data
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Domain() IString
	HostName() IString
	IncludesPeerToPeer() bool
	SetIncludesPeerToPeer(value bool)
	Name() IString
	Port() int
	Type() IString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NetServiceClass) Alloc() NetService {
	rv := objc.Send[NetService](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NetServiceClass) New() NetService {
	rv := objc.Send[NetService](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NetService) Init() NetService {
	rv := objc.Send[NetService](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NetService) Autorelease() NetService {
	rv := objc.Send[NetService](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetService creates a new NetService instance.
func NewNetService() NetService {
	return getNetServiceClass().New()
}





// A network service that broadcasts its availability using multicast DNS.
//
// The class represents a network service, either one your application publishes or is a client of. This class and the class use multicast DNS to convey information about network services to and from your application. The API of provides a convenient way to publish the services offered by your application and to resolve the socket address for a service. The types of services you access using are the same types that you access directly using BSD sockets. HTTP and FTP are two services commonly provided by systems. (For a list of common services and the ports used by those services, see the file .) Applications can also define their own custom services to provide specific data to clients. You can use the class as either a publisher of a service or a client of a service. If your application publishes a service, your code must acquire a port and prepare a socket to communicate with clients. Once your socket is ready, you use the class to notify clients that your service is ready. If your application is the client of a network service, you can either create an object directly (if you know the exact host and port information) or use an object to browse for services. To publish a service, initialize your object with the service name, domain, type, and port information. All of this information must be valid for the socket created by your application. Once initialized, call the method to broadcast your service information to the network. When connecting to a service, use the class to locate the service on the network and obtain the corresponding object. Once you have the object, call the method to verify that the service is available and ready for your application. If it is, the property provides the socket information you can use to connect to the service. The methods of operate asynchronously so your application is not impacted by the speed of the network. All information about a service is returned to your application through the object’s delegate. You must provide a delegate object to respond to messages and to handle errors appropriately.


// A network service that broadcasts its availability using multicast DNS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService
type NetService struct {
	objectivec.Object
}

// NetServiceFrom constructs a [NetService] from an unsafe.Pointer.
//
// A network service that broadcasts its availability using multicast DNS.
func NetServiceFrom(ptr unsafe.Pointer) NetService {
	return NetService{objectivec.Object{objc.ID(ptr)}}
}






// Returns the receiver, initialized as a network service of a given type and sets the initial host information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:)
func NewNetServiceWithDomainTypeName(domain IString, type_ IString, name IString) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[NetService](instance.ID, objc.Sel("initWithDomain:type:name:"), domain, type_, name)
	rv.Autorelease()
	return rv
}


// Initializes the receiver for publishing a network service of type at the socket location specified by , , and .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:port:)
func NewNetServiceWithDomainTypeNamePort(domain IString, type_ IString, name IString, port int) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[NetService](instance.ID, objc.Sel("initWithDomain:type:name:port:"), domain, type_, name, port)
	rv.Autorelease()
	return rv
}







// Returns an object representing a TXT record formed from a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/data(fromTXTRecord:)
func (nc _NetServiceClass) DataFromTXTRecordDictionary(txtDictionary IDictionary) IData {
	rv := objc.Send[Data](objc.ID(nc.class), objc.Sel("dataFromTXTRecordDictionary:"), txtDictionary)
	return rv
}


// Returns a dictionary representing a TXT record given as an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/dictionary(fromTXTRecord:)
func (nc _NetServiceClass) DictionaryFromTXTRecordData(txtData IData) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(nc.class), objc.Sel("dictionaryFromTXTRecordData:"), txtData)
	return rv
}

















// A read-only array containing objects, each of which contains a socket address for the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/addresses
func (n_ NetService) Addresses() []Data {
	rv := objc.Send[[]Data](n_.ID, objc.Sel("addresses"))
	return rv
}


// The delegate for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/delegate
func (n_ NetService) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/delegate
func (n_ NetService) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}


// A string containing the domain for this service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/domain
func (n_ NetService) Domain() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("domain"))
	return rv
}


// A string containing the DNS hostname for this service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/hostName
func (n_ NetService) HostName() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("hostName"))
	return rv
}


// Specifies whether to also publish, resolve, or monitor this service over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/includesPeerToPeer
func (n_ NetService) IncludesPeerToPeer() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("includesPeerToPeer"))
	return rv
}


// Specifies whether to also publish, resolve, or monitor this service over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/includesPeerToPeer
func (n_ NetService) SetIncludesPeerToPeer(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludesPeerToPeer:"), value)
}


// A string containing the name of this service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/name
func (n_ NetService) Name() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("name"))
	return rv
}


// The port on which the service is listening for connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/port
func (n_ NetService) Port() int {
	rv := objc.Send[int](n_.ID, objc.Sel("port"))
	return rv
}


// The type of the published service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/type
func (n_ NetService) Type() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("type"))
	return rv
}







