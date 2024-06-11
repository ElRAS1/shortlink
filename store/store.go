package store

type Store struct {
	Db    *database
	Cache *cache
	Data  *data
}
