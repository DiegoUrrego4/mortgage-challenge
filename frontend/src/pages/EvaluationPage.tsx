import { useNavigate } from 'react-router-dom';
import { FaPerson, FaHouseChimney, FaWpforms } from "react-icons/fa6";
import { CiMoneyBill } from "react-icons/ci";
import { useForm } from "@/hooks/useForm.ts";
import { InputField } from "@/components/InputField.tsx";
import { submitEvaluation } from '@/services/evaluationApi.ts';
import "./evaluationPage.scss";

interface FormFields {
    monthly_income: number;
    monthly_debts: number;
    loan_amount: number;
    property_value: number;
    credit_score: number;
    occupancy_type: string;
}

export const EvaluationPage = () => {
    const navigate = useNavigate();

    const { formState, onInputChange } = useForm<FormFields>({
        occupancy_type: "Primary",
        monthly_income: 0,
        monthly_debts: 0,
        credit_score: 0,
        loan_amount: 0,
        property_value: 0
    });

    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const apiData = {
                monthly_income:   Number(formState.monthly_income || 0),
                monthly_debts:    Number(formState.monthly_debts || 0),
                loan_amount:      Number(formState.loan_amount || 0),
                property_value:   Number(formState.property_value || 0),
                credit_score:     Number(formState.credit_score || 0),
                occupancy_type:   formState.occupancy_type,
            };

            await submitEvaluation(apiData);
            navigate('/history');
        } catch (err) {
            console.error(err);
            alert('Hubo un error al enviar la evaluación.');
        }
    };

    return (
        <form onSubmit={handleFormSubmit} className="evaluation-form">
            <div className="card-header">
                <div className="icon-background">
                    <FaWpforms size={24} />
                </div>
                <h2>Mortgage Loan Application</h2>
                <p>Complete the borrower's information to assess loan eligibility.</p>
            </div>

            <div className="form-section">
                <div className="section-title">
                    <FaPerson/>
                    <p>Occupancy Information</p>
                </div>
                <div className="input-group">
                    <label htmlFor="occupancy_type">Occupancy Type</label>
                    <select
                        id="occupancy_type"
                        name="occupancy_type"
                        value={formState.occupancy_type}
                        onChange={onInputChange}
                    >
                        <option value="Primary">Primary Residence</option>
                        <option value="Secondary">Secondary Residence</option>
                        <option value="Investment">Investment</option>
                    </select>
                </div>
            </div>

            <div className="form-section">
                <div className="section-title">
                    <CiMoneyBill/>
                    <p>Financial Information</p>
                </div>
                <div className="grid-2-cols">
                    <InputField label="Monthly Income" name="monthly_income" type="number" value={formState.monthly_income} onChange={onInputChange} />
                    <InputField label="Monthly Debts" name="monthly_debts" type="number" value={formState.monthly_debts} onChange={onInputChange} />
                    <InputField label="Credit Score" name="credit_score" type="number" value={formState.credit_score} onChange={onInputChange} className="full-width" />
                </div>
            </div>

            <div className="form-section">
                <div className="section-title">
                    <FaHouseChimney/>
                    <p>Property Information</p>
                </div>
                <div className="grid-2-cols">
                    <InputField label="Loan Amount" name="loan_amount" type="number" value={formState.loan_amount} onChange={onInputChange} />
                    <InputField label="Property Value" name="property_value" type="number" value={formState.property_value} onChange={onInputChange} />
                </div>
            </div>

            <button type="submit" className="submit-button">Evaluate Loan</button>
        </form>
    );
};