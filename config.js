/***
 * 如果开启了mongoDB,将下面代码注释去掉，
 * 并将dbUserName, dbPassword和dbName都
 * 替换成分配得到的值。即可查看 mongoDB
 * 测试程序。否则则开启hello world程序。
 ***/
/*
var mongo = require("mongoskin");
var db_url = exports.db_url = "dbUserName:dbPassword@127.0.0.1:20088/dbName";
exports.db = mongo.db(db_url);
*/
var mongo = require("mongoskin");
var db_url = exports.db_url = "8b2qv5kfx2imv:7okgyd5s6d2@127.0.0.1:20088/x7mxzznp7i2of";
exports.db = mongo.db(db_url);
