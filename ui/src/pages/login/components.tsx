import {
  Button as ButtonMui,
  ButtonProps as ButtonPropsMui,
  Paper as PaperMui,
  TextField as TextFieldMui,
  TextFieldProps,
  Typography,
} from '@mui/material'
import { FC, ReactNode } from 'react'

export const Paper: FC<{ children: ReactNode }> = ({ children }) => (
  <PaperMui
    sx={{
      backgroundColor: 'rgb(33,41,70)',
      width: '450px',
      padding: '24px',
      display: 'flex',
      flexDirection: 'column',
      gap: '16px',
    }}
  >
    {children}
  </PaperMui>
)

export const TextWelcome = () => <Typography sx={{ fontWeight: 'bold', fontSize: 20 }}>Selamat Datang</Typography>

export const TextBody: FC<{ children: ReactNode }> = ({ children }) => (
  <Typography sx={{ color: '#8492d4' }}>{children}</Typography>
)

export const TextField: FC<TextFieldProps> = (props) => <TextFieldMui {...props} fullWidth />

interface ButtonProps extends ButtonPropsMui {
  children: ReactNode
  isLoading: boolean
}

export const Button: FC<ButtonProps> = ({ children, isLoading, ...rest }) => (
  <ButtonMui {...rest} fullWidth variant="contained" disableElevation disabled={isLoading}>
    {children}
  </ButtonMui>
)
