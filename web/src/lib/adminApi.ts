import type { Category } from "./api";

export interface AdminPost {
  id: string;
  slug: string;
  title: string;
  summary: string;
  category: Category;
  tags: string[];
  coverImage: string;
  status: "draft" | "published";
  contentJSON: unknown;
  contentHTML: string;
  createdAt: string;
  updatedAt: string;
  publishedAt?: string;
}

export interface AdminPostInput {
  slug: string;
  title: string;
  summary: string;
  category: Category;
  tags: string[];
  coverImage: string;
  contentJSON: unknown;
  contentHTML: string;
}

const API_BASE = "/api/admin";
const TOKEN_KEY = "adminToken";

export function getAdminToken(): string {
  return localStorage.getItem(TOKEN_KEY) ?? "";
}

export function setAdminToken(token: string) {
  localStorage.setItem(TOKEN_KEY, token);
}

async function adminJSON<T>(path: string, options: RequestInit = {}): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${getAdminToken()}`,
      ...options.headers,
    },
  });
  if (!res.ok) {
    const message = await res.text();
    throw new Error(message || `Request failed with status ${res.status}`);
  }
  if (res.status === 204) return undefined as T;
  return res.json() as Promise<T>;
}

export function listAdminPosts(): Promise<AdminPost[]> {
  return adminJSON<AdminPost[]>("/posts");
}

export function getAdminPost(id: string): Promise<AdminPost> {
  return adminJSON<AdminPost>(`/posts/${id}`);
}

export function createAdminPost(input: AdminPostInput): Promise<AdminPost> {
  return adminJSON<AdminPost>("/posts", { method: "POST", body: JSON.stringify(input) });
}

export function updateAdminPost(id: string, input: AdminPostInput): Promise<AdminPost> {
  return adminJSON<AdminPost>(`/posts/${id}`, { method: "PUT", body: JSON.stringify(input) });
}

export function publishAdminPost(id: string): Promise<AdminPost> {
  return adminJSON<AdminPost>(`/posts/${id}/publish`, { method: "POST" });
}

export function unpublishAdminPost(id: string): Promise<AdminPost> {
  return adminJSON<AdminPost>(`/posts/${id}/unpublish`, { method: "POST" });
}

export function deleteAdminPost(id: string): Promise<void> {
  return adminJSON<void>(`/posts/${id}`, { method: "DELETE" });
}
