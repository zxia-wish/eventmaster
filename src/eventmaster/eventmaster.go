package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ContextLogic/eventmaster/eventmaster"
	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type dbConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Keyspace    string `json:"keyspace"`
	Consistency string `json:"consistency"`
}

func startUIServer(store *EventStore) {
	funcMap := template.FuncMap{
		"parseTime": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format(time.ANSIC)
		},
		"getVisibility": func(pd pageData, id string) string {
			if id == "topicFilter" && pd.Topic != "" {
				return "visible"
			} else if id == "hostFilter" && pd.Host != "" {
				return "visible"
			}
			return "hidden"
		},
	}
	mux := http.NewServeMux()
	mph := &mainPageHandler{
		store: store,
		fm:    funcMap,
	}
	geh := &getEventHandler{
		store: store,
		fm:    funcMap,
	}
	cph := &createPageHandler{
		store: store,
		fm:    funcMap,
	}
	ceh := &createEventHandler{
		store: store,
		fm:    funcMap,
	}
	mux.Handle("/", mph)
	mux.Handle("/get_events", geh)
	mux.Handle("/create", cph)
	mux.Handle("/create_event", ceh)
	go func() {
		fmt.Println("uiserver starting on port 8080")
		http.ListenAndServe(":8081", mux)
	}()
}

func startAPIServer(store *EventStore) {
	mux := http.NewServeMux()
	eah := &eventAPIHandler{
		store: store,
	}
	mux.Handle("/v1/event", eah)
	go func() {
		fmt.Println("http server starting on port 8080")
		http.ListenAndServe(":8080", mux)
	}()
}

func main() {
	var config Config
	parser := flags.NewParser(&config, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	exp.Exp(metrics.DefaultRegistry)
	sock, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		fmt.Println("go-metrics server listening at port 12345")
		http.Serve(sock, nil)
	}()

	// Create listening socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := metrics.NewGauge()
	metrics.Register("goroutines", g)
	go func() {
		for {
			g.Update(int64(runtime.NumGoroutine()))
			time.Sleep(time.Duration(10) * time.Second)
		}
	}()

	store, err := NewEventStore()
	if err != nil {
		log.Fatalf("Unable to create event store: %v", err)
	}
	startUIServer(store)
	startAPIServer(store)

	// Create the EventMaster server
	server, err := NewServer(&config, store)
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	maxMsgSizeOpt := grpc.MaxMsgSize(1024 * 1024 * 100)
	// Create the gRPC server and register our service
	s := grpc.NewServer(maxMsgSizeOpt)
	eventmaster.RegisterEventMasterServer(s, server)
	reflection.Register(s)
	fmt.Println("grpc server listening on port:", config.Port)
	go func() {
		if err := s.Serve(lis); err != nil {
			// Because we graceful stop, just log this out
			// GracefulStop will kill lis, but we should not
			// throw an error to let it shut down gracefully
			fmt.Println("failed to serve:", err)

		}
	}()

	<-stopChan
	fmt.Println("Got shutdown signal, gracefully shutting down")
	store.CloseSession()
	s.GracefulStop()
}
