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
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddReportVisitor - Report an infected visitor
func (s *Server) AddReportVisitor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
