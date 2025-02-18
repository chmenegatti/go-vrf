package nsxt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-vrf/src/configs"
)

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

func GetEdgeCluster(edge string) (edgeCluster EdgeCluster, err error) {

	var (
		res *http.Response
	)

	path := "/api/v1/edge-clusters"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return edgeCluster, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &edgeCluster)
	if err != nil {
		return
	}

	return edgeCluster, nil
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

func GetTier0Gateways(edge string) (tier0 Tier0Gateway, err error) {

	var (
		res *http.Response
	)

	path := "/policy/api/v1/infra/tier-0s"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return tier0, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &tier0)
	if err != nil {
		return
	}

	return tier0, nil
}

type Tier1Gateway struct {
	Results []struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
		Path        string `json:"path"`
	}
}

func GetTier1Gateways(edge string) (tier1 Tier1Gateway, err error) {

	var (
		res *http.Response
	)

	path := "/policy/api/v1/infra/tier-1s"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return tier1, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &tier1)
	if err != nil {
		return
	}

	return tier1, nil
}

type DistributedFirewalPolicy struct {
	Results []struct {
		Id           string `json:"id"`
		DisplayName  string `json:"display_name"`
		Path         string `json:"path"`
		RelativePath string `json:"relative_path"`
	}
}

func GetDistributedFirewallPolicy(edge string) (dfp DistributedFirewalPolicy, err error) {

	var (
		res *http.Response
	)

	path := "/policy/api/v1/infra/domains/default/security-policies"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return dfp, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &dfp)
	if err != nil {
		return
	}

	return dfp, nil

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

func GetTransportZones(edge string) (transportZones TransportZones, err error) {

	var (
		res *http.Response
	)

	path := "/api/v1/transport-zones"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return transportZones, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &transportZones)
	if err != nil {
		return
	}

	return transportZones, nil

}

type LogicalSwitches struct {
	Results []struct {
		Id          string `json:"id,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
	}
}

func GetLogicalSwitchs(edge string) (switchIds LogicalSwitches, err error) {

	var (
		res *http.Response
	)

	path := "/api/v1/logical-switches?sort_by=display_name"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return switchIds, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &switchIds)
	if err != nil {
		return
	}

	return switchIds, nil
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

func GetSegments(edge string) (segments Segments, err error) {

	var (
		res *http.Response
	)

	path := "/policy/api/v1/infra/segments"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return segments, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &segments)
	if err != nil {
		return
	}

	return segments, nil
}

type Groups struct {
	Results []struct {
		Id          string `json:"id,omitempty"`
		DisplayName string `json:"display_name"`
		Path        string `json:"path"`
	}
}

func GetGroups(edge string) (groups Groups, err error) {

	var (
		res *http.Response
	)

	path := "/policy/api/v1/infra/domains/default/groups?sort_by=display_name"
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return groups, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &groups)
	if err != nil {
		return
	}

	return groups, nil
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

func GetProfiles(segmentId, edge string) (profiles Profiles, err error) {
	var (
		res *http.Response
	)

	path := fmt.Sprintf("/policy/api/v1/infra/segments/%s/segment-discovery-profile-binding-maps", segmentId)
	url := fmt.Sprintf("%s%s", configs.GetEnvKeys(fmt.Sprintf("%s_BASEPATH", edge)), path)

	if res, err = RequestNSXTApi(url, edge); err != nil {
		return profiles, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = json.Unmarshal(bodyBytes, &profiles)
	if err != nil {
		return
	}

	return profiles, nil
}
