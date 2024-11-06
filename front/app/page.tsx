import { WorldIDLogin } from "@/components/world-id-login";
import { PasskeyLogin } from "@/components/passkey-login";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useState } from "react";

type AuthMethod = 'select' | 'passkey' | 'worldid';

export default function Home() {
	const [authMethod, setAuthMethod] = useState<AuthMethod>('select');

	const renderAuthMethod = () => {
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
			);
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
			);
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
		);
	};

	return (
		<main className="flex min-h-screen flex-col items-center justify-between p-24">
			{renderAuthMethod()}
		</main>
	);
}
