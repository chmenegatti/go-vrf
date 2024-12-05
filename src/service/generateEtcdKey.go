package service

import (
	"fmt"

	"go-vrf/src/model"
	"go-vrf/src/nsxt"
	"go-vrf/src/objects"
)

func GenerateEtcdKey(etcd objects.EdgeClusterEtcd) (data model.EdgeCluster, err error) {

	var (
		ec nsxt.EdgeCluster
		tz nsxt.TransportZones
	)

	if etcd.NsxtEdgeClusterName == "" {
		return data, fmt.Errorf("nsxt cluster name is empty")
	}

	if ec, err = nsxt.GetEdgeCluster(etcd.Edge); err != nil {
		return
	}

	for _, v := range ec.Results {
		if v.DisplayName == etcd.NsxtEdgeClusterName {
			data.NsxtEdgeClusterID = v.Id
			break
		}
	}

	if tz, err = nsxt.GetTransportZones(etcd.Edge); err != nil {
		return
	}

	for _, v := range tz.Results {
		if v.DisplayName == etcd.TransportZoneName {
			data.NsxtTransportZoneID = v.Id
			break
		}
	}
	return
}
