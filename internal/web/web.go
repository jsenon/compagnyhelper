package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/opengintracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	myopentracing "github.com/jsenon/compagnyhelper/internal/opentracing"

	"github.com/rs/zerolog/log"
)

// Serve launch http server
func Serve() {
	log.Info().Msg("Startin Web Server on port 8080")
	var tracer opentracing.Tracer
	if !viper.GetBool("DISABLETRACE") {
		jaeger := viper.GetString("JAEGERURL")

		tracer, closer, err := myopentracing.ConfigureTracing(jaeger)
		if err != nil {
			log.Fatal().Msgf("Can't start: %v", err)
		}
		setupRoutes(tracer)
		defer closer.Close()
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

	setRouterApi(router, tracer)

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

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown:", err)
	}

	log.Info().Msg("Server exiting")
}

func setRouterApi(router *gin.Engine, tracer opentracing.Tracer) {

	router.GET("/healthz",
		opengintracing.NewSpan(tracer, "healthz"),
		opengintracing.InjectToHeaders(tracer, true),
		healthz)

	router.POST("/retrieve-links",
		opengintracing.NewSpan(tracer, "retrieve-links"),
		opengintracing.InjectToHeaders(tracer, true),
		retrieveLinks)

	router.POST("/describe-link",
		opengintracing.NewSpan(tracer, "retrieve-links"),
		opengintracing.InjectToHeaders(tracer, true),
		describeLink)

	router.POST("/open-link",
		opengintracing.NewSpan(tracer, "retrieve-links"),
		opengintracing.InjectToHeaders(tracer, true),
		openLink)

}
