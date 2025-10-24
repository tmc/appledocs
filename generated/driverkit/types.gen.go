// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit
import (
	"unsafe"
)


// C struct types
// IOBufferMemoryDescriptor - A memory buffer allocated in the caller’s address space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor
type IOBufferMemoryDescriptor struct {
}/* debug [types.gen.go/struct]: IOBufferMemoryDescriptor */

// Create - Creates a new memory buffer descriptor object in the current process space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/Create
type Create struct {
}/* debug [types.gen.go/struct]: Create */

// free - Performs any final cleanup for the memory buffer descriptor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/free
type free struct {
}/* debug [types.gen.go/struct]: free */

// GetAddressRange - Returns the address and length of the memory buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/GetAddressRange
type GetAddressRange struct {
}/* debug [types.gen.go/struct]: GetAddressRange */

// init - Initializes the buffer memory descriptor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/init
type init struct {
}/* debug [types.gen.go/struct]: init */

// SetLength - Changes the length of the memory buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOBufferMemoryDescriptor/SetLength
type SetLength struct {
}/* debug [types.gen.go/struct]: SetLength */

// IODataQueueDispatchSource - A dispatch source that manages a shared-memory data queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource
type IODataQueueDispatchSource struct {
}/* debug [types.gen.go/struct]: IODataQueueDispatchSource */

// Cancel - Cancels all callbacks from the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/Cancel
type Cancel struct {
}/* debug [types.gen.go/struct]: Cancel */

// CanEnqueueData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/CanEnqueueData-3yyz2
type CanEnqueueData struct {
}/* debug [types.gen.go/struct]: CanEnqueueData */

// CheckForWork - Checks for events to handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/CheckForWork
type CheckForWork struct {
}/* debug [types.gen.go/struct]: CheckForWork */

// CopyDataAvailableHandler - A private method that the dispatch source uses to detect enqueued data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/CopyDataAvailableHandler
type CopyDataAvailableHandler struct {
}/* debug [types.gen.go/struct]: CopyDataAvailableHandler */

// CopyDataServicedHandler - A private method that the dispatch source uses to detect dequeued data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/CopyDataServicedHandler
type CopyDataServicedHandler struct {
}/* debug [types.gen.go/struct]: CopyDataServicedHandler */

// CopyMemory - A private method that the dispatch source uses to copy memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/CopyMemory
type CopyMemory struct {
}/* debug [types.gen.go/struct]: CopyMemory */

// DataAvailable - Responds to the addition of new data to the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/DataAvailable
type DataAvailable struct {
}/* debug [types.gen.go/struct]: DataAvailable */

// DataServiced - Responds to the removal of data from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/DataServiced
type DataServiced struct {
}/* debug [types.gen.go/struct]: DataServiced */

// Dequeue - Removes the next entry from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/Dequeue
type Dequeue struct {
}/* debug [types.gen.go/struct]: Dequeue */

// DequeueWithCoalesce - Removes the next queue entry, but doesn’t automatically send notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/DequeueWithCoalesce
type DequeueWithCoalesce struct {
}/* debug [types.gen.go/struct]: DequeueWithCoalesce */

// Enqueue - Adds a single entry to the shared data queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/Enqueue
type Enqueue struct {
}/* debug [types.gen.go/struct]: Enqueue */

// EnqueueWithCoalesce - Adds an entry to the queue, but doesn’t automatically send a data-available notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/EnqueueWithCoalesce
type EnqueueWithCoalesce struct {
}/* debug [types.gen.go/struct]: EnqueueWithCoalesce */

// GetDataQueueEntryHeaderSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/GetDataQueueEntryHeaderSize
type GetDataQueueEntryHeaderSize struct {
}/* debug [types.gen.go/struct]: GetDataQueueEntryHeaderSize */

// IsDataAvailable - Checks whether the data queue contains data to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/IsDataAvailable
type IsDataAvailable struct {
}/* debug [types.gen.go/struct]: IsDataAvailable */

// Peek - Returns the next queue entry without removing it from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/Peek
type Peek struct {
}/* debug [types.gen.go/struct]: Peek */

// SendDataAvailable - Sends a notification to observers that indicates more data is available for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SendDataAvailable
type SendDataAvailable struct {
}/* debug [types.gen.go/struct]: SendDataAvailable */

// SendDataServiced - Notifies interested parties that you removed data from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SendDataServiced
type SendDataServiced struct {
}/* debug [types.gen.go/struct]: SendDataServiced */

// SetDataAvailableHandler - Sets the handler block to run when another object adds data to the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SetDataAvailableHandler
type SetDataAvailableHandler struct {
}/* debug [types.gen.go/struct]: SetDataAvailableHandler */

// SetDataServicedHandler - Installs the handler block to execute when data is removed from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SetDataServicedHandler
type SetDataServicedHandler struct {
}/* debug [types.gen.go/struct]: SetDataServicedHandler */

// SetEnableWithCompletion - Controls the enable state of the interrupt source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODataQueueDispatchSource/SetEnableWithCompletion
type SetEnableWithCompletion struct {
}/* debug [types.gen.go/struct]: SetEnableWithCompletion */

// IODispatchQueue - An object that manages the serial execution of blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue
type IODispatchQueue struct {
}/* debug [types.gen.go/struct]: IODispatchQueue */

// DispatchAsync - Schedule a block for asynchronous execution on the current queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchAsync
type DispatchAsync struct {
}/* debug [types.gen.go/struct]: DispatchAsync */

// DispatchAsync_f - Schedule a C-style function for asynchrous execution on the current queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchAsync_f
type DispatchAsync_f struct {
}/* debug [types.gen.go/struct]: DispatchAsync_f */

// DispatchConcurrent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchConcurrent
type DispatchConcurrent struct {
}/* debug [types.gen.go/struct]: DispatchConcurrent */

// DispatchConcurrent_f
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchConcurrent_f
type DispatchConcurrent_f struct {
}/* debug [types.gen.go/struct]: DispatchConcurrent_f */

// DispatchSync - Schedule a block for synchronous execution on the current queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchSync
type DispatchSync struct {
}/* debug [types.gen.go/struct]: DispatchSync */

// DispatchSync_f - Schedule a C-style function for synchronous execution on the current queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/DispatchSync_f
type DispatchSync_f struct {
}/* debug [types.gen.go/struct]: DispatchSync_f */

// GetName - Returns the name of the queue as a C string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/GetName
type GetName struct {
}/* debug [types.gen.go/struct]: GetName */

// Log - Log the current execution context with respect to any queues the current thread holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/Log
type Log struct {
}/* debug [types.gen.go/struct]: Log */

// OnQueue - Returns a Boolean value that indicates whether the current thread matches the dispatch queue’s thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/OnQueue
type OnQueue struct {
}/* debug [types.gen.go/struct]: OnQueue */

// RunAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/RunAction
type RunAction struct {
}/* debug [types.gen.go/struct]: RunAction */

// Sleep
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/Sleep
type Sleep struct {
}/* debug [types.gen.go/struct]: Sleep */

// SleepWithDeadline
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/SleepWithDeadline
type SleepWithDeadline struct {
}/* debug [types.gen.go/struct]: SleepWithDeadline */

// SleepWithTimeout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/SleepWithTimeout
type SleepWithTimeout struct {
}/* debug [types.gen.go/struct]: SleepWithTimeout */

// Wakeup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/Wakeup
type Wakeup struct {
}/* debug [types.gen.go/struct]: Wakeup */

// WakeupWithOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchQueue/WakeupWithOptions
type WakeupWithOptions struct {
}/* debug [types.gen.go/struct]: WakeupWithOptions */

// IODispatchSource - The common base class for dispatch sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchSource
type IODispatchSource struct {
}/* debug [types.gen.go/struct]: IODispatchSource */

// SetEnable - Enables or disables the delivery of events to your code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODispatchSource/SetEnable
type SetEnable struct {
}/* debug [types.gen.go/struct]: SetEnable */

// IOHistogramReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter
type IOHistogramReporter struct {
}/* debug [types.gen.go/struct]: IOHistogramReporter */

// addChannel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/addChannel
type addChannel struct {
}/* debug [types.gen.go/struct]: addChannel */

// initWith
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/initWith
type initWith struct {
}/* debug [types.gen.go/struct]: initWith */

// overrideBucketValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/overrideBucketValues
type overrideBucketValues struct {
}/* debug [types.gen.go/struct]: overrideBucketValues */

// tallyValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/tallyValue
type tallyValue struct {
}/* debug [types.gen.go/struct]: tallyValue */

// with
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter/with
type with struct {
}/* debug [types.gen.go/struct]: with */

// IOInterruptDispatchSource - A dispatch source that reports hardware-related interrupt events to your driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource
type IOInterruptDispatchSource struct {
}/* debug [types.gen.go/struct]: IOInterruptDispatchSource */

// GetInterruptType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource/GetInterruptType
type GetInterruptType struct {
}/* debug [types.gen.go/struct]: GetInterruptType */

// GetLastInterrupt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource/GetLastInterrupt
type GetLastInterrupt struct {
}/* debug [types.gen.go/struct]: GetLastInterrupt */

// InterruptOccurred - Executes custom code when an interrupt occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource/InterruptOccurred
type InterruptOccurred struct {
}/* debug [types.gen.go/struct]: InterruptOccurred */

// SetHandler - Set the handler block to run when the interrupt fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSource/SetHandler
type SetHandler struct {
}/* debug [types.gen.go/struct]: SetHandler */

// IOMemoryDescriptor - The base class for describing a location in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor
type IOMemoryDescriptor struct {
}/* debug [types.gen.go/struct]: IOMemoryDescriptor */

// CreateMapping - Maps the contents of the memory block to the address space of the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/CreateMapping
type CreateMapping struct {
}/* debug [types.gen.go/struct]: CreateMapping */

// CreateSubMemoryDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/CreateSubMemoryDescriptor
type CreateSubMemoryDescriptor struct {
}/* debug [types.gen.go/struct]: CreateSubMemoryDescriptor */

// CreateWithMemoryDescriptors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/CreateWithMemoryDescriptors
type CreateWithMemoryDescriptors struct {
}/* debug [types.gen.go/struct]: CreateWithMemoryDescriptors */

// GetLength - Returns the length of the memory block represented by this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/GetLength
type GetLength struct {
}/* debug [types.gen.go/struct]: GetLength */

// Map - Maps memory internally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryDescriptor/Map
type Map struct {
}/* debug [types.gen.go/struct]: Map */

// IOMemoryMap - A reference to an existing block of memory in the current process or in a different process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap
type IOMemoryMap struct {
}/* debug [types.gen.go/struct]: IOMemoryMap */

// GetAddress - Returns the address of the memory block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/GetAddress
type GetAddress struct {
}/* debug [types.gen.go/struct]: GetAddress */

// GetOffset - Returns the offset from the original start of the memory block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMemoryMap/GetOffset
type GetOffset struct {
}/* debug [types.gen.go/struct]: GetOffset */

// IOReportLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend
type IOReportLegend struct {
}/* debug [types.gen.go/struct]: IOReportLegend */

// addLegendEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend/addLegendEntry
type addLegendEntry struct {
}/* debug [types.gen.go/struct]: addLegendEntry */

// addReporterLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend/addReporterLegend-c.method
type addReporterLegend struct {
}/* debug [types.gen.go/struct]: addReporterLegend */

// getLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportLegend/getLegend
type getLegend struct {
}/* debug [types.gen.go/struct]: getLegend */

// IOReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter
type IOReporter struct {
}/* debug [types.gen.go/struct]: IOReporter */

// configureAllReports - Calls 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/configureAllReports
type configureAllReports struct {
}/* debug [types.gen.go/struct]: configureAllReports */

// configureReport - Track IOService::configureReport(), provide sizing info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/configureReport
type configureReport struct {
}/* debug [types.gen.go/struct]: configureReport */

// createLegend - Create a legend entry represending this reporter’s channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/createLegend
type createLegend struct {
}/* debug [types.gen.go/struct]: createLegend */

// updateAllReports - Calls 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/updateAllReports
type updateAllReports struct {
}/* debug [types.gen.go/struct]: updateAllReports */

// updateReport - Produce standard reply to IOService::updateReport()
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter/updateReport
type updateReport struct {
}/* debug [types.gen.go/struct]: updateReport */

// IOService - The base class for managing the setup and registration of your driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService
type IOService struct {
}/* debug [types.gen.go/struct]: IOService */

// AdjustBusy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/AdjustBusy
type AdjustBusy struct {
}/* debug [types.gen.go/struct]: AdjustBusy */

// ChangePowerState - Changes the device’s power state to the specified level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/ChangePowerState
type ChangePowerState struct {
}/* debug [types.gen.go/struct]: ChangePowerState */

// ClientCrashed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/ClientCrashed
type ClientCrashed struct {
}/* debug [types.gen.go/struct]: ClientCrashed */

// ConfigureReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/ConfigureReport
type ConfigureReport struct {
}/* debug [types.gen.go/struct]: ConfigureReport */

// CopyDispatchQueue - Gets the dispatch queue with the specified name from the current service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopyDispatchQueue
type CopyDispatchQueue struct {
}/* debug [types.gen.go/struct]: CopyDispatchQueue */

// CopyName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopyName
type CopyName struct {
}/* debug [types.gen.go/struct]: CopyName */

// CopyProperties - Returns the registry properties associated with the current service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopyProperties
type CopyProperties struct {
}/* debug [types.gen.go/struct]: CopyProperties */

// CopyProviderProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopyProviderProperties
type CopyProviderProperties struct {
}/* debug [types.gen.go/struct]: CopyProviderProperties */

// CopySystemStateNotificationService
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CopySystemStateNotificationService
type CopySystemStateNotificationService struct {
}/* debug [types.gen.go/struct]: CopySystemStateNotificationService */

// CoreAnalyticsSendEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CoreAnalyticsSendEvent
type CoreAnalyticsSendEvent struct {
}/* debug [types.gen.go/struct]: CoreAnalyticsSendEvent */

// CreateDefaultDispatchQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreateDefaultDispatchQueue
type CreateDefaultDispatchQueue struct {
}/* debug [types.gen.go/struct]: CreateDefaultDispatchQueue */

// CreateKernelClassMatchingDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreateKernelClassMatchingDictionary-3uqly
type CreateKernelClassMatchingDictionary struct {
}/* debug [types.gen.go/struct]: CreateKernelClassMatchingDictionary */

// CreateNameMatchingDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreateNameMatchingDictionary-206ej
type CreateNameMatchingDictionary struct {
}/* debug [types.gen.go/struct]: CreateNameMatchingDictionary */

// CreatePMAssertion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreatePMAssertion
type CreatePMAssertion struct {
}/* debug [types.gen.go/struct]: CreatePMAssertion */

// CreatePropertyMatchingDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreatePropertyMatchingDictionary-4tuca
type CreatePropertyMatchingDictionary struct {
}/* debug [types.gen.go/struct]: CreatePropertyMatchingDictionary */

// CreateUserClassMatchingDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/CreateUserClassMatchingDictionary-4gpbj
type CreateUserClassMatchingDictionary struct {
}/* debug [types.gen.go/struct]: CreateUserClassMatchingDictionary */

// GetBusyState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/GetBusyState
type GetBusyState struct {
}/* debug [types.gen.go/struct]: GetBusyState */

// GetProvider
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/GetProvider
type GetProvider struct {
}/* debug [types.gen.go/struct]: GetProvider */

// GetRegistryEntryID - Returns the registry ID for the current service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/GetRegistryEntryID
type GetRegistryEntryID struct {
}/* debug [types.gen.go/struct]: GetRegistryEntryID */

// JoinPMTree
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/JoinPMTree
type JoinPMTree struct {
}/* debug [types.gen.go/struct]: JoinPMTree */

// NewUserClient - Requests the creation of a new user client for the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/NewUserClient
type NewUserClient struct {
}/* debug [types.gen.go/struct]: NewUserClient */

// RegisterService - Starts the registration process for the service and performs any additional matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/RegisterService
type RegisterService struct {
}/* debug [types.gen.go/struct]: RegisterService */

// ReleasePMAssertion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/ReleasePMAssertion
type ReleasePMAssertion struct {
}/* debug [types.gen.go/struct]: ReleasePMAssertion */

// RemoveProperty
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/RemoveProperty
type RemoveProperty struct {
}/* debug [types.gen.go/struct]: RemoveProperty */

// RequireMaxBusStall
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/RequireMaxBusStall
type RequireMaxBusStall struct {
}/* debug [types.gen.go/struct]: RequireMaxBusStall */

