import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

interface PracticeTask {
  type: string;
  category: string;
  question: string;
  imageKey?: string;
}

const defaultTask: PracticeTask = {
  type: "Task 2",
  category: "Opinion",
  question:
    "Some people think that the best way to increase road safety is to raise the minimum legal age for driving cars or riding motorbikes. To what extent do you agree or disagree?",
};

export default function PracticePage() {
  const location = useLocation();
  const task = (location.state as PracticeTask | null) ?? defaultTask;
  const [answer, setAnswer] = useState("");
  const [imageFailed, setImageFailed] = useState(false);

  useEffect(() => setImageFailed(false), [task.imageKey]);

  function handleCheckAnswer() {}

  return (
    <div className="flex flex-col gap-5">
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-6 md:p-8">
        <div className="flex items-center gap-2 mb-3">
          <span className="text-[11px] font-semibold tracking-wide uppercase text-sage bg-sage-50 px-2.5 py-1 rounded-full">
            {task.type}
          </span>
          <span className="text-[11px] font-medium tracking-wide uppercase text-charcoal-light/60">
            {task.category}
          </span>
        </div>
        <p className="font-serif text-lg md:text-xl text-charcoal leading-relaxed">{task.question}</p>
      </section>

      {task.type === "Task 1" && task.imageKey ? (
        <section className="rounded-2xl border border-white/60 bg-white/40 overflow-hidden">
          {imageFailed ? (
            <div className="flex flex-col items-center justify-center py-14 text-charcoal-light/50 text-sm">
              Image unavailable
            </div>
          ) : (
            <img
              src={task.imageKey}
              alt={`${task.category} chart for this Writing Task 1 question`}
              className="w-full max-h-[420px] object-contain bg-white"
              onError={() => setImageFailed(true)}
            />
          )}
        </section>
      ) : (
        <section className="min-h-[100px] rounded-2xl border border-dashed border-gray-300 bg-white/40" />
      )}

      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-2">
        <textarea
          value={answer}
          onChange={(e) => setAnswer(e.target.value)}
          placeholder="Start writing your response here..."
          className="w-full min-h-[360px] resize-y rounded-2xl bg-transparent p-4 text-[15px] text-charcoal leading-relaxed focus:outline-none"
        />
      </section>

      <div className="flex justify-end">
        <button
          type="button"
          onClick={handleCheckAnswer}
          className="px-6 py-3 rounded-xl bg-sage text-white text-sm font-semibold shadow-soft-lg hover:bg-sage-dark transition-colors"
        >
          Check Answer
        </button>
      </div>
    </div>
  );
}
