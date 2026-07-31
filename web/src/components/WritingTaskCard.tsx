export function WritingTaskCard({
  badge,
  question,
  taskLabel,
  onSelect,
  loading,
}: {
  badge: string;
  question: string;
  taskLabel: string;
  onSelect: () => void;
  loading?: boolean;
}) {
  return (
    <button
      type="button"
      onClick={onSelect}
      disabled={loading}
      className="snap-start flex-shrink-0 w-64 md:w-72 text-left bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-5 hover:shadow-card-hover hover:-translate-y-1 transition-all duration-300 disabled:opacity-60 disabled:cursor-wait disabled:hover:-translate-y-0"
    >
      <div className="flex items-center gap-2 mb-3">
        <span className="text-[11px] font-semibold tracking-wide uppercase text-sage bg-sage-50 px-2.5 py-1 rounded-full">
          {taskLabel}
        </span>
        <span className="text-[11px] font-medium tracking-wide uppercase text-charcoal-light/60">{badge}</span>
      </div>
      <p className="font-serif text-[15px] text-charcoal leading-relaxed line-clamp-4">
        {loading ? "Loading..." : question}
      </p>
    </button>
  );
}
