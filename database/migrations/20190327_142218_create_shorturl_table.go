package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateShorturlTable_20190327_142218 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateShorturlTable_20190327_142218{}
	m.Created = "20190327_142218"

	migration.Register("CreateShorturlTable_20190327_142218", m)
}

// Run the migrations
func (m *CreateShorturlTable_20190327_142218) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	m.SQL("CREATE TABLE `tb_shorturl` (`id` int(11) unsigned NOT NULL,`hash_id` char(10) DEFAULT NULL COMMENT '哈希ID',`url` varchar(255) DEFAULT NULL COMMENT 'URL地址',`host` varchar(255) DEFAULT NULL COMMENT '域名',`created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',`create_user_id` varchar(50) DEFAULT NULL COMMENT '创建用户',PRIMARY KEY (`id`),KEY `url` (`url`),KEY `hash_id` (`hash_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
}

// Reverse the migrations
func (m *CreateShorturlTable_20190327_142218) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
