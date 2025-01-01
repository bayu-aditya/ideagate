import { createBrowserRouter } from 'react-router-dom'

import { DashboardLayout } from './layout/dashboard'
import DataSourcePage from './pages/datasource'
import LoginPage from './pages/login'
import Root from './pages/root'

export const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <DashboardLayout>
        <Root />
      </DashboardLayout>
    ),
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
  {
    path: '/datasource',
    element: (
      <DashboardLayout>
        <DataSourcePage />
      </DashboardLayout>
    ),
  },
])
