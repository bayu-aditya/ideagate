import { createBrowserRouter } from 'react-router-dom'

import ApplicationPage from '#/pages/application'

// import { DashboardLayout } from './layout/dashboard'
import MainLayout from './layout/MainLayout'
import DataSourcePage from './pages/datasource'
import LoginPage from './pages/login'
import Root from './pages/root'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <MainLayout />,
    children: [
      {
        path: '/',
        element: <Root />,
      },
      {
        path: '/application',
        element: <ApplicationPage />,
      },
      {
        path: '/datasource',
        element: <DataSourcePage />,
      },
    ],
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
])
