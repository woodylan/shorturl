package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateJumpLogTable_20190403_135355 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateJumpLogTable_20190403_135355{}
	m.Created = "20190403_135355"

	migration.Register("CreateJumpLogTable_20190403_135355", m)
}

// Run the migrations
func (m *CreateJumpLogTable_20190403_135355) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `tb_jump_log` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT,`url_id` int(11) DEFAULT NULL COMMENT '短网址ID',`user_agent` varchar(255) DEFAULT NULL COMMENT 'UA',`ip` char(15) DEFAULT NULL COMMENT 'IP地址',`referer` varchar(255) DEFAULT NULL COMMENT '来源',`created_at` timestamp NULL DEFAULT NULL COMMENT '访问时间',PRIMARY KEY (`id`),KEY `url_id` (`url_id`),KEY `created_at` (`created_at`)) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;")

}

// Reverse the migrations
func (m *CreateJumpLogTable_20190403_135355) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
