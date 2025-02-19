import AddIcon from '@mui/icons-material/Add'
import { Card, CardContent, CardHeader, Grid2, IconButton, Paper, Typography, useTheme } from '@mui/material'
import { useQuery } from '@tanstack/react-query'
import moment from 'moment'
import { FC } from 'react'
import { Link, useParams } from 'react-router-dom'

import { getListApps } from './api/grpc-web'

const ApplicationListPage: FC = () => {
  const theme = useTheme()

  const { project_id } = useParams()
  const projectID = project_id as string

  const { data } = useQuery({ queryKey: ['apps', project_id], queryFn: () => getListApps(projectID) })
  const apps = data?.applications || []

  return (
    <Card sx={{ minHeight: '100%' }}>
      <CardHeader
        title="Application List"
        action={
          <IconButton sx={{ backgroundColor: theme.palette.secondary.main, color: 'white' }}>
            <AddIcon />
          </IconButton>
        }
      />
      <CardContent>
        <Grid2 container spacing={2}>
          {apps.map((app) => (
            <Grid2 size={4}>
              <Link to={`/${projectID}/application/${app.id}`}>
                <Paper
                  key={app.id}
                  variant="outlined"
                  sx={{
                    backgroundColor: theme.opts.colors.grey50,
                    padding: theme.spacing(2),
                    cursor: 'pointer',
                  }}
                >
                  <Typography variant="h3" sx={{ color: theme.palette.secondary.main }}>
                    {app.name}
                  </Typography>
                  <Typography variant="subtitle2" sx={{ marginBottom: theme.spacing(1) }}>
                    {app.description}
                  </Typography>
                  <Typography variant="subtitle2">
                    Created at {moment.unix(Number(app.createdAt?.seconds)).startOf('hours').fromNow()}
                  </Typography>
                  <Typography variant="subtitle2">
                    Updated at {moment.unix(Number(app.updatedAt?.seconds)).startOf('hours').fromNow()}
                  </Typography>
                </Paper>
              </Link>
            </Grid2>
          ))}
        </Grid2>
      </CardContent>
    </Card>
  )
}

export default ApplicationListPage
