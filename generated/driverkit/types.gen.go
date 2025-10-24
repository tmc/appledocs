// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit


// C struct types
// IOAddressSegment - A structure that describes the location and size of a block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOAddressSegment
type IOAddressSegment struct {
}// IOBufferMemoryDescriptor - A memory buffer allocated in the caller’s address space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor
type IOBufferMemoryDescriptor struct {
}// Create - Creates a new memory buffer descriptor object in the current process space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/Create
type Create struct {
}// SetLength - Changes the length of the memory buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/SetLength
type SetLength struct {
}// IODataQueueDispatchSource - A dispatch source that manages a shared-memory data queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource
type IODataQueueDispatchSource struct {
}// DataAvailable - Responds to the addition of new data to the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/DataAvailable
type DataAvailable struct {
}// SetEnableWithCompletion - Controls the enable state of the interrupt source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SetEnableWithCompletion
type SetEnableWithCompletion struct {
}// free - Performs any final cleanup for the data-queue dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/free
type free struct {
}// IODispatchQueue - An object that manages the serial execution of blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue
type IODispatchQueue struct {
}// Create - Creates a new dispatch queue object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/Create
type Create struct {
}// GetName - Returns the name of the queue as a C string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/GetName
type GetName struct {
}// IODispatchSource - The common base class for dispatch sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchSource
type IODispatchSource struct {
}// free - Performs any final cleanup for the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchSource/free
type free struct {
}// init - Handles the basic initialization of the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchSource/init
type init struct {
}// IOHistogramReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter
type IOHistogramReporter struct {
}// initWith
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/initWith
type initWith struct {
}// tallyValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/tallyValue
type tallyValue struct {
}// IOInterruptDispatchSource - A dispatch source that reports hardware-related interrupt events to your driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource
type IOInterruptDispatchSource struct {
}// free - Performs any final cleanup for the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource/free
type free struct {
}// IOMemoryDescriptor - The base class for describing a location in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor
type IOMemoryDescriptor struct {
}// CreateMapping - Maps the contents of the memory block to the address space of the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/CreateMapping
type CreateMapping struct {
}// CreateWithMemoryDescriptors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/CreateWithMemoryDescriptors
type CreateWithMemoryDescriptors struct {
}// free - Performs any final cleanup for the memory descriptor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/free
type free struct {
}// IOMemoryMap - A reference to an existing block of memory in the current process or in a different process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap
type IOMemoryMap struct {
}// GetOffset - Returns the offset from the original start of the memory block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/GetOffset
type GetOffset struct {
}// free - Performs any final cleanup for the memory map object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/free
type free struct {
}// GetAddress - Returns the address of the memory block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/GetAddress
type GetAddress struct {
}// GetLength - Returns the length of the memory block in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/GetLength
type GetLength struct {
}// init - Initializes the memory map object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/init
type init struct {
}// IORPCMessageErrorReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessageErrorReturn
type IORPCMessageErrorReturn struct {
}// IOReportChannelType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportChannelType
type IOReportChannelType struct {
}// IOReportElementValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportElementValues
type IOReportElementValues struct {
}// IOReportLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend
type IOReportLegend struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend/free
type free struct {
}// with
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend/with
type with struct {
}// IOReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter
type IOReporter struct {
}// createLegend - Create a legend entry represending this reporter’s channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/createLegend
type createLegend struct {
}// updateReport - Produce standard reply to IOService::updateReport()
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/updateReport
type updateReport struct {
}// IOService - The base class for managing the setup and registration of your driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService
type IOService struct {
}// CopyProperties - Returns the registry properties associated with the current service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopyProperties
type CopyProperties struct {
}// CreateNameMatchingDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreateNameMatchingDictionary-2nzta
type CreateNameMatchingDictionary struct {
}// GetProvider
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/GetProvider
type GetProvider struct {
}// JoinPMTree
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/JoinPMTree
type JoinPMTree struct {
}// NewUserClient - Requests the creation of a new user client for the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/NewUserClient
type NewUserClient struct {
}// SearchProperty - Searches for a property with the specified name in the current service or one of its parent services, and returns the corresponding value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SearchProperty
type SearchProperty struct {
}// SetDispatchQueue - Associates a custom dispatch queue with the service and assigns the specified name to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetDispatchQueue
type SetDispatchQueue struct {
}// Start - Starts the current service and associates it with the specified provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/Start
type Start struct {
}// init - Handles the basic initialization of the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/init
type init struct {
}// IOServiceNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource
type IOServiceNotificationDispatchSource struct {
}// Cancel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/Cancel
type Cancel struct {
}// DeliverNotifications
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/DeliverNotifications
type DeliverNotifications struct {
}// ServiceNotificationReady
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/ServiceNotificationReady
type ServiceNotificationReady struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/free
type free struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/init
type init struct {
}// IOServiceStateNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource
type IOServiceStateNotificationDispatchSource struct {
}// SetEnableWithCompletion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource/SetEnableWithCompletion
type SetEnableWithCompletion struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource/free
type free struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource/init
type init struct {
}// IOSimpleReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter
type IOSimpleReporter struct {
}// incrementValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/incrementValue
type incrementValue struct {
}// initWith
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/initWith
type initWith struct {
}// with
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/with
type with struct {
}// IOStateReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter
type IOStateReporter struct {
}// getStateLastTransitionTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/getStateLastTransitionTime
type getStateLastTransitionTime struct {
}// setState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setState-1puxp
type setState struct {
}// setStateID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setStateID
type setStateID struct {
}// IOStateReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars
type IOStateReporter_IVars struct {
}// IOStateReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/IOStateReporter_IVars
type IOStateReporter_IVars struct {
}// handleAddChannelSwap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleAddChannelSwap
type handleAddChannelSwap struct {
}// handleIncrementChannelStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleIncrementChannelStateByIndices
type handleIncrementChannelStateByIndices struct {
}// handleOverrideChannelStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleOverrideChannelStateByIndices
type handleOverrideChannelStateByIndices struct {
}// handleSetStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSetStateByIndices
type handleSetStateByIndices struct {
}// handleSetStateID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSetStateID
type handleSetStateID struct {
}// handleSwapCleanup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSwapCleanup
type handleSwapCleanup struct {
}// handleSwapPrepare
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSwapPrepare
type handleSwapPrepare struct {
}// updateChannelValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/updateChannelValues
type updateChannelValues struct {
}// ~IOStateReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/~IOStateReporter_IVars
type ~IOStateReporter_IVars struct {
}// IOTimerDispatchSource - A dispatch source that notifies your driver at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource
type IOTimerDispatchSource struct {
}// Cancel - Cancels all callbacks from the event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/Cancel
type Cancel struct {
}// CheckForWork - Checks for events to handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/CheckForWork
type CheckForWork struct {
}// SetEnableWithCompletion - Enables or disables the timer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/SetEnableWithCompletion
type SetEnableWithCompletion struct {
}// init - Handles the basic initialization of the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/init
type init struct {
}// IOUserClient - A connection to another service that the system manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient
type IOUserClient struct {
}// ExternalMethod - Receive arguments from IOKit.framework IOConnectMethod calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/ExternalMethod
type ExternalMethod struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/free
type free struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/init
type init struct {
}// IOUserServer - A system-managed service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer
type IOUserServer struct {
}// Create
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/Create
type Create struct {
}// Exit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/Exit
type Exit struct {
}// LoadModule
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/LoadModule
type LoadModule struct {
}// RegisterService
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/RegisterService
type RegisterService struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/init
type init struct {
}// OSAction - An object that executes your driver’s custom behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction
type OSAction struct {
}// Aborted - Calls the abort handler of the action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/Aborted
type Aborted struct {
}// Wait
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/Wait
type Wait struct {
}// free - Performs any final cleanup for the action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/free
type free struct {
}// OSArray - A container for an ordered, random-access collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray
type OSArray struct {
}// getCapacity - Returns count of currently allocated capacity for members in array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getCapacity
type getCapacity struct {
}// OSBoolean - A container for a true or false value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBoolean
type OSBoolean struct {
}// release - Releases the OSObject instance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBoolean/release
type release struct {
}// OSBundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle
type OSBundle struct {
}// createFromPath
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle/createFromPath
type createFromPath struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle/init
type init struct {
}// OSCollection - The base class for DriverKit collection objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollection
type OSCollection struct {
}// OSContainer - The base class for DriverKit data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSContainer
type OSContainer struct {
}// OSData - A container for untyped data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData
type OSData struct {
}// isEqualTo - Compares the data with an OSObject
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/isEqualTo-4xz0j
type isEqualTo struct {
}// withData - Allocates an OSData object with a copy of bytes from a subset of another OSData.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/withData-4rd8n
type withData struct {
}// OSDictionary - A container for a collection with elements that are key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionary
type OSDictionary struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionary/free
type free struct {
}// setObject - Add or replace an object in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionary/setObject-9b4z0
type setObject struct {
}// OSMappedFile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile
type OSMappedFile struct {
}// createFromPath
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/createFromPath
type createFromPath struct {
}// data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/data
type data struct {
}// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/free
type free struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/init
type init struct {
}// size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/size
type size struct {
}// OSNumber - A container for an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber
type OSNumber struct {
}// isEqualTo - Compares the string with an OSObject
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/isEqualTo-333kh
type isEqualTo struct {
}// numberOfBits - Returns the number of bits the OSNumber was created with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/numberOfBits
type numberOfBits struct {
}// unsigned8BitValue - Returns the value of the OSNumber as a uint8_t value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/unsigned8BitValue
type unsigned8BitValue struct {
}// OSObject - The base class for DriverKit objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObject
type OSObject struct {
}// CopyDispatchQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObject/CopyDispatchQueue
type CopyDispatchQueue struct {
}// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObject/init
type init struct {
}// OSSerialization - A container for one or more objects, serialized in a binary data format that is suitable for messaging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization
type OSSerialization struct {
}// finalizeBuffer - Obtain the result of the serialization performed by createFromObject().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/finalizeBuffer
type finalizeBuffer struct {
}// freeBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/freeBuffer
type freeBuffer struct {
}// OSString - A container for managing an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString
type OSString struct {
}// withCStringNoCopy - Allocates an OSString object with a copy of a c-string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/withCStringNoCopy
type withCStringNoCopy struct {
}// withString - Allocates an OSString object with a copy of an OString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/withString
type withString struct {
}// SCSICmd_INQUIRY_Page83_Header_SPC_16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_Header_SPC_16
type SCSICmd_INQUIRY_Page83_Header_SPC_16 struct {
}// mach_msg_port_descriptor_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_port_descriptor_t
type mach_msg_port_descriptor_t struct {
}// queue_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/queue_entry
type queue_entry struct {
}



