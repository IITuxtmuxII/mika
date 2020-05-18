package http

import (
	"context"
	"fmt"
	"github.com/leighmacdonald/mika/examples/api"
	"github.com/leighmacdonald/mika/store"
	"github.com/leighmacdonald/mika/store/memory"
	"os"
	"testing"
	"time"
)

const (
	listenAddr = "localhost:35999"
	pathPrefix = "/testing"
)

func TestUserStore(t *testing.T) {
	store.TestUserStore(t, NewUserStore(listenAddr))
}

func TestPeerStore(t *testing.T) {
	store.TestPeerStore(t, NewPeerStore(listenAddr), memory.NewTorrentStore(), memory.NewUserStore())
}

func TestTorrentStore(t *testing.T) {
	store.TestTorrentStore(t, NewTorrentStore(listenAddr))
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	server := api.New(listenAddr, pathPrefix, api.DefaultAuthKey)
	go func() { _ = server.ListenAndServe() }()
	defer func() { _ = server.Shutdown(ctx) }()
	// Give the server enough time to init
	time.Sleep(500 * time.Millisecond)
	retVal := m.Run()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down test api server: %s", err.Error())
	}
	os.Exit(retVal)
}
