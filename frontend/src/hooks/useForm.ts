import {type ChangeEvent, useState} from 'react';

type FormFields<T> = {
    [key in keyof T]: T[key];
};

export const useForm = <T>(initialForm: FormFields<T>) => {
    const [formState, setFormState] = useState<FormFields<T>>(initialForm);

    const onInputChange = (event: ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = event.target;
        setFormState({
            ...formState,
            [name]: value,
        });
    };

    const onResetForm = () => {
        setFormState(initialForm);
    };

    return {
        ...formState,
        formState,
        onInputChange,
        onResetForm,
    };
};