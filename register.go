//go:build windows

package discordrpc

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// Registers a command by which Discord can launch your game.
func (c *Client) RegisterCommand(command string, defaultIcon string) error {
	rootKey := fmt.Sprintf(`SOFTWARE\Classes\discord-%s`, c.ClientID)
	defaultIconKey := fmt.Sprintf(`%s\DefaultIcon`, rootKey)
	shellCommandKey := fmt.Sprintf(`%s\shell\open\command`, rootKey)

	// Create `SOFTWARE\Classes\discord-CLIENT_ID` key and register it as a URL Protocol
	newK, _, err := registry.CreateKey(registry.CURRENT_USER, rootKey, registry.WRITE)
	if err != nil {
		return fmt.Errorf("could not create registry key %s: %v", rootKey, err)
	}
	defer newK.Close()

	err = newK.SetStringValue("", fmt.Sprintf("URL:Run game %s protocol", c.ClientID))
	if err != nil {
		return fmt.Errorf("could not set default value of key %s: %v", rootKey, err)
	}

	err = newK.SetStringValue("URL Protocol", "")
	if err != nil {
		return fmt.Errorf("could not set value 'URL Protocol' of key %s: %v", rootKey, err)
	}

	// Create `SOFTWARE\Classes\discord-CLIENTID\DefaultIcon` key to let Discord know the icon of your game
	// It's the path to the executable
	newK, _, err = registry.CreateKey(registry.CURRENT_USER, defaultIconKey, registry.WRITE)
	if err != nil {
		return fmt.Errorf("could not create registry key %s: %v", defaultIconKey, err)
	}
	defer newK.Close()

	err = newK.SetStringValue("", defaultIcon)
	if err != nil {
		return fmt.Errorf("could not set default value of key %s: %v", defaultIconKey, err)
	}

	// Create `SOFTWARE\Classes\discord-CLIENTID\shell\open\command` key and set the command for Discord to open our game
	newK, _, err = registry.CreateKey(registry.CURRENT_USER, shellCommandKey, registry.WRITE)
	if err != nil {
		return fmt.Errorf("could not create registry key %s: %v", shellCommandKey, err)
	}
	defer newK.Close()

	err = newK.SetStringValue("", command)
	if err != nil {
		return fmt.Errorf("could not set default value of key %s: %v", shellCommandKey, err)
	}

	return nil
}
