package email

type Repository interface {
	InsertVerificationRecord(record VerificationRecord) error
	UpdateVerificationRecord(recordUUID string) error
}
