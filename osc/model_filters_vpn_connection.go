/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// FiltersVpnConnection One or more filters.
type FiltersVpnConnection struct {
	// The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
	BgpAsns []int32 `json:"BgpAsns,omitempty"`
	// The IDs of the client gateways.
	ClientGatewayIds []string `json:"ClientGatewayIds,omitempty"`
	// The types of the VPN connections (only `ipsec.1` is supported).
	ConnectionTypes []string `json:"ConnectionTypes,omitempty"`
	// The destination IP ranges.
	RouteDestinationIpRanges []string `json:"RouteDestinationIpRanges,omitempty"`
	// The states of the VPN connections (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States []string `json:"States,omitempty"`
	// If `false`, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If `true`, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute).
	StaticRoutesOnly bool `json:"StaticRoutesOnly,omitempty"`
	// The IDs of the virtual gateways.
	VirtualGatewayIds []string `json:"VirtualGatewayIds,omitempty"`
	// The IDs of the VPN connections.
	VpnConnectionIds []string `json:"VpnConnectionIds,omitempty"`
}