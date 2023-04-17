package madmin

import (
	"context"
	"fmt"
	"testing"
)

func TestAdminClient_GetClusterInfo(t *testing.T) {
	c, err := New("127.0.0.1:19000", "minioadmin", "minioadmin", false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.GetClusterInfo(context.Background(), false)
	if err != nil {
		t.Fatal(err.Error())
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
	c, err := New("127.0.0.1:19000", "minioadmin", "minioadmin", false)
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
