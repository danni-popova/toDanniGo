package account

type Repository interface {
	InsertAccountData(acc AccountData) error
	SelectAuthDetails(email string) (authDetails AuthDetails, err error)
	SelectAccountDetails(id int) (acc AccountData, err error)
}
