-- =========================
-- Удаление индексов
-- =========================
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_projects_user_id;
DROP INDEX IF EXISTS idx_project_blocks_block_type_id;
DROP INDEX IF EXISTS idx_project_blocks_project_id;
DROP INDEX IF EXISTS idx_block_types_tag_id;

-- =========================
-- Удаление таблиц
-- =========================
DROP TABLE IF EXISTS project_blocks;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS block_types;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS users;
