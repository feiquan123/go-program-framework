package ops

import "github.com/google/wire"

// ProviderSet is wire provider set of app
var ProviderSet = wire.NewSet(NewAPP)
