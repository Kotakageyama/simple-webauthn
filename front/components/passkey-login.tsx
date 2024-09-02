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
	PublicKeyCredentialCreationOptionsJSON,
	PublicKeyCredentialRequestOptionsJSON,
} from "@simplewebauthn/types";
import apiClient from "@/api/client";
import { cookies } from "next/headers";

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
					}
				);
				if (error) {
					throw new Error("サーバーからのレスポンスが不正です");
				}
				if (!response.ok) {
					throw new Error("サーバーからのレスポンスが不正です");
				}

				const setCookie = response.headers.get("Set-Cookie") || "";
				const challengeData: PublicKeyCredentialCreationOptionsJSON =
					JSON.parse(JSON.stringify(data));

				// Passkeyの登録
				const regResult = await startRegistration(challengeData);

				// 登録結果をサーバーに送信
				const verificationResponse = await apiClient.POST(
					"/passkey/register",
					{
						params: {
							cookie: {
								__attestation__: setCookie,
							},
							body: JSON.stringify(regResult),
						},
					}
				);

				if (verificationResponse.response.ok) {
					alert("Passkeyの登録が完了しました");
				} else {
					throw new Error("登録に失敗しました");
				}
			} else {
				// ログイン認証時のチャレンジを取得
				const { response: challengeResponse } = await apiClient.POST(
					"/passkey/login-challenge"
				);
				const setCookie =
					challengeResponse.headers.get("Set-Cookie") || "";
				const challengeResponseText = await challengeResponse.text();
				const challengeData: PublicKeyCredentialRequestOptionsJSON =
					JSON.parse(challengeResponseText);

				// Passkeyでの認証
				const authResult = await startAuthentication(challengeData);

				// 認証結果をサーバーに送信
				const verificationResponse = await apiClient.POST(
					"/passkey/login",
					{
						params: {
							cookie: {
								__assertion__: setCookie,
							},
							body: JSON.stringify(authResult),
						},
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
