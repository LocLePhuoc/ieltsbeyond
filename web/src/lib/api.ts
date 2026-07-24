export type Category = "tech" | "gaming" | "travelling";

export interface Post {
  slug: string;
  title: string;
  summary: string;
  category: Category;
  tags: string[];
  date: string;
  coverImage: string;
  htmlContent?: string;
}

export interface CategoryInfo {
  slug: Category;
  description: string;
}

const API_BASE = "/api";

async function getJSON<T>(path: string): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`);
  if (!res.ok) {
    throw new Error(`Request to ${path} failed with status ${res.status}`);
  }
  return res.json() as Promise<T>;
}

export function getRecentPosts(limit: number): Promise<Post[]> {
  return getJSON<Post[]>(`/posts?limit=${limit}`);
}

export function getPostsByCategory(category: Category): Promise<Post[]> {
  return getJSON<Post[]>(`/posts?category=${category}`);
}

export function getPostBySlug(slug: string): Promise<Post> {
  return getJSON<Post>(`/posts/${encodeURIComponent(slug)}`);
}

export function getCategories(): Promise<CategoryInfo[]> {
  return getJSON<CategoryInfo[]>(`/categories`);
}
