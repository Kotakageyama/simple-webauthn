// 型宣言をインポート
import createClient from "openapi-fetch";
import type { paths } from "./api";

const apiClient = createClient<paths>({
    baseUrl: process.env.BACKEND_URL || "http://localhost:8080",
});

export default apiClient;
