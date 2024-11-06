import { IDKitWidget } from '@worldcoin/idkit'
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { Button } from './ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card'

export function WorldIDLogin() {
  const router = useRouter()
  const [error, setError] = useState<string>('')

  const onSuccess = async (response: any) => {
    try {
      const res = await fetch('/api/worldid/verify', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          proof: response.proof,
          action: process.env.NEXT_PUBLIC_WORLD_ID_ACTION || 'auth',
          signal: process.env.NEXT_PUBLIC_WORLD_ID_SIGNAL || 'login',
        }),
      })

      if (!res.ok) {
        throw new Error('Verification failed')
      }

      // Redirect to dashboard on success
      router.push('/dashboard')
    } catch (err) {
      setError('Authentication failed. Please try again.')
    }
  }

  return (
    <Card className="w-[350px] mx-auto">
      <CardHeader>
        <CardTitle>World ID Login</CardTitle>
        <CardDescription>
          Verify with World ID to authenticate
        </CardDescription>
      </CardHeader>
      <CardContent>
        {error && (
          <div className="text-red-500 mb-4 text-sm">
            {error}
          </div>
        )}
        <IDKitWidget
          app_id={process.env.NEXT_PUBLIC_WORLD_ID_APP_ID || 'app_ff9b0b40f0aaf55e2b6a6a967de21ede'}
          action="auth"
          signal="login"
          onSuccess={onSuccess}
          handleVerify={async () => true}
        >
          {({ open }) => <Button onClick={open}>Verify with World ID</Button>}
        </IDKitWidget>
      </CardContent>
    </Card>
  )
}
