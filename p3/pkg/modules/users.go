package modules

type User struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Age        int    `db:"age" json:"age"`
	Hobby      string `db:"hobby" json:"hobby"`
	Profession string `db:"profession" json:"profession"`
}
