-- +goose Up
-- +goose StatementBegin
INSERT INTO public.authors ("name", bio) VALUES('Heitor', 'ABC');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM public.authors WHERE "name" = 'Heitor';
-- +goose StatementEnd
