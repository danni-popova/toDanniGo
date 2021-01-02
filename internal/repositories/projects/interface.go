package projects

type Repository interface {
	Create(userID int, project Project) (Project, error)
	List(userID int) (projects []Project, err error)
	AddMember(projID, userID int) error
}
