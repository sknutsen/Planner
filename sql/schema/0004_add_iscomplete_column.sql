-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks 
ADD COLUMN is_complete integer NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks 
DROP COLUMN is_complete;
-- +goose StatementEnd
