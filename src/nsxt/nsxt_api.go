package nsxt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-vrf/src/configs"
)

// fetch performs an authenticated GET against the NSX-T API for the given
// edge and decodes the JSON response into T. It centralises URL building,
// response-body handling, and error wrapping for every endpoint below.
func fetch[T any](edge, path string) (T, error) {
	var out T

	base := configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge))
	url := base + path

	res, err := RequestNSXTApi(url, edge)
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

type EdgeCluster struct {
	Results []struct {
		DeploymentType string `json:"deployment_type"`
		MemberNodeType string `json:"member_node_type"`
		ResourceType   string `json:"resource_type"`
		Id             string `json:"id"`
		DisplayName    string `json:"display_name"`
		Description    string `json:"description"`
		CreateUser     string `json:"_create_user"`
	} `json:"results"`
}

func GetEdgeCluster(edge string) (EdgeCluster, error) {
	return fetch[EdgeCluster](edge, "/api/v1/edge-clusters")
}

type Tier0Gateway struct {
	Results []struct {
		ResourceType string `json:"resource_type"`
		Id           string `json:"id"`
		DisplayName  string `json:"display_name"`
		Path         string `json:"path"`
		ParentPath   string `json:"parent_path"`
		RelativePath string `json:"relative_path"`
	} `json:"results"`
}

func GetTier0Gateways(edge string) (Tier0Gateway, error) {
	return fetch[Tier0Gateway](edge, "/policy/api/v1/infra/tier-0s")
}

type Tier1Gateway struct {
	Results []struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
		Path        string `json:"path"`
	}
}

func GetTier1Gateways(edge string) (Tier1Gateway, error) {
	return fetch[Tier1Gateway](edge, "/policy/api/v1/infra/tier-1s")
}

type DistributedFirewalPolicy struct {
	Results []struct {
		Id           string `json:"id"`
		DisplayName  string `json:"display_name"`
		Path         string `json:"path"`
		RelativePath string `json:"relative_path"`
	}
}

func GetDistributedFirewallPolicy(edge string) (DistributedFirewalPolicy, error) {
	return fetch[DistributedFirewalPolicy](edge, "/policy/api/v1/infra/domains/default/security-policies")
}

type TransportZones struct {
	Results []struct {
		Id             string `json:"id"`
		ResourceType   string `json:"resource_type"`
		TransportType  string `json:"transport_type"`
		HostSwitchName string `json:"host_switch_name"`
		DisplayName    string `json:"display_name"`
	} `json:"results"`
}

func GetTransportZones(edge string) (TransportZones, error) {
	return fetch[TransportZones](edge, "/api/v1/infra/sites/default/enforcement-points/default/transport-zones")
}

type LogicalSwitches struct {
	Results []struct {
		Id          string `json:"id,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
	}
}

func GetLogicalSwitchs(edge string) (LogicalSwitches, error) {
	return fetch[LogicalSwitches](edge, "/api/v1/logical-switches?sort_by=display_name")
}

type Segments struct {
	Results []struct {
		Id           string `json:"id,omitempty"`
		DisplayName  string `json:"display_name"`
		Path         string `json:"path"`
		ResourceType string `json:"resource_type"`
		Subnets      []struct {
			GatwayAddress string `json:"gateway_address"`
			Network       string `json:"network"`
		} `json:"subnets"`
	} `json:"results"`
}

func GetSegments(edge string) (Segments, error) {
	return fetch[Segments](edge, "/policy/api/v1/infra/segments")
}

type Groups struct {
	Results []struct {
		Id          string `json:"id,omitempty"`
		DisplayName string `json:"display_name"`
		Path        string `json:"path"`
	}
}

func GetGroups(edge string) (Groups, error) {
	return fetch[Groups](edge, "/policy/api/v1/infra/domains/default/groups?sort_by=display_name")
}

type Profiles struct {
	Results []struct {
		ResourceType string `json:"resource_type,omitempty"`
		Id           string `json:"id,omitempty"`
		DisplayName  string `json:"display_name,omitempty"`
		Path         string `json:"path,omitempty"`
		RelativePath string `json:"relative_path,omitempty"`
		ParentPath   string `json:"parent_path,omitempty"`
	} `json:"results"`
}

func GetProfiles(segmentId, edge string) (Profiles, error) {
	path := fmt.Sprintf("/policy/api/v1/infra/segments/%s/segment-discovery-profile-binding-maps", segmentId)
	return fetch[Profiles](edge, path)
}
