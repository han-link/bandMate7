import { AppProvider } from './provider';
import { AppRouter } from './router.tsx';

export const App = () => {
    return (
        <AppProvider>
            <AppRouter />
        </AppProvider>
    );
};