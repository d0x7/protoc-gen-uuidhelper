package test

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"testing"
	gen "xiam.li/uuidhelper/internal/test/gen/go"
)

func TestUUID(t *testing.T) {
	sess, internal := uuid.Must(uuid.NewRandom()), uuid.Must(uuid.NewRandom())
	player := &gen.Player{}
	player.SetSessionUUID(sess)
	player.SetInternalUUID(internal)

	bytes, err := proto.Marshal(player)
	if err != nil {
		t.Fatalf("Failed to marshal player: %v", err)
	}
	newPlayer := &gen.Player{}
	if err := proto.Unmarshal(bytes, newPlayer); err != nil {
		t.Fatalf("Failed to unmarshal player: %v", err)
	}
	if newPlayer.GetSessionUUID() != sess {
		t.Fatalf("Session UUID mismatch: expected %v, got %v", sess, newPlayer.GetSessionUUID())
	}
	if newPlayer.GetInternalUUID() != internal {
		t.Fatalf("Internal UUID mismatch: expected %v, got %v", internal, newPlayer.GetInternalUUID())
	}
}

func TestUUIDs(t *testing.T) {
	uuids := []uuid.UUID{uuid.Must(uuid.NewRandom()), uuid.Must(uuid.NewRandom()), uuid.Must(uuid.NewRandom()), uuid.Must(uuid.NewRandom())}
	player := &gen.Player{}
	player.SetGameUUIDs(uuids[:2])
	player.AddGameUUIDs(uuids[2], uuids[3])

	bytes, err := proto.Marshal(player)
	if err != nil {
		t.Fatalf("Failed to marshal player: %v", err)
	}
	newPlayer := &gen.Player{}
	if err := proto.Unmarshal(bytes, newPlayer); err != nil {
		t.Fatalf("Failed to unmarshal player: %v", err)
	}
	if len(newPlayer.GetGameUUIDs()) != len(uuids) {
		t.Fatalf("UUIDs length mismatch: expected %d, got %d", len(uuids), len(newPlayer.GetGameUUIDs()))
	}
	for i, uid := range uuids {
		if newPlayer.GetGameUUIDs()[i] != uid {
			t.Fatalf("UUID mismatch at index %d: expected %v, got %v", i, uid, newPlayer.GetGameUUIDs()[i])
		}
	}
}
