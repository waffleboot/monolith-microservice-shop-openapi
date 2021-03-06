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

// OrdersAddress struct for OrdersAddress
type OrdersAddress struct {
	Name string `json:"name"`
	Street string `json:"street"`
	City string `json:"city"`
	PostCode string `json:"post_code"`
	Country string `json:"country"`
}

// NewOrdersAddress instantiates a new OrdersAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersAddress(name string, street string, city string, postCode string, country string) *OrdersAddress {
	this := OrdersAddress{}
	this.Name = name
	this.Street = street
	this.City = city
	this.PostCode = postCode
	this.Country = country
	return &this
}

// NewOrdersAddressWithDefaults instantiates a new OrdersAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersAddressWithDefaults() *OrdersAddress {
	this := OrdersAddress{}
	return &this
}

// GetName returns the Name field value
func (o *OrdersAddress) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrdersAddress) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrdersAddress) SetName(v string) {
	o.Name = v
}

// GetStreet returns the Street field value
func (o *OrdersAddress) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *OrdersAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *OrdersAddress) SetStreet(v string) {
	o.Street = v
}

// GetCity returns the City field value
func (o *OrdersAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *OrdersAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *OrdersAddress) SetCity(v string) {
	o.City = v
}

// GetPostCode returns the PostCode field value
func (o *OrdersAddress) GetPostCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
func (o *OrdersAddress) GetPostCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostCode, true
}

// SetPostCode sets field value
func (o *OrdersAddress) SetPostCode(v string) {
	o.PostCode = v
}

// GetCountry returns the Country field value
func (o *OrdersAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *OrdersAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *OrdersAddress) SetCountry(v string) {
	o.Country = v
}

func (o OrdersAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["post_code"] = o.PostCode
	}
	if true {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableOrdersAddress struct {
	value *OrdersAddress
	isSet bool
}

func (v NullableOrdersAddress) Get() *OrdersAddress {
	return v.value
}

func (v *NullableOrdersAddress) Set(val *OrdersAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersAddress(val *OrdersAddress) *NullableOrdersAddress {
	return &NullableOrdersAddress{value: val, isSet: true}
}

func (v NullableOrdersAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


