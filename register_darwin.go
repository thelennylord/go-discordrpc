//go:build darwin

package discordrpc

import (
	"fmt"
	"os"
	"path/filepath"
)

// Reference: https://github.com/discord/discord-rpc/blob/master/src/discord_register_osx.m
// NOTE: Needs testing
var fileFormat = `{ "command":	"%s" }`

// Registers a command by which Discord can launch your game.
func (c *Client) RegisterCommand(command string, defaultIcon string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	path := filepath.Join(homeDir, "Library", "Application Support", "discord", "games")

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	fileContent := fmt.Sprintf(fileFormat, command)
	fileName := fmt.Sprintf(`%s.json`, c.ClientID)

	err = os.WriteFile(filepath.Join(path, fileName), []byte(fileContent), 0755)
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	return nil
}
