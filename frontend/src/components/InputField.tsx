import type { ChangeEvent } from 'react';

interface InputFieldProps {
    label: string;
    name: string;
    value: string | number | undefined;
    onChange: (event: ChangeEvent<HTMLInputElement>) => void;
    type?: 'text' | 'number';
    placeholder?: string;
    className?: string;
}

export const InputField = ({
                               label,
                               name,
                               value,
                               onChange,
                               type = 'text',
                               placeholder,
                               className = ''
                           }: InputFieldProps) => {
    return (
        <div className={`input-group ${className}`}>
            <label htmlFor={name}>{label}</label>
            <input
                id={name}
                name={name}
                type={type}
                placeholder={placeholder}
                value={value}
                onChange={onChange}
            />
        </div>
    );
};