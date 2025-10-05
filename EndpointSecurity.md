# EndpointSecurity Framework

> **Note:** This documentation covers Apple's EndpointSecurity framework for building security products.

## Overview

Develop system extensions that enhance user security.

## Platforms

- Mac Catalyst 13.0
- macOS 10.15

## API Reference

## Topic Groups

<details>
<summary>Event Monitoring (4 items)</summary>

- [Client](/documentation/endpointsecurity/client) - An opaque type that maintains Endpoint Security client state, and functions related to this type.
- [Message](/documentation/endpointsecurity/message) - A type used by Endpoint Security to notify your client when a monitored action occurs.
- [Event Types](/documentation/endpointsecurity/event-types) - Types used by messages to deliver details specific to different kinds of Endpoint Security events.
- [Monitoring System Events with Endpoint Security](/documentation/endpointsecurity/monitoring-system-events-with-endpoint-security) - Receive notifications and authorization requests for sensitive operations by creating an Endpoint...

</details>

<details>
<summary>Entitlements (1 items)</summary>

- [com.apple.developer.endpoint-security.client](/documentation/BundleResources/Entitlements/com.apple.developer.endpoint-security.client) - The entitlement required to monitor system events for potentially malicious activity.

</details>

<details>
<summary>Reference (5 items)</summary>

- [EndpointSecurity Constants](/documentation/endpointsecurity/endpointsecurity-constants)
- [EndpointSecurity Data Types](/documentation/endpointsecurity/endpointsecurity-data-types)
- [EndpointSecurity Functions](/documentation/endpointsecurity/endpointsecurity-functions)
- [EndpointSecurity Structures](/documentation/endpointsecurity/endpointsecurity-structures)
- [EndpointSecurity Enumerations](/documentation/endpointsecurity/endpointsecurity-enumerations)

</details>

<details>
<summary>Structures (5 items)</summary>

- [es_event_tcc_modify_t](/documentation/endpointsecurity/es_event_tcc_modify_t)
- [es_tcc_authorization_reason_t](/documentation/endpointsecurity/es_tcc_authorization_reason_t) - ess_tcc_authorization_reason_t
- [es_tcc_authorization_right_t](/documentation/endpointsecurity/es_tcc_authorization_right_t) - ess_tcc_authorization_right_t
- [es_tcc_event_type_t](/documentation/endpointsecurity/es_tcc_event_type_t)
- [es_tcc_identity_type_t](/documentation/endpointsecurity/es_tcc_identity_type_t) - es_tcc_identity_type_t

</details>

<details>
<summary>Variables (30 items)</summary>

- [ES_EVENT_TYPE_NOTIFY_TCC_MODIFY](/documentation/endpointsecurity/es_event_type_notify_tcc_modify)
- [ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY](/documentation/endpointsecurity/es_tcc_authorization_reason_app_type_policy) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_ENTITLED](/documentation/endpointsecurity/es_tcc_authorization_reason_entitled) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_ERROR](/documentation/endpointsecurity/es_tcc_authorization_reason_error)
- [ES_TCC_AUTHORIZATION_REASON_MDM_POLICY](/documentation/endpointsecurity/es_tcc_authorization_reason_mdm_policy) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING](/documentation/endpointsecurity/es_tcc_authorization_reason_missing_usage_string) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_NONE](/documentation/endpointsecurity/es_tcc_authorization_reason_none)
- [ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN](/documentation/endpointsecurity/es_tcc_authorization_reason_preflight_unknown) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL](/documentation/endpointsecurity/es_tcc_authorization_reason_prompt_cancel) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT](/documentation/endpointsecurity/es_tcc_authorization_reason_prompt_timeout) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY](/documentation/endpointsecurity/es_tcc_authorization_reason_service_override_policy) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY](/documentation/endpointsecurity/es_tcc_authorization_reason_service_policy) - A system process changed the authorization right
- [ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET](/documentation/endpointsecurity/es_tcc_authorization_reason_system_set) - User changed the authorization right via Preferences
- [ES_TCC_AUTHORIZATION_REASON_USER_CONSENT](/documentation/endpointsecurity/es_tcc_authorization_reason_user_consent)
- [ES_TCC_AUTHORIZATION_REASON_USER_SET](/documentation/endpointsecurity/es_tcc_authorization_reason_user_set) - User answered a prompt
- [ES_TCC_AUTHORIZATION_RIGHT_ADD_MODIFY_ADDED](/documentation/endpointsecurity/es_tcc_authorization_right_add_modify_added)
- [ES_TCC_AUTHORIZATION_RIGHT_ALLOWED](/documentation/endpointsecurity/es_tcc_authorization_right_allowed)
- [ES_TCC_AUTHORIZATION_RIGHT_DENIED](/documentation/endpointsecurity/es_tcc_authorization_right_denied)
- [ES_TCC_AUTHORIZATION_RIGHT_LEARN_MORE](/documentation/endpointsecurity/es_tcc_authorization_right_learn_more)
- [ES_TCC_AUTHORIZATION_RIGHT_LIMITED](/documentation/endpointsecurity/es_tcc_authorization_right_limited)
- [ES_TCC_AUTHORIZATION_RIGHT_SESSION_PID](/documentation/endpointsecurity/es_tcc_authorization_right_session_pid)
- [ES_TCC_AUTHORIZATION_RIGHT_UNKNOWN](/documentation/endpointsecurity/es_tcc_authorization_right_unknown)
- [ES_TCC_EVENT_TYPE_CREATE](/documentation/endpointsecurity/es_tcc_event_type_create)
- [ES_TCC_EVENT_TYPE_DELETE](/documentation/endpointsecurity/es_tcc_event_type_delete)
- [ES_TCC_EVENT_TYPE_MODIFY](/documentation/endpointsecurity/es_tcc_event_type_modify)
- [ES_TCC_EVENT_TYPE_UNKNOWN](/documentation/endpointsecurity/es_tcc_event_type_unknown)
- [ES_TCC_IDENTITY_TYPE_BUNDLE_ID](/documentation/endpointsecurity/es_tcc_identity_type_bundle_id)
- [ES_TCC_IDENTITY_TYPE_EXECUTABLE_PATH](/documentation/endpointsecurity/es_tcc_identity_type_executable_path)
- [ES_TCC_IDENTITY_TYPE_FILE_PROVIDER_DOMAIN_ID](/documentation/endpointsecurity/es_tcc_identity_type_file_provider_domain_id)
- [ES_TCC_IDENTITY_TYPE_POLICY_ID](/documentation/endpointsecurity/es_tcc_identity_type_policy_id)

</details>

---

*Generated EndpointSecurity reference*
