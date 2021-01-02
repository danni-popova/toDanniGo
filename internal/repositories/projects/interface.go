package projects

type Repository interface {
	Create(project Project) (Project, error)
	List(userID int) (projects []Project, err error)
	AddMember(projID, userID int) error
}
