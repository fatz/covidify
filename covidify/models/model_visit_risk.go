/*
 * Covidify
 *
 * Simple API collecting guest data.
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type VisitRisk struct {

	Id int64 `json:"id,omitempty" gorm:"primary_key"`

	Risk string `json:"risk,omitempty" cql:"risk"`

	Description string `json:"description,omitempty" cql:"description"`
}
