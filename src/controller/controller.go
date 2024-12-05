package controller

import (
	"fmt"
	"strings"

	"go-vrf/src/model"
	"go-vrf/src/objects"
	"go-vrf/src/service"
	utilities "go-vrf/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GenerateEtcdKey(c *fiber.Ctx) error {

	var (
		payload    objects.EdgeClusterEtcd
		edgeResult model.EdgeCluster
		err        error
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Error parsing request",
			},
		)
	}

	if edgeResult, err = service.GenerateEtcdKey(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error generating etcd key",
			},
		)
	}

	uuids := utilities.GenerateUUIDs(2)

	var result = model.EdgeCluster{
		Index:                   11,
		NsxtTier0ID:             payload.VrfName,
		NsxtTier0DisplayName:    payload.VrfName,
		NsxtEdgeClusterID:       edgeResult.NsxtEdgeClusterID,
		GatewayEdgeClusterID:    uuids[0],
		NsxtTransportZoneID:     edgeResult.NsxtTransportZoneID,
		DatabaseTierID:          uuids[1],
		PhysicalFirewall:        "fisico",
		VirtualFirewall:         payload.VirtualFirewall,
		VpnSite:                 "physical",
		FirewallExternalAddress: payload.FirewallExternalAddress,
		MaxOrganization:         payload.MaxOrganization,
		Enable:                  payload.Enable,
		RubrikDatabaseCluster:   payload.RubrikDatabaseCluster,
	}

	if err := utilities.SaveToFile(payload.VrfName, result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": fmt.Errorf("error on save to file: %v", err),
			},
		)
	}

	vrf := fmt.Sprintf("%s", payload.VrfName)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			vrf:        result,
			"filename": fmt.Sprintf("%s.json", payload.VrfName),
		},
	)
}

func CreateOrganizationVRF(c *fiber.Ctx) error {

	var (
		payload objects.OrganizationVRF
		org     model.Organizations
		err     error
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Error parsing request",
			},
		)
	}

	if org, err = service.CreateOrganizationVRF(payload.NameTier1, payload.Edge); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error creating organization VRF",
			},
		)
	}

	uuids := utilities.GenerateUUIDs(2)

	jsonData, err := utilities.ReadT0Json(fmt.Sprintf("%s.json", payload.VrfName))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": fmt.Sprintf("error reading JSON data: %v", err),
			},
		)

	}

	var result = model.Organizations{
		Id:                      uuids[0],
		Name:                    payload.NameTier1,
		Description:             payload.NameTier1,
		TierProvider:            jsonData[payload.VrfName].NsxtTier0ID,
		TierProviderID:          jsonData[payload.VrfName].NsxtTier0ID,
		EdgeCluster:             jsonData[payload.VrfName].GatewayEdgeClusterID,
		Tier1GatewayID:          org.Tier1GatewayID,
		PolicyID:                org.PolicyID,
		LocaleServiceID:         "default",
		LoadBalanceID:           uuids[1],
		BackupCluster:           jsonData[payload.VrfName].RubrikDatabaseCluster,
		PhysicalFirewall:        jsonData[payload.VrfName].PhysicalFirewall,
		VirtualFirewall:         jsonData[payload.VrfName].VirtualFirewall,
		FirewallExternalAddress: jsonData[payload.VrfName].FirewallExternalAddress,
		LoadBalanceSize:         strings.ToUpper("small"),
		Status:                  "COMPLETED",
	}

	if err := utilities.SaveToFile(payload.NameTier1, result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": fmt.Errorf("error on save to file: %v", err),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data": result,
			"sql":  organizationsInsert(result),
		},
	)
}

func CreateNetworksProducts(c *fiber.Ctx) error {
	var (
		payload objects.NetworksProdutcsVRF
		net     []model.Networks
		err     error
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Error parsing JSON request",
			},
		)
	}

	if net, err = service.CreateNetworksVRF(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error creating networks VRF",
			},
		)
	}

	uuids := utilities.GenerateUUIDs(6)

	jsonData, err := utilities.ReadT1Json(fmt.Sprintf("%s.json", payload.NameTier1))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": fmt.Sprintf("error reading JSON data: %v", err),
			},
		)
	}

	for key := range net {
		net[key].Id = uuids[key]
		net[key].Organization = jsonData[payload.NameTier1].Id
		net[key].Status = "COMPLETED"
	}

	jsonFile := fmt.Sprintf("%s_Networks", payload.NameTier1)

	if err := utilities.SaveToFile(jsonFile, net); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": fmt.Errorf("error on save to file: %v", err),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data": net,
			"sql":  networksInsert(net),
		},
	)

}

func organizationsInsert(o model.Organizations) string {
	return fmt.Sprintf(
		"INSERT INTO organizations (id, name, description, tier_provider, tier_provider_id, edge_cluster, tier1_gateway_id, policy_id, locale_service_id, load_balance_id, backup_cluster, physical_firewall, virtual_firewall, firewall_external_address, load_balance_size, status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		o.Id, o.Name, o.Description, o.TierProvider, o.TierProviderID, o.EdgeCluster, o.Tier1GatewayID, o.PolicyID,
		o.LocaleServiceID, o.LoadBalanceID, o.BackupCluster, o.PhysicalFirewall, o.VirtualFirewall,
		o.FirewallExternalAddress, o.LoadBalanceSize, o.Status,
	)
}

func networksInsert(n []model.Networks) []string {
	var (
		sqlString []string
	)

	for _, v := range n {
		sqlString = append(
			sqlString, fmt.Sprintf(
				"INSERT INTO networks (id, organization, name, description, address, segment_id, switch_id, group_id, profile_id, display_name, enable_side_communication, status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %t, '%s');",
				v.Id, v.Organization, v.Name, v.Description, v.Address, v.SegmentID, v.SwitchID, v.GroupID, v.ProfileID,
				v.DisplayName, v.EnableSideCommunication, v.Status,
			),
		)
	}

	return sqlString
}
