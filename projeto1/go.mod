module projeto1

require internal/crypt v0.0.0

require internal/fmts v0.0.0

replace internal/crypt => ./internal/crypt

replace internal/fmts => ./internal/fmts

go 1.16
