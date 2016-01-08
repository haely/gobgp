// Copyright (C) 2014,2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import "fmt"

// typedef for typedef openconfig-types:std-regexp
type StdRegexp string

// typedef for typedef openconfig-types:percentage
type Percentage uint8

// typedef for typedef bgp-types:rr-cluster-id-type
type RrClusterIdType string

// typedef for typedef bgp-types:remove-private-as-option
type RemovePrivateAsOption string

const (
	REMOVE_PRIVATE_AS_OPTION_ALL     RemovePrivateAsOption = "all"
	REMOVE_PRIVATE_AS_OPTION_REPLACE RemovePrivateAsOption = "replace"
)

func (v RemovePrivateAsOption) ToInt() int {
	for i, vv := range []string{"all", "replace"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v RemovePrivateAsOption) FromInt(i int) RemovePrivateAsOption {
	for j, vv := range []string{"all", "replace"} {
		if i == j {
			return RemovePrivateAsOption(vv)
		}
	}
	return RemovePrivateAsOption("")
}

func (v RemovePrivateAsOption) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid RemovePrivateAsOption: %s", v)
	}
	return nil
}

// typedef for typedef bgp-types:bgp-community-regexp-type
type BgpCommunityRegexpType StdRegexp

// typedef for typedef bgp-types:community-type
type CommunityType string

const (
	COMMUNITY_TYPE_STANDARD CommunityType = "standard"
	COMMUNITY_TYPE_EXTENDED CommunityType = "extended"
	COMMUNITY_TYPE_BOTH     CommunityType = "both"
	COMMUNITY_TYPE_NONE     CommunityType = "none"
)

func (v CommunityType) ToInt() int {
	for i, vv := range []string{"standard", "extended", "both", "none"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v CommunityType) FromInt(i int) CommunityType {
	for j, vv := range []string{"standard", "extended", "both", "none"} {
		if i == j {
			return CommunityType(vv)
		}
	}
	return CommunityType("")
}

func (v CommunityType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid CommunityType: %s", v)
	}
	return nil
}

// typedef for typedef bgp-types:bgp-ext-community-type
type BgpExtCommunityType string

// typedef for typedef bgp-types:bgp-std-community-type
type BgpStdCommunityType string

// typedef for typedef bgp-types:peer-type
type PeerTypeDef string

const (
	PEER_TYPE_INTERNAL PeerTypeDef = "internal"
	PEER_TYPE_EXTERNAL PeerTypeDef = "external"
)

func (v PeerTypeDef) ToInt() int {
	for i, vv := range []string{"internal", "external"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v PeerTypeDef) FromInt(i int) PeerTypeDef {
	for j, vv := range []string{"internal", "external"} {
		if i == j {
			return PeerTypeDef(vv)
		}
	}
	return PeerTypeDef("")
}

func (v PeerTypeDef) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid PeerTypeDef: %s", v)
	}
	return nil
}

// typedef for typedef bgp-types:bgp-session-direction
type BgpSessionDirection string

const (
	BGP_SESSION_DIRECTION_INBOUND  BgpSessionDirection = "inbound"
	BGP_SESSION_DIRECTION_OUTBOUND BgpSessionDirection = "outbound"
)

func (v BgpSessionDirection) ToInt() int {
	for i, vv := range []string{"inbound", "outbound"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v BgpSessionDirection) FromInt(i int) BgpSessionDirection {
	for j, vv := range []string{"inbound", "outbound"} {
		if i == j {
			return BgpSessionDirection(vv)
		}
	}
	return BgpSessionDirection("")
}

func (v BgpSessionDirection) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid BgpSessionDirection: %s", v)
	}
	return nil
}

// typedef for typedef ptypes:match-set-options-restricted-type
type MatchSetOptionsRestrictedType string

const (
	MATCH_SET_OPTIONS_RESTRICTED_TYPE_ANY    MatchSetOptionsRestrictedType = "any"
	MATCH_SET_OPTIONS_RESTRICTED_TYPE_INVERT MatchSetOptionsRestrictedType = "invert"
)

func (v MatchSetOptionsRestrictedType) ToInt() int {
	for i, vv := range []string{"any", "invert"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v MatchSetOptionsRestrictedType) FromInt(i int) MatchSetOptionsRestrictedType {
	for j, vv := range []string{"any", "invert"} {
		if i == j {
			return MatchSetOptionsRestrictedType(vv)
		}
	}
	return MatchSetOptionsRestrictedType("")
}

func (v MatchSetOptionsRestrictedType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid MatchSetOptionsRestrictedType: %s", v)
	}
	return nil
}

func (v MatchSetOptionsRestrictedType) Default() MatchSetOptionsRestrictedType {
	return MATCH_SET_OPTIONS_RESTRICTED_TYPE_ANY
}

func (v MatchSetOptionsRestrictedType) DefaultAsNeeded() MatchSetOptionsRestrictedType {
	if string(v) == "" {
		return v.Default()
	}
	return v
}

// typedef for typedef ptypes:match-set-options-type
type MatchSetOptionsType string

const (
	MATCH_SET_OPTIONS_TYPE_ANY    MatchSetOptionsType = "any"
	MATCH_SET_OPTIONS_TYPE_ALL    MatchSetOptionsType = "all"
	MATCH_SET_OPTIONS_TYPE_INVERT MatchSetOptionsType = "invert"
)

func (v MatchSetOptionsType) ToInt() int {
	for i, vv := range []string{"any", "all", "invert"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v MatchSetOptionsType) FromInt(i int) MatchSetOptionsType {
	for j, vv := range []string{"any", "all", "invert"} {
		if i == j {
			return MatchSetOptionsType(vv)
		}
	}
	return MatchSetOptionsType("")
}

func (v MatchSetOptionsType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid MatchSetOptionsType: %s", v)
	}
	return nil
}

func (v MatchSetOptionsType) Default() MatchSetOptionsType {
	return MATCH_SET_OPTIONS_TYPE_ANY
}

func (v MatchSetOptionsType) DefaultAsNeeded() MatchSetOptionsType {
	if string(v) == "" {
		return v.Default()
	}
	return v
}

// typedef for typedef ptypes:tag-type
type TagType string

// typedef for typedef rpol:default-policy-type
type DefaultPolicyType string

const (
	DEFAULT_POLICY_TYPE_ACCEPT_ROUTE DefaultPolicyType = "accept-route"
	DEFAULT_POLICY_TYPE_REJECT_ROUTE DefaultPolicyType = "reject-route"
)

func (v DefaultPolicyType) ToInt() int {
	for i, vv := range []string{"accept-route", "reject-route"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v DefaultPolicyType) FromInt(i int) DefaultPolicyType {
	for j, vv := range []string{"accept-route", "reject-route"} {
		if i == j {
			return DefaultPolicyType(vv)
		}
	}
	return DefaultPolicyType("")
}

func (v DefaultPolicyType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid DefaultPolicyType: %s", v)
	}
	return nil
}

// typedef for typedef bgp-pol:bgp-next-hop-type
type BgpNextHopType string

// typedef for typedef bgp-pol:bgp-as-path-prepend-repeat
type BgpAsPathPrependRepeat uint8

// typedef for typedef bgp-pol:bgp-set-med-type
type BgpSetMedType string

// typedef for typedef bgp-pol:bgp-set-community-option-type
type BgpSetCommunityOptionType string

