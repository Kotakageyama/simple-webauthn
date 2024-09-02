"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { Key } from "lucide-react";
import {
	startRegistration,
	startAuthentication,
} from "@simplewebauthn/browser";
import {
	create,
	parseCreationOptionsFromJSON,
	parseRequestOptionsFromJSON,
	get,
	CredentialCreationOptionsJSON,
	CredentialRequestOptionsJSON,
} from "@github/webauthn-json/browser-ponyfill";
import apiClient from "@/api/client";

export function PasskeyLogin() {
	const [email, setEmail] = useState("");
	const [isRegistering, setIsRegistering] = useState(false);
	const [error, setError] = useState("");

	const handlePasskeyAction = async (isRegistration: boolean) => {
		setError("");
		try {
			if (isRegistration) {
				// 登録時のチャレンジを取得
				const { data, error, response } = await apiClient.POST(
					"/passkey/register-challenge",
					{
						params: {},
						body: { email },
						credentials: "include",
					}
				);
				if (error) {
					throw new Error("サーバーからのレスポンスが不正です");
				}
				if (!response.ok) {
					throw new Error("サーバーからのレスポンスが不正です");
				}

				const challengeResponseJSON: CredentialCreationOptionsJSON =
					JSON.parse(JSON.stringify(data));
				const credentialData = parseCreationOptionsFromJSON(
					challengeResponseJSON
				);

				// Passkeyの登録
				let credential: Credential | null;
				try {
					credential = await create(credentialData);
				} catch (e) {
					throw new Error("Passkeyの登録に失敗しました");
				}
				if (!credential) {
					throw new Error("Passkeyの登録に失敗しました");
				}

				// 登録結果をサーバーに送信
				const verificationResponse = await apiClient.POST(
					"/passkey/register",
					{
						params: {},
						credentials: "include",
						body: JSON.stringify(credential) as unknown as Record<
							string,
							never
						>,
					}
				);

				if (verificationResponse.response.ok) {
					alert("Passkeyの登録が完了しました");
				} else {
					throw new Error("登録に失敗しました");
				}
			} else {
				// ログイン認証時のチャレンジを取得
				const { data, error, response } = await apiClient.POST(
					"/passkey/login-challenge",
					{
						credentials: "include",
					}
				);
				if (error) {
					throw new Error("サーバーからのレスポンスが不正です");
				}
				if (!response.ok) {
					throw new Error("サーバーからのレスポンスが不正です");
				}

				const setCookie = response.headers.get("Set-Cookie") || "";
				const challengeResponseJSON: CredentialRequestOptionsJSON =
					JSON.parse(JSON.stringify(data));
				let credentialData = parseRequestOptionsFromJSON(
					challengeResponseJSON
				);

				// Passkeyの認証
				let credential: Credential;
				try {
					credential = await get(credentialData);
				} catch (e) {
					throw new Error("Passkeyの認証に失敗しました");
				}
				if (!credential) {
					throw new Error("Passkeyの認証に失敗しました");
				}

				// 認証結果をサーバーに送信
				const verificationResponse = await apiClient.POST(
					"/passkey/login",
					{
						body: JSON.stringify(credential) as unknown as Record<
							string,
							never
						>,
						credentials: "include",
					}
				);

				if (verificationResponse.response.ok) {
					alert("ログインに成功しました");
				} else {
					throw new Error("ログインに失敗しました");
				}
			}
		} catch (err) {
			setError(
				err instanceof Error
					? err.message
					: "予期せぬエラーが発生しました"
			);
		}
	};

	return (
		<Card className="w-[350px] mx-auto">
			<CardHeader>
				<CardTitle>パスキーログイン</CardTitle>
				<CardDescription>
					メールアドレスを入力してパスキーで認証してください
				</CardDescription>
			</CardHeader>
			<CardContent>
				<div className="space-y-4">
					<div className="space-y-2">
						<Label htmlFor="email">メールアドレス</Label>
						<Input
							id="email"
							type="email"
							placeholder="you@example.com"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
							required
						/>
					</div>
					<Button
						onClick={() => handlePasskeyAction(isRegistering)}
						className="w-full"
						disabled={!email}
					>
						<Key className="mr-2 h-4 w-4" />
						{isRegistering
							? "パスキーを登録"
							: "パスキーでログイン"}
					</Button>
					<Button
						variant="outline"
						className="w-full"
						onClick={() => setIsRegistering(!isRegistering)}
					>
						{isRegistering
							? "ログインに戻る"
							: "パスキーを登録する"}
					</Button>
					{error && <p className="text-sm text-red-500">{error}</p>}
				</div>
			</CardContent>
		</Card>
	);
}
