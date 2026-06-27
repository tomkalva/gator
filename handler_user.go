package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Username is required")
	}

	name := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.cfgPointer.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Username is required")
	}
	name := cmd.arguments[0]

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("User '%v' created\n", user.Name)

	err = s.cfgPointer.SetUser(name)
	if err != nil {
		return err
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database reset successfully")
	return nil
}