const (
	BGP_SET_COMMUNITY_OPTION_TYPE_ADD     BgpSetCommunityOptionType = "add"
	BGP_SET_COMMUNITY_OPTION_TYPE_REMOVE  BgpSetCommunityOptionType = "remove"
	BGP_SET_COMMUNITY_OPTION_TYPE_REPLACE BgpSetCommunityOptionType = "replace"
)

func (v BgpSetCommunityOptionType) ToInt() int {
	for i, vv := range []string{"add", "remove", "replace"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v BgpSetCommunityOptionType) FromInt(i int) BgpSetCommunityOptionType {
	for j, vv := range []string{"add", "remove", "replace"} {
		if i == j {
			return BgpSetCommunityOptionType(vv)
		}
	}
	return BgpSetCommunityOptionType("")
}

func (v BgpSetCommunityOptionType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid BgpSetCommunityOptionType: %s", v)
	}
	return nil
}

// typedef for typedef gobgp:bmp-route-monitoring-policy-type
type BmpRouteMonitoringPolicyType string

const (
	BMP_ROUTE_MONITORING_POLICY_TYPE_PRE_POLICY  BmpRouteMonitoringPolicyType = "pre-policy"
	BMP_ROUTE_MONITORING_POLICY_TYPE_POST_POLICY BmpRouteMonitoringPolicyType = "post-policy"
	BMP_ROUTE_MONITORING_POLICY_TYPE_BOTH        BmpRouteMonitoringPolicyType = "both"
)

func (v BmpRouteMonitoringPolicyType) ToInt() int {
	for i, vv := range []string{"pre-policy", "post-policy", "both"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v BmpRouteMonitoringPolicyType) FromInt(i int) BmpRouteMonitoringPolicyType {
	for j, vv := range []string{"pre-policy", "post-policy", "both"} {
		if i == j {
			return BmpRouteMonitoringPolicyType(vv)
		}
	}
	return BmpRouteMonitoringPolicyType("")
}

func (v BmpRouteMonitoringPolicyType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid BmpRouteMonitoringPolicyType: %s", v)
	}
	return nil
}

// typedef for typedef gobgp:rpki-validation-result-type
type RpkiValidationResultType string

const (
	RPKI_VALIDATION_RESULT_TYPE_NONE      RpkiValidationResultType = "none"
	RPKI_VALIDATION_RESULT_TYPE_NOT_FOUND RpkiValidationResultType = "not-found"
	RPKI_VALIDATION_RESULT_TYPE_VALID     RpkiValidationResultType = "valid"
	RPKI_VALIDATION_RESULT_TYPE_INVALID   RpkiValidationResultType = "invalid"
)

func (v RpkiValidationResultType) ToInt() int {
	for i, vv := range []string{"none", "not-found", "valid", "invalid"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v RpkiValidationResultType) FromInt(i int) RpkiValidationResultType {
	for j, vv := range []string{"none", "not-found", "valid", "invalid"} {
		if i == j {
			return RpkiValidationResultType(vv)
		}
	}
	return RpkiValidationResultType("")
}

func (v RpkiValidationResultType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid RpkiValidationResultType: %s", v)
	}
	return nil
}

// typedef for typedef gobgp:bgp-origin-attr-type
type BgpOriginAttrType string

const (
	BGP_ORIGIN_ATTR_TYPE_IGP        BgpOriginAttrType = "igp"
	BGP_ORIGIN_ATTR_TYPE_EGP        BgpOriginAttrType = "egp"
	BGP_ORIGIN_ATTR_TYPE_INCOMPLETE BgpOriginAttrType = "incomplete"
)

func (v BgpOriginAttrType) ToInt() int {
	for i, vv := range []string{"igp", "egp", "incomplete"} {
		if string(v) == vv {
			return i
		}
	}
	return -1
}

func (v BgpOriginAttrType) FromInt(i int) BgpOriginAttrType {
	for j, vv := range []string{"igp", "egp", "incomplete"} {
		if i == j {
			return BgpOriginAttrType(vv)
		}
	}
	return BgpOriginAttrType("")
}

func (v BgpOriginAttrType) Validate() error {
	if v.ToInt() < 0 {
		return fmt.Errorf("invalid BgpOriginAttrType: %s", v)
	}
	return nil
}

//struct for container gobgp:state
type BmpServerState struct {
}

//struct for container gobgp:config
type BmpServerConfig struct {
	// original -> gobgp:address
	//gobgp:address's original type is inet:ip-address
	Address string
	// original -> gobgp:port
	Port uint32
	// original -> gobgp:route-monitoring-policy
	RouteMonitoringPolicy BmpRouteMonitoringPolicyType
}

//struct for container gobgp:bmp-server
type BmpServer struct {
	// original -> gobgp:address
	//gobgp:address's original type is inet:ip-address
	Address string
	// original -> gobgp:bmp-server-config
	Config BmpServerConfig
	// original -> gobgp:bmp-server-state
	State BmpServerState
}

//struct for container gobgp:bmp-servers
type BmpServers struct {
	// original -> gobgp:bmp-server
	BmpServerList []BmpServer
}

//struct for container gobgp:rpki-received
type RpkiReceived struct {
	// original -> gobgp:serial-notify
	SerialNotify int64
	// original -> gobgp:cache-reset
	CacheReset int64
	// original -> gobgp:cache-response
	CacheResponse int64
	// original -> gobgp:ipv4-prefix
	Ipv4Prefix int64
	// original -> gobgp:ipv6-prefix
	Ipv6Prefix int64
	// original -> gobgp:end-of-data
	EndOfData int64
	// original -> gobgp:error
	Error int64
}

//struct for container gobgp:rpki-sent
type RpkiSent struct {
	// original -> gobgp:serial-query
	SerialQuery int64
	// original -> gobgp:reset-query
	ResetQuery int64
	// original -> gobgp:error
	Error int64
}

//struct for container gobgp:rpki-messages
type RpkiMessages struct {
	// original -> gobgp:rpki-sent
	RpkiSent RpkiSent
	// original -> gobgp:rpki-received
	RpkiReceived RpkiReceived
}

//struct for container gobgp:state
type RpkiServerState struct {
	// original -> gobgp:uptime
	Uptime int64
	// original -> gobgp:downtime
	Downtime int64
	// original -> gobgp:last-pdu-recv-time
	LastPduRecvTime int64
	// original -> gobgp:rpki-messages
	RpkiMessages RpkiMessages
}

//struct for container gobgp:config
type RpkiServerConfig struct {
	// original -> gobgp:address
	//gobgp:address's original type is inet:ip-address
	Address string
	// original -> gobgp:port
	Port uint32
	// original -> gobgp:refresh-time
	RefreshTime int64
	// original -> gobgp:hold-time
	HoldTime int64
	// original -> gobgp:record-lifetime
	RecordLifetime int64
	// original -> gobgp:preference
	Preference uint8
}

//struct for container gobgp:rpki-server
type RpkiServer struct {
	// original -> gobgp:address
	//gobgp:address's original type is inet:ip-address
	Address string
	// original -> gobgp:rpki-server-config
	Config RpkiServerConfig
	// original -> gobgp:rpki-server-state
	State RpkiServerState
}

//struct for container gobgp:rpki-servers
type RpkiServers struct {
	// original -> gobgp:rpki-server
	RpkiServerList []RpkiServer
}