// SearchProperty - Searches for a property with the specified name in the current service or one of its parent services, and returns the corresponding value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SearchProperty
type SearchProperty struct {
}/* debug [types.gen.go/struct]: SearchProperty */

// SetDispatchQueue - Associates a custom dispatch queue with the service and assigns the specified name to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetDispatchQueue
type SetDispatchQueue struct {
}/* debug [types.gen.go/struct]: SetDispatchQueue */

// SetLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetLegend
type SetLegend struct {
}/* debug [types.gen.go/struct]: SetLegend */

// SetName - Sets the name of the service in the system’s registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetName
type SetName struct {
}/* debug [types.gen.go/struct]: SetName */

// SetPowerOverride
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetPowerOverride
type SetPowerOverride struct {
}/* debug [types.gen.go/struct]: SetPowerOverride */

// SetPowerState - Updates the service in response to power-related changes for a provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetPowerState
type SetPowerState struct {
}/* debug [types.gen.go/struct]: SetPowerState */

// SetProperties - Sends the dictionary of properties to the current service object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/SetProperties
type SetProperties struct {
}/* debug [types.gen.go/struct]: SetProperties */

// Start - Starts the current service and associates it with the specified provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/Start
type Start struct {
}/* debug [types.gen.go/struct]: Start */

// StateNotificationItemCopy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/StateNotificationItemCopy
type StateNotificationItemCopy struct {
}/* debug [types.gen.go/struct]: StateNotificationItemCopy */

// StateNotificationItemCreate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/StateNotificationItemCreate
type StateNotificationItemCreate struct {
}/* debug [types.gen.go/struct]: StateNotificationItemCreate */

// StateNotificationItemSet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/StateNotificationItemSet
type StateNotificationItemSet struct {
}/* debug [types.gen.go/struct]: StateNotificationItemSet */

// Stop - Stops the service associated with the specified provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/Stop
type Stop struct {
}/* debug [types.gen.go/struct]: Stop */

// Stop_async
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/Stop_async
type Stop_async struct {
}/* debug [types.gen.go/struct]: Stop_async */

// StringFromReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/StringFromReturn
type StringFromReturn struct {
}/* debug [types.gen.go/struct]: StringFromReturn */

// Terminate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/Terminate
type Terminate struct {
}/* debug [types.gen.go/struct]: Terminate */

// UpdateReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOService/UpdateReport
type UpdateReport struct {
}/* debug [types.gen.go/struct]: UpdateReport */

// IOServiceStateNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource
type IOServiceStateNotificationDispatchSource struct {
}/* debug [types.gen.go/struct]: IOServiceStateNotificationDispatchSource */

// StateNotificationBegin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource/StateNotificationBegin
type StateNotificationBegin struct {
}/* debug [types.gen.go/struct]: StateNotificationBegin */

// StateNotificationReady
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource/StateNotificationReady
type StateNotificationReady struct {
}/* debug [types.gen.go/struct]: StateNotificationReady */

// IOSimpleReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter
type IOSimpleReporter struct {
}/* debug [types.gen.go/struct]: IOSimpleReporter */

// getValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/getValue
type getValue struct {
}/* debug [types.gen.go/struct]: getValue */

// incrementValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/incrementValue
type incrementValue struct {
}/* debug [types.gen.go/struct]: incrementValue */

// setValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter/setValue
type setValue struct {
}/* debug [types.gen.go/struct]: setValue */

// IOStateReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter
type IOStateReporter struct {
}/* debug [types.gen.go/struct]: IOStateReporter */

// getStateInTransitions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/getStateInTransitions
type getStateInTransitions struct {
}/* debug [types.gen.go/struct]: getStateInTransitions */

// getStateLastChannelUpdateTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/getStateLastChannelUpdateTime
type getStateLastChannelUpdateTime struct {
}/* debug [types.gen.go/struct]: getStateLastChannelUpdateTime */

// getStateLastTransitionTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/getStateLastTransitionTime
type getStateLastTransitionTime struct {
}/* debug [types.gen.go/struct]: getStateLastTransitionTime */

// getStateResidencyTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/getStateResidencyTime
type getStateResidencyTime struct {
}/* debug [types.gen.go/struct]: getStateResidencyTime */

// incrementChannelState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/incrementChannelState
type incrementChannelState struct {
}/* debug [types.gen.go/struct]: incrementChannelState */

// overrideChannelState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/overrideChannelState
type overrideChannelState struct {
}/* debug [types.gen.go/struct]: overrideChannelState */

// setChannelState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setChannelState-7n3or
type setChannelState struct {
}/* debug [types.gen.go/struct]: setChannelState */

// setState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setState-1puxp
type setState struct {
}/* debug [types.gen.go/struct]: setState */

// setStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setStateByIndices-13fxh
type setStateByIndices struct {
}/* debug [types.gen.go/struct]: setStateByIndices */

// setStateID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter/setStateID
type setStateID struct {
}/* debug [types.gen.go/struct]: setStateID */

// IOTimerDispatchSource - A dispatch source that notifies your driver at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource
type IOTimerDispatchSource struct {
}/* debug [types.gen.go/struct]: IOTimerDispatchSource */

// TimerOccurred - Executes custom code when the timer fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/TimerOccurred
type TimerOccurred struct {
}/* debug [types.gen.go/struct]: TimerOccurred */

// WakeAtTime - Schedules a callback from the timer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOTimerDispatchSource/WakeAtTime
type WakeAtTime struct {
}/* debug [types.gen.go/struct]: WakeAtTime */

// IOUserClient - A connection to another service that the system manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient
type IOUserClient struct {
}/* debug [types.gen.go/struct]: IOUserClient */

// AsyncCompletion - Send asynchronous arguments to a completion supplied by ExternalMethod().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/AsyncCompletion
type AsyncCompletion struct {
}/* debug [types.gen.go/struct]: AsyncCompletion */

// CopyClientEntitlements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/CopyClientEntitlements
type CopyClientEntitlements struct {
}/* debug [types.gen.go/struct]: CopyClientEntitlements */

// CopyClientMemoryForType - Return an IOMemoryDescriptor to be mapped into the client task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/CopyClientMemoryForType
type CopyClientMemoryForType struct {
}/* debug [types.gen.go/struct]: CopyClientMemoryForType */

// CreateMemoryDescriptorFromClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/CreateMemoryDescriptorFromClient
type CreateMemoryDescriptorFromClient struct {
}/* debug [types.gen.go/struct]: CreateMemoryDescriptorFromClient */

// ExternalMethod - Receive arguments from IOKit.framework IOConnectMethod calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/ExternalMethod
type ExternalMethod struct {
}/* debug [types.gen.go/struct]: ExternalMethod */

// KernelCompletion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClient/KernelCompletion
type KernelCompletion struct {
}/* debug [types.gen.go/struct]: KernelCompletion */

// IOUserServer - A system-managed service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer
type IOUserServer struct {
}/* debug [types.gen.go/struct]: IOUserServer */

// Exit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/Exit
type Exit struct {
}/* debug [types.gen.go/struct]: Exit */

// LoadModule
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserServer/LoadModule
type LoadModule struct {
}/* debug [types.gen.go/struct]: LoadModule */

// OSAction - An object that executes your driver’s custom behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction
type OSAction struct {
}/* debug [types.gen.go/struct]: OSAction */

// Aborted - Calls the abort handler of the action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/Aborted
type Aborted struct {
}/* debug [types.gen.go/struct]: Aborted */

// EndWait
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/EndWait
type EndWait struct {
}/* debug [types.gen.go/struct]: EndWait */

// GetReference - Returns a pointer to any additional memory allocated by the action object on your behalf.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/GetReference
type GetReference struct {
}/* debug [types.gen.go/struct]: GetReference */

// SetAbortedHandler - Install a handler for the system to call when no other processes reference the action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/SetAbortedHandler
type SetAbortedHandler struct {
}/* debug [types.gen.go/struct]: SetAbortedHandler */

// Wait
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/Wait
type Wait struct {
}/* debug [types.gen.go/struct]: Wait */

// WillWait
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSAction/WillWait
type WillWait struct {
}/* debug [types.gen.go/struct]: WillWait */

// OSArray - A container for an ordered, random-access collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray
type OSArray struct {
}/* debug [types.gen.go/struct]: OSArray */

