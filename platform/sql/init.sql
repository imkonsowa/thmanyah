CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS unaccent;

CREATE TYPE category_type AS ENUM (
    'CATEGORY_TYPE_PODCAST',
    'CATEGORY_TYPE_DOCUMENTARY',
    'CATEGORY_TYPE_SPORTS_EVENT',
    'CATEGORY_TYPE_EDUCATIONAL',
    'CATEGORY_TYPE_NEWS',
    'CATEGORY_TYPE_ENTERTAINMENT'
    );

CREATE TYPE program_status AS ENUM (
    'PROGRAM_STATUS_DRAFT',
    'PROGRAM_STATUS_PUBLISHED',
    'PROGRAM_STATUS_ARCHIVED'
    );

CREATE TYPE episode_status AS ENUM (
    'EPISODE_STATUS_DRAFT',
    'EPISODE_STATUS_PUBLISHED',
    'EPISODE_STATUS_SCHEDULED',
    'EPISODE_STATUS_ARCHIVED'
    );

CREATE TYPE import_status AS ENUM (
    'IMPORT_STATUS_PENDING',
    'IMPORT_STATUS_PROCESSING',
    'IMPORT_STATUS_COMPLETED',
    'IMPORT_STATUS_FAILED'
    );

CREATE TABLE IF NOT EXISTS users
(
    id         uuid primary key,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,

    name       text      not null,
    email      text      not null unique,
    password   text      null
);

CREATE TABLE IF NOT EXISTS categories
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(255)  NOT NULL,
    description TEXT,
    type        category_type NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    created_by  VARCHAR(255),
    metadata    JSONB     DEFAULT '{}'::jsonb,
    UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS programs
(
    id             UUID PRIMARY KEY,
    title          text   NOT NULL,
    description    TEXT,
    category_id    UUID           NOT NULL REFERENCES categories (id),
    status         program_status NOT NULL DEFAULT 'PROGRAM_STATUS_DRAFT',
    created_at     TIMESTAMP               DEFAULT NOW(),
    updated_at     TIMESTAMP               DEFAULT NOW(),
    published_at   TIMESTAMP,
    created_by     UUID           NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    updated_by     UUID           NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    thumbnail_url  TEXT,
    tags           TEXT[]                  DEFAULT '{}',
    metadata       JSONB                   DEFAULT '{}'::jsonb,
    source_url     TEXT,
    episodes_count INTEGER                 DEFAULT 0,
    is_featured    BOOLEAN                 DEFAULT FALSE,
    view_count     INTEGER                 DEFAULT 0,
    rating         DECIMAL(3, 2)           DEFAULT 0.0,
    search_vector  TSVECTOR
);

CREATE TABLE IF NOT EXISTS episodes
(
    id               UUID PRIMARY KEY,
    program_id       UUID           NOT NULL REFERENCES programs (id) ON DELETE CASCADE,
    title            TEXT   NOT NULL,
    description      TEXT,
    duration_seconds INTEGER                 DEFAULT 0,
    episode_number   INTEGER        NOT NULL,
    season_number    INTEGER        NOT NULL DEFAULT 1,
    status           episode_status NOT NULL DEFAULT 'EPISODE_STATUS_DRAFT',
    created_at       TIMESTAMP               DEFAULT NOW(),
    updated_at       TIMESTAMP               DEFAULT NOW(),
    published_at     TIMESTAMP,
    scheduled_at     TIMESTAMP,
    created_by       UUID           NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    updated_by       UUID           NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    media_url        TEXT,
    thumbnail_url    TEXT,
    tags             TEXT[]                  DEFAULT '{}',
    metadata         JSONB                   DEFAULT '{}'::jsonb,
    view_count       INTEGER                 DEFAULT 0,
    rating           DECIMAL(3, 2)           DEFAULT 0.0,
    file_size_bytes  BIGINT                  DEFAULT 0,
    search_vector    TSVECTOR,
    UNIQUE (program_id, season_number, episode_number)
);

