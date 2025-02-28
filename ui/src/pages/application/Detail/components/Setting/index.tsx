import { Divider, Grid2, Paper, TextField, Typography, useTheme } from '@mui/material'
import { FC } from 'react'

const SettingPage: FC = () => {
  const theme = useTheme()

  return (
    <Grid2 container justifyContent={'center'}>
      <Grid2 size={8}>
        <Paper
          sx={{
            my: 1,
            border: 'solid 1px',
            borderColor: 'divider',
            borderRadius: theme.opts.borderRadius,
            maxWidth: '800px',
          }}
        >
          <Typography variant="h5" sx={{ p: 3 }}>
            General Information
          </Typography>
          <Divider />
          <Grid2 container sx={{ p: 3 }} spacing={2}>
            <TextField label="Name" fullWidth />
            <TextField label="Description" fullWidth multiline minRows={3} />
          </Grid2>
        </Paper>
      </Grid2>
    </Grid2>
  )
}

export default SettingPage
