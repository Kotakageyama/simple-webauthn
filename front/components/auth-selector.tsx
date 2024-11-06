import { useState } from 'react'
import { Button } from './ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card'
import { PasskeyLogin } from './passkey-login'
import { WorldIDLogin } from './world-id-login'

type AuthMethod = 'select' | 'passkey' | 'worldid'

export function AuthSelector() {
  const [authMethod, setAuthMethod] = useState<AuthMethod>('select')

  if (authMethod === 'passkey') {
    return (
      <div className="w-full max-w-md">
        <Button
          variant="outline"
          className="mb-4"
          onClick={() => setAuthMethod('select')}
        >
          ← Back to selection
        </Button>
        <PasskeyLogin />
      </div>
    )
  }

  if (authMethod === 'worldid') {
    return (
      <div className="w-full max-w-md">
        <Button
          variant="outline"
          className="mb-4"
          onClick={() => setAuthMethod('select')}
        >
          ← Back to selection
        </Button>
        <WorldIDLogin />
      </div>
    )
  }

  return (
    <Card className="w-full max-w-md">
      <CardHeader>
        <CardTitle>Choose Authentication Method</CardTitle>
        <CardDescription>
          Select how you would like to authenticate
        </CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        <Button
          className="w-full"
          onClick={() => setAuthMethod('passkey')}
        >
          Continue with Passkey
        </Button>
        <Button
          className="w-full"
          onClick={() => setAuthMethod('worldid')}
        >
          Continue with World ID
        </Button>
      </CardContent>
    </Card>
  )
}