// ensureCapacity - Allocates capacity for members in array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/ensureCapacity
type ensureCapacity struct {
}/* debug [types.gen.go/struct]: ensureCapacity */

// flushCollection - Removes and drops references to all members of array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/flushCollection
type flushCollection struct {
}/* debug [types.gen.go/struct]: flushCollection */

// getCapacity - Returns count of currently allocated capacity for members in array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getCapacity
type getCapacity struct {
}/* debug [types.gen.go/struct]: getCapacity */

// getCount - Returns count of members in array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getCount
type getCount struct {
}/* debug [types.gen.go/struct]: getCount */

// getLastObject - Returns the last member of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getLastObject
type getLastObject struct {
}/* debug [types.gen.go/struct]: getLastObject */

// getNextIndexOfObject - Searches the array for an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getNextIndexOfObject
type getNextIndexOfObject struct {
}/* debug [types.gen.go/struct]: getNextIndexOfObject */

// getObject - Returns a member of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/getObject
type getObject struct {
}/* debug [types.gen.go/struct]: getObject */

// isEqualTo - Compares all members of two arrays with isEqualTo().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/isEqualTo-5w7om
type isEqualTo struct {
}/* debug [types.gen.go/struct]: isEqualTo */

// iterateObjects - Iterates the array calling a callback block for each member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/iterateObjects
type iterateObjects struct {
}/* debug [types.gen.go/struct]: iterateObjects */

// merge - Appends all members of an array to this array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/merge
type merge struct {
}/* debug [types.gen.go/struct]: merge */

// removeObject - Removes a current member of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/removeObject
type removeObject struct {
}/* debug [types.gen.go/struct]: removeObject */

// replaceObject - Removes a current member of the array and replaces it with another object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/replaceObject
type replaceObject struct {
}/* debug [types.gen.go/struct]: replaceObject */

// setObject - Appends an object as the last member of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/setObject-3bore
type setObject struct {
}/* debug [types.gen.go/struct]: setObject */

// withArray - Allocates an OSArray object with given members and preallocated capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/withArray
type withArray struct {
}/* debug [types.gen.go/struct]: withArray */

// withCapacity - Allocates an OSArray object with preallocated capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/withCapacity
type withCapacity struct {
}/* debug [types.gen.go/struct]: withCapacity */

// withObjects - Allocates an OSArray object with given members and preallocated capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArray/withObjects
type withObjects struct {
}/* debug [types.gen.go/struct]: withObjects */

// OSBoolean - A container for a true or false value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBoolean
type OSBoolean struct {
}/* debug [types.gen.go/struct]: OSBoolean */

// release - Releases the OSObject instance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBoolean/release
type release struct {
}/* debug [types.gen.go/struct]: release */

// retain - Retains the OSObject instance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBoolean/retain
type retain struct {
}/* debug [types.gen.go/struct]: retain */

// OSBundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle
type OSBundle struct {
}/* debug [types.gen.go/struct]: OSBundle */

// createFromPath
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle/createFromPath
type createFromPath struct {
}/* debug [types.gen.go/struct]: createFromPath */

// loadResource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle/loadResource
type loadResource struct {
}/* debug [types.gen.go/struct]: loadResource */

// mainBundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSBundle/mainBundle
type mainBundle struct {
}/* debug [types.gen.go/struct]: mainBundle */

// OSCollection - The base class for DriverKit collection objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollection
type OSCollection struct {
}/* debug [types.gen.go/struct]: OSCollection */

// copyCollection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollection/copyCollection
type copyCollection struct {
}/* debug [types.gen.go/struct]: copyCollection */

// OSContainer - The base class for DriverKit data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSContainer
type OSContainer struct {
}/* debug [types.gen.go/struct]: OSContainer */

// OSData - A container for untyped data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData
type OSData struct {
}/* debug [types.gen.go/struct]: OSData */

// appendBytes - Appends a buffer of bytes to the OSData object’s internal data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/appendBytes-38hs5
type appendBytes struct {
}/* debug [types.gen.go/struct]: appendBytes */

// getBytesNoCopy - Returns a pointer to the OSData object’s internal data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/getBytesNoCopy-53puz
type getBytesNoCopy struct {
}/* debug [types.gen.go/struct]: getBytesNoCopy */

// getLength - Returns length of data present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/getLength
type getLength struct {
}/* debug [types.gen.go/struct]: getLength */

// withBytes - Allocates an OSData object with a copy of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/withBytes
type withBytes struct {
}/* debug [types.gen.go/struct]: withBytes */

// withBytesNoCopy - Allocates an OSData object with a copy of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/withBytesNoCopy
type withBytesNoCopy struct {
}/* debug [types.gen.go/struct]: withBytesNoCopy */

// withData - Allocates an OSData object with a copy of bytes from a subset of another OSData.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSData/withData-4rd8n
type withData struct {
}/* debug [types.gen.go/struct]: withData */

// OSDictionary - A container for a collection with elements that are key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionary
type OSDictionary struct {
}/* debug [types.gen.go/struct]: OSDictionary */

// withDictionary - Allocates an OSDictionary object with given members and preallocated capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionary/withDictionary
type withDictionary struct {
}/* debug [types.gen.go/struct]: withDictionary */

// OSMappedFile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile
type OSMappedFile struct {
}/* debug [types.gen.go/struct]: OSMappedFile */

// data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/data
type data struct {
}/* debug [types.gen.go/struct]: data */

// size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMappedFile/size
type size struct {
}/* debug [types.gen.go/struct]: size */

// OSNumber - A container for an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber
type OSNumber struct {
}/* debug [types.gen.go/struct]: OSNumber */

// numberOfBits - Returns the number of bits the OSNumber was created with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/numberOfBits
type numberOfBits struct {
}/* debug [types.gen.go/struct]: numberOfBits */

// unsigned16BitValue - Returns the value of the OSNumber as a uint16_t value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/unsigned16BitValue
type unsigned16BitValue struct {
}/* debug [types.gen.go/struct]: unsigned16BitValue */

// unsigned32BitValue - Returns the value of the OSNumber as a uint32_t value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/unsigned32BitValue
type unsigned32BitValue struct {
}/* debug [types.gen.go/struct]: unsigned32BitValue */

// unsigned64BitValue - Returns the value of the OSNumber as a uint64_t value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/unsigned64BitValue
type unsigned64BitValue struct {
}/* debug [types.gen.go/struct]: unsigned64BitValue */

// unsigned8BitValue - Returns the value of the OSNumber as a uint8_t value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/unsigned8BitValue
type unsigned8BitValue struct {
}/* debug [types.gen.go/struct]: unsigned8BitValue */

// withNumber - Allocates an OSNumber object with value from a c-string and size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumber/withNumber-2c6oe
type withNumber struct {
}/* debug [types.gen.go/struct]: withNumber */

// OSObject - The base class for DriverKit objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObject
type OSObject struct {
}/* debug [types.gen.go/struct]: OSObject */

// OSSerialization - A container for one or more objects, serialized in a binary data format that is suitable for messaging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization
type OSSerialization struct {
}/* debug [types.gen.go/struct]: OSSerialization */

// copyObject - Obtain the result of the deserialization performed by createFromBytes().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/copyObject
type copyObject struct {
}/* debug [types.gen.go/struct]: copyObject */

// createFromBytes - Allocates an OSSerialization object from the serialized data of a previous serialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/createFromBytes
type createFromBytes struct {
}/* debug [types.gen.go/struct]: createFromBytes */

// createFromObject - Allocates an OSSerialization object with the serialized data of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/createFromObject
type createFromObject struct {
}/* debug [types.gen.go/struct]: createFromObject */

// finalizeBuffer - Obtain the result of the serialization performed by createFromObject().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/finalizeBuffer
type finalizeBuffer struct {
}/* debug [types.gen.go/struct]: finalizeBuffer */

// freeBuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerialization/freeBuffer
type freeBuffer struct {
}/* debug [types.gen.go/struct]: freeBuffer */

// OSString - A container for managing an array of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString
type OSString struct {
}/* debug [types.gen.go/struct]: OSString */

// getCStringNoCopy - Returns a pointer to the OSString object’s internal data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/getCStringNoCopy
type getCStringNoCopy struct {
}/* debug [types.gen.go/struct]: getCStringNoCopy */

