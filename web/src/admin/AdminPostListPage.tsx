import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { deleteAdminPost, getAdminToken, listAdminPosts, setAdminToken, type AdminPost } from "../lib/adminApi";
import { useDocumentTitle } from "../lib/useDocumentTitle";

export default function AdminPostListPage() {
  useDocumentTitle("Admin Posts - My Blog");
  const [posts, setPosts] = useState<AdminPost[]>([]);
  const [token, setToken] = useState(getAdminToken());
  const [message, setMessage] = useState("");

  async function load() {
    try {
      setPosts(await listAdminPosts());
      setMessage("");
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Failed to load posts");
    }
  }

  useEffect(() => {
    void load();
  }, []);

  function saveToken() {
    setAdminToken(token);
    void load();
  }

  async function remove(id: string) {
    if (!confirm("Delete this post?")) return;
    await deleteAdminPost(id);
    await load();
  }

  return (
    <main className="mx-auto max-w-5xl p-6">
      <div className="mb-6 flex flex-wrap items-center justify-between gap-4">
        <div>
          <p className="text-sm font-semibold uppercase tracking-wide text-sage">CMS</p>
          <h1 className="font-serif text-4xl font-bold text-charcoal">Admin posts</h1>
        </div>
        <Link to="/admin/posts/new" className="rounded-2xl bg-sage px-5 py-2.5 text-sm font-semibold text-white">New post</Link>
      </div>

      <div className="mb-6 flex gap-2 rounded-3xl border border-white/60 bg-white/80 p-4 shadow-glass">
        <input
          value={token}
          onChange={(event) => setToken(event.target.value)}
          placeholder="Admin bearer token"
          className="min-w-0 flex-1 rounded-xl border border-gray-200 px-3 py-2"
          type="password"
        />
        <button onClick={saveToken} className="rounded-xl bg-navy-400 px-4 py-2 text-sm font-semibold text-white">Use token</button>
      </div>

      {message && <p className="mb-4 rounded-2xl bg-rose-50 px-4 py-3 text-sm text-rose-700">{message}</p>}

      <div className="overflow-hidden rounded-3xl border border-white/60 bg-white/80 shadow-glass">
        {posts.map((post) => (
          <div key={post.id} className="flex flex-wrap items-center justify-between gap-4 border-b border-gray-100 p-4 last:border-b-0">
            <div>
              <h2 className="font-serif text-xl font-bold text-charcoal">{post.title}</h2>
              <p className="text-sm text-charcoal-light/60">/{post.slug} · {post.category} · {post.status}</p>
            </div>
            <div className="flex gap-2">
              <Link to={`/admin/posts/${post.id}/edit`} className="rounded-xl bg-gray-100 px-4 py-2 text-sm font-semibold text-charcoal">Edit</Link>
              <button onClick={() => void remove(post.id)} className="rounded-xl bg-rose-50 px-4 py-2 text-sm font-semibold text-rose-700">Delete</button>
            </div>
          </div>
        ))}
        {posts.length === 0 && <p className="p-8 text-center text-charcoal-light/60">No posts yet.</p>}
      </div>
    </main>
  );
}
