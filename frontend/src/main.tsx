import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './index.css'
import App from './App.tsx'
import {EvaluationPage} from "@/pages/EvaluationPage.tsx";
import {HistoryPage} from "@/pages/HistoryPage.tsx";

const router = createBrowserRouter([
    {
        path: '/',
        element: <App />, // App será nuestro "cascarón" o layout principal
        children: [
            {
                index: true, // Esta es la ruta por defecto (/)
                element: <EvaluationPage />,
            },
            {
              path: 'history',
              element: <HistoryPage />,
            },
        ],
    },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
