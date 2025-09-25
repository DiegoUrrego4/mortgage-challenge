import { NavLink } from 'react-router-dom';
import { FaPencilAlt, FaHistory } from 'react-icons/fa';
import './tabs.scss'; // Crea un archivo SCSS para los tabs

export const Tabs = () => {
    return (
        <div className="tabs-container">
            <NavLink to="/" className="tab">
                <FaPencilAlt />
                <span>Evaluación</span>
            </NavLink>
            <NavLink to="/history" className="tab">
                <FaHistory />
                <span>Historial</span>
            </NavLink>
        </div>
    );
};
