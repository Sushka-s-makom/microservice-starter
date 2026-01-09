package buildinfo

// Эти значения по умолчанию видны при go run.
// При сборке их можно заменить через -ldflags "-X ...=..."
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
