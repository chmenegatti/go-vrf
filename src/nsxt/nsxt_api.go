package nsxt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go-vrf/src/configs"
)

// fetch performs an authenticated GET against the NSX-T API for the given
// edge and decodes the JSON response into T. It centralises URL building,
// response-body handling, and error wrapping for every endpoint below.
func fetch[T any](ctx context.Context, edge, path string) (T, error) {
	var out T

	base := configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge))
	url := base + path

	res, err := RequestNSXTApi(ctx, url, edge)
	if err != nil {
		return out, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return out, fmt.Errorf("nsxt GET %s: unexpected status %d", path, res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return out, fmt.Errorf("nsxt GET %s: decode response: %w", path, err)
	}

	return out, nil
}

type EdgeClusterMember struct {
	DeploymentType string `json:"deployment_type"`
	MemberNodeType string `json:"member_node_type"`
	ResourceType   string `json:"resource_type"`
	Id             string `json:"id"`
	DisplayName    string `json:"display_name"`
	Description    string `json:"description"`
	CreateUser     string `json:"_create_user"`
}

type EdgeCluster struct {
	Results []EdgeClusterMember `json:"results"`
}

func GetEdgeCluster(ctx context.Context, edge string) (EdgeCluster, error) {
	return fetch[EdgeCluster](ctx, edge, "/api/v1/edge-clusters")
}

type Tier0 struct {
	ResourceType string `json:"resource_type"`
	Id           string `json:"id"`
	DisplayName  string `json:"display_name"`
	Path         string `json:"path"`
	ParentPath   string `json:"parent_path"`
	RelativePath string `json:"relative_path"`
}

type Tier0Gateway struct {
	Results []Tier0 `json:"results"`
}

func GetTier0Gateways(ctx context.Context, edge string) (Tier0Gateway, error) {
	return fetch[Tier0Gateway](ctx, edge, "/policy/api/v1/infra/tier-0s")
}

type Tier1 struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Path        string `json:"path"`
}

type Tier1Gateway struct {
	Results []Tier1 `json:"results"`
}

func GetTier1Gateways(ctx context.Context, edge string) (Tier1Gateway, error) {
	return fetch[Tier1Gateway](ctx, edge, "/policy/api/v1/infra/tier-1s")
}

type SecurityPolicy struct {
	Id           string `json:"id"`
	DisplayName  string `json:"display_name"`
	Path         string `json:"path"`
	RelativePath string `json:"relative_path"`
}

type DistributedFirewallPolicy struct {
	Results []SecurityPolicy `json:"results"`
}

func GetDistributedFirewallPolicy(ctx context.Context, edge string) (DistributedFirewallPolicy, error) {
	return fetch[DistributedFirewallPolicy](ctx, edge, "/policy/api/v1/infra/domains/default/security-policies")
}

type TransportZone struct {
	Id             string `json:"id"`
	ResourceType   string `json:"resource_type"`
	TransportType  string `json:"transport_type"`
	HostSwitchName string `json:"host_switch_name"`
	DisplayName    string `json:"display_name"`
}

type TransportZones struct {
	Results []TransportZone `json:"results"`
}

func GetTransportZones(ctx context.Context, edge string) (TransportZones, error) {
	return fetch[TransportZones](ctx, edge, "/api/v1/infra/sites/default/enforcement-points/default/transport-zones")
}

type LogicalSwitch struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type LogicalSwitches struct {
	Results []LogicalSwitch `json:"results"`
}

func GetLogicalSwitches(ctx context.Context, edge string) (LogicalSwitches, error) {
	return fetch[LogicalSwitches](ctx, edge, "/api/v1/logical-switches?sort_by=display_name")
}

type Subnet struct {
	GatewayAddress string `json:"gateway_address"`
	Network        string `json:"network"`
}

type Segment struct {
	Id           string   `json:"id,omitempty"`
	DisplayName  string   `json:"display_name"`
	Path         string   `json:"path"`
	ResourceType string   `json:"resource_type"`
	Subnets      []Subnet `json:"subnets"`
}

type Segments struct {
	Results []Segment `json:"results"`
}

func GetSegments(ctx context.Context, edge string) (Segments, error) {
	return fetch[Segments](ctx, edge, "/policy/api/v1/infra/segments")
}

type Group struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name"`
	Path        string `json:"path"`
}

type Groups struct {
	Results []Group `json:"results"`
}

func GetGroups(ctx context.Context, edge string) (Groups, error) {
	return fetch[Groups](ctx, edge, "/policy/api/v1/infra/domains/default/groups?sort_by=display_name")
}

type Profile struct {
	ResourceType string `json:"resource_type,omitempty"`
	Id           string `json:"id,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	Path         string `json:"path,omitempty"`
	RelativePath string `json:"relative_path,omitempty"`
	ParentPath   string `json:"parent_path,omitempty"`
}

type Profiles struct {
	Results []Profile `json:"results"`
}

func GetProfiles(ctx context.Context, segmentId, edge string) (Profiles, error) {
	path := fmt.Sprintf("/policy/api/v1/infra/segments/%s/segment-discovery-profile-binding-maps", segmentId)
	return fetch[Profiles](ctx, edge, path)
}
