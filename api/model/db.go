package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DNS struct {
	ServerID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServerTag       string             `json:"server_tag,omitempty" bson:"server_tag,omitempty"`
	Address         string             `json:"address,omitempty" bson:"address,omitempty"`
	AddressResolver string             `json:"address_resolver,omitempty" bson:"address_resolver,omitempty"`
	AddressStrategy string             `json:"address_strategy,omitempty" bson:"address_strategy,omitempty"`
	Detour          string             `json:"detour,omitempty" bson:"detour,omitempty"`
}

type Route struct {
	RuleID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rules  []Rule             `json:"rules,omitempty" bson:"rules,omitempty"`
	Detour string             `json:"detour,omitempty" bson:"detour,omitempty"`
}

type Rule struct {
	// Strategy includes "geo*" etc.
	Strategy string `json:"strategy,omitempty" bson:"strategy,omitempty"`
	// Contents includes the rule content, like "google" etc.
	Contents []string `json:"contents,omitempty" bson:"contents,omitempty"`
}
