package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Username is required")
	}

	err := s.cfgPointer.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}
