package server

import "github.com/google/wire"

// ProviderSet for cron job ...
var ProviderSet = wire.NewSet(NewMainCron, NewServer)
