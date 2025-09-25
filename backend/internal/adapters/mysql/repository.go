package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/DiegoUrrego4/backend/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS evaluations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		monthly_income DECIMAL(15, 2) NOT NULL,
		monthly_debts DECIMAL(15, 2) NOT NULL,
		loan_amount DECIMAL(15, 2) NOT NULL,
		property_value DECIMAL(15, 2) NOT NULL,
		credit_score INT NOT NULL,
		occupancy_type VARCHAR(50) NOT NULL,
		decision VARCHAR(10) NOT NULL,
		dti DECIMAL(5, 4) NOT NULL,
		ltv DECIMAL(5, 4) NOT NULL,
		reasons TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Save(eval domain.Evaluation) error {
	reasonsJSON, _ := json.Marshal(eval.Reasons)

	query := `INSERT INTO evaluations (monthly_income, monthly_debts, loan_amount, property_value, credit_score, occupancy_type, decision, dti, ltv, reasons, created_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, eval.MonthlyIncome, eval.MonthlyDebts, eval.LoanAmount, eval.PropertyValue, eval.CreditScore, eval.OccupancyType, eval.Decision, eval.DTI, eval.LTV, string(reasonsJSON), eval.CreatedAt)
	return err
}

func (r *Repository) GetAll() ([]domain.Evaluation, error) {
	rows, err := r.db.Query("SELECT id, monthly_income, monthly_debts, loan_amount, property_value, credit_score, occupancy_type, decision, dti, ltv, reasons, created_at FROM evaluations ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	evaluations := make([]domain.Evaluation, 0)

	for rows.Next() {
		var eval domain.Evaluation
		var reasonsJSON string

		err := rows.Scan(
			&eval.ID, &eval.MonthlyIncome, &eval.MonthlyDebts, &eval.LoanAmount,
			&eval.PropertyValue, &eval.CreditScore, &eval.OccupancyType, &eval.Decision,
			&eval.DTI, &eval.LTV, &reasonsJSON, &eval.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		json.Unmarshal([]byte(reasonsJSON), &eval.Reasons)
		evaluations = append(evaluations, eval)
	}
	return evaluations, nil
}
