import { Endpoint } from '@bayu-aditya/ideagate-model-js/core/endpoint/endpoint'
import { Box, Typography, useTheme } from '@mui/material'
import { FC } from 'react'

const EnptypointPage: FC = () => {
  const theme = useTheme()
  const entrypoints = mockEntrypoints

  return (
    <div>
      {entrypoints.map((entrypoint) => (
        <Box
          key={entrypoint.id}
          sx={{ my: 1, border: 'solid 1px', borderColor: 'divider', borderRadius: theme.opts.borderRadius, padding: 1 }}
        >
          <Typography>
            {entrypoint.method} {entrypoint.path}
          </Typography>
          <Typography variant="subtitle2">{entrypoint.setting?.name}</Typography>
        </Box>
      ))}
    </div>
  )
}

export default EnptypointPage

const mockEntrypoints: Endpoint[] = [
  {
    id: '1',
    setting: {
      name: 'Healthcheck',
      description: 'Description of Healthcheck',
      timeoutMs: 0n,
      numWorkers: 0,
    },
    projectId: 'test',
    method: 'GET',
    path: '/health',
  },
  {
    id: '2',
    setting: {
      name: 'Get User',
      description: 'Description of Get User',
      timeoutMs: 0n,
      numWorkers: 0,
    },
    projectId: 'test',
    method: 'GET',
    path: '/user',
  },
]
