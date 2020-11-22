// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Trait - represents a wit-ai trait.
//
// https://wit.ai/docs/http/20200513/#post__traits_link
type Trait struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Values []TraitValue `json:"values"`
}

// TraitValue - represents the value of a Trait.
//
// https://wit.ai/docs/http/20200513/#get__traits__trait_link
type TraitValue struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// GetTraits - returns a list of traits.
//
// https://wit.ai/docs/http/20200513/#get__traits_link
func (c *Client) GetTraits() ([]Trait, error) {
	resp, err := c.request(http.MethodGet, "/traits", "application/json", nil)
	if err != nil {
		return []Trait{}, err
	}

	defer resp.Close()

	var traits []Trait
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&traits)
	return traits, err
}

// CreateTrait - creates a new trait with the given values.
//
// https://wit.ai/docs/http/20200513/#post__traits_link
func (c *Client) CreateTrait(name string, values []string) (*Trait, error) {
	type trait struct {
		Name   string   `json:"name"`
		Values []string `json:"values"`
	}

	traitJSON, err := json.Marshal(trait{Name: name, Values: values})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/traits", "application/json", bytes.NewBuffer(traitJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var traitResp *Trait
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&traitResp)
	return traitResp, err
}

// GetTrait - returns all available information about a trait.
//
// https://wit.ai/docs/http/20200513/#get__traits__trait_link
func (c *Client) GetTrait(name string) (*Trait, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/traits/%s", url.PathEscape(name)), "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	var traitResp *Trait
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&traitResp)
	return traitResp, err
}

// DeleteTrait - permanently deletes a trait.
//
// https://wit.ai/docs/http/20200513/#delete__traits__trait_link
func (c *Client) DeleteTrait(name string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/traits/%s", url.PathEscape(name)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// AddTraitValue - create a new value for a trait.
//
// https://wit.ai/docs/http/20200513/#post__traits__trait_values_link
func (c *Client) AddTraitValue(traitName string, value string) (*Trait, error) {
	type traitValue struct {
		Value string `json:"value"`
	}

	valueJSON, err := json.Marshal(traitValue{Value: value})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/traits/%s/values", url.PathEscape(traitName)), "application/json", bytes.NewBuffer(valueJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var traitResp *Trait
	decoder := json.NewDecoder(resp)
	if err = decoder.Decode(&traitResp); err != nil {
		return nil, err
	}

	return traitResp, nil
}

// DeleteTraitValue - permanently deletes the trait value.
//
// https://wit.ai/docs/http/20200513/#delete__traits__trait_values__value_link
func (c *Client) DeleteTraitValue(traitName string, value string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/traits/%s/values/%s", url.PathEscape(traitName), url.PathEscape(value)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}
