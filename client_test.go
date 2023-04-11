package discordrpc

import (
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	drpc, err := New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}
	defer drpc.Socket.Close()
}

func TestSetActivity(t *testing.T) {
	drpc, err := New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}
	defer drpc.Socket.Close()

	err = drpc.SetActivity(Activity{
		Details: "Foo",
		State:   "Bar",
		Assets: &Assets{
			SmallImage: "keyart_hero",
			LargeImage: "keyart_hero",
		},
		Timestamps: &Timestamps{
			Start: &Epoch{Time: time.Now()},
		},
		Buttons: []*Button{
			{
				Label: "Button 1",
				Url:   "https://www.google.com",
			},
			{
				Label: "Button 2",
				Url:   "https://www.youtube.com",
			},
		},
	})

	if err != nil {
		t.Fatalf("Failed to set activity: %v", err)
	}
}

func TestRegisterGame(t *testing.T) {
	drpc, err := New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}
	defer drpc.Socket.Close()

	err = drpc.RegisterCommand("foo://bar", "")
	if err != nil {
		t.Fatalf("Failed to register game: %v", err)
	}
}