//struct for container bgp:state
type PeerGroupState struct {
	// original -> bgp:peer-as
	//bgp:peer-as's original type is inet:as-number
	PeerAs uint32
	// original -> bgp:local-as
	//bgp:local-as's original type is inet:as-number
	LocalAs uint32
	// original -> bgp:peer-type
	PeerType PeerTypeDef
	// original -> bgp:auth-password
	AuthPassword string
	// original -> bgp:remove-private-as
	RemovePrivateAs RemovePrivateAsOption
	// original -> bgp:route-flap-damping
	//bgp:route-flap-damping's original type is boolean
	RouteFlapDamping bool
	// original -> bgp:send-community
	SendCommunity CommunityType
	// original -> bgp:description
	Description string
	// original -> bgp:peer-group-name
	PeerGroupName string
	// original -> bgp-op:total-paths
	TotalPaths uint32
	// original -> bgp-op:total-prefixes
	TotalPrefixes uint32
}

//struct for container bgp:config
type PeerGroupConfig struct {
	// original -> bgp:peer-as
	//bgp:peer-as's original type is inet:as-number
	PeerAs uint32
	// original -> bgp:local-as
	//bgp:local-as's original type is inet:as-number
	LocalAs uint32
	// original -> bgp:peer-type
	PeerType PeerTypeDef
	// original -> bgp:auth-password
	AuthPassword string
	// original -> bgp:remove-private-as
	RemovePrivateAs RemovePrivateAsOption
	// original -> bgp:route-flap-damping
	//bgp:route-flap-damping's original type is boolean
	RouteFlapDamping bool
	// original -> bgp:send-community
	SendCommunity CommunityType
	// original -> bgp:description
	Description string
	// original -> bgp:peer-group-name
	PeerGroupName string
}

//struct for container bgp:peer-group
type PeerGroup struct {
	// original -> bgp:peer-group-name
	PeerGroupName string
	// original -> bgp:peer-group-config
	Config PeerGroupConfig
	// original -> bgp:peer-group-state
	State PeerGroupState
	// original -> bgp:timers
	Timers Timers
	// original -> bgp:transport
	Transport Transport
	// original -> bgp:error-handling
	ErrorHandling ErrorHandling
	// original -> bgp:logging-options
	LoggingOptions LoggingOptions
	// original -> bgp:ebgp-multihop
	EbgpMultihop EbgpMultihop
	// original -> bgp:route-reflector
	RouteReflector RouteReflector
	// original -> bgp:as-path-options
	AsPathOptions AsPathOptions
	// original -> bgp:add-paths
	AddPaths AddPaths
	// original -> bgp:afi-safis
	AfiSafis AfiSafis
	// original -> bgp:graceful-restart
	GracefulRestart GracefulRestart
	// original -> rpol:apply-policy
	ApplyPolicy ApplyPolicy
	// original -> bgp-mp:use-multiple-paths
	UseMultiplePaths UseMultiplePaths
	// original -> gobgp:route-server
	RouteServer RouteServer
}

//struct for container bgp:peer-groups
type PeerGroups struct {
	// original -> bgp:peer-group
	PeerGroupList []PeerGroup
}

//struct for container gobgp:state
type RouteServerState struct {
	// original -> gobgp:route-server-client
	//gobgp:route-server-client's original type is boolean
	RouteServerClient bool
}

//struct for container gobgp:config
type RouteServerConfig struct {
	// original -> gobgp:route-server-client
	//gobgp:route-server-client's original type is boolean
	RouteServerClient bool
}

//struct for container gobgp:route-server
type RouteServer struct {
	// original -> gobgp:route-server-config
	Config RouteServerConfig
	// original -> gobgp:route-server-state
	State RouteServerState
}

//struct for container bgp-op:prefixes
type Prefixes struct {
	// original -> bgp-op:received
	Received uint32
	// original -> bgp-op:sent
	Sent uint32
	// original -> bgp-op:installed
	Installed uint32
}

//struct for container bgp:state
type AddPathsState struct {
	// original -> bgp:receive
	//bgp:receive's original type is boolean
	Receive bool
	// original -> bgp:send-max
	SendMax uint8
}

//struct for container bgp:config
type AddPathsConfig struct {
	// original -> bgp:receive
	//bgp:receive's original type is boolean
	Receive bool
	// original -> bgp:send-max
	SendMax uint8
}

//struct for container bgp:add-paths
type AddPaths struct {
	// original -> bgp:add-paths-config
	Config AddPathsConfig
	// original -> bgp:add-paths-state
	State AddPathsState
}

//struct for container bgp:state
type AsPathOptionsState struct {
	// original -> bgp:allow-own-as
	AllowOwnAs uint8
	// original -> bgp:replace-peer-as
	//bgp:replace-peer-as's original type is boolean
	ReplacePeerAs bool
}

//struct for container bgp:config
type AsPathOptionsConfig struct {
	// original -> bgp:allow-own-as
	AllowOwnAs uint8
	// original -> bgp:replace-peer-as
	//bgp:replace-peer-as's original type is boolean
	ReplacePeerAs bool
}

//struct for container bgp:as-path-options
type AsPathOptions struct {
	// original -> bgp:as-path-options-config
	Config AsPathOptionsConfig
	// original -> bgp:as-path-options-state
	State AsPathOptionsState
}

//struct for container bgp:state
type RouteReflectorState struct {
	// original -> bgp:route-reflector-cluster-id
	RouteReflectorClusterId RrClusterIdType
	// original -> bgp:route-reflector-client
	//bgp:route-reflector-client's original type is boolean
	RouteReflectorClient bool
}

//struct for container bgp:config
type RouteReflectorConfig struct {
	// original -> bgp:route-reflector-cluster-id
	RouteReflectorClusterId RrClusterIdType
	// original -> bgp:route-reflector-client
	//bgp:route-reflector-client's original type is boolean
	RouteReflectorClient bool
}

//struct for container bgp:route-reflector
type RouteReflector struct {
	// original -> bgp:route-reflector-config
	Config RouteReflectorConfig
	// original -> bgp:route-reflector-state
	State RouteReflectorState
}

//struct for container bgp:state
type EbgpMultihopState struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:multihop-ttl
	MultihopTtl uint8
}

//struct for container bgp:config
type EbgpMultihopConfig struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:multihop-ttl
	MultihopTtl uint8
}

//struct for container bgp:ebgp-multihop
type EbgpMultihop struct {
	// original -> bgp:ebgp-multihop-config
	Config EbgpMultihopConfig
	// original -> bgp:ebgp-multihop-state
	State EbgpMultihopState
}

//struct for container bgp:state
type LoggingOptionsState struct {
	// original -> bgp:log-neighbor-state-changes
	//bgp:log-neighbor-state-changes's original type is boolean
	LogNeighborStateChanges bool
}

//struct for container bgp:config
type LoggingOptionsConfig struct {
	// original -> bgp:log-neighbor-state-changes
	//bgp:log-neighbor-state-changes's original type is boolean
	LogNeighborStateChanges bool
}

//struct for container bgp:logging-options
type LoggingOptions struct {
	// original -> bgp:logging-options-config
	Config LoggingOptionsConfig
	// original -> bgp:logging-options-state
	State LoggingOptionsState
}

//struct for container bgp:state
type ErrorHandlingState struct {
	// original -> bgp:treat-as-withdraw
	//bgp:treat-as-withdraw's original type is boolean
	TreatAsWithdraw bool
	// original -> bgp-op:erroneous-update-messages
	ErroneousUpdateMessages uint32
}

