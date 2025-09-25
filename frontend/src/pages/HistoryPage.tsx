import { useFetch } from '@/hooks/useFetch';
import type { Evaluation } from '@/types/evaluation';
import { EvaluationCard } from '../components/EvaluationCard.tsx';
import './historyPage.scss';

export const HistoryPage = () => {
    const { data: evaluations, isLoading, hasError } = useFetch<Evaluation[]>('/api/evaluations');

    if (isLoading) return <p className={"loadingMessage"}>Cargando historial...</p>;
    if (hasError) return <p className={"errorMessage"}>Error loading history</p>;

    return (
        <div className={"historyContainer"}>
            <div className={"historyHeader"}>
                <div>
                    <h2>Evaluation History</h2>
                    <p>{evaluations?.length ?? 0} evaluation(s) performed</p>
                </div>
            </div>

            <div className={"evaluationsList"}>
                {evaluations && evaluations.length > 0 ? (
                    evaluations.map(evaluation => (
                        <EvaluationCard key={evaluation.id} evaluation={evaluation} />
                    ))
                ) : (
                    <p>No evaluations in history yet.</p>
                )}
            </div>
        </div>
    );
};