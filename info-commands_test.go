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

	info, err := c.GetClusterInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, pool := range info.Pools {
		t.Log(fmt.Sprintf("%+v", pool))
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