//struct for container bgp:config
type ErrorHandlingConfig struct {
	// original -> bgp:treat-as-withdraw
	//bgp:treat-as-withdraw's original type is boolean
	TreatAsWithdraw bool
}

//struct for container bgp:error-handling
type ErrorHandling struct {
	// original -> bgp:error-handling-config
	Config ErrorHandlingConfig
	// original -> bgp:error-handling-state
	State ErrorHandlingState
}

//struct for container bgp:state
type TransportState struct {
	// original -> bgp:tcp-mss
	TcpMss uint16
	// original -> bgp:mtu-discovery
	//bgp:mtu-discovery's original type is boolean
	MtuDiscovery bool
	// original -> bgp:passive-mode
	//bgp:passive-mode's original type is boolean
	PassiveMode bool
	// original -> bgp-op:local-port
	//bgp-op:local-port's original type is inet:port-number
	LocalPort uint16
	// original -> bgp-op:remote-address
	//bgp-op:remote-address's original type is inet:ip-address
	RemoteAddress string
	// original -> bgp-op:remote-port
	//bgp-op:remote-port's original type is inet:port-number
	RemotePort uint16
	// original -> gobgp:local-address
	//gobgp:local-address's original type is inet:ip-address
	LocalAddress string
}

//struct for container bgp:config
type TransportConfig struct {
	// original -> bgp:tcp-mss
	TcpMss uint16
	// original -> bgp:mtu-discovery
	//bgp:mtu-discovery's original type is boolean
	MtuDiscovery bool
	// original -> bgp:passive-mode
	//bgp:passive-mode's original type is boolean
	PassiveMode bool
	// original -> gobgp:local-address
	//gobgp:local-address's original type is inet:ip-address
	LocalAddress string
}

//struct for container bgp:transport
type Transport struct {
	// original -> bgp:transport-config
	Config TransportConfig
	// original -> bgp:transport-state
	State TransportState
}

//struct for container bgp:state
type TimersState struct {
	// original -> bgp:connect-retry
	//bgp:connect-retry's original type is decimal64
	ConnectRetry float64
	// original -> bgp:hold-time
	//bgp:hold-time's original type is decimal64
	HoldTime float64
	// original -> bgp:keepalive-interval
	//bgp:keepalive-interval's original type is decimal64
	KeepaliveInterval float64
	// original -> bgp:minimum-advertisement-interval
	//bgp:minimum-advertisement-interval's original type is decimal64
	MinimumAdvertisementInterval float64
	// original -> bgp-op:uptime
	//bgp-op:uptime's original type is yang:timeticks
	Uptime int64
	// original -> bgp-op:negotiated-hold-time
	//bgp-op:negotiated-hold-time's original type is decimal64
	NegotiatedHoldTime float64
	// original -> gobgp:idle-hold-time-after-reset
	//gobgp:idle-hold-time-after-reset's original type is decimal64
	IdleHoldTimeAfterReset float64
	// original -> gobgp:downtime
	//gobgp:downtime's original type is yang:timeticks
	Downtime int64
	// original -> gobgp:update-recv-time
	UpdateRecvTime int64
}

//struct for container bgp:config
type TimersConfig struct {
	// original -> bgp:connect-retry
	//bgp:connect-retry's original type is decimal64
	ConnectRetry float64
	// original -> bgp:hold-time
	//bgp:hold-time's original type is decimal64
	HoldTime float64
	// original -> bgp:keepalive-interval
	//bgp:keepalive-interval's original type is decimal64
	KeepaliveInterval float64
	// original -> bgp:minimum-advertisement-interval
	//bgp:minimum-advertisement-interval's original type is decimal64
	MinimumAdvertisementInterval float64
	// original -> gobgp:idle-hold-time-after-reset
	//gobgp:idle-hold-time-after-reset's original type is decimal64
	IdleHoldTimeAfterReset float64
}

//struct for container bgp:timers
type Timers struct {
	// original -> bgp:timers-config
	Config TimersConfig
	// original -> bgp:timers-state
	State TimersState
}

//struct for container bgp:queues
type Queues struct {
	// original -> bgp-op:input
	Input uint32
	// original -> bgp-op:output
	Output uint32
}

//struct for container bgp:received
type Received struct {
	// original -> bgp-op:UPDATE
	Update uint64
	// original -> bgp-op:NOTIFICATION
	Notification uint64
	// original -> gobgp:OPEN
	Open uint64
	// original -> gobgp:REFRESH
	Refresh uint64
	// original -> gobgp:KEEPALIVE
	Keepalive uint64
	// original -> gobgp:DYNAMIC-CAP
	DynamicCap uint64
	// original -> gobgp:DISCARDED
	Discarded uint64
	// original -> gobgp:TOTAL
	Total uint64
}

//struct for container bgp:sent
type Sent struct {
	// original -> bgp-op:UPDATE
	Update uint64
	// original -> bgp-op:NOTIFICATION
	Notification uint64
	// original -> gobgp:OPEN
	Open uint64
	// original -> gobgp:REFRESH
	Refresh uint64
	// original -> gobgp:KEEPALIVE
	Keepalive uint64
	// original -> gobgp:DYNAMIC-CAP
	DynamicCap uint64
	// original -> gobgp:DISCARDED
	Discarded uint64
	// original -> gobgp:TOTAL
	Total uint64
}

//struct for container bgp:messages
type Messages struct {
	// original -> bgp:sent
	Sent Sent
	// original -> bgp:received
	Received Received
}

//struct for container bgp:state
type NeighborState struct {
	// original -> bgp:peer-as
	//bgp:peer-as's original type is inet:as-number
	PeerAs uint32
	// original -> bgp:local-as
	//bgp:local-as's original type is inet:as-number
	LocalAs uint32
	// original -> bgp:peer-type
	PeerType PeerTypeDef
	// original -> bgp:auth-password
	AuthPassword string
	// original -> bgp:remove-private-as
	RemovePrivateAs RemovePrivateAsOption
	// original -> bgp:route-flap-damping
	//bgp:route-flap-damping's original type is boolean
	RouteFlapDamping bool
	// original -> bgp:send-community
	SendCommunity CommunityType
	// original -> bgp:description
	Description string
	// original -> bgp:peer-group
	PeerGroup string
	// original -> bgp:neighbor-address
	//bgp:neighbor-address's original type is inet:ip-address
	NeighborAddress string
	// original -> bgp-op:session-state
	//bgp-op:session-state's original type is enumeration
	SessionState uint32
	// original -> bgp-op:supported-capabilities
	//original type is list of identityref
	SupportedCapabilities []string
	// original -> bgp:messages
	Messages Messages
	// original -> bgp:queues
	Queues Queues
	// original -> gobgp:admin-down
	//gobgp:admin-down's original type is boolean
	AdminDown bool
	// original -> gobgp:established-count
	EstablishedCount uint32
	// original -> gobgp:flops
	Flops uint32
}

//struct for container bgp:config
type NeighborConfig struct {
	// original -> bgp:peer-as
	//bgp:peer-as's original type is inet:as-number
	PeerAs uint32
	// original -> bgp:local-as
	//bgp:local-as's original type is inet:as-number
	LocalAs uint32
	// original -> bgp:peer-type
	PeerType PeerTypeDef
	// original -> bgp:auth-password
	AuthPassword string
	// original -> bgp:remove-private-as
	RemovePrivateAs RemovePrivateAsOption
	// original -> bgp:route-flap-damping
	//bgp:route-flap-damping's original type is boolean
	RouteFlapDamping bool
	// original -> bgp:send-community
	SendCommunity CommunityType
	// original -> bgp:description
	Description string
	// original -> bgp:peer-group
	PeerGroup string
	// original -> bgp:neighbor-address
	//bgp:neighbor-address's original type is inet:ip-address
	NeighborAddress string
}

