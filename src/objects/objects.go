package objects

type EdgeClusterEtcd struct {
	Edge                    string `json:",omitempty"`
	VrfName                 string `json:",omitempty"`
	NsxtEdgeClusterName     string `json:",omitempty"`
	TransportZoneName       string `json:",omitempty"`
	VirtualFirewall         string `json:",omitempty"`
	FirewallExternalAddress string `json:",omitempty"`
	MaxOrganization         int    `json:",omitempty"`
	Enable                  bool   `json:",omitempty"`
	RubrikDatabaseCluster   string `json:",omitempty"`
	NextHopVpn              string `json:",omitempty"`
}

type OrganizationVRF struct {
	Edge      string `json:",omitempty"`
	NameTier1 string `json:",omitempty"`
	VrfName   string `json:",omitempty"`
}

type NetworksProdutcsVRF struct {
	Edge      string   `json:",omitempty"`
	NameTier1 string   `json:",omitempty"`
	Products  []string `json:",omitempty"`
}
