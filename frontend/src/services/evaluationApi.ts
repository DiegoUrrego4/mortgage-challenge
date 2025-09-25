interface ApiInput {
    monthly_income: number | undefined;
    monthly_debts: number | undefined;
    loan_amount: number | undefined;
    property_value: number | undefined;
    credit_score: number | undefined;
    occupancy_type: string | undefined;}

export interface ApiResponse {
    decision: 'Approve' | 'Refer' | 'Decline';
    dti: number;
    ltv: number;
    reasons: string[];
}

export const submitEvaluation = async (formData: ApiInput): Promise<ApiResponse> => {
    const response = await fetch('/api/evaluate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData),
    });

    if (!response.ok) {
        throw new Error('Network response was not ok');
    }

    return response.json();
};