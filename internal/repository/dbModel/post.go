package dbModel

type Post struct {
	Title    string `db:"title"`
	Body     string `db:"body"`
	ImageURL string `db:"image"`
	Author   string `db:"author"`
}
