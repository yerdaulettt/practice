package modules

type User struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Age        int    `db:"age"`
	Hobby      string `db:"hobby"`
	Profession string `db:"profession"`
}
