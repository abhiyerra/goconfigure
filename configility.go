package configility

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
)

type ParseType int

const (
	Env ParseType = iota
	Etcd
)

type Config struct {
	EtcdClient    etcd.Client
	EtcdNamespace string
	ParseType     ParseType
	keys          map[string]Key
}

type Key struct {
	Description  string
	DefaultValue string
	EnvKey       string
	EtcdKey      string
}

func (c *Config) SetParser(parseType ParseType) {
	c.ParseType = parseType
}

func (c *Config) Add(name, desc, defaultValue, envKey, etcdKey string) {
	// Initialize the keys map if it isn't set.
	if c.keys == nil {
		c.keys = make(map[string]Key)
	}

	c.keys[name] = Key{
		Description:  desc,
		DefaultValue: defaultValue,
		EnvKey:       envKey,
		EtcdKey:      etcdKey,
	}
}

func (c *Config) Get(key string) string {
	k := c.keys[key]

	switch c.ParseType {
	case Env:
		value := os.Getenv(k.EnvKey)
		if value == "" {
			return k.DefaultValue
		}
	case Etcd:
		resp, err := c.EtcdClient.Get(
			fmt.Sprintf("/%s/%s", c.EtcdNamespace, k),
			false,
			false)
		if err != nil {
			return k.DefaultValue
		}

		return resp.Node.Value
	}

	return ""
}
