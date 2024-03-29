package nsxt

import (
	"fmt"
	"testing"
)

func TestApiServiceNSXT_GetEdgeCluster(t *testing.T) {
	var (
		err error
		r   EdgeCluster
	)

	if r, err = GetEdgeCluster("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}

func TestApiServiceNSXT_GetTier0Gateways(t *testing.T) {
	var (
		err error
		r   Tier0Gateway
	)

	if r, err = GetTier0Gateways("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}

func TestApiServiceNSXT_GetTransportZones(t *testing.T) {
	var (
		err error
		r   TransportZones
	)

	if r, err = GetTransportZones("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}

func TestApiServiceNSXT_GetLogicalSwitchs(t *testing.T) {
	var (
		err error
		r   LogicalSwitches
	)

	if r, err = GetLogicalSwitchs("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}

func TestApiServiceNSXT_GetSegments(t *testing.T) {
	var (
		err error
		r   Segments
	)

	if r, err = GetSegments("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}

func TestApiServiceNSXT_GetGroups(t *testing.T) {
	var (
		err error
		r   Groups
	)

	if r, err = GetGroups("TESP5"); err != nil {
		t.Fatalf("error: %s", err)
	}

	fmt.Println(r)
}
