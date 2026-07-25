import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { getAdminPost, type AdminPost } from "../lib/adminApi";
import { useDocumentTitle } from "../lib/useDocumentTitle";
import AdminPostForm from "./AdminPostForm";

export default function AdminPostEditorPage() {
  const { id } = useParams<{ id: string }>();
  const isNew = !id;
  const [post, setPost] = useState<AdminPost | null>(null);
  const [error, setError] = useState("");

  useDocumentTitle(isNew ? "New Post - Admin" : "Edit Post - Admin");

  useEffect(() => {
    if (!id) return;
    getAdminPost(id)
      .then(setPost)
      .catch((caught) => setError(caught instanceof Error ? caught.message : "Failed to load post"));
  }, [id]);

  if (!isNew && !post && !error) return null;

  return (
    <main className="mx-auto max-w-5xl p-6">
      <div className="mb-6 flex items-center justify-between gap-4">
        <div>
          <Link to="/admin/posts" className="text-sm font-semibold text-sage">← Admin posts</Link>
          <h1 className="mt-2 font-serif text-4xl font-bold text-charcoal">{isNew ? "New post" : "Edit post"}</h1>
        </div>
      </div>
      {error ? <p className="rounded-2xl bg-rose-50 px-4 py-3 text-sm text-rose-700">{error}</p> : <AdminPostForm post={post ?? undefined} />}
    </main>
  );
}
