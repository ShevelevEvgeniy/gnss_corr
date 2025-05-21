package gnss_corr

import "gnss_corr/internal/config/modules"

func (c Config) Service() modules.Service { return c.service }

func (c Config) Postgres() modules.Postgres { return c.storage.postgres }