//struct for container bgp:neighbor
type Neighbor struct {
	// original -> bgp:neighbor-address
	//bgp:neighbor-address's original type is inet:ip-address
	NeighborAddress string
	// original -> bgp:neighbor-config
	Config NeighborConfig
	// original -> bgp:neighbor-state
	State NeighborState
	// original -> bgp:timers
	Timers Timers
	// original -> bgp:transport
	Transport Transport
	// original -> bgp:error-handling
	ErrorHandling ErrorHandling
	// original -> bgp:logging-options
	LoggingOptions LoggingOptions
	// original -> bgp:ebgp-multihop
	EbgpMultihop EbgpMultihop
	// original -> bgp:route-reflector
	RouteReflector RouteReflector
	// original -> bgp:as-path-options
	AsPathOptions AsPathOptions
	// original -> bgp:add-paths
	AddPaths AddPaths
	// original -> bgp:afi-safis
	AfiSafis AfiSafis
	// original -> bgp:graceful-restart
	GracefulRestart GracefulRestart
	// original -> rpol:apply-policy
	ApplyPolicy ApplyPolicy
	// original -> bgp-mp:use-multiple-paths
	UseMultiplePaths UseMultiplePaths
	// original -> gobgp:route-server
	RouteServer RouteServer
}

//struct for container bgp:neighbors
type Neighbors struct {
	// original -> bgp:neighbor
	NeighborList []Neighbor
}

//struct for container gobgp:listen-config
type ListenConfig struct {
	// original -> gobgp:port
	Port int32
}

//struct for container gobgp:mpls-label-range
type MplsLabelRange struct {
	// original -> gobgp:min-label
	MinLabel uint32
	// original -> gobgp:max-label
	MaxLabel uint32
}

//struct for container gobgp:zebra
type Zebra struct {
	// original -> gobgp:enabled
	//gobgp:enabled's original type is boolean
	Enabled bool
	// original -> gobgp:url
	Url string
	// original -> gobgp:redistribute-route-type
	//original type is list of identityref
	RedistributeRouteType []string
}

//struct for container gobgp:mrt
type Mrt struct {
	// original -> gobgp:file-name
	FileName string
}

//struct for container bgp-mp:l2vpn-evpn
type L2vpnEvpn struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:l2vpn-vpls
type L2vpnVpls struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:l3vpn-ipv6-multicast
type L3vpnIpv6Multicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:l3vpn-ipv4-multicast
type L3vpnIpv4Multicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:l3vpn-ipv6-unicast
type L3vpnIpv6Unicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:l3vpn-ipv4-unicast
type L3vpnIpv4Unicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:ipv6-labelled-unicast
type Ipv6LabelledUnicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:ipv4-labelled-unicast
type Ipv4LabelledUnicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
}

//struct for container bgp-mp:state
type Ipv6UnicastState struct {
	// original -> bgp-mp:send-default-route
	//bgp-mp:send-default-route's original type is boolean
	SendDefaultRoute bool
}

//struct for container bgp-mp:config
type Ipv6UnicastConfig struct {
	// original -> bgp-mp:send-default-route
	//bgp-mp:send-default-route's original type is boolean
	SendDefaultRoute bool
}

//struct for container bgp-mp:ipv6-unicast
type Ipv6Unicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
	// original -> bgp-mp:ipv6-unicast-config
	Config Ipv6UnicastConfig
	// original -> bgp-mp:ipv6-unicast-state
	State Ipv6UnicastState
}

//struct for container bgp-mp:state
type Ipv4UnicastState struct {
	// original -> bgp-mp:send-default-route
	//bgp-mp:send-default-route's original type is boolean
	SendDefaultRoute bool
}

//struct for container bgp-mp:config
type Ipv4UnicastConfig struct {
	// original -> bgp-mp:send-default-route
	//bgp-mp:send-default-route's original type is boolean
	SendDefaultRoute bool
}

//struct for container bgp-mp:state
type PrefixLimitState struct {
	// original -> bgp-mp:max-prefixes
	MaxPrefixes uint32
	// original -> bgp-mp:shutdown-threshold-pct
	ShutdownThresholdPct Percentage
	// original -> bgp-mp:restart-timer
	//bgp-mp:restart-timer's original type is decimal64
	RestartTimer float64
}

//struct for container bgp-mp:config
type PrefixLimitConfig struct {
	// original -> bgp-mp:max-prefixes
	MaxPrefixes uint32
	// original -> bgp-mp:shutdown-threshold-pct
	ShutdownThresholdPct Percentage
	// original -> bgp-mp:restart-timer
	//bgp-mp:restart-timer's original type is decimal64
	RestartTimer float64
}

//struct for container bgp-mp:prefix-limit
type PrefixLimit struct {
	// original -> bgp-mp:prefix-limit-config
	Config PrefixLimitConfig
	// original -> bgp-mp:prefix-limit-state
	State PrefixLimitState
}

//struct for container bgp-mp:ipv4-unicast
type Ipv4Unicast struct {
	// original -> bgp-mp:prefix-limit
	PrefixLimit PrefixLimit
	// original -> bgp-mp:ipv4-unicast-config
	Config Ipv4UnicastConfig
	// original -> bgp-mp:ipv4-unicast-state
	State Ipv4UnicastState
}

//struct for container rpol:state
type ApplyPolicyState struct {
	// original -> rpol:import-policy
	ImportPolicy []string
	// original -> rpol:default-import-policy
	DefaultImportPolicy DefaultPolicyType
	// original -> rpol:export-policy
	ExportPolicy []string
	// original -> rpol:default-export-policy
	DefaultExportPolicy DefaultPolicyType
	// original -> gobgp:in-policy
	InPolicy []string
	// original -> gobgp:default-in-policy
	DefaultInPolicy DefaultPolicyType
}

//struct for container rpol:config
type ApplyPolicyConfig struct {
	// original -> rpol:import-policy
	ImportPolicy []string
	// original -> rpol:default-import-policy
	DefaultImportPolicy DefaultPolicyType
	// original -> rpol:export-policy
	ExportPolicy []string
	// original -> rpol:default-export-policy
	DefaultExportPolicy DefaultPolicyType
	// original -> gobgp:in-policy
	InPolicy []string
	// original -> gobgp:default-in-policy
	DefaultInPolicy DefaultPolicyType
}

//struct for container rpol:apply-policy
type ApplyPolicy struct {
	// original -> rpol:apply-policy-config
	Config ApplyPolicyConfig
	// original -> rpol:apply-policy-state
	State ApplyPolicyState
}

//struct for container bgp-mp:state
type AfiSafiState struct {
	// original -> bgp-mp:afi-safi-name
	AfiSafiName string
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
	// original -> bgp-op:total-paths
	TotalPaths uint32
	// original -> bgp-op:total-prefixes
	TotalPrefixes uint32
}

//struct for container bgp-mp:config
type AfiSafiConfig struct {
	// original -> bgp-mp:afi-safi-name
	AfiSafiName string
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
}

