package main

import (
	"database/sql"
	"github.com/NeroQue/GATOR/internal/database"
	_ "github.com/lib/pq"
)
import (
	"fmt"
	"github.com/NeroQue/GATOR/internal/config"
	"github.com/joho/godotenv"
	"os"

	"log"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dbURL)

	dbQueries := database.New(db)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: not enough arguments\n")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cfg.DBURL = dbURL
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("register", handlerRegister)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("reset", HandlerReset)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("users", handlerListUsers)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("agg", handlerAgg)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("feeds", handlerListFeeds)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("follow", middlewareLoggedIn(handlerFollowFeed))
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("following", middlewareLoggedIn(handlerFollowingFeedsByUser))
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
