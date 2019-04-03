package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTokenTable_20190403_135232 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTokenTable_20190403_135232{}
	m.Created = "20190403_135232"

	migration.Register("CreateTokenTable_20190403_135232", m)
}

// Run the migrations
func (m *CreateTokenTable_20190403_135232) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	m.SQL("CREATE TABLE `tb_token` (`id` char(32) NOT NULL DEFAULT '',`name` varchar(32) DEFAULT NULL,`token` char(64) DEFAULT NULL,`created_at` timestamp NULL DEFAULT NULL,PRIMARY KEY (`id`),KEY `token` (`token`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
}

// Reverse the migrations
func (m *CreateTokenTable_20190403_135232) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
