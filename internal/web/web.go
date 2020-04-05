//Package web will serve the web server part
package web

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/opengintracing"
	"github.com/gin-gonic/gin"
	myopentracing "github.com/jsenon/compagnyhelper/internal/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"github.com/rs/zerolog/log"
)

const nbrTimeOut time.Duration = 5

// Serve launch http server
func Serve() {
	log.Info().Msg("Startin Web Server on port 8080")

	var tracer opentracing.Tracer

	if !viper.GetBool("DISABLETRACE") {
		var closer io.Closer

		var err error

		jaeger := viper.GetString("JAEGERURL")

		tracer, closer, err = myopentracing.ConfigureTracing(jaeger)
		if err != nil {
			log.Fatal().Msgf("Can't start: %v", err)
		}

		setupRoutes(tracer)

		defer closer.Close() //nolint: errcheck
	}

	setupRoutes(tracer)
}

func setupRoutes(tracer opentracing.Tracer) {
	log.Debug().Msgf("tracer: %v", tracer)

	if tracer == nil {
		tracer = opentracing.GlobalTracer()
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger.SetLogger())

	prometheus := ginprometheus.NewPrometheus("compagnyhelper")
	prometheus.Use(router)

	setRouterAPI(router, tracer)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //nolint: staticcheck
	<-quit
	log.Info().Msg("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), nbrTimeOut*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown: %v", err)
	}

	log.Info().Msg("Server exiting")
}

func setRouterAPI(router *gin.Engine, tracer opentracing.Tracer) {
	router.GET("/healthz",
		opengintracing.NewSpan(tracer, "healthz"),
		opengintracing.InjectToHeaders(tracer, true),
		healthz)

	// router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/retrieve-links",
		opengintracing.NewSpan(tracer, "retrieve-links"),
		opengintracing.InjectToHeaders(tracer, true),
		retrieveLinks)

	router.POST("/retrieve-link/:name",
		opengintracing.NewSpan(tracer, "retrieve-link"),
		opengintracing.InjectToHeaders(tracer, true),
		retrievedescribeLink)

	router.POST("/describe-link/:name",
		opengintracing.NewSpan(tracer, "describe-link"),
		opengintracing.InjectToHeaders(tracer, true),
		retrievedescribeLink)

	router.POST("/open-link/:name",
		opengintracing.NewSpan(tracer, "open-link"),
		opengintracing.InjectToHeaders(tracer, true),
		openLink)
}