// withCString - Allocates an OSString object with a copy of a c-string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/withCString-4wsql
type withCString struct {
}/* debug [types.gen.go/struct]: withCString */

// withCStringNoCopy - Allocates an OSString object with a copy of a c-string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/withCStringNoCopy
type withCStringNoCopy struct {
}/* debug [types.gen.go/struct]: withCStringNoCopy */

// withString - Allocates an OSString object with a copy of an OString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSString/withString
type withString struct {
}/* debug [types.gen.go/struct]: withString */

// IOAddressSegment - A structure that describes the location and size of a block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOAddressSegment
type IOAddressSegment struct {
	Address uint64 // The address of the buffer in the current process’ virtual memory space.
	Length uint64 // The size of the memory buffer, in bytes.
}/* debug [types.gen.go/struct]: IOAddressSegment */

// IOCallOnceFlag
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCallOnceFlag
type IOCallOnceFlag struct {
	Opaque unsafe.Pointer
}/* debug [types.gen.go/struct]: IOCallOnceFlag */

// IOCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommand
type IOCommand struct {
}/* debug [types.gen.go/struct]: IOCommand */

// CommandChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommand/CommandChain
type CommandChain struct {
}/* debug [types.gen.go/struct]: CommandChain */

// FromChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommand/FromChain
type FromChain struct {
}/* debug [types.gen.go/struct]: FromChain */

// IOCommandPool
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool
type IOCommandPool struct {
}/* debug [types.gen.go/struct]: IOCommandPool */

// gatedGetCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/gatedGetCommand
type gatedGetCommand struct {
}/* debug [types.gen.go/struct]: gatedGetCommand */

// gatedReturnCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/gatedReturnCommand
type gatedReturnCommand struct {
}/* debug [types.gen.go/struct]: gatedReturnCommand */

// getCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/getCommand
type getCommand struct {
}/* debug [types.gen.go/struct]: getCommand */

// initWithQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/initWithQueue
type initWithQueue struct {
}/* debug [types.gen.go/struct]: initWithQueue */

// returnCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/returnCommand
type returnCommand struct {
}/* debug [types.gen.go/struct]: returnCommand */

// withQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCommandPool/withQueue
type withQueue struct {
}/* debug [types.gen.go/struct]: withQueue */

// IODMACommand - An object that converts memory references to I/O bus addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommand
type IODMACommand struct {
}/* debug [types.gen.go/struct]: IODMACommand */

// CompleteDMA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommand/CompleteDMA
type CompleteDMA struct {
}/* debug [types.gen.go/struct]: CompleteDMA */

// GetPreparation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommand/GetPreparation
type GetPreparation struct {
}/* debug [types.gen.go/struct]: GetPreparation */

// PerformOperation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommand/PerformOperation
type PerformOperation struct {
}/* debug [types.gen.go/struct]: PerformOperation */

// PrepareForDMA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommand/PrepareForDMA
type PrepareForDMA struct {
}/* debug [types.gen.go/struct]: PrepareForDMA */

// IODMACommandSpecification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODMACommandSpecification
type IODMACommandSpecification struct {
	MaxAddressBits uint64
	Options uint64
}/* debug [types.gen.go/struct]: IODMACommandSpecification */

// IOHistogramReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter_IVars
type IOHistogramReporter_IVars struct {
	BucketBounds unsafe.Pointer
	BucketCount int
	HistogramSegmentsConfig IOHistogramSegmentConfig
	SegmentCount int
}/* debug [types.gen.go/struct]: IOHistogramReporter_IVars */

// handleCreateLegend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter_IVars/handleCreateLegend
type handleCreateLegend struct {
}/* debug [types.gen.go/struct]: handleCreateLegend */

// ~IOHistogramReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReporter_IVars/~IOHistogramReporter_IVars
type ~IOHistogramReporter_IVars struct {
}/* debug [types.gen.go/struct]: ~IOHistogramReporter_IVars */

// IOHistogramReportValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramReportValues
type IOHistogramReportValues struct {
	Bucket_hits uint64
	Bucket_max int64
	Bucket_min int64
	Bucket_sum int64
}/* debug [types.gen.go/struct]: IOHistogramReportValues */

// IOHistogramSegmentConfig
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOHistogramSegmentConfig
type IOHistogramSegmentConfig struct {
	Base_bucket_width uint32
	Scale_flag uint32
	Segment_bucket_count uint32
	Segment_idx uint32
}/* debug [types.gen.go/struct]: IOHistogramSegmentConfig */

// IOInterruptDispatchSourcePayload - A private structure for an interrupt dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOInterruptDispatchSourcePayload
type IOInterruptDispatchSourcePayload struct {
	Count uint64
	Time uint64
}/* debug [types.gen.go/struct]: IOInterruptDispatchSourcePayload */

// IONormDistReportValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IONormDistReportValues
type IONormDistReportValues struct {
	Mean uint64
	Reserved uint64
	Samples uint64
	Variance uint64
}/* debug [types.gen.go/struct]: IONormDistReportValues */

// IOReportChannel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportChannel
type IOReportChannel struct {
	Channel_id uint64
	Channel_type ReportChannelType
}/* debug [types.gen.go/struct]: IOReportChannel */

// IOReportChannelList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportChannelList
type IOReportChannelList struct {
	Channels ReportChannel
	Nchannels uint32
}/* debug [types.gen.go/struct]: IOReportChannelList */

// IOReportChannelType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportChannelType
type IOReportChannelType struct {
	Categories uint16
	Element_idx int16
	Nelements uint16
	Report_format uint8
	Reserved uint8
}/* debug [types.gen.go/struct]: IOReportChannelType */

// IOReportElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportElement
type IOReportElement struct {
	Channel_id uint64
	Channel_type ReportChannelType
	Provider_id uint64
	Timestamp uint64
	Values ReportElementValues
}/* debug [types.gen.go/struct]: IOReportElement */

// IOReportElementValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportElementValues
type IOReportElementValues struct {
	V uint64
}/* debug [types.gen.go/struct]: IOReportElementValues */

// IOReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars
type IOReporter_IVars struct {
	ChannelDimension uint16
	ChannelNames OSArray
	ChannelType ReportChannelType
	ConfigLock unsafe.Pointer
	Driver_id uint64
	Elements IOReportElement
	EnableCounts []int
	Enabled int
	NChannels int
	NElements int
	ReporterConfigIsLocked bool
	ReporterIsLocked bool
	ReporterLock unsafe.Pointer
	SwapElements IOReportElement
	SwapEnableCounts []int
	Unit ReportUnit
}/* debug [types.gen.go/struct]: IOReporter_IVars */

// copyChannelIDs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/copyChannelIDs
type copyChannelIDs struct {
}/* debug [types.gen.go/struct]: copyChannelIDs */

// copyElementValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/copyElementValues
type copyElementValues struct {
}/* debug [types.gen.go/struct]: copyElementValues */

// getChannelIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/getChannelIndex
type getChannelIndex struct {
}/* debug [types.gen.go/struct]: getChannelIndex */

// getChannelIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/getChannelIndices
type getChannelIndices struct {
}/* debug [types.gen.go/struct]: getChannelIndices */

// getElementValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/getElementValues
type getElementValues struct {
}/* debug [types.gen.go/struct]: getElementValues */

// getFirstElementIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/getFirstElementIndex
type getFirstElementIndex struct {
}/* debug [types.gen.go/struct]: getFirstElementIndex */

// handleAddChannelSwap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/handleAddChannelSwap
type handleAddChannelSwap struct {
}/* debug [types.gen.go/struct]: handleAddChannelSwap */

// handleConfigureReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/handleConfigureReport
type handleConfigureReport struct {
}/* debug [types.gen.go/struct]: handleConfigureReport */

// handleSwapCleanup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/handleSwapCleanup
type handleSwapCleanup struct {
}/* debug [types.gen.go/struct]: handleSwapCleanup */

// handleSwapPrepare
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/handleSwapPrepare
type handleSwapPrepare struct {
}/* debug [types.gen.go/struct]: handleSwapPrepare */

// handleUpdateReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/handleUpdateReport
type handleUpdateReport struct {
}/* debug [types.gen.go/struct]: handleUpdateReport */

// legendWith
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/legendWith
type legendWith struct {
}/* debug [types.gen.go/struct]: legendWith */

// lockReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/lockReporter
type lockReporter struct {
}/* debug [types.gen.go/struct]: lockReporter */

