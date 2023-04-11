//go:build linux

package discordrpc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Reference: https://github.com/discord/discord-rpc/blob/master/src/discord_register_linux.cpp
var fileFormat = `[Desktop Entry]
Name=Game %s
Exec=%s %%u
Type=Application
NoDisplay=true
Categories=Discord;Games;
MimeType=x-scheme-handler/discord-%s;
`

// Registers a command by which Discord can launch your game.
func (c *Client) RegisterCommand(command string, defaultIcon string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	path := filepath.Join(homeDir, ".local", "share", "applications")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	fileContent := fmt.Sprintf(fileFormat, c.ClientID, command, c.ClientID)
	fileName := fmt.Sprintf(`discord-%s.desktop`, c.ClientID)

	err = os.WriteFile(filepath.Join(path, fileName), []byte(fileContent), 0755)
	if err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	cmd := exec.Command(
		"xdg-mime",
		"default",
		fileName,
		fmt.Sprintf("x-scheme-handler/discord-%s", c.ClientID),
	)

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("could not register command; got error: %v", err)
	}

	return nil
}
