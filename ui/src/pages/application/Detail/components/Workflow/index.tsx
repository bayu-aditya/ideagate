import { Grid2 } from '@mui/material'
import { FC } from 'react'

import Pipeline from './components/Pipeline'
import StepDetail from './components/StepDetail'

const WorkflowTab: FC = () => {
  return (
    <Grid2 container spacing={2}>
      <Grid2 size="grow">
        <Pipeline />
      </Grid2>
      <Grid2 size={3}>
        <StepDetail />
      </Grid2>
    </Grid2>
  )
}

export default WorkflowTab