// lockReporterConfig
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/lockReporterConfig
type lockReporterConfig struct {
}/* debug [types.gen.go/struct]: lockReporterConfig */

// setElementValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/setElementValues
type setElementValues struct {
}/* debug [types.gen.go/struct]: setElementValues */

// unlockReporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/unlockReporter
type unlockReporter struct {
}/* debug [types.gen.go/struct]: unlockReporter */

// unlockReporterConfig
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/unlockReporterConfig
type unlockReporterConfig struct {
}/* debug [types.gen.go/struct]: unlockReporterConfig */

// updateChannelValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/updateChannelValues
type updateChannelValues struct {
}/* debug [types.gen.go/struct]: updateChannelValues */

// updateReportChannel - Internal method to extract channel data to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/updateReportChannel
type updateReportChannel struct {
}/* debug [types.gen.go/struct]: updateReportChannel */

// valid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/valid
type valid struct {
}/* debug [types.gen.go/struct]: valid */

// ~IOReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReporter_IVars/~IOReporter_IVars
type ~IOReporter_IVars struct {
}/* debug [types.gen.go/struct]: ~IOReporter_IVars */

// IOReportInterest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportInterest
type IOReportInterest struct {
	Channel ReportChannel
	Provider_id uint64
}/* debug [types.gen.go/struct]: IOReportInterest */

// IOReportInterestList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOReportInterestList
type IOReportInterestList struct {
	Interests ReportInterest
	Ninterests uint32
}/* debug [types.gen.go/struct]: IOReportInterestList */

// IORPC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPC
type IORPC struct {
	Message IORPCMessageMach
	Reply IORPCMessageMach
	ReplySize uint32
	SendSize uint32
}/* debug [types.gen.go/struct]: IORPC */

// IORPCMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessage
type IORPCMessage struct {
	Flags uint64
	Msgid uint64
	ObjectRefs uint64
	Objects unsafe.Pointer
}/* debug [types.gen.go/struct]: IORPCMessage */

// IORPCMessageErrorReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessageErrorReturn
type IORPCMessageErrorReturn struct {
	Content RPCMessageErrorReturnContent
	Mach RPCMessageMach
}/* debug [types.gen.go/struct]: IORPCMessageErrorReturn */

// IORPCMessageErrorReturnContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessageErrorReturnContent
type IORPCMessageErrorReturnContent struct {
	Hdr RPCMessage
	Pad uint32
	Result unsafe.Pointer
}/* debug [types.gen.go/struct]: IORPCMessageErrorReturnContent */

// IORPCMessageMach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessageMach
type IORPCMessageMach struct {
	Msgh Mach_msg_header_t
	Msgh_body Mach_msg_body_t
	Objects Mach_msg_port_descriptor_t
}/* debug [types.gen.go/struct]: IORPCMessageMach */

// IOServiceNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource
type IOServiceNotificationDispatchSource struct {
}/* debug [types.gen.go/struct]: IOServiceNotificationDispatchSource */

// CopyNextNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/CopyNextNotification
type CopyNextNotification struct {
}/* debug [types.gen.go/struct]: CopyNextNotification */

// DeliverNotifications
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/DeliverNotifications
type DeliverNotifications struct {
}/* debug [types.gen.go/struct]: DeliverNotifications */

// ServiceNotificationReady
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceNotificationDispatchSource/ServiceNotificationReady
type ServiceNotificationReady struct {
}/* debug [types.gen.go/struct]: ServiceNotificationReady */

// IOSimpleArrayReportValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleArrayReportValues
type IOSimpleArrayReportValues struct {
	Simple_values int64
}/* debug [types.gen.go/struct]: IOSimpleArrayReportValues */

// IOSimpleReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReporter_IVars
type IOSimpleReporter_IVars struct {
}/* debug [types.gen.go/struct]: IOSimpleReporter_IVars */

// IOSimpleReportValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSimpleReportValues
type IOSimpleReportValues struct {
	Reserved1 uint64
	Reserved2 uint64
	Reserved3 uint64
	Simple_value int64
}/* debug [types.gen.go/struct]: IOSimpleReportValues */

// IOStateReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars
type IOStateReporter_IVars struct {
	CurrentStates []int
	LastUpdateTimes []uint64
	SwapCurrentStates []int
	SwapLastUpdateTimes []uint64
}/* debug [types.gen.go/struct]: IOStateReporter_IVars */

// handleIncrementChannelStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleIncrementChannelStateByIndices
type handleIncrementChannelStateByIndices struct {
}/* debug [types.gen.go/struct]: handleIncrementChannelStateByIndices */

// handleOverrideChannelStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleOverrideChannelStateByIndices
type handleOverrideChannelStateByIndices struct {
}/* debug [types.gen.go/struct]: handleOverrideChannelStateByIndices */

// handleSetStateByIndices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSetStateByIndices
type handleSetStateByIndices struct {
}/* debug [types.gen.go/struct]: handleSetStateByIndices */

// handleSetStateID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/handleSetStateID
type handleSetStateID struct {
}/* debug [types.gen.go/struct]: handleSetStateID */

// ~IOStateReporter_IVars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_IVars/~IOStateReporter_IVars
type ~IOStateReporter_IVars struct {
}/* debug [types.gen.go/struct]: ~IOStateReporter_IVars */

// IOStateReportValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReportValues
type IOStateReportValues struct {
	Intransitions uint64
	Last_intransition uint64
	State_id uint64
	Upticks uint64
}/* debug [types.gen.go/struct]: IOStateReportValues */

// IOUserClientMethodArguments - Arguments to pass to IOConnectMethod calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClientMethodArguments
type IOUserClientMethodArguments struct {
	Completion OSAction // An action for processing asynchronous data received from the service.
	ScalarInput []uint64 // An array of scalars from the caller.
	ScalarInputCount uint32 // The number of scalars provided by the caller.
	ScalarOutput []uint64 // An array of scalars to return to the caller.
	ScalarOutputCount uint32 // The number of scalars to return to the caller.
	Selector uint64 // The index of the method you want to execute.
	StructureInput OSData // A data object containing the structure input from the IOKit connect method.
	StructureInputDescriptor IOMemoryDescriptor // A memory descriptor containing structure input from the IOKit connect method.
	StructureOutput OSData // A data object to return to the caller as structure output.
	StructureOutputDescriptor IOMemoryDescriptor // An IOMemoryDescriptor specified by the caller for structure output.
	StructureOutputMaximumSize uint64 // The maximum size of the output structure that you specified.
	Version uint64
}/* debug [types.gen.go/struct]: IOUserClientMethodArguments */

// IOUserClientMethodDispatch - A structure that specifies how to validate the arguments passed to a client method function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOUserClientMethodDispatch
type IOUserClientMethodDispatch struct {
	CheckCompletionExists uint32 // An integer value indicating whether to check for the existence of a completion action.
	CheckScalarInputCount uint32 // The expected number of scalar inputs.
	CheckScalarOutputCount uint32 // The expected number of scalar outputs.
	CheckStructureInputSize uint32 // The expected size of the scalar inputs.
	CheckStructureOutputSize uint32 // The expected size of the scalar outputs.
	Function UserClientMethodFunction // The function to call after validating all of the specified values.
}/* debug [types.gen.go/struct]: IOUserClientMethodDispatch */

// IVarsInvalidator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IVarsInvalidator
type IVarsInvalidator struct {
	Ivars IOReporter_IVars
}/* debug [types.gen.go/struct]: IVarsInvalidator */

// disarm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IVarsInvalidator/disarm
type disarm struct {
}/* debug [types.gen.go/struct]: disarm */

// ~IVarsInvalidator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IVarsInvalidator/~IVarsInvalidator
type ~IVarsInvalidator struct {
}/* debug [types.gen.go/struct]: ~IVarsInvalidator */

// mach_msg_body_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_body_t
type mach_msg_body_t struct {
	Msgh_descriptor_count unsafe.Pointer
}/* debug [types.gen.go/struct]: mach_msg_body_t */

// mach_msg_header_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_header_t
type mach_msg_header_t struct {
	Msgh_bits unsafe.Pointer
	Msgh_id unsafe.Pointer
	Msgh_remote_port unsafe.Pointer
	Msgh_size unsafe.Pointer
	Msgh_voucher_port unsafe.Pointer
}/* debug [types.gen.go/struct]: mach_msg_header_t */

