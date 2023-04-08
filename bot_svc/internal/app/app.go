package app

import (
	"bot_svc/internal/config"
	"bot_svc/internal/service"
	"bot_svc/internal/storage"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type App struct {
	db     *sqlx.DB
	server *grpc.Server
	quit   chan os.Signal
}

func NewApp(conf *config.Conf) *App {
	a := App{
		db:   newStorage(&conf.Db),
		quit: make(chan os.Signal, 1),
	}
	signal.Notify(a.quit, os.Interrupt)

	// storage
	st := storage.New(a.db)
	// sevice
	svc := service.New(st)
	// grpc server
	a.server = newGrpc(svc)

	return &a
}

func (a *App) ListenAndServe(conf *config.Conf) {
	lis, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		log.Fatalf("failed to listen on port: %v\n", err)
	}

	log.Printf("server started %s/...\n", conf.Addr)

	go func() {
		if err := a.server.Serve(lis); err != nil {
			log.Printf("serve: on start: %v\n", err)
		}
		a.quit <- os.Interrupt
	}()
	<-a.quit
}

func (a *App) Stop() {
	log.Println("stopping server: shut down env...")
	if a.server != nil {
		a.server.GracefulStop()
	}
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			log.Printf("db: stop: %v\n", err)
		}
	}
}
