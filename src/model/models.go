package model

type EdgeCluster struct {
	Index                   int    `json:"index,omitempty"`
	NsxtTier0ID             string `json:"nsxt_tier0_id,omitempty"`
	NsxtTier0DisplayName    string `json:"nsxt_tier0_display_name,omitempty"`
	NsxtEdgeClusterID       string `json:"nsxt_edge_cluster_id,omitempty"`
	GatewayEdgeClusterID    string `json:"gateway_edge_cluster_id,omitempty"`
	NsxtTransportZoneID     string `json:"nsxt_transport_zone_id,omitempty"`
	DatabaseTierID          string `json:"database_tier_id,omitempty"`
	PhysicalFirewall        string `json:"physical_firewall,omitempty"`
	VirtualFirewall         string `json:"virtual_firewall,omitempty"`
	VpnSite                 string `json:"vpn_site,omitempty"`
	FirewallExternalAddress string `json:"firewall_external_address,omitempty"`
	MaxOrganization         int    `json:"max_organization,omitempty"`
	Enable                  bool   `json:"enable,omitempty"`
	RubrikDatabaseCluster   string `json:"rubrik_database_cluster,omitempty"`
	NextHopVpn              string `json:"next_hop_vpn,omitempty"`
}

type Organizations struct {
	Id                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	TierProvider            string `json:"tier_provider,omitempty"`
	EdgeCluster             string `json:"edge_cluster,omitempty"`
	TierProviderID          string `json:"tier_provider_id"`
	Tier1GatewayID          string `json:"tier1_gateway_id"`
	PolicyID                string `json:"policy_id"`
	LocaleServiceID         string `json:"locale_service_id"`
	BackupCluster           string `json:"backup_cluster,omitempty"`
	LoadBalanceID           string `json:"load_balance_id"`
	PhysicalFirewall        string `json:"physical_firewall"`
	VirtualFirewall         string `json:"virtual_firewall"`
	FirewallExternalAddress string `json:"firewall_external_address"`
	LoadBalanceSize         string `json:"load_balance_size"`
	Status                  string `json:"status"`
}

type Networks struct {
	Id                      string `json:"id"`
	Organization            string `json:"organization"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Address                 string `json:"address"`
	SegmentID               string `json:"segment_id"`
	SwitchID                string `json:"switch_id"`
	GroupID                 string `json:"group_id"`
	ProfileID               string `json:"profile_id"`
	DisplayName             string `json:"display_name"`
	EnableSideCommunication bool   `json:"enable_side_communication"`
	Status                  string `json:"status"`
}
