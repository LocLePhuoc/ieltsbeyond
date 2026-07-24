import { Link } from "react-router-dom";
import type { Post } from "../lib/api";
import {
  capitalize,
  cardBadgeClass,
  cardGradientClass,
  cardSubtextClass,
  cardTextClass,
  formatDate,
} from "../lib/cards";

export function BentoCard({ post, index, large }: { post: Post; index: number; large: boolean }) {
  return (
    <Link
      to={`/posts/${post.slug}`}
      className={`bento-card group block rounded-3xl overflow-hidden transition-all duration-500 hover:shadow-card-hover relative ${cardGradientClass(index)} ${
        large ? "sm:col-span-2 sm:row-span-2" : ""
      }`}
    >
      {/* Subtle inner overlay for depth */}
      <div className="absolute inset-0 bg-gradient-to-t from-black/20 via-transparent to-white/5 pointer-events-none" />
      {/* Decorative circles */}
      <div className="absolute -top-12 -right-12 w-48 h-48 rounded-full bg-white/5 pointer-events-none" />
      <div className="absolute -bottom-8 -left-8 w-32 h-32 rounded-full bg-black/5 pointer-events-none" />

      <div className={`flex flex-col justify-between h-full p-5 relative z-10 ${large ? "sm:p-8 min-h-[340px]" : "min-h-[240px]"}`}>
        {/* Top: category badge + arrow */}
        <div className="flex items-start justify-between">
          <span className={`inline-flex items-center px-3 py-1 rounded-full text-[11px] font-semibold tracking-wide uppercase ${cardBadgeClass(index)}`}>
            {capitalize(post.category)}
          </span>
          <div className={`w-8 h-8 rounded-full bg-white/10 flex items-center justify-center opacity-0 group-hover:opacity-100 -translate-x-2 group-hover:translate-x-0 transition-all duration-300 ${cardTextClass(index)}`}>
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25" />
            </svg>
          </div>
        </div>
        {/* Bottom: title, summary, meta */}
        <div className="mt-auto">
          <h3 className={`font-serif font-bold leading-tight tracking-tight mb-2 ${cardTextClass(index)} ${large ? "text-2xl sm:text-[28px]" : "text-[17px]"}`}>
            {post.title}
          </h3>
          {large && (
            <p className={`text-sm leading-relaxed mb-3 line-clamp-2 ${cardSubtextClass(index)}`}>{post.summary}</p>
          )}
          <div className={`text-[11px] font-medium tracking-wide ${cardSubtextClass(index)}`}>{formatDate(post.date)}</div>
        </div>
      </div>
    </Link>
  );
}
