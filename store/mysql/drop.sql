DROP PROCEDURE IF EXISTS user_by_passkey; -- STMT
DROP PROCEDURE IF EXISTS user_by_id; -- STMT
DROP PROCEDURE IF EXISTS user_delete; -- STMT
DROP PROCEDURE IF EXISTS user_add; -- STMT
DROP PROCEDURE IF EXISTS user_update; -- STMT
DROP PROCEDURE IF EXISTS user_update_stats; -- STMT
DROP PROCEDURE IF EXISTS torrent_by_infohash; -- STMT
DROP PROCEDURE IF EXISTS torrent_delete; -- STMT
DROP PROCEDURE IF EXISTS torrent_disable; -- STMT
DROP PROCEDURE IF EXISTS torrent_add; -- STMT
DROP PROCEDURE IF EXISTS torrent_update_stats; -- STMT
DROP PROCEDURE IF EXISTS whitelist_all; -- STMT
DROP PROCEDURE IF EXISTS whitelist_add; -- STMT
DROP PROCEDURE IF EXISTS whitelist_delete_by_prefix; -- STMT
DROP PROCEDURE IF EXISTS peer_update_stats; -- STMT
DROP PROCEDURE IF EXISTS peer_reap; -- STMT
DROP PROCEDURE IF EXISTS peer_add; -- STMT
DROP PROCEDURE IF EXISTS peer_delete; -- STMT
DROP PROCEDURE IF EXISTS peer_get; -- STMT
DROP PROCEDURE IF EXISTS peer_get_n; -- STMT
DROP TABLE IF EXISTS user_multi; -- STMT
DROP TABLE IF EXISTS user_roles; -- STMT
DROP TABLE IF EXISTS roles; -- STMT
DROP TABLE IF EXISTS peers; -- STMT
DROP TABLE IF EXISTS users; -- STMT
DROP TABLE IF EXISTS whitelist; -- STMT
DROP TABLE IF EXISTS torrent; -- STMT
