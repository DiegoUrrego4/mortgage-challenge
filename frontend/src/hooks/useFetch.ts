import { useEffect, useState } from 'react';

interface FetchState<T> {
    data: T | null;
    isLoading: boolean;
    hasError: string | null;
}

export const useFetch = <T>(url: string) => {
    const [state, setState] = useState<FetchState<T>>({
        data: null,
        isLoading: true,
        hasError: null,
    });

    const getFetch = async () => {
        setState({ ...state, isLoading: true });
        try {
            const resp = await fetch(url, {
                headers: {
                    'Cache-Control': 'no-cache'
                }
            });

            const data = await resp.json();
            setState({
                data,
                isLoading: false,
                hasError: null,
            });
        } catch (error) {
            setState({
                data: null,
                isLoading: false,
                hasError: `${error}`
            })
        }
    };

    useEffect(() => {
        getFetch();
    }, [url]);

    return {
        data: state.data,
        isLoading: state.isLoading,
        hasError: state.hasError,
    };
};