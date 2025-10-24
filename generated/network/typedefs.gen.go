// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

// Type aliases and typedefs
// Nw_advertise_descriptor_t - A description used to advertise the Bonjour service that a listener provides.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_t
// nw_advertise_descriptor_t has base type: NSObject<OS_nw_advertise_descriptor> *
type Nw_advertise_descriptor_t uintptr
// Nw_browse_descriptor_t - A service description used to discover Bonjour services.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_t
// nw_browse_descriptor_t has base type: NSObject<OS_nw_browse_descriptor> *
type Nw_browse_descriptor_t uintptr
// Nw_browse_result_change_t - Flags describing ways in which discovered services can change between specific results.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_change_t
// nw_browse_result_change_t has base type: uint64_t
type Nw_browse_result_change_t uintptr
// Nw_browse_result_t - A discovered service and metadata about the service.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_t
// nw_browse_result_t has base type: NSObject<OS_nw_browse_result> *
type Nw_browse_result_t uintptr
// Nw_browser_t - An object you use to browse for available network services.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_t
// nw_browser_t has base type: NSObject<OS_nw_browser> *
type Nw_browser_t uintptr
// Nw_connection_group_t - An object you use to communicate with a group of endpoints, such as an IP multicast group on a local network.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_t
// nw_connection_group_t has base type: NSObject<OS_nw_connection_group> *
type Nw_connection_group_t uintptr
// Nw_connection_t - A bidirectional data connection between a local endpoint and a remote endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_t
// nw_connection_t has base type: NSObject<OS_nw_connection> *
type Nw_connection_t uintptr
// Nw_content_context_t - A representation of a message to send or receive, containing protocol metadata and send properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_t
// nw_content_context_t has base type: NSObject<OS_nw_content_context> *
type Nw_content_context_t uintptr
// Nw_data_transfer_report_t - A report that provides metrics about data being sent and received on a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_t
// nw_data_transfer_report_t has base type: NSObject<OS_nw_data_transfer_report> *
type Nw_data_transfer_report_t uintptr
// Nw_endpoint_t - A local or remote endpoint in a network connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_t
// nw_endpoint_t has base type: NSObject<OS_nw_endpoint> *
type Nw_endpoint_t uintptr
// Nw_error_t - The errors returned by the Network framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_t
// nw_error_t has base type: NSObject<OS_nw_error> *
type Nw_error_t uintptr
// Nw_establishment_report_t - A report that provides metrics about how a connection was established.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_t
// nw_establishment_report_t has base type: NSObject<OS_nw_establishment_report> *
type Nw_establishment_report_t uintptr
// Nw_ethernet_address_t - A 48-bit Ethernet address.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_address_t
// nw_ethernet_address_t has base type: unsigned char[6]
type Nw_ethernet_address_t uintptr
// Nw_ethernet_channel_t - An object you use to send and receive custom Ethernet frames.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_t
// nw_ethernet_channel_t has base type: NSObject<OS_nw_ethernet_channel> *
type Nw_ethernet_channel_t uintptr
// Nw_framer_message_t - A message for a custom protocol, in which you can store arbitrary key-value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_t
// nw_framer_message_t has base type: nw_protocol_metadata_t
type Nw_framer_message_t uintptr
// Nw_framer_t - An object that represents a single instance of your custom protocol running in a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_t
// nw_framer_t has base type: NSObject<OS_nw_framer> *
type Nw_framer_t uintptr
// Nw_group_descriptor_t - A type that defines a group of endpoints with which you can communicate, such as a multicast group.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_t
// nw_group_descriptor_t has base type: NSObject<OS_nw_group_descriptor> *
type Nw_group_descriptor_t uintptr
// Nw_interface_t - An interface that a network connection uses to send and receive data.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_t
// nw_interface_t has base type: NSObject<OS_nw_interface> *
type Nw_interface_t uintptr
// Nw_listener_t - An object you use to listen for incoming network connections.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_t
// nw_listener_t has base type: NSObject<OS_nw_listener> *
type Nw_listener_t uintptr
// Nw_object_t - The generic type for objects in the Network framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_object_t
// nw_object_t has base type: NSObject<OS_nw_object> *
type Nw_object_t uintptr
// Nw_parameters_t - An object that stores the protocols to use for connections, options for sending data, and network path constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_t
// nw_parameters_t has base type: NSObject<OS_nw_parameters> *
type Nw_parameters_t uintptr
// Nw_path_monitor_t - An observer that you use to monitor and react to network changes.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_t
// nw_path_monitor_t has base type: NSObject<OS_nw_path_monitor> *
type Nw_path_monitor_t uintptr
// Nw_path_t - An object that contains information about the properties of the network that a connection uses, or that are available to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_t
// nw_path_t has base type: NSObject<OS_nw_path> *
type Nw_path_t uintptr
// Nw_privacy_context_t - An object that defines the privacy requirements for a set of connections.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_t
// nw_privacy_context_t has base type: NSObject<OS_nw_privacy_context> *
type Nw_privacy_context_t uintptr
// Nw_protocol_definition_t - The abstract superclass for identifying a network protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_definition_t
// nw_protocol_definition_t has base type: NSObject<OS_nw_protocol_definition> *
type Nw_protocol_definition_t uintptr
// Nw_protocol_metadata_t - The abstract superclass for specifying metadata about a network protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_t
// nw_protocol_metadata_t has base type: NSObject<OS_nw_protocol_metadata> *
type Nw_protocol_metadata_t uintptr
// Nw_protocol_options_t - The abstract superclass for configuring the options of a network protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_options_t
// nw_protocol_options_t has base type: NSObject<OS_nw_protocol_options> *
type Nw_protocol_options_t uintptr
// Nw_protocol_stack_t - An ordered set of protocol options that define the protocols that connections and listeners use.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_t
// nw_protocol_stack_t has base type: NSObject<OS_nw_protocol_stack> *
type Nw_protocol_stack_t uintptr
// Nw_resolution_report_t - A description of a single DNS resolution step.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_t
// nw_resolution_report_t has base type: NSObject<OS_nw_resolution_report> *
type Nw_resolution_report_t uintptr
// Nw_txt_record_t - A dictionary representing a TXT record in a DNS packet.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_t
// nw_txt_record_t has base type: NSObject<OS_nw_txt_record> *
type Nw_txt_record_t uintptr
// Nw_ws_request_t - A WebSocket handshake request sent from a client to a server.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_request_t
// nw_ws_request_t has base type: NSObject<OS_nw_ws_request> *
type Nw_ws_request_t uintptr
// Nw_ws_response_t - A WebSocket handshake reponse sent from a server to a client.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_t
// nw_ws_response_t has base type: NSObject<OS_nw_ws_response> *
type Nw_ws_response_t uintptr

