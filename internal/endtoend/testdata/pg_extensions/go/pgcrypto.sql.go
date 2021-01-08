// Code generated by sqlc. DO NOT EDIT.
// source: pgcrypto.sql

package querytest

import (
	"context"
)

const encodeDigest = `-- name: EncodeDigest :one
SELECT encode(digest($1, 'sha1'), 'hex')
`

func (q *Queries) EncodeDigest(ctx context.Context, digest string) (string, error) {
	row := q.db.QueryRowContext(ctx, encodeDigest, digest)
	var encode string
	err := row.Scan(&encode)
	return encode, err
}
