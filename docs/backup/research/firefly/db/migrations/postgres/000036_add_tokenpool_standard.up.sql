BEGIN;
ALTER TABLE tokenpool ADD COLUMN standard VARCHAR(64);
COMMIT;