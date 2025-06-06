package main

import "fmt"

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.commands[cmd.name]
	if !exists {
		return fmt.Errorf("the command %s does not exist", cmd.name)
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