//struct for container bgp-mp:state
type MpGracefulRestartState struct {
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
	// original -> bgp-op:received
	//bgp-op:received's original type is boolean
	Received bool
	// original -> bgp-op:advertised
	//bgp-op:advertised's original type is boolean
	Advertised bool
}

//struct for container bgp-mp:config
type MpGracefulRestartConfig struct {
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
}

//struct for container bgp-mp:graceful-restart
type MpGracefulRestart struct {
	// original -> bgp-mp:mp-graceful-restart-config
	Config MpGracefulRestartConfig
	// original -> bgp-mp:mp-graceful-restart-state
	State MpGracefulRestartState
}

//struct for container bgp-mp:afi-safi
type AfiSafi struct {
	// original -> bgp-mp:afi-safi-name
	//bgp-mp:afi-safi-name's original type is identityref
	AfiSafiName string
	// original -> bgp-mp:mp-graceful-restart
	MpGracefulRestart MpGracefulRestart
	// original -> bgp-mp:afi-safi-config
	Config AfiSafiConfig
	// original -> bgp-mp:afi-safi-state
	State AfiSafiState
	// original -> rpol:apply-policy
	ApplyPolicy ApplyPolicy
	// original -> bgp-mp:ipv4-unicast
	Ipv4Unicast Ipv4Unicast
	// original -> bgp-mp:ipv6-unicast
	Ipv6Unicast Ipv6Unicast
	// original -> bgp-mp:ipv4-labelled-unicast
	Ipv4LabelledUnicast Ipv4LabelledUnicast
	// original -> bgp-mp:ipv6-labelled-unicast
	Ipv6LabelledUnicast Ipv6LabelledUnicast
	// original -> bgp-mp:l3vpn-ipv4-unicast
	L3vpnIpv4Unicast L3vpnIpv4Unicast
	// original -> bgp-mp:l3vpn-ipv6-unicast
	L3vpnIpv6Unicast L3vpnIpv6Unicast
	// original -> bgp-mp:l3vpn-ipv4-multicast
	L3vpnIpv4Multicast L3vpnIpv4Multicast
	// original -> bgp-mp:l3vpn-ipv6-multicast
	L3vpnIpv6Multicast L3vpnIpv6Multicast
	// original -> bgp-mp:l2vpn-vpls
	L2vpnVpls L2vpnVpls
	// original -> bgp-mp:l2vpn-evpn
	L2vpnEvpn L2vpnEvpn
	// original -> bgp-mp:route-selection-options
	RouteSelectionOptions RouteSelectionOptions
	// original -> bgp-mp:use-multiple-paths
	UseMultiplePaths UseMultiplePaths
}

//struct for container bgp:afi-safis
type AfiSafis struct {
	// original -> bgp-mp:afi-safi
	AfiSafiList []AfiSafi
}

//struct for container bgp:state
type GracefulRestartState struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:restart-time
	RestartTime uint16
	// original -> bgp:stale-routes-time
	//bgp:stale-routes-time's original type is decimal64
	StaleRoutesTime float64
	// original -> bgp:helper-only
	//bgp:helper-only's original type is boolean
	HelperOnly bool
	// original -> bgp-op:peer-restart-time
	PeerRestartTime uint16
	// original -> bgp-op:peer-restarting
	//bgp-op:peer-restarting's original type is boolean
	PeerRestarting bool
	// original -> bgp-op:local-restarting
	//bgp-op:local-restarting's original type is boolean
	LocalRestarting bool
	// original -> bgp-op:mode
	//bgp-op:mode's original type is enumeration
	Mode uint32
}

//struct for container bgp:config
type GracefulRestartConfig struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:restart-time
	RestartTime uint16
	// original -> bgp:stale-routes-time
	//bgp:stale-routes-time's original type is decimal64
	StaleRoutesTime float64
	// original -> bgp:helper-only
	//bgp:helper-only's original type is boolean
	HelperOnly bool
}

//struct for container bgp:graceful-restart
type GracefulRestart struct {
	// original -> bgp:graceful-restart-config
	Config GracefulRestartConfig
	// original -> bgp:graceful-restart-state
	State GracefulRestartState
}

//struct for container bgp-mp:state
type IbgpState struct {
	// original -> bgp-mp:maximum-paths
	MaximumPaths uint32
}

//struct for container bgp-mp:config
type IbgpConfig struct {
	// original -> bgp-mp:maximum-paths
	MaximumPaths uint32
}

//struct for container bgp-mp:ibgp
type Ibgp struct {
	// original -> bgp-mp:ibgp-config
	Config IbgpConfig
	// original -> bgp-mp:ibgp-state
	State IbgpState
}

//struct for container bgp-mp:state
type EbgpState struct {
	// original -> bgp-mp:allow-multiple-as
	//bgp-mp:allow-multiple-as's original type is boolean
	AllowMultipleAs bool
	// original -> bgp-mp:maximum-paths
	MaximumPaths uint32
}

//struct for container bgp-mp:config
type EbgpConfig struct {
	// original -> bgp-mp:allow-multiple-as
	//bgp-mp:allow-multiple-as's original type is boolean
	AllowMultipleAs bool
	// original -> bgp-mp:maximum-paths
	MaximumPaths uint32
}

//struct for container bgp-mp:ebgp
type Ebgp struct {
	// original -> bgp-mp:ebgp-config
	Config EbgpConfig
	// original -> bgp-mp:ebgp-state
	State EbgpState
}

//struct for container bgp-mp:state
type UseMultiplePathsState struct {
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
}

//struct for container bgp-mp:config
type UseMultiplePathsConfig struct {
	// original -> bgp-mp:enabled
	//bgp-mp:enabled's original type is boolean
	Enabled bool
}

//struct for container bgp-mp:use-multiple-paths
type UseMultiplePaths struct {
	// original -> bgp-mp:use-multiple-paths-config
	Config UseMultiplePathsConfig
	// original -> bgp-mp:use-multiple-paths-state
	State UseMultiplePathsState
	// original -> bgp-mp:ebgp
	Ebgp Ebgp
	// original -> bgp-mp:ibgp
	Ibgp Ibgp
}

//struct for container bgp:state
type ConfederationState struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:identifier
	//bgp:identifier's original type is inet:as-number
	Identifier uint32
	// original -> bgp:member-as
	//original type is list of inet:as-number
	MemberAs []uint32
}

//struct for container bgp:config
type ConfederationConfig struct {
	// original -> bgp:enabled
	//bgp:enabled's original type is boolean
	Enabled bool
	// original -> bgp:identifier
	//bgp:identifier's original type is inet:as-number
	Identifier uint32
	// original -> bgp:member-as
	//original type is list of inet:as-number
	MemberAs []uint32
}

//struct for container bgp:confederation
type Confederation struct {
	// original -> bgp:confederation-config
	Config ConfederationConfig
	// original -> bgp:confederation-state
	State ConfederationState
}

//struct for container bgp:state
type DefaultRouteDistanceState struct {
	// original -> bgp:external-route-distance
	ExternalRouteDistance uint8
	// original -> bgp:internal-route-distance
	InternalRouteDistance uint8
}

//struct for container bgp:config
type DefaultRouteDistanceConfig struct {
	// original -> bgp:external-route-distance
	ExternalRouteDistance uint8
	// original -> bgp:internal-route-distance
	InternalRouteDistance uint8
}

