package migrations

type IMigrations interface {
	Up() error
	Down() error
}