/*
 * Covidify
 *
 * Simple API collecting guest data.
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package covidify

import (
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func (s *Server) NewRouter() *gin.Engine {
	router := s.NewRouterWithMiddleware(nil, nil, nil)

	return router
}

func (s *Server) NewRouterWithMiddleware(pre, mid, post []gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	router.Use(pre...)
	router.Use(ginlogrus.Logger(s.config.Logger))
	router.Use(gin.Recovery())

	router.Use(mid...)

	router.GET("/health", s.Health)
	router.POST("/visit", s.AddVisit)
	router.GET("/visit/:visitID", s.CheckVisit)
	router.POST("/report/visitor", s.AddReportVisitor)

	router.Use(post...)

	return router
}
