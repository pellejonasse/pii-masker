package piimasker_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

// added tests for external packages to see if the masker can handle them.
type cloudWatchWrapper struct {
	Event  events.CloudWatchEvent `Pii:"mask"`
	UserID string                 `Pii:"mask"`
}

type cloudWatchWrapperShow struct {
	Event  events.CloudWatchEvent `Pii:"show"`
	UserID string                 `Pii:"show"`
}

type cloudWatchWrapperAnonymize struct {
	Event  events.CloudWatchEvent `Pii:"anonymize"`
	UserID string                 `Pii:"anonymize"`
}

type cloudWatchWrapperNoTag struct {
	Event  events.CloudWatchEvent
	UserID string
}

func newCloudWatchWrapperFixture() cloudWatchWrapper {
	return cloudWatchWrapper{
		UserID: "user-123",
		Event: events.CloudWatchEvent{
			Version:    "0",
			ID:         "abc-123",
			DetailType: "EC2 Instance State-change Notification",
			Source:     "aws.ec2",
			AccountID:  "123456789012",
			Time:       time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			Region:     "eu-west-1",
			Resources:  []string{"arn:aws:ec2:eu-west-1:123456789012:instance/i-abc123"},
			Detail:     json.RawMessage(`{"state":"running"}`),
		},
	}
}

func TestMask_ExternalStruct_Propagation(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	result := masker.Mask(fixture).(cloudWatchWrapper)

	t.Run("userid_masked", func(t *testing.T) {
		if !validateStringMask(result.UserID, fixture.UserID, testConfig.MaxPiiStringLength) {
			t.Errorf("UserID: expected masked, got %q", result.UserID)
		}
	})

	t.Run("external_strings_masked", func(t *testing.T) {
		e := result.Event
		if !validateStringMask(e.Version, fixture.Event.Version, testConfig.MaxPiiStringLength) {
			t.Errorf("Version: expected masked, got %q", e.Version)
		}
		if !validateStringMask(e.ID, fixture.Event.ID, testConfig.MaxPiiStringLength) {
			t.Errorf("ID: expected masked, got %q", e.ID)
		}
		if !validateStringMask(e.DetailType, fixture.Event.DetailType, testConfig.MaxPiiStringLength) {
			t.Errorf("DetailType: expected masked, got %q", e.DetailType)
		}
		if !validateStringMask(e.Source, fixture.Event.Source, testConfig.MaxPiiStringLength) {
			t.Errorf("Source: expected masked, got %q", e.Source)
		}
		if !validateStringMask(e.AccountID, fixture.Event.AccountID, testConfig.MaxPiiStringLength) {
			t.Errorf("AccountID: expected masked, got %q", e.AccountID)
		}
		if !validateStringMask(e.Region, fixture.Event.Region, testConfig.MaxPiiStringLength) {
			t.Errorf("Region: expected masked, got %q", e.Region)
		}
	})

	t.Run("external_slice_masked", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if !validateStringMask(r, fixture.Event.Resources[i], testConfig.MaxPiiStringLength) {
				t.Errorf("Resources[%d]: expected masked, got %q", i, r)
			}
		}
	})
}

func TestMask_ExternalStruct_Show(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	input := cloudWatchWrapperShow{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(cloudWatchWrapperShow)

	t.Run("userid_preserved", func(t *testing.T) {
		if result.UserID != fixture.UserID {
			t.Errorf("UserID: want %q, got %q", fixture.UserID, result.UserID)
		}
	})

	t.Run("external_strings_preserved", func(t *testing.T) {
		e := result.Event
		if e.Version != fixture.Event.Version {
			t.Errorf("Version: want %q, got %q", fixture.Event.Version, e.Version)
		}
		if e.ID != fixture.Event.ID {
			t.Errorf("ID: want %q, got %q", fixture.Event.ID, e.ID)
		}
		if e.DetailType != fixture.Event.DetailType {
			t.Errorf("DetailType: want %q, got %q", fixture.Event.DetailType, e.DetailType)
		}
		if e.Source != fixture.Event.Source {
			t.Errorf("Source: want %q, got %q", fixture.Event.Source, e.Source)
		}
		if e.AccountID != fixture.Event.AccountID {
			t.Errorf("AccountID: want %q, got %q", fixture.Event.AccountID, e.AccountID)
		}
		if e.Region != fixture.Event.Region {
			t.Errorf("Region: want %q, got %q", fixture.Event.Region, e.Region)
		}
	})

	t.Run("external_slice_preserved", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if r != fixture.Event.Resources[i] {
				t.Errorf("Resources[%d]: want %q, got %q", i, fixture.Event.Resources[i], r)
			}
		}
	})
}

