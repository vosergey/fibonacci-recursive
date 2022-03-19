module fibonacci-recursive

go 1.17

replace internal/console => ./internal/console
replace internal/fibonacci => ./internal/fibonacci

require (
	internal/console v1.0.0
	internal/fibonacci v1.0.0
)