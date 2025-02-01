import { Grid2 } from '@mui/material'
import { FC } from 'react'

import { StepDetail, Workflow } from './components'

const ApplicationPage: FC = () => {
  return (
    <Grid2 container spacing={2}>
      <Grid2 size="grow">
        <Workflow />
      </Grid2>
      <Grid2 size={3}>
        <StepDetail />
      </Grid2>
    </Grid2>
  )
}

export default ApplicationPage
