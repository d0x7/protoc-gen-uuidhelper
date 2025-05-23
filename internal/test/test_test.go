package test

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"testing"
	gen "xiam.li/uuidhelper/internal/test/gen/go"
)

func transmit(t *testing.T, player *gen.Player) *gen.Player {
	bytes, err := proto.Marshal(player)
	if err != nil {
		t.Fatalf("Failed to marshal player: %v", err)
	}
	newPlayer := &gen.Player{}
	if err := proto.Unmarshal(bytes, newPlayer); err != nil {
		t.Fatalf("Failed to unmarshal player: %v", err)
	}
	return newPlayer
}

func TestUUID(t *testing.T) {
	sess, internal := uuid.Must(uuid.NewRandom()), uuid.Must(uuid.NewRandom())
	player := &gen.Player{}
	player.SetSessionUUID(sess)
	player.SetInternalUUID(internal)

	newPlayer := transmit(t, player)

	// Check if GetSessionUUID and GetInternalUUID of both players are equal
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

	newPlayer := transmit(t, player)

	// Check if GetGameUUIDs of both players are equal
	if len(newPlayer.GetGameUUIDs()) != len(uuids) {
		t.Fatalf("UUIDs length mismatch: expected %d, got %d", len(uuids), len(newPlayer.GetGameUUIDs()))
	}
	for i, uid := range uuids {
		if newPlayer.GetGameUUIDs()[i] != uid {
			t.Fatalf("UUID mismatch at index %d: expected %v, got %v", i, uid, newPlayer.GetGameUUIDs()[i])
		}
	}
}

func TestMaps(t *testing.T) {
	player := &gen.Player{}

	player.GetMapStringUUIDs()
	player.SetMapStringUUIDs(map[string]uuid.UUID{
		"key1": uuid.Must(uuid.NewRandom()),
		"key2": uuid.Must(uuid.NewRandom()),
	})
	player.GetMapStringUUIDs()

	newPlayer := transmit(t, player)

	// Check if GetMapStringUUIDs of both players are equal
	if len(player.GetMapStringUUIDs()) != len(newPlayer.GetMapStringUUIDs()) {
		t.Fatalf("Map UUIDs length mismatch: expected %d, got %d", len(player.GetMapStringUUIDs()), len(newPlayer.GetMapStringUUIDs()))
	}
	for k, v := range player.GetMapStringUUIDs() {
		if newPlayer.GetMapStringUUIDs()[k] != v {
			t.Fatalf("Map UUID mismatch at key %s: expected %v, got %v", k, v, newPlayer.GetMapStringUUIDs()[k])
		}
	}
}

func TestOptionalUUID(t *testing.T) {
	player := &gen.Player{}

	player.GetOptUUID()
	player.SetOptUUID(uuid.Must(uuid.NewRandom()))
	player.GetOptUUID()
	player.SetOptUUID(uuid.Nil) // TODO: Check if the behavior is correct; when setting the UUID to uuid.Nil, it might be preferable to set it to nil instead of an empty byte slice
	player.GetOptUUID()

	bytes, err := proto.Marshal(player)
	if err != nil {
		t.Fatalf("Failed to marshal player: %v", err)
	}
	newPlayer := &gen.Player{}
	if err := proto.Unmarshal(bytes, newPlayer); err != nil {
		t.Fatalf("Failed to unmarshal player: %v", err)
	}

	// Check if GetOptUUID of both players are equal
	if newPlayer.GetOptUUID() != player.GetOptUUID() {
		t.Fatalf("Optional UUID mismatch: expected %v, got %v", player.GetOptUUID(), newPlayer.GetOptUUID())
	}
}

func TestNestedUUID(t *testing.T) {
	player := &gen.Player{}

	player.Nested = &gen.Player_NestedMsg{}
	player.GetNested().SetNestedUUID(uuid.Must(uuid.NewRandom()))

	newPlayer := transmit(t, player)

	// Check if GetNestedUUID of both players are equal
	if newPlayer.GetNested().GetNestedUUID() != player.GetNested().GetNestedUUID() {
		t.Fatalf("Nested UUID mismatch: expected %v, got %v", player.GetNested().GetNestedUUID(), newPlayer.GetNested().GetNestedUUID())
	}
}
