import type { Evaluation } from '@/types/evaluation';
import { FaUserCircle } from 'react-icons/fa';

// Props que el componente recibe
interface Props {
    evaluation: Evaluation;
}

// Función de ayuda para formatear la fecha a un formato legible
const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString('en-US', {
        year: 'numeric', month: 'long', day: 'numeric',
        hour: '2-digit', minute: '2-digit', hour12: false
    });
};

// Función de ayuda para formatear números como moneda USD sin decimales
const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
        minimumFractionDigits: 0,
        maximumFractionDigits: 0,
    }).format(amount);
};

export const EvaluationCard = ({ evaluation }: Props) => {
    // Generamos el nombre de la clase para el badge dinámicamente
    // ej. "badge-approve", "badge-decline"
    const decisionClass = `badge-${evaluation.decision.toLowerCase()}`;

    return (
        <div className="evaluationCard">
            {/* --- Parte Superior: Usuario y Decisión --- */}
            <div className="cardTop">
                <div className="userInfo">
                    <FaUserCircle size={40} className="avatar" />
                    <div>
                        {/* El backend no nos da un nombre, usamos el ID por ahora */}
                        <div className="userName">Application #{evaluation.id}</div>
                        <div className="date">{formatDate(evaluation.created_at)}</div>
                    </div>
                </div>
                <div className={`badge ${decisionClass}`}>
                    {evaluation.decision}
                </div>
            </div>

            {/* --- Parte Media: Datos Financieros --- */}
            <div className="cardData">
                <div className="dataPoint">
                    <span className="dataLabel">Loan</span>
                    <strong className="dataValue">{formatCurrency(evaluation.loan_amount)}</strong>
                </div>
                <div className="dataPoint">
                    <span className="dataLabel">DTI</span>
                    <strong className="dataValue">{(evaluation.dti * 100).toFixed(1)}%</strong>
                </div>
                <div className="dataPoint">
                    <span className="dataLabel">LTV</span>
                    <strong className="dataValue">{(evaluation.ltv * 100).toFixed(0)}%</strong>
                </div>
                <div className="dataPoint">
                    <span className="dataLabel">Credit Score</span>
                    <strong className="dataValue">{evaluation.credit_score}</strong>
                </div>
            </div>

            {/* --- Parte Inferior: Razones --- */}
            <div className="cardReasons">
                <strong>Decision Reasons:</strong>
                <ul>
                    {evaluation.reasons.map((reason, index) => <li key={index}>{reason}</li>)}
                </ul>
            </div>
        </div>
    );
};