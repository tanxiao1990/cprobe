package tomcat

import (
	"context"

	"github.com/BurntSushi/toml"
	"github.com/cprobe/cprobe/plugins"
	"github.com/cprobe/cprobe/types"
)

func init() {
	plugins.RegisterPlugin(types.PluginTomcat, &Tomcat{})
}

type Tomcat struct{}

func (*Tomcat) ParseConfig(baseDir string, bs []byte) (any, error) {
	var c Config
	err := toml.Unmarshal(bs, &c)
	if err != nil {
		return nil, err
	}
	c.BaseDir = baseDir
	return &c, nil
}

func (*Tomcat) Scrape(ctx context.Context, target string, c any, ss *types.Samples) error {
	cfg := c.(*Config)
	return cfg.Scrape(ctx, target, ss)
}