//struct for container bgp:default-route-distance
type DefaultRouteDistance struct {
	// original -> bgp:default-route-distance-config
	Config DefaultRouteDistanceConfig
	// original -> bgp:default-route-distance-state
	State DefaultRouteDistanceState
}

//struct for container bgp-mp:state
type RouteSelectionOptionsState struct {
	// original -> bgp-mp:always-compare-med
	//bgp-mp:always-compare-med's original type is boolean
	AlwaysCompareMed bool
	// original -> bgp-mp:ignore-as-path-length
	//bgp-mp:ignore-as-path-length's original type is boolean
	IgnoreAsPathLength bool
	// original -> bgp-mp:external-compare-router-id
	//bgp-mp:external-compare-router-id's original type is boolean
	ExternalCompareRouterId bool
	// original -> bgp-mp:advertise-inactive-routes
	//bgp-mp:advertise-inactive-routes's original type is boolean
	AdvertiseInactiveRoutes bool
	// original -> bgp-mp:enable-aigp
	//bgp-mp:enable-aigp's original type is boolean
	EnableAigp bool
	// original -> bgp-mp:ignore-next-hop-igp-metric
	//bgp-mp:ignore-next-hop-igp-metric's original type is boolean
	IgnoreNextHopIgpMetric bool
}

//struct for container bgp-mp:config
type RouteSelectionOptionsConfig struct {
	// original -> bgp-mp:always-compare-med
	//bgp-mp:always-compare-med's original type is boolean
	AlwaysCompareMed bool
	// original -> bgp-mp:ignore-as-path-length
	//bgp-mp:ignore-as-path-length's original type is boolean
	IgnoreAsPathLength bool
	// original -> bgp-mp:external-compare-router-id
	//bgp-mp:external-compare-router-id's original type is boolean
	ExternalCompareRouterId bool
	// original -> bgp-mp:advertise-inactive-routes
	//bgp-mp:advertise-inactive-routes's original type is boolean
	AdvertiseInactiveRoutes bool
	// original -> bgp-mp:enable-aigp
	//bgp-mp:enable-aigp's original type is boolean
	EnableAigp bool
	// original -> bgp-mp:ignore-next-hop-igp-metric
	//bgp-mp:ignore-next-hop-igp-metric's original type is boolean
	IgnoreNextHopIgpMetric bool
}

//struct for container bgp-mp:route-selection-options
type RouteSelectionOptions struct {
	// original -> bgp-mp:route-selection-options-config
	Config RouteSelectionOptionsConfig
	// original -> bgp-mp:route-selection-options-state
	State RouteSelectionOptionsState
}

//struct for container bgp:state
type GlobalState struct {
	// original -> bgp:as
	//bgp:as's original type is inet:as-number
	As uint32
	// original -> bgp:router-id
	//bgp:router-id's original type is inet:ipv4-address
	RouterId string
	// original -> bgp-op:total-paths
	TotalPaths uint32
	// original -> bgp-op:total-prefixes
	TotalPrefixes uint32
}

//struct for container bgp:config
type GlobalConfig struct {
	// original -> bgp:as
	//bgp:as's original type is inet:as-number
	As uint32
	// original -> bgp:router-id
	//bgp:router-id's original type is inet:ipv4-address
	RouterId string
}

//struct for container bgp:global
type Global struct {
	// original -> bgp:global-config
	Config GlobalConfig
	// original -> bgp:global-state
	State GlobalState
	// original -> bgp-mp:route-selection-options
	RouteSelectionOptions RouteSelectionOptions
	// original -> bgp:default-route-distance
	DefaultRouteDistance DefaultRouteDistance
	// original -> bgp:confederation
	Confederation Confederation
	// original -> bgp-mp:use-multiple-paths
	UseMultiplePaths UseMultiplePaths
	// original -> bgp:graceful-restart
	GracefulRestart GracefulRestart
	// original -> bgp:afi-safis
	AfiSafis AfiSafis
	// original -> rpol:apply-policy
	ApplyPolicy ApplyPolicy
	// original -> gobgp:mrt
	Mrt Mrt
	// original -> gobgp:zebra
	Zebra Zebra
	// original -> gobgp:mpls-label-range
	MplsLabelRange MplsLabelRange
	// original -> gobgp:listen-config
	ListenConfig ListenConfig
}

//struct for container bgp:bgp
type Bgp struct {
	// original -> bgp:global
	Global Global
	// original -> bgp:neighbors
	Neighbors Neighbors
	// original -> bgp:peer-groups
	PeerGroups PeerGroups
	// original -> gobgp:rpki-servers
	RpkiServers RpkiServers
	// original -> gobgp:bmp-servers
	BmpServers BmpServers
}

//struct for container bgp-pol:set-ext-community-method
type SetExtCommunityMethod struct {
	// original -> bgp-pol:communities
	//original type is list of union
	Communities []string
	// original -> bgp-pol:ext-community-set-ref
	ExtCommunitySetRef string
}

//struct for container bgp-pol:set-ext-community
type SetExtCommunity struct {
	// original -> bgp-pol:set-ext-community-method
	SetExtCommunityMethod SetExtCommunityMethod
	// original -> bgp-pol:options
	//bgp-pol:options's original type is bgp-set-community-option-type
	Options string
}

//struct for container bgp-pol:set-community-method
type SetCommunityMethod struct {
	// original -> bgp-pol:communities
	//original type is list of union
	Communities []string
	// original -> bgp-pol:community-set-ref
	CommunitySetRef string
}

//struct for container bgp-pol:set-community
type SetCommunity struct {
	// original -> bgp-pol:set-community-method
	SetCommunityMethod SetCommunityMethod
	// original -> bgp-pol:options
	//bgp-pol:options's original type is bgp-set-community-option-type
	Options string
}

//struct for container bgp-pol:set-as-path-prepend
type SetAsPathPrepend struct {
	// original -> bgp-pol:repeat-n
	RepeatN uint8
	// original -> gobgp:as
	//gobgp:as's original type is union
	As string
}

//struct for container bgp-pol:bgp-actions
type BgpActions struct {
	// original -> bgp-pol:set-as-path-prepend
	SetAsPathPrepend SetAsPathPrepend
	// original -> bgp-pol:set-community
	SetCommunity SetCommunity
	// original -> bgp-pol:set-ext-community
	SetExtCommunity SetExtCommunity
	// original -> bgp-pol:set-route-origin
	SetRouteOrigin BgpOriginAttrType
	// original -> bgp-pol:set-local-pref
	SetLocalPref uint32
	// original -> bgp-pol:set-next-hop
	SetNextHop BgpNextHopType
	// original -> bgp-pol:set-med
	SetMed BgpSetMedType
}

//struct for container rpol:igp-actions
type IgpActions struct {
	// original -> rpol:set-tag
	SetTag TagType
}

//struct for container rpol:route-disposition
type RouteDisposition struct {
	// original -> rpol:accept-route
	//rpol:accept-route's original type is empty
	AcceptRoute bool
	// original -> rpol:reject-route
	//rpol:reject-route's original type is empty
	RejectRoute bool
}

//struct for container rpol:actions
type Actions struct {
	// original -> rpol:route-disposition
	RouteDisposition RouteDisposition
	// original -> rpol:igp-actions
	IgpActions IgpActions
	// original -> bgp-pol:bgp-actions
	BgpActions BgpActions
}

