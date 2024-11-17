package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clc3-CloudComputing/clc3-ws23/4-Continuous-Deployment/exercise4.1/src/internal/format"
)

const plainOutput = `Good day! Time now: %02d:%02d`
const htmlOutput = `
<html>
<h1>Good day!</h1>
<body>
Time now: %02d:%02d
</body>
</html>
`

type myHandler struct {
}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	var output string
	if format.GetOutputFormat() == format.Html {
		output = htmlOutput
	} else {
		output = plainOutput
	}
	fmt.Fprintf(w, output, t.Hour(), getMinute(t.Minute(), t.Second()))
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

func main() {

	srv := &http.Server{
		Addr:         ":8888",
		Handler:      &myHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		<-sigs

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()

		// We received an interrupt signal, shut down
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
