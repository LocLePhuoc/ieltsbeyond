import { useEffect, useState } from "react";
import { getRecentPosts, type Post } from "../lib/api";
import { BentoCard } from "../components/BentoCard";
import { useDocumentTitle } from "../lib/useDocumentTitle";

export default function HomePage() {
  useDocumentTitle("My Blog");
  const [posts, setPosts] = useState<Post[] | null>(null);

  useEffect(() => {
    let cancelled = false;
    getRecentPosts(6).then((data) => {
      if (!cancelled) setPosts(data);
    });
    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <div>
      <section className="mb-8">
        <div className="relative">
          <h1 className="font-serif text-4xl md:text-5xl font-bold text-charcoal leading-[1.1] tracking-tight mb-3">
            Discover
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-navy-400 to-sage"> stories</span>
          </h1>
          <p className="text-charcoal-light text-[15px] leading-relaxed max-w-lg">
            Exploring the intersections of technology, gaming, and the places we wander.
          </p>
        </div>
      </section>

      {posts && posts.length > 0 && (
        <section>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3.5 auto-rows-auto">
            {posts.map((post, i) => (
              <BentoCard key={post.slug} post={post} index={i} large={i === 0} />
            ))}
          </div>
        </section>
      )}
    </div>
  );
}
