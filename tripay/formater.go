package tripay

type UserFormatter struct {
	ID                 int    `json:"id"`
	StudentID          int    `json:"student_id"`
	TeacherID          int    `json:"teacher_id"`
	NominalTransaction int    `json:"nominal_transaction"`
	StatusTransaction  string `json:"status_transaction"`
	AvatarBuktiTf      string `json:"avatar_bukti_tf"`
}

func FormatUser(res ClosedTransactionResponse) UserFormatter {
	formatter := UserFormatter{
		// ID:                 teacher.ID,
		// StudentID:          teacher.StudentID,
		// TeacherID:          teacher.TeacherID,
		// NominalTransaction: teacher.NominalTransaction,
		// StatusTransaction:  teacher.StatusTransaction,
		// AvatarBuktiTf:      teacher.AvatarBuktiTf,
	}

	return formatter
}
