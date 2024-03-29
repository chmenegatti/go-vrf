package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go-vrf/src/model"
	"go-vrf/src/objects"
	"go-vrf/src/service"
	utilities "go-vrf/src/utils"
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
		Created:                 time.Now(),
		UpdatedAt:               time.Now(),
		DeletedAt:               time.Time{},
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
		},
	)
}
