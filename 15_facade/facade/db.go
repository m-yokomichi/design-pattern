package facade

type DB struct {
	connect string // 接続先
}

func (db *DB) Select(sql string) string {
	// クエリ結果を返す（処理はダミー）
	return "太郎"
}