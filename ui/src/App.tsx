import './App.css'

import { CssBaseline, ThemeProvider } from '@mui/material'
import { RouterProvider } from 'react-router-dom'

import theme from '#/themes'

import { router } from './router'

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  )
}

export default App
