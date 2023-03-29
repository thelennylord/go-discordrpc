package discordrpc

import (
	"strconv"
	"time"
)

// Activity holds the data for discord rich presence
type Activity struct {
	Details string `json:"details,omitempty"`
	State   string `json:"state,omitempty"`

	Timestamps *Timestamps `json:"timestamps,omitempty"`
	Assets     *Assets     `json:"assets,omitempty"`
	Party      *Party      `json:"party,omitempty"`
	Secrets    *Secrets    `json:"secrets,omitempty"`
	Buttons    []*Button   `json:"buttons,omitempty"`

	Instance bool `json:"instance"`
}

// Timestamps holds unix timestamps for start and/or end of the game
type Timestamps struct {
	Start *Epoch `json:"start,omitempty"`
	End   *Epoch `json:"end,omitempty"`
}

// Epoch wrapper around time.Time to ensure times are sent as a unix epoch int
type Epoch struct{ time.Time }

// MarshalJSON converts time.Time to unix time int
func (t Epoch) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

// Assets passes image references for inclusion in rich presence
type Assets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

// Party holds information for the current party of the player
type Party struct {
	ID   string `json:"id"`
	Size []int  `json:"size"` // seems to be element [0] is count and [1] is max
}

// Secrets holds secrets for Rich Presence joining and spectating
type Secrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type Button struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}

// SetActivity sets the Rich Presence activity for the running application
func (c *Client) SetActivity(activity Activity) error {
	return c.sendCommand(SetActivityCommand, activity)
}
