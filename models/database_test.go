package models

import (
	"testing"
)

// Example user data
const (
	testUsername = "testuser"
	testPassword = "secret123"
)

func TestAddUser(t *testing.T) {
	db := NewDB()

	user, err := db.Add(testUsername, testPassword)
	if err != nil {
		t.Fatalf("expected no error from Add, got %v", err)
	}

	// Check that the user is not nil
	if user == nil {
		t.Fatal("expected user to be created, got nil")
	}

	// Check that the name and ID were set correctly
	if user.Name != testUsername {
		t.Errorf("expected name %s, got %s", testUsername, user.Name)
	}
	if user.ID != 1 {
		t.Errorf("expected ID 1, got %d", user.ID)
	}

	// Check if the password hash is valid (matches the original password)
	ok, err := db.Validate(user.ID, testPassword)
	if err != nil || !ok {
		t.Errorf("expected password to validate correctly, got err: %v, ok: %v", err, ok)
	}
}

func TestValidateFailsOnWrongPassword(t *testing.T) {
	db := NewDB()
	user, _ := db.Add("badtester", "correct-password")

	ok, err := db.Validate(user.ID, "wrong-password")
	if ok || err == nil {
		t.Errorf("expected validation to fail on wrong password, got ok=%v, err=%v", ok, err)
	}
}

func TestDeleteUser(t *testing.T) {
	db := NewDB()

	// Add a user first
	user, err := db.Add("deleteme", "password123")
	if err != nil {
		t.Fatalf("failed to add user: %v", err)
	}

	// Delete the user
	deleted := db.DeleteUserFromID(user.ID)
	if deleted == nil {
		t.Fatal("expected a deleted user, got nil")
	}

	// Check that returned user matches
	if deleted.ID != user.ID || deleted.Name != user.Name {
		t.Errorf("deleted user does not match original; got %+v, want %+v", deleted, user)
	}

	// Check that user is actually removed
	if found := db.GetUserFromID(user.ID); found != nil {
		t.Error("expected user to be removed from ID map, but found")
	}
	if found := db.GetUserFromName(user.Name); found != nil {
		t.Error("expected user to be removed from name map, but found")
	}
}

func TestDeleteNonexistentUser(t *testing.T) {
	db := NewDB()

	// Try deleting a user that doesn't exist
	deleted := db.DeleteUserFromID(999)
	if deleted != nil {
		t.Errorf("expected nil when deleting nonexistent user, got %+v", deleted)
	}
}

func TestValidateCorrectPassword(t *testing.T) {
	db := NewDB()
	user, err := db.Add("validuser", "securepass")
	if err != nil {
		t.Fatalf("failed to add user: %v", err)
	}

	ok, err := db.Validate(user.ID, "securepass")
	if err != nil {
		t.Errorf("unexpected error from Validate: %v", err)
	}
	if !ok {
		t.Error("expected password validation to succeed, got false")
	}
}

func TestValidateWrongPassword(t *testing.T) {
	db := NewDB()
	user, _ := db.Add("wrongpassuser", "correctpass")

	ok, err := db.Validate(user.ID, "incorrectpass")
	if err == nil {
		t.Error("expected error from Validate on wrong password, got nil")
	}
	if ok {
		t.Error("expected password validation to fail, got true")
	}
}

func TestValidateNonExistentUser(t *testing.T) {
	db := NewDB()

	ok, err := db.Validate(999, "anything")
	if err == nil {
		t.Error("expected error for nonexistent user, got nil")
	}
	if ok {
		t.Error("expected validation to fail for nonexistent user, got true")
	}
}

func TestValidateDeletedUser(t *testing.T) {
	db := NewDB()
	user, _ := db.Add("deletetest", "to-be-removed")

	db.DeleteUserFromID(user.ID)

	ok, err := db.Validate(user.ID, "to-be-removed")
	if err == nil {
		t.Error("expected error when validating deleted user, got nil")
	}
	if ok {
		t.Error("expected validation to fail for deleted user, got true")
	}
}
