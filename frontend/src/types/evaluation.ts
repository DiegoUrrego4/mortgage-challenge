export interface Evaluation {
    id:              number;
    monthly_income:   number;
    monthly_debts:    number;
    loan_amount:      number;
    property_value:   number;
    credit_score:     number;
    occupancy_type:   string;
    decision:        'Approve' | 'Refer' | 'Decline';
    dti:             number;
    ltv:             number;
    reasons:         string[];
    created_at:       string;
}