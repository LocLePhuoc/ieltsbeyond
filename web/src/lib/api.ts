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
  contentJSON?: unknown;
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

export interface WritingTask1 {
  id: string;
  question: string;
  type: string;
  image_key: string;
}

export interface WritingTask2 {
  id: string;
  question: string;
  category: string;
}

export function getWritingTask1List(limit = 20): Promise<WritingTask1[]> {
  return getJSON<WritingTask1[]>(`/writing/task1?limit=${limit}`);
}

export function getWritingTask2List(limit = 20): Promise<WritingTask2[]> {
  return getJSON<WritingTask2[]>(`/writing/task2?limit=${limit}`);
}
