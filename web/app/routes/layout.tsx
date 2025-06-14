import { Outlet } from 'react-router';

export default function Layout() {
  return (
    <div className='flex flex-col min-h-screen px-10'>
      <div className='flex-grow'>
        <Outlet />
      </div>
    </div>
  );
}
