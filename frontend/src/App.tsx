import { Outlet } from 'react-router-dom';
import { Tabs } from './components/Tabs.tsx';
import './App.scss';

function App() {
    return (
        <div className="page-container">
            <header className="page-header">
                <h1>Mortgage Subscription Engine</h1>
                <p>Automatically evaluates mortgage loan applications based on DTI, LTV, and credit score</p>
            </header>

            <Tabs />

            <main className="floating-card">
                <Outlet />
            </main>
        </div>
    );
}

export default App;