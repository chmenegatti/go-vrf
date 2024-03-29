package service

import (
	"fmt"

	"go-vrf/src/model"
	"go-vrf/src/nsxt"
)

func CreateOrganizationVRF(name, edge string) (org model.Organizations, err error) {
	var (
		tp  nsxt.Tier1Gateway
		dfp nsxt.DistributedFirewalPolicy
	)

	if name == "" {
		return org, fmt.Errorf("name is required")
	}

	if edge == "" {
		return org, fmt.Errorf("edge is required")
	}

	if tp, err = nsxt.GetTier1Gateways(edge); err != nil {
		return
	}

	for _, v := range tp.Results {
		if v.DisplayName == name {
			org.Tier1GatewayID = v.Id
			break
		}
	}

	if dfp, err = nsxt.GetDistributedFirewallPolicy(edge); err != nil {
		return
	}

	for _, v := range dfp.Results {
		if v.DisplayName == name {
			org.PolicyID = v.Id
			break
		}
	}

	return
}