func TestMask_ExternalStruct_Anonymize(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	input := cloudWatchWrapperAnonymize{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(cloudWatchWrapperAnonymize)

	t.Run("userid_anonymized", func(t *testing.T) {
		if !validateAnonymization(result.UserID, fixture.UserID, testConfig.MaxPiiStringLength) {
			t.Errorf("UserID: want anonymized string of len %d, got %q", len(fixture.UserID), result.UserID)
		}
	})

	t.Run("external_strings_anonymized", func(t *testing.T) {
		e := result.Event
		if !validateAnonymization(e.Version, fixture.Event.Version, testConfig.MaxPiiStringLength) {
			t.Errorf("Version: want anonymized string of len %d, got %q", len(fixture.Event.Version), e.Version)
		}
		if !validateAnonymization(e.ID, fixture.Event.ID, testConfig.MaxPiiStringLength) {
			t.Errorf("ID: want anonymized string of len %d, got %q", len(fixture.Event.ID), e.ID)
		}
		if !validateAnonymization(e.Source, fixture.Event.Source, testConfig.MaxPiiStringLength) {
			t.Errorf("Source: want anonymized string of len %d, got %q", len(fixture.Event.Source), e.Source)
		}
		if !validateAnonymization(e.AccountID, fixture.Event.AccountID, testConfig.MaxPiiStringLength) {
			t.Errorf("AccountID: want anonymized string of len %d, got %q", len(fixture.Event.AccountID), e.AccountID)
		}
		if !validateAnonymization(e.Region, fixture.Event.Region, testConfig.MaxPiiStringLength) {
			t.Errorf("Region: want anonymized string of len %d, got %q", len(fixture.Event.Region), e.Region)
		}
	})

	t.Run("external_slice_anonymized", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if !validateAnonymization(r, fixture.Event.Resources[i], testConfig.MaxPiiStringLength) {
				t.Errorf("Resources[%d]: want anonymized string of len %d, got %q", i, len(fixture.Event.Resources[i]), r)
			}
		}
	})
}

func TestMask_ExternalStruct_NoTag(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	input := cloudWatchWrapperNoTag{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(cloudWatchWrapperNoTag)

	t.Run("userid_masked", func(t *testing.T) {
		if !validateStringMask(result.UserID, fixture.UserID, testConfig.MaxPiiStringLength) {
			t.Errorf("UserID: expected masked, got %q", result.UserID)
		}
	})

	t.Run("external_strings_masked", func(t *testing.T) {
		e := result.Event
		if !validateStringMask(e.Version, fixture.Event.Version, testConfig.MaxPiiStringLength) {
			t.Errorf("Version: expected masked, got %q", e.Version)
		}
		if !validateStringMask(e.ID, fixture.Event.ID, testConfig.MaxPiiStringLength) {
			t.Errorf("ID: expected masked, got %q", e.ID)
		}
		if !validateStringMask(e.DetailType, fixture.Event.DetailType, testConfig.MaxPiiStringLength) {
			t.Errorf("DetailType: expected masked, got %q", e.DetailType)
		}
		if !validateStringMask(e.Source, fixture.Event.Source, testConfig.MaxPiiStringLength) {
			t.Errorf("Source: expected masked, got %q", e.Source)
		}
		if !validateStringMask(e.AccountID, fixture.Event.AccountID, testConfig.MaxPiiStringLength) {
			t.Errorf("AccountID: expected masked, got %q", e.AccountID)
		}
		if !validateStringMask(e.Region, fixture.Event.Region, testConfig.MaxPiiStringLength) {
			t.Errorf("Region: expected masked, got %q", e.Region)
		}
	})

	t.Run("external_slice_masked", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if !validateStringMask(r, fixture.Event.Resources[i], testConfig.MaxPiiStringLength) {
				t.Errorf("Resources[%d]: expected masked, got %q", i, r)
			}
		}
	})
}
