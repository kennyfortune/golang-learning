package mocklearning

// DB 假设 DB 是代码中负责与数据库交互的部分(在这里用 map 模拟)，测试用例中不能创建真实的数据库连接。这个时候，如果我们需要测试 GetFromDB 这个函数内部的逻辑，就需要 mock 接口 DB。
type DB interface {
	Get(key string) (int, error)
}

// mockgen -source ./mocklearning/db.go -destination ./mocklearning/db_mock.go -package mocklearning
func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err != nil {
		return value
	}
	return -1
}
