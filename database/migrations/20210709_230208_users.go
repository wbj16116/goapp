package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20210709_230208 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20210709_230208{}
	m.Created = "20210709_230208"

	migration.Register("Users_20210709_230208", m)
}

// Run the migrations
func (m *Users_20210709_230208) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) DEFAULT NULL,`username` varchar(128) NOT NULL,`phone` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`password` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *Users_20210709_230208) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
