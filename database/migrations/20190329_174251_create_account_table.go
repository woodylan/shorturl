package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAccountTable_20190329_174251 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAccountTable_20190329_174251{}
	m.Created = "20190329_174251"

	migration.Register("CreateAccountTable_20190329_174251", m)
}

// Run the migrations
func (m *CreateAccountTable_20190329_174251) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `tb_account` (`id` char(32) NOT NULL DEFAULT '',`name` varchar(32) DEFAULT NULL,`access_key` char(64) DEFAULT NULL,`secret_key` char(64) DEFAULT NULL,`created_at` timestamp NULL DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")

}

// Reverse the migrations
func (m *CreateAccountTable_20190329_174251) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
