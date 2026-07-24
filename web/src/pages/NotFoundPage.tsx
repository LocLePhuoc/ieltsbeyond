import { Link } from "react-router-dom";
import { useDocumentTitle } from "../lib/useDocumentTitle";

export default function NotFoundPage() {
  useDocumentTitle("Not Found - My Blog");

  return (
    <div className="flex flex-col items-center justify-center py-24 bg-white/60 backdrop-blur-sm rounded-3xl border border-white/60">
      <div className="text-7xl font-serif font-bold text-transparent bg-clip-text bg-gradient-to-r from-navy-300 to-sage mb-4">
        404
      </div>
      <p className="text-charcoal-light mb-6">The page you're looking for doesn't exist.</p>
      <Link
        to="/"
        className="inline-flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-navy-400 to-sage text-white rounded-2xl text-sm font-medium hover:shadow-glass transition-all duration-300"
      >
        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
          <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
        </svg>
        Back to Home
      </Link>
    </div>
  );
}
