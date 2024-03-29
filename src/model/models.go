package model

import "time"

type EdgeCluster struct {
	Index                   int    `json:",omitempty"`
	NsxtTier0ID             string `json:",omitempty"`
	NsxtTier0DisplayName    string `json:",omitempty"`
	NsxtEdgeClusterID       string `json:",omitempty"`
	GatewayEdgeClusterID    string `json:",omitempty"`
	NsxtTransportZoneID     string `json:",omitempty"`
	DatabaseTierID          string `json:",omitempty"`
	PhysicalFirewall        string `json:",omitempty"`
	VirtualFirewall         string `json:",omitempty"`
	VpnSite                 string `json:",omitempty"`
	FirewallExternalAddress string `json:",omitempty"`
	MaxOrganization         int    `json:",omitempty"`
	Enable                  bool   `json:",omitempty"`
	RubrikDatabaseCluster   string `json:",omitempty"`
	NextHopVpn              string `json:",omitempty"`
}

type Organizations struct {
	Id                      string    `json:"id,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Description             string    `json:"description,omitempty"`
	TierProvider            string    `json:"tier_provider,omitempty"`
	EdgeCluster             string    `json:"edge_cluster,omitempty"`
	TierProviderID          string    `json:"tier_provider_id"`
	Tier1GatewayID          string    `json:"tier1_gateway_id"`
	PolicyID                string    `json:"policy_id"`
	LocaleServiceID         string    `json:"locale_service_id"`
	BackupCluster           string    `json:"backup_cluster,omitempty"`
	LoadBalanceID           string    `json:"load_balance_id"`
	PhysicalFirewall        string    `json:"physical_firewall"`
	VirtualFirewall         string    `json:"virtual_firewall"`
	FirewallExternalAddress string    `json:"firewall_external_address"`
	LoadBalanceSize         string    `json:"load_balance_size"`
	Status                  string    `json:"status"`
	Created                 time.Time `json:"created"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               time.Time `json:"deleted_at"`
}

type Networks struct {
	Id                      string    `json:"id"`
	Organization            string    `json:"organization"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	Address                 string    `json:"address"`
	SegmentID               string    `json:"segment_id"`
	SwitchID                string    `json:"switch_id"`
	GroupID                 string    `json:"group_id"`
	ProfileID               string    `json:"profile_id"`
	DisplayName             string    `json:"display_name"`
	EnableSideCommunication bool      `json:"enable_side_communication"`
	Status                  string    `json:"status"`
	Created                 time.Time `json:"created"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               time.Time `json:"deleted_at"`
}
