import { useMemo, useState } from "react";
import type { FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import type { JSONContent } from "@tiptap/core";
import type { Category } from "../lib/api";
import {
  createAdminPost,
  publishAdminPost,
  unpublishAdminPost,
  updateAdminPost,
  type AdminPost,
  type AdminPostInput,
} from "../lib/adminApi";
import TiptapEditor from "./TiptapEditor";

const categories: Category[] = ["tech", "gaming", "travelling"];
const starterDoc: JSONContent = { type: "doc", content: [{ type: "paragraph" }] };

function toJSONContent(value: unknown): JSONContent {
  return value && typeof value === "object" ? (value as JSONContent) : starterDoc;
}

export default function AdminPostForm({ post }: { post?: AdminPost }) {
  const navigate = useNavigate();
  const [savedPost, setSavedPost] = useState<AdminPost | undefined>(post);
  const [slug, setSlug] = useState(post?.slug ?? "");
  const [title, setTitle] = useState(post?.title ?? "");
  const [summary, setSummary] = useState(post?.summary ?? "");
  const [category, setCategory] = useState<Category>(post?.category ?? "tech");
  const [tagsText, setTagsText] = useState(post?.tags.join(", ") ?? "");
  const [coverImage, setCoverImage] = useState(post?.coverImage ?? "/static/images/placeholder.svg");
  const [contentJSON, setContentJSON] = useState<JSONContent>(toJSONContent(post?.contentJSON));
  const [contentHTML, setContentHTML] = useState(post?.contentHTML ?? "<p></p>");
  const [message, setMessage] = useState("");
  const [saving, setSaving] = useState(false);

  const input: AdminPostInput = useMemo(
    () => ({
      slug,
      title,
      summary,
      category,
      tags: tagsText
        .split(",")
        .map((tag) => tag.trim())
        .filter(Boolean),
      coverImage,
      contentJSON,
      contentHTML,
    }),
    [slug, title, summary, category, tagsText, coverImage, contentJSON, contentHTML],
  );

  async function save(event?: FormEvent) {
    event?.preventDefault();
    setSaving(true);
    setMessage("");
    try {
      const result = savedPost ? await updateAdminPost(savedPost.id, input) : await createAdminPost(input);
      setSavedPost(result);
      setMessage("Draft saved.");
      if (!savedPost) navigate(`/admin/posts/${result.id}/edit`, { replace: true });
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Save failed");
    } finally {
      setSaving(false);
    }
  }

  async function publish() {
    const current = savedPost ?? (await createAdminPost(input));
    const result = await publishAdminPost(current.id);
    setSavedPost(result);
    setMessage("Post published.");
  }

  async function unpublish() {
    if (!savedPost) return;
    const result = await unpublishAdminPost(savedPost.id);
    setSavedPost(result);
    setMessage("Post unpublished.");
  }

  return (
    <form onSubmit={save} className="space-y-5 rounded-3xl border border-white/60 bg-white/80 p-6 shadow-glass">
      {message && <p className="rounded-2xl bg-sage-50 px-4 py-3 text-sm text-charcoal-light">{message}</p>}

      <div className="grid gap-4 md:grid-cols-2">
        <label className="space-y-1 text-sm font-semibold text-charcoal">
          Slug
          <input value={slug} onChange={(event) => setSlug(event.target.value)} className="w-full rounded-xl border border-gray-200 px-3 py-2 font-normal" />
        </label>
        <label className="space-y-1 text-sm font-semibold text-charcoal">
          Category
          <select value={category} onChange={(event) => setCategory(event.target.value as Category)} className="w-full rounded-xl border border-gray-200 px-3 py-2 font-normal">
            {categories.map((item) => (
              <option key={item} value={item}>{item}</option>
            ))}
          </select>
        </label>
      </div>

      <label className="block space-y-1 text-sm font-semibold text-charcoal">
        Title
        <input value={title} onChange={(event) => setTitle(event.target.value)} className="w-full rounded-xl border border-gray-200 px-3 py-2 font-normal" />
      </label>

      <label className="block space-y-1 text-sm font-semibold text-charcoal">
        Summary
        <textarea value={summary} onChange={(event) => setSummary(event.target.value)} className="min-h-24 w-full rounded-xl border border-gray-200 px-3 py-2 font-normal" />
      </label>

      <div className="grid gap-4 md:grid-cols-2">
        <label className="space-y-1 text-sm font-semibold text-charcoal">
          Tags, comma-separated
          <input value={tagsText} onChange={(event) => setTagsText(event.target.value)} className="w-full rounded-xl border border-gray-200 px-3 py-2 font-normal" />
        </label>
        <label className="space-y-1 text-sm font-semibold text-charcoal">
          Cover image URL/path
          <input value={coverImage} onChange={(event) => setCoverImage(event.target.value)} className="w-full rounded-xl border border-gray-200 px-3 py-2 font-normal" />
        </label>
      </div>

      <TiptapEditor
        content={contentJSON}
        onChange={(json, html) => {
          setContentJSON(json);
          setContentHTML(html);
        }}
      />

      <div className="flex flex-wrap gap-3">
        <button disabled={saving} className="rounded-2xl bg-navy-400 px-5 py-2.5 text-sm font-semibold text-white disabled:opacity-50" type="submit">
          {saving ? "Saving..." : "Save draft"}
        </button>
        <button type="button" onClick={publish} className="rounded-2xl bg-sage px-5 py-2.5 text-sm font-semibold text-white">
          Publish
        </button>
        {savedPost?.status === "published" && (
          <button type="button" onClick={unpublish} className="rounded-2xl bg-gray-100 px-5 py-2.5 text-sm font-semibold text-charcoal">
            Unpublish
          </button>
        )}
        {savedPost && <span className="self-center text-sm text-charcoal-light/70">Status: {savedPost.status}</span>}
      </div>
    </form>
  );
}
