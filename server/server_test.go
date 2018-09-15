package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"../../jenkins-x/test/myliles"
	"github.com/lileio/lile"
)

var s = MylilesServer{}
var cli myliles.MylilesClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		myliles.RegisterMylilesServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = myliles.NewMylilesClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
