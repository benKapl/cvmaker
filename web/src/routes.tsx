import { createBrowserRouter } from 'react-router-dom';
import Home from './routes/home';
import Register from './routes/register';

export const router = createBrowserRouter([
  { path: '/', Component: Home },
  { path: '/register', Component: Register },
]);
