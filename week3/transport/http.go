package transport

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func corsHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	return
}

// corsMiddleware allows cors
func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			corsHandler(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func grpcGateway(ctx context.Context, grpcPort string) (http.Handler, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterFeedbackServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts)

	if err != nil {
		return nil, err
	}

	return mux, nil
}

// RunHTTPServer runs HTTP gateway
func RunHTTPServer(ctx context.Context, grpcPort string, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gw, err := grpcGateway(ctx, grpcPort)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: corsMiddleware(mux),
	}

	// graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("Shutting down HTTP server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown HTTP server: %v", err)
		}
	}()

	log.Println("Starting HTTP server...")
	return srv.ListenAndServe()
}
