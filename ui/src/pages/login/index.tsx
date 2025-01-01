import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

import { Button, Paper, TextBody, TextField, TextWelcome } from '#/pages/login/components.tsx'
import { sleep } from '#/utils'

export default function LoginPage() {
  const navigate = useNavigate()
  const [isLoading, setIsLoading] = useState(false)

  const handleLogin = async () => {
    setIsLoading(true)
    await sleep(1000)
    navigate('/')
  }

  return (
    <div className="flex justify-center items-center bg-[#111936] min-h-screen text-center">
      <Paper>
        <TextBody>IdeaGate</TextBody>
        <TextWelcome />
        <TextBody>Silahkan isi username dan password anda</TextBody>
        <TextField label="Username" />
        <TextField label="Password" type="password" />
        <Button isLoading={isLoading} onClick={handleLogin}>
          Masuk
        </Button>
      </Paper>
    </div>
  )
}
