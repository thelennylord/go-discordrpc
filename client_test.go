package discordrpc

import (
	"os"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	drpc, err := New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		t.Fatalf("failed to connect to discord ipc: %v", err)
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
		t.Fatalf("failed to set activity: %v", err)
	}
}
