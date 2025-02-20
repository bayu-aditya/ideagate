import { FC } from 'react'
import { useParams } from 'react-router-dom'

import { Workflow } from './components'

const ApplicationPage: FC = () => {
  const { app_id } = useParams()
  console.log(app_id)

  return <Workflow />
}

export default ApplicationPage
