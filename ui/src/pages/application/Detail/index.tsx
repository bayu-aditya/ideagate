import ApiIcon from '@mui/icons-material/Api'
import SettingsIcon from '@mui/icons-material/Settings'
import { Card, CardContent, CardHeader, Divider, Tab, TabProps, Tabs } from '@mui/material'
import { FC, useState } from 'react'
import { useParams } from 'react-router-dom'

import { Entrypoints, Setting } from './components'

const ApplicationPage: FC = () => {
  const { app_id } = useParams()

  const [tab, setTab] = useState(0)

  const Body = () => {
    switch (tab) {
      case 0:
        return <Entrypoints />
      case 1:
        return <Setting />
      default:
        return null
    }
  }

  return (
    <Card sx={{ minHeight: '100%' }}>
      <CardHeader title={app_id} />
      <CardContent sx={{ py: 0 }}>
        <Tabs value={tab} onChange={(_, value) => setTab(value)} sx={{ minHeight: 0 }}>
          <TabCustom icon={<ApiIcon />} iconPosition="start" label="Entrypoints" />
          <TabCustom icon={<SettingsIcon />} iconPosition="start" label="Settings" />
        </Tabs>
        <Divider sx={{ marginBottom: 2 }} />

        {Body()}
      </CardContent>
    </Card>
  )
}

const TabCustom = (props: TabProps) => <Tab {...props} sx={{ ...props.sx, minHeight: 0 }} />

export default ApplicationPage
