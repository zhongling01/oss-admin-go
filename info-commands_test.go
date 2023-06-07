package madmin

import (
	"context"
	"fmt"
	"testing"
)

const (
	EndpointDefault        = "127.0.0.1:" + PortDefault
	PortDefault            = "19000"
	AccessKeyIDDefault     = "minioadmin"
	SecretAccessKeyDefault = "minioadmin"
	RegionDefault          = "us-east-1"
)

func TestAdminClient_GetClusterInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.GetClusterInfo(context.Background(), false)
	if err != nil {
		t.Fatal(err.Error())
	}
	for key, value := range info.Config {
		t.Log(fmt.Sprintf("%s: %s", key, value))
	}
	for _, pool := range info.Pools {
		t.Log(fmt.Sprintf("%+v", pool))
	}

	info, err = c.GetClusterInfo(context.Background(), true)
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, pool := range info.Pools {
		for _, set := range pool.Sets {
			t.Log(fmt.Sprintf("=== pool %d set %d ===", pool.Index, set.Index))
			for _, member := range set.Member {
				t.Log(fmt.Sprintf("%s", member.Endpoint))
			}
		}
	}
}

func TestAdminClient_GetBucketInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.GetBucketInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	for bucket, info := range info {
		t.Log(fmt.Sprintf("%s: %+v", bucket, info))
	}
}

func TestAdminClient_ServerInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.ServerInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(fmt.Sprintf("%+v", info))
}

func TestAdminClient_DataUsageInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.DataUsageInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(fmt.Sprintf("%+v", info))
}

func TestAdminClient_StorageInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.StorageInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(fmt.Sprintf("%+v", info))
}

func TestAdminClient_AccountInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.AccountInfo(context.Background(), AccountOpts{})
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(fmt.Sprintf("%+v", info))
}
