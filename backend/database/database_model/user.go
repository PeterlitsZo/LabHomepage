package databaseModel

const (
	UnKnow = "unknow"
	Admin  = "admin"
	Normal = "normal"
)

var (
	Role2RoleId = map[string]int32{
		UnKnow: 0,
		Admin:  1,
		Normal: 2,
	}
	RoleId2Role = map[int32]string{
		0: UnKnow,
		1: Admin,
		2: Normal,
	}
)

type User struct {
	ID       uint   `gorm:"primaryKey;index;column:id"`
	Name     string `gorm:"unique;uniqueIndex;column:name"`
	Password string `gorm:"column:password"`
	RoleId   int32  `gorm:"column:role_id;"`
	Extra    string `gorm:"column:extra"`
}

func (u *User) IsAdmin() bool {
	return RoleId2Role[u.RoleId] == Admin
}

type UserFilter User