CREATE TABLE IF NOT EXISTS imports
(
    id              UUID PRIMARY KEY,
    source_type     VARCHAR(50)   NOT NULL,
    source_url      TEXT,
    source_config   JSONB                  DEFAULT '{}'::jsonb,
    category_id     UUID REFERENCES categories (id),
    status          import_status NOT NULL DEFAULT 'IMPORT_STATUS_PENDING',
    total_items     INTEGER                DEFAULT 0,
    processed_items INTEGER                DEFAULT 0,
    success_count   INTEGER                DEFAULT 0,
    error_count     INTEGER                DEFAULT 0,
    errors          TEXT[]                 DEFAULT '{}',
    warnings        TEXT[]                 DEFAULT '{}',
    created_at      TIMESTAMP              DEFAULT NOW(),
    updated_at      TIMESTAMP              DEFAULT NOW(),
    created_by      UUID          NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    updated_by      UUID          NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    field_mapping   JSONB                  DEFAULT '{}'::jsonb, -- this helps to map data from external source structure to internal structure
    metadata        JSONB                  DEFAULT '{}'::jsonb  -- this helps to map data from external source structure to internal structure
);

-- Performance Indexes for Programs table
CREATE INDEX IF NOT EXISTS idx_programs_category_id ON programs (category_id);
CREATE INDEX IF NOT EXISTS idx_programs_status ON programs (status);
CREATE INDEX IF NOT EXISTS idx_programs_featured ON programs (is_featured);
CREATE INDEX IF NOT EXISTS idx_programs_created_at ON programs (created_at);
CREATE INDEX IF NOT EXISTS idx_programs_view_count ON programs (view_count);
CREATE INDEX IF NOT EXISTS idx_programs_search_vector ON programs USING gin (search_vector);
CREATE INDEX IF NOT EXISTS idx_programs_tags_gin ON programs USING gin (tags);

-- Performance Indexes for Episodes table
CREATE INDEX IF NOT EXISTS idx_episodes_program_id ON episodes (program_id);
CREATE INDEX IF NOT EXISTS idx_episodes_status ON episodes (status);
CREATE INDEX IF NOT EXISTS idx_episodes_season ON episodes (season_number);
CREATE INDEX IF NOT EXISTS idx_episodes_number ON episodes (episode_number);
CREATE INDEX IF NOT EXISTS idx_episodes_created_at ON episodes (created_at);
CREATE INDEX IF NOT EXISTS idx_episodes_search_vector ON episodes USING gin (search_vector);

-- Performance Indexes for Categories table
CREATE INDEX IF NOT EXISTS idx_categories_type ON categories (type);
CREATE INDEX IF NOT EXISTS idx_categories_name_gin ON categories USING gin (to_tsvector('english', name));

-- Performance Indexes for Imports table
CREATE INDEX IF NOT EXISTS idx_imports_status ON imports (status);
CREATE INDEX IF NOT EXISTS idx_imports_created_at ON imports (created_at);

-- Functions and triggers for automatic search vector updates
CREATE OR REPLACE FUNCTION update_programs_search_vector()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.search_vector :=
            setweight(to_tsvector('simple', unaccent(coalesce(NEW.title, ''))), 'A') ||
            setweight(to_tsvector('simple', unaccent(coalesce(NEW.description, ''))), 'B') ||
            setweight(to_tsvector('simple', unaccent(array_to_string(NEW.tags, ' '))), 'C');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_episodes_search_vector()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.search_vector :=
            setweight(to_tsvector('simple', unaccent(coalesce(NEW.title, ''))), 'A') ||
            setweight(to_tsvector('simple', unaccent(coalesce(NEW.description, ''))), 'B') ||
            setweight(to_tsvector('simple', unaccent(array_to_string(NEW.tags, ' '))), 'C');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers
CREATE TRIGGER programs_search_vector_update
    BEFORE INSERT OR UPDATE
    ON programs
    FOR EACH ROW
EXECUTE FUNCTION update_programs_search_vector();

CREATE TRIGGER episodes_search_vector_update
    BEFORE INSERT OR UPDATE
    ON episodes
    FOR EACH ROW
EXECUTE FUNCTION update_episodes_search_vector();
