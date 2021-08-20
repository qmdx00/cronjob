package cron

import "github.com/google/wire"

// ProviderSet for cron ...
var ProviderSet = wire.NewSet(NewRootCron)
