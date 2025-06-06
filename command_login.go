package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login needs a username")
	}

	config := s.cfg
	err := config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %s\n", cmd.args[0])

	return nil
}
