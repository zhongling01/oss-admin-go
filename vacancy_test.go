package madmin

import (
	"context"
	"math/rand"
	"testing"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	checkInterval := rand.Intn(1000) + 1
	vacancyThreshold := rand.Intn(100) + 1
	err = c.SetVacancy(context.Background(), VacancyInfo{
		Enabled:          true,
		CheckInterval:    checkInterval,
		VacancyThreshold: vacancyThreshold,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	info, err := c.GetVacancyInfo(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	if info.Enabled != true || info.CheckInterval != checkInterval || info.VacancyThreshold != vacancyThreshold {
		t.Fatal("set failed")
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
