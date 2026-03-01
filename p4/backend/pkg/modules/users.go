package modules

type User struct {
	Id         int    `db:"id" json:"id" redis:"id"`
	Name       string `db:"name" json:"name" redis:"name"`
	Age        int    `db:"age" json:"age" redis:"age"`
	Hobby      string `db:"hobby" json:"hobby" redis:"hobby"`
	Profession string `db:"profession" json:"profession" redis:"profession"`
}
