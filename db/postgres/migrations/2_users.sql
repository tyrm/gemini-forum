-- +migrate Up
CREATE TABLE "public"."users" (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    certhash character varying NOT NULL,
    username character varying NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    PRIMARY KEY ("id")
);

CREATE INDEX users_certhash_deleted_at_idx ON "public"."users" (certhash, deleted_at);
CREATE INDEX users_username_deleted_at_idx ON "public"."users" (username, deleted_at);

CREATE TABLE "public"."group_membership" (
     id uuid NOT NULL DEFAULT uuid_generate_v4 (),
     user_id uuid NOT NULL,
     group_id uuid NOT NULL,
     created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
     deleted_at timestamp without time zone,
     PRIMARY KEY ("id"),
     FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE "public"."group_membership";
DROP INDEX users_username_deleted_at_idx;
DROP INDEX users_certhash_deleted_at_idx;
DROP TABLE "public"."users";
