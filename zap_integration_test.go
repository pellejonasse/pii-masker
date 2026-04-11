package piimasker_test

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type zapUser struct {
	ID    string `Pii:"show"`
	Email string `Pii:"mask"`
	Age   int    `Pii:"mask"`
	Name  string `Pii:"anonymize"`
}

func TestZapIntegration(t *testing.T) {
	masker := newTestMasker(t)

	user := zapUser{
		ID:    "u-123",
		Email: "alice@example.com",
		Age:   31,
		Name:  "Alice",
	}

	// Build an in-memory observed zap logger so we can inspect what was logged.
	core, logs := observer.New(zapcore.InfoLevel)
	logger := zap.New(zapcore.NewTee(core, zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)))

	logger.Info("processing user", zap.Any("user", masker.Mask(user)))

	if logs.Len() != 1 {
		t.Fatalf("expected 1 log entry, got %d", logs.Len())
	}

	fields := logs.All()[0].ContextMap()
	logged, ok := fields["user"]
	if !ok {
		t.Fatal("expected 'user' field in log entry")
	}

	logged_user, ok := logged.(zapUser)
	if !ok {
		t.Fatalf("expected logged 'user' to be of type zapUser, got %T", logged)
	}

	// ID tagged show — must be original value
	if logged_user.ID != user.ID {
		t.Errorf("ID: expected %q (show), got %q", user.ID, logged_user.ID)
	}

	// Email tagged mask — must be all '*', same length
	if !validateStringMask(logged_user.Email, user.Email, testConfig.MaxPiiStringLength) {
		t.Errorf("Email: expected masked string, got %q", logged_user.Email)
	}

	// Age tagged mask — must be zero
	if logged_user.Age != 0 {
		t.Errorf("Age: expected 0 (masked), got %d", logged_user.Age)
	}

	// Name tagged anonymize — must be same length but different value
	if !validateAnonymization(logged_user.Name, user.Name, testConfig.MaxPiiStringLength) {
		t.Errorf("Name: expected anonymized string (different, same length), got %q", logged_user.Name)
	}

	// Sanity check: original is not mutated
	if user.Email != "alice@example.com" {
		t.Error("original user.Email was mutated")
	}
	if user.Age != 31 {
		t.Error("original user.Age was mutated")
	}
}
