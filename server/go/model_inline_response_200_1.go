/*
 * monolith-microservice-shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse2001 struct {

	Id int64 `json:"id,omitempty"`

	IsPaid bool `json:"is_paid,omitempty"`
}

// AssertInlineResponse2001Required checks if the required fields are not zero-ed
func AssertInlineResponse2001Required(obj InlineResponse2001) error {
	return nil
}

// AssertRecurseInlineResponse2001Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineResponse2001 (e.g. [][]InlineResponse2001), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineResponse2001Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineResponse2001, ok := obj.(InlineResponse2001)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineResponse2001Required(aInlineResponse2001)
	})
}
