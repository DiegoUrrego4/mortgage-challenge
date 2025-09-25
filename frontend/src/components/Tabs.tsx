import { NavLink } from 'react-router-dom';
import { FaPencilAlt, FaHistory } from 'react-icons/fa';
import './tabs.scss';

export const Tabs = () => {
    return (
        <div className="tabs-container">
            <NavLink to="/" className="tab">
                <FaPencilAlt />
                <span>Evaluation</span>
            </NavLink>
            <NavLink to="/history" className="tab">
                <FaHistory />
                <span>History</span>
            </NavLink>
        </div>
    );
};
