import { Avatar, ButtonBase, Tooltip } from '@mui/material'
import { FC, ReactNode } from 'react'
import { ImDatabase } from 'react-icons/im'
import { MdHome, MdLogout, MdSettings } from 'react-icons/md'
import { Link } from 'react-router-dom'

export const DashboardLayout: FC<{ children: ReactNode }> = ({ children }) => {
  return (
    <>
      <AppBar />
      <div className="h-[calc(100vh-48px)] grid grid-cols-[min-content_auto]">
        <Drawer />
        {children}
      </div>
    </>
  )
}

const AppBar: FC = () => {
  return (
    <div className="flex items-center justify-between h-[48px] px-[12px] bg-tertiary z-20 sticky">
      <div className="flex items-center gap-3">
        <div className="font-bold text-[1rem] text-primary">IdeaGate</div>
        <ButtonBase>Workspace 1</ButtonBase>
      </div>

      <Avatar sx={{ height: '32px', width: '32px' }} src="https://mui.com/static/images/avatar/1.jpg" />
    </div>
  )
}

const Drawer: FC = () => {
  const items: Array<{ name: string; icon: JSX.Element; to: string; isBottom?: boolean }> = [
    {
      name: 'Home',
      icon: <MdHome className="text-primary" />,
      to: '/',
    },
    {
      name: 'Data Source',
      icon: <ImDatabase className="text-primary" />,
      to: '/datasource',
    },
    {
      name: 'Settings',
      icon: <MdSettings className="text-primary" />,
      to: '/settings',
    },
    {
      name: 'Logout',
      icon: <MdLogout className="text-primary" />,
      to: '/logout',
      isBottom: true,
    },
  ]

  return (
    <div className="w-[56px] bg-[#191B1F] flex flex-col justify-between">
      <div>
        {items
          .filter((item) => !item.isBottom)
          .map((item) => (
            <Tooltip title={item.name} placement="right">
              <Link to={item.to}>
                <ButtonBase className="w-full h-[48px] [&>svg]:w-[24px] [&>svg]:h-[24px]">{item.icon}</ButtonBase>
              </Link>
            </Tooltip>
          ))}
      </div>
      <div>
        {items
          .filter((item) => item.isBottom)
          .map((item) => (
            <Tooltip title={item.name} placement="right">
              <Link to={item.to}>
                <ButtonBase className="w-full h-[48px] [&>svg]:w-[24px] [&>svg]:h-[24px]">{item.icon}</ButtonBase>
              </Link>
            </Tooltip>
          ))}
      </div>
    </div>
  )
}
