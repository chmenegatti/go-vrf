package service

import (
	"context"

	"go-vrf/src/model"
	"go-vrf/src/nsxt"
	"go-vrf/src/objects"
)

func GenerateEtcdKey(ctx context.Context, etcd objects.EdgeClusterEtcd) (data model.EdgeCluster, err error) {

	var (
		ec nsxt.EdgeCluster
		tz nsxt.TransportZones
	)

	if etcd.NsxtEdgeClusterName == "" {
		return data, ErrEdgeClusterNameRequired
	}

	if ec, err = nsxt.GetEdgeCluster(ctx, etcd.Edge); err != nil {
		return
	}

	for _, v := range ec.Results {
		if v.DisplayName == etcd.NsxtEdgeClusterName {
			data.NsxtEdgeClusterID = v.Id
			break
		}
	}

	if tz, err = nsxt.GetTransportZones(ctx, etcd.Edge); err != nil {
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
