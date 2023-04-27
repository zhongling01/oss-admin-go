package madmin

import (
	"context"
	"testing"
)

func TestAdminClient_GetVacancyInfo(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.GetVacancyInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(info)
}

func TestAdminClient_SetVacancy(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = c.SetVacancy(context.Background(), VacancyInfo{
		Enabled:          false,
		CheckInterval:    30,
		VacancyThreshold: 50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAdminClient_ManualMergeVacancy(t *testing.T) {
	c, err := New(EndpointDefault, AccessKeyIDDefault, SecretAccessKeyDefault, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = c.ManualMergeVacancy(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}
