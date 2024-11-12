import { createBrowserRouter } from 'react-router-dom'

import { DashboardLayout } from './layout/dashboard'
import DataSourcePage from './pages/datasource'
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
    path: '/datasource',
    element: (
      <DashboardLayout>
        <DataSourcePage />
      </DashboardLayout>
    ),
  },
])
