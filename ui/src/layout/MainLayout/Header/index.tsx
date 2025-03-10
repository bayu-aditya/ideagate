import Avatar from '@mui/material/Avatar'
import Box from '@mui/material/Box'
import ButtonBase from '@mui/material/ButtonBase'
import { useTheme } from '@mui/material/styles'
import { IconMenu2 } from '@tabler/icons-react'
import { FC } from 'react'

import LogoSection from '#/components/atoms/logo'

import ProfileSection from './ProfileSection'
import SearchSection from './SearchSection'

interface HeaderProps {
  handleLeftDrawerToggle: () => void
}

const Header: FC<HeaderProps> = ({ handleLeftDrawerToggle }) => {
  const theme = useTheme()

  return (
    <>
      {/* logo & toggler button */}
      <Box
        sx={{
          display: 'flex',
          gap: '16px',
          width: '200px',
          [theme.breakpoints.down('md')]: {
            width: 'auto',
          },
        }}
      >
        <ButtonBase sx={{ borderRadius: '8px', overflow: 'hidden' }}>
          <Avatar
            variant="rounded"
            sx={{
              ...theme.typography.commonAvatar,
              ...theme.typography.mediumAvatar,
              transition: 'all .2s ease-in-out',
              background: theme.palette.secondary.light,
              color: theme.palette.secondary.dark,
              '&:hover': {
                background: theme.palette.secondary.dark,
                color: theme.palette.secondary.light,
              },
            }}
            onClick={handleLeftDrawerToggle}
            color="inherit"
          >
            <IconMenu2 stroke={1.5} size="1.3rem" />
          </Avatar>
        </ButtonBase>
        <Box component="span" sx={{ display: { xs: 'none', md: 'block' }, flexGrow: 1 }}>
          <LogoSection />
        </Box>
      </Box>

      {/* header search */}
      <SearchSection />
      <Box sx={{ flexGrow: 1 }} />
      <Box sx={{ flexGrow: 1 }} />

      {/* profile */}
      <ProfileSection />
    </>
  )
}

export default Header
