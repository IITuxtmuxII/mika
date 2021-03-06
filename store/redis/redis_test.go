package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/leighmacdonald/mika/config"
	"github.com/leighmacdonald/mika/store"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRedisTorrentStore(t *testing.T) {
	ts, e := store.NewStore(config.TorrentStore)
	require.NoError(t, e, e)
	store.TestStore(t, ts)
}

func clearDB(c *redis.Client) {
	keys, err := c.Keys("*").Result()
	if err != nil {
		log.Panicf("Could not initialize redis db: %s", err.Error())
	}
	for _, k := range keys {
		c.Del(k)
	}
}

func setupDB(t *testing.T, c *redis.Client) {
	clearDB(c)
	t.Cleanup(func() {
		clearDB(c)
	})
}

func TestMain(m *testing.M) {
	if err := config.Read("mika_testing_redis"); err != nil {
		log.Info("Skipping database tests, failed to find config: mika_testing_redis.yaml")
		os.Exit(0)
		return
	}
	if config.General.RunMode != "test" {
		log.Info("Skipping database tests, not running in testing mode")
		os.Exit(0)
		return
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}
