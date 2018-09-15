package main

import (
	_ "net/http/pprof"

	"../../jenkins-x/test/myliles"
	"../../jenkins-x/test/myliles/myliles/cmd"
	"../../jenkins-x/test/myliles/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.MylilesServer{}

	lile.Name("myliles")
	lile.Server(func(g *grpc.Server) {
		myliles.RegisterMylilesServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