// mach_msg_max_trailer_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_max_trailer_t
type mach_msg_max_trailer_t struct {
	Val unsafe.Pointer
}/* debug [types.gen.go/struct]: mach_msg_max_trailer_t */

// mach_msg_ool_descriptor_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_ool_descriptor_t
type mach_msg_ool_descriptor_t struct {
	Address unsafe.Pointer
	Copy unsafe.Pointer
	Deallocate int
	Pad1 unsafe.Pointer
	Size unsafe.Pointer
	Type unsafe.Pointer
}/* debug [types.gen.go/struct]: mach_msg_ool_descriptor_t */

// mach_msg_port_descriptor_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_msg_port_descriptor_t
type mach_msg_port_descriptor_t struct {
	Disposition unsafe.Pointer
	Name unsafe.Pointer
	Pad1 unsafe.Pointer
	Pad2 unsafe.Pointer
	Type unsafe.Pointer
}/* debug [types.gen.go/struct]: mach_msg_port_descriptor_t */

// mach_timebase_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_timebase_info-c.struct
type mach_timebase_info struct {
	Denom uint32
	Numer uint32
}/* debug [types.gen.go/struct]: mach_timebase_info */

// OS_no_retain_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OS_no_retain_t
type OS_no_retain_t struct {
}/* debug [types.gen.go/struct]: OS_no_retain_t */

// OS_retain_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OS_retain_t
type OS_retain_t struct {
}/* debug [types.gen.go/struct]: OS_retain_t */

// OSClassDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSClassDescription
type OSClassDescription struct {
	DescriptionSize uint32
	DispatchNames unsafe.Pointer
	Flags uint64
	MetaMethodNames unsafe.Pointer
	MetaMethodNamesOffset uint32
	MetaMethodNamesSize uint32
	MetaMethodOptions uint64
	MetaMethodOptionsOffset uint32
	MetaMethodOptionsSize uint32
	MethodNames unsafe.Pointer
	MethodNamesOffset uint32
	MethodNamesSize uint32
	MethodOptions uint64
	MethodOptionsOffset uint32
	MethodOptionsSize uint32
	Name unsafe.Pointer
	QueueNamesOffset uint32
	QueueNamesSize uint32
	Resv1 uint64
	SuperName unsafe.Pointer
}/* debug [types.gen.go/struct]: OSClassDescription */

// OSClassLoadInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSClassLoadInformation
type OSClassLoadInformation struct {
	Description OSClassDescription
	InstanceSize uint32
	MetaPointer OSMetaClass
	New unsafe.Pointer
	Resv2 uint64
	Resv3 uint64
	Version uint32
}/* debug [types.gen.go/struct]: OSClassLoadInformation */

// OSInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSInterface
type OSInterface struct {
}/* debug [types.gen.go/struct]: OSInterface */

// OSMetaClass - Base class for DriverKit runtime class system. Not called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClass
type OSMetaClass struct {
	MetaClassPrivate unsafe.Pointer
}/* debug [types.gen.go/struct]: OSMetaClass */

// New
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClass/New
type New struct {
}/* debug [types.gen.go/struct]: New */

// OSMetaClassBase - Base class for DriverKit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase
type OSMetaClassBase struct {
	Meta OSMetaClass
	Refcount unsafe.Pointer
	Reserved int32
}/* debug [types.gen.go/struct]: OSMetaClassBase */

// Dispatch - Runtime internals
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/Dispatch
type Dispatch struct {
}/* debug [types.gen.go/struct]: Dispatch */

// GetClass - Internal helper for GetClassName. Not to be called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/GetClass
type GetClass struct {
}/* debug [types.gen.go/struct]: GetClass */

// GetClassName - Returns the name of the class given an OSObject pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/GetClassName-50b5y
type GetClassName struct {
}/* debug [types.gen.go/struct]: GetClassName */

// getMetaClass - Internal helper for GetClassName. Not to be called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/getMetaClass
type getMetaClass struct {
}/* debug [types.gen.go/struct]: getMetaClass */

// getRetainCount
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/getRetainCount
type getRetainCount struct {
}/* debug [types.gen.go/struct]: getRetainCount */

// Invoke - Runtime internals
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/Invoke
type Invoke struct {
}/* debug [types.gen.go/struct]: Invoke */

// IsRemote
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/IsRemote
type IsRemote struct {
}/* debug [types.gen.go/struct]: IsRemote */

// operator new
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/operator_new
type operator new struct {
}/* debug [types.gen.go/struct]: operator new */

// requiredMetaCast - Internal helper for OSRequiredCast. Not to be called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/requiredMetaCast
type requiredMetaCast struct {
}/* debug [types.gen.go/struct]: requiredMetaCast */

// safeMetaCast - Internal helper for OSDynamicCast. Not to be called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSMetaClassBase/safeMetaCast
type safeMetaCast struct {
}/* debug [types.gen.go/struct]: safeMetaCast */

// OSOrderedSet - OSOrderedSet provides an ordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSOrderedSet
type OSOrderedSet struct {
}/* debug [types.gen.go/struct]: OSOrderedSet */

// getFirstObject - The object at index 0 in the ordered set if there is one, otherwise 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSOrderedSet/getFirstObject
type getFirstObject struct {
}/* debug [types.gen.go/struct]: getFirstObject */

// orderObject - Calls the ordered set’s order function against a 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSOrderedSet/orderObject
type orderObject struct {
}/* debug [types.gen.go/struct]: orderObject */

// setFirstObject - Adds an object to the OSOrderedSet at index 0 if it is not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSOrderedSet/setFirstObject
type setFirstObject struct {
}/* debug [types.gen.go/struct]: setFirstObject */

// setLastObject - Adds an object at the end of the OSOrderedSet if it is not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSOrderedSet/setLastObject
type setLastObject struct {
}/* debug [types.gen.go/struct]: setLastObject */

// OSSet - OSSet provides an unordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSet
type OSSet struct {
}/* debug [types.gen.go/struct]: OSSet */

// containsObject
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSet/containsObject
type containsObject struct {
}/* debug [types.gen.go/struct]: containsObject */

// getAnyObject - Returns an arbitrary (not random) object from the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSet/getAnyObject
type getAnyObject struct {
}/* debug [types.gen.go/struct]: getAnyObject */

// member - Checks the set for the presence of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSet/member
type member struct {
}/* debug [types.gen.go/struct]: member */

// queue_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/queue_entry
type queue_entry struct {
	Next queue_entry
	Prev queue_entry
}/* debug [types.gen.go/struct]: queue_entry */

