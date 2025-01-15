import { IconApps, IconDashboard, IconKey, IconProps } from '@tabler/icons-react'
import { FC, ReactElement } from 'react'

export type Menu = {
  id: string
  title: string
  type: 'group' | 'collapse' | 'item'
  disabled?: boolean
  caption?: string
  url?: string
  icon?: FC<IconProps>
  chip?: {
    label: string
    color?: 'default' | 'primary' | 'secondary' | 'error' | 'info' | 'success' | 'warning'
    variant?: 'filled' | 'outlined'
    size?: 'small' | 'medium'
    avatar?: ReactElement
  }
  children?: Menu[]
}

export const menu: Menu[] = [
  {
    id: 'dashboard',
    title: 'Dashboard',
    type: 'group',
    children: [
      {
        id: 'default',
        title: 'Dashboard',
        type: 'item',
        url: '/',
        icon: IconDashboard,
      },
      {
        id: 'application',
        title: 'Aplikasi',
        type: 'item',
        url: '/application',
        icon: IconApps,
      },
    ],
  },
  {
    id: 'pages',
    title: 'Pages',
    caption: 'Pages Caption',
    type: 'group',
    children: [
      {
        id: 'authentication',
        title: 'Authentication',
        type: 'collapse',
        icon: IconKey,
        children: [
          {
            id: 'login',
            title: 'Login',
            type: 'item',
            url: '/login',
          },
          {
            id: 'register',
            title: 'Register',
            type: 'item',
            url: '/register',
          },
        ],
      },
    ],
  },
]
