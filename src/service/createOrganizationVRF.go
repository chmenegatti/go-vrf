package service

import (
	"context"

	"go-vrf/src/model"
	"go-vrf/src/nsxt"
)

func CreateOrganizationVRF(ctx context.Context, name, edge string) (org model.Organizations, err error) {
	var (
		tp  nsxt.Tier1Gateway
		dfp nsxt.DistributedFirewallPolicy
	)

	if name == "" {
		return org, ErrNameRequired
	}

	if edge == "" {
		return org, ErrEdgeRequired
	}

	if tp, err = nsxt.GetTier1Gateways(ctx, edge); err != nil {
		return
	}

	for _, v := range tp.Results {
		if v.DisplayName == name {
			org.Tier1GatewayID = v.Id
			break
		}
	}

	if dfp, err = nsxt.GetDistributedFirewallPolicy(ctx, edge); err != nil {
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
