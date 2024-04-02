package service

import (
	"fmt"
	"strings"

	"go-vrf/src/model"
	"go-vrf/src/nsxt"
	"go-vrf/src/objects"
)

func CreateNetworksVRF(obj objects.NetworksProdutcsVRF) (net []model.Networks, err error) {
	var (
		sg   nsxt.Segments
		gp   nsxt.Groups
		pf   nsxt.Profiles
		sw   nsxt.LogicalSwitches
		edge = obj.Edge
	)

	if obj.Edge == "" {
		return net, fmt.Errorf("edge is required")
	}

	if obj.NameTier1 == "" {
		return net, fmt.Errorf("db-shared name on tier 1 is required")
	}

	if sg, err = nsxt.GetSegments(edge); err != nil {
		return
	}

	if gp, err = nsxt.GetGroups(edge); err != nil {
		return
	}

	if sw, err = nsxt.GetLogicalSwitchs(edge); err != nil {
		return
	}

	for idx, key := range obj.Products {

		net = append(
			net, model.Networks{
				Name:        key,
				Description: key,
			},
		)

		for _, v := range sg.Results {
			if strings.Contains(v.DisplayName, key) {
				net[idx].SegmentID = v.Id
				net[idx].DisplayName = v.DisplayName
				net[idx].Address = v.Subnets[0].Network
				break
			}
		}

		for _, v := range gp.Results {
			if strings.Contains(v.DisplayName, key) {
				net[idx].GroupID = v.Id
				break
			}
		}

		segmentId := fmt.Sprintf("DB-Shared_%s", key)

		if pf, err = nsxt.GetProfiles(segmentId, edge); err != nil {
			return
		}

		net[idx].ProfileID = pf.Results[0].Id

		for _, v := range sw.Results {
			if strings.Contains(v.DisplayName, key) {
				net[idx].SwitchID = v.Id
				break
			}
		}
	}

	return
}
