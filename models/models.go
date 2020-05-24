package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
	Roles     []*Role            `json:"roles,omitempty" bson:"roles,omitempty"`
	Address   *Address           `json:"address,omitempty" bson:"address,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty" bson:"street,omitempty"`
	City   string `json:"city,omitempty" bson:"city,omitempty"`
	Zip    int    `json:"zip,omitempty" bson:"zip,omitempty"`
}

type Role struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rolename    string             `json:"rolename,omitempty" bson:"rolename,omitempty"`
	Permissions []*Permission      `json:"permissions,omitempty" bson:"permissions,omitempty"`
}

type Permission struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Permissionname string             `json:"permissionname,omitempty" bson:"permissionname,omitempty"`
}
