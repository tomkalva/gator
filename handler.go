package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
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

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.cfgPointer.CurrentUserName == user.Name {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", *feed)
	return nil
}

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Feed name and url are required")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("ID:", feed.ID)
	fmt.Println("CreatedAt:", feed.CreatedAt)
	fmt.Println("UpdatedAt:", feed.UpdatedAt)
	fmt.Println("Name:", feed.Name)
	fmt.Println("Url:", feed.Url)
	fmt.Println("UserID:", feed.UserID)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		user, err := s.db.GetUserFromId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Println("--------------------------")
		fmt.Println("Feed:", i+1)
		fmt.Println("ID:", feed.ID)
		fmt.Println("CreatedAt:", feed.CreatedAt)
		fmt.Println("UpdatedAt:", feed.UpdatedAt)
		fmt.Println("Name:", feed.Name)
		fmt.Println("Url:", feed.Url)
		fmt.Println("Created by:", user.Name)
		fmt.Println("UserID:", feed.UserID)
		fmt.Println("--------------------------")
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Url is required")
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Feed follow created successfully")
	fmt.Println("Feed name:", feedFollow.FeedName)
	fmt.Println("Current user:", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range feedFollows {
		fmt.Println(follow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Url is required")
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
