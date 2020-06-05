package repo

type IUser interface {
	GetUser(login, pass string) (map[string]string, error) 
}