//struct for container bgp-pol:as-path-length
type AsPathLength struct {
	// original -> ptypes:operator
	Operator string
	// original -> ptypes:value
	Value uint32
}

//struct for container bgp-pol:community-count
type CommunityCount struct {
	// original -> ptypes:operator
	Operator string
	// original -> ptypes:value
	Value uint32
}

//struct for container bgp-pol:match-as-path-set
type MatchAsPathSet struct {
	// original -> bgp-pol:as-path-set
	AsPathSet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsType
}

//struct for container bgp-pol:match-ext-community-set
type MatchExtCommunitySet struct {
	// original -> bgp-pol:ext-community-set
	ExtCommunitySet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsType
}

//struct for container bgp-pol:match-community-set
type MatchCommunitySet struct {
	// original -> bgp-pol:community-set
	CommunitySet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsType
}

//struct for container bgp-pol:bgp-conditions
type BgpConditions struct {
	// original -> bgp-pol:match-community-set
	MatchCommunitySet MatchCommunitySet
	// original -> bgp-pol:match-ext-community-set
	MatchExtCommunitySet MatchExtCommunitySet
	// original -> bgp-pol:match-as-path-set
	MatchAsPathSet MatchAsPathSet
	// original -> bgp-pol:med-eq
	MedEq uint32
	// original -> bgp-pol:origin-eq
	OriginEq BgpOriginAttrType
	// original -> bgp-pol:next-hop-in
	//original type is list of inet:ip-address
	NextHopIn []string
	// original -> bgp-pol:afi-safi-in
	//original type is list of identityref
	AfiSafiIn []string
	// original -> bgp-pol:local-pref-eq
	LocalPrefEq uint32
	// original -> bgp-pol:community-count
	CommunityCount CommunityCount
	// original -> bgp-pol:as-path-length
	AsPathLength AsPathLength
	// original -> bgp-pol:route-type
	//bgp-pol:route-type's original type is enumeration
	RouteType uint32
	// original -> gobgp:rpki-validation-result
	RpkiValidationResult RpkiValidationResultType
}

//struct for container rpol:igp-conditions
type IgpConditions struct {
}

//struct for container rpol:match-tag-set
type MatchTagSet struct {
	// original -> rpol:tag-set
	TagSet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsRestrictedType
}

//struct for container rpol:match-neighbor-set
type MatchNeighborSet struct {
	// original -> rpol:neighbor-set
	NeighborSet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsRestrictedType
}

//struct for container rpol:match-prefix-set
type MatchPrefixSet struct {
	// original -> rpol:prefix-set
	PrefixSet string
	// original -> rpol:match-set-options
	MatchSetOptions MatchSetOptionsRestrictedType
}

//struct for container rpol:conditions
type Conditions struct {
	// original -> rpol:call-policy
	CallPolicy string
	// original -> rpol:match-prefix-set
	MatchPrefixSet MatchPrefixSet
	// original -> rpol:match-neighbor-set
	MatchNeighborSet MatchNeighborSet
	// original -> rpol:match-tag-set
	MatchTagSet MatchTagSet
	// original -> rpol:install-protocol-eq
	InstallProtocolEq string
	// original -> rpol:igp-conditions
	IgpConditions IgpConditions
	// original -> bgp-pol:bgp-conditions
	BgpConditions BgpConditions
}

//struct for container rpol:statement
type Statement struct {
	// original -> rpol:name
	Name string
	// original -> rpol:conditions
	Conditions Conditions
	// original -> rpol:actions
	Actions Actions
}

//struct for container rpol:statements
type Statements struct {
	// original -> rpol:statement
	StatementList []Statement
}

//struct for container rpol:policy-definition
type PolicyDefinition struct {
	// original -> rpol:name
	Name string
	// original -> rpol:statements
	Statements Statements
}

//struct for container rpol:policy-definitions
type PolicyDefinitions struct {
	// original -> rpol:policy-definition
	PolicyDefinitionList []PolicyDefinition
}

//struct for container bgp-pol:as-path-set
type AsPathSet struct {
	// original -> bgp-pol:as-path-set-name
	AsPathSetName string
	// original -> gobgp:as-path
	AsPath []string
}

//struct for container bgp-pol:as-path-sets
type AsPathSets struct {
	// original -> bgp-pol:as-path-set
	AsPathSetList []AsPathSet
}

//struct for container bgp-pol:ext-community-set
type ExtCommunitySet struct {
	// original -> bgp-pol:ext-community-set-name
	ExtCommunitySetName string
	// original -> gobgp:ext-community
	ExtCommunity []string
}

//struct for container bgp-pol:ext-community-sets
type ExtCommunitySets struct {
	// original -> bgp-pol:ext-community-set
	ExtCommunitySetList []ExtCommunitySet
}

//struct for container bgp-pol:community-set
type CommunitySet struct {
	// original -> bgp-pol:community-set-name
	CommunitySetName string
	// original -> gobgp:community
	Community []string
}

//struct for container bgp-pol:community-sets
type CommunitySets struct {
	// original -> bgp-pol:community-set
	CommunitySetList []CommunitySet
}

//struct for container bgp-pol:bgp-defined-sets
type BgpDefinedSets struct {
	// original -> bgp-pol:community-sets
	CommunitySets CommunitySets
	// original -> bgp-pol:ext-community-sets
	ExtCommunitySets ExtCommunitySets
	// original -> bgp-pol:as-path-sets
	AsPathSets AsPathSets
}

//struct for container rpol:tag
type Tag struct {
	// original -> rpol:value
	Value TagType
}

//struct for container rpol:tag-set
type TagSet struct {
	// original -> rpol:tag-set-name
	TagSetName string
	// original -> rpol:tag
	TagList []Tag
}

//struct for container rpol:tag-sets
type TagSets struct {
	// original -> rpol:tag-set
	TagSetList []TagSet
}

//struct for container rpol:neighbor-set
type NeighborSet struct {
	// original -> rpol:neighbor-set-name
	NeighborSetName string
	// original -> gobgp:neighbor-info
	//original type is list of inet:ip-address
	NeighborInfo []string
}

//struct for container rpol:neighbor-sets
type NeighborSets struct {
	// original -> rpol:neighbor-set
	NeighborSetList []NeighborSet
}

//struct for container rpol:prefix
type Prefix struct {
	// original -> rpol:ip-prefix
	//rpol:ip-prefix's original type is inet:ip-prefix
	IpPrefix string
	// original -> rpol:masklength-range
	MasklengthRange string
}

//struct for container rpol:prefix-set
type PrefixSet struct {
	// original -> rpol:prefix-set-name
	PrefixSetName string
	// original -> rpol:prefix
	PrefixList []Prefix
}

//struct for container rpol:prefix-sets
type PrefixSets struct {
	// original -> rpol:prefix-set
	PrefixSetList []PrefixSet
}

//struct for container rpol:defined-sets
type DefinedSets struct {
	// original -> rpol:prefix-sets
	PrefixSets PrefixSets
	// original -> rpol:neighbor-sets
	NeighborSets NeighborSets
	// original -> rpol:tag-sets
	TagSets TagSets
	// original -> bgp-pol:bgp-defined-sets
	BgpDefinedSets BgpDefinedSets
}

//struct for container rpol:routing-policy
type RoutingPolicy struct {
	// original -> rpol:defined-sets
	DefinedSets DefinedSets
	// original -> rpol:policy-definitions
	PolicyDefinitions PolicyDefinitions
}
