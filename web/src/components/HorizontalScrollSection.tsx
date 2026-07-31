import { useRef, type ReactNode } from "react";

export function HorizontalScrollSection({ title, children }: { title: string; children: ReactNode }) {
  const trackRef = useRef<HTMLDivElement>(null);

  function scrollByAmount(direction: 1 | -1) {
    const track = trackRef.current;
    if (!track) return;
    track.scrollBy({ left: direction * track.clientWidth * 0.8, behavior: "smooth" });
  }

  return (
    <section>
      <div className="flex items-center justify-between mb-3">
        <h2 className="font-serif text-2xl font-bold text-charcoal tracking-tight">{title}</h2>
        <div className="flex items-center gap-2">
          <button
            type="button"
            onClick={() => scrollByAmount(-1)}
            aria-label={`Scroll ${title} left`}
            className="w-9 h-9 rounded-full bg-white/80 backdrop-blur-xl shadow-soft border border-white/60 flex items-center justify-center text-charcoal hover:bg-sage hover:text-white hover:border-sage transition-colors"
          >
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
            </svg>
          </button>
          <button
            type="button"
            onClick={() => scrollByAmount(1)}
            aria-label={`Scroll ${title} right`}
            className="w-9 h-9 rounded-full bg-white/80 backdrop-blur-xl shadow-soft border border-white/60 flex items-center justify-center text-charcoal hover:bg-sage hover:text-white hover:border-sage transition-colors"
          >
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M8.25 4.5l7.5 7.5-7.5 7.5" />
            </svg>
          </button>
        </div>
      </div>
      <div ref={trackRef} className="hide-scrollbar flex gap-3.5 overflow-x-auto scroll-smooth snap-x snap-mandatory pb-1">
        {children}
      </div>
    </section>
  );
}
