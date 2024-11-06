# simple-webauthn

## 説明 / Description

この Repository は Passkeys と World ID を利用した認証システムの学習用に作成されました。
This repository was created to learn about authentication using Passkeys and World ID.

## 認証方式 / Authentication Methods

### Passkeys
- WebAuthn/FIDO2 準拠の認証方式
- パスワードレス認証を実現
- Standard WebAuthn/FIDO2 compliant authentication
- Enables passwordless authentication

### World ID
- World ID による生体認証
- プライバシーを保護しながら人間であることを証明
- Biometric authentication using World ID
- Proves humanity while preserving privacy

## FEのコンポーネントは[v0](https://v0.dev/chat)を利用しました / Frontend components use [v0](https://v0.dev/chat)
![image](https://github.com/user-attachments/assets/30f4e5d0-21b5-4b1e-80cd-471007fa8611)

## セットアップ / Setup

### 1. World ID の設定 / World ID Configuration

```bash
# .env ファイルを作成 / Create .env file
NEXT_PUBLIC_WORLD_ID_APP_ID=your_app_id
NEXT_PUBLIC_WORLD_ID_ACTION=auth
NEXT_PUBLIC_WORLD_ID_SIGNAL=login
```

### 2. tool install

```bash
$ make install
```

### 3. build container

```bash
$ make build
```

### 4. start container

```bash
$ make up
```

## 使用方法 / Usage

### アクセス / Access Points
- フロントエンド / Frontend: http://localhost:3000
- バックエンド / Backend: http://localhost:8080

### 認証方法の選択 / Choosing Authentication Method
1. Passkeys
   - デバイスに登録された生体認証やPINを使用
   - Use device's biometric authentication or PIN
2. World ID
   - World Appでの生体認証を使用
   - Use World App for biometric verification

## 技術スタック / Tech Stack
- Frontend: Next.js, TypeScript
- Backend: Go
- Authentication:
  - WebAuthn/Passkeys
  - World ID (@worldcoin/idkit)
