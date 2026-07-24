import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { getPostBySlug, type Post } from "../lib/api";
import { capitalize, formatDate } from "../lib/cards";
import { useActiveTab } from "../components/ActiveTabContext";
import { useDocumentTitle } from "../lib/useDocumentTitle";
import NotFoundPage from "./NotFoundPage";

export default function PostDetailPage() {
  const { slug } = useParams<{ slug: string }>();
  const { setActiveTab } = useActiveTab();
  const [post, setPost] = useState<Post | null>(null);
  const [notFound, setNotFound] = useState(false);

  useEffect(() => {
    let cancelled = false;
    setPost(null);
    setNotFound(false);
    getPostBySlug(slug!)
      .then((data) => {
        if (cancelled) return;
        setPost(data);
        setActiveTab(data.category);
      })
      .catch(() => {
        if (!cancelled) setNotFound(true);
      });
    return () => {
      cancelled = true;
    };
  }, [slug, setActiveTab]);

  useDocumentTitle(post ? `${post.title} - My Blog` : "My Blog");

  if (notFound) return <NotFoundPage />;
  if (!post) return null;

  const categoryUrl = `/${post.category}`;

  return (
    <article className="max-w-[720px]">
      <div className="mb-6">
        <Link
          to={categoryUrl}
          className="inline-flex items-center gap-1.5 text-[13px] font-medium text-charcoal-light/60 hover:text-charcoal transition-colors duration-200 group"
        >
          <svg className="w-4 h-4 transition-transform duration-200 group-hover:-translate-x-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
            <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
          </svg>
          {capitalize(post.category)}
        </Link>
      </div>

      <header className="mb-8">
        <span className="inline-flex items-center px-3 py-1 rounded-full text-[11px] font-semibold bg-sage-50 text-sage tracking-wide uppercase mb-4">
          {capitalize(post.category)}
        </span>
        <h1 className="font-serif text-3xl md:text-[40px] font-bold text-charcoal leading-[1.15] tracking-tight mb-4">{post.title}</h1>
        <div className="flex items-center gap-3">
          <time className="text-[13px] font-medium text-charcoal-light/50 tracking-wide">{formatDate(post.date)}</time>
        </div>
      </header>

      <div className="bg-white/70 backdrop-blur-sm rounded-3xl border border-white/60 shadow-glass p-6 md:p-10">
        <div className="prose-blog" dangerouslySetInnerHTML={{ __html: post.htmlContent ?? "" }} />
      </div>

      <div className="mt-8 flex">
        <Link
          to={categoryUrl}
          className="inline-flex items-center gap-2 px-5 py-2.5 bg-white/80 backdrop-blur-sm border border-gray-200/60 text-charcoal rounded-2xl text-[13px] font-medium hover:bg-white hover:shadow-soft transition-all duration-200"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
            <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
          </svg>
          More in {capitalize(post.category)}
        </Link>
      </div>
    </article>
  );
}
