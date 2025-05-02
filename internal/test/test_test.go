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
