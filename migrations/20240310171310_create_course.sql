-- +goose Up

CREATE TABLE courses (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name varchar(255) NOT NULL,
    platform varchar(255) NOT NULL,
    price integer NOT NULL,
    CONSTRAINT "pk_id" PRIMARY KEY (id ASC)
);

Insert into courses (name, platform, price) values ('JavaScript', 'YouTube', 5000);

-- +goose Down
DROP TABLE courses;