// SCSI_Capacity_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSI_Capacity_Data
type SCSI_Capacity_Data struct {
	BLOCK_LENGTH_IN_BYTES unsafe.Pointer
	RETURNED_LOGICAL_BLOCK_ADDRESS unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSI_Capacity_Data */

// SCSI_Capacity_Data_Long
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSI_Capacity_Data_Long
type SCSI_Capacity_Data_Long struct {
	BLOCK_LENGTH_IN_BYTES unsafe.Pointer
	Reserved unsafe.Pointer
	RETURNED_LOGICAL_BLOCK_ADDRESS unsafe.Pointer
	RTO_EN_PROT_EN unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSI_Capacity_Data_Long */

// SCSI_Sense_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSI_Sense_Data
type SCSI_Sense_Data struct {
	ADDITIONAL_SENSE_CODE unsafe.Pointer
	ADDITIONAL_SENSE_CODE_QUALIFIER unsafe.Pointer
	ADDITIONAL_SENSE_LENGTH unsafe.Pointer
	COMMAND_SPECIFIC_INFORMATION_1 unsafe.Pointer
	COMMAND_SPECIFIC_INFORMATION_2 unsafe.Pointer
	COMMAND_SPECIFIC_INFORMATION_3 unsafe.Pointer
	COMMAND_SPECIFIC_INFORMATION_4 unsafe.Pointer
	FIELD_REPLACEABLE_UNIT_CODE unsafe.Pointer
	INFORMATION_1 unsafe.Pointer
	INFORMATION_2 unsafe.Pointer
	INFORMATION_3 unsafe.Pointer
	INFORMATION_4 unsafe.Pointer
	SEGMENT_NUMBER unsafe.Pointer
	SENSE_KEY unsafe.Pointer
	SENSE_KEY_SPECIFIC_LSB unsafe.Pointer
	SENSE_KEY_SPECIFIC_MID unsafe.Pointer
	SKSV_SENSE_KEY_SPECIFIC_MSB unsafe.Pointer
	VALID_RESPONSE_CODE unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSI_Sense_Data */

// SCSICmd_INQUIRY_Page00_Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page00_Header
type SCSICmd_INQUIRY_Page00_Header struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page00_Header */

// SCSICmd_INQUIRY_Page00_Header_SPC_16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page00_Header_SPC_16
type SCSICmd_INQUIRY_Page00_Header_SPC_16 struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page00_Header_SPC_16 */

// SCSICmd_INQUIRY_Page80_Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page80_Header
type SCSICmd_INQUIRY_Page80_Header struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	PRODUCT_SERIAL_NUMBER unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page80_Header */

// SCSICmd_INQUIRY_Page80_Header_SPC_16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page80_Header_SPC_16
type SCSICmd_INQUIRY_Page80_Header_SPC_16 struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	PRODUCT_SERIAL_NUMBER unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page80_Header_SPC_16 */

// SCSICmd_INQUIRY_Page83_Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_Header
type SCSICmd_INQUIRY_Page83_Header struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_Header */

// SCSICmd_INQUIRY_Page83_Header_SPC_16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_Header_SPC_16
type SCSICmd_INQUIRY_Page83_Header_SPC_16 struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_Header_SPC_16 */

// SCSICmd_INQUIRY_Page83_Identification_Descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_Identification_Descriptor
type SCSICmd_INQUIRY_Page83_Identification_Descriptor struct {
	CODE_SET unsafe.Pointer
	IDENTIFIER unsafe.Pointer
	IDENTIFIER_LENGTH unsafe.Pointer
	IDENTIFIER_TYPE unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_Identification_Descriptor */

// SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier
type SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier struct {
	LOGICAL_UNIT_GROUP unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier */

// SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier
type SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier struct {
	OBSOLETE unsafe.Pointer
	RELATIVE_TARGET_PORT_IDENTIFIER unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier */

// SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier
type SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier struct {
	RESERVED unsafe.Pointer
	TARGET_PORT_GROUP unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier */

// SCSICmd_INQUIRY_Page89_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_Page89_Data
type SCSICmd_INQUIRY_Page89_Data struct {
	ATA_DEVICE_SIGNATURE unsafe.Pointer
	COMMAND_CODE unsafe.Pointer
	IDENTIFY_DATA unsafe.Pointer
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	Reserved unsafe.Pointer
	Reserved2 unsafe.Pointer
	SAT_PRODUCT_IDENTIFICATION unsafe.Pointer
	SAT_PRODUCT_REVISION_LEVEL unsafe.Pointer
	SAT_VENDOR_IDENTIFICATION unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_Page89_Data */

// SCSICmd_INQUIRY_PageB0_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageB0_Data
type SCSICmd_INQUIRY_PageB0_Data struct {
	ATOMIC_ALIGNMENT unsafe.Pointer
	ATOMIC_TRANSFER_LENGTH_GRANULARITY unsafe.Pointer
	MAXIMUM_ATOMIC_BOUNDARY_SIZE unsafe.Pointer
	MAXIMUM_ATOMIC_TRANSFER_LENGTH unsafe.Pointer
	MAXIMUM_ATOMIC_TRANSFER_LENGTH_WITH_ATOMIC_BOUNDARY unsafe.Pointer
	MAXIMUM_COMPARE_AND_WRITE_LENGTH unsafe.Pointer
	MAXIMUM_PREFETCH_LENGTH unsafe.Pointer
	MAXIMUM_TRANSFER_LENGTH unsafe.Pointer
	MAXIMUM_UNMAP_BLOCK_DESCRIPTOR_COUNT unsafe.Pointer
	MAXIMUM_UNMAP_LBA_COUNT unsafe.Pointer
	MAXIMUM_WRITE_SAME_LENGTH unsafe.Pointer
	OPTIMAL_TRANSFER_LENGTH unsafe.Pointer
	OPTIMAL_TRANSFER_LENGTH_GRANULARITY unsafe.Pointer
	OPTIMAL_UNMAP_GRANULARITY unsafe.Pointer
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	UNMAP_GRANULARITY_ALIGNMENT unsafe.Pointer
	WSNZ unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageB0_Data */

// SCSICmd_INQUIRY_PageB1_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageB1_Data
type SCSICmd_INQUIRY_PageB1_Data struct {
	MEDIUM_ROTATION_RATE unsafe.Pointer
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	Reserved unsafe.Pointer
	Reserved2 unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageB1_Data */

// SCSICmd_INQUIRY_PageB2_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageB2_Data
type SCSICmd_INQUIRY_PageB2_Data struct {
	LBP_FLAGS unsafe.Pointer
	MINIMUM_PERCENTAGE unsafe.Pointer
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	THRESHOLD_EXPONENT unsafe.Pointer
	THRESHOLD_PERCENTAGE unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageB2_Data */

// SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor
type SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor struct {
	DESIGNATION_DESCRIPTOR unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor */

// SCSICmd_INQUIRY_PageC0_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageC0_Data
type SCSICmd_INQUIRY_PageC0_Data struct {
	FEncryptionType unsafe.Pointer
	FFeatures unsafe.Pointer
	FHeader SICmd_INQUIRY_PAGECx_Header
	FInstalledRAMSize unsafe.Pointer
	FMacModelId unsafe.Pointer
	FMaxReadSize unsafe.Pointer
	FMaxWriteSize unsafe.Pointer
	FNativeBlockSize unsafe.Pointer
	FPreferredIOSize unsafe.Pointer
	FReserved1 unsafe.Pointer
	FReserved2 unsafe.Pointer
	FReserved3 unsafe.Pointer
	FSerialNumber unsafe.Pointer
	FTdmPageVersion unsafe.Pointer
	FTdmProtocolVersion unsafe.Pointer
	FWorkArounds unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageC0_Data */

// SCSICmd_INQUIRY_PageC1_Data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageC1_Data
type SCSICmd_INQUIRY_PageC1_Data struct {
	FHeader SICmd_INQUIRY_PAGECx_Header
	FPowerRequired unsafe.Pointer
	FReserved1 unsafe.Pointer
	FReserved2 unsafe.Pointer
	FTdmPowerRequirementsPageVersion unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageC1_Data */

// SCSICmd_INQUIRY_PageCx_Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_PageCx_Header-c.struct
type SCSICmd_INQUIRY_PageCx_Header struct {
	PAGE_CODE unsafe.Pointer
	PAGE_LENGTH unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	RESERVED unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_PageCx_Header */

// SCSICmd_INQUIRY_StandardData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_StandardData
type SCSICmd_INQUIRY_StandardData struct {
	ADDITIONAL_LENGTH unsafe.Pointer
	Flags1 unsafe.Pointer
	Flags2 unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	PRODUCT_IDENTIFICATION unsafe.Pointer
	PRODUCT_REVISION_LEVEL unsafe.Pointer
	RESPONSE_DATA_FORMAT unsafe.Pointer
	RMB unsafe.Pointer
	SCCSReserved unsafe.Pointer
	VENDOR_IDENTIFICATION unsafe.Pointer
	VERSION unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_StandardData */

// SCSICmd_INQUIRY_StandardDataAll
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSICmd_INQUIRY_StandardDataAll
type SCSICmd_INQUIRY_StandardDataAll struct {
	ADDITIONAL_LENGTH unsafe.Pointer
	Flags1 unsafe.Pointer
	Flags2 unsafe.Pointer
	Flags3 unsafe.Pointer
	PERIPHERAL_DEVICE_TYPE unsafe.Pointer
	PRODUCT_IDENTIFICATION unsafe.Pointer
	PRODUCT_REVISION_LEVEL unsafe.Pointer
	Reserved1 unsafe.Pointer
	Reserved2 unsafe.Pointer
	RESPONSE_DATA_FORMAT unsafe.Pointer
	RMB unsafe.Pointer
	SCCSReserved unsafe.Pointer
	VENDOR_IDENTIFICATION unsafe.Pointer
	VendorSpecific1 unsafe.Pointer
	VendorSpecific2 unsafe.Pointer
	VERSION unsafe.Pointer
	VERSION_DESCRIPTOR unsafe.Pointer
}/* debug [types.gen.go/struct]: SCSICmd_INQUIRY_StandardDataAll */





