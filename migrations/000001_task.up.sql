-- =========================
-- USERS
-- =========================
CREATE TABLE users
(
    id            serial        NOT NULL UNIQUE,
    name          varchar(255)  NOT NULL,
    username      varchar(255)  NOT NULL UNIQUE,
    password_hash varchar(255)  NOT NULL,
    created_at    timestamp     DEFAULT now(),
    updated_at    timestamp     DEFAULT now(),
    deleted_at    timestamp
);

-- =========================
-- TAGS
-- =========================
CREATE TABLE tags
(
    id          serial      NOT NULL UNIQUE,
    name        text        NOT NULL,
    created_at  timestamp   DEFAULT now(),
    updated_at  timestamp   DEFAULT now(),
    deleted_at  timestamp
);

-- =========================
-- BLOCK TYPES
-- =========================
CREATE TABLE block_types
(
    id          serial       NOT NULL UNIQUE,
    tag_id      integer      REFERENCES tags (id) ON DELETE SET NULL,
    name        text         NOT NULL,
    description text,
    template    text         NOT NULL,
    schema      JSONB        NOT NULL,
    preview     text,
    created_at  timestamp    DEFAULT now(),
    updated_at  timestamp    DEFAULT now(),
    deleted_at  timestamp
);

-- =========================
-- PROJECTS
-- =========================
CREATE TABLE projects
(
    id          serial       NOT NULL UNIQUE,
    user_id     integer      REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    name        text         NOT NULL,
    data        JSONB        DEFAULT '{}'::jsonb,
    created_at  timestamp    DEFAULT now(),
    updated_at  timestamp    DEFAULT now(),
    deleted_at  timestamp
);

-- =========================
-- PROJECT_BLOCKS
-- (Каждый блок, добавленный пользователем в проект)
-- =========================
CREATE TABLE project_blocks
(
    id             serial       NOT NULL UNIQUE,
    project_id     integer      REFERENCES projects (id) ON DELETE CASCADE NOT NULL,
    block_type_id  integer      REFERENCES block_types (id) ON DELETE CASCADE NOT NULL,
    position       integer      DEFAULT 0,
    data           JSONB        DEFAULT '{}'::jsonb,
    created_at     timestamp    DEFAULT now(),
    updated_at     timestamp    DEFAULT now(),
    deleted_at     timestamp
);

-- =========================
-- Индексы
-- =========================
CREATE INDEX idx_block_types_tag_id ON block_types(tag_id);
CREATE INDEX idx_project_blocks_project_id ON project_blocks(project_id);
CREATE INDEX idx_project_blocks_block_type_id ON project_blocks(block_type_id);
CREATE INDEX idx_projects_user_id ON projects(user_id);
CREATE INDEX idx_users_username ON users(username);
