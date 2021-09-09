/*
monolith-microservice-shop

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineObject struct for InlineObject
type InlineObject struct {
	ProductId int64 `json:"product_id"`
	Address OrdersAddress `json:"address"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject(productId int64, address OrdersAddress) *InlineObject {
	this := InlineObject{}
	this.ProductId = productId
	this.Address = address
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *InlineObject) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *InlineObject) GetProductIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *InlineObject) SetProductId(v int64) {
	o.ProductId = v
}

// GetAddress returns the Address field value
func (o *InlineObject) GetAddress() OrdersAddress {
	if o == nil {
		var ret OrdersAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineObject) GetAddressOk() (*OrdersAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineObject) SetAddress(v OrdersAddress) {
	o.Address = v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product_id"] = o.ProductId
	}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


