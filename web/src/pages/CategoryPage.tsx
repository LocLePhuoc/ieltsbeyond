import { useEffect, useState } from "react";
import { getPostsByCategory, getCategories, type Post, type Category } from "../lib/api";
import { BentoCard } from "../components/BentoCard";
import { capitalize } from "../lib/cards";
import { useDocumentTitle } from "../lib/useDocumentTitle";

export default function CategoryPage({ category }: { category: Category }) {
  const [posts, setPosts] = useState<Post[] | null>(null);
  const [description, setDescription] = useState("");

  useDocumentTitle(`${capitalize(category)} - My Blog`);

  useEffect(() => {
    let cancelled = false;
    setPosts(null);
    Promise.all([getPostsByCategory(category), getCategories()]).then(([postsData, categories]) => {
      if (cancelled) return;
      setPosts(postsData);
      setDescription(categories.find((c) => c.slug === category)?.description ?? "");
    });
    return () => {
      cancelled = true;
    };
  }, [category]);

  if (!posts) return null;

  return (
    <div>
      <section className="mb-8">
        <h1 className="font-serif text-4xl md:text-5xl font-bold text-charcoal mb-2 tracking-tight">{capitalize(category)}</h1>
        <p className="text-charcoal-light text-[15px]">{description}</p>
      </section>
      {posts.length > 0 ? (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3.5 auto-rows-auto">
          {posts.map((post, i) => (
            <BentoCard key={post.slug} post={post} index={i} large={i === 0 && posts.length > 2} />
          ))}
        </div>
      ) : (
        <div className="flex flex-col items-center justify-center py-20 bg-white/60 backdrop-blur-sm rounded-3xl border border-white/60">
          <div className="w-12 h-12 rounded-2xl bg-gray-100 flex items-center justify-center mb-4">
            <svg className="w-6 h-6 text-charcoal-light/40" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={1.5}>
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z"
              />
            </svg>
          </div>
          <p className="text-charcoal-light/50 text-sm">No posts yet in this category.</p>
        </div>
      )}
    </div>
  );
}
