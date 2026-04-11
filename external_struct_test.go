package piimasker_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type CloudWatchWrapper struct {
	Event  events.CloudWatchEvent `Pii:"mask"`
	UserID string                 `Pii:"mask"`
}

type CloudWatchWrapperShow struct {
	Event  events.CloudWatchEvent `Pii:"show"`
	UserID string                 `Pii:"show"`
}

type CloudWatchWrapperAnonymize struct {
	Event  events.CloudWatchEvent `Pii:"anonymize"`
	UserID string                 `Pii:"anonymize"`
}

type CloudWatchWrapperNoTag struct {
	Event  events.CloudWatchEvent
	UserID string
}

func newCloudWatchWrapperFixture() CloudWatchWrapper {
	return CloudWatchWrapper{
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
	result := masker.Mask(fixture).(CloudWatchWrapper)

	t.Run("userid_masked", func(t *testing.T) {
		if result.UserID != "" {
			t.Errorf("UserID: want %q, got %q", "", result.UserID)
		}
	})

	t.Run("external_strings_masked", func(t *testing.T) {
		e := result.Event
		if e.Version != "" {
			t.Errorf("Version: want %q, got %q", "", e.Version)
		}
		if e.ID != "" {
			t.Errorf("ID: want %q, got %q", "", e.ID)
		}
		if e.DetailType != "" {
			t.Errorf("DetailType: want %q, got %q", "", e.DetailType)
		}
		if e.Source != "" {
			t.Errorf("Source: want %q, got %q", "", e.Source)
		}
		if e.AccountID != "" {
			t.Errorf("AccountID: want %q, got %q", "", e.AccountID)
		}
		if e.Region != "" {
			t.Errorf("Region: want %q, got %q", "", e.Region)
		}
	})

	t.Run("external_slice_masked", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if r != "" {
				t.Errorf("Resources[%d]: want %q, got %q", i, "", r)
			}
		}
	})
}

func TestMask_ExternalStruct_Show(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	input := CloudWatchWrapperShow{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(CloudWatchWrapperShow)

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
	input := CloudWatchWrapperAnonymize{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(CloudWatchWrapperAnonymize)

	t.Run("userid_anonymized", func(t *testing.T) {
		if got := result.UserID; len(got) != len(fixture.UserID) || got == fixture.UserID {
			t.Errorf("UserID: want anonymized string of len %d, got %q", len(fixture.UserID), got)
		}
	})

	t.Run("external_strings_anonymized", func(t *testing.T) {
		e := result.Event
		if got := e.Version; len(got) != len(fixture.Event.Version) || got == fixture.Event.Version {
			t.Errorf("Version: want anonymized string of len %d, got %q", len(fixture.Event.Version), got)
		}
		if got := e.ID; len(got) != len(fixture.Event.ID) || got == fixture.Event.ID {
			t.Errorf("ID: want anonymized string of len %d, got %q", len(fixture.Event.ID), got)
		}
		if got := e.Source; len(got) != len(fixture.Event.Source) || got == fixture.Event.Source {
			t.Errorf("Source: want anonymized string of len %d, got %q", len(fixture.Event.Source), got)
		}
		if got := e.AccountID; len(got) != len(fixture.Event.AccountID) || got == fixture.Event.AccountID {
			t.Errorf("AccountID: want anonymized string of len %d, got %q", len(fixture.Event.AccountID), got)
		}
		if got := e.Region; len(got) != len(fixture.Event.Region) || got == fixture.Event.Region {
			t.Errorf("Region: want anonymized string of len %d, got %q", len(fixture.Event.Region), got)
		}
	})

	t.Run("external_slice_anonymized", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			orig := fixture.Event.Resources[i]
			if len(r) != len(orig) || r == orig {
				t.Errorf("Resources[%d]: want anonymized string of len %d, got %q", i, len(orig), r)
			}
		}
	})
}

func TestMask_ExternalStruct_NoTag(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newCloudWatchWrapperFixture()
	input := CloudWatchWrapperNoTag{
		UserID: fixture.UserID,
		Event:  fixture.Event,
	}
	result := masker.Mask(input).(CloudWatchWrapperNoTag)

	t.Run("userid_masked", func(t *testing.T) {
		if result.UserID != "" {
			t.Errorf("UserID: want %q, got %q", "", result.UserID)
		}
	})

	t.Run("external_strings_masked", func(t *testing.T) {
		e := result.Event
		if e.Version != "" {
			t.Errorf("Version: want %q, got %q", "", e.Version)
		}
		if e.ID != "" {
			t.Errorf("ID: want %q, got %q", "", e.ID)
		}
		if e.DetailType != "" {
			t.Errorf("DetailType: want %q, got %q", "", e.DetailType)
		}
		if e.Source != "" {
			t.Errorf("Source: want %q, got %q", "", e.Source)
		}
		if e.AccountID != "" {
			t.Errorf("AccountID: want %q, got %q", "", e.AccountID)
		}
		if e.Region != "" {
			t.Errorf("Region: want %q, got %q", "", e.Region)
		}
	})

	t.Run("external_slice_masked", func(t *testing.T) {
		for i, r := range result.Event.Resources {
			if r != "" {
				t.Errorf("Resources[%d]: want %q, got %q", i, "", r)
			}
		}
	})
}
