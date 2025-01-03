import { Typography } from '@mui/material'
import { FC } from 'react'

import { menu } from '#/store/menu'

import NavGroup from './NavGroup'

const MenuList: FC = () => {
  const navItems = menu.map((item) => {
    switch (item.type) {
      case 'group':
        return <NavGroup key={item.id} item={item} />
      default:
        return (
          <Typography key={item.id} variant="h6" color="error" align="center">
            Menu Items Error
          </Typography>
        )
    }
  })

  return <>{navItems}</>
}

export default MenuList
