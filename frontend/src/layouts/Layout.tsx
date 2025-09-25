import type {FC, ReactNode} from "react";
import "./layout.scss";
import {Tabs} from "@/layouts/Tabs.tsx";

interface Props {
    children: ReactNode
}

export const Layout: FC<Props> = ({children}) => {
    return (
        <>
            <Tabs/>
            <div className="info-container">
                <h1>Mortgage Subscription Engine</h1>
                <p>Automatically evaluates mortgage loan applications based on DTI, LTV, and credit score</p>
            </div>
            <div className="main-container">
                <div className="floating-card">
                    {children}
                </div>
            </div>
        </>
    )
}