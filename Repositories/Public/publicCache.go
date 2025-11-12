package Public

type publicCacheRepository interface {
	FindKeysRedis(formatKey string) (keys []string, err error)
}

func (p public) FindKeysRedis(formatKey string) (keys []string, err error) {
	return
}